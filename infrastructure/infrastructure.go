package infrastructure

import (
	"flag"
	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

const (
	ENV        = "ENV"
	APPPORT    = "APP_PORT"
	DBHOST     = "DB_HOST"
	DBPORT     = "DB_PORT"
	DBUSER     = "DB_USER"
	DBPASSWORD = "DB_PASSWORD"
	DBNAME     = "DB_NAME"

	HTTPSWAGGER = "HTTP_SWAGGER"
	ROOTPATH    = "ROOT_PATH"

	PRIVATEPASSWORD = "PRIVATE_PASSWORD"
	PRIVATEPATH     = "PRIVATE_PATH"
	PUBLICPATH      = "PUBLIC_PATH"

	REDISURL = "REDIS_URL"
	BASEURL  = "BASE_URL"

	EXTENDHOUR         = "EXTEND_ACCESS_HOUR"
	EXTENDACCESSMINUTE = "EXTEND_ACCESS_MINUTE"
	EXTENDREFRESHHOUR  = "EXTEND_REFRESH_HOUR"

	KEYMATCHMODEL = "KEY_MATCH_MODEL"

	MAILSERVER  = "MAIL_SERVER"
	MAILPORT    = "MAIL_PORT"
	MAILACCOUNT = "MAIL_ACCOUNT"
	MAILPASS    = "MAIL_PASS"

	RATETERGETDIRECTOFFER = 0.1
	RATETARGETSTUDYRECORD = 0.4
	RATETARGETEXAM        = 0.5
)

var (
	env        string
	appPort    string
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string

	httpSwagger       string
	rootPath          string
	storagePath       string
	storagePublicPath string
	storageAvatarPath string
	storageFilePath   string

	InfoLog *log.Logger
	ErrLog  *log.Logger

	db         *gorm.DB
	encodeAuth *jwtauth.JWTAuth
	decodeAuth *jwtauth.JWTAuth
	//privateKey *rsa.PrivateKey
	publicKey interface{}

	baseURL string

	redisURL    string
	redisClient *redis.Client
	//enforcer    *casbin.Enforcer

	privatePassword    string
	privatePath        string
	extendAccessMinute int
	extendRefreshHour  int

	publicPath string

	extendHour int

	keyMatchModel string

	NameRefreshTokenInCookie string
	NameAccessTokenInCookie  string

	mailServer   string
	mailPort     string
	mailAccount  string
	mailPassword string

	documentFetchingChan    chan string
	documentFetchingStatus  bool
	documentFetchingErrChan chan error
)

func getStringEnvParameter(envParam string, defaultValue string) string {
	if value, ok := os.LookupEnv(envParam); ok {
		return value
	}
	return defaultValue
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func loadEnvParameters(version int, dbNameArg string, dbPwdArg string) {
	root, _ := os.Getwd()
	env = getStringEnvParameter(ENV, goDotEnvVariable(ENV))
	appPort = getStringEnvParameter(APPPORT, goDotEnvVariable(APPPORT))

	dbPort = getStringEnvParameter(DBPORT, goDotEnvVariable(DBPORT))
	switch version {
	case 0:
		dbHost = getStringEnvParameter(DBHOST, "localhost")
		dbUser = getStringEnvParameter(DBUSER, "postgres")
		dbPassword = getStringEnvParameter(DBPASSWORD, dbPwdArg)
		dbName = getStringEnvParameter(DBNAME, dbNameArg)
		log.Println("Enviroment: LOCALHOST")
		break

	default:
		dbHost = getStringEnvParameter(DBHOST, goDotEnvVariable(DBHOST))
		dbUser = getStringEnvParameter(DBUSER, goDotEnvVariable(DBUSER))
		dbPassword = getStringEnvParameter(DBPASSWORD, goDotEnvVariable(DBPASSWORD))
		dbName = getStringEnvParameter(DBNAME, goDotEnvVariable(DBNAME))

		// dbHost = getStringEnvParameter(DBHOST, "159.65.143.187")
		// dbUser = getStringEnvParameter(DBUSER, "cdio_user")
		// dbName = getStringEnvParameter(DBNAME, "cdio")
		// dbPassword = getStringEnvParameter(DBPASSWORD, "s3#fj@dAnU")
		log.Println("Enviroment: Development Default")
	}

	privatePassword = getStringEnvParameter(PRIVATEPASSWORD, "Nhuanhthu1")
	privatePath = getStringEnvParameter(PRIVATEPATH, root+"/infrastructure/private.pem")
	publicPath = getStringEnvParameter(PUBLICPATH, root+"/infrastructure/public.pem")

	extendHour, _ = strconv.Atoi(getStringEnvParameter(EXTENDHOUR, "720"))
	extendAccessMinute, _ = strconv.Atoi(getStringEnvParameter(EXTENDACCESSMINUTE, goDotEnvVariable(EXTENDACCESSMINUTE)))
	extendRefreshHour, _ = strconv.Atoi(getStringEnvParameter(EXTENDREFRESHHOUR, goDotEnvVariable(EXTENDREFRESHHOUR)))

	keyMatchModel = getStringEnvParameter(KEYMATCHMODEL, root+"/infrastructure/keymatch_model.conf")

	httpSwagger = getStringEnvParameter(HTTPSWAGGER, goDotEnvVariable(HTTPSWAGGER))

	baseURL = getStringEnvParameter(BASEURL, goDotEnvVariable("BASE_URL"))
	redisURL = getStringEnvParameter(REDISURL, goDotEnvVariable("REDIS_URL"))

	rootPath = getStringEnvParameter(goDotEnvVariable(ROOTPATH), root)
	storagePath = rootPath + string(os.PathSeparator) + "storage"
	storagePublicPath = rootPath + string(os.PathSeparator) + "storagePublic"
	storageAvatarPath = storagePath + string(os.PathSeparator) + "avatar"
	storageFilePath = storagePath + string(os.PathSeparator) + "files"

	NameRefreshTokenInCookie = "refreshTokenEP"
	NameAccessTokenInCookie = "accessTokenEP"

	mailServer = getStringEnvParameter(MAILSERVER, goDotEnvVariable(MAILSERVER))
	mailPort = getStringEnvParameter(MAILPORT, goDotEnvVariable(MAILPORT))
	mailAccount = getStringEnvParameter(MAILACCOUNT, goDotEnvVariable(MAILACCOUNT))
	mailPassword = getStringEnvParameter(MAILPASS, goDotEnvVariable(MAILPASS))
}

func init() {
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	ErrLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Get version ARGS
	var version int
	flag.IntVar(&version, "v", 1, "select version dev v1 or dev v2")
	// flag.Parse()
	var dbNameArg string
	flag.StringVar(&dbNameArg, "dbname", "postgres", "database name need to connect")

	var dbPwdArg string
	flag.StringVar(&dbPwdArg, "dbpwd", "123456", "password in database need to connect")

	var initDB bool
	flag.BoolVar(&initDB, "db", false, "allow recreate model database in postgres")

	flag.Parse()
	log.Println("database version: ", version)

	loadEnvParameters(version, dbNameArg, dbPwdArg)
	//if err := loadAuthToken(); err != nil {
	//	log.Println("error load auth token: ", err)
	//}

	if err := InitRedis(); err != nil {
		log.Fatal("error initialize redis: ", err)
	}

	if err := InitDatabase(initDB); err != nil {
		ErrLog.Println("error initialize database: ", err)
	}

	//if err := InitAuthorization(); err != nil {
	//	log.Fatal(err)
	//}

	documentFetchingChan = make(chan string)
	documentFetchingErrChan = make(chan error)
	documentFetchingStatus = false
}

func GetDBName() string {
	return dbName
}

// GetDB export db
func GetDB() *gorm.DB {
	return db
}

// GetHTTPSwagger export link swagger
func GetHTTPSwagger() string {
	return httpSwagger
}

// GetAppPort export app port
func GetAppPort() string {
	return appPort
}

// GetStoragePath export storage path
func GetStoragePath() string {
	return storagePath
}

// GetStoragePublicPath export storage path
func GetStoragePublicPath() string {
	return storagePublicPath
}

// GetEncodeAuth get token auth
func GetEncodeAuth() *jwtauth.JWTAuth {
	return encodeAuth
}

// GetDecodeAuth export decode auth
func GetDecodeAuth() *jwtauth.JWTAuth {
	return decodeAuth
}

// GetExtendAccessMinute export access extend minute
func GetExtendAccessHour() int {
	return extendHour
}

// GetExtendAccessMinute export access extend minute
func GetExtendAccessMinute() int {
	return extendAccessMinute
}

// GetExtendRefreshHour export refresh extends hour
func GetExtendRefreshHour() int {
	return extendRefreshHour
}

// GetKeyMatchModel get key match model path
func GetKeyMatchModel() string {
	return keyMatchModel
}

// GetEnforcer export enforcer
//func GetEnforcer() *casbin.Enforcer {
//	return enforcer
//}

// GetMailParam
func GetMailParam() (string, string, string, string) {
	return mailServer, mailPort, mailAccount, mailPassword
}

// GetRedisClient export redis client
//func GetRedisClient() *redis.Client {
//	return redisClient
//}

// GetPublicKey get public key
func GetPublicKey() interface{} {
	return publicKey
}

// GetRootPath get path of storage
func GetRootPath() string {
	return rootPath
}

// GetRPFilePath get path of storage
func GetAvatarFilePath() string {
	return storageAvatarPath
}

// GetRPFilePath get path of storage
func GetStorageFilePath() string {
	return storageFilePath
}

// GetBaseURL get base url
func GetBaseURL() string {
	return baseURL
}

// GetDocumentFetchingChannel
func GetDocumentFetchingChannel() chan string {
	return documentFetchingChan
}

// GetDocumentFetchingErrChannel
func GetDocumentFetchingErrChannel() chan error {
	return documentFetchingErrChan
}

func GetDocumentFetchingStatus() bool {
	return documentFetchingStatus
}

func SetDocumentFetchingStatus(status bool) {
	documentFetchingStatus = status
}
func GetEnvironments() string {
	return env
}
