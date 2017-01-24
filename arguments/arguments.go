package arguments

import "strings"

const file_arg_prefix = "--file="
const DefaultFile = "timio.yml"

func GetConfigFile(argsWithoutProg []string) string {
	for _, element := range argsWithoutProg {
		if strings.HasPrefix(element, file_arg_prefix) {
			return strings.Replace(element, file_arg_prefix, "", 1)
		}
	}
	return DefaultFile
}