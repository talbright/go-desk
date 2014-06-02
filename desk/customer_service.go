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
  restful := Restful{}
  customer := new(Customer)
  resp, err := restful.
    Get(fmt.Sprintf("customers/%v",id)).
    Json(customer).
    Client(c.client).
    Do()
  return customer, resp, err
}

// List customers with filtering and pagination. 
// See Desk API: http://dev.desk.com/API/customers/#list 
func (c* CustomerService) List(params *url.Values) (*Page,*http.Response,error) {
  restful := Restful{}
  page := new(Page)
  resp,err := restful.
    Get("customers").
    Json(page).
    Params(params).
    Client(c.client).
    Do()
	if err != nil {
		return nil,resp,err
	}
  err = c.unravelPage(page)
  if err != nil {
    return nil,nil,err
  }
	return page, resp, err
}

// Search customers with filtering and pagination.
// See Desk API: http://dev.desk.com/API/customers/#search 
func (c *CustomerService) Search(params *url.Values, q *string) (*Page, *http.Response, error) {
  restful := Restful{}
  page := new(Page)
  resp,err := restful.
    Get("customers/search").
    Json(page).
    Params(params).
    Client(c.client).
    Do()
	if err != nil {
		return nil,resp,err
	}
  err = c.unravelPage(page)
  if err != nil {
    return nil,nil,err
  }
	return page, resp, err
}

// Create a customer.
// See Desk API: http://dev.desk.com/API/customers/#create
func (c *CustomerService) Create(customer *Customer) (*Customer, *http.Response, error) {
  restful:=Restful{}
	createdCustomer := new(Customer)
  resp,err:=restful.
    Post("customers").
    Body(customer).
    Json(createdCustomer).
    Client(c.client).
    Do()
  return createdCustomer,resp,err
}

// Update a customer.
// See Desk API: http://dev.desk.com/API/customers/#update
func (c *CustomerService) Update(customer *Customer) (*Customer, *http.Response, error) {
  restful:=Restful{}
  updatedCustomer:=new(Customer)
  resp,err:=restful.
    Patch(fmt.Sprintf("customers/%d", *customer.ID)).
    Body(customer).
    Json(updatedCustomer).
    Client(c.client).
    Do()
  return updatedCustomer,resp,err
}

// Cases provides a list of cases associated with a customer.
// See Desk API: http://dev.desk.com/API/customers/#list-cases
func (c *CustomerService) Cases(id string, params *url.Values) (*Page, *http.Response, error) {
  restful := Restful{}
  page := new(Page)
  resp,err := restful.
    Get(fmt.Sprintf("customers/%v/cases", id)).
    Json(page).
    Params(params).
    Client(c.client).
    Do()
	if err != nil {
		return nil,resp,err
	}
  err = c.unravelPage(page)
  if err != nil {
    return nil,nil,err
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
