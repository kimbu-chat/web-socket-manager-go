package config

import "os"

type env uint8

const (
	Development env = iota
	Production
)

func (e env) Dev() bool {
	return e == Development
}

func (e env) Prod() bool {
	return e == Production
}

func Env() env {
	if os.Getenv("ENV") == "production" {
		return Production
	}

	return Development
}
