package jsonv_test

import (
    "testing"
    j ".."
)

func TestOr(t *testing.T) {
    test("asd is String, Value:asd OR String, value:qwe", true,
        []byte(`"asd"`),
        &j.Or{&j.String{Value:"asd"}, &j.String{Value:"qwe"}},
    )
    
    test("asd is String, Value:asd OR String, value:qwe", true,
        []byte(`"qwe"`),
        &j.Or{&j.String{Value:"asd"}, &j.String{Value:"qwe"}},
    )
    
    test("asd is String, Value:asd OR j.Number", false,
        []byte(`"qwe"`),
        &j.Or{&j.String{Value:"asd"}, &j.Number{}},
    )
}
