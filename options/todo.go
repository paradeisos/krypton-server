package options

import "time"

type CreateTodoOpts struct {
	Uid     string    `json:"uid"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Due     time.Time `json:"due"`
}

type ListTodoOpts struct {
	Uid        string    `json:"uid"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	From       time.Time `json:"from"`
	To         time.Time `json:"to"`
	IsFinished bool      `json:"is_finished"`
}

type UpdateTodoOpts struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Due      time.Time `json:"due"`
	Finished bool      `json:"finished"`
}
