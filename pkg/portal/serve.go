package portal

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
	"net/http"
)

func InstallWith(e *echo.Echo) {
	handler := createStaticHandler("/",
		&assetfs.AssetFS{
			Asset:     Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo})

	e.GET("/", handler)
	e.HEAD("/", handler)
}

func createStaticHandler(path string, fs http.FileSystem) echo.HandlerFunc {
	fileServer := http.StripPrefix(path, http.FileServer(fs))
	return func(c echo.Context) error {
		fileServer.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
