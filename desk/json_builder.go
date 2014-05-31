package desk

import (
  "time"
  "github.com/lann/builder"
)

type jsonBuilder builder.Builder

func (b jsonBuilder) SetString(field string,value string) jsonBuilder {
  return builder.Set(b, field, &value).(jsonBuilder)
}

func (b jsonBuilder) SetInt(field string,value int) jsonBuilder {
  return builder.Set(b, field, &value).(jsonBuilder)
}

func (b jsonBuilder) SetTimestamp(field string,value Timestamp) jsonBuilder {
  return builder.Set(b, field, &value).(jsonBuilder)
}

func (b jsonBuilder) SetTimestampNow(field string) jsonBuilder {
  timet:=Timestamp{time.Now()}
  return builder.Set(b, field, &timet).(jsonBuilder)
}

func (b jsonBuilder) SetLinkCollection(value LinkCollection) jsonBuilder {
  return builder.Set(b, "LinkCollection", value).(jsonBuilder)
}

func (b jsonBuilder) SetMessage(value Message) jsonBuilder {
  return builder.Set(b, "Message", &value).(jsonBuilder)
}

func (b jsonBuilder) AddHrefLink(class string,href string) jsonBuilder {
  val,_ := builder.Get(b, "LinkCollection")
  if val==nil {
    val = LinkCollection{}
  }
  coll,_ := val.(LinkCollection)
  coll.AddHrefLink(class,href)
  return builder.Set(b,"LinkCollection",coll).(jsonBuilder)
}

func (b jsonBuilder) AddCustomField(name string,value interface{}) jsonBuilder {
  val,_ := builder.Get(b,"CustomFields")
  if val==nil {
    val = make(map[string]interface{})
  }
  fields,_ := val.(map[string]interface{})
  fields[name]=value
  return builder.Set(b,"CustomFields",fields).(jsonBuilder)
}

func (b jsonBuilder) BuildCustomer() Customer {
  return builder.GetStructLike(b, Customer{}).(Customer)
}

func (b jsonBuilder) BuildMessage() Message {
  return builder.GetStructLike(b, Message{}).(Message)
}

func (b jsonBuilder) BuildCase() Case {
  return builder.GetStructLike(b, Case{}).(Case)
}

func (b jsonBuilder) BuildLinkCollection() LinkCollection {
  return builder.GetStructLike(b, LinkCollection{}).(LinkCollection)
}

var CaseBuilder = builder.Register(jsonBuilder{}, Case{}).(jsonBuilder)
var MessageBuilder = builder.Register(jsonBuilder{}, Message{}).(jsonBuilder)
var CustomerBuilder = builder.Register(jsonBuilder{}, Customer{}).(jsonBuilder)
var LinkCollectionBuilder = builder.Register(jsonBuilder{}, LinkCollection{}).(jsonBuilder)

