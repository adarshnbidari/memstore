package get

import (
	"log"

	response "github.com/techslaim/memstore/components/response"
)

//Get for getting the value from the key
func Get(key string, memstore map[string]interface{}) response.JSONResponse {

	log.Println("getting the key data from the store")

	exists := memstore[key]

	if exists == nil {

		log.Println("key:", key, "nout found")

		return response.JSONResponse{

			Status:  "failed",
			Message: "key not found",
		}

	}

	log.Println("key found")

	return response.JSONResponse{

		Status: "success",
		Body:   exists,
	}

}
