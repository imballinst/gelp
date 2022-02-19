package gelp

import "fmt"

var colorReset = "\033[0m"
var colorBlue = "\033[36m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"

func GetBlueText(str string) string {
	return fmt.Sprint(string(colorBlue), str, string(colorReset))
}

func GetGreenText(str string) string {
	return fmt.Sprint(string(colorGreen), str, string(colorReset))
}

func GetRedText(str string) string {
	return fmt.Sprint(string(colorRed), str, string(colorReset))
}
