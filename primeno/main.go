package main
import (
	"fmt"
)

func main(){
	var num,flag int
	flag =0
	fmt.Printf("enter number \n")
	fmt.Scanf("%d",&num)
	if num == 1 ||num == 0{
		fmt.Printf("number is not prime\n")
		return
	}
	for i:=2;i<=num/10;i++{
		if num%2 == 0{
			flag =1
		}

	}
	if flag == 0{
		fmt.Printf("number is prime \n")
	}else{
		fmt.Printf("number is not prime \n")
	}
	
}