package main

import (
    "database/sql"
    "time"
    "fmt"

	"log"
	"net/http"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
    "github.com/tatsuhidehirakawa/STGprd_devpkg/controllers"
    "github.com/tatsuhidehirakawa/STGprd_devpkg/config"
)

func main() {

	cfg := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Sslmode) // 「dsn」に結合した文字列を格納
	db, err := sql.Open("postgres", dsn)

	accountAttributeController := controllers.NewAccountAttribute(db)

    if err != nil {
        log.Fatal(err)
    }

	router := gin.Default()

	// // Logon function
	// // セッションCookieの設定
	// store := cookie.NewStore([]byte("secret"))
	// router.Use(sessions.Sessions("mysession", store))
	// // ログイン用のhandler
	// router.POST("/login", handler.Login)
	// // 認証済のみアクセス可能なグループ
	// authUserGroup := router.Group("/auth")
	// authUserGroup.Use(middleware.LoginCheckMiddleware())
	// {
	// 	authUserGroup.GET("/getSample", handler.getSample)
	// }

	// Access Allowance
	setCors := func(r *gin.Engine) {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"}, // or host.docker.internal, or 110wbs_dev
			AllowMethods:     []string{http.MethodPut, http.MethodPatch},
			AllowHeaders:     []string{"Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	setCors(router)
	router.GET("/someGet", accountAttributeController.Get)
	router.POST("/somePost", accountAttributeController.Post)
	router.PUT("/somePut", accountAttributeController.Put)
	router.DELETE("/someDelete/:accountID", accountAttributeController.Delete)

	router.Run()
}