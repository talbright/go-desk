package main

import (
  "fmt"
  "reflect"
  "flag"
  "net/url"
  "time"
  "github.com/talbright/go-desk/desk"
)

//We could also create a map/slice of functions, but I want to play with reflection...
type Example struct {}

func main() {
  siteUrl := flag.String("site-url", "", "site URL to use ie: mysite.desk.com")
  userEmail := flag.String("email", "", "email for authentication") 
  userPassword := flag.String("password", "", "password for authentication") 
  exampleName := flag.String("example","","example to run")
  flag.Parse()
  client := desk.NewClient(nil,*siteUrl,*userEmail,*userPassword)
  inputs := make([]reflect.Value, 1)
  inputs[0] = reflect.ValueOf(client)
  reflect.ValueOf(&Example{}).MethodByName(*exampleName).Call(inputs)
}

func (e *Example) GetCaseMessage(client *desk.Client) {
  cse,_,err := client.Case.Message.Get("1")
  if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n",cse.String())
	}
}

func (e *Example) GetCase(client *desk.Client) {
  cse,_,err := client.Case.Get("1")
  if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n",cse.String())
	}
}

func (e *Example) ListCase(client *desk.Client) {
  listParams := url.Values{}
  listParams.Add("sort_field","created_at")
  listParams.Add("sort_direction","asc")
  collection,_,err := client.Case.List(&listParams)
  if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
    fmt.Printf("%v\n\n",collection.String()) 
  } 
}

func (e *Example) SearchCase(client *desk.Client) {
  searchParams := url.Values{}
  searchParams.Add("sort_field","created_at")
  searchParams.Add("sort_direction","asc")
  searchParams.Add("status","new")
  collection,_,err := client.Case.Search(&searchParams,nil)
  if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
    fmt.Printf("%v\n\n",collection.String()) 
  } 
}

func (e *Example) UpdateCase(client *desk.Client) {
  subject := fmt.Sprintf("updated case at %v",time.Now())
  id := 1
  caze := desk.Case{ ID: &id, Subject: &subject}
  new_case,_,err := client.Case.Update(&caze)
  if err != nil {
    fmt.Printf("error: %v\n\n", err)
  } else {
    fmt.Printf("%v\n\n",new_case) 
  }
}

func (e *Example) CreateCase(client *desk.Client) {
  caze := desk.Case { }
  customer := desk.Customer {}
  message := desk.Message {}
  new_case,_,err := client.Case.Create(&caze,&customer,&message)
  if err!= nil {
    fmt.Printf("error: %v\n\n", err)
  } else {
    fmt.Printf("%v\n\n",new_case)
  }
}

