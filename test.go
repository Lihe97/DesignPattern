package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func test(){
	once.Do(func() {
		fmt.Println("我是创建的")
	})
	fmt.Println("test")
}
func main(){

	for i := 0 ; i < 10 ; i ++{
		test()
	}

}