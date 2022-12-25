package jasypt

import (
	"errors"
	"reflect"
	"strings"
)

func DecryptObj(j Jasypt, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	} else {
		return errors.New("not a pointer to a struct")
	}
	return DecryptField(j, val)
}

func DecryptStruct(j Jasypt, f reflect.Value) error {
	for i := 0; i < f.NumField(); i++ {
		err := DecryptField(j, f.Field(i))
		if err != nil {
			return err
		}
	}
	return nil
}

func DecryptField(j Jasypt, f reflect.Value) error {
	switch f.Kind() {
	case reflect.Ptr:
		err := DecryptField(j, f.Elem())
		if err != nil {
			return err
		}
	case reflect.Struct:
		err := DecryptStruct(j, f)
		if err != nil {
			return err
		}
	case reflect.String:
		fv := f.String()
		if f.CanSet() && isEncryptStr(fv) {
			dv, err := j.Decrypt(getEncryptStr(fv))
			if err != nil {
				return err
			}
			f.SetString(dv)
		}
	}
	return nil
}

func isEncryptStr(v string) bool {
	return v != "" && strings.HasPrefix(v, "ENC(") && strings.HasSuffix(v, ")")
}

func getEncryptStr(v string) string {
	return strings.TrimSuffix(strings.TrimPrefix(v, "ENC("), ")")
}
