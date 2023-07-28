package tpl

import (
	"embed"
	_ "embed"
)

//go:embed main.go.tmpl
var MainGo embed.FS

//go:embed root.go.tmpl
var RootGo embed.FS

//go:embed signal.go.tmpl
var SignalGo embed.FS

//go:embed config.go.tmpl
var ConfigGo embed.FS

//go:embed db.config.go.tmpl
var DBConfigGo embed.FS

//go:embed db.go.tmpl
var DBGo embed.FS

//go:embed http.go.tmpl
var HttpGo embed.FS

//go:embed route.go.tmpl
var RouteGo embed.FS

//go:embed util.go.tmpl
var UtilGo embed.FS

//go:embed version.go.tmpl
var VersionGo embed.FS

//go:embed Makefile.tmpl
var Makefile embed.FS

//go:embed dockerfile.tmpl
var Dockerfile embed.FS

//go:embed dockerfile.tmpl
var GitIgnore []byte
