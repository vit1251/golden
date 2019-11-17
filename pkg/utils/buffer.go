package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"reflect"
)

// ReadStructFromBuffer is reading from buffer to struct
func ReadStructFromBuffer(buffer *bytes.Buffer, data interface{}) error {
	k := reflect.TypeOf(data).Kind()
	if k != reflect.Ptr {
		return errors.New("the second parameter must be a pointer")
	}

	v := reflect.ValueOf(data).Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("invaild type Not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			err := binary.Read(buffer, binary.LittleEndian, v.Field(i).Addr().Interface())
			if err != nil {
				return err
			}
			break
		case reflect.String:
			str, err := buffer.ReadString(0x00)
			if err != nil && err != io.EOF {
				return err
			}
			v.Field(i).SetString(str)
			break
		case reflect.Array:
			err := binary.Read(buffer, binary.LittleEndian, v.Field(i).Addr().Interface())
			if err != nil {
				return err
			}
			break
		default:
			log.Printf("%#v", v.Field(i).Type().Kind())
			return errors.New("invail type Unspport reflect type")
		}
	}

	return nil
}

// WriteStructToBuffer is reading from struct to buffer
func WriteStructToBuffer(buffer *bytes.Buffer, data interface{}) error {
	k := reflect.TypeOf(data).Kind()
	if k != reflect.Ptr {
		return errors.New("the second parameter must be a pointer")
	}

	v := reflect.ValueOf(data).Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("invaild type Not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			err := binary.Write(buffer, binary.LittleEndian, v.Field(i).Addr().Interface())
			if err != nil {
				return err
			}
			break
		case reflect.String:
			_, err := buffer.WriteString(v.Field(i).Interface().(string))
			if err != nil {
				return err
			}
			break
		case reflect.Array:
			err := binary.Write(buffer, binary.LittleEndian, v.Field(i).Addr().Interface())
			if err != nil {
				return err
			}
			break
		default:
			log.Printf("%#v", v.Field(i).Type().Kind())
			return errors.New("invail type Unspport reflect type")
		}
	}

	return nil
}
