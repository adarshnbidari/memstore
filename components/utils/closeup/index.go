package closeup

import (
	"log"
	"os"

	"github.com/techslaim/memstore/components/utils/closeup/savetodisk"
	storeformat "github.com/techslaim/memstore/components/utils/storeformat"
)

var backupFolderName = "storage.backup"

//Save for saving the memstore in the disk
func Save(memstore storeformat.Store) {

	savetodisk.SaveToDisk(createNewFolder, backupFolderName, memstore)

}

func createNewFolder() {

	_, err := os.Stat(backupFolderName)

	log.Println("Checking existance of backup folder", backupFolderName)

	if err != nil && os.IsNotExist(err) {

		log.Println("Backup folder not found")
		log.Println("Created new backup folder", backupFolderName)

		os.Mkdir(backupFolderName, 0777)

	}

}
