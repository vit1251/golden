package area

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vit1251/golden/pkg/msg"
	"os/user"
	"path/filepath"
	"log"
)

type AreaManager struct {
	Path       string       /* Area path        */
	AreaList   AreaList
}

func NewAreaManager() (*AreaManager) {
	am := new(AreaManager)

	/* Search user home directory */
	usr, err1 := user.Current()
	if err1 != nil {
		panic( err1 )
	}
	userHomeDirectory := usr.HomeDir
	log.Printf("userHomeDirectory = %+v", userHomeDirectory)

	/* Set parameter storage */
	am.Path = filepath.Join(userHomeDirectory, "golden.sqlite3")

	return am
}

func (self *AreaManager) Reset() {
	self.AreaList.Reset()
}

func (self *AreaManager) Register(a *Area) {
	self.AreaList.Areas = append(self.AreaList.Areas, a)
}



func (self *AreaManager) checkSchema(conn *sql.DB) (error) {

	sqlStmt := "CREATE TABLE IF NOT EXISTS area (" +
		    "    areaId INTEGER NOT NULL PRIMARY KEY," +
		    "    areaName CHAR(64) NOT NULL," +
		    "    areaType CHAR(64) NOT NULL," +
		    "    areaPath CHAR(64) NOT NULL," +
		    "    areaSummary CHAR(64) NOT NULL," +
		    "    areaOrder INTEGER NOT NULL," +
		    "    UNIQUE(areaName)" +
		    ")"
	log.Printf("sqlStmt = %s", sqlStmt)
	conn.Exec(sqlStmt)

	return nil

}

func (self *AreaManager) Rescan() {

	/* Open message base */
	messageManager := msg.NewMessageManager()

	/* Preload echo areas */
	areas, err3 := messageManager.GetAreaList2()
	if err3 != nil {
		panic(err3)
	}

	/* Reset areas */
	self.Reset()
	for _, area := range areas {
		log.Printf("area = %q", area)
		a := NewArea()
		a.Name = area.Name
		a.MessageCount = area.Count
		self.Register(a)
	}

}

func (self *AreaManager) GetAreas() ([]*Area, error) {
	self.Rescan()
	return self.AreaList.Areas, nil
}

func (self *AreaManager) GetAreas2() ([]*Area, error) {

	var areas []*Area

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
	if err1 != nil {
		return nil, err1
	}
	defer db.Close()

	/* Check schema */
	self.checkSchema(db)

	/* Restore parameters */
	sqlStmt := "SELECT `areaName` FROM `area`"
	rows, err2 := db.Query(sqlStmt)
	if err2 != nil {
		return nil, err2
	}
	defer rows.Close()
	for rows.Next() {

		var areaName string

		err3 := rows.Scan(&areaName)
		if err3 != nil {
			return nil, err3
		}

		area := NewArea()
		area.Name = areaName

		areas = append(areas, area)

	}

	return areas, nil
}


func (self *AreaManager) GetAreaByName(echoTag string) (*Area, error) {
	self.Rescan()
	return self.AreaList.SearchByName(echoTag)
}

