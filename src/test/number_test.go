package jsonv_test

import (
    "testing"
    j ".."
)

func TestNumber(t *testing.T) {
    test("123 is Number", true,
        []byte(`123`),
        &j.Number{},
    )
    
    test("1.23 is Number", true,
        []byte(`1.23`),
        &j.Number{},
    )
    
    test("123 is Number, Min:1", true,
        []byte(`123`),
        &j.Number{Min:1},
    )
    
    test("123 is Number, Min:124", false,
        []byte(`123`),
        &j.Number{Min:124},
    )
    
    test("123 is Number, Max:124", true,
        []byte(`123`),
        &j.Number{Max:124},
    )
    
    test("123 is Number, Max:123", false,
        []byte(`123`),
        &j.Number{Max:122},
    )
    
    test("123 is Number, Min:1, Max:123", true,
        []byte(`123`),
        &j.Number{Min:1, Max:123},
    )
    
    test("123 is Number, Min:124, Max:122", false,
        []byte(`123`),
        &j.Number{Min:124, Max:122},
    )
}
