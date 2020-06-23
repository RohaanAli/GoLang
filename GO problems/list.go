package main

import(
	"fmt"
	"strconv"
	"container/list"
)


func main()  {

	linkedlist := list.New()
	linkedlist.PushBack("yo")
	linkedlist.PushBack("hello")

	for i:=20; i>=0 ; i--{
		linkedlist.PushFront(strconv.Itoa(i))

	}

	for i:= linkedlist.Front(); i!=nil; i = i.Next(){
		fmt.Println(i.Value)

	}
	
}