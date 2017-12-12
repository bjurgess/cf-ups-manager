package utils

type rawUserProvidedService struct {
	Name string `yaml:"name,omitempty"`
	Syslog string `yaml:"syslog,omitempty"`
	RouteService string `yaml:"route-service,omitempty"`
	Credentials map[string]string `yaml:"credentials,omitempty"`
}