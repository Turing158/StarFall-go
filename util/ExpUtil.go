package util

import "starfall-go/entity"

func GetMaxExp(level int64) int64 {
	switch level {
	case 1:
		return 100
	case 2:
		return 530
	case 3:
		return 1200
	case 4:
		return 2140
	case 5:
		return 3310
	case 6:
		return 4710
	case 7:
		return 6340
	case 8:
		return 8200
	case 9:
		return 10290
	case 10:
		return 12610
	}
	return 999999
}

func CheckAndLevelUp(currentExp, currentLevel int64) int64 {
	return currentExp - GetMaxExp(currentLevel)
}

func SetUserMaxExp(user *entity.UserOut) {
	user.MaxExp = GetMaxExp(user.Level)
}

func SetTopicUserMaxExp(topicOut *entity.TopicOut) {
	topicOut.MaxExp = GetMaxExp(topicOut.Level)
}
