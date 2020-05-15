package Strategy

import "fmt"

type Strategy interface {
	backSchool()
}

type StrategyA struct {

}

func (s *StrategyA) backSchool(){
	fmt.Println("坐飞机回学校，预订机票")
}


type StrategyB struct {

}

func (s *StrategyB) backSchool(){
	fmt.Println("坐高铁回学校，预订车票")
}



type Context struct {
	strategy Strategy
}



func (c *Context) SetStrategy(strategy Strategy){
	c.strategy = strategy
}

func (c *Context) Execute(){
	c.strategy.backSchool()
}