package sqlite

import (
	"log"
)

func CheckSchema() {

	/* Step 1. Create "message" store */
	sqlStmt :=  "CREATE TABLE IF NOT EXISTS message (" +
		    "    msgId INTEGER NOT NULL PRIMARY KEY," +
		    "    msgHash CHAR(16) NOT NULL," +
		    "    msgDate INTEGER," +
		    "    msgArea CHAR(64) NOT NULL," +
		    "    msgFrom TEXT NOT NULL," +
		    "    msgTo TEXT NOT NULL," +
		    "    msgSubject TEXT NOT NULL," +
		    "    msgContent TEXT NOT NULL" +
		    ")"
	log.Printf("sql = %q", sqlStmt)

}
