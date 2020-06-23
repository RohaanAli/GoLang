package main

import (
	"fmt"
)



func say(i interface{}){
	switch v:= i.(type) {
	case int:
		fmt.Println("its an integer" ,v)
	case float64:
		fmt.Println("its a float",v)
		
	}
}


func main(){
	say(14)
	say(18.0)
}