package main

import (
	"flag"
	"fmt"
	"gintest/constant"
	"gintest/controller"
	"gintest/driver"
	"gintest/helper"
	"gintest/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dbUsage := flag.Bool("db_usage", false, "Config for flag usage")
	dbHost := flag.String("db_host", "", "Config for db host")
	dbUsername := flag.String("db_username", "", "Config for username")
	dbPassword := flag.String("db_password", "", "Config for db password")
	dbName := flag.String("db_name", "", "Config for db name")
	dbSSLMode := flag.String("db_ssl_mode", "", "Config for db ssl mode")
	dbDialect := flag.String("db_dialect", "", "Config for db_dialect")
	dbPath := flag.String("db_path", "", "Config for db_path")
	dbHostParam := flag.String("db_host_param", "", "Config db port and db instance, db port ex: :8080 or db instance ex: /sql2014 for sql server" +
		" and db port ex : 8080 for other than sql server")

	flag.Parse()
	driver.SetParamGorm(driver.Parameter{
		UseCli:*dbUsage,
		SslMode:*dbSSLMode,
		DbName:*dbName,
		Password:*dbPassword,
		User:*dbUsername,
		HostParam:*dbHostParam,
		Host:*dbHost,
		Dialect:*dbDialect,
		DbPath:*dbPath,
	})

	db,errInitDB := driver.InitGorm()
	if errInitDB!=nil {
		fmt.Print(errInitDB.Error())
	}else{
		dialect := *dbDialect
		if dialect==""{
			dialect = constant.POSTGRESQL_DIALECT
		}
		err := helper.InitDriver(dialect)
		if err!=nil{
			log.Fatal(err)
		}
		defer db.Close()
	}
	router := gin.Default()
	router.POST("/login", controller.Login)
	router.GET("/get_events", middleware.AuthMiddleware(), controller.GetEvents)
	router.GET("/get_sermon", middleware.AuthMiddleware(), controller.GetSermon)
	router.Run(":3000")

	//r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//
	//// the jwt middleware
	//authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
	//	Realm:       "test zone",
	//	Key:         []byte("secret key"),
	//	Timeout:     time.Hour,
	//	MaxRefresh:  time.Hour,
	//	IdentityKey: identityKey,
	//	PayloadFunc: func(data interface{}) jwt.MapClaims {
	//		if v, ok := data.(*dto.User); ok {
	//			return jwt.MapClaims{
	//				"user_id": v.UserID,
	//				"username": v.Username,
	//				"app_id": v.AppID,
	//				"app_name": v.AppName,
	//				"subscription_type": v.SubscriptionType,
	//				"subscription_end": v.SubscriptionEnd,
	//				"status_design": v.StatusDesign,
	//			}
	//		}
	//		return jwt.MapClaims{}
	//	},
	//	IdentityHandler: func(c *gin.Context) interface{} {
	//		claims := jwt.ExtractClaims(c)
	//		return &dto.User{
	//			UserID: claims["user_id"].(string),
	//		}
	//	},
	//	Authenticator: func(c *gin.Context) (interface{}, error) {
	//		var loginVals login
	//		if err := c.ShouldBind(&loginVals); err != nil {
	//			return "", jwt.ErrMissingLoginValues
	//		}
	//		userID := loginVals.Username
	//		password := loginVals.Password
	//
	//		user, err := repository.Login(dto.Login{Username:userID, Password:password})
	//		if err != nil {
	//			return nil, jwt.ErrFailedAuthentication
	//		}
	//
	//		return user, nil
	//	},
	//	Authorizator: func(data interface{}, c *gin.Context) bool {
	//		if _, ok := data.(*dto.User); ok{
	//			return true
	//		}
	//
	//		return false
	//	},
	//	Unauthorized: func(c *gin.Context, code int, message string) {
	//		c.JSON(code, gin.H{
	//			"code":    code,
	//			"message": message,
	//		})
	//	},
	//	// TokenLookup is a string in the form of "<source>:<name>" that is used
	//	// to extract token from the request.
	//	// Optional. Default value "header:Authorization".
	//	// Possible values:
	//	// - "header:<name>"
	//	// - "query:<name>"
	//	// - "cookie:<name>"
	//	// - "param:<name>"
	//	TokenLookup: "header: Authorization, query: token, cookie: jwt",
	//	// TokenLookup: "query:token",
	//	// TokenLookup: "cookie:token",
	//
	//	// TokenHeadName is a string in the header. Default value is "Bearer"
	//	TokenHeadName: "Bearer",
	//
	//	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	//	TimeFunc: time.Now,
	//})
	//
	//if err != nil {
	//	log.Fatal("JWT Error:" + err.Error())
	//}
	//
	//r.POST("/login", authMiddleware.LoginHandler)
	//r.GET("/get_events", authMiddleware.MiddlewareFunc(),controller.GetEvents)
	//r.GET("/get_sermon", authMiddleware.MiddlewareFunc(),controller.GetSermon)
	//
	//http.ListenAndServe(":3000", r)
}
