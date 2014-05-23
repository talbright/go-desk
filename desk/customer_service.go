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

func (c* CustomerService) List(params *url.Values) (*Page,*http.Response,error) {
  path := fmt.Sprintf("customers")
  if params!=nil && len(*params) > 0 {
    path = fmt.Sprintf("%v?%v",path,params.Encode())
  }
  req, err := c.client.NewRequest("GET",path,nil)
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

func (c *CustomerService) Search(params *url.Values, q *string) (*Page, *http.Response, error) {
	path := fmt.Sprintf("customers/search")
	if params != nil && len(*params) > 0 {
		path = fmt.Sprintf("%v?%v", path, params.Encode())
	} else if q != nil {
		path = fmt.Sprintf("%v?%v", path, q)
	}
	req, err := c.client.NewRequest("GET", path, nil)
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
