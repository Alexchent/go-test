package main

import "fmt"

func main() {
   var a int = 4
   var b int32
   var c float32
   var ptr *int

   /* 运算符实例 */
   fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
   fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
   fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );
   fmt.Printf("第 4 行 - c 变量类型为 = %T\n", ptr );
   fmt.Println(ptr); //空指针为 nil

   /*  & 和 * 运算符实例 */
   ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
   d := &a
   fmt.Printf("a 的值为  %d\n", a);
   fmt.Printf("d 的值为  %d\n", d);
   fmt.Printf("ptr 为 %d\n", ptr);
   fmt.Printf("*ptr 为 %d\n", *ptr);
}