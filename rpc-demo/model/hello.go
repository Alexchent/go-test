package model

type Person struct {
	Name string
	Age  int
}

type PersonParams struct {
	Name string
}

func (r *Person) MyName(p PersonParams, ret *string) error {
	*ret = "hello " + p.Name + "我的名字是：" + r.Name
	return nil
}
