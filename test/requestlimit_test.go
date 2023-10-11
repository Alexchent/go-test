package main

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestHystrixRequestWithErr(t *testing.T) {
	hystrix.ConfigureCommand("mycommand", hystrix.CommandConfig{
		Timeout:                1000, // 超时时间1秒
		MaxConcurrentRequests:  20,   // 最大并发数量20
		SleepWindow:            1000, // 窗口时间1秒，熔断开启1秒后尝试重试
		RequestVolumeThreshold: 5,    // 10秒钟请求数量超过5次，启动熔断器判断
		ErrorPercentThreshold:  50,   // 请求数超过5并且错误率达到百分之50，开启熔断
	})
	wg := new(sync.WaitGroup)

	// 模拟并发10次请求，5次返回err，导致熔断器开启
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = hystrix.Do("mycommand", func() error {
				_, err := http.Get("https://baidu.com")
				if err != nil {
					fmt.Printf("请求失败, err:%s\n", err.Error())
					return err
				}
				if i%2 == 0 {
					return errors.New("测试错误!")
				}
				fmt.Println("success!")
				return nil
			}, func(err error) error {
				fmt.Printf("handle  error:%v\n", err)
				return nil
			})
		}(i)
	}
	wg.Wait()
	fmt.Println("----------------")

	// 继续模拟10次请求，熔断器应该为开启状态
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = hystrix.Do("mycommand", func() error {
				_, err := http.Get("https://baidu.com")
				if err != nil {
					fmt.Printf("请求失败, err:%s\n", err.Error())
					return err
				}
				fmt.Println("success!")
				return nil
			}, func(err error) error {
				fmt.Printf("handle  error:%v\n", err)
				return nil
			})
		}(i)
	}
	wg.Wait()
	fmt.Println("----------------")

	// 睡眠1秒，转换为半开状态，并发请求10次，应该会有一个goroutine真正去请求，返回成功，其它请求直接走fallback逻辑
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = hystrix.Do("mycommand", func() error {
				_, err := http.Get("https://baidu.com")
				if err != nil {
					fmt.Printf("请求失败, err:%s\n", err.Error())
					return err
				}
				fmt.Println("success!")
				return nil
			}, func(err error) error {
				fmt.Printf("handle  error:%v\n", err)
				return nil
			})
		}(i)
	}
	wg.Wait()
	fmt.Println("----------------")

	// 熔断器已经由半开转为关闭状态，请求应该全部成功
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = hystrix.Do("mycommand", func() error {
				//_, err := http.Get("https://www.qixiashu.com/yue/465/465479/940464.html")
				//if err != nil {
				//	fmt.Printf("请求失败, err:%s\n", err.Error())
				//	return err
				//}
				fmt.Println("success!")
				return nil
			}, func(err error) error {
				fmt.Printf("handle  error:%v\n", err)
				return nil
			})
		}(i)
	}
	wg.Wait()
	fmt.Println("----------------")
}

func TestHystrixRequestWithOverLimit(t *testing.T) {
	hystrix.ConfigureCommand("mycommand2", hystrix.CommandConfig{
		Timeout:                1000,  // 超时时间1秒
		MaxConcurrentRequests:  10,    // 最大并发数量20
		SleepWindow:            10000, // 窗口时间1秒，熔断开启10秒后尝试重试
		RequestVolumeThreshold: 5,     // 10秒钟请求数量超过5次，启动熔断器判断
		ErrorPercentThreshold:  50,    // 请求数超过5并且错误率达到百分之50，开启熔断
	})

	wg := sync.WaitGroup{}
	// 模拟并发大于10
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = hystrix.Do("mycommand2", func() error {
				_, err := http.Get("https://baidu.com")
				if err != nil {
					fmt.Printf("请求失败, err:%s\n", err.Error())
					return err
				}
				fmt.Println("success!", i)
				return nil
			}, func(err error) error {
				fmt.Printf("handle  error:%v\n", err)
				return nil
			})
		}(i)
	}
	wg.Wait()
	fmt.Println("------熔断后----------")

	// 继续模拟10次请求，熔断器应该为开启状态
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = hystrix.Do("mycommand2", func() error {
				_, err := http.Get("https://baidu.com")
				if err != nil {
					fmt.Printf("请求失败, err:%s\n", err.Error())
					return err
				}
				fmt.Println("success!", i)
				return nil
			}, func(err error) error {
				fmt.Printf("handle  error:%v\n", err)
				return nil
			})
		}(i)
	}
	wg.Wait()
	fmt.Println("-------继续熔断---------")

	// 继续模拟10次请求，熔断器应该为半开启状态
	time.Sleep(10 * time.Second)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = hystrix.Do("mycommand2", func() error {
				_, err := http.Get("https://baidu.com")
				if err != nil {
					fmt.Printf("请求失败, err:%s\n", err.Error())
					return err
				}
				fmt.Println("success!", i)
				return nil
			}, func(err error) error {
				fmt.Printf("handle  error:%v\n", err)
				return nil
			})
		}(i)
	}
	wg.Wait()
	fmt.Println("--------成功1次--------")

	// 继续模拟10次请求，熔断器应该为关闭状态
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = hystrix.Do("mycommand2", func() error {
				_, err := http.Get("https://baidu.com")
				if err != nil {
					fmt.Printf("请求失败, err:%s\n", err.Error())
					return err
				}
				fmt.Println("success!", i)
				return nil
			}, func(err error) error {
				fmt.Printf("handle  error:%v\n", err)
				return nil
			})
		}(i)
	}
	wg.Wait()
	fmt.Println("----------------")
}
