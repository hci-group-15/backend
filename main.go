package main

import (
	"os"
	"time"

	"github.com/TwiN/logr"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hci-group-15/backend/config"
	"github.com/hci-group-15/backend/controller"
	"github.com/hci-group-15/backend/model"
	"github.com/hci-group-15/backend/util"
	"github.com/janishutz/oidclogin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
)

// @title janishutz Store
// @version         1.0
// @description     This is the backend for store.janishutz.com.
// @termsOfService  https://janishutz.com/legal/tos/
//
// @contact.name   Support
// @contact.url    https://janishutz.com/support
// @contact.email  support@janishutz.com
//
// @license.name  GNU Affero General Public License 3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.en.html
//
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	logr.Debug("Starting...")
	r := gin.Default()

	conf := config.Load()

	// Load templates
	r.LoadHTMLGlob("templates/*.tmpl")

	// ┌                                               ┐
	// │            Security configuration             │
	// └                                               ┘
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{conf.Urls.FrontendURL},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(util.DefaultHeaders)
	r.Use(util.RateLimiter())
	// r.SetTrustedProxies(conf.Urls.TrustedProxies)

	// ┌                                               ┐
	// │                    Session                    │
	// └                                               ┘
	secret, ok := os.LookupEnv("STORE_SESSION_SECRET")
	if !ok {
		logr.Warn("Missing secret (STORE_SESSION_SECRET is unset), using unsafe default")
		secret = "secret"
	}
	store := memstore.NewStore([]byte(secret))
	r.Use(sessions.Sessions("jhid", store))

	// ┌                                               ┐
	// │               Routes, DB, Login               │
	// └                                               ┘
	db := model.Init()
	oidclogin.ConfigureFull(r, conf.Urls.BackendURL, conf.Urls.DefaultRedirect, func(userid, name, email string, c *gin.Context) {
		// uid := db.UpdateOrCreateUser(userid, name, email)
		uid := ""
		ses := sessions.Default(c)
		ses.Set("userid", uid)
		ses.Set("username", name)
	}, true, nil, os.Getenv("PROD") == "true")

	// Routes
	c := controller.NewController(db)
	controller.AddRoutes(c, r)

	r.Run()
}
