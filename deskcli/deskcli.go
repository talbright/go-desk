package main

import (
  "fmt"
  "flag"
  "github.com/talbright/go-desk/desk"
)

func main() {
  siteUrl := flag.String("site-url", "", "site URL to use ie: mysite.desk.com")
  userEmail := flag.String("email", "", "email for authentication") 
  userPassword := flag.String("password", "", "password for authentication") 
  flag.Parse()
  client := desk.NewClient(nil,*siteUrl,*userEmail,*userPassword)
  cse,_,err := client.Cases.Get("1")
  if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n",cse.String())
	}
}

