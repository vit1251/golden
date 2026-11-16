package mapper

import (
	"database/sql"
	"log"
	
	"github.com/huandu/go-sqlbuilder"

	"github.com/vit1251/golden/internal/utils"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
)

type EchoAreaMapper struct {
	Mapper
}

func NewEchoAreaMapper(r *registry.Container) *EchoAreaMapper {
	newEchoAreaMapper := new(EchoAreaMapper)
	newEchoAreaMapper.SetRegistry(r)
	return newEchoAreaMapper
}

func (self *EchoAreaMapper) Register(a *Area) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	var areaName string = a.GetName()
	var areaType string = ""
	var areaPath string = ""
	var areaSummary string = a.GetSummary()
	var areaOrder int64 = a.GetOrder()
	var areaIndex string = a.GetAreaIndex()

	/* Make unique area index */
	if areaIndex == "" {
		areaIndex = utils.IndexHelper_makeUUID()
		a.SetAreaIndex(areaIndex)
	}

	//                                   1          2           3            4             5            6
	query1 := "INSERT INTO `area` ( `areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder`, `areaIndex` ) VALUES ( ?, ?, ?, ?, ?, ? )"
	var params []interface{}
	params = append(params, areaName)    // 1
	params = append(params, areaType)    // 2
	params = append(params, areaPath)    // 3
	params = append(params, areaSummary) // 4
	params = append(params, areaOrder)   // 5
	params = append(params, areaIndex)   // 6

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return err
	})

	return err1
}

func (self *EchoAreaMapper) GetAreas() ([]*Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)
	mapperManager := RestoreMapperManager(self.registry)
	echoMapper := mapperManager.GetEchoMapper()

	var result []*Area

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("areaIndex", "areaName", "areaSummary", "areaCharset", "areaOrder")
	sb.From("area")
	sb.OrderBy("areaOrder ASC", "areaName ASC")
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {

		var areaIndex string
		var areaName string
		var areaSummary string
		var areaCharset string
		var areaOrder int64

		err3 := rows.Scan(&areaIndex, &areaName, &areaSummary, &areaCharset, &areaOrder)
		if err3 != nil {
			return err3
		}

		area := NewArea()
		area.SetAreaIndex(areaIndex)
		area.SetName(areaName)
		area.SetSummary(areaSummary)
		area.SetCharset(areaCharset)
		area.SetOrder(areaOrder)

		result = append(result, area)

		return nil
	})
	if err != nil {
		return nil, err
	}

	if err1 := echoMapper.UpdateAreaMessageCounters(result); err1 != nil {
		return nil, err1
	}

	return result, nil
}

func (self *EchoAreaMapper) GetAreaByName(areaName string) (*Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result *Area

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("areaIndex", "areaName", "areaSummary", "areaCharset", "areaOrder")
	sb.From("area")
	sb.Where(sb.Equal("areaName", areaName))
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {

		var areaIndex string
		var origAreaName string
		var areaSummary string
		var areaCharset string
		var areaOrder int64

		err3 := rows.Scan(&areaIndex, &origAreaName, &areaSummary, &areaCharset, &areaOrder)
		if err3 != nil {
			return err3
		}

		area := NewArea()
		area.SetAreaIndex(areaIndex)
		area.SetName(origAreaName)
		area.SetSummary(areaSummary)
		area.SetCharset(areaCharset)
		area.SetOrder(areaOrder)

		result = area

		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (self *EchoAreaMapper) GetAreaByAreaIndex(areaIndex string) (*Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result *Area

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("areaIndex", "areaName", "areaSummary", "areaCharset", "areaOrder")
	sb.From("area")
	sb.Where(sb.Equal("areaIndex", areaIndex))
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {

		var newAreaIndex string
		var areaName string
		var areaSummary string
		var areaCharset string
		var areaOrder int64

		err3 := rows.Scan(&newAreaIndex, &areaName, &areaSummary, &areaCharset, &areaOrder)
		if err3 != nil {
			return err3
		}

		area := NewArea()
		area.SetAreaIndex(newAreaIndex)
		area.SetName(areaName)
		area.SetSummary(areaSummary)
		area.SetCharset(areaCharset)
		area.SetOrder(areaOrder)

		result = area

		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (self *EchoAreaMapper) Update(area *Area) error {
	if area.areaIndex == "" {
		return self.updateByAreaName(area)
	} else {
		// TODO - implement update by `areaIndex` here ...
		return self.updateByAreaName(area)
	}
}

func (self *EchoAreaMapper) updateByAreaName(area *Area) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("area")
	ub.Set(
	    ub.Assign("areaIndex", area.GetAreaIndex()),
	    ub.Assign("areaSummary", area.GetSummary()),
	    ub.Assign("areaCharset", area.GetCharset()),
	    ub.Assign("areaOrder", area.GetOrder()),
	)
	ub.Where(ub.Equal("areaName", area.GetName()))
	query1, args := ub.Build()

	return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
		return err
	})
}

func (self *EchoAreaMapper) RemoveAreaByName(echoName string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	db := sqlbuilder.NewDeleteBuilder()
	db.DeleteFrom("area")
	db.Where(db.Equal("areaName", echoName))
	query1, args := db.Build()

	return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
		return err
	})
}
