package helpers

import (
	"strconv"

	"github.com/gofrs/uuid"
)

//StringToInt converts string to int
func StringToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

//StringToUUID converts string to uuid
func StringToUUID(str string) (uuid.UUID, error) {
	u, err := uuid.FromString(str)
	if err != nil {
		return uuid.Nil, err
	}
	return u, nil
}
