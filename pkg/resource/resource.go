package resource

// Object a basic object for a Kubernetes object.
type Resource struct {
	Group        string            `json:"group"`
	Version      string            `json:"version"`
	Kind         string            `json:"kind"`
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	HealthStatus string            `json:"healthStatus"`
	Labels       map[string]string `json:"-"`
	Images       []string          `json:"-"`
}