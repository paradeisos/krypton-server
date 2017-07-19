package options

import "time"

type ListTodoOpts struct {
	Page       int
	Limit      int
	From       time.Time
	To         time.Time
	Uid        string
	IsAll      bool
	IsFinished bool
}
