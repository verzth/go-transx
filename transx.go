package transx

import (
	"reflect"
)

func Transform(source any, dest any) error {
	sourceT := reflect.TypeOf(source)
	sourceV := reflect.ValueOf(source)
	destT := reflect.TypeOf(dest)
	destV := reflect.ValueOf(dest)
	if sourceV.Kind() != reflect.Struct {
		return ErrSrcNotStruct
	}
	if destV.Kind() != reflect.Ptr {
		return ErrDstNotPtrStruct
	} else {
		if destV.Elem().Kind() != reflect.Struct {
			return ErrDstNotPtrStruct
		}
	}

	var valMaps = make(map[string]reflect.Value)
	for i := 0; i < sourceT.NumField(); i++ {
		srcField, ok := sourceT.FieldByName(sourceT.Field(i).Name)
		if ok {
			tag := srcField.Tag.Get("transx")
			if tag != "" && tag != "-" {
				valMaps[tag] = sourceV.Field(i)
			}
		}
	}

	for i := 0; i < destT.Elem().NumField(); i++ {
		destField, ok := destT.Elem().FieldByName(destT.Elem().Field(i).Name)
		if ok {
			tag := destField.Tag.Get("transx")
			if tag != "" && tag != "-" {
				if val, ok := valMaps[tag]; ok {
					destV.Elem().Field(i).Set(val)
				}
			}
		}
	}

	return nil
}
