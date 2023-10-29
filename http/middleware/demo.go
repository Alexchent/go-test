package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func Chain(handlerFunc http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handlerFunc = middleware(handlerFunc)
	}
	return handlerFunc
}

// Logging 一个简单的中间件，记录请求的路径
func Logging() Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {

			// 中间件处理逻辑
			start := time.Now()
			defer func() {
				log.Printf("%s %s %v", request.Method, request.URL.Path, time.Since(start))
			}()

			// 调用下一个中间件或者最终的handlerFunc
			handlerFunc(writer, request)
		}
	}
}

// Recover 一个简单的中间件，恢复程序运行并返回 500 错误
func Recover() Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {

			// 中间件处理逻辑
			defer func() {
				if err := recover(); err != nil {
					log.Printf("panic: %+v", err)
					http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			// 调用下一个中间件或者最终的handlerFunc
			handlerFunc(writer, request)
		}
	}
}
