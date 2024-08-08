package template

import "fmt"

func Base(name string) string {
	return fmt.Sprintf("package response\n\n// Write Return parameter\ntype Write struct {\n\tStatus int         `json:\"status\"`           //Status Code\n\tMsg    string      `json:\"msg\"`              //Msg Prompt message\n\tResult interface{} `json:\"result,omitempty\"` //Data\n\tError  interface{} `json:\"error,omitempty\"`  //Error message\n}\n\n// Page Pagination return\ntype Page struct {\n\tTotal int64       `json:\"total\"` //Total total pages\n\tList  interface{} `json:\"list\"`  //List json data\n}\n\n// Success Successfully returned\nfunc Success(msg string, status int, data ...interface{}) *Write {\n\tvar lenData = len(data)\n\tif lenData == 1 {\n\t\treturn &Write{Status: status, Msg: msg, Result: data[0]}\n\t} else if lenData > 1 {\n\t\treturn &Write{Status: status, Msg: msg, Result: data}\n\t}\n\n\treturn &Write{Status: status, Msg: msg}\n}\n\n// Fail Error return, the second parameter is passed back to the front end and printed\nfunc Fail(msg string, status int, err ...string) *Write {\n\tvar lenErr = len(err)\n\n\tif lenErr == 1 {\n\t\treturn &Write{Status: status, Msg: msg, Error: err[0]}\n\t} else if lenErr > 1 {\n\t\treturn &Write{Status: status, Msg: msg, Error: err}\n\t}\n\n\treturn &Write{Status: status, Msg: msg}\n}\n")
	//package response
	//
	//// Write Return parameter
	//type Write struct {
	//	Status int         `json:"status"`           //Status Code
	//	Msg    string      `json:"msg"`              //Msg Prompt message
	//	Result interface{} `json:"result,omitempty"` //Data
	//	Error  interface{} `json:"error,omitempty"`  //Error message
	//}
	//
	//// Page Pagination return
	//type Page struct {
	//	Total int64       `json:"total"` //Total total pages
	//	List  interface{} `json:"list"`  //List json data
	//}
	//
	//// Success Successfully returned
	//func Success(msg string, status int, data ...interface{}) *Write {
	//var lenData = len(data)
	//if lenData == 1 {
	//return &Write{Status: status, Msg: msg, Result: data[0]}
	//} else if lenData > 1 {
	//return &Write{Status: status, Msg: msg, Result: data}
	//}
	//
	//return &Write{Status: status, Msg: msg}
	//}
	//
	//// Fail Error return, the second parameter is passed back to the front end and printed
	//func Fail(msg string, status int, err ...string) *Write {
	//var lenErr = len(err)
	//
	//if lenErr == 1 {
	//return &Write{Status: status, Msg: msg, Error: err[0]}
	//} else if lenErr > 1 {
	//return &Write{Status: status, Msg: msg, Error: err}
	//}
	//
	//return &Write{Status: status, Msg: msg}
	//}

}
