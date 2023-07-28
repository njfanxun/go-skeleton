package util

import (
	"bytes"
	"io/fs"
	"os"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

var osFs = afero.NewOsFs()

func MKDir(dir string) error {
	exist, err := afero.DirExists(osFs, dir)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	return osFs.MkdirAll(dir, os.ModePerm)
}

func DirExist(dir string) (bool, error) {
	return afero.DirExists(osFs, dir)
}
func FileExist(filePath string) (bool, error) {
	return afero.Exists(osFs, filePath)
}
func DirIsEmpty(dir string) (bool, error) {
	return afero.IsEmpty(osFs, dir)
}

func TemplateParseFS(fss fs.FS, data any, patterns ...string) (*bytes.Buffer, error) {
	t, err := template.ParseFS(fss, patterns...)
	if err != nil {
		return nil, errors.Wrap(err, "模板文件解析错误")
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return nil, errors.Wrap(err, "模版文件执行错误")
	}
	return &buf, nil
}
