package scylldb

import "github.com/ormushq/ormus/manager/entity"

func (a StorageAdapter) InsertSource(source *entity.Source) error {

	panic("implement me")
}
func (a StorageAdapter) UpdateSource(id string, source *entity.Source) error {

	panic("implement me")
}
func (a StorageAdapter) DeleteSource(id string) error {

	panic("implement me")
}
func (a StorageAdapter) GetUserSourceById(ownerID, id string) (*entity.Source, error) {

	panic("implement me")
}
func (a StorageAdapter) IsSourceAlreadyCreatedByName(name string) (bool, error) {

	panic("implement me")
}
