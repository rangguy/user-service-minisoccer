package constant

const (
	success         = "success"
	failed          = "failed"
	notFound        = "data not found"
	error           = "internal server error"
	unauthorized    = "unauthorized"
	unauthenticated = "unauthenticated"
)

func Success() string {
	return success
}

func Failed() string {
	return failed
}

func NotFound() string {
	return notFound
}

func Error() string {
	return error
}

func Unauthorized() string {
	return unauthorized
}

func Unauthenticated() string {
	return unauthenticated
}
