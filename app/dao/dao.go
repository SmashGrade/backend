package dao

import "time"

// General DAO methods
type DAO interface {
	Create(model *any) error
	Update(model *any) error
	Find(model *any) (models []*any, err error)
	GetAll() (models []*any, err error)
}

// DAO methods for models with only an id
type IdDAO interface {
	Get(id uint) (entity *any, err error)
	Delete(id uint) error
}

// DAO methods for models with id and version
type VersionedDAO interface {
	Get(id, version uint) (entity *any, err error)
	Delete(id, version uint) error
}

// DAO methods for models with id and start date
type TimedDAO interface {
	Get(id uint, startDate time.Time) (entity *any, err error)
	Delete(id uint, startDate time.Time) error
}
