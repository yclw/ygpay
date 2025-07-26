package middleware

var DefaultMiddleware = NewMiddleware()

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}
