package server

type Reflector struct {
	services []string
}

func NewReflector() *Reflector {
	return &Reflector{}
}

func (r *Reflector) Names() []string {
	return r.services
}

func (r *Reflector) AddService(s string) {
	r.services = append(r.services, s)
}
