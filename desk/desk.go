package desk

import (
	"bytes"
	"encoding/json"
	/* "errors" */
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	/* "reflect" */
	/* "strconv" */
	/* "strings" */
	/* "time" */)

const (
	DeskLibVersion = "0.1"
	DeskApiVersion = "v2"
	DeskHost       = "desk.com"
	DeskUserAgent  = "go-desk/" + DeskLibVersion
)

type Client struct {
	client       *http.Client
	BaseURL      *url.URL
	UserEmail    string
	UserPassword string
	Case         *CaseService
	Customer     *CustomerService
}

func init() {
	log.SetPrefix("[desk] ")
	log.Println("init")
	log.Printf("Desk client library (%v) for desk.com API %v\n", DeskLibVersion, DeskApiVersion)
}

func NewClient(httpClient *http.Client, endpointURL string, userEmail string, userPassword string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(fmt.Sprintf("%s/api/%s/", endpointURL, DeskApiVersion))
	c := &Client{client: httpClient, BaseURL: baseURL, UserEmail: userEmail, UserPassword: userPassword}
	c.Case = &CaseService{client: c}
	c.Customer = &CustomerService{client: c}
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)

	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	req.SetBasicAuth(c.UserEmail, c.UserPassword)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", DeskUserAgent)
	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	log.Printf("Do %v", req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)

	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err
}

type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
