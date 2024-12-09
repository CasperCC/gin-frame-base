package response

const (
	CODE_SUCCESS = iota

	CODE_SYSTEM_ERROR = -(iota + 999)
	CODE_NO_API
	CODE_LOGIN_FAILED
	CODE_MISSING_PARAMS
	CODE_DB_NOT_FOUND
	CODE_UNAUTHORIZED
	CODE_DB_ERROR
	CODE_PASSWORD_INCORRECT
	CODE_USER_EXISTS
	CODE_STRONG_PASSWORD_REQUIRED
)

const (
	MSG_SUCCESS = "ok"

	MSG_SYSTEM_ERROR             = "System error."
	MSG_NO_API                   = "API not exist."
	MSG_LOGIN_FAILED             = "Login failed."
	MSG_MISSING_PARAMS           = "Missing parameters."
	MSG_DB_NOT_FOUND             = "Data not exist."
	MSG_UNAUTHORIZED             = "Login invalid."
	MSG_DB_ERROR                 = "System error."
	MSG_PASSWORD_INCORRECT       = "Password incorrect."
	MSG_USER_EXISTS              = "User exists."
	MSG_STRONG_PASSWORD_REQUIRED = "Strong password required."
)
