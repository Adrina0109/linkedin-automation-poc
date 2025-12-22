package messaging

import "strings"

func RenderMessage(template, name string) string {
	msg := strings.ReplaceAll(template, "{{name}}", name)
	return msg
}
