package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

// BindJsonToStruct 绑定JSON数据到结构体，并检查字段名严格匹配
func BindJsonToStruct(c *gin.Context, obj interface{}) error {
	data, err := c.GetRawData()
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields() // 禁止未知字段
	if err = decoder.Decode(&obj); err != nil {
		return err
	}
	if err = checkStrictFields(data, obj); err != nil {
		return err
	}
	return nil
}

// checkStrictFields 检查JSON数据中的字段名是否严格匹配结构体的JSON标签
func checkStrictFields(data []byte, obj interface{}) error {
	structFields, _ := getStructFields(obj)
	for _, field := range structFields {
		if !bytes.Contains(data, []byte(field)) {
			return errors.New("The case of the field names does not match: " + field)
		}
	}
	return nil
}

// getStructFields 获取结构体的JSON标签字段名
func getStructFields(obj interface{}) ([]string, error) {
	var fields []string

	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return nil, errors.New("传入的对象不是结构体指针")
	}

	typ := val.Elem().Type()

	for i := 0; i < val.Elem().NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			fields = append(fields, jsonTag)
		}
	}
	return fields, nil
}
