package commands

import (
	"log"
	"strings"

	response "github.com/techslaim/memstore/components/response"
	storeformat "github.com/techslaim/memstore/components/utils/storeformat"
)

type dsFormat struct {
	command  string
	key      string
	value    interface{}
	memstore storeformat.Store
}

//Exec is
func Exec(command string, key string, value interface{}, dataType string, memstore storeformat.Store) response.JSONResponse {

	dsData := dsFormat{command, key, value, memstore}

	log.Println("executing the command", command)

	command = strings.ToLower(command)

	var responseData response.JSONResponse

	switch dataType {

	case "key-value":
		keyValue(dsData, &responseData)
	default:
		responseData = response.JSONResponse{

			Status:  "failed",
			Message: "invalid Data Type",
		}

	}

	return responseData

}
