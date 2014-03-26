package jsonv

import (
    "fmt"
)

/*
Validates if `raw` or (pointer only) `*raw` is `nil`, otherwise returns result of `V.Validate(raw)`
*/
type Optional struct {
    V Validator
}

func (self *Optional) Validate(raw interface{}) (string, error) {
    
    if err, unsungValue, _ := ptrOrNil(raw); err != nil {
        // as per ptrOrNil's documentation, the error occurs only if `raw` or (if pointer) `*raw` is `nil`,
        // which Optional is supposed to validate
        return "Optional(Not present)", nil
        
    } else {
        path, err := self.V.Validate(unsungValue.Interface())
        return fmt.Sprintf("Optional->%s", path), err
    }
    
}
