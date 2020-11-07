package tracker

import (
	"github.com/vit1251/golden/pkg/charset"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/mailer/cache"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"log"
	"os"
	"path"
	"time"
)

type Tracker struct {
	registry *registry.Container
}

func NewTracker(r *registry.Container) *Tracker {
	newTracker := new(Tracker)
	newTracker.registry = r
	return newTracker
}

func (self Tracker) Track() {

	trackerStart := time.Now()
	log.Printf("Start tracker session")

	err1 := self.ProcessInbound()
	if err1 != nil {
		log.Printf("err = %+v", err1)
	}
	err2 := self.ProcessOutbound()
	if err2 != nil {
		log.Printf("err = %+v", err2)
	}

	log.Printf("Stop tracker session")
	elapsed := time.Since(trackerStart)

	log.Printf("Tracker session: %+v", elapsed)
}

func (self *Tracker) ProcessInbound() error {

	/* New mailer inbound */
	mi := cache.NewMailerInbound(self.registry)

	/* Scan inbound */
	items, err2 := mi.Scan()
	if err2 != nil {
		return err2
	}
	log.Printf("items = %+v", items)

	for _, item := range items {
		if item.Type == cache.TypeTICmail {
			log.Printf("Tracker: TIC packet: name = %s", item.Name)
			if err := self.processTICmail(item); err != nil {
				log.Printf("Tracker: process TIC with error: err = %+v", err)
			}
		} else {
			// TODO - message about skip ...
		}
	}

	return nil
}

func (self *Tracker) ProcessOutbound() error {
	return nil
}

func (self *Tracker) processTICmail(item cache.FileEntry) error {

	mapperManager := self.restoreMapperManager()
	fileMapper := mapperManager.GetFileMapper()
	statMapper := mapperManager.GetStatMapper()

	/* Parse */
	newTicParser := NewTicParser(self.registry)
	tic, err1 := newTicParser.ParseFile(item.AbsolutePath)
	if err1 != nil {
		return err1
	}
	log.Printf("tic = %+v", tic)

	areaName := tic.GetArea()

	/* Search area */
	fa, err1 := fileMapper.GetAreaByName(areaName)
	if err1 != nil {
		return err1
	}

	/* Prepare area directory */
	boxDirectory := cmn.GetFilesDirectory()
	inboundDirectory := cmn.GetInboundDirectory()

	areaLocation := path.Join(boxDirectory, areaName)
	os.MkdirAll(areaLocation, 0755)

	/* Create area */
	if fa == nil {
		/* Prepare area */
		newFa := mapper.NewFileArea()
		newFa.SetName(areaName)
		newFa.Path = areaLocation
		/* Create area */
		if err := fileMapper.CreateFileArea(newFa); err != nil {
			log.Printf("Fail CreateFileArea on fileMapper: area = %s err = %+v", areaName, err)
			return err
		}
	}

	/* Create new path */
	inboxTicLocation := path.Join(inboundDirectory, tic.File)
	areaFileLocation := path.Join(areaLocation, tic.File)
	log.Printf("inboxTicLocation = %s areaFileLocation = %s", inboxTicLocation, areaFileLocation)

	/* Move */
	if err := os.Rename(inboxTicLocation, areaFileLocation); err != nil {
		log.Printf("Fail on Rename: err = %+v", err)
	}

	/* Register file */
	newFile := mapper.NewFile()
	newFile.SetArea(tic.GetArea())
	newFile.SetDesc(tic.Desc)
	newFile.SetUnixTime(tic.UnixTime)
	newFile.SetFile(tic.File)
	fileMapper.RegisterFile(*newFile)

	/* Register status */
	statMapper.RegisterInFile(tic.File)

	/* Move TIC */
	areaTicLocation := path.Join(areaLocation, item.Name)
	log.Printf("Move %+v -> %+v", item.AbsolutePath, areaTicLocation)
	if err := os.Rename(item.AbsolutePath, areaTicLocation);err != nil {
		log.Printf("Fail on Rename: err = %+v", err)
	}

	return nil
}

func (self Tracker) restoreCharsetManager() *charset.CharsetManager {
	managerPtr := self.registry.Get("CharsetManager")
	if manager, ok := managerPtr.(*charset.CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
}

func (self Tracker) restoreMapperManager() *mapper.MapperManager {
	managerPtr := self.registry.Get("MapperManager")
	if manager, ok := managerPtr.(*mapper.MapperManager); ok {
		return manager
	} else {
		panic("no mapper manager")
	}
}