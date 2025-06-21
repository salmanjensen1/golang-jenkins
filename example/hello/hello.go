package main

import ("fmt"
"example/greetings"
"log"
)

func main(){
	log.SetPrefix("greetings")
	log.SetFlags(0)
	name:="Vergil"
	message,err:=greetings.Hello(name)
	if(err!=nil){
		log.Fatal(err)
	}
	fmt.Println(message)
}


