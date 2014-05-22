package desk

import (
  "net/http"
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
