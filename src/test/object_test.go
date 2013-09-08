package jsonv_test

import (
    "testing"
    j ".."
)

func TestObject(t *testing.T) {
    test("{} is Object", true,
        []byte(`{}`),
        &j.Object{},
    )
    
    test("{a:5,b:1} is Object, Each:String,Number", true,
        []byte(`{"a":5,"b":1}`),
        &j.Object{Each:j.ObjectEach{
            &j.String{}, &j.Number{},
        }},
    )
    
    test("{a:5,b:1} is Object, Each:Number,String", false,
        []byte(`{"a":5,"b":1}`),
        &j.Object{Each:j.ObjectEach{
            &j.Number{}, &j.String{},
        }},
    )
    
    test("{a:c,b:d} is Object, Each:String,Number", false,
        []byte(`{"a":"c","b":"d"}`),
        &j.Object{Each:j.ObjectEach{
            &j.String{}, &j.Number{},
        }},
    )
    
    test("{a:5,b:c} is Object, Properties:a=Number,b=String", true,
        []byte(`{"a":5,"b":"c"}`),
        &j.Object{Properties:[]j.ObjectItem{
            {"a", &j.Number{}},
            {"b", &j.String{}},
        }},
    )
    
    test("{a:5,b:c} is Object, Properties:a=String,b=Number", false,
        []byte(`{"a":5,"b":"c"}`),
        &j.Object{Properties:[]j.ObjectItem{
            {"a", &j.String{}},
            {"b", &j.Number{}},
        }},
    )
    
    test("{a:null} is Object, Properties:a=String", false,
        []byte(`{"a":null}`),
        &j.Object{Properties:[]j.ObjectItem{
            {"a", &j.String{}},
        }},
    )
    
    test("{} is Object, Properties:a=String", false,
        []byte(`{}`),
        &j.Object{Properties:[]j.ObjectItem{
            {"a", &j.String{}},
        }},
    )
}
