package middleware

import (
	"github.com/sde-kiran-sadvilkar/sope"
	"myapp/data"
)

type Middleware struct {
	Core   *sope.Sope
	Models data.Models
}
