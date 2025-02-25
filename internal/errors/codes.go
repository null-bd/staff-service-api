package errors

const (
	// region request error codes
	ErrBadRequest ErrorCode = "DOMSAPI_101"

	// region service error codes
	ErrDomExists ErrorCode = "DOMSAPI_201"
	ErrDomActive ErrorCode = "DOMSAPI_202"

	// region repository error codes
	ErrDatabaseConnection ErrorCode = "DOMSAPI_301"
	ErrDatabaseQuery      ErrorCode = "DOMSAPI_302"
	ErrCacheConnection    ErrorCode = "DOMSAPI_303"
	ErrDatabaseOperation  ErrorCode = "DOMSAPI_304"
	ErrDomNotFound        ErrorCode = "DOMSAPI_305"
)
