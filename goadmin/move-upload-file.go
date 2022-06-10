package goadmin

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/code206/goutils/inslice"
	"github.com/code206/goutils/pathfunc"

	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
)

type MoveFuncParam struct {
	FieldName    string   // 表单中上传文件字段名称
	UploadsPath  string   // goadmin默认上传目录
	Exts         []string // 允许上传的扩展名集合
	LevelsStr    string   // 要做 level 的字符串
	LevelsDirSet []int    // 通过hash字符串生成多级目录的设置
	IdStr        string   // 唯一标识字符串，用作文件名
	UrlPrefix    string   // url 前缀
	PublishPath  string   // 发布目录绝对路径
}

// move upload file
func MoveUploadFile(values form.Values, mfp *MoveFuncParam) (string, error) {
	uploadFileName := values.Get(mfp.FieldName) // 获取表单中上传文件字段名称
	if uploadFileName == "" {                   // 如果表单中上传文件字段名称为空，表示没有上传文件，直接返回
		return "", nil
	}

	// 获取上传文件路径，在结束函数时删除此文件
	goadminUploadFile := filepath.Join(mfp.UploadsPath, uploadFileName)
	defer func() {
		os.Remove(goadminUploadFile)
	}()

	// 检查上传文件扩展名是否在允许范围内
	ext := strings.ToLower(path.Ext(uploadFileName))
	if !inslice.InSlice(ext, mfp.Exts) {
		return "", errors.New("file type error")
	}

	// 生成保存文件绝对路径 和 存入数据库的url路径
	levelsDir, err := pathfunc.PathLevels(mfp.LevelsStr, mfp.LevelsDirSet)
	if err != nil {
		return "", err
	}
	subPath := filepath.Join(mfp.UrlPrefix, levelsDir, mfp.IdStr+ext)
	urlPath := filepath.ToSlash(subPath)
	fileStorePath := filepath.Join(mfp.PublishPath, subPath)

	// 建立目录
	if err = os.MkdirAll(filepath.Dir(fileStorePath), 0755); err != nil {
		return "", err
	}
	// 移动文件
	if err = os.Rename(goadminUploadFile, fileStorePath); err != nil {
		return "", err
	}
	return urlPath, nil
}
