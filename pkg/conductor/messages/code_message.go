package messages

import "fmt"

func NewCodeMessageInfo(code string, distEmail string) *MessageInfo {
	return &MessageInfo{
		TemplateName: "code_template",
		DistEmail:    distEmail,
		Subject:      fmt.Sprintf("%s - authorization code to log in", code),
		FieldValues: map[string]string{
			"code": code,
		},
	}
}
