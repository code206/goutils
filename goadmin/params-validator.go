package goadmin

import (
	"errors"
	"strconv"
	"strings"

	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"

	"github.com/code206/goutils/str"
)

type ParamsValidateRule struct {
	Field string // 要检查的字段名称
	Type  string //要检查的字段类型，number 和 text
	Start int    // 开始位置
	End   int    // 结束位置
}

// 参数验证，不满足规则要求时，就跳过
func ParamsValidator(values form.Values, pvr []ParamsValidateRule) error {
	for _, validate := range pvr {
		if values.Has(validate.Field) {
			// 检查数字参数是否在规则限定范围内
			if validate.Type == "number" {
				// 数字参数在规则限定范围内，continue
				if i, err := strconv.Atoi(values.Get(validate.Field)); err == nil {
					if validate.Start <= i && i <= validate.End {
						continue
					}
				}
				// 数字参数不在规则限定范围内，如果是id字段，就报错
				if validate.Field == "id" {
					return errors.New("id error")
				}
				// 数字参数不在规则限定范围内，使用规则中开始位置的数字做默认值
				values.Add(validate.Field, strconv.Itoa(validate.Start))
			}

			// 检查字符串的长度是否在规则限定范围内
			if validate.Type == "text" {
				values.Add(validate.Field, strings.TrimSpace(values.Get(validate.Field)))
				// 字符串的长度在规则限定范围，continue
				if len(values.Get(validate.Field)) <= validate.End {
					continue
				}
				// 字符串的长度超过规则限定范围，截取规则内的部分
				values.Add(validate.Field, str.SubString(values.Get(validate.Field), validate.Start, validate.End))
			}

		} // 完成一条规则的处理
	} // 完成所有规则的处理

	return nil
}
