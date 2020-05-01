package Singleton

import "sync"


var (
	gua *zhangZhe
	once sync.Once
)
func getInstance() *zhangZhe{
	once.Do(func() {
		gua = &zhangZhe{}
	})
	return gua
}

type zhangZhe struct {
	age int
	mux sync.Mutex
}
func (gua *zhangZhe) getAge() int {
	return gua.age
}
func (gua *zhangZhe) incrementAge(n int) {
	gua.mux.Lock()
	defer gua.mux.Unlock()
	gua.age++
}
