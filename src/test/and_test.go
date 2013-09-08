package jsonv_test

import (
    "testing"
    "fmt"
    "reflect"
    j ".."
)

func TestAnd(t *testing.T) {
    test("asd is String, first letter 'a' AND asd is String, second letter 's'", true,
        []byte(`"asd"`),
        &j.And{&j.Function{validationFunc2(0, 'a')}, &j.Function{validationFunc2(1, 's')}},
    )
    
    test("xsd is String, first letter 'a' AND asd is String, second letter 's'", false,
        []byte(`"xsd"`),
        &j.And{&j.Function{validationFunc2(0, 'a')}, &j.Function{validationFunc2(1, 's')}},
    )
    
    test("axd is String, first letter 'a' AND asd is String, second letter 's'", false,
        []byte(`"axd"`),
        &j.And{&j.Function{validationFunc2(0, 'a')}, &j.Function{validationFunc2(1, 's')}},
    )
}

func validationFunc2(pos int, value uint8) (func(*interface{}) (string, error)) {
    return func(data *interface{}) (desc string, err error) {
        desc = "validationFunc2"
        
        if validate, ok := (*data).(string); ok {
            if validate[pos] == value { err = nil
            } else { err = fmt.Errorf("expected letter at pos %d to be \"%c\" but it was \"%c\"", pos, value, validate[pos]) }
        } else {
            err = fmt.Errorf("expected *string, got %v", reflect.TypeOf(*data))
        }
        
        return
    }
}
