package startup

import (
	"log"
	"os"

	loadfromdisk "github.com/techslaim/memstore/components/utils/startup/loadfromdisk"
	storeformat "github.com/techslaim/memstore/components/utils/storeformat"
)

var backupFolderName = "storage.backup"

func init() {

	log.Println("checking existance of the backup folder", backupFolderName)

	_, err := os.Stat(backupFolderName)

	if err != nil && os.IsNotExist(err) {

		log.Println("backup folder not found")

		os.Mkdir(backupFolderName, 0777)

	}

}

//Load for loading the store from the disk
func Load(memstore storeformat.Store) {

	loadfromdisk.LoadFromDisk(memstore, backupFolderName)

}
