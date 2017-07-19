package models

import (
	"krypton-server/options"
	"testing"
	"time"

	"github.com/golib/assert"
	uuid "github.com/satori/go.uuid"
)

func Test_Todo_CRUD(t *testing.T) {
	assertion := assert.New(t)

	due := time.Now().Add(24 * time.Hour)
	uid := uuid.NewV4().String()
	title := "Buy Milk"
	content := "Go to market."

	todo := Todo.NewModel(uid, title, content, due)

	// create
	assertion.Nil(todo.Save())
	assertion.Equal(title, todo.Title)
	assertion.Equal(content, todo.Content)
	assertion.Equal(due, todo.Due)
	assertion.Equal(false, todo.Finished)

	// update
	todo.Finished = true
	assertion.Nil(todo.Save())
	assertion.Equal(true, todo.Finished)

	// find
	total, todos, err := Todo.List(&options.ListTodoOpts{
		Page:  1,
		Limit: 2,
	})
	assertion.Nil(err)
	assertion.Equal(1, len(todos))
	assertion.Equal(1, total)
	assertion.Equal(todo.Id, todos[0].Id)

	// delete
	assertion.Nil(Todo.Delete(todo.Id.Hex()))
}
