package models

import (
	"time"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

var (
	Tomato *_Tomato

	tomatoCollection = "tomato"
	tomatoIndexes    = []mgo.Index{}
)

type _Tomato struct{}

type TomatoModel struct {
	Id        bson.ObjectId `bson:"_id" json:"-"`
	Uid       string        `bson:"uid" json:"uid"`
	Start     time.Time     `bson:"start" json:"start"`
	End       time.Time     `bson:"end" json:"end"`
	Desc      string        `bson:"desc" json:"desc"`
	Finished  bool          `bson:"finished" json:"finished"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`

	isNewRecord bool `bson:"-"`
}

func (_ *_Tomato) NewTomatoModel(uid string, start, end time.Time, desc string, finished bool) *TomatoModel {
	return &TomatoModel{
		Id:          bson.NewObjectId(),
		Uid:         uid,
		Start:       start,
		End:         end,
		Desc:        desc,
		Finished:    finished,
		isNewRecord: true,
	}
}

func (tomato *TomatoModel) Save() (err error) {
	if !tomato.Id.Valid() {
		return ErrInvalidId
	}

	Tomato.Query(func(c *mgo.Collection) {
		t := time.Now()

		if tomato.isNewRecord {
			tomato.CreatedAt = t
			tomato.UpdatedAt = t

			if err = c.Insert(tomato); err == nil {
				tomato.isNewRecord = false
			}
		} else {
			migration := bson.M{
				"$set": bson.M{
					"desc":       tomato.Desc,
					"start":      tomato.Start,
					"end":        tomato.End,
					"finished":   tomato.Finished,
					"updated_at": t,
				},
			}

			err = c.UpdateId(tomato.Id, migration)
		}
	})

	return
}

func (tomato *TomatoModel) Delete() (err error) {
	Tomato.Query(func(c *mgo.Collection) {
		err = c.RemoveId(tomato.Id)
	})

	return
}

func (_ *_Tomato) Find(id string) (tomato *TomatoModel, err error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidId
	}

	bsonID := bson.ObjectIdHex(id)

	Tomato.Query(func(c *mgo.Collection) {
		err = c.FindId(bsonID).One(&tomato)
	})

	return
}

func (_ *_Tomato) AllByUid(uid string) (tomatoes []*TomatoModel, err error) {
	if uid == "" || uid != "" && !bson.IsObjectIdHex(uid) {
		err = ErrInvalidId
		return
	}

	Tomato.Query(func(c *mgo.Collection) {
		err = c.Find(bson.M{
			"uid": uid,
		}).All(&tomatoes)
	})

	return
}

func (_ *_Tomato) Query(query func(c *mgo.Collection)) {
	Mongo().Query(tomatoCollection, tomatoIndexes, query)
}
