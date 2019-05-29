package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type H struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string)  {
	h := H{
		Code:code,
		Msg:msg,
		Data:data,
	}
	// 结构体转json字符串
	ret,err := json.Marshal(h)
	if(err != nil){
		log.Println(err.Error())
	}

	// header的Content-Type为application/json
	w.Header().Set("Content-Type","application/json")
	// 设置200状态
	w.WriteHeader(http.StatusOK)
	// 输出
	w.Write([]byte(ret))
}
func userLogin(writer http.ResponseWriter, request *http.Request) {
	// 数据库操作
	// 逻辑处理
	// restapi json/xml 返回
	// 获取前段传递的post表单参数
	// 普通的post表单请求，Content-Type=application/x-www-form-urlencoded
	// 解析参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	loginok := false
	if(mobile == "15812341234" && passwd == "000000"){
		loginok = true
	}
	if(loginok){
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer,0,data,"")
	}else{
		Resp(writer,-1,nil,"用户名密码错误")
	}

}

func main() {

	// 绑定请求和处理函数
	http.HandleFunc("/user/login", userLogin)

	// 提供静态资源目录支持
	//http.Handle("/",http.FileServer(http.Dir(".")))
	http.Handle("/asset/",http.FileServer(http.Dir(".")))

	// 启动web服务器
	http.ListenAndServe(":8080",nil)
}
