package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql Mysql
}

type Mysql struct {
	DataSource string
}
