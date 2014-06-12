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

func (b jsonBuilder) SetLinks(value map[string]map[string]interface{}) jsonBuilder {
  return builder.Set(b, "Links", value).(jsonBuilder)
}

func (b jsonBuilder) SetMessage(value Message) jsonBuilder {
  return builder.Set(b, "Message", &value).(jsonBuilder)
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

func (b jsonBuilder) AddEmail(value string,valueType string) jsonBuilder {
  customer:=builder.GetStructLike(b,Customer{}).(Customer)
  customer.AddEmail(value,valueType)
  return builder.Set(b,"Emails",customer.Emails).(jsonBuilder)
}

func (b jsonBuilder) AddAddress(value string,valueType string) jsonBuilder {
  customer:=builder.GetStructLike(b,Customer{}).(Customer)
  customer.AddAddress(value,valueType)
  return builder.Set(b,"Addresses",customer.Addresses).(jsonBuilder)
}

func (b jsonBuilder) AddPhoneNumber(value string,valueType string) jsonBuilder {
  customer:=builder.GetStructLike(b,Customer{}).(Customer)
  customer.AddPhoneNumber(value,valueType)
  return builder.Set(b,"PhoneNumbers",customer.PhoneNumbers).(jsonBuilder)
}

func (b jsonBuilder) AddHrefLink(class string,href string) jsonBuilder {
  val,_ := builder.Get(b, "Hal")
  if val==nil {
    val = Hal{}
  }
  hal,_ := val.(Hal)
  hal.AddHrefLink(class,href)
  return builder.Set(b,"Hal",hal).(jsonBuilder)
}

func (b jsonBuilder) BuildCustomer() Customer {
  return builder.GetStructLike(b, Customer{}).(Customer)
}

func (b jsonBuilder) BuildMessage() Message {
  return builder.GetStructLike(b, Message{}).(Message)
}

func (b jsonBuilder) BuildReply() Reply {
  return builder.GetStructLike(b, Reply{}).(Reply)
}

func (b jsonBuilder) BuildCase() Case {
  return builder.GetStructLike(b, Case{}).(Case)
}

var CaseBuilder = builder.Register(jsonBuilder{}, Case{}).(jsonBuilder)
var MessageBuilder = builder.Register(jsonBuilder{}, Message{}).(jsonBuilder)
var ReplyBuilder = builder.Register(jsonBuilder{}, Reply{}).(jsonBuilder)
var CustomerBuilder = builder.Register(jsonBuilder{}, Customer{}).(jsonBuilder)

