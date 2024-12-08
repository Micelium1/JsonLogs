package model

type ContextStruct struct {
	Ip string `json:"ip"`
}

type LogStruct struct {
	Date    string        `json:"date"`
	Level   string        `json:"level"`
	Message string        `json:"message"`
	Context ContextStruct `json:"context"`
	Logger  string        `json:"logger"`
}
