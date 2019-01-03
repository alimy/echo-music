package cmd

import (
	"github.com/alimy/echo-music/api/v1"
	"github.com/alimy/echo-music/cmd"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
	"net/http"
	"time"

	_ "github.com/alimy/echo-music/module/serve"
	_ "github.com/alimy/echo-music/pkg/portal"
)

const (
	certFilePathDefault = "cert.pem" // certificate file default path
	keyFilePathDefault  = "key.pem"  // key file used in https server default path
)

var (
	certFile    string
	keyFile     string
	enableHttps bool
	inDebug     bool
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start to echoMusic service",
		Long:  "this cmd will start a https server to provide echoMusic service",
		Run:   serveRun,
	}

	// Parse flags for serveCmd
	serveCmd.Flags().StringVarP(&certFile, "cert", "c", certFilePathDefault, "certificate path used in https connect")
	serveCmd.Flags().StringVarP(&keyFile, "key", "k", keyFilePathDefault, "key path used in https connect")
	serveCmd.Flags().BoolVarP(&enableHttps, "https", "s", false, "whether use https serve connect")
	serveCmd.Flags().BoolVarP(&inDebug, "debug", "d", false, "whether in debug mode")

	// Register serveCmd as sub-command
	cmd.Register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	e := echo.New()

	// Setup debug mode
	e.Debug = inDebug

	// Use echo middleware
	e.Use(middleware.Recover())
	if inDebug {
		// Use echo middleware
		e.Use(middleware.Logger())
	}

	// Install api router
	api.InstallDefault(e)

	// Setup http.Server
	server := &http.Server{
		Addr: "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start http.Server
	e.StartServer(server)
}
