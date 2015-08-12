package resource

type Resourceful interface {
	InitializeResource(model interface{})
	GetResourceId() (id string)
	SetResourceId(id string)
	GetResourceName() (name string)
	GetResourcePath(resource Resourceful, options ...func(*ResourcePath)) (path ResourcePath)
}

type Resource struct {
	Hal
	Naming
}

func (r *Resource) InitializeResource(model interface{}) {
	r.SetResourceName(model)
}
