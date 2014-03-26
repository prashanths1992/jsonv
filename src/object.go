package jsonv

import (
    "fmt"
    "reflect"
    "container/list"
)

type Object map[interface{}]Validator

func (self *Object) Validate(raw interface{}) (string, error) {
    err, unsungValue, unsungType := indirectOrNil(raw)
    if err != nil {
        return "Object", fmt.Errorf("expected map or struct, or pointer to either. was %s", err)
    }
    
    // switch data-to-be-validated type
    switch unsungType.Kind() {
    
    case reflect.Map:
        return self.validateMap(unsungValue, unsungType)
        
    // case reflect.Struct:
        // return validateStruct(unsungValue, unsungType)
        
    default:
        return "Object", fmt.Errorf("expected map or struct. was %s", unsungType)
    }
}

func (self *Object) validateMap(unsungValue reflect.Value, unsungType reflect.Type) (string, error) {
    
    selfValue := reflect.ValueOf(*self)
    selfType := selfValue.Type()
    
    // list of schema keys that are used to validate unsung keys that have not been visited
    schemaValidatorKeys := list.New()
    // list of unsung keys that were handled by schema non-validator keys
    unsungVisitedKeys := make(map[string]bool)
    
    // types used to determine assignability
    validatorType := reflect.TypeOf((*Validator)(nil)).Elem()
    stringType := reflect.TypeOf((*String)(nil)).Elem()
    optionalType := reflect.TypeOf((*Optional)(nil)).Elem()
    
    // validate with schema[key] Validator each corresponding unsung[key]
    for _, schemaKeyValueRaw := range selfValue.MapKeys() {
        err, schemaKeyValue, schemaKeyType := indirectOrNil(schemaKeyValueRaw)
        
        if schemaKeyType.Implements(validatorType) {
            // schema key is Validator. stash for later handling in the next (level-1) loop
            schemaValidatorKeys.PushBack(schemaKeyValue)
            
        } else if schemaKeyType.AssignableTo(stringType) {
            // schema key is String
            
            // check if exists: unsung[schema key]
            if
            
        }
        // switch schemaKey := pSchemaKey.(type) {
        // case *Validator:
            // schemaValidatorKeys.PushBack(key)
            
        // case *String:
            // if unsungKeyValue := pUnsungValue.MapIndex(); unsungKeyValue == Value{} {
                // return fmt.Sprintf("Object(Key:%v)", *schemaKey), fmt.Errorf("doesn't exist", err)
            // } else {
                // unsungKey
            // }
        // }
        // if type(key) == Validator {
            // schemaValidatorKeys.push(key)
        // } else if type(key) == String {
            // if dtv[key] exists {
                // if err = schema[key].Validate(dtv[key]); err {
                    // // object (key:%key): err
                // }
                // dtvKeysVisited[key] = true
            // } else {
                // // object (key:%key): not found
            // }
        // } else {
            // // object (key:%key): expected string or validator
        // }
    }
    
    // for dataKey = dtv keys {
        // if dtvKeysVisited[key] { continue }
        // //loop through left-over schema validator keys (jsonv.Object doc prefers only one validator-key, lest n^n(?) time):
        // for schemaValidatorKey = schema keys {
            // if err = schema[schemaValidatorKey].Validate(dtv[dataKey]); err {
                // // object (key:%key): err
            // }
        // }
    // }
    
    /*
    
    
    fmt.Println("IS map")
    fmt.Println("keytype:", (*unsungType).Key())
    fmt.Println("elemtype:", (*unsungType).Elem())
    */
    
    // --
    //var unsungKeysNotInSchema = every dtv key
    //loop through schema(self) as key
        //if TypeOf key == string:
            //ensure validator:[key] validates data:dtv[key]
            //remove key from unsungKeysNotInSchema
        //else if TypeOf key == validator, save to validatorKeys
        //else error: unknown schema key type
    //loop through unsungKeysNotInSchema as keyNE
        //loop through validatorKeys as keyV (jsonv.Object doc prefers only one validator-key):
            // ensure keyV validates keyNIS && [keyNIS] validates [keyV]
    // --
    
    //var unsungKeysNotInSchema = every dtv key
    /*
    unsungKeysNotInSchema := make(map[string]bool)
    for _, unsungKeyValue := range unsungValue.MapKeys() {
        if err, key := isString(&unsungKeyValue); err != nil {
            return fmt.Sprintf("Object(Key:%v)", unsungKeyValue), fmt.Errorf("expected string. was %s", err)
        } else {
            unsungKeysNotInSchema[key] = true
        }
    }
    
    //loop through schema(self) as key
    for _, rawSchemaKey := range selfValue.MapKeys() {
        if err, rawSchemaKeyPtr, _, _ := ptrOrNil(rawSchemaKey.Interface()); err != nil {
            return "Object(Schema Key)", fmt.Errorf("expected string or validator, or pointer to either. was %s", err)
        } else {
            switch schemaKeyPtr := rawSchemaKeyPtr.(type) {
                case *String:
                    rawSchemaKeyValue := self.MapIndex(rawSchemaKey)
                    //key exists, there's no doubt
                    if rawSchemaKeyValue.type
                    
                    unsungValue
                    self[*schemaKeyPtr]()
                    //ensure validator:[key] validates data:dtv[key]
                    //remove key from unsungKeysNotInSchema
                case *Validator:
                    // save to validatorKeys
                default:
                    // error: unknown schema key type
            }
        }
    }
    */
    
    return "",nil
    
    /*
    fmt.Println("IS map")
    fmt.Println("keytype:", rawType.Key())
    fmt.Println("elemtype:", rawType.Elem())
    // loop schema
    for mapKey, mapValue := range self {
    switch keyValue := mapKey.(type) {
        case string:
        case Validator:
        default:
            _ = keyValue
            return fmt.Sprintf("Object(Map Property:%v)", mapKey), fmt.Errorf("expected string or Validator, was %v", reflect.TypeOf(mapKey))
        }
        _ = mapValue
        }
    }
    */
}

func isString(mapKey *reflect.Value) (error, string) {
    kind := mapKey.Kind()
    if kind != reflect.String {
        return fmt.Errorf("%s", kind), ""
    } else {
        return nil, mapKey.Interface().(string)
    }
}

// func validateStruct(unsungValue reflect.Value, unsungType reflect.Type) (string, error) {
    // fmt.Println("Object.Validate struct. NOT IMPLEMENTED YET")
    // return "", nil
// }

type ExactObject Object
func (self ExactObject) Validate(data interface{}) (string, error) {
    fmt.Println("ExactObject.Validate NOT IMPLEMENTED YET")
    return "ExactObject", nil
}

