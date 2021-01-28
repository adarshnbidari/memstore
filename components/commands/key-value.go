package commands

import (
	delete "github.com/techslaim/memstore/components/commands/DS/KeyValue/delete"
	get "github.com/techslaim/memstore/components/commands/DS/KeyValue/get"
	set "github.com/techslaim/memstore/components/commands/DS/KeyValue/set"
	response "github.com/techslaim/memstore/components/response"
)

func keyValue(dsData dsFormat, responseData *response.JSONResponse) {

	key := dsData.key
	value := dsData.value
	memstore := dsData.memstore.KeyValue

	switch dsData.command {
	case "set":
		*responseData = set.Set(key, value, memstore)
	case "delete":
		*responseData = delete.Delete(key, memstore)
	case "get":
		*responseData = get.Get(key, memstore)
	default:
		*responseData = response.JSONResponse{

			Status:  "failed",
			Message: "invalid command",
		}

	}

}
