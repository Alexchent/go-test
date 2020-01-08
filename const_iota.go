package main

import "fmt"


const (
    o=1<<iota
    j=3<<iota
    k
    l
)

func main(){

	const (
		a = iota   //0
        b          //1
        c          //2
        d = "ha"   //独立值，iota += 1
        e          //"ha"   iota += 1
        f = 100    //iota +=1
        g          //100  iota +=1
        h = iota   //7,恢复计数
        i          //8
	)

	fmt.Println(a,b,c,d,e,f,g,h,i)
    //0 1 2 ha ha 100 100 7 8

    fmt.Println(o,j,k,l)
    //1 6 12 24
    //iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，
    //关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3
}