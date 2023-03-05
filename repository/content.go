package repository

import "github.com/mochafiqri/simple-crud/entities"

type contentRepo struct {
	Contents *map[string]entities.Content
}

func NewContentRepo(content *map[string]entities.Content) entities.ContentRepo {
	return &contentRepo{}
}

func (r *contentRepo) GetAll() ([]entities.Content, error) {
	var data = make([]entities.Content, 0)
	for _, v := range *r.Contents {
		data = append(data, v)
	}
	return data, nil
}

func (r *contentRepo) Get(id string) (entities.Content, error) {
	var (
		tmp  = *r.Contents
		data entities.Content
	)

	if _, ok := tmp[id]; ok {
		data = tmp[id]
	}
	return data, nil
}

func (r *contentRepo) Save(data entities.Content) error {
	//TODO implement me
	panic("implement me")
}

func (r *contentRepo) Update(data entities.Content) error {
	//TODO implement me
	panic("implement me")
}

func (r *contentRepo) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
