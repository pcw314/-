package conf

import "fmt"

type Config struct {
	App   *App   `json:"app" yaml:"app"`
	Mysql *Mysql `json:"mysql" yaml:"mysql"`
	Redis *Redis `json:"redis" yaml:"redis"`
	Http  *Http  `json:"http" yaml:"http"`
	Grpc  *Grpc  `json:"grpc" yaml:"grpc"`
}

func NewConfig() *Config {
	return &Config{
		App:   NewDefaultAppConfig(),
		Mysql: NewDefaultMysqlConfig(),
		Redis: NewDefaultRedisConfig(),
		Http:  NewDefaultHttpConfig(),
		Grpc:  NewDefaultGrpcConfig(),
	}
}

func NewDefaultConfig() *Config {
	return NewConfig()
}

type App struct {
	Name   string `json:"name" yaml:"name"`
	Secret string `json:"secret" yaml:"secret"`
}

func NewDefaultAppConfig() *App {
	return &App{
		Name:   "authorization",
		Secret: "123456789",
	}
}

type Mysql struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Database string `json:"database" yaml:"database"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func NewDefaultMysqlConfig() *Mysql {
	return &Mysql{
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "campus_second-hand",
		Username: "root",
		Password: "123456",
	}
}

type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Database int    `json:"database" yaml:"database"`
	Password string `json:"password" yaml:"password"`
}

func NewDefaultRedisConfig() *Redis {
	return &Redis{
		Host:     "127.0.0.1",
		Port:     6379,
		Database: 0,
		Password: "",
	}
}

func (r *Redis) Address() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

type Http struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func NewDefaultHttpConfig() *Http {
	return &Http{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

func (h *Http) Address() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

type Grpc struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func NewDefaultGrpcConfig() *Grpc {
	return &Grpc{
		Host: "127.0.0.1",
		Port: 8090,
	}
}

func (g *Grpc) Address() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}
