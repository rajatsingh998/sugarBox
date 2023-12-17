package Utility

import "github.com/google/uuid"

func GenerateUUID() (string, error) {
	var (
		newId uuid.UUID
		err   error
	)
	if newId, err = uuid.NewUUID(); err != nil {
		return "", err
	}
	return newId.String(), nil
}
