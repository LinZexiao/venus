package typeutil

import (
	"fmt"
	"go/ast"
	"reflect"
	"sync"
)

var exportedFieldsCache = struct {
	sync.RWMutex
	fields map[reflect.Type][]reflect.StructField
}{
	fields: make(map[reflect.Type][]reflect.StructField),
}

func ExportedFields(obj interface{}) ([]reflect.StructField, error) {
	typ, ok := obj.(reflect.Type)
	if !ok {
		typ = reflect.TypeOf(obj)
	}

	if kind := typ.Kind(); kind != reflect.Struct {
		return nil, fmt.Errorf("%s is not struct", typ)
	}

	exportedFieldsCache.RLock()
	fields, ok := exportedFieldsCache.fields[typ]
	exportedFieldsCache.RUnlock()

	if ok {
		return fields, nil
	}

	num := typ.NumField()
	fields = make([]reflect.StructField, 0, num)
	for i := 0; i < num; i++ {
		field := typ.Field(i)
		if !ast.IsExported(field.Name) {
			continue
		}

		fields = append(fields, field)
	}

	exportedFieldsCache.Lock()
	exportedFieldsCache.fields[typ] = fields
	exportedFieldsCache.Unlock()

	return fields, nil
}
