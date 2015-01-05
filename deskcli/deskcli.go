package main

import (
	"flag"
	"fmt"
	"github.com/talbright/go-desk/desk"
	"net/url"
	"reflect"
	"time"
)

//We could also create a map/slice of functions, but I want to play with reflection...
type Example struct{}

const DefaultCustomerId int = 192220782

func main() {
	siteUrl := flag.String("site-url", "", "site URL to use ie: mysite.desk.com")
	userEmail := flag.String("email", "", "email for authentication")
	userPassword := flag.String("password", "", "password for authentication")
	exampleName := flag.String("example", "", "example to run")
	flag.Parse()
	client := desk.NewClient(nil, *siteUrl, *userEmail, *userPassword)
	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(client)
	reflect.ValueOf(&Example{}).MethodByName(*exampleName).Call(inputs)
}

//-----------------------------------------------------------------------------
//Utilities
//-----------------------------------------------------------------------------
func HandleResults(resource desk.Stringable, err error) {
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n", resource.String())
	}
}

func BuildSampleCase() *desk.Case {
	message := desk.MessageBuilder.
		SetString("Direction", "in").
		SetString("Status", "received").
		SetString("To", "someone@desk.com").
		SetString("From", "someone-else@desk.com").
		SetString("Subject", "Case created by API via desk-go").
		SetString("Body", "Please assist me with this case").
		BuildMessage()
	caze := desk.CaseBuilder.
		SetString("Type", "email").
		SetString("Subject", "Case created by API via desk-go").
		SetInt("Priority", 4).
		SetString("Status", "received").
		SetMessage(message).
		AddHrefLink("customer", fmt.Sprintf("/api/v2/customers/%d", DefaultCustomerId)).
		BuildCase()
	return &caze
}

//-----------------------------------------------------------------------------
//Case Drafts
//-----------------------------------------------------------------------------
func (e *Example) CreateCaseDraft(client *desk.Client) {
	cze := BuildSampleCase()
	createdCase, _, err := client.Case.Create(cze)
	HandleResults(createdCase, err)
	draft := desk.DraftBuilder.
		SetString("Body", "some body").
		SetString("Direction", "out").
		SetString("Status", "draft").
		BuildDraft()
	newDraft, _, err := client.Case.Draft.Create(createdCase.GetResourceId(),&draft)
	HandleResults(newDraft, err)
}

func (e *Example) ShowCaseDraft(client *desk.Client) {
	cze := BuildSampleCase()
	createdCase, _, err := client.Case.Create(cze)
	HandleResults(createdCase, err)
	draft := desk.DraftBuilder.
		SetString("Body", "some body").
		SetString("Direction", "out").
		SetString("Status", "draft").
		BuildDraft()
	client.Case.Draft.Create(createdCase.GetResourceId(), &draft)
	showDraft, _, err := client.Case.Draft.Get(createdCase.GetResourceId())
	HandleResults(showDraft, err)
}

func (e *Example) UpdateCaseDraft(client *desk.Client) {
	cze := BuildSampleCase()
	createdCase, _, err := client.Case.Create(cze)
	HandleResults(createdCase, err)
	draft := desk.DraftBuilder.
		SetString("Body", "some body").
		SetString("Direction", "out").
		SetString("Status", "draft").
		BuildDraft()
	newDraft, _, err := client.Case.Draft.Create(createdCase.GetResourceId(), &draft)
	HandleResults(newDraft, err)
	body := fmt.Sprintf("body updated at %v", time.Now())
	newDraft.Body = &body
	//TODO this marshalls to null, but the API cannot handle null
	delete(newDraft.Links, "outbound_mailbox")
	updatedDraft, _, err := client.Case.Draft.Update(createdCase.GetResourceId(), newDraft)
	HandleResults(updatedDraft, err)
}

//-----------------------------------------------------------------------------
//Case Replies
//-----------------------------------------------------------------------------
func (e *Example) ListCaseReplies(client *desk.Client) {
	listParams := url.Values{}
	listParams.Add("sort_field", "created_at")
	listParams.Add("sort_direction", "asc")
	collection, _, err := client.Case.Reply.List("1", &listParams)
	HandleResults(collection, err)
}

