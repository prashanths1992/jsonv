package jsonv

import (
    "fmt"
    "reflect"
)

/*
Return reflection to `in`, unless `in` is a pointer, in which case return reflection to indirection of `in`

Errors if `in` == `nil` or `*in` == `nil`
*/
func indirectOrNil(in interface{}) (error, reflect.Value, reflect.Type) {
    return indirectOrPtrOrNil(in, false)
}

/*
Return reflection to pointer of `in`'s copy, unless `in` is a pointer, in which case return reflection to it.

Errors if `in` == `nil` or `*in` == `nil`
*/
func ptrOrNil(in interface{}) (error, reflect.Value, reflect.Type) {
    return indirectOrPtrOrNil(in, true)
}

func indirectOrPtrOrNil(in interface{}, wantPtr bool) (error, reflect.Value, reflect.Type) {
    if (in == nil) {
        return fmt.Errorf("nil"), reflect.Value{}, nil
    }
    
    rawValue := reflect.ValueOf(in)
    rawType := rawValue.Type()
    
    if (wantPtr) {
        // output pointer
        switch rawType.Kind() {
        
        case reflect.Ptr:
            // return reflection to `in`
            return nil, rawValue, rawType
        
        default:
            // return reflection to address of a copy of `in`
            ptr := reflect.New(rawType)
            ptrValue := ptr.Elem()
            ptrValue.Set(rawValue)
            return nil, ptrValue, ptrValue.Type()
        }
        
    } else {
        // output !pointer
        switch rawType.Kind() {
        
        case reflect.Ptr:
            // return reflection to indirection of `n`
            if rawValue.IsNil() {
                return fmt.Errorf("nil pointer (%s)", rawType), reflect.Value{}, nil
            } else {
                rawValue = rawValue.Elem()
                rawType = rawValue.Type()
                return nil, rawValue, rawType
            }
            
        default:
            // return reflection to `in`
            return nil, rawValue, rawType
        }
        
    }
}
