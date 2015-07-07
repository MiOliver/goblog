package models

import ()

const (
	SuccessCode int = 200
	FailCode    int = 201
)

type Header struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

type Response struct {
	Head Header      `json:"header"`
	Body interface{} `json:"body"`
}

func CreateResponse(code int, desc string, body interface{}) Response {
	var mRes Response
	var mHeader Header
	var mBody interface{}
	mHeader.Code = code
	mHeader.Desc = desc
	mBody = body
	mRes.Head = mHeader
	mRes.Body = mBody
	// b, err := json.Marshal(mRes)
	// if err != nil {
	// 	fmt.Println("createResponse error!")
	// 	return "error"
	// }
	// fmt.Println(string(b))
	return mRes
}
