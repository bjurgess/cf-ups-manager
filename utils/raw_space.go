package utils

type rawSpace struct {
	Name string `yaml:"name,omitempty"`
	UserProvidedServices []UserProvidedService `yaml:"user-provided-services,omitempty"`
}