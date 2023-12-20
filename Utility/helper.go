package Utility

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

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

func ConvertTimestampTo_YYYY_MM_DD(timeStamp int64) (time.Time, error) {
	var (
		err        error
		parsedTime time.Time
	)
	epochTime := time.Unix(timeStamp, 0)
	parsedTime, err = time.Parse("2006-01-02", epochTime.Format("2006-01-02"))
	if err != nil {
		err = errors.New("Encountered Error In Date Conversion ")
		return time.Time{}, err
	}
	return parsedTime, nil
}
