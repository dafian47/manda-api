package util

import (
	"github.com/satori/go.uuid"
	"strconv"
)

func GenerateChannelID() (string, error) {

	channelID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return channelID.String(), nil
}

func GenerateThreadID() (string, error) {

	threadID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return threadID.String(), nil
}

func GenerateCommentID() (string, error) {

	commentID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return commentID.String(), nil
}

func GenerateUserID() (string, error) {

	userID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return userID.String(), nil
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
