package main

import (
	"fmt"
)

func reverse(s string){
	var str string
	fmt.Printf("original string is %s\n",s)
	for _,v:=range s{
		str = string(v) + str
	}
	fmt.Printf("reverse of string is %s \n",str)
}



func main(){
	var s string
	fmt.Println("enter the string")
	fmt.Scanf("%s",&s)
	reverse(s)
}