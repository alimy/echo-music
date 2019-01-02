package serve

import (
	"github.com/labstack/echo"
	"github.com/unisx/logus"
	"net/http"
)

func getAlbums(context echo.Context) error {
	// TODO
	logus.Debug("get albums")
	return context.String(http.StatusOK, "Albums")
}

func createAlbums(context echo.Context) error {
	// TODO
	logus.Debug("create albums")
	return context.String(http.StatusCreated, "Albums item to update")
}

func updateAlbums(context echo.Context) error {
	// TODO
	logus.Debug("update albums")
	return context.String(http.StatusCreated, "Albums item to update")
}

func getAlbumsById(context echo.Context) error {
	// TODO
	albumId := context.Param("albumId")
	logus.Debug("get albums by id", logus.String("albumId", albumId))
	return context.String(http.StatusOK, "albums")
}

func deleteAlbumsById(context echo.Context) error {
	// TODO
	albumId := context.Param("albumId")
	logus.Info("delete albums", logus.String("albumId", albumId))
	return context.String(http.StatusOK, "Item deleted")
}
