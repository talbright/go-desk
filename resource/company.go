package resource

import(
	. "github.com/talbright/go-desk/types"
)

type Company struct {
	ExternalID   *string                `json:"external_id,omitempty"`
	Name         *string                `json:"name,omitempty"`
	Domains      []string               `json:"domains,omitempty"`
	CreatedAt    *Timestamp             `json:"created_at,omitempty"`
	UpdatedAt    *Timestamp             `json:"updated_at,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Resource
}

func NewCompany() *Company {
	company := &Company{}
	company.InitializeResource(company)
	return company
}

func (c Company) String() string {
	return Stringify(c)
}

func (c *Company) AddDomain(domain string) {
	c.Domains = append(c.Domains, domain)
}