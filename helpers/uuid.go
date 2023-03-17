package helpers

import "github.com/google/uuid"

func Uuid() string {
	uuid := uuid.New().String()
	return uuid
}
