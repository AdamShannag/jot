package types

type Service struct {
	Name      string   `yaml:"name"`
	Port      int      `yaml:"port"`
	Endpoints []string `yaml:"endpoints"`
}

func NewService(name string, port int, endpoints []string) *Service {
	return &Service{
		Name:      name,
		Port:      port,
		Endpoints: endpoints,
	}
}

func IsExistingService(services []Service, service string) (bool, int) {
	for i, srv := range services {
		if srv.Name == service {
			return true, i
		}
	}
	return false, -1
}
