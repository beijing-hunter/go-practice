package main
import "fmt"

type Point struct{
	x int64
	y int64
	z [6]int64
}

type Rect struct{
	leftPoint,rigthPoint Point
}

type A struct{
	number int
}

type B struct{
	number int
}

func main()  {
	
	r1:=Rect{Point{x:12,y:12},Point{x:15,y:16}}
fmt.Printf("r1.leftPoint.x=%p,r1.leftPoint.y=%p,r1.rigthPoint.x=%p,r1.rigthPoint.y=%p\n",
&r1.leftPoint.x,&r1.leftPoint.y,&r1.rigthPoint.x,&r1.rigthPoint.y)

fmt.Printf("r1.leftPoint.z=%p\n",&r1.leftPoint.z)

var a A
var b B

a.number=1
b=B(a)//字段名称、个数和类型完全相同才能转换，否则会提示错误

fmt.Println(a,b)
}