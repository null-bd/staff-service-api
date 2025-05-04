package errors

const (
	// region request error codes
	ErrBadRequest ErrorCode = "STAFFAPI_101"

	// region service error codes
	ErrStaffExists ErrorCode = "STAFFAPI_201"
	ErrStaffActive ErrorCode = "STAFFAPI_202"

	// region repository error codes
	ErrDatabaseConnection ErrorCode = "STAFFAPI_301"
	ErrDatabaseQuery      ErrorCode = "STFFAPI_302"
	ErrCacheConnection    ErrorCode = "STAFFAPI_303"
	ErrDatabaseOperation  ErrorCode = "STAFFAPI_304"
	ErrStaffNotFound      ErrorCode = "STAFFAPI_305"
	ErrInvalidInput       ErrorCode = "STAFFAPI_306"
)
