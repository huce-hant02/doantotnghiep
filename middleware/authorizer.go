package middlewares

//
//import (
//	"doantotnghiep/infrastructure"
//	"doantotnghiep/model"
//	"bytes"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"io"
//	"io/ioutil"
//	"log"
//	"net/http"
//
//	"regexp"
//	"strings"
//	"time"
//)
//
//// Authorizer authorize middleware
//// func Authorizer(e *casbin.Enforcer) func(next http.Handler) http.Handler {
//// 	return func(next http.Handler) http.Handler {
//// 		fn := func(w http.ResponseWriter, r *http.Request) {
//// 			_, claims, _ := jwtauth.FromContext(r.Context())
//func (s *middlewareService) Authorizer() func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		fn := func(w http.ResponseWriter, r *http.Request) {
//			var apiCallLog model.APICallLog
//			t := time.Now()
//			apiCallLog.CallTime = &t
//			// _, claims, _ := jwtauth.FromContext(r.Context())
//			url := r.URL.Path
//			method := r.Method
//			if method == "" {
//				method = "GET"
//			}
//			apiCallLog.URL = &url
//			apiCallLog.Method = &method
//			apiCallLog.Authorized = false
//			role, claims, err := s.GetRole(r)
//			if err != nil {
//				// infrastructure.InfoLog.Println(err)
//				apiCallLog.Error = err.Error()
//				if errSaveLog := s.apiCallLogService.SaveLog(&apiCallLog); errSaveLog != nil {
//					w.WriteHeader(http.StatusInternalServerError)
//					http.Error(w, errSaveLog.Error(), http.StatusInternalServerError)
//					return
//				}
//				infrastructure.InfoLog.Println(err)
//				w.WriteHeader(http.StatusUnauthorized)
//				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//				return
//			}
//
//			// username, ok := claims["username"].(string)
//			// if !ok {
//			// 	err := errors.New("couldn't parse username")
//			// 	apiCallLog.Error = err.Error()
//			// 	if errSaveLog := s.apiCallLogService.SaveLog(&apiCallLog); errSaveLog != nil {
//			// 		w.WriteHeader(http.StatusInternalServerError)
//			// 		http.Error(w, errSaveLog.Error(), http.StatusInternalServerError)
//			// 		return
//			// 	}
//			// 	infrastructure.InfoLog.Println(err)
//			// 	w.WriteHeader(http.StatusUnauthorized)
//			// 	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//			// 	return
//			// }
//			userId, ok := claims["user_id"].(float64)
//			if !ok {
//				err := errors.New("couldn't parse user_id")
//				apiCallLog.Error = err.Error()
//				if errSaveLog := s.apiCallLogService.SaveLog(&apiCallLog); errSaveLog != nil {
//					w.WriteHeader(http.StatusInternalServerError)
//					http.Error(w, errSaveLog.Error(), http.StatusInternalServerError)
//					return
//				}
//				infrastructure.InfoLog.Println(err)
//				w.WriteHeader(http.StatusUnauthorized)
//				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//				return
//			}
//
//			active, err := s.accessService.GetSystemStatus()
//			if err != nil || active == nil || *active == false {
//				w.WriteHeader(http.StatusUnprocessableEntity)
//				http.Error(w, http.StatusText(422), 422)
//				return
//			}
//
//			// Get User Permission ===============================================================================
//			permission, err := s.accessService.GetUserPermission(uint(userId), role)
//			if err != nil {
//				w.WriteHeader(http.StatusUnprocessableEntity)
//				http.Error(w, http.StatusText(422), 422)
//				return
//			}
//			if permission.UserActive == nil || *permission.UserActive != true {
//				w.WriteHeader(http.StatusUnauthorized)
//				http.Error(w, http.StatusText(401), 401)
//				return
//			}
//			// =========================================================================
//			roles := []string{}
//			for _, r := range permission.Roles {
//				roles = append(roles, r.Code)
//			}
//
//			if !utils.Contains(roles, role) {
//				err := errors.New("role '" + role + "' is not in this API required scope")
//				w.WriteHeader(403)
//				http.Error(w, err.Error(), 403)
//				return
//			}
//			uid := uint(userId)
//			apiCallLog.UserId = &uid
//			apiCallLog.Role = &role
//			queryValues := r.URL.Query()
//			queryParam := make(model.ListParamKeyValue, 0)
//			for k, v := range queryValues {
//				queryParam = append(queryParam, model.ParamKeyValue{
//					ParamKey:   k,
//					ParamValue: v,
//				})
//			}
//			if len(queryParam) > 0 {
//				apiCallLog.QueryParameters = queryParam
//			}
//
//			regex, _ := regexp.Compile("^.*login.*$")
//			if !regex.MatchString(r.URL.Path) {
//				if r.Method == "POST" || r.Method == "PUT" || r.Method == "DELETE" {
//					// Read the body content
//					var bodyBytes []byte
//					if r.Body != nil {
//						bodyBytes, _ = ioutil.ReadAll(r.Body)
//					}
//					// Restore the io.ReadCloser to its original state
//					r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
//					// Use the content
//					// bodyString := string(bodyBytes)
//					// paramString := r.URL.Query()
//					// param, err := json.Marshal(paramString)
//					// regexNum := regexp.MustCompile("[0-9]+")
//					// trimPrefixURL := strings.TrimPrefix(url, "/api/v1/")
//					// trimmedHasNum := regexNum.MatchString(trimPrefixURL)
//					// if trimmedHasNum {
//					// 	url = trimPrefixURL
//					// 	replacedURL := regexNum.ReplaceAll([]byte(url), []byte("*"))
//					// 	url = "/api/v1/" + string(replacedURL)
//					// }
//					mapAPI := []string{}
//					for _, item := range permission.APIs {
//						mapAPI = append(mapAPI, strings.Join([]string{item.Method, item.Url}, "::"))
//					}
//
//					if !utils.Contains(mapAPI, strings.Join([]string{method, url}, "::")) {
//						w.WriteHeader(http.StatusForbidden)
//						http.Error(w, err.Error(), http.StatusForbidden)
//						return
//					} else {
//						goto SERVE
//					}
//					//infrastructure.ErrLog.Println(bodyString)
//					// userID, ok := claims["user_id"].(float64)
//					// if !ok {
//					// 	w.WriteHeader(http.StatusInternalServerError)
//					// 	http.Error(w, http.StatusText(500), 500)
//					// 	return
//					// }
//					// logRepo := repository.NewApiLogRepository()
//					// err = logRepo.Save(model.ApiLog{
//					// 	UserID: int(userID),
//					// 	URL:    r.URL.Path,
//					// 	Method: r.Method,
//					// 	Param:  string(param),
//					// 	Data:   bodyString,
//					// 	Time:   time.Now(),
//					// })
//					// if err != nil {
//					// 	infrastructure.ErrLog.Println("ERROR : ", err)
//					// }
//				}
//			}
//		SERVE:
//			apiCallLog.Authorized = true
//			next.ServeHTTP(w, r)
//			// if ignore := IsRouteIgnored(url); !ignore {
//			// 	if err := s.apiCallLogService.SaveLog(&apiCallLog); err != nil {
//			// 		infrastructure.ErrLog.Println(err)
//			// 	}
//			// }
//		}
//		return http.HandlerFunc(fn)
//	}
//}
//
//func (s *middlewareService) BasicQueryAuthorizerV2() func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		fn := func(w http.ResponseWriter, r *http.Request) {
//			role, _, err := s.GetRole(r)
//			if err != nil {
//				w.WriteHeader(401)
//				http.Error(w, err.Error(), 401)
//				return
//			}
//			var payload controller.BasicQueryPayload
//			body, err := io.ReadAll(r.Body)
//			if err != nil {
//				w.WriteHeader(http.StatusUnprocessableEntity)
//				http.Error(w, http.StatusText(422), 422)
//				return
//			}
//			r.Body = io.NopCloser(bytes.NewBuffer(body))
//			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
//				log.Println("here")
//				w.WriteHeader(401)
//				http.Error(w, err.Error(), 401)
//				return
//			}
//			modelType := payload.ModelType
//
//			// Get User Permission ===============================================================================
//			role, claims, err := s.GetRole(r)
//			if err != nil {
//				fmt.Println("dm1", err)
//				w.WriteHeader(401)
//				http.Error(w, err.Error(), 401)
//				return
//			}
//			// username, ok := claims["username"].(string)
//			// if !ok {
//			// 	w.WriteHeader(http.StatusUnprocessableEntity)
//			// 	http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
//			// 	return
//			// }
//			userId, ok := claims["user_id"].(float64)
//			if !ok {
//				w.WriteHeader(http.StatusUnprocessableEntity)
//				http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
//				return
//			}
//			permission, err := s.accessService.GetUserPermission(uint(userId), role)
//			if err != nil {
//				w.WriteHeader(http.StatusUnprocessableEntity)
//				http.Error(w, http.StatusText(422), 422)
//				return
//			}
//
//			// Query  ===============================================================================
//			switch r.Method {
//			case "POST":
//				// is insert
//				if !permission.ModelTypes[modelType].Create {
//					w.WriteHeader(401)
//					http.Error(w, err.Error(), 401)
//					return
//				}
//			case "PUT":
//				path := r.URL.Path
//				if strings.Contains(path, "delete") {
//					queryValues := r.URL.Query()
//					modelType := queryValues.Get("modelType")
//					// is delete
//					if !permission.ModelTypes[modelType].Delete {
//						w.WriteHeader(401)
//						http.Error(w, err.Error(), 401)
//						return
//					}
//				} else {
//					if !permission.ModelTypes[modelType].Update {
//						w.WriteHeader(401)
//						http.Error(w, err.Error(), 401)
//						return
//					}
//				}
//			}
//			r.Body = io.NopCloser(bytes.NewBuffer(body))
//			next.ServeHTTP(w, r)
//		}
//		return http.HandlerFunc(fn)
//	}
//}
//
//func (s *middlewareService) AdvanceFilterAuthorizerV2() func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		fn := func(w http.ResponseWriter, r *http.Request) {
//			// get role from the claims
//			role, claims, err := s.GetRole(r)
//			if err != nil {
//				fmt.Println("dm1", err)
//				w.WriteHeader(401)
//				http.Error(w, err.Error(), 401)
//				return
//			}
//			userID, ok := claims["user_id"].(float64)
//			if !ok {
//				w.WriteHeader(http.StatusUnprocessableEntity)
//				http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
//				return
//			}
//
//			// check if the user has read access to this table
//			var payload controller.FilterGeneralV2Payload
//			body, err := io.ReadAll(r.Body)
//			if err != nil {
//				w.WriteHeader(http.StatusUnprocessableEntity)
//				http.Error(w, http.StatusText(422), 422)
//				return
//			}
//			r.Body = io.NopCloser(bytes.NewBuffer(body))
//			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
//				fmt.Println("dm2", err)
//				w.WriteHeader(401)
//				http.Error(w, err.Error(), 401)
//				return
//			}
//
//			permission, err := s.accessService.GetUserPermission(uint(userID), role)
//			if err != nil {
//				w.WriteHeader(http.StatusUnprocessableEntity)
//				http.Error(w, http.StatusText(422), 422)
//				return
//			}
//
//			if permission.ModelTypes[payload.ModelType].Read != true {
//				w.WriteHeader(403)
//				http.Error(w, http.StatusText(403), 403)
//				return
//			} else {
//				r.Body = io.NopCloser(bytes.NewBuffer(body))
//				next.ServeHTTP(w, r)
//			}
//		}
//		return http.HandlerFunc(fn)
//	}
//}
//
//func IsRouteIgnored(url string) bool {
//	for _, ignored := range model.APILoggerIgnoreRoutes {
//		if strings.Contains(url, ignored) {
//			return true
//		}
//	}
//	return false
//}
