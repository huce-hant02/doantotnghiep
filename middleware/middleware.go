package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"q-cdio/controller"
	"q-cdio/infrastructure"
	"q-cdio/model"
	"q-cdio/repository"
	"q-cdio/service"
	"q-cdio/utils"
	"reflect"
	"strings"
	"time"

	"github.com/go-chi/render"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

type middlewareService struct {
	accessService     service.AccessService
	apiCallLogService service.APICallLogService
	db                *gorm.DB
	basicQueryRepo    repository.BasicQueryV2Repository
}

type MiddlewareService interface {
	// DeveloperAuthenticator(next http.Handler) http.Handler
	GetRole(r *http.Request) (string, map[string]interface{}, error)
	Authorizer() func(next http.Handler) http.Handler
	BasicQueryAuthorizerV2() func(next http.Handler) http.Handler
	AdvanceFilterAuthorizerV2() func(next http.Handler) http.Handler
}

func validate(w http.ResponseWriter, payload interface{}) error {
	v := validator.New()

	errs := v.Struct(payload)

	if errs != nil {
		validationErrorResponse(w, errs)
		return errs
	}
	return nil
}

func ValidatePayload(payload interface{}) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			payloadModel := reflect.ValueOf(payload)
			datas := reflect.New(payloadModel.Type().Elem())
			// if payloadModelType.Kind() == reflect.Ptr {
			// 	payloadModelType = payloadModelType.Elem()
			// }
			// datas := reflect.New(payloadModelType)
			// log.Println(payloadModelType)
			// initializeStruct(payloadModelType, datas.Elem())
			err := json.NewDecoder(r.Body).Decode(datas.Interface())

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				http.Error(w, http.StatusText(400), 400)
				res := &controller.Response{
					Data:    nil,
					Message: "Bad request: " + err.Error(),
					Success: false,
				}
				render.JSON(w, r, res)
				return
			}

			defer r.Body.Close()
			err = validate(w, datas.Interface())
			if err != nil {
				return
			}

			ctx := context.WithValue(r.Context(), "payload", datas.Interface())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}

func validationErrorResponse(w http.ResponseWriter, err error) {
	errResponse := make([]string, 0)

	for _, e := range err.(validator.ValidationErrors) {
		errResponse = append(errResponse, fmt.Sprint(e))
	}

	response := map[string][]string{"errors": errResponse}
	render.JSON(w, nil, response)
}

