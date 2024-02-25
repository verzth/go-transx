package transx

import (
	"reflect"
)

func Transform(source any, dest any) error {
	sourceT := reflect.TypeOf(source)
	sourceV := reflect.ValueOf(source)
	destT := reflect.TypeOf(dest)
	destV := reflect.ValueOf(dest)
	var valMaps = make(map[string]reflect.Value)
	if sourceV.Kind() == reflect.Struct {
		for i := 0; i < sourceT.NumField(); i++ {
			srcField, ok := sourceT.FieldByName(sourceT.Field(i).Name)
			if ok {
				tag := srcField.Tag.Get("transx")
				if tag != "" && tag != "-" {
					valMaps[tag] = sourceV.Field(i)
				}
			}
		}
	} else if sourceV.Kind() == reflect.Ptr && sourceV.Elem().Kind() == reflect.Struct {
		for i := 0; i < sourceT.Elem().NumField(); i++ {
			srcField, ok := sourceT.Elem().FieldByName(sourceT.Elem().Field(i).Name)
			if ok {
				tag := srcField.Tag.Get("transx")
				if tag != "" && tag != "-" {
					valMaps[tag] = sourceV.Elem().Field(i)
				}
			}
		}
	} else {
		return ErrSrcNotStruct
	}
	if destV.Kind() != reflect.Ptr {
		return ErrDstNotPtrStruct
	} else {
		if destV.Elem().Kind() != reflect.Struct {
			return ErrDstNotPtrStruct
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

func TransformSlice(source any, dest any) error {
	if source == nil {
		return ErrSrcSliceNil
	}
	if dest == nil {
		return ErrDstSliceNil
	}
	sourceV := reflect.ValueOf(source)
	destV := reflect.ValueOf(dest)

	if destV.Kind() != reflect.Ptr {
		return ErrDstSliceNotPtr
	}

	if destV.Elem().Kind() != reflect.Slice {
		return ErrDstNotSlice
	}

	if sourceV.Kind() == reflect.Slice {
		for i := 0; i < sourceV.Len(); i++ {
			destR := reflect.New(destV.Elem().Type().Elem())
			err := Transform(sourceV.Index(i).Interface(), destR.Interface())
			if err != nil {
				return err
			}
			destV.Elem().Set(reflect.Append(destV.Elem(), destR.Elem()))
		}
	} else if sourceV.Kind() == reflect.Ptr && sourceV.Elem().Kind() == reflect.Slice {
		for i := 0; i < sourceV.Elem().Len(); i++ {
			var destR reflect.Value
			destC := destV.Elem().Type().Elem()
			if destC.Kind() == reflect.Ptr && destC.Elem().Kind() == reflect.Struct {
				destR = reflect.New(destC.Elem())
			} else {
				destR = reflect.New(destC)
			}
			err := Transform(sourceV.Elem().Index(i).Interface(), destR.Interface())
			if err != nil {
				return err
			}
			if destC.Kind() == reflect.Ptr {
				destV.Elem().Set(reflect.Append(destV.Elem(), destR))
			} else {
				destV.Elem().Set(reflect.Append(destV.Elem(), destR.Elem()))
			}
		}
	} else {
		return ErrSrcNotSlice
	}

	return nil
}