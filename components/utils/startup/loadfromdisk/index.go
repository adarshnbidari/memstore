package loadfromdisk

import (
	"encoding/json"
	"io/ioutil"
	"log"

	storeformat "github.com/techslaim/memstore/components/utils/storeformat"
)

//LoadFromDisk for loading the data from the disk
func LoadFromDisk(memstore storeformat.Store, backupFolderName string) {

	log.Println("Getting data from the backup folder")

	data, err := ioutil.ReadFile(backupFolderName + "/memstore_key-value.json")

	if err != nil {

		log.Println(err)

		return

	}

	err = json.Unmarshal(data, &memstore.KeyValue)

	if err != nil {

		log.Println(err)
		return

	}

	log.Println("Loaded data from the backup folder")

}
