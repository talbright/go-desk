package main

import (
  "fmt"
  "flag"
  "net/url"
  "github.com/talbright/go-desk/desk"
)

func main() {
  siteUrl := flag.String("site-url", "", "site URL to use ie: mysite.desk.com")
  userEmail := flag.String("email", "", "email for authentication") 
  userPassword := flag.String("password", "", "password for authentication") 
  flag.Parse()
  client := desk.NewClient(nil,*siteUrl,*userEmail,*userPassword)
  SearchExample(client) 
}

func GetExample(client *desk.Client) {
  cse,_,err := client.Case.Get("1")
  if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n",cse.String())
	}
}

func ListExample(client *desk.Client) {
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

func SearchExample(client *desk.Client) {
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
