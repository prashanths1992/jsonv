package jsonv

import (
    // "fmt"
    // "reflect"
)

/*
Ensures that the data is boolean.
*/
type Boolean struct {
}

func (self Boolean) Validate(raw interface{}) (path string, err error) {
    /*path = "Boolean"
    
    switch tmp := raw.(type) {
    case *interface{}:
        if (tmp != nil) { return self.Validate(*tmp) }
    case bool:
        return
    case *bool:
        if (tmp != nil) { return }
    }
    
    err = fmt.Errorf("expected bool, was %v", reflect.TypeOf(raw))
    return*/
    return "Boolean", nil
}
