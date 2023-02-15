package routes

type Router struct {
	Services map[string]interface{}
}

var RData *Router

func InitConfig() {
	RData = &Router{
		Services: make(map[string]interface{}),
	}
}
func (r *Router) AddService(name string, service interface{}) {
	r.Services[name] = service
}
func (r *Router) GetService(name string) interface{} {
	return r.Services[name]
}
