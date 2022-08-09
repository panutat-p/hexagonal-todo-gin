package port

type Handler func(ctx Context)

type Context interface {
	Bind(interface{}) error
	Json(int, interface{})
	TransactionId() string
	Audience() string
}
