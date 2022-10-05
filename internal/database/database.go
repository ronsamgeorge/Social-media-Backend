package database

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Client struct {
	path string
}
type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

//creates an instance of 'Client' and returns it 
func NewClient(path string) *Client{
	clientInstance := Client{path: path}
	return &clientInstance
}

//Create a New DB at the path associated with the Client instance
func (c Client) createDB() error{

	data, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})

	if err != nil{
		return err
	}

	err = os.WriteFile(c.path,data,0600)

	if err != nil {
		err2 := c.createDB()
		return err2
	}

	return nil
}


func (c Client) EnsureDB() error{
	_,err := os.ReadFile(c.path)

	if errors.Is(err, os.ErrNotExist){
		return c.createDB()
	}

	return nil
}