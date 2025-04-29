package messages

import "fmt"

func NewCodeMessageInfo(code string, distEmail string) *MessageInfo {
	return &MessageInfo{
		TemplateName: "code_template",
		DistEmail:    distEmail,
		Subject:      fmt.Sprintf("%s - код подтверждения на платформе Mirai.", code),
		FieldValues: map[string]string{
			"code": code,
		},
	}
}
