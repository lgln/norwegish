package norwegish

import "strings"

func Translate(s string) string {
	r := strings.NewReplacer("o", "ø", "c", "k", "th", "t", "w", "v", "W", "V")
	return r.Replace(s)
}
