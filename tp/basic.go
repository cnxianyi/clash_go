package tp

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	tag    = map[string]string{}
	allTag = map[AllTag]string{}
)

type AllTag struct {
	yaml string
	meta bool
}

func Tag(c *Clash, s string, value string) Clash {
	fieldName, exists := tag[s] // 读取反射出来的字段
	if !exists {
		fmt.Println("非法字段")
		return *c
	}

	// 反射 值
	v := reflect.ValueOf(c).Elem()

	// 获取 字段名
	field := v.FieldByName(fieldName)
	// 检查字段是否存在 且支持写
	if !field.IsValid() || !field.CanSet() {
		fmt.Println("非法字段类型", s, value)
		return *c
	}

	// 获取 字段类型
	switch field.Kind() {
	case reflect.Int:
		// 去除字符串两侧空格
		i, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			fmt.Println("字符串转换错误")
			return *c
		}
		field.SetInt(int64(i))

	case reflect.String:
		strings.TrimSpace(value)
		field.SetString(value)

	case reflect.Bool:
		if value == "true" {
			field.SetBool(true)
		}
		field.SetBool(false)
	}

	return *c
}

func SetTagMap() {
	// 反射 type
	t := reflect.TypeOf(Clash{})

	tags := map[string]string{}    // 缓存
	allTags := map[AllTag]string{} // 缓存

	// 读取 每一个 type
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// 获取 yaml tag
		ya := field.Tag.Get("yaml")
		me := false // 默认值

		if field.Tag.Get("meta") == "true" {
			me = true
		}

		tagMap := map[string]string{
			"yaml": ya,
		}
		allTagMap := map[string]AllTag{
			"all": {
				yaml: ya,
				meta: me,
			},
		}

		tags[tagMap["yaml"]] = field.Name
		allTags[allTagMap["all"]] = field.Name
	}
	tag = tags
	allTag = allTags
}
