package models

//响应体数据
type ResponseMod struct {
	Code string
	Message string
	Data interface{}
}