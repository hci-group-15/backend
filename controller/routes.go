package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/docs"
	"github.com/janishutz/oidclogin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(c *Controller, r *gin.Engine) {
	users := r.Group("/user", oidclogin.EnsureLogin(false))
	users.GET("profile", c.GetUserProfile)
	subjects := users.Group("subjects")
	subjects.GET("get", c.GetUserSubjects)

	// ── Healthcheck ──────────────────────────────────────────────────
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": true}) })

	// ── Swagger ──────────────────────────────────────────────────────
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = "/"
}
