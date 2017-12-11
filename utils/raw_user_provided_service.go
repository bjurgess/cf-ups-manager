package utils

type rawUserProvidedService struct {
	Name string `yaml:"name,omitempty"`
	Credentials map[string]string `yaml:"credentials"`
}