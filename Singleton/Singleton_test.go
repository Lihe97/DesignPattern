package Singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T){

	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0 ; i < 1000; i++{
		go func() {
			obj := getInstance()
			obj.incrementAge(1)
			wg.Done()
		}()
	}
	wg.Wait()
	gua := getInstance()
	fmt.Printf("长者的年龄是%d岁！ ",gua.age)
}

