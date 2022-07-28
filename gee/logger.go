package gee

import (
	"log"
	"time"
)

func Logger() HandleFunc {
	return func(ctx *Context) {
		t := time.Now()
		ctx.Next()
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Request.URL, time.Since(t))
	}
}
