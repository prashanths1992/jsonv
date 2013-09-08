package jsonv_test

import (
    "testing"
    j ".."
)

func TestBoolean(t *testing.T) {
    test("true is Boolean", true,
        []byte(`true`),
        &j.Boolean{},
    )
    
    test("false is Boolean", true,
        []byte(`false`),
        &j.Boolean{},
    )
}