func AuthenticatorCookieLDAP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		domain := ""
		full_domain := r.Header.Get("Origin")
		if strings.Contains(full_domain, "localhost") {
			domain = "localhost"
		}
		if strings.Contains(full_domain, "piditi.com") {
			domain = ".piditi.com"
		}
		if strings.Contains(full_domain, "phenikaa-uni.edu.vn") {
			domain = ".phenikaa-uni.edu.vn"
		}
		if r.RequestURI == "/api/v1/logout" {
			var accessToken string
			cookie_access, errCookie := r.Cookie("AccessToken")
			if errCookie != nil {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, http.StatusText(401), 401)
				return
			}

			accessToken = cookie_access.Value
			r.Header.Set("Authorization", "Bearer "+accessToken)
			c_access := http.Cookie{
				Name:   "AccessToken",
				Domain: full_domain,
				Path:   "/",
				Value:  "",
				// HttpOnly: true,
				// Secure:   true,
				MaxAge: -1,
			}
			c_refresh := http.Cookie{
				Name:   "RefreshToken",
				Domain: full_domain,
				Path:   "/",
				Value:  "",
				// HttpOnly: true,
				// Secure:   true,
				MaxAge: -1,
			}
			c_ldap := http.Cookie{
				Name:     "LDAPToken",
				Domain:   domain,
				Path:     "/",
				Value:    "",
				HttpOnly: true,
				// Secure:   true,
				MaxAge: -1,
			}
			http.SetCookie(w, &c_access)
			http.SetCookie(w, &c_refresh)
			http.SetCookie(w, &c_ldap)
		} else {
			accessCookie, errAccessCookie := r.Cookie("AccessToken")
			if accessCookie == nil || errAccessCookie != nil {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, http.StatusText(401), 401)
				return
			}
			accessToken := accessCookie.Value
			accessClaims, err := controller.GetAndDecodeToken("Bearer " + accessToken)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, http.StatusText(401), 401)
				return
			}
			utype, ok := accessClaims["utype"].(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, http.StatusText(401), 401)
				return
			}

			if domain != ".phenikaa-uni.edu.vn" || !utils.Contains([]string{"NHANSU", "PHONGBAN"}, utype) {
				r.Header.Set("Authorization", "Bearer "+accessToken)
				// fmt.Println("set token ok")
			} else {
				var ldapToken string
				ldap_cookie, errCookie := r.Cookie("LDAPToken")
				if errCookie != nil {
					infrastructure.InfoLog.Println("here", errCookie)
					w.WriteHeader(http.StatusUnauthorized)
					http.Error(w, http.StatusText(401), 401)
					return
				}
				ldapToken = ldap_cookie.Value

				accessService := service.NewAccessService()
				if errCookie != nil {
					ldap_token, err := infrastructure.GetDecodeAuth().Decode(ldapToken)
					if err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						http.Error(w, err.Error(), http.StatusUnauthorized)
						return
					}
					if time.Since(ldap_token.Expiration()).Milliseconds() < 0 {
						// Check token valid
						if _, err := ldap_token.AsMap(context.Background()); err != nil {
							infrastructure.InfoLog.Println("here")
							w.WriteHeader(http.StatusUnauthorized)
							http.Error(w, err.Error(), http.StatusUnauthorized)
							return
						}

						// Get the uuid
						claims, err := ldap_token.AsMap(context.Background())
						if err != nil {
							infrastructure.InfoLog.Println("here")
							w.WriteHeader(http.StatusUnauthorized)
							http.Error(w, err.Error(), http.StatusUnauthorized)
							return
						}
						ldapUUID, ok := claims["ldap_uuid"].(string)
						if !ok {
							w.WriteHeader(http.StatusUnprocessableEntity)
							http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
							return
						}
						role, ok := claims["role"].(string)
						if !ok {
							w.WriteHeader(http.StatusUnauthorized)
							http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
							return
						}
						email, ok := claims["email"].(string)
						if !ok {
							w.WriteHeader(http.StatusUnprocessableEntity)
							http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
							return
						}
						code, ok := claims["code"].(string)
						if !ok {
							w.WriteHeader(http.StatusUnprocessableEntity)
							http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
							return
						}
						username, ok := claims["ldap_username"].(string)
						if !ok {
							w.WriteHeader(http.StatusUnprocessableEntity)
							http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
							return
						}
						utype, ok := claims["utype"].(string)
						if !ok {
							w.WriteHeader(http.StatusUnprocessableEntity)
							http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
							return
						}
						password, ok := claims["ldap_password"].(string)
						if !ok {
							w.WriteHeader(http.StatusUnprocessableEntity)
							http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
							return
						}
						passwordDecrypt, err := infrastructure.RsaDecrypt(password)
						if err != nil {
							w.WriteHeader(http.StatusBadRequest)
							http.Error(w, http.StatusText(400), 400)
							return
						}
						// Delete the previous Refresh Token
						_, err = accessService.DeleteAuth(ldapUUID)
						if err != nil {
							infrastructure.InfoLog.Println("here")
							w.WriteHeader(http.StatusUnauthorized)
							http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
							return
						}
						isPassLDAP, _, err := accessService.TryLDAP(username, string(password))
						if err != nil {
							infrastructure.InfoLog.Println("here")
							w.WriteHeader(http.StatusInternalServerError)
							http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
							return
						}
						if !isPassLDAP {
							infrastructure.InfoLog.Println("here")
							w.WriteHeader(http.StatusUnauthorized)
							http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
							return
						}

						LDAPToken, err := accessService.CreateLDAPToken(username, string(passwordDecrypt), utype, code, email, role)
						if err != nil {
							infrastructure.InfoLog.Println("here")
							w.WriteHeader(http.StatusUnauthorized)
							http.Error(w, err.Error(), http.StatusUnauthorized)
							return
						}
						userService := service.NewUserService()
						employee, err := userService.GetEmployeeByEmail(email)
						if err != nil {
							infrastructure.InfoLog.Println("here")
							w.WriteHeader(http.StatusUnauthorized)
							http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
							return
						}
						tokenDetail, err := accessService.CreateToken(employee.User.ID, employee.User.Role)
						if err != nil {
							w.WriteHeader(http.StatusUnprocessableEntity)
							http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
							return
						}

						if saveErr := accessService.CreateAuth(employee.User.ID, tokenDetail); saveErr != nil {
							w.WriteHeader(http.StatusUnprocessableEntity)
							http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
							return
						}
						c_access := http.Cookie{
							Name:   "AccessToken",
							Domain: full_domain,
							Path:   "/",
							Value:  tokenDetail.AccessToken,
							// HttpOnly: true,
							// Secure:   true,
							Expires: time.Now().Add(time.Hour * time.Duration(model.AccessTokenTime)),
						}
						c_refresh := http.Cookie{
							Name:   "RefreshToken",
							Domain: full_domain,
							Path:   "/",
							Value:  tokenDetail.RefreshToken,
							// HttpOnly: true,
							// Secure:   true,
							Expires: time.Now().Add(time.Hour * time.Duration(model.RefreshTokenTime)),
						}
						ldap_cookie := http.Cookie{
							Name:     "LDAPToken",
							Domain:   domain,
							Path:     "/",
							Value:    LDAPToken,
							HttpOnly: true,
							// Secure:   true,
							Expires: time.Now().Add(time.Hour * time.Duration(model.LDAPTokenTime)),
						}
						http.SetCookie(w, &ldap_cookie)
						http.SetCookie(w, &c_access)
						http.SetCookie(w, &c_refresh)
						r.Header.Set("Authorization", "Bearer "+tokenDetail.AccessToken)
					}
				} else {
					r.Header.Set("Authorization", "Bearer "+accessToken)
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

func AuthenticatorExternalCookieLDAP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fullDomain := r.Header.Get("Origin")
		fullDomain := r.Host
		domain := ""
		if strings.Contains(fullDomain, "localhost") || strings.Contains(fullDomain, "[::1]") {
			domain = "localhost"
		}
		if strings.Contains(fullDomain, "piditi.com") {
			domain = ".piditi.com"
		}
		if strings.Contains(fullDomain, "phenikaa-uni.edu.vn") {
			domain = ".phenikaa-uni.edu.vn"
		}
		if domain != "" {
			goto SERVE
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

	SERVE:
		next.ServeHTTP(w, r)
	})
}

