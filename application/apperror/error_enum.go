package apperror

const (
	FailUnmarshalResponseBodyError ErrorType = "ER400 Fail to unmarshal response body"  // used by controller
	ObjectNotFound                 ErrorType = "ER404 Object %s is not found"           // used by injected repo in interactor
	UnrecognizedEnum               ErrorType = "ER500 %s is not recognized %s enum"     // used by enum
	DatabaseNotFoundInContextError ErrorType = "ER500 Database is not found in context" // used by repoimpl
	NameMustNotEmpty               ErrorType = "ER400 name must not empty"              //
	EmailMustNotEmpty              ErrorType = "ER400 email must not empty"             //
	PasswordMustNotEmpty           ErrorType = "ER400 password must not empty"          //
	ERR500                         ErrorType = "ER500 %s"                               //
	ERR400                         ErrorType = "ER400 %s"                               //
	EmailAlreadyExist              ErrorType = "ER400 email already exist"              //
	UnsupportedImageFormat         ErrorType = "ER400 unsupported image format"         //
)
