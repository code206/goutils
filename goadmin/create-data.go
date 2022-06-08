package goadmin

import (
	"errors"

	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"

	"github.com/code206/goutils/inslice"
)

// 创建 insert 或者 update 需要的 data
func CreateData(values form.Values, fields ...string) (dialect.H, error) {
	if len(fields) == 0 {
		return nil, errors.New("fields is empty")
	}

	data := dialect.H{}
	for k, v := range values.ToMap() {
		if inslice.InSlice(k, fields) {
			data[k] = v
		}
	}

	if len(data) == 0 {
		return nil, errors.New("data is empty")
	} else {
		return data, nil
	}
}
