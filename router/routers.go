package router

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var (
	infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLog  = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

// Router Root Router
func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// api swagger for develope mode
	//r.Get("/api/v1/swagger/*", httpSwagger.Handler(
	//	httpSwagger.URL(infrastructure.GetHTTPSwagger()),
	//	httpSwagger.DocExpansion("none"),
	//))

	//declare controller
	//accessController := controller.NewAccessController()
	//middlewareService := middlewares.NewMiddlewareService()
	//connectController := controller.NewConnectController()
	//syncController := controller.NewSyncController()
	//reportController := controller.NewReportController()
	//statisticController := controller.NewStatisticController()
	//userController := controller.NewUserController()
	//
	//advanceFilterV2Controller := controller.NewFilterController()
	//basicQueryV2Controller := controller.NewBasicQueryV2Controller()
	//fileController := controller.NewFileController()
	//
	//apiLogControl := controller.NewApiLogController()
	//logController := controller.NewLogController()
	//
	//scholasticController := controller.NewScholasticController()
	//
	//educationProgramController := controller.NewEducationProgramController()
	//outlineController := controller.NewOutlineController()
	//
	//abetProgramController := controller.NewABETProgramController()
	//abetOutlineController := controller.NewABETOutlineController()
	//
	//r.Route("/api/v1", func(router chi.Router) {
	//	// Public routes
	//	router.Post("/login", accessController.Login) // .With(middlewares.ValidatePayload(&controller.LoginPayload{}))
	//	router.With(middlewares.ValidatePayload(&controller.LoginGooglePayload{})).Post("/login/google", accessController.LoginWithGoogle)
	//	router.Post("/login/ldap", accessController.LoginLDAP)
	//	router.Post("/login/ldap-token", accessController.LoginLDAPWithToken)
	//	router.Post("/auth/reset-roles", accessController.ResetRoles)
	//	router.Post("/auth/reset-system", accessController.ResetSystem)
	//
	//	router.Group(func(externalRoute chi.Router) {
	//		externalRoute.Use(middlewares.AuthenticatorExternalCookieLDAP)
	//
	//		externalRoute.Route("/external", func(er chi.Router) {
	//			er.Get("/statistic", statisticController.GetStatistic)
	//			er.Get("/system", statisticController.GetSystem)
	//		})
	//
	//		externalRoute.Route("/connect", func(er chi.Router) {
	//			er.Get("/lay-cong-thuc-diem", connectController.LayCongThucDiem)
	//			er.Get("/lay-ds-hoc-phan", connectController.LayDSHocPhan)
	//			er.Get("/lay-stc-chuong-trinh", connectController.LaySTCChuongTrinh)
	//		})
	//	})
	//
	//	router.Group(func(protectedRoute chi.Router) {
	//		protectedRoute.Use(middlewares.AuthenticatorCookieLDAP)
	//		protectedRoute.Use(jwtauth.Verifier(infrastructure.GetEncodeAuth()))
	//		protectedRoute.Use(jwtauth.Authenticator)
	//		protectedRoute.Use(middlewareService.Authorizer())
	//		protectedRoute.Post("/token/refresh", accessController.Refresh)
	//		protectedRoute.Post("/logout", accessController.Logout)
	//		protectedRoute.Post("/auth/switch-role", accessController.SelectRole)
	//		protectedRoute.Post("/auth/apply-user-roles", accessController.ApplyUserRoles)
	//		protectedRoute.Post("/auth/reset-token", accessController.ResetToken)
	//		// protectedRoute.Post("/auth/reset-roles", accessController.ResetRoles)
	//
	//		protectedRoute.Route("/report", func(reportR chi.Router) {
	//			reportR.Get("/outline", reportController.ReportOutline)
	//			reportR.Get("/used-document", reportController.ReportUsedDocument)
	//		})
	//
	//		protectedRoute.Route("/log", func(logSubr chi.Router) {
	//			logSubr.Get("/filter", apiLogControl.FilterLog)
	//		})
	//
	//		protectedRoute.Route("/logs/log", func(logR chi.Router) {
	//			logR.With(middlewares.ValidatePayload(&controller.CreateLogPayload{})).Post("/", logController.Create)
	//			logR.With(middlewares.ValidatePayload(&model.Log{})).Put("/", logController.Update)
	//			logR.Delete("/{id}", logController.Delete)
	//		})
	//
	//		// =============================================================================
	//		protectedRoute.Route("/users/user", func(userR chi.Router) {
	//			userR.With(middlewares.ValidatePayload(&controller.ListUserPayload{})).Post("/", userController.Create)
	//			userR.With(middlewares.ValidatePayload(&model.User{})).Put("/", userController.Update)
	//			userR.With(middlewares.ValidatePayload(&controller.UserPasswordPayload{})).Put("/password", userController.UpdatePassword)
	//			userR.Delete("/", userController.Delete)
	//			userR.Put("/reset", userController.ResetPassword)
	//		})
	//
	//		// =============================================================================
	//		protectedRoute.Route("/scholastics/scholastic", func(scholasticR chi.Router) {
	//			scholasticR.Post("/copy", scholasticController.CopyData)
	//		})
	//
	//		// CDIO =============================================================================
	//		protectedRoute.Route("/cdio-program", func(eduProgramR chi.Router) {
	//			eduProgramR.Post("/copy", educationProgramController.CopyVersion)
	//			eduProgramR.Post("/submit", educationProgramController.SubmitProgram)
	//		})
	//		protectedRoute.Route("/cdio-outline", func(outlineR chi.Router) {
	//			outlineR.Post("/copy", outlineController.CopyVersion)
	//			outlineR.Post("/submit", outlineController.SubmitOutline)
	//		})
	//
	//		// ABET =============================================================================
	//		protectedRoute.Route("/abet-program", func(abetProgR chi.Router) {
	//			abetProgR.Post("/copy", abetProgramController.Copy)
	//			abetProgR.Post("/submit", abetProgramController.SubmitABETProgram)
	//		})
	//		protectedRoute.Route("/abet-outline", func(abetOutlineR chi.Router) {
	//			abetOutlineR.Post("/copy", abetOutlineController.Copy)
	//			abetOutlineR.Post("/submit", abetOutlineController.SubmitABETOutline)
	//		})
	//
	//		// =============================================================================
	//		/* Filter & Basic Query */
	//		protectedRoute.Route("/filters", func(filterR chi.Router) {
	//			filterR.Use(middlewareService.AdvanceFilterAuthorizerV2())
	//			filterR.Post("/filter/", advanceFilterV2Controller.AdvanceFilter)
	//		})
	//
	//		protectedRoute.Route("/basicQueries/basicQueryV2", func(basicQueryR chi.Router) {
	//			basicQueryR.Use(middlewareService.BasicQueryAuthorizerV2())
	//			basicQueryR.Put("/delete", basicQueryV2Controller.Delete)
	//			basicQueryR.Post("/", basicQueryV2Controller.Insert)
	//			basicQueryR.Put("/", basicQueryV2Controller.Update)
	//		})
	//
	//		// ----fileController ----
	//		protectedRoute.Route("/file", func(subR chi.Router) {
	//			subR.Post("/storage", fileController.UploadFile)
	//			subR.Delete("/storage", fileController.DeleteFile)
	//		})
	//
	//		// ---syncController ---
	//		protectedRoute.Route("/sync", func(syncR chi.Router) {
	//			syncR.Get("/get-sync-status", syncController.GetSyncStatus)
	//			syncR.Get("/hrm", syncController.PullHRM)
	//			syncR.Get("/document/koha", syncController.FetchDataFromLibraryKOHA)
	//			syncR.Get("/document/dspace", syncController.FetchDataFromLibraryDSPACE)
	//			syncR.Get("/document-num-of-items", syncController.FetchNumOfItemFromLibrary)
	//		})
	//	})
	//	// Protected routes
	//	// Create serve files api
	//	router.Group(func(protectedRoute chi.Router) {
	//		// Middleware authentication
	//		protectedRoute.Use(middlewares.AuthenticatorCookieLDAP)
	//		protectedRoute.Use(jwtauth.Verifier(infrastructure.GetEncodeAuth()))
	//		protectedRoute.Use(jwtauth.Authenticator)
	//
	//		fs := http.StripPrefix("/api/v1/storage", http.FileServer(http.Dir(infrastructure.GetStoragePath())))
	//		protectedRoute.Get("/storage/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//			fs.ServeHTTP(w, r)
	//		}))
	//	})
	//
	//	router.Group(func(er chi.Router) {
	//		fs := http.StripPrefix("/api/v1/storagePublic", http.FileServer(http.Dir(infrastructure.GetStoragePublicPath())))
	//		er.Get("/storagePublic/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//			fs.ServeHTTP(w, r)
	//		}))
	//	})
	//})
	//
	return r
}
