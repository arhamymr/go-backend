package server

type (
	Route struct {
		url     string
		method  string
		handler HandlerFunc
		Middleware []HandlerFunc
	}
	
	HandlerFunc func (ctx *Context)
)



// func (r *Route) AddMiddleware(middleware ...func(HandlerFunc) HandlerFunc) *Route {
// 	r.Middleware = append(r.Middleware, middleware...)
// 	return r
// }

func (r *Route) AddMiddleware(middleware ...HandlerFunc) *Route {
	r.Middleware = append(r.Middleware, middleware...)
	return r
}


func (r *Route) applyMiddleware(ctx *Context) {
	for _, m := range r.Middleware {
		m(ctx)
	}
}
