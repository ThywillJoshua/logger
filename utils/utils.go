package utils

import (
	"fmt"
	"strings"
)

func RemoveSlashesAndConvertToString(v any) string  {
	var str string
    str = fmt.Sprint(v)
    
    str = strings.ReplaceAll(str, "/", "")
    str = strings.ReplaceAll(str, "\\", "")

    return str
}