package helpers

import "reflect"

func ObjectAssign(target interface{}, object interface{}) {
    tVal := reflect.ValueOf(target).Elem()
    oVal := reflect.ValueOf(object).Elem()
    
    for i := 0; i < oVal.NumField(); i++ {
        srcField := oVal.Type().Field(i)
        srcFieldName := srcField.Name
        
        tgtField := tVal.FieldByName(srcFieldName)
        if !tgtField.IsValid() {
            continue
        }
        
        if srcField.Type != tgtField.Type() {
            continue 
        }
        
        if tgtField.CanSet() {
            tgtField.Set(oVal.Field(i))
        }
    }
}