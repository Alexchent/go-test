package soild

// 策略模式
type IStrategy interface {
	do(int2 int, int3 int) int
}

type Add struct{}

func (a *Add) do(int2 int, int3 int) int {
	return int2 + int3
}

type Sub struct{}

func (s *Sub) do(int2 int, int3 int) int {
	return int2 - int3
}

type Operation struct {
	strategy IStrategy
}

func (o *Operation) SetStrategy(strategy IStrategy) {
	o.strategy = strategy
}

func (o *Operation) ExecuteStrategy(int2 int, int3 int) int {
	return o.strategy.do(int2, int3)
}
