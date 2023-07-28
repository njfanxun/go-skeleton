package skeleton

import (
	"os/exec"
	"path"

	"github.com/njfanxun/go-skeleton/config"
	"github.com/njfanxun/go-skeleton/pkg/util"
)

// ShellGo 检查是否可执行go命令
func ShellGo() error {
	cmd := exec.Command("go", "version")
	return cmd.Run()
}

// CreateProject 创建项目目录
func CreateProject(cfg *config.ProjectConfig) error {
	cfg.ProjectDir = path.Join(cfg.ProjectDir, cfg.ProjectName)
	exist, err := util.DirExist(cfg.ProjectDir)
	if err != nil {
		return err
	}
	if exist {
		//isEmpty, err := util.DirIsEmpty(cfg.ProjectDir)
		//if err != nil {
		//	return err
		//}
		//if !isEmpty {
		//	return errors.Errorf("项目目录[%q]包含其他文件，请选择空目录创建项目", cfg.ProjectDir)
		//}
		return nil
	}
	return util.MKDir(cfg.ProjectDir)

}

// ShellModInit 初始化go.mod文件
func ShellModInit(cfg *config.ProjectConfig) error {
	exist, err := util.FileExist(path.Join(cfg.ProjectDir, "go.mod"))
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	cmd := exec.Command("go", "mod", "init", cfg.ModulePath)
	cmd.Dir = cfg.ProjectDir
	return cmd.Run()
}

func ShellGoModTidy(cfg *config.ProjectConfig) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = cfg.ProjectDir
	return cmd.Run()
}
