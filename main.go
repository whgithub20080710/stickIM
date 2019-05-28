package main

import (
	"net/http"
)

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", func(writer http.ResponseWriter, request *http.Request) {
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
		str := `{"code":0,"data":{"id":1,"token":"test"}}`
		if(!loginok){
			str = `{"code":-1,"msg":"账号或密码错误"}`
		}
		// header的Content-Type为application/json
		writer.Header().Set("Content-Type","application/json")
		// 设置200状态
		writer.WriteHeader(http.StatusOK)
		// 输出
		writer.Write([]byte(str))
	})
	// 启动web服务器
	http.ListenAndServe(":8080",nil)
}
