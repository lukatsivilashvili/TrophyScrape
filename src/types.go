package main

import "math/rand"

type Account struct {
	ID                int    `json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	TotalAchievements int64  `json:"totalAchievements"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		FirstName: firstName,
		LastName:  lastName,
	}
}
