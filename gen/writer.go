package gen

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

const GOPHERS = 10

func (g *Gen) InsertUsers() error {
	transaction, err := g.p.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := transaction.Prepare(pq.CopyInSchema("public", "users", "id", "name"))
	if err != nil {
		return err
	}

	g.wg.Add(USERS_COUNT)
	if n := USERS_COUNT / GOPHERS; n >= 1 {
		for i := 0; i < GOPHERS; i++ {
			go func() {
				for j := 0; j < n; j++ {
					g.insertUser(stmt)
				}
			}()
		}
	} else {
		for i := 0; i < USERS_COUNT; i++ {
			g.insertUser(stmt)
		}
	}
	g.wg.Wait()

	err = g.closeTransaction(transaction, stmt)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gen) insertUser(stmt *sql.Stmt) {
	g.mu.Lock()
	defer func() {
		g.mu.Unlock()
		g.wg.Done()
	}()

	var user *User
	user, users = users[0], users[1:]

	_, err := stmt.Exec(user.Id, user.Name)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) InsertCategories() error {
	transaction, err := g.p.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := transaction.Prepare(pq.CopyInSchema("public", "categories", "id", "name", "parent_id"))
	if err != nil {
		return nil
	}

	g.wg.Add(CATEGORIES_COUNT)
	if n := CATEGORIES_COUNT / GOPHERS; n >= 1 {
		for i := 0; i < GOPHERS; i++ {
			go func() {
				for j := 0; j < n; j++ {
					g.insertCategory(stmt)
				}
			}()
		}
	} else {
		for i := 0; i < CATEGORIES_COUNT; i++ {
			go g.insertCategory(stmt)
		}
	}
	g.wg.Wait()

	err = g.closeTransaction(transaction, stmt)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gen) insertCategory(stmt *sql.Stmt) {
	g.mu.Lock()
	defer func() {
		g.mu.Unlock()
		g.wg.Done()
	}()

	var category *Category

	category, categories = categories[0], categories[1:]

	_, err := stmt.Exec(category.Id, category.Name, category.ParentId)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) InsertMessages() error {
	transaction, err := g.p.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := transaction.Prepare(pq.CopyInSchema("public", "messages", "id", "text", "category_id", "posted_at", "author_id"))
	if err != nil {
		return err
	}

	g.wg.Add(MESSAGES_COUNT)
	if n := MESSAGES_COUNT / GOPHERS; n >= 1 {
		for i := 0; i < GOPHERS; i++ {
			go func() {
				for j := 0; j < n; j++ {
					g.insertMessage(stmt)
				}
			}()
		}
	} else {
		for i := 0; i < MESSAGES_COUNT; i++ {
			go g.insertMessage(stmt)
		}
	}
	g.wg.Wait()
	
	err = g.closeTransaction(transaction, stmt)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gen) insertMessage(stmt *sql.Stmt) {
	g.mu.Lock()
	defer func() {
		g.mu.Unlock()
		g.wg.Done()
	}()

	var message *Message

	message, messages = messages[len(messages)-1], messages[:len(messages)-1]

	_, err := stmt.Exec(message.Id, message.Text, message.CategoryId, message.PostedAt, message.AuthorId)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gen) closeTransaction(transaction *sql.Tx, stmt *sql.Stmt) error {
	_, err := stmt.Exec()
	if err != nil {
		return transaction.Rollback()
	}
	err = stmt.Close()
	if err != nil {
		return transaction.Rollback()
	}
	err = transaction.Commit()
	if err != nil {
		return transaction.Rollback()
	}

	return nil
}
