package mapstruct

import (
	"database/sql"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func MapStruct(from, to interface{}) {
	fromValue := reflect.ValueOf(from)
	toValue := reflect.ValueOf(to).Elem()

	for i := 0; i < fromValue.NumField(); i++ {
		fromField := fromValue.Field(i)
		toField := toValue.FieldByName(fromValue.Type().Field(i).Name)

		if toField.IsValid() && toField.CanSet() {
			toFieldType := toField.Type()

			if fromField.Kind() == reflect.Ptr && fromField.Type() == toFieldType {
				if !fromField.IsNil() {
					toField.Set(fromField.Elem())
				}
			} else if fromField.Type() == toFieldType {
				toField.Set(fromField)
			} else if fromField.Kind() == reflect.Struct && fromField.Type() == toFieldType {
				MapStruct(fromField.Interface(), toField.Addr().Interface())
			} else {
				mapField(fromField, toField)
			}
		}
	}

}

func mapField(fromField, toField reflect.Value) {
	if fromField.Kind() == reflect.Ptr {
		fromField = fromField.Elem()
	}

	if toField.Kind() == reflect.Ptr {
		toField = toField.Elem()
	}

	if !fromField.IsValid() && !toField.IsValid() {
		return
	}

	toFieldType := toField.Type()
	fromFieldType := fromField.Type()

	switch toFieldType.Kind() {
	case reflect.String:
		switch fromFieldType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			toField.SetString(strconv.FormatInt(fromField.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			toField.SetString(strconv.FormatUint(fromField.Uint(), 10))
		case reflect.Float32:
			toField.SetString(strconv.FormatFloat(fromField.Float(), 'f', -1, 32))
		case reflect.Float64:
			toField.SetString(strconv.FormatFloat(fromField.Float(), 'f', -1, 64))
		case reflect.Struct:
			switch fromFieldType {
			case reflect.TypeOf(sql.NullString{}):
				nullString := fromField.Interface().(sql.NullString)
				if nullString.Valid {
					toField.SetString(nullString.String)
				} else {
					toField.SetString(" ")
				}
			case reflect.TypeOf(sql.NullInt64{}):
				nullInt64 := fromField.Interface().(sql.NullInt64)
				if nullInt64.Valid {
					toField.SetString(strconv.FormatInt(nullInt64.Int64, 10))
				}
			case reflect.TypeOf(sql.NullInt32{}):
				nullInt32 := fromField.Interface().(sql.NullInt32)
				if nullInt32.Valid {
					toField.SetString(strconv.FormatInt(int64(nullInt32.Int32), 10))
				}
			case reflect.TypeOf(sql.NullInt16{}):
				nullInt16 := fromField.Interface().(sql.NullInt16)
				if nullInt16.Valid {
					toField.SetString(strconv.FormatInt(int64(nullInt16.Int16), 10))
				}
			case reflect.TypeOf(sql.NullFloat64{}):
				nullFloat64 := fromField.Interface().(sql.NullFloat64)
				if nullFloat64.Valid {
					toField.SetString(strconv.FormatFloat(nullFloat64.Float64, 'f', -1, 64))
				}
			}
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch fromFieldType.Kind() {
		case reflect.String:
			fromVal := fromField.String()
			regex := regexp.MustCompile(`\.0+$`)
			fromVal = regex.ReplaceAllString(fromVal, "")
			fromVal = strings.Replace(fromVal, ",", "", -1)
			intVal, err := strconv.ParseInt(fromVal, 10, toFieldType.Bits())
			if err == nil {
				toField.SetInt(intVal)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			toField.SetInt(int64(fromField.Uint()))
		case reflect.Float32, reflect.Float64:
			toField.SetInt(int64(fromField.Float()))
		case reflect.Interface:
			var uintVal int64
			switch interfaceType := fromField.Interface().(type) {
			case int, int8, int16, int32, int64:
				uintVal = int64(interfaceType.(int64))
			case uint, uint8, uint16, uint32, uint64:
				uintVal = int64(interfaceType.(uint64))
			case float32:
				uintVal = int64(interfaceType)
			case float64:
				uintVal = int64(interfaceType)
			case string:
				floatval, err := strconv.ParseFloat(interfaceType, toFieldType.Bits())
				if err == nil {
					uintVal = int64(floatval)
				}
			}
			toField.SetInt(uintVal)
		case reflect.Struct:
			switch fromFieldType {
			case reflect.TypeOf(sql.NullInt64{}):
				nullInt64 := fromField.Interface().(sql.NullInt64)
				if nullInt64.Valid {
					toField.SetInt(nullInt64.Int64)
				} else {
					toField.SetInt(0)
				}
			case reflect.TypeOf(sql.NullInt32{}):
				nullInt32 := fromField.Interface().(sql.NullInt32)
				if nullInt32.Valid {
					toField.SetInt(int64(nullInt32.Int32))
				} else {
					toField.SetInt(0)
				}
			case reflect.TypeOf(sql.NullInt16{}):
				nullInt16 := fromField.Interface().(sql.NullInt16)
				if nullInt16.Valid {
					toField.SetInt(int64(nullInt16.Int16))
				} else {
					toField.SetInt(0)
				}
			case reflect.TypeOf(sql.NullFloat64{}):
				nullFloat64 := fromField.Interface().(sql.NullFloat64)
				if nullFloat64.Valid {
					toField.SetInt(int64(nullFloat64.Float64))
				} else {
					toField.SetInt(0)
				}
			}

		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch fromFieldType.Kind() {
		case reflect.String:
			fromVal := fromField.String()
			fromVal = strings.Replace(fromVal, ",", "", -1)
			uintVal, err := strconv.ParseFloat(fromVal, toFieldType.Bits())
			if err == nil {
				toField.SetUint(uint64(uintVal))
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal := fromField.Int()
			if intVal >= 0 {
				toField.SetUint(uint64(intVal))
			}
		case reflect.Float32, reflect.Float64:
			toField.SetUint(uint64(fromField.Float()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			toField.SetUint(fromField.Uint())
		case reflect.Interface:
			var uintVal uint64
			switch interfaceType := fromField.Interface().(type) {
			case int, int8, int16, int32, int64:
				uintVal = uint64(interfaceType.(int64))
			case uint, uint8, uint16, uint32, uint64:
				uintVal = uint64(interfaceType.(uint64))
			case float32:
				uintVal = uint64(interfaceType)
			case float64:
				uintVal = uint64(interfaceType)
			case string:
				floatval, err := strconv.ParseFloat(interfaceType, toFieldType.Bits())
				if err == nil {
					uintVal = uint64(floatval)
				}
			}
			toField.SetUint(uintVal)
		case reflect.Struct:
			switch fromFieldType {
			case reflect.TypeOf(sql.NullInt64{}):
				nullInt64 := fromField.Interface().(sql.NullInt64)
				if nullInt64.Valid {
					toField.SetUint(uint64(nullInt64.Int64))
				} else {
					toField.SetUint(0)
				}
			case reflect.TypeOf(sql.NullInt32{}):
				nullInt32 := fromField.Interface().(sql.NullInt32)
				if nullInt32.Valid {
					toField.SetUint(uint64(nullInt32.Int32))
				} else {
					toField.SetUint(0)
				}
			case reflect.TypeOf(sql.NullInt16{}):
				nullInt16 := fromField.Interface().(sql.NullInt16)
				if nullInt16.Valid {
					toField.SetUint(uint64(nullInt16.Int16))
				} else {
					toField.SetUint(0)
				}
			case reflect.TypeOf(sql.NullFloat64{}):
				nullFloat64 := fromField.Interface().(sql.NullFloat64)
				if nullFloat64.Valid {
					toField.SetUint(uint64(nullFloat64.Float64))
				} else {
					toField.SetUint(0)
				}
			}
		}
	}
}
