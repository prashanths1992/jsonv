package jsonv_test

import (
    "log"
    "encoding/json"
    j "."
    "fmt"
    "reflect"
    "testing"
)

/*
    $ go test ./src/...
Ensure that the example passes (log.Fatal kills the process -> test fails)
*/
func TestValidator(t *testing.T) {
    ExampleValidator()
}

func ExampleValidator() {
    // import j "github.com/gima/jsonv/src"
    
    rawJson := []byte(`
        {
            "status": true,
            "data": {
                "token": "CAFED00D",
                "debug": 69306,
                "items": [
                    { "url": "https://news.ycombinator.com/", "comment": "why wont people share?" },
                    { "url": "http://golang.org/", "comment": "some darn gophers" },
                    { "url": "http://www.kickstarter.com/", "comment": "\"opensource\" projects. yeah, right.." }
                ],
                "ghost2": null,
                "meta": {
                    "g": 1,
                    "xyzzy": 0.25,
                    "blöö": 0.5
                }
            }
        }
    `)
    
    decoded := new(interface{})
    if err := json.Unmarshal(rawJson, decoded); err != nil {
        log.Fatal("JSON parsing failed:", err)
    }
    
    myValidatorFunc := func(data *interface{}) (desc string, err error) {
        desc = "myValidatorFunc"
        if validate, ok := (*data).(string); !ok {
            err = fmt.Errorf("expected string, was %v", reflect.TypeOf(*data))
        } else {
            if validate != "CAFED00D" { err = fmt.Errorf("expected CAFED00D, was %s", validate)
            } else { err = nil }
        }
        return
    }
    
    schema := &j.Object{Properties:[]j.ObjectItem{
        {"status", &j.Boolean{}},
        {"data", &j.Object{Properties:[]j.ObjectItem{
            {"token", &j.Function{myValidatorFunc}},
            {"debug", &j.Number{Min:1, Max:99999}},
            {"items", &j.Array{Each:&j.Object{Properties:[]j.ObjectItem{
                {"url", &j.String{MinLen:1}},
                {"comment", &j.Optional{&j.String{}}},
            }}}},
            {"ghost", &j.Optional{&j.String{}}},
            {"ghost2", &j.Optional{&j.String{}}},
            {"meta", &j.Object{Each:j.ObjectEach{
                &j.String{}, &j.Or{&j.Number{Min:.01, Max:1.1}, &j.String{}},
            }}},
        }}},
    }}
    
    if path, err := schema.Validate(decoded); err == nil {
        log.Println("OK!")
    } else {
        log.Fatalf("Failed (%s). Path: %s", err, path)
    }
}
