// +build portal

package portal

import (
	"github.com/alimy/echo-music/api/v1"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
	"net/http"
)

func init() {
	staticHandler := createStaticHandler("/static", &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo})
	api.Register(api.ApiGetStaticAssets, staticHandler)
	api.Register(api.ApiHeadStaticAssets, staticHandler)
}

func createStaticHandler(path string, fs http.FileSystem) echo.HandlerFunc {
	fileServer := http.StripPrefix(path, http.FileServer(fs))
	return func(c echo.Context) error {
		fileServer.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
