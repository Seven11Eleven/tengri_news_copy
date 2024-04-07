package main

import(
	"tengri_news/handlers"
	
)

func main(){
	r := handlers.SetupRouter()
	r.Run(":8080")
	
}