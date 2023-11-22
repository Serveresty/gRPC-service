package utils

import "github.com/enescakir/emoji"

func GetEmojiById(id int64) string {
	switch id {
	case 1:
		return string(emoji.House)
	case 2:
		return string("")
	case 3:
		return string(emoji.Airplane)
	case 4:
		return string(emoji.Airplane)
	case 5:
		return string(emoji.Thermometer)
	case 6:
		return string(emoji.Thermometer)
	case 7:
		return string("")
	case 8:
		return string("")
	case 9:
		return string(emoji.GraduationCap)
	case 10:
		return string(emoji.House)
	case 11:
		return string(emoji.Sun)
	case 12:
		return string(emoji.Sun)
	case 13:
		return string(emoji.Sun)
	}
	return ""
}
