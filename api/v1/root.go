package api

import (
	"github.com/labstack/echo"
)

var (
	// OperationIds for api handler
	ApiGetAlbums        = apiGet("/albums")             // get albums
	ApiCreateAlbums     = apiPut("/albums")             // create an albums
	ApiUpdateAlbums     = apiPost("/albums")            // update an albums
	ApiGetAlbumsById    = apiGet("/albums/:albumId")    // Get album by Id
	ApiDeleteAlbumsById = apiDelete("/albums/:albumId") // Delete an albums by Id
)

type Operation struct {
	Method string
	Path   string
}

type OperationIds map[Operation]echo.HandlerFunc

func InstallWith(e *echo.Echo, ids OperationIds) {
	g := e.Group(ApiVersion)

	for opt, handler := range ids {
		g.Add(opt.Method, opt.Path, handler)
	}
}

func apiGet(path string) Operation {
	return Operation{Method: echo.GET, Path: path}
}

func apiPut(path string) Operation {
	return Operation{Method: echo.PUT, Path: path}
}

func apiPost(path string) Operation {
	return Operation{Method: echo.POST, Path: path}
}

func apiDelete(path string) Operation {
	return Operation{Method: echo.DELETE, Path: path}
}
