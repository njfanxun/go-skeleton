package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/njfanxun/go-skeleton/config"
	"github.com/njfanxun/go-skeleton/pkg/skeleton"
	"github.com/njfanxun/go-skeleton/pkg/util"
	"github.com/pterm/pterm"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "go-skeleton",
		Short: "go项目骨架生成器",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},

		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   true,
			DisableDescriptions: true,
			HiddenDefaultCmd:    true,
		},
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       util.GetVersion().Version,
	}

	//--project api-backend -d $GOPATH$/src/go-awesome/test
	cmd.PersistentFlags().StringP("project", "p", "", "项目名称")
	_ = cmd.MarkPersistentFlagRequired("project")
	cmd.PersistentFlags().StringP("dir", "d", "", "项目所在文件夹路径 (default 当前路径)")

	cmd.PersistentFlags().StringSliceP("without", "w", []string{}, fmt.Sprintf("不启用%s框架", strings.Join(lo.Keys[string](config.DefaultModules), ",")))

	cmd.PersistentFlags().StringP("mod", "m", "", ".mod 文件的module路径 (default 项目名称)")
	_ = viper.BindPFlags(cmd.PersistentFlags())
	return cmd
}

func run() {
	cfg, err := flagsHandler()
	if err != nil {
		pterm.Error.Printfln("参数处理错误:%v", err)
		return
	}
	// 1. 检查是否安装go程序
	if err := skeleton.ShellGo(); err != nil {
		pterm.Error.Printfln("无法获取go执行命令")
		return
	}

	// 2. 检查创建项目目录，不是空目录，清除警告
	if err := skeleton.CreateProject(cfg); err != nil {
		pterm.Error.Printfln("创建项目目录错误:%s", err.Error())
		return
	}

	// 3. 执行go mod init 命令
	if err := skeleton.ShellModInit(cfg); err != nil {
		pterm.Error.Printfln("创建go.mod错误:%s", err.Error())
		return
	}

	if err := skeleton.GoMainFile(cfg); err != nil {
		pterm.Error.Printfln("创建 mail.go 错误:%s", err.Error())
		return
	}

	if err := skeleton.GoRootFile(cfg); err != nil {
		pterm.Error.Printfln("创建 root.go 错误:%s", err.Error())
		return
	}

	if err := skeleton.GoSignalFile(cfg); err != nil {
		pterm.Error.Printfln("创建 signal.go 错误:%s", err.Error())
		return
	}

	if err := skeleton.GoGlobalConfigFile(cfg); err != nil {
		pterm.Error.Printfln("创建 global_config.go 错误:%s", err.Error())
		return
	}
	if _, ok := cfg.Modules["gorm"]; ok {
		if err := skeleton.GoDBConfigFile(cfg); err != nil {
			pterm.Error.Printfln("创建 db_config.go 错误:%s", err.Error())
			return
		}
		if err := skeleton.GoDBFile(cfg); err != nil {
			pterm.Error.Printfln("创建 database.go 错误:%s", err.Error())
			return
		}
	}
	if _, ok := cfg.Modules["gin"]; ok {
		if err := skeleton.GoHttpFile(cfg); err != nil {
			pterm.Error.Printfln("创建 http.go 错误:%s", err.Error())
			return
		}
		if err := skeleton.GoRouteFile(cfg); err != nil {
			pterm.Error.Printfln("创建 route.go 错误:%s", err.Error())
			return
		}
	}
	if err := skeleton.GoUtilFile(cfg); err != nil {
		pterm.Error.Printfln("创建 util.go 错误:%s", err.Error())
		return
	}
	if err := skeleton.GoVersionFile(cfg); err != nil {
		pterm.Error.Printfln("创建 version.go 错误:%s", err.Error())
		return
	}
	if err := skeleton.MakeFile(cfg); err != nil {
		pterm.Error.Printfln("创建 Makefile 错误:%s", err.Error())
		return
	}
	if err := skeleton.DockerFile(cfg); err != nil {
		pterm.Error.Printfln("创建 Dockerfile 错误:%s", err.Error())
		return
	}
	if err := skeleton.GitIgnore(cfg); err != nil {
		pterm.Error.Printfln("创建 Dockerfile 错误:%s", err.Error())
		return
	}
	if err := skeleton.ShellGoModTidy(cfg); err != nil {
		pterm.Error.Printfln("执行 go mod tidy 命令错误:%s", err.Error())
		return
	}
}

func flagsHandler() (*config.ProjectConfig, error) {
	dir := viper.GetString("dir")
	if dir == "" {
		dir, _ = os.Getwd()
	}
	name := viper.GetString("project")
	modulePath := viper.GetString("mod")
	if modulePath == "" {
		modulePath = name
	}

	opts := []config.ConfigOption{
		config.WithProjectName(name),
		config.WithProjectDir(dir),
		config.WithModulePath(modulePath),
	}

	out := viper.GetStringSlice("without")
	if len(out) > 0 {
		opts = append(opts, config.WithOutModules(out...))
	}

	cfg := config.NewProjectConfig(opts...)
	if errs := cfg.Validate(); len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return cfg, nil
}
