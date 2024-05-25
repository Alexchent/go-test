package soild

import "fmt"

// 模板模式
// 定义一个接口
// 定义一个抽象类, 定义具体的步骤, 交给具体的子类实现, 封装具体步骤, 定义具体的子类, 调用具体的子类, 执行具体的步骤

type Cooker interface {
	fire()
	cooke()
	outfire()
}

// 类似于一个抽象类
type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜，交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

// XiHongShi 具体的子类
type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("做西红柿")
}

type ChaoJiDan struct {
	CookMenu
}

func (*ChaoJiDan) cooke() {
	fmt.Println("做炒鸡蛋")
}

type XiaoChaorou struct {
	CookMenu
}

func (*XiaoChaorou) cooke() {
	fmt.Println("做小炒肉")
}
