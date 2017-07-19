package models

import (
	"time"

	"krypton-server/options"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	Todo *_Todo

	todoCollection = "todo"
	todoIndexes    = []mgo.Index{}
)

type _Todo struct{}

type TodoModel struct {
	Id        bson.ObjectId `bson:"_id" json:"-"`
	Uid       string        `bson:"uid" json:"uid"`
	Title     string        `bson:"title" json:"title"`
	Content   string        `bson:"content" json:"content"`
	Due       time.Time     `bson:"due" json:"due"`
	Finished  bool          `bson:"finished" json:"finished"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`

	isNewRecord bool `bson:"-"`
}

func (_ *_Todo) NewModel(uid, title, content string, due time.Time) *TodoModel {
	return &TodoModel{
		Id:          bson.NewObjectId(),
		Uid:         uid,
		Title:       title,
		Content:     content,
		Due:         due,
		Finished:    false,
		isNewRecord: true,
	}
}

func (todo *TodoModel) Save() (err error) {
	if !todo.Id.Valid() {
		return ErrInvalidId
	}

	Todo.Query(func(c *mgo.Collection) {
		t := time.Now()

		if todo.isNewRecord {
			todo.CreatedAt = t
			todo.UpdatedAt = t

			if err = c.Insert(todo); err == nil {
				todo.isNewRecord = false
			}
		} else {
			migration := bson.M{
				"$set": bson.M{
					"title":      todo.Title,
					"content":    todo.Content,
					"due":        todo.Due,
					"finished":   todo.Finished,
					"updated_at": t,
				},
			}

			err = c.UpdateId(todo.Id, migration)
		}
	})

	return
}

func (_ *_Todo) Find(id string) (todo *TodoModel, err error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidId
	}

	bsonID := bson.ObjectIdHex(id)

	Todo.Query(func(c *mgo.Collection) {
		err = c.FindId(bsonID).One(&todo)
	})

	return
}

func (_ *_Todo) List(opts *options.ListTodoOpts) (total int, todos []*TodoModel, err error) {
	offset := (opts.Page - 1) * opts.Limit
	if offset < 0 {
		offset = 0
	}

	query := bson.M{}
	if opts.From.After(opts.To) {
		return 0, nil, ErrInvalidParams
	}

	dueRange := bson.M{}
	if !opts.From.IsZero() {
		dueRange["$gte"] = opts.From
	}
	if !opts.To.IsZero() {
		dueRange["$lte"] = opts.To
	}
	if len(dueRange) > 0 {
		query["due"] = dueRange
	}

	if opts.Uid != "" {
		query["uid"] = opts.Uid
	}

	if !opts.IsAll {
		query["finished"] = opts.IsFinished
	}

	Todo.Query(func(c *mgo.Collection) {
		err = c.Find(query).Skip(offset).Limit(opts.Limit).Sort("due").All(&todos)
		if err == nil {
			total, err = c.Find(query).Count()
		}
	})

	return
}

func (_ *_Todo) Delete(id string) (err error) {
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidId
	}

	Todo.Query(func(c *mgo.Collection) {
		err = c.RemoveId(bson.ObjectIdHex(id))
	})

	return
}

func (_ *_Todo) Query(query func(c *mgo.Collection)) {
	Mongo().Query(todoCollection, todoIndexes, query)
}