func (e *Example) CreateCaseReply(client *desk.Client) {
	cze := BuildSampleCase()
	createdCase, _, err := client.Case.Create(cze)
	HandleResults(createdCase, err)
	reply := desk.ReplyBuilder.
		SetString("Body", "some body").
		SetString("Direction", "out").
		SetString("Status", "draft").
		BuildReply()
	newReply, _, err := client.Case.Reply.Create(createdCase.GetResourceId(), &reply)
	newReply.GetId()
	HandleResults(newReply, err)
}

func (e *Example) UpdateCaseReply(client *desk.Client) {
	cze := BuildSampleCase()
	createdCase, _, err := client.Case.Create(cze)
	HandleResults(createdCase, err)
	reply := desk.ReplyBuilder.
		SetString("Body", "some body").
		SetString("Direction", "out").
		SetString("Status", "draft").
		BuildReply()
	newReply, _, err := client.Case.Reply.Create(createdCase.GetResourceId(), &reply)
	HandleResults(newReply, err)
	body := fmt.Sprintf("some body updated")
	newReply.Body = &body
	updatedReply, _, err := client.Case.Reply.Update(createdCase.GetResourceId(), newReply)
	HandleResults(updatedReply, err)
}

func (e *Example) DeleteCaseReply(client *desk.Client) {
	cze := BuildSampleCase()
	createdCase, _, err := client.Case.Create(cze)
	HandleResults(createdCase, err)
	reply := desk.ReplyBuilder.
		SetString("Body", "some body").
		SetString("Direction", "out").
		SetString("Status", "draft").
		BuildReply()
	newReply, _, err := client.Case.Reply.Create(createdCase.GetResourceId(), &reply)
	HandleResults(newReply, err)
	resp, _ := client.Case.Reply.Delete(createdCase.GetResourceId(), newReply.GetResourceId())
	fmt.Printf("Delete results: %v\n", resp)
}

//-----------------------------------------------------------------------------
//Case Message
//-----------------------------------------------------------------------------
func (e *Example) GetCaseMessage(client *desk.Client) {
	cse, _, err := client.Case.Message.Get("1")
	HandleResults(cse, err)
}

func (e *Example) UpdateCaseMessage(client *desk.Client) {
	message := desk.MessageBuilder.
		SetString("Direction", "out").
		SetString("Status", "draft").
		SetString("To", "someone@desk.com").
		SetString("From", "someone-else@desk.com").
		SetString("Subject", "Case created by API via desk-go").
		SetString("Body", "Request for assistance denied").
		BuildMessage()
	caze := desk.CaseBuilder.
		SetString("Type", "email").
		SetString("Subject", "Case created by API via desk-go").
		SetInt("Priority", 4).
		SetString("Status", "received").
		SetMessage(message).
		AddHrefLink("customer", fmt.Sprintf("/api/v2/customers/%d", DefaultCustomerId)).
		BuildCase()
	newCase, _, err := client.Case.Create(&caze)
	HandleResults(newCase, err)
	updateMsg := desk.MessageBuilder.
		SetString("Subject", fmt.Sprintf("Case updated by API via desk-go at %v", time.Now())).
		BuildMessage()
	newMsg, _, err := client.Case.Message.Update(newCase.GetResourceId(), &updateMsg, nil)
	HandleResults(newMsg, err)
}

func (e *Example) DeleteCaseMessage(client *desk.Client) {
	message := desk.MessageBuilder.
		SetString("Direction", "out").
		SetString("Status", "draft").
		SetString("To", "someone@desk.com").
		SetString("From", "someone-else@desk.com").
		SetString("Subject", "Case created by API via desk-go").
		SetString("Body", "Request for assistance denied").
		BuildMessage()
	caze := desk.CaseBuilder.
		SetString("Type", "email").
		SetString("Subject", "Case created by API via desk-go").
		SetInt("Priority", 4).
		SetString("Status", "received").
		SetMessage(message).
		AddHrefLink("customer", fmt.Sprintf("/api/v2/customers/%d", DefaultCustomerId)).
		BuildCase()
	newCase, _, err := client.Case.Create(&caze)
	HandleResults(newCase, err)
	res, _ := client.Case.Message.Delete(newCase.GetResourceId())
	fmt.Printf("Delete results: %v\n", res)
}

