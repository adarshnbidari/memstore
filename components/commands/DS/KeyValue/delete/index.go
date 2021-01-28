package delete

import (
	"log"

	response "github.com/techslaim/memstore/components/response"
)

//Delete is
func Delete(key string, memstore map[string]interface{}) response.JSONResponse {

	log.Println("Deleting the key", key)

	delete(memstore, key)

	log.Println("Successfully deleted the data with key", key)

	return response.JSONResponse{

		Status:  "success",
		Message: "existing key/value is deleted",
	}

}
