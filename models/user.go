package models

import "time"

type User struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	Email        string    `josn:"email" bson:"email"`
	PasswordHash string    `json:"passwordHash" bson:"passwordHash"`
	RefreshToken string    `json:"refreshToken" bson:"refreshToken"`
	AccessToken  string    `json:"accessToken" bson:"accessToken"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
}
