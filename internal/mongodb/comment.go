package mongodb

import "time"

type Comment struct {
	Username  string    `bson:"username"`
	Content   string    `bson:"content"`
	CreatedAt time.Time `bson:"created_at"`
}
