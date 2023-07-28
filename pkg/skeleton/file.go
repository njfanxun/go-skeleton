package skeleton

import (
	"os"
	"path"

	"github.com/njfanxun/go-skeleton/config"
	"github.com/njfanxun/go-skeleton/pkg/tpl"
	"github.com/njfanxun/go-skeleton/pkg/util"
)

func GoMainFile(cfg *config.ProjectConfig) error {
	buf, err := util.TemplateParseFS(tpl.MainGo, cfg, "main.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "main.go"), buf.Bytes(), os.ModePerm)
}

func GoRootFile(cfg *config.ProjectConfig) error {
	if err := util.MKDir(path.Join(cfg.ProjectDir, "cmd")); err != nil {
		return err
	}
	buf, err := util.TemplateParseFS(tpl.RootGo, cfg, "root.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "cmd", "root.go"), buf.Bytes(), os.ModePerm)
}

func GoSignalFile(cfg *config.ProjectConfig) error {
	if err := util.MKDir(path.Join(cfg.ProjectDir, "pkg", "signals")); err != nil {
		return err
	}
	buf, err := util.TemplateParseFS(tpl.SignalGo, cfg, "signal.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "pkg", "signals", "signal.go"), buf.Bytes(), os.ModePerm)
}

func GoGlobalConfigFile(cfg *config.ProjectConfig) error {
	if err := util.MKDir(path.Join(cfg.ProjectDir, "config")); err != nil {
		return err
	}
	buf, err := util.TemplateParseFS(tpl.ConfigGo, cfg, "config.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "config", "global_config.go"), buf.Bytes(), os.ModePerm)
}
func GoDBConfigFile(cfg *config.ProjectConfig) error {
	if err := util.MKDir(path.Join(cfg.ProjectDir, "config")); err != nil {
		return err
	}
	buf, err := util.TemplateParseFS(tpl.DBConfigGo, cfg, "db.config.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "config", "db_config.go"), buf.Bytes(), os.ModePerm)
}

func GoDBFile(cfg *config.ProjectConfig) error {
	if err := util.MKDir(path.Join(cfg.ProjectDir, "pkg", "db")); err != nil {
		return err
	}
	buf, err := util.TemplateParseFS(tpl.DBGo, cfg, "db.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "pkg", "db", "database.go"), buf.Bytes(), os.ModePerm)
}

func GoHttpFile(cfg *config.ProjectConfig) error {
	if err := util.MKDir(path.Join(cfg.ProjectDir, "pkg", "server")); err != nil {
		return err
	}
	buf, err := util.TemplateParseFS(tpl.HttpGo, cfg, "http.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "pkg", "server", "http.go"), buf.Bytes(), os.ModePerm)

}
func GoRouteFile(cfg *config.ProjectConfig) error {
	if err := util.MKDir(path.Join(cfg.ProjectDir, "pkg", "server")); err != nil {
		return err
	}

	buf, err := util.TemplateParseFS(tpl.RouteGo, cfg, "route.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "pkg", "server", "route.go"), buf.Bytes(), os.ModePerm)
}

func GoUtilFile(cfg *config.ProjectConfig) error {
	if err := util.MKDir(path.Join(cfg.ProjectDir, "pkg", "util")); err != nil {
		return err
	}

	buf, err := util.TemplateParseFS(tpl.UtilGo, cfg, "util.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "pkg", "util", "util.go"), buf.Bytes(), os.ModePerm)
}
func GoVersionFile(cfg *config.ProjectConfig) error {
	if err := util.MKDir(path.Join(cfg.ProjectDir, "pkg", "util")); err != nil {
		return err
	}

	buf, err := util.TemplateParseFS(tpl.VersionGo, cfg, "version.go.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "pkg", "util", "version.go"), buf.Bytes(), os.ModePerm)
}

func MakeFile(cfg *config.ProjectConfig) error {

	buf, err := util.TemplateParseFS(tpl.Makefile, cfg, "Makefile.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "Makefile"), buf.Bytes(), os.FileMode(0644))
}
func DockerFile(cfg *config.ProjectConfig) error {

	buf, err := util.TemplateParseFS(tpl.Dockerfile, cfg, "dockerfile.tmpl")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(cfg.ProjectDir, "Dockerfile"), buf.Bytes(), os.FileMode(0644))
}

func GitIgnore(cfg *config.ProjectConfig) error {
	return os.WriteFile(path.Join(cfg.ProjectDir, ".gitignore"), tpl.GitIgnore, os.FileMode(0644))
}
