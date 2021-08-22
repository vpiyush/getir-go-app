// Package model defines the types which reflect the structure of
// database reponses
package models

import (
	"time"
)

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

type Pair struct {
	Key   string `validate:"required"json:"key"`
	Value string `validate:"required"json:"value"`
}
