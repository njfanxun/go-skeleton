package config

import (
    "fmt"
    
    "github.com/pkg/errors"
)

type DBConfig struct {
    Host           string `json:"host" yaml:"host"`
    Port           int    `json:"port" yaml:"port"`
    Username       string `json:"username" yaml:"username"`
    Password       string `json:"password" yaml:"password"`
    Database       string `json:"database" yaml:"database"`
    MaxConnections int    `json:"maxConnections,omitempty" yaml:"maxConnections,omitempty"`
}

func (t *DBConfig) Validate() []error {
    var errs = make([]error, 0)
    if t.Username == "" || t.Password == "" {
        errs = append(errs, errors.Errorf("连接的数据库用户名或密码为空"))
    }
    if t.Database == "" {
        errs = append(errs, errors.Errorf("没有指定需要连接的数据库名称"))
    }
    return errs
}

func NewDefaultDBConfig() *DBConfig {
    return &DBConfig{
        Host:           "127.0.0.1",
        Port:           3306,
        Username:       "",
        Password:       "",
        Database:       "",
        MaxConnections: 10,
    }
}
func (t *DBConfig) DSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", t.Username, t.Password, t.Host, t.Port, t.Database, "charset=utf8mb4&parseTime=true&loc=Asia%2fShanghai")
}
