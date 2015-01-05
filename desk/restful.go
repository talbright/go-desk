package desk

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"log"
)

type Restful struct {
	method string
	path   string
	params *url.Values
	query  *string
	body   interface{}
	json   interface{}
	client *Client
}

func (r Restful) String() string {
	return Stringify(r)
}

func (r *Restful) methodAndPath(m string, p string) *Restful {
	r.method = m
	r.path = p
	return r
}

func (r *Restful) Post(p string) *Restful {
	return r.methodAndPath("POST", p)
}

func (r *Restful) Get(p string) *Restful {
	return r.methodAndPath("GET", p)
}

func (r *Restful) Put(p string) *Restful {
	return r.methodAndPath("PUT", p)
}

func (r *Restful) Patch(p string) *Restful {
	return r.methodAndPath("PATCH", p)
}

func (r *Restful) Delete(p string) *Restful {
	return r.methodAndPath("DELETE", p)
}

func (r *Restful) Method(m string) *Restful {
	r.method = m
	return r
}

func (r *Restful) Path(p string) *Restful {
	r.path = p
	return r
}

func (r *Restful) Params(p *url.Values) *Restful {
	r.params = p
	return r
}

func (r *Restful) Query(q *string) *Restful {
	r.query = q
	return r
}

func (r *Restful) Client(c *Client) *Restful {
	r.client = c
	return r
}

func (r *Restful) Body(b interface{}) *Restful {
	r.body = b
	return r
}

func (r *Restful) Json(j interface{}) *Restful {
	r.json = j
	return r
}

func (r *Restful) Do() (*http.Response, error) {
	path := r.path
	if r.params != nil && len(*r.params) > 0 {
		path = fmt.Sprintf("%v?%v", path, r.params.Encode())
	} else if r.query != nil {
		path = fmt.Sprintf("%v?%v", path, r.query)
	}
	log.Printf("Request %v %v",strings.ToUpper(r.method),path)
	req, err := r.client.NewRequest(r.method, path, r.body)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req, r.json)
	if err != nil {
		return resp, err
	}
	return resp, err
}
