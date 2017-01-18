package helpers

import (
	"reflect"
	"regexp"
	"strings"
	"time"
	"errors"
)

func WalkinMap(dataMap map[string]interface{}, result interface{}) error {
	for k, v := range dataMap {

		if value, found := v.(map[string]interface{});found {
			return WalkinMap(value, result)
		}

		iFieldVal, i, err := FetchField("json", k, result)
		if err != nil {
			continue
		}
		fieldVal, found := iFieldVal.(reflect.Value)
		if !found {
			continue
		}
		fieldVal.Field(i).Set(reflect.ValueOf(v))
	}

	return nil
}

func SetField(name, tag string, value interface{}, t reflect.Type, v reflect.Value) {
	for i:=0;i < t.NumField(); i++ {
		field := t.Field(i)
		var dest reflect.Value
		if name == FieldName(tag, field) {
			switch field.Type.Kind() {
			case reflect.String:
				dest = reflect.ValueOf(value)
				v.Field(i).Set(dest)
			case reflect.Uint:
				tmp, found := value.(uint)
				if !found {
					continue
				}
				dest = reflect.ValueOf(tmp)
				v.Field(i).Set(dest)
			case reflect.Int:
				tmp, found := value.(int)
				if !found{
					continue
				}
				dest = reflect.ValueOf(tmp)
				v.Field(i).Set(dest)
			case reflect.Struct:
				if v.Field(i).Type().Name() == "Time" {
					tmp, found := value.(time.Time)
					if !found {
						continue
					}
					dest = reflect.ValueOf(tmp)
					v.Field(i).Set(dest)
				} else {
					v.Field(i).Set(reflect.ValueOf(value))
				}
			}
		}

	}
}

func Map2Struct(data map[string]interface{}, tag string, output interface{}) {
	val := reflect.Indirect(reflect.ValueOf(output))
	t := val.Type()

	if val.Kind() == reflect.Struct {
		for k, v := range data {
			SetField(k, tag, v, t, val)
		}
	}
}

func Struct2Map(data interface{}, tag string) (map[string]interface{}, error) {
	val := reflect.Indirect(reflect.ValueOf(data))
	t := val.Type()

	if val.Kind() == reflect.Struct {
		result := make(map[string]interface{})
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldVal := reflect.Indirect(val.Field(i))
			value := fieldVal.Interface()
			if !isZero(fieldVal, value) {
				fieldStr := FieldName(tag, field)
				result[fieldStr] = value
			}
		}
		return result, nil
	} else {
		return nil, errors.New("Submitted type must in struct type")
	}
}

func isZero(v reflect.Value, dv interface{}) bool {

	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i), v)
		}
		return z
	case reflect.Struct:
		z := true
		if v.Type().String() == "time.Time" {
			zeroTime := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
			comparedTime, found := v.Interface().(time.Time)
			if found {
				return zeroTime.Equal(comparedTime)
			}
		}
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i), v)
		}
		return z
	}

	// Other types is compared directly
	z := reflect.Zero(v.Type())
	return dv == z.Interface()
}

func FieldName(tag string, field reflect.StructField) string {
	if t := field.Tag.Get(tag); t != "" {
		return t
	}
	return field.Name
}

func FetchField(tag string, name string, dStruct interface{}) (interface{}, int, error) {
	ref := reflect.Indirect(reflect.ValueOf(dStruct))
	t := ref.Type()

	if ref.Kind() == reflect.Struct {
		tmp := ref
		for i:=0; i < t.NumField(); i++ {
			field := t.Field(i)

			fieldVal := reflect.Indirect(ref.Field(i))

			t := field.Tag.Get(tag)

			if (t != "" && t == name) || (field.Name == name) {
				return tmp, i, nil
			}

			if fieldVal.Kind() == reflect.Struct && (field.Type.Name() != "Time") {
				return FetchField(tag, name, fieldVal.Interface())
			}
		}
	}

	return nil, 0, errors.New("Result must be in type struct")
}

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func underscore(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}
