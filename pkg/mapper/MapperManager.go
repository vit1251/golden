package mapper

import "github.com/vit1251/golden/pkg/registry"

type MapperManager struct {
	echoMapper       *EchoMapper
	echoAreaMapper   *EchoAreaMapper
	netmailMapper    *NetmailMapper
	fileMapper       *FileMapper
	fileAreaMapper   *FileAreaMapper
	configMapper     *ConfigMapper
	twitMapper       *TwitMapper
	draftMapper      *DraftMapper
	statMailerMapper *StatMailerMapper
}

func NewMapperManager(r *registry.Container) *MapperManager {
	newMapperManager := new(MapperManager)
	newMapperManager.echoMapper = NewEchoMapper(r)
	newMapperManager.echoAreaMapper = NewEchoAreaMapper(r)
	newMapperManager.statMailerMapper = NewStatMailerMapper(r)
	newMapperManager.netmailMapper = NewNetmailMapper(r)
	newMapperManager.fileMapper = NewFileMapper(r)
	newMapperManager.fileAreaMapper = NewFileAreaMapper(r)
	newMapperManager.configMapper = NewConfigMapper(r)
	newMapperManager.twitMapper = NewTwitMapper(r)
	newMapperManager.draftMapper = NewDraftMapper(r)
	return newMapperManager
}

// / GetEchoMapper provide echo mapper
func (self MapperManager) GetEchoMapper() *EchoMapper {
	return self.echoMapper
}

func (self MapperManager) GetEchoAreaMapper() *EchoAreaMapper {
	return self.echoAreaMapper
}

func (self MapperManager) GetStatMailerMapper() *StatMailerMapper {
	return self.statMailerMapper
}

func (self MapperManager) GetNetmailMapper() *NetmailMapper {
	return self.netmailMapper
}

func (self MapperManager) GetFileAreaMapper() *FileAreaMapper {
	return self.fileAreaMapper
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
