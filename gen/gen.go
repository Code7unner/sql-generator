package gen

import (
	"math/rand"
	"time"
	"github.com/brianvoe/gofakeit"
)

const (
	MESSAGES_COUNT = 10000000
	CATEGORIES_COUNT = 5000
	USERS_COUNT = 500000
)

type Category struct {
	Id       string
	Name     string
	ParentId string
}

func generateCategories() ([]Category, []string) {
	var categoryIds []string
	var categories []Category

	for i := 0; i < CATEGORIES_COUNT; i++ {
		c := Category{
			Id: gofakeit.UUID(),
			Name: gofakeit.Company(),
		}

		if i == 0 {	c.ParentId = c.Id } else { c.Id = categoryIds[len(categoryIds) - 1] }

		categoryIds = append(categoryIds, c.Id)
		categories = append(categories, c)
	}

	return categories, categoryIds
}

type User struct {
	Id   string
	Name string
}

func generateUsers() ([]User, []string) {
	var userIds []string
	var users []User

	for i := 0; i < USERS_COUNT; i++ {
		user := User{
			Id:   gofakeit.UUID(),
			Name: gofakeit.Name(),
		}

		userIds = append(userIds, user.Id)
		users = append(users, user)
	}

	return users, userIds
}

type Message struct {
	Id         string `fake:"{ misc.uuid}"`
	Text       string `fake:"{ string.letter }"`
	CategoryId string
	PostedAt   time.Time `fake:"{ date.date }"`
	AuthorId   string
}

func generateMessages(userIds []string, categoryIds []string) []Message {
	var messages []Message

	for i := 0; i < MESSAGES_COUNT; i++ {
		message := Message{
			Id:         gofakeit.UUID(),
			Text:       gofakeit.ProgrammingLanguage(),
			CategoryId: categoryIds[rand.Intn(len(categoryIds))],
			PostedAt:   gofakeit.Date(),
			AuthorId:   userIds[rand.Intn(len(userIds))],
		}

		messages = append(messages, message)
	}

	return messages
}

func Generate() ([]Category, []User, []Message) {
	gofakeit.Seed(time.Now().UnixNano())

	categories, categoryIds := generateCategories()
	users, userIds := generateUsers()
	messages := generateMessages(userIds, categoryIds)

	return categories, users, messages
}
