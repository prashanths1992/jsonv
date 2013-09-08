package jsonv_test

import (
    "testing"
    j ".."
)

func TestString(t *testing.T) {
    test("abc is String", true,
        []byte(`"abc"`),
        &j.String{},
    )
        
    test("abc is String, MinLen:1", true,
        []byte(`"abc"`),
        &j.String{MinLen:1},
    )
    
    test("abc is String, MinLen:4", false,
        []byte(`"abc"`),
        &j.String{MinLen:4},
    )
    
    test("abc is String, MaxLen:3", true,
        []byte(`"abc"`),
        &j.String{MaxLen:3},
    )
    
    test("abc is String, MaxLen:2", false,
        []byte(`"abc"`),
        &j.String{MaxLen:2},
    )
    
    test("abc is String, MinLen:2, MaxLen:4", true,
        []byte(`"abc"`),
        &j.String{MinLen:2, MaxLen:4},
    )
    
    test("abc is String, MinLen:4, MaxLen:1", false,
        []byte(`"abc"`),
        &j.String{MinLen:4, MaxLen:1},
    )
    
    test("abc is String, Content:abc", true,
        []byte(`"abc"`),
        &j.String{Value:"abc"},
    )
    
    test("abc is String, Content:asd", false,
        []byte(`"abc"`),
        &j.String{Value:"asd"},
    )
    
    test("abc is String, Value:abc, MinLen:1, MaxLen:3", true,
        []byte(`"abc"`),
        &j.String{Value:"abc", MinLen:1, MaxLen:3},
    )
}
