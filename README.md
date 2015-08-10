[![GoDoc](https://godoc.org/github.com/talbright/go-desk?status.png)](https://godoc.org/github.com/talbright/go-desk)
[![Build Status](https://travis-ci.org/talbright/go-desk.png?branch=master)](https://travis-ci.org/talbright/go-desk)

Desk.com APIv2 client implementation in Go (http://dev.desk.com/API)

Most major endpoints have been implemented, leaving the more isoteric ones
still to do. Those likely won't be worked on unless there's a demand for it.
See the project issues section for up-to-date information on what's on the
roadmap.

### Examples

There's two ways to create request bodies.

Using the builder pattern:

```go
message:=resource.MessageBuilder.
	SetString("Direction","in").
	SetString("Status","received").
	SetString("To","someone@desk.com").
	SetString("From","someone-else@desk.com").
	SetString("Subject","Case created by API via desk-go").
	SetString("Body","Please assist me with this case").
	BuildMessage()
```

Using a constructor:

```go
message:=resource.NewMessage()
message.Direction=types.String("in")
message.Status=types.String("received")
message.To=types.String("someone@desk.com")
message.From=types.String("someone-else@desk.com")
message.Subject=types.String("Case created by API via desk-go")
message.Body=types.String("Please assist me with this case")
```

Struct literal composition is not supported, as the constructor
performs some additional initialization(s).

For additional examples of the desk API usage, look at the tests in the integration directory.

#### Create a new case

```go
func main() {

  //create http client
	siteUrl := "mysite.desk.com"
	userEmail := "mysite@somewhere.com"
	userPassword := "mysite.desk.com pass"
	client := service.NewClient(nil,siteUrl,userEmail,userPassword)

  //create a new case
	message:=resource.MessageBuilder.
		SetString("Direction","in").
		SetString("Status","received").
		SetString("To","someone@desk.com").
		SetString("From","someone-else@desk.com").
		SetString("Subject","Case created by API via desk-go").
		SetString("Body","Please assist me with this case").
		BuildMessage()
	caze:=resource.CaseBuilder.
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
