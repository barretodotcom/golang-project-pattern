package uuid

import (
	"fmt"

	"github.com/google/uuid"
)

func GetId() string {
	uuid, err := uuid.NewRandom()

	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}

	return uuid.String()
}
