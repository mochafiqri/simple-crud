package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/mochafiqri/simple-crud/commons/entities"
	"github.com/mochafiqri/simple-crud/commons/interfaces"
	"gorm.io/gorm"
)

type ContentRepo struct {
	db  *gorm.DB
	rds *redis.Client
}

func NewContentRepo(db *gorm.DB, rds *redis.Client) interfaces.ContentRepo {
	return &ContentRepo{
		db:  db,
		rds: rds,
	}
}

func (r ContentRepo) Create(cmd *entities.Content) error {
	cmd.Id = uuid.NewString()
	var err = r.db.Create(&cmd).Error
	return err
}

func (r ContentRepo) Read() (res []entities.Content, err error) {
	err = r.db.Find(&res).Error
	return res, err
}

func (r ContentRepo) Get(id string) (res entities.Content, err error) {
	err = r.db.Where("id = ?", id).Debug().Find(&res).Error
	return res, err
}

func (r ContentRepo) Update(cmd *entities.Content) error {
	var err = r.db.Where("id = ?", cmd.Id).Updates(&cmd).Error
	return err
}

func (r ContentRepo) Delete(id string) error {
	var err = r.db.Where("id = ?", id).Delete(&entities.Content{}).Error
	return err
}
