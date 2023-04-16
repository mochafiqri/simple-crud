package usecases

import (
	"errors"
	"github.com/mochafiqri/simple-crud/commons/constants"
	"github.com/mochafiqri/simple-crud/commons/entities"
	"github.com/mochafiqri/simple-crud/commons/interfaces"
	"net/http"
	"time"
)

type ContentUseCase struct {
	repo interfaces.ContentRepo
}

func NewContentUseCase(r interfaces.ContentRepo) interfaces.ContentUseCase {
	return &ContentUseCase{repo: r}
}

func (u ContentUseCase) Create(cmd *entities.Content) (int, string, error) {
	if cmd.Title == "" || cmd.Body == "" {
		return http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errors.New("parameter required")
	}

	var err = u.repo.Create(cmd)
	if err != nil {
		return http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err
	}
	return http.StatusOK, http.StatusText(http.StatusOK), nil
}

func (u ContentUseCase) Read() ([]entities.Content, int, error) {
	contents, err := u.repo.Read()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return contents, http.StatusOK, nil
}

func (u ContentUseCase) Get(id string) (res entities.Content, code int, err error) {
	content, err := u.repo.Get(id)
	if err != nil {
		code = http.StatusInternalServerError
		return
	}

	if content.Id == "" {
		return res, http.StatusNotFound, errors.New(constants.MsgDataNotFound)
	}

	return content, http.StatusOK, nil
}

func (u ContentUseCase) Update(cmd *entities.Content) (code int, err error) {
	if cmd.Title == "" || cmd.Body == "" {
		return http.StatusBadRequest, errors.New("title dan body is required")
	}
	content, err := u.repo.Get(cmd.Id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if content.Id == "" {
		return http.StatusNotFound, errors.New(constants.MsgDataNotFound)
	}

	cmd.CreatedAt = content.CreatedAt
	cmd.UpdateAt = time.Now()

	err = u.repo.Update(cmd)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, nil
}

func (u ContentUseCase) Delete(id string) (code int, err error) {
	content, err := u.repo.Get(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if content.Id == "" {
		return http.StatusNotFound, errors.New(constants.MsgDataNotFound)
	}

	err = u.repo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
