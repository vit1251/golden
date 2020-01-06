package sqlite

import (
	"log"
)

func CheckSchema() {

	/* Step 1. Create "message" store */
	sqlStmt :=  "CREATE TABLE IF NOT EXISTS message (" +
		    "    msgId INTEGER NOT NULL PRIMARY KEY," +
		    "    msgHash CHAR(16)," +
		    "    msgDate INTEGER," +
		    "    msgArea CHAR(64)," +
		    "    msgFrom TEXT," +
		    "    msgTo TEXT," +
		    "    msgSubject TEXT," +
		    "    msgContent TEXT" +
		    ")"
	log.Printf("sql = %q", sqlStmt)

}
