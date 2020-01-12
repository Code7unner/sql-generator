package gen

import (
	"github.com/brianvoe/gofakeit"
	"log"
	"math/rand"
	"sql-generator/db"
	"sync"
	"time"
)

const (
	MESSAGES_COUNT = 10000000
	CATEGORIES_COUNT = 5000
	USERS_COUNT = 500000
)

var (
	users 		[]*User
	messages 	[]*Message
	categories  []*Category
)

type Gen struct {
	p *db.Postgres
	mu *sync.Mutex
	wg *sync.WaitGroup
}

func InitGenerator(p *db.Postgres) (*Gen, error) {
	return &Gen{
		p: p,
		mu: &sync.Mutex{},
		wg: &sync.WaitGroup{},
	}, nil
}

func (g *Gen) Generate() {
	gofakeit.Seed(time.Now().UnixNano())

	g.generateCategories()
	log.Printf("Successfully created %v categories.\n", len(categories))

	g.generateUsers()
	log.Printf("Successfully created %v users.\n", len(users))

	g.generateMessages()
	log.Printf("Successfully created %v messages.\n", len(messages))
}

type Category struct {
	Id       string
	Name     string
	ParentId string
}

func (g *Gen) generateCategories() {
	for i := 0; i < CATEGORIES_COUNT; i ++ {
		c := g.generateCategory()
		categories = append(categories, c)
	}
}

func (g *Gen) generateCategory() *Category {
	category :=  &Category{
		Id: gofakeit.UUID(),
		Name: gofakeit.Company(),
	}
	category.ParentId = category.Id

	return category
}

type User struct {
	Id   string
	Name string
}

func (g *Gen) generateUsers() {
	for i := 0; i < USERS_COUNT; i++ {
		u := g.generateUser()
		users = append(users, u)
	}
}

func (g *Gen) generateUser() *User {
	user := &User{
		Id:   gofakeit.UUID(),
		Name: gofakeit.Name(),
	}

	return user
}

type Message struct {
	Id         string    `fake:"{ misc.uuid}"`
	Text       string 	 `fake:"{ string.letter }"`
	CategoryId string
	PostedAt   time.Time `fake:"{ date.date }"`
	AuthorId   string
}

func (g *Gen) generateMessages() {
	for i := 0; i < MESSAGES_COUNT; i++ {
		m := g.generateMessage()
		messages = append(messages, m)
	}
}

func (g *Gen) generateMessage() *Message {
	category := categories[rand.Intn(len(categories))]
	author := users[rand.Intn(len(users))]

	message := &Message{
		Id:         gofakeit.UUID(),
		Text:       gofakeit.ProgrammingLanguage(),
		CategoryId: category.Id,
		PostedAt:   gofakeit.Date(),
		AuthorId:   author.Id,
	}

	return message
}