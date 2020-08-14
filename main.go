package main

import (
	log "gcode/utils/logger"

	"gcode/orm"
	"gcode/server"
)

func main() {
	// Create a new ORM instance to send it to our
	orm, err := orm.Factory()
	defer orm.DB.Close()
	if err != nil {
		log.Panic(err)
	}
	// Send: ORM instance
	server.Run(orm)
}
