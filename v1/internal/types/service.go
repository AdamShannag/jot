package types

type Service struct {
	Name        string   `yaml:"name"`
	Port        int      `yaml:"port"`
	Endpoints   []string `yaml:"endpoints"`
	Middlewares []string `yaml:"middlewares"`
}

func NewService(name string, port int, endpoints []string, middelwares []string) *Service {
	return &Service{
		Name:        name,
		Port:        port,
		Endpoints:   endpoints,
		Middlewares: middelwares,
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
