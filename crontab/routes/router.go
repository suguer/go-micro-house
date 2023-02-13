package routes

type Router struct {
	Services map[string]interface{}
}

func NewRouter() *Router {
	return &Router{
		Services: make(map[string]interface{}),
	}
}
func (r *Router) AddService(name string, service interface{}) {
	r.Services[name] = service
}
