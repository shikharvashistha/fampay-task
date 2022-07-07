package types

var (
	InternalServerError = &err{msg: "Oops! Something went wrong. Please try again."}
	InvalidInput        = &err{msg: "Invalid input. Please try again."}
	Failed              = &err{msg: "Request Failed"}
)

type err struct {
	msg string
}

func (e *err) Error() string {
	return e.msg
}

func (e *err) String() string {
	return e.msg
}
