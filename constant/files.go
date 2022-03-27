package constant

const (
	UploadHTMLFile    = "./htmlfile/upload.html"
	MaxFileUploadSize = 8 << 20 // MB
	ImageFolderPath   = "./image"
)

// server status codes
const (
	StatusCodeOk             = "200"
	StatusCodeCreated        = "201"
	StatusCodeBadRequest     = "400"
	StatusCodeForbidden      = "403"
	StatusCodeSessionExpired = "440"
	StatusCodeServerError    = "500"
	StatusCodeDuplicateEntry = "1000"
)
