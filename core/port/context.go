package port

type WrappedHandler func(ctx Context)

type Context interface {
	Bind(interface{}) error
	Json(int, interface{})
	TransactionId() string
	Audience() string
}
