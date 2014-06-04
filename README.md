Desk.com APIv2 client implementation in Go (http://dev.desk.com/API)

With over 200 API endpoints, the Desk API is a comprehensize one. To check
on which endpoints have been implemented, please see the [status](STATUS.md) 
page.

The end goal is implement most of the API endpoints, but prioritizing the ones
that are most important first. If you need a particular endpoint that hasn't
been built yet, feel free to open an issue request, or better yet [submit](CONTRIBUTING.MD) a 
patch.

### Example

```go
func main() {

  //create http client 
  siteUrl := "mysite.desk.com"
  userEmail := "mysite@somewhere.com"
  userPassword := "mysite.desk.com pass" 
  client := desk.NewClient(nil,siteUrl,userEmail,userPassword)
  
  //create a new case
  message:=desk.MessageBuilder.
    SetString("Direction","in").
    SetString("Status","received").
    SetString("To","someone@desk.com").
    SetString("From","someone-else@desk.com").
    SetString("Subject","Case created by API via desk-go").
    SetString("Body","Please assist me with this case").
    BuildMessage()
  caze:=desk.CaseBuilder.
    SetString("Type","email").
    SetString("Subject","Case created by API via desk-go").
    SetInt("Priority",4).
    SetString("Status","received").
    SetMessage(message).
    AddHrefLink("customer",fmt.Sprintf("/api/v2/customers/%d",192220782)).
    BuildCase()
  newCase,_,err := client.Case.Create(caze)
  if err != nil {
    fmt.Printf("error: %v\n\n", err)
  } else {
    fmt.Printf("%v\n\n",newCase.String())
  }   
}
```

### Other Libraries

Libraries in other languages are also available:

* https://github.com/tstachl/desk_api [ruby]
* https://github.com/tstachl/desk.js [node]
* https://github.com/chriswarren/desk [ruby]
* https://github.com/eventbrite/deskapi [python]
* https://github.com/bradfeehan/desk-php [php]
