package cmd

import (
	"context"
	"github.com/alimy/echo-music/cmd"
	"github.com/alimy/echo-music/module/serve"
	"github.com/spf13/cobra"
)

var AppContext = context.Background()

const (
	certFilePathDefault = "cert.pem" // certificate file default path
	keyFilePathDefault  = "key.pem"  // key file used in https server default path
)

var (
	config = &serve.Config{}
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start to okmall service",
		Long:  "this cmd will start a https server to provide okmall service",
		Run:   serveRun,
	}

	// Parse flags for serveCmd
	serveCmd.Flags().StringVarP(&config.CertFile, "cert", "c", certFilePathDefault, "certificate path used in https connect")
	serveCmd.Flags().StringVarP(&config.KeyFile, "key", "k", keyFilePathDefault, "key path used in https connect")
	serveCmd.Flags().BoolVarP(&config.EnableHttps, "https", "s", false, "whether use https serve connect")

	// Register serveCmd as sub-command
	cmd.Register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	serve.StartService(config)
}
