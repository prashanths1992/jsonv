package jsonv

import (
    // "fmt"
    // "reflect"
)

// /*
// Ensures that the data is (json) array, with optional length checks and validation of elements. Value in `Each` must be present

// Note: Value 0 in either `AtLeast` or `AtMost` is ignored.
// */
type Array struct {
    AtLeast int
    AtMost int
    Each Validator
}


func (self Array) Validate(raw interface{}) (string, error) {
    /*switch tmp := raw.(type) {
    case *interface{}:
        if (tmp != nil) { return self.Validate(*tmp) }
    case []interface{}:
        return self._validate(&tmp)
    case *[]interface{}:
        if (tmp != nil) { return self._validate(tmp) }
    }
    
    return "Array", fmt.Errorf("expected []interface{}, was %v", reflect.TypeOf(raw))
    */
    return "Array", nil
}
/*
func (self *Array) _validate(unsung *[]interface{}) (path string, err error) {

    if self.AtLeast != 0 {
        if len(*unsung) < self.AtLeast { return "Array", fmt.Errorf("length should >=%d, was %d", self.AtLeast, len(*unsung)) }
    }
    
    if self.AtMost != 0 {
        if len(*unsung) > self.AtMost { return "Array", fmt.Errorf("length should <=%d, was %d", self.AtMost, len(*unsung)) }
    }
    
    // cannot range iterate, since wish is to avoid copying data
    if (self.Each != nil) { for i := 0; i < len(*unsung); i++ {
        if desc, err := self.Each.Validate(&(*unsung)[i]); err != nil {
            return fmt.Sprintf("Array[%d]->%s", i, desc), err
        }
    }}
    
    return "Array", nil
}
*/