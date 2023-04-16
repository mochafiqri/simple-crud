package interfaces

import (
	"github.com/mochafiqri/simple-crud/commons/entities"
)

type ContentUseCase interface {
	Create(cmd *entities.Content) (int, string, error)
	Read() ([]entities.Content, int, error)
	Get(id string) (entities.Content, int, error)
	Update(cmd *entities.Content) (int, error)
	Delete(id string) (int, error)
}

type ContentRepo interface {
	Create(cmd *entities.Content) error
	Read() ([]entities.Content, error)
	Get(id string) (entities.Content, error)
	Update(cmd *entities.Content) error
	Delete(id string) error
}
