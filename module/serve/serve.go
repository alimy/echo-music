package serve

import (
	"github.com/alimy/echo-music/api/v1"
	"github.com/alimy/echo-music/pkg/portal"
	"github.com/labstack/echo"
	"github.com/unisx/logus"
	"net/http"
	"time"
)

type Config struct {
	CertFile    string
	KeyFile     string
	EnableHttps bool
}

func StartService(config *Config) {
	e := echo.New()

	// Enable debug
	e.Debug = true

	// Install portal router
	portal.InstallWith(e)

	// Install api router
	api.InstallWith(e, api.OperationIds{
		api.ApiGetAlbums:        getAlbums,
		api.ApiCreateAlbums:     createAlbums,
		api.ApiUpdateAlbums:     updateAlbums,
		api.ApiGetAlbumsById:    getAlbumsById,
		api.ApiDeleteAlbumsById: deleteAlbumsById,
	})

	// Setup http.Server
	server := &http.Server{
		Addr: "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start http.Server
	if config.EnableHttps {
		logus.Info("listen and serve in https://:8080")
		e.StartServer(server)
	} else {
		logus.Info("listen and serve in http://:8080")
		e.StartServer(server)
	}
}
