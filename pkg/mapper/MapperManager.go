package mapper

import "github.com/vit1251/golden/pkg/registry"

type MapperManager struct {
	echoMapper     *EchoMapper
	echoAreaMapper *EchoAreaMapper
	statMapper     *StatMapper
	netmailMapper  *NetmailMapper
	fileMapper     *FileMapper
	configMapper   *ConfigMapper
	twitMapper     *TwitMapper
	draftMapper    *DraftMapper
}

func NewMapperManager(r *registry.Container) *MapperManager {
	newMapperManager := new(MapperManager)
	newMapperManager.echoMapper = NewEchoMapper(r)
	newMapperManager.echoAreaMapper = NewEchoAreaMapper(r)
	newMapperManager.statMapper = NewStatMapper(r)
	newMapperManager.netmailMapper = NewNetmailMapper(r)
	newMapperManager.fileMapper = NewFileMapper(r)
	newMapperManager.configMapper = NewConfigMapper(r)
	newMapperManager.twitMapper = NewTwitMapper(r)
	newMapperManager.draftMapper = NewDraftMapper(r)
	return newMapperManager
}

/// GetEchoMapper provide echo mapper
func (self MapperManager) GetEchoMapper() *EchoMapper {
	return self.echoMapper
}

func (self MapperManager) GetEchoAreaMapper() *EchoAreaMapper {
	return self.echoAreaMapper
}

func (self MapperManager) GetStatMapper() *StatMapper {
	return self.statMapper
}

func (self MapperManager) GetNetmailMapper() *NetmailMapper {
	return self.netmailMapper
}

func (self MapperManager) GetFileMapper() *FileMapper {
	return self.fileMapper
}

func (self MapperManager) GetConfigMapper() *ConfigMapper {
	return self.configMapper
}

func (self MapperManager) GetTwitMapper() *TwitMapper {
	return self.twitMapper
}

func (self MapperManager) GetDraftMapper() *DraftMapper {
	return self.draftMapper
}
