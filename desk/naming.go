package desk

import (
	"reflect"
	"strings"
	"github.com/gedex/inflector"
)

type Naming struct {
	ResourceName string
}

func (n* Naming) SetResourceName(thing interface{}) {
	if thing != nil {
		nameParts := strings.Split(reflect.TypeOf(thing).String(),".")
		typeName := nameParts[len(nameParts)-1]
		n.ResourceName = strings.ToLower(typeName)
		n.Pluralize()
	}
}

func (n* Naming) GetResourceName() (string) {
	return n.ResourceName
}

func (n* Naming) Pluralize() {
	n.ResourceName = inflector.Pluralize(n.ResourceName)
}

func (n* Naming) Singularize() {
	n.ResourceName = inflector.Singularize(n.ResourceName)
}

