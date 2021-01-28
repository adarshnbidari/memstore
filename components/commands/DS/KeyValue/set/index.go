package set

import (
	"log"

	response "github.com/techslaim/memstore/components/response"
)

//Set is
func Set(key string, value interface{}, memstore map[string]interface{}) response.JSONResponse {

	log.Println("setting the data with the key", key)

	if _, ok := value.(string); ok {

	} else if _, ok := value.(float64); ok {

	} else {

		return response.JSONResponse{

			Status:  "failed",
			Message: "Invalid data type",
		}

	}

	memstore[key] = value

	log.Println("data has been saved")

	return response.JSONResponse{

		Status: "success",
		Body:   memstore,
	}

}
