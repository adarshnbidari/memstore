package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/techslaim/memstore/components/commands"
	"github.com/techslaim/memstore/components/utils/closeup"
	"github.com/techslaim/memstore/components/utils/startup"
	storeformat "github.com/techslaim/memstore/components/utils/storeformat"
)

var memstore storeformat.Store = storeformat.Store{

	KeyValue: make(map[string]interface{}),
	List:     make([]interface{}, 0),
}

type commandFormat struct {
	Command  string      `json:"command"`
	Key      string      `json:"key"`
	Value    interface{} `json:"value,omitempty"`
	DataType string      `json:"datatype,omitempty"`
}

func main() {

	defer cleanup()

	var mutex sync.Mutex

	startupOperation()

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {

		mutex.Lock()
		defer mutex.Unlock()

		var commandData commandFormat

		c.BindJSON(&commandData)

		res := commands.Exec(commandData.Command, commandData.Key, commandData.Value, commandData.DataType, memstore)

		c.JSON(200, res)

	})

	cleanupOperation()

	r.Run(":80")

}

func startupOperation() {

	log.Println("Searching for memstore files in the directory sotrage.backup")
	startup.Load(memstore)

}

func cleanupOperation() {

	stopSignal := make(chan os.Signal, 1)

	signal.Notify(stopSignal, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	<-stopSignal

	signal.Stop(stopSignal)

	cleanup()

	os.Exit(0)

}

func cleanup() {

	log.Println("saving the existing in memory data into the backup directory")
	log.Println("cleaning up the resources")
	closeup.Save(memstore) //cleanup operation

	log.Println("Shutting down...")

}
