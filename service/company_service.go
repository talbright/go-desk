package service

import (
	"encoding/json"
	. "github.com/talbright/go-desk/resource"
	"net/http"
	"net/url"
)

type CompanyService struct {
	client *Client
}

// Get retrieves a company.
// See Desk API: http://dev.desk.com/API/companies/#show
func (c *CompanyService) Get(id string) (*Company, *http.Response, error) {
	restful := Restful{}
	company := NewCompany()
	path := NewIdentityResourcePath(id, company)
	resp, err := restful.
		Get(path.Path()).
		Json(company).
		Client(c.client).
		Do()
	return company, resp, err
}

// List companies with filtering and pagination.
// See Desk API: http://dev.desk.com/API/companies/#list
func (c *CompanyService) List(params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewResourcePath(NewCompany())
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(c.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = c.unravelPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

// Search companies with filtering and pagination.
// See Desk API: http://dev.desk.com/API/companies/#search
func (c *CompanyService) Search(params *url.Values, q *string) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewResourcePath(NewCompany()).SetAction("search")
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(c.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = c.unravelPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

// Create a company.
// See Desk API: http://dev.desk.com/API/companies/#create
func (c *CompanyService) Create(company *Company) (*Company, *http.Response, error) {
	restful := Restful{}
	createdCompany := new(Company)
	path := NewResourcePath(NewCompany())
	resp, err := restful.
		Post(path.Path()).
		Body(company).
		Json(createdCompany).
		Client(c.client).
		Do()
	return createdCompany, resp, err
}

// Update a company.
// See Desk API: http://dev.desk.com/API/companies/#update
func (c *CompanyService) Update(company *Company) (*Company, *http.Response, error) {
	restful := Restful{}
	updatedCompany := new(Company)
	path := NewIdentityResourcePath(company.GetResourceId(), company)
	resp, err := restful.
		Patch(path.Path()).
		Body(company).
		Json(updatedCompany).
		Client(c.client).
		Do()
	return updatedCompany, resp, err
}

// Cases provides a list of companies associated with a company.
// See Desk API: http://dev.desk.com/API/companies/#list-cases
func (c *CompanyService) Cases(id string, params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewIdentityResourcePath(id, NewCompany()).SetNested(NewCase())
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(c.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = c.unravelPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

// Customers provides a list of companies associated with a company.
// See Desk API: http://dev.desk.com/API/companies/#customers-list
func (c *CompanyService) Customers(id string, params *url.Values) (*Page, *http.Response, error) {
	restful := Restful{}
	page := new(Page)
	path := NewIdentityResourcePath(id, NewCompany()).SetNested(NewCustomer())
	resp, err := restful.
		Get(path.Path()).
		Json(page).
		Params(params).
		Client(c.client).
		Do()
	if err != nil {
		return nil, resp, err
	}
	err = c.unravelPage(page)
	if err != nil {
		return nil, nil, err
	}
	return page, resp, err
}

func (c *CompanyService) unravelPage(page *Page) error {
	companies := new([]Company)
	err := json.Unmarshal(*page.Embedded.RawEntries, &companies)
	if err != nil {
		return err
	}
	page.Embedded.Entries = make([]interface{}, len(*companies))
	for i, v := range *companies {
		v.InitializeResource(v)
		page.Embedded.Entries[i] = interface{}(v)
	}
	page.Embedded.RawEntries = nil
	return err
}