//-----------------------------------------------------------------------------
//Case
//-----------------------------------------------------------------------------
func (e *Example) GetCase(client *desk.Client) {
	cse, _, err := client.Case.Get("1")
	HandleResults(cse, err)
}

func (e *Example) ListCase(client *desk.Client) {
	listParams := url.Values{}
	listParams.Add("sort_field", "created_at")
	listParams.Add("sort_direction", "asc")
	collection, _, err := client.Case.List(&listParams)
	HandleResults(collection, err)
}

func (e *Example) SearchCase(client *desk.Client) {
	searchParams := url.Values{}
	searchParams.Add("sort_field", "created_at")
	searchParams.Add("sort_direction", "asc")
	searchParams.Add("status", "new")
	collection, _, err := client.Case.Search(&searchParams, nil)
	HandleResults(collection, err)
}

func (e *Example) UpdateCase(client *desk.Client) {
	caze := desk.CaseBuilder.
		SetString("Subject", fmt.Sprintf("updated case at %v", time.Now())).
		SetInt("ID", 1).
		BuildCase()
	newCase, _, err := client.Case.Update(&caze)
	HandleResults(newCase, err)
}

func (e *Example) CreateCase(client *desk.Client) {
	caze := BuildSampleCase()
	newCase, _, err := client.Case.Create(caze)
	HandleResults(newCase, err)
}

func (e *Example) DeleteCase(client *desk.Client) {
	caze := BuildSampleCase()
	newCase, _, err := client.Case.Create(caze)
	HandleResults(newCase, err)
	results, err := client.Case.Delete(newCase.GetResourceId())
	fmt.Printf("Delete results: %v\n", results)
	foundCase, results, err := client.Case.Get(newCase.GetResourceId())
	HandleResults(foundCase, err)
}

func (e *Example) ForwardCase(client *desk.Client) {
	resp, _ := client.Case.Forward("1", "someone@desk.com", "some note")
	fmt.Printf("Forward results: %v\n", resp)
}

//-----------------------------------------------------------------------------
//Customers
//-----------------------------------------------------------------------------
func (e *Example) GetCustomer(client *desk.Client) {
	customer, _, err := client.Customer.Get(fmt.Sprintf("%d", DefaultCustomerId))
	HandleResults(customer, err)
}

func (e *Example) ListCustomer(client *desk.Client) {
	listParams := url.Values{}
	listParams.Add("sort_field", "created_at")
	listParams.Add("sort_direction", "asc")
	collection, _, err := client.Customer.List(&listParams)
	HandleResults(collection, err)
}

func (e *Example) SearchCustomer(client *desk.Client) {
	searchParams := url.Values{}
	searchParams.Add("sort_field", "created_at")
	searchParams.Add("sort_direction", "asc")
	searchParams.Add("max_id", "200000000")
	collection, _, err := client.Customer.Search(&searchParams, nil)
	HandleResults(collection, err)
}

func (e *Example) CreateCustomer(client *desk.Client) {
	firstName := "James"
	lastName := "Dean"
	customer := desk.Customer{FirstName: &firstName, LastName: &lastName}
	newCustomer, _, err := client.Customer.Create(&customer)
	HandleResults(newCustomer, err)
}

func (e *Example) UpdateCustomer(client *desk.Client) {
	id := DefaultCustomerId
	background := fmt.Sprintf("background updated at %v", time.Now())
	customer := desk.Customer{Background: &background}
	customer.Id = &id
	updatedCustomer, _, err := client.Customer.Update(&customer)
	HandleResults(updatedCustomer, err)
}

func (e *Example) CustomerCases(client *desk.Client) {
	params := url.Values{}
	params.Add("sort_field", "created_at")
	params.Add("sort_direction", "asc")
	page, _, err := client.Customer.Cases(fmt.Sprintf("%d", DefaultCustomerId), &params)
	HandleResults(page, err)
}
