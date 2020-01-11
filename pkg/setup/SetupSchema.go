package setup

import (
	"log"
)

type SetupSchema struct {
}

func (self *SetupSchema) Create(section string, name string, summary string) {
}

func (self *SetupSchema) Set(section string, name string, value string) {
}

func (self *SetupSchema) Get(section string, name string, defaultValue *string) (*string, error) {
	return nil, nil
}

func (self *SetupSchema) Init() {

	/* Step 2. Create "area" store */
	sqlStmt2 := "CREATE TABLE IF NOT EXISTS area (" +
		    "    areaId INTEGER NOT NULL PRIMARY KEY," +
		    "    areaName CHAR(64) NOT NULL," +
		    "    areaType CHAR(64) NOT NULL," +
		    "    areaPath CHAR(64) NOT NULL," +
		    "    areaSummary CHAR(64) NOT NULL," +
		    "    areaOrder INTEGER NOT NULL," +
		    "    UNIQUE(areaName)" +
		    ")"
	log.Printf("sql = %s", sqlStmt2)

}
