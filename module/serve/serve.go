package serve

import (
	"github.com/alimy/echo-music/api/v1"
	"github.com/labstack/echo"
	"github.com/unisx/logus"
	"net/http"
)

func init() {
	api.Register(api.ApiGetAppInfo, getAppInfo)
	api.Register(api.ApiGetAlbums, getAlbums)
	api.Register(api.ApiCreateAlbums, createAlbums)
	api.Register(api.ApiUpdateAlbums, updateAlbums)
	api.Register(api.ApiGetAlbumsById, getAlbumsById)
	api.Register(api.ApiDeleteAlbumsById, deleteAlbumsById)
}

func getAppInfo(context echo.Context) error {
	// TODO
	logus.Debug("get application information")
	return context.String(http.StatusOK, "get application information")
}

func getAlbums(context echo.Context) error {
	// TODO
	logus.Debug("get albums")
	return context.String(http.StatusOK, "get albums")
}

func createAlbums(context echo.Context) error {
	// TODO
	logus.Debug("create albums")
	return context.String(http.StatusCreated, "Albums item created")
}

func updateAlbums(context echo.Context) error {
	// TODO
	logus.Debug("update albums")
	return context.String(http.StatusCreated, "Albums item updated")
}

func getAlbumsById(context echo.Context) error {
	// TODO
	albumId := context.Param("albumId")
	logus.Debug("get albums by id", logus.String("albumId", albumId))
	return context.String(http.StatusOK, "get albums by id")
}

func deleteAlbumsById(context echo.Context) error {
	// TODO
	albumId := context.Param("albumId")
	logus.Info("delete albums", logus.String("albumId", albumId))
	return context.String(http.StatusOK, "Albums item deleted")
}
