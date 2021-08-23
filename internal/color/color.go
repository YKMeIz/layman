package color

import "fmt"

const (
	whiteSchema  = "\033[97m"
	blueSchema   = "\033[34m"
	greenSchema  = "\033[97;32m"
	yellowSchema = "\033[33m"
	redSchema    = "\033[31m"
	resetSchema  = "\033[0m"
	boldSchema   = "\033[1m"
)

var (
	White = func(v ...interface{}) string {
		return whiteSchema + fmt.Sprint(v...) + resetSchema
	}

	Blue = func(v ...interface{}) string {
		return blueSchema + fmt.Sprint(v...) + resetSchema
	}

	Green = func(v ...interface{}) string {
		return greenSchema + fmt.Sprint(v...) + resetSchema
	}

	Yellow = func(v ...interface{}) string {
		return yellowSchema + fmt.Sprint(v...) + resetSchema
	}

	Red = func(v ...interface{}) string {
		return redSchema + fmt.Sprint(v...) + resetSchema
	}

	Bold = func(v ...interface{}) string {
		return boldSchema + fmt.Sprint(v...) + resetSchema
	}
)
