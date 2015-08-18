package resource

import (
	. "github.com/talbright/go-desk/types"
)

type Group struct {
	Name *string `json:"name,omitempty"`
	Resource
}

func NewGroup() *Group {
	group := &Group{}
	group.InitializeResource(group)
	return group
}

func (c Group) String() string {
	return Stringify(c)
}
