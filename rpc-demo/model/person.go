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

func (r Person) Wife(p PersonParams, ret *string) error {
	if p.Name == "谢晓峰" {
		*ret = "慕容秋荻"
	}

	if p.Name == "张翠山" {
		*ret = "殷素素"
	}

	return nil
}
