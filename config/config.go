package config

import (
	"regexp"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

var DefaultModules = map[string]string{
	"zap":  "go.uber.org/zap",
	"gorm": "gorm.io/gorm",
	"gin":  "github.com/gin-gonic/gin",
}

type IConfig interface {
	Validate() []error
}

type ProjectConfig struct {
	ProjectName string
	ProjectDir  string
	Modules     map[string]string
	ModulePath  string
}

func NewProjectConfig(opts ...ConfigOption) *ProjectConfig {
	cfg := &ProjectConfig{
		Modules: DefaultModules,
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}
	return cfg
}

var projectNameRegex = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]*$")

func (p *ProjectConfig) Validate() []error {
	var errs = make([]error, 0)

	if !projectNameRegex.MatchString(p.ProjectName) {
		if p.ProjectName == "" {
			errs = append(errs, errors.New("项目名称不能为空"))
		} else {
			errs = append(errs, errors.Errorf("项目名称[%s]必须以字母开头,只能包含字母[a-zA-z]、数字[0-9]、下划线(_)和连字符(-)", p.ProjectName))
		}

	}
	if p.ProjectDir == "" {
		errs = append(errs, errors.Errorf("项目目录路径不能为空"))
	}
	fs := afero.NewOsFs()
	exist, err := afero.DirExists(fs, p.ProjectDir)
	if err != nil {
		errs = append(errs, err)
	} else {
		if !exist {
			errs = append(errs, errors.Errorf("项目目录[%s]不存在", p.ProjectDir))
		}
	}
	return errs
}
