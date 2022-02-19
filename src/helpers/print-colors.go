package gelp

import "fmt"

var colorReset = "\033[0m"
var colorBlue = "\033[36m"

func GetBlueText(str string) string {
	return fmt.Sprint(string(colorBlue), str, string(colorReset))
}
