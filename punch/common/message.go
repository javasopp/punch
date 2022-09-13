package common

type Message struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Count   int         `json:"count"`
}

func Success() *Message {
	return &Message{
		Code:    codeSuccess,
		Message: messageSuccess,
	}
}

func SuccessWithData(data interface{}, counts ...int) *Message {
	count := 0
	if counts != nil {
		count = counts[0]
	}
	return &Message{
		Code:    codeSuccess,
		Message: messageSuccess,
		Data:    data,
		Count:   count,
	}
}

func Error() *Message {
	return &Message{
		Code:    codeError,
		Message: messageError,
	}
}

func NoData() *Message {
	return &Message{
		Code:    codeNoData,
		Message: messageNoData,
	}
}

func NoPermission() *Message {
	return &Message{
		Code:    codeNoPermission,
		Message: messageNoPermission,
	}
}
