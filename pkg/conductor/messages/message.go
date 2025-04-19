package messages

import "encoding/json"

type MessageInfo struct {
	TemplateName string            `json:"templateName"`
	DistEmail    string            `json:"distEmail"`
	FieldValues  map[string]string `json:"fieldValues"`
	Subject      string            `json:"subject"`
}

func (c *MessageInfo) ToJson() ([]byte, error) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return jsonData, nil

}
