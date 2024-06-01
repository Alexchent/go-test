package main

import (
	"errors"
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
)

var errorNotExists = errors.New("not exists")

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := getData("mongo")
			if err != nil {
				//log.Println(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
}

var gsf singleflight.Group

func getData(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExists {
		//data, err = getDataFromDB(key)
		// 利用singleflight.Group实现缓存击穿
		v, err, _ := gsf.Do(key, func() (interface{}, error) {
			return getDataFromDB(key)
		})
		if err != nil {
			//log.Println(err)
			return data, err
		}
		data = v.(string)
	} else if err != nil {
		return data, err
	}
	return data, nil
}

func getDataFromCache(key string) (string, error) {
	return "", errorNotExists
}

func getDataFromDB(key string) (string, error) {
	log.Printf("get %s from db", key)
	return "data", nil
}
