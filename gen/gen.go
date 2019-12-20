package gen

import (
	"time"
)

const (
	MESSAGES_COUNT = 10000000
	CATEGORIES_COUNT = 5000
	USERS_COUNT = 500000
)

type Message struct {
	Id         string `fake:"{ misc.uuid}"`
	Text       string `fake:"{ string.letter }"`
	CategoryId string
	PostedAt   time.Time `fake:"{ date.date }"`
	AuthorId   string
}

type Category struct {
	Id       string
	Name     string
	ParentId string
}

type User struct {
	Id   string
	Name string
}
