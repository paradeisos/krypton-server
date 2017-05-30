package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	User *_User

	userCollection = "user" // NOTE: this should migrate to user
	userIndexes    = []mgo.Index{
		{
			Key:    []string{"email"},
			Unique: true,
		},
		{
			Key:    []string{"username"},
			Unique: true,
		},
	}
)

type _User struct{}

func (_ *_User) NewUserModel(username, email, password, desc string) *UserModel {
	now := time.Now()
	return &UserModel{
		Id:          bson.NewObjectId(),
		Username:    username,
		Email:       email,
		Password:    password,
		Status:      UserStatusActive,
		Description: desc,
		CreatedAt:   now,
		UpdatedAt:   now,
		isNewRecord: true,
	}
}

type UserModel struct {
	Id          bson.ObjectId `bson:"_id" json:"-"`
	Username    string        `bson:"username" json:"username"`
	Email       string        `bson:"email" json:"email"`
	Password    string        `bson:"password" json:"password"`
	Status      string        `bson:"status" json:"status"`
	Description string        `bson:"description" json:"descriptioon"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`

	isNewRecord bool `bson:"-"`
}

func (user *UserModel) Save() (err error) {
	if !user.Id.Valid() {
		return ErrInvalidId
	}

	User.Query(func(c *mgo.Collection) {
		t := time.Now()

		if user.isNewRecord {
			user.CreatedAt = t
			user.UpdatedAt = t

			err = c.Insert(user)
			if err == nil {
				user.isNewRecord = false
			}

		} else {
			migration := bson.M{
				"$set": bson.M{
					"username":   user.Username,
					"password":   user.Password,
					"status":     user.Status,
					"updated_at": t,
				},
			}

			err = c.UpdateId(user.Id, migration)
		}
	})

	return
}

func (_ *_User) FindByUsername(username string) (user *UserModel, err error) {
	User.Query(func(c *mgo.Collection) {
		err = c.Find(bson.M{
			"username": username,
		}).One(&user)
	})

	return
}

func (_ *_User) FindByEmail(email string) (user *UserModel, err error) {
	User.Query(func(c *mgo.Collection) {
		err = c.Find(bson.M{
			"email": email,
		}).One(&user)
	})

	return
}

func (_ *_User) Query(query func(c *mgo.Collection)) {
	Mongo().Query(userCollection, userIndexes, query)
}
