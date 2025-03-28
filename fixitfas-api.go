package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	"errors"
)

type Ticket struct {
	ID    		int    	`json:"id"`
	Reporter 	string	`json:"reporter"`
	Description string 	`json:"description"`
	Status 		string 	`json:"status"`
	CreatedAt 	string 	`json:"created_at"`
}

func main(){
	router := gin.Default()
}