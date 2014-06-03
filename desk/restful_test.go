package desk

import (
  "fmt"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestRestful(t *testing.T) {
  fmt.Println("")
  Convey("Post", t, func() {
    Convey("should build fields",func(){
      r:=Restful{}
      r.Post("/path")
      So(r.path,ShouldEqual,"/path")
      So(r.method,ShouldEqual,"POST")
    })
  })
  Convey("Get", t, func() {
    Convey("should build fields",func(){
      r:=Restful{}
      r.Get("/path")
      So(r.path,ShouldEqual,"/path")
      So(r.method,ShouldEqual,"GET")
    })
  })
  Convey("Put", t, func() {
    Convey("should build fields",func(){
      r:=Restful{}
      r.Put("/path")
      So(r.path,ShouldEqual,"/path")
      So(r.method,ShouldEqual,"PUT")
    })
  })
  Convey("Patch", t, func() {
    Convey("should build fields",func(){
      r:=Restful{}
      r.Patch("/path")
      So(r.path,ShouldEqual,"/path")
      So(r.method,ShouldEqual,"PATCH")
    })
  })
  Convey("Delete", t, func() {
    Convey("should build fields",func(){
      r:=Restful{}
      r.Delete("/path")
      So(r.path,ShouldEqual,"/path")
      So(r.method,ShouldEqual,"DELETE")
    })
  })
}

