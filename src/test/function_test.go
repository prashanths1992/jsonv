package jsonv_test

import (
    "testing"
    "fmt"
    "reflect"
    j ".."
)

func TestFunction(t *testing.T) {
    test("aSd is String, second letter is S", true,
        []byte(`"aSd"`),
        &j.Function{validationFunc1},
    )
    
    test("asd is String, second letter is S", false,
        []byte(`"asd"`),
        &j.Function{validationFunc1},
    )
}

func validationFunc1(data *interface{}) (desc string, err error){
    desc = "validationFunc1"
    
    if validate, ok := (*data).(string); ok {
        if validate[1] == 'S' { err = nil
        } else { err = fmt.Errorf("expected second letter to be \"S\" but it was \"%c\"", validate[1]) }
    } else {
        err = fmt.Errorf("expected *string, got %v", reflect.TypeOf(*data))
    }
    
    return
}
