package config

type ConfigOption interface {
	apply(config *ProjectConfig) *ProjectConfig
}

type configOptionFunc func(cfg *ProjectConfig) *ProjectConfig

func (fn configOptionFunc) apply(cfg *ProjectConfig) *ProjectConfig {
	return fn(cfg)
}

func WithProjectName(name string) ConfigOption {
	return configOptionFunc(func(cfg *ProjectConfig) *ProjectConfig {
		cfg.ProjectName = name
		return cfg
	})
}
func WithProjectDir(dir string) ConfigOption {
	return configOptionFunc(func(cfg *ProjectConfig) *ProjectConfig {
		cfg.ProjectDir = dir
		return cfg
	})
}

func WithOutModules(modules ...string) ConfigOption {
	return configOptionFunc(func(cfg *ProjectConfig) *ProjectConfig {
		for _, module := range modules {
			if _, ok := cfg.Modules[module]; ok {
				delete(cfg.Modules, module)
			}
		}
		return cfg
	})
}

func WithModulePath(path string) ConfigOption {
	return configOptionFunc(func(cfg *ProjectConfig) *ProjectConfig {
		cfg.ModulePath = path
		return cfg
	})
}
