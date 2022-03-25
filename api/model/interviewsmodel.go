package model

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/zeromicro/go-zero/core/stores/mongo"
)

type InterviewsModel interface {
	Insert(ctx context.Context, data *Interviews) error
	FindOne(ctx context.Context, id string) (*Interviews, error)
	Update(ctx context.Context, data *Interviews) error
	Delete(ctx context.Context, id string) error
}

type defaultInterviewsModel struct {
	*mongo.Model
}

func NewInterviewsModel(url, collection string) InterviewsModel {
	return &defaultInterviewsModel{
		Model: mongo.MustNewModel(url, collection),
	}
}

func (m *defaultInterviewsModel) Insert(ctx context.Context, data *Interviews) error {
	if !data.ID.Valid() {
		data.ID = bson.NewObjectId()
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	return m.GetCollection(session).Insert(data)
}

func (m *defaultInterviewsModel) FindOne(ctx context.Context, id string) (*Interviews, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data Interviews

	err = m.GetCollection(session).FindId(bson.ObjectIdHex(id)).One(&data)
	switch err {
	case nil:
		return &data, nil
	case mongo.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInterviewsModel) Update(ctx context.Context, data *Interviews) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).UpdateId(data.ID, data)
}

func (m *defaultInterviewsModel) Delete(ctx context.Context, id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)

	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id))
}