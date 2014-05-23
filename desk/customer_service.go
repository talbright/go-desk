package desk

import (
  "encoding/json"
  "net/http"
  "net/url"
  "fmt"
)

type CustomerService struct {
	client *Client
}

// Get retrieves a customer.
// See Desk API: http://dev.desk.com/API/customers/#show 
func (c* CustomerService) Get(id string) (*Customer, *http.Response, error) {
  path := fmt.Sprintf("customers/%v",id)
  req, err := c.client.NewRequest("GET",path,nil)
  if err!=nil {
    return nil, nil, err
  }
  customer := new(Customer)
  resp, err := c.client.Do(req,customer)
  if err!=nil {
    return nil, resp, err
  }
  return customer, resp, err
}

// List customers with filtering and pagination. 
// See Desk API: http://dev.desk.com/API/customers/#list 
func (c* CustomerService) List(params *url.Values) (*Page,*http.Response,error) {
  path := fmt.Sprintf("customers")
  if params!=nil && len(*params) > 0 {
    path = fmt.Sprintf("%v?%v",path,params.Encode())
  }
  req, err := c.client.NewRequest("GET",path,nil)
  if err!=nil {
    return nil, nil, err
  }
  page := new(Page)
  resp, err := c.client.Do(req,page)
  if err != nil {
    return nil,resp,err
  }
  err = c.unravelPage(page)
  if err != nil {
    return nil,nil,err
  }
  return page,resp,err
}

// Search customers with filtering and pagination.
// See Desk API: http://dev.desk.com/API/customers/#search 
func (c *CustomerService) Search(params *url.Values, q *string) (*Page, *http.Response, error) {
	path := fmt.Sprintf("customers/search")
	if params != nil && len(*params) > 0 {
		path = fmt.Sprintf("%v?%v", path, params.Encode())
	} else if q != nil {
		path = fmt.Sprintf("%v?%v", path, q)
	}
	req, err := c.client.NewRequest("GET", path, nil)
  if err!=nil {
    return nil, nil, err
  }
	page := new(Page)
	resp, err := c.client.Do(req, page)
	if err != nil {
		return nil, resp, err
	}
  err = c.unravelPage(page) 
  if err != nil {
    return nil, nil, err
  }
	return page, resp, err
}

// Create a customer.
// See Desk API: http://dev.desk.com/API/customers/#create
func (c *CustomerService) Create(customer *Customer) (*Customer, *http.Response, error) {
	path := fmt.Sprintf("customers")
	req, err := c.client.NewRequest("POST", path, customer)
	if err != nil {
		return nil, nil, err
	}
	createdCustomer := new(Customer)
	resp, err := c.client.Do(req, createdCustomer)
	if err != nil {
		return nil, resp, err
	}
	return createdCustomer,resp,err
}

// Update a customer.
// See Desk API: http://dev.desk.com/API/customers/#update
func (c *CustomerService) Update(customer *Customer) (*Customer, *http.Response, error) {
	path := fmt.Sprintf("customers/%d", *customer.ID)
	req, err := c.client.NewRequest("PATCH",path,customer)
	if err != nil {
		return nil, nil, err
	}
	updatedCustomer := new(Customer)
	resp, err := c.client.Do(req, updatedCustomer)
	if err != nil {
		return nil, resp, err
	}
	return updatedCustomer, resp, err
}

// Cases provides a list of cases associated with a customer.
// See Desk API: http://dev.desk.com/API/customers/#list-cases
func (c *CustomerService) Cases(id string, params *url.Values) (*Page, *http.Response, error) {
	path := fmt.Sprintf("customers/%v/cases", id)
  if params!=nil && len(*params) > 0 {
    path = fmt.Sprintf("%v?%v",path,params.Encode())
  }
  req, err := c.client.NewRequest("GET",path,nil)
  if err != nil {
    return nil,nil,err
  }
	page := new(Page)
	resp, err := c.client.Do(req, page)
	if err != nil {
		return nil, resp, err
	}
  err = c.unravelPage(page) 
  if err != nil {
    return nil, nil, err
  }
	return page, resp, err
}

func (c* CustomerService)unravelPage(page *Page) (error) {
  customers := new([]Customer)
  err := json.Unmarshal(*page.Embedded.RawEntries,&customers)
  if err!= nil {
    return err
  }
  page.Embedded.Entries = make([]interface{},len(*customers))
  for i, v:= range *customers {
    page.Embedded.Entries[i] = interface{}(v)
  }
  page.Embedded.RawEntries=nil
  return err
} 
