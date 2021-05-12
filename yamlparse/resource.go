package yamlparse

// Yaml2Go
type Yaml2GoResource struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

// Metadata
type Metadata struct {
	Labels Labels `yaml:"labels"`
	Name   string `yaml:"name"`
}

// Labels
type Labels struct {
	App string `yaml:"app"`
}

// Spec
type Spec struct {
	Template Template `yaml:"template"`
}

// Template
type Template struct {
	Spec TemplateSpec `yaml:"spec"`
}

// TemplateSpec
type TemplateSpec struct {
	Containers []Containers `yaml:"containers"`
}

// Containers
type Containers struct {
	Name      string    `yaml:"name"`
	Resources Resources `yaml:"resources"`
}

// Resources
type Resources struct {
	Requests Requests `yaml:"requests"`
	Limits   Limits   `yaml:"limits"`
}

// Requests
type Requests struct {
	Cpu string `yaml:"cpu"`
}

// Limits
type Limits struct {
	Cpu string `yaml:"cpu"`
}