package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	Title     string    `bson:"title" json:"title"`
	Content   string    `bson:"content" json:"content"`
	Pinned    bool      `bson:"pinned" json:"pinned"`
	CreateAt  time.Time `bson:"createAt" json:"createAt'`
	UpdatedAt time.Time `bson:"UpdatedAt" json:"updatedAt"`
}

type CreateNoteRequest struct {
	Title   string `json:"title" binding: "required"`
	Content string `json:"content" binding:"required"`
	Pinned  bool   `json:"pinned"`
}

type UpdateNoteRequest struct {
	Title   string `json:"title" binding: "required"`
	Content string `json:"content" binding:"required"`
	Pinned  bool   `json:"pinned"`
}
