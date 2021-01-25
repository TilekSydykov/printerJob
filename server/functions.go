package server

import(
	"strings"
)

func parseString(p string) (string, string){
	p = strings.ReplaceAll(p, "\n", "")
	p = strings.ReplaceAll(p, "\r", "")
	p = strings.ReplaceAll(p, "\f", "")
	slice := strings.Split(p, "=")
	if(len(slice)<2){
		return "", ""
	}
	return slice[0], slice[1]
}