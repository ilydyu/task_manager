package middlware

type Middleware struct {
	secret string
}

func NewMiddlware(secret string) *Middleware {
	return &Middleware{
		secret: secret,
	}
}
