package models

import (
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

type Model struct {
	mux        sync.RWMutex
	session    *mgo.Session
	collection *mgo.Collection

	option  *Option
	indexes map[string]bool
}

const (
	MongoPoolMax     = 4096
	MongoSyncTimeout = 5
)

func NewModel(option *Option) *Model {
	dsn := "mongodb://"
	if option.User != "" && option.Password != "" {
		dsn += option.User + ":" + option.Password + "@"
	}
	dsn += option.Host
	if option.Database != "" {
		dsn += "/" + option.Database
	}

	session, err := mgo.Dial(dsn)
	if err != nil {
		beego.Error(err.Error())
		panic(err)
	}

	if err := session.Ping(); err != nil {
		beego.Error(err.Error())
		panic(err)
	}

	// set session mode
	switch option.Mode {
	case "Strong":
		session.SetMode(mgo.Strong, true)
	case "Monotonic":
		session.SetMode(mgo.Monotonic, true)
	case "Eventual":
		session.SetMode(mgo.Eventual, true)
	default:
		session.SetMode(mgo.Strong, true)
	}

	// set session safe
	session.SetSafe(&mgo.Safe{
		W:        1,
		WTimeout: 200,
	})

	// set pool size
	if option.Pool > 0 {
		if option.Pool > MongoPoolMax {
			option.Pool = MongoPoolMax
		}

		session.SetPoolLimit(option.Pool)
	}

	// set op response timeout
	if option.Timeout == 0 {
		option.Timeout = MongoSyncTimeout
	}
	session.SetSyncTimeout(time.Duration(option.Timeout) * time.Second)

	// panic as early as possible
	if err := session.Ping(); err != nil {
		panic(err.Error())
	}

	return &Model{
		session: session,
		option:  option,
		indexes: make(map[string]bool),
	}
}

func (model *Model) Use(database string) *Model {
	model.option.Database = database

	return model
}

func (model *Model) Copy() *Model {
	return &Model{
		session: model.session.Copy(),
		option:  model.option,
	}
}

func (model *Model) C(name string) *Model {
	copiedDB := model.Copy()
	copiedDB.collection = copiedDB.session.DB(model.Database()).C(name)

	return copiedDB
}

func (model *Model) Database() string {
	return model.option.Database
}

func (model *Model) Session() *mgo.Session {
	return model.session
}

func (model *Model) Collection() *mgo.Collection {
	return model.collection
}

func (model *Model) Query(collectionName string, collectionIndexes []mgo.Index, query func(*mgo.Collection)) {
	copiedDB := model.C(collectionName)
	defer copiedDB.Close()

	copiedCollection := copiedDB.Collection()

	if !model.indexes[collectionName] {
		model.mux.Lock()
		if !model.indexes[collectionName] {
			for _, index := range collectionIndexes {
				if err := copiedCollection.EnsureIndex(index); err != nil {
					model.indexes[collectionName] = false

					beego.Error(fmt.Sprintf("Ensure index of %s (%#v) : %v", collectionName, index, err))
				}
			}

			model.indexes[collectionName] = true
		}
		model.mux.Unlock()
	}

	query(copiedCollection)
}

func (model *Model) Close() {
	model.session.Close()
}
