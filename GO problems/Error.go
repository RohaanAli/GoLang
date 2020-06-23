package main

import (
	"fmt"
	//"strings"
)

func div(x,y float64) (float64, error){
	if y == 0{
		err := ("The divisor cant be zero")
		return 0 , fmt.Errorf(err)
	}
	return x/y , nil
}
func main(){

	x := 5.0
	y := 7.0
	
	result, err := div(x,y)

	if err!=nil{
		fmt.Println(result , err)
	}
	fmt.Println(result, nil)

}

