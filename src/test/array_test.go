package jsonv_test

import (
    "testing"
    j ".."
)

func TestArray(t *testing.T) {
    test("[] is Array", true,
        []byte(`[]`),
        &j.Array{},
    )
    
    test("[0] is Array, AtLeast:1", true,
        []byte(`[0]`),
        &j.Array{AtLeast:1},
    )
    
    test("[0] is Array, AtLeast:2", false,
        []byte(`[0]`),
        &j.Array{AtLeast:2},
    )
    
    test("[0, 0] is Array, AtMost:2", true,
        []byte(`[0, 0]`),
        &j.Array{AtMost:2},
    )
    
    test("[0, 0] is Array, AtMost:1", false,
        []byte(`[0, 0]`),
        &j.Array{AtMost:1},
    )
    
    test("[0, 0] is Array, AtMost:1", false,
        []byte(`[0, 0]`),
        &j.Array{AtMost:1},
    )
    
    test("[0, 0] is Array, Each:Number{}", true,
        []byte(`[0, 0]`),
        &j.Array{Each:&j.Number{}},
    )
    
    test("[0, a] is Array, Each:Number{}", false,
        []byte(`[0, "a"]`),
        &j.Array{Each:&j.Number{}},
    )
}
