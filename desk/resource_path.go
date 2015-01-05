package desk

import (
	"strings"
)

type ResourcePath struct {
	Target Resourceful
	Nested Resourceful
	Member bool
	Action string
	Prefix string
	Suffix string
	additionalPaths []*ResourcePath
}

func NewIdentityResourcePath(id string,target Resourceful) *ResourcePath {
	target.SetResourceId(id)
	rp := &ResourcePath{Target: target, Member:true}
	return rp
}

func NewResourcePath(target Resourceful) *ResourcePath {
	rp := &ResourcePath{Target: target}
	return rp
}

func (p* ResourcePath) Path() string {
	parts := make([]string,0,7)
	if p.Prefix != "" {
		parts = append(parts,p.Prefix)
	}
	parts = append(parts,p.Target.GetResourceName())
	if p.Member && p.Target.GetResourceId() != "" {
		parts = append(parts,p.Target.GetResourceId())
	}
	if p.Action != "" {
		parts = append(parts,p.Action)
	}
	if p.Suffix != "" {
		parts = append(parts,p.Suffix)
	}
	if p.Nested != nil {
		parts = append(parts,p.Nested.GetResourcePath(p.Nested).String())
	}
	if p.additionalPaths != nil {
		for _,v := range p.additionalPaths {
			parts = append(parts,v.Path())
		}
	}
	return strings.Join(parts,"/")
}

func (p* ResourcePath) AppendPath(rp *ResourcePath) *ResourcePath {
	if p.additionalPaths == nil {
		p.additionalPaths = make([]*ResourcePath,0)
	}
	p.additionalPaths = append(p.additionalPaths,rp)
	return p
}

func (p* ResourcePath) SetMember() *ResourcePath {
	p.Member = true
	return p
}

func (p* ResourcePath) SetCollection() *ResourcePath {
	p.Member = false
	return p
}

func (p* ResourcePath) SetAction(action string) *ResourcePath {
	p.Action = action
	return p
}

func (p* ResourcePath) SetNested(resource Resourceful) *ResourcePath {
	p.Nested = resource
	return p
}

func (p* ResourcePath) SetTarget(resource Resourceful) *ResourcePath {
	p.Target = resource
	return p
}

func (p* ResourcePath) SetSuffix(suffix string) *ResourcePath {
	p.Suffix = suffix
	return p
}

func (p* ResourcePath) SetPrefix(prefix string) *ResourcePath {
	p.Prefix = prefix
	return p
}

func (p ResourcePath) String() string {
	return p.Path()
}

func ResourcePathOptionSetCollection(rp *ResourcePath) {
	rp.Member = false
}

