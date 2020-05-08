package SimpleFactory

import "fmt"


//抽象产品
type FactoryA interface {
	produce()
}

//具体产品
type FactoryB struct {

}
func (b *FactoryB) produce(){
	fmt.Println("生产B工厂需要的产品")
}
type FactoryC struct {

}
func (c *FactoryC) produce(){
	fmt.Println("生产C工厂需要的产品")
}

//工厂类
type Factory struct {
	need byte
}
func (a *Factory)factory() FactoryA{
	switch a.need {
	case 'b':
		return &FactoryB{}
	case 'c':
		return &FactoryC{}
	}
	return nil
}
