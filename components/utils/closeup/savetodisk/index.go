package savetodisk

import (
	"encoding/json"
	"io/ioutil"
	"log"

	storeformat "github.com/techslaim/memstore/components/utils/storeformat"
)

//SaveToDisk for saving the memstore in the disk
func SaveToDisk(createNewFolder func(), backupFolderName string, memstore storeformat.Store) bool {

	encoded, err := json.Marshal(memstore.KeyValue)

	if err != nil {

		log.Println(err)
		return false

	}

	createNewFolder()

	err = ioutil.WriteFile("./"+backupFolderName+"/memstore_key-value.json", encoded, 0777)

	if err != nil {

		log.Println(err)
		return false

	}

	log.Println("store data successfully saved in the backup folder")

	return true

}
