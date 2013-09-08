package jsonv_test

import (
    "encoding/json"
    "log"
    j ".."
    "testing"
)

/*
Testing for type cross-matching (shouldn't be possible)
*/
func TestValidator(t *testing.T) {
    test("Number should not match String", false, []byte(`"abc"`), &j.Number{})
    test("Number should not match null", false, []byte(`null`), &j.Number{})
    
    test("Array should not match String", false, []byte(`"abc"`), &j.Array{})
    test("Array should not match null", false, []byte(`null`), &j.Array{})
    
    test("Object should not match String", false, []byte(`"abc"`), &j.Object{})
    test("Object should not match null", false, []byte(`null`), &j.Object{})
    
    test("Boolean should not match String", false, []byte(`"abc"`), &j.Boolean{})
    test("Boolean should not match null", false, []byte(`null`), &j.Boolean{})
    
    // String is left over. test it too
    test("String should not match Number", false, []byte(`123`), &j.String{})
    test("String should not match null", false, []byte(`null`), &j.String{})
}

/*
Helper function that's given description of the test to be performed, expectation of pass/fail, raw json and schema used to validate the json.

This function will output nothing if everything went as expected, but if not, it'll output detailed error and kill the process.
*/
func test(title string, expectPass bool, rawJson []byte, schema j.Validator) {
    decoded := new(interface{})
    if err := json.Unmarshal(rawJson, decoded); err != nil {
        log.Fatalf("'%s' failed. JSON parsing failed: %s", title, err)
    }
    
    path, err := schema.Validate(decoded)
    
    if expectPass && err != nil {
        log.Fatalf("'%s' failed (%s). Path: %s", title, err, path)
    } else if !expectPass && err == nil {
        log.Fatalf("'%s' failed. expected validation to fail, but it passed", title)
    }
}
