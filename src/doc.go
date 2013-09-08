/*
Validate data produced by `json.Unmarshal`. Data given to validation functions, but not found in schema, passes validation without inspection (i.e. is ignored).

Creating a schema looks like this:

    schema := &j.Object{Properties:[]j.ObjectItem{
        {"status", &j.Boolean{}},
        {"data", &j.Object{Properties:[]j.ObjectItem{
            {"token", &j.Function{myValidatorFunc}},
            {"debug", &j.Number{Min:1, Max:99999}},
            {"items", &j.Array{Each:&j.Object{Properties:[]j.ObjectItem{
                {"url", &j.String{MinLen:1}},
                {"comment", &j.Optional{&j.String{}}},
            }}}},
            {"ghost", &j.Optional{&j.String{}}},
            {"ghost2", &j.Optional{&j.String{}}},
            {"meta", &j.Object{Each:j.ObjectEach{
                &j.String{}, &j.Or{&j.Number{Min:.01, Max:1.1}, &j.String{}},
            }}},
        }}},
    }}

Validation using the schema looks like this:

    if path, err := schema.Validate(decoded); err == nil {
        log.Println("OK!")
    } else {
        log.Fatalf("Failed (%s). Path: %s", err, path)
    }

And an example of an error produced from failed validation with the above code looks like this:

    // Failed (expected float64, was bool). Path: Object(Items, key("data").Value)->Object(Items, key("debug").Value)->Number
*/
package jsonv
