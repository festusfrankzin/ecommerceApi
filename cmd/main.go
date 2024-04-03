package main
import (
	"log"
	"github.com/festusfrankzin/ecommerceApi/cmd/api"
)


func main (){
	server := api.CreateApiServer(":8000", nil)
	if err:= server.Run(); err != nil{
		log.Fatal("Error creating api server")
	}
}