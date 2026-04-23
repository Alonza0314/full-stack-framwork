package internal

import (
	"backend/config"
	"backend/constant"
	"backend/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/free-ran-ue/util"
	"github.com/gin-gonic/gin"
)

type jwt struct {
	secret    string
	expiresIn time.Duration
}

type backend struct {
	router *gin.Engine
	server *http.Server

	username string
	password string

	port int

	jwt

	frontendFilePath string

	*logger.BackendLogger
}

func NewBackend(config *config.Config, logger *logger.BackendLogger) *backend {
	b := &backend{
		router: nil,
		server: nil,

		username: config.Backend.Username,
		password: config.Backend.Password,

		port: config.Backend.Port,

		jwt: jwt{
			secret:    config.Backend.JWT.Secret,
			expiresIn: config.Backend.JWT.ExpiresIn,
		},

		frontendFilePath: config.Backend.FrontendFilePath,

		BackendLogger: logger,
	}

	b.router = util.NewGinRouter(constant.API_PREFIX, b.iniRoutes())
	b.router.NoRoute(b.returnPages())

	addMiddleware(b.router)

	return b
}

func (b *backend) returnPages() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodGet {

			destPath := filepath.Join(b.frontendFilePath, c.Request.URL.Path)
			if _, err := os.Stat(destPath); err == nil {
				c.File(filepath.Clean(destPath))
				return
			}

			c.File(filepath.Clean("build/console/index.html"))
		} else {
			c.Next()
		}
	}
}

func (b *backend) Start() {
	b.BckLog.Infoln("Starting backend server...")

	b.server = &http.Server{
		Addr:    ":" + strconv.Itoa(b.port),
		Handler: b.router,
	}

	go func() {
		if err := b.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			b.BckLog.Errorf("Failed to start server: %s\n", err)
		}
	}()
	time.Sleep(500 * time.Millisecond)

	b.BckLog.Infof("Backend server started on port: %d", b.port)
}

func (b *backend) Stop() {
	fmt.Println()
	b.BckLog.Infoln("Stopping backend server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := b.server.Shutdown(shutdownCtx); err != nil {
		b.BckLog.Errorf("Failed to stop backend server: %v", err)
	} else {
		b.BckLog.Infoln("Backend server stopped successfully")
	}
}

func (b *backend) iniRoutes() util.Routes {
	return util.Routes{
		{
			Name:        "Login",
			Method:      http.MethodPost,
			Pattern:     "/login",
			HandlerFunc: b.handleLogin,
		},
		{
			Name:        "Logout",
			Method:      http.MethodPost,
			Pattern:     "/logout",
			HandlerFunc: b.handleLogout,
		},
	}
}

func (b *backend) handleLogin(c *gin.Context) {
}

func (b *backend) handleLogout(c *gin.Context) {

}
