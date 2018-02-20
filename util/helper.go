package util

import (
	"github.com/satori/go.uuid"
	"strconv"
)

func GenerateUserID() (string, error) {

	userId, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return userId.String(), nil
}

func ConvertStringToUint(origin string) uint {

	uid64, _ := strconv.ParseUint(origin, 10, 32)
	change := uint(uid64)

	return change
}

func ConvertStringToInt(origin string) int {

	idInt, _ := strconv.Atoi(origin)

	return idInt
}

func GetLimitAndOffset(perPage string, page string) (int, int) {

	limit, _ := strconv.Atoi(perPage)
	offset, _ := strconv.Atoi(page)

	if offset == 1 {
		offset = 0
	} else {
		offset = (offset * limit) - limit
	}

	return limit, offset
}
