package core

import "reflect"

//CoreModel is the interface that describes any core model
type CoreModel interface {
	TableName() string
	//	GenerateInsertQuery() string
	//	GenerateUpdateQuery(oldValue *CoreModel) string
	//	GenerateDeleteQuery() string
	GenerateRollbackQueries(blockID int64, ecosysteID int64) string
	IsEcosystemModel() bool
	UpdateCache() error
	Type() reflect.Type
}

type Cache interface {
	GetCurrent(tablePrefix int64, id int64) CoreModel
	GetVerified(tablePrefix int64, id int64) CoreModel
	Update(tablePrefix int64, id int64, val CoreModel) error
	Push(tablePrefix int64, id int64, val CoreModel)
	Insert(tablePrefix int64, id int64, val CoreModel) error
	Delete(tablePrefix int64, id int64, val CoreModel) error
	Flush(blockID int64) error
}
