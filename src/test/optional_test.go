package jsonv_test

import (
    "testing"
    j ".."
)

func TestOptional(t *testing.T) {
    test("true is Boolean and present", true,
        []byte(`true`),
        &j.Optional{&j.Boolean{}},
    )
    
    test("asd is not Boolean and present", false,
        []byte(`"asd"`),
        &j.Optional{&j.Boolean{}},
    )
    
    /*
    test("not present (really)", true,
        []byte(``),
        &j.Optional{&Boolean{}},
    )
    */
    
    test("not present (null)", true,
        []byte(`null`),
        &j.Optional{&j.Boolean{}},
    )
}
