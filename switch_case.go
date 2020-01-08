package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var grade string = "B"
   var marks int = 70

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }

   switch grade {
      case "A" :
         fmt.Printf("优秀!\n" )    
      case "B","C" :
         fmt.Printf("良好\n" )      
      case "D" :
         fmt.Printf("及格\n" )      
      case "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" );
   }
   fmt.Printf("你的等级是 %s\n", grade );      
}