func GetUserId(r *http.Request) (uint, error) {
	accessCookie, errAccessCookie := r.Cookie("AccessToken")
	if accessCookie == nil || errAccessCookie != nil {
		return 0, errors.New("unauthorized")
	}

	accessToken := accessCookie.Value
	accessClaims, err := controller.GetAndDecodeToken("Bearer " + accessToken)
	if err != nil {
		return 0, errors.New("unauthorized")
	}
	userId, ok := accessClaims["user_id"].(float64)
	if !ok {
		return 0, errors.New("unauthorized")
	}
	return uint(userId), nil
}

func (s *middlewareService) GetRole(r *http.Request) (string, map[string]interface{}, error) {
	accessCookie, errAccessCookie := r.Cookie("AccessToken")
	if accessCookie == nil {
		return "guest", nil, nil
	}
	if errAccessCookie != nil {
		return "guest", nil, errors.New("unauthorized")
	}
	accessToken := accessCookie.Value
	accessClaims, err := controller.GetAndDecodeToken("Bearer " + accessToken)
	if err != nil {
		return "guest", nil, errors.New("unauthorized")
	}
	role, ok := accessClaims["role"].(string)
	if !ok {
		return "guest", accessClaims, errors.New("unauthorized")
	}

	return role, accessClaims, nil
}

func NewMiddlewareService() MiddlewareService {
	apiCallLogService := service.NewAPICallLogService()
	db := infrastructure.GetDB()
	basicQueryRepo := repository.NewBasicQueryV2Repo()
	accessService := service.NewAccessService()
	return &middlewareService{
		apiCallLogService: apiCallLogService,
		db:                db,
		basicQueryRepo:    basicQueryRepo,
		accessService:     accessService,
	}
}
