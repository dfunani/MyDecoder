package models

type EncodeMessage struct {
	Data map[string]any `json:"data"`
}

type DecodeMessage struct {
	Message string `json:"message"`
}
