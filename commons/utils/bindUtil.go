package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"

	"github.com/gin-gonic/gin"
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

// BindQueryStrict 绑定 GET 请求的 query 到结构体，并进行严格字段检查
func BindQueryStrict(c *gin.Context, obj interface{}) error {
	values := c.Request.URL.Query()
	if err := mapQueryToStruct(values, obj); err != nil {
		return err
	}
	if err := checkStrictQueryFields(values, obj); err != nil {
		return err
	}
	return nil
}

// mapQueryToStruct 将 query 参数映射到结构体
func mapQueryToStruct(values url.Values, obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return errors.New("obj 必须是结构体指针")
	}
	valElem := val.Elem()
	typ := valElem.Type()

	for i := 0; i < valElem.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		if v, ok := values[jsonTag]; ok && len(v) > 0 {
			switch valElem.Field(i).Kind() {
			case reflect.Int:
				var intVal int
				_, err := fmt.Sscanf(v[0], "%d", &intVal)
				if err != nil {
					return errors.New("参数类型错误: " + jsonTag)
				}
				valElem.Field(i).SetInt(int64(intVal))
			case reflect.String:
				valElem.Field(i).SetString(v[0])
			default:
				return errors.New("暂不支持该字段类型: " + jsonTag)
			}
		}
	}
	return nil
}

// checkStrictQueryFields 检查 query 参数中字段是否严格匹配结构体 JSON 标签
func checkStrictQueryFields(values url.Values, obj interface{}) error {
	structFields, _ := getStructFields(obj)
	for _, field := range structFields {
		if _, ok := values[field]; !ok {
			return errors.New("缺少必填参数或字段名不匹配: " + field)
		}
	}

	for k := range values {
		if !contains(structFields, k) {
			return errors.New("存在多余参数: " + k)
		}
	}
	return nil
}

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
