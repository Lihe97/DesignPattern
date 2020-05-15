package Strategy

import "testing"

func TestContext_Execute(t *testing.T) {
	strategyA := &StrategyA{}
	c := &Context{}
	c.SetStrategy(strategyA)
	c.Execute()

	strategyB := &StrategyB{}
	c.SetStrategy(strategyB)
	c.Execute()
}
