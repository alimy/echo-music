package api

const (
	// OperationIds for api handler
	ApiGetStaticAssets  = iota // Get static assets
	ApiHeadStaticAssets        // Head static assets
	ApiGetAppInfo              // Get application information
	ApiGetAlbums               // Get albums
	ApiCreateAlbums            // Create an albums
	ApiUpdateAlbums            // Update an albums
	ApiGetAlbumsById           // Get album by Id
	ApiDeleteAlbumsById        // Delete an albums by Id
)

// OperationIds contains id-operation map used for install and register handler info
var OperationIds = map[int]*Operation{
	ApiGetStaticAssets:  apiGet("/static/*"),
	ApiHeadStaticAssets: apiHead("/static/*"),
	ApiGetAppInfo:       apiGet("/appinfo/", ApiVersion),
	ApiGetAlbums:        apiGet("/albums/", ApiVersion),
	ApiCreateAlbums:     apiPut("/albums/", ApiVersion),
	ApiUpdateAlbums:     apiPost("/albums/", ApiVersion),
	ApiGetAlbumsById:    apiGet("/albums/:albumId", ApiVersion),
	ApiDeleteAlbumsById: apiDelete("/albums/:albumId", ApiVersion),
}
