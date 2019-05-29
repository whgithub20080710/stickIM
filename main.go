package main

import (
	"net/http"
	"encoding/json"
	"log"
	"html/template"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var DbEngin *xorm.Engine
func init()  {
	drivename := "mysql"
	DsName := "root:123456@(127.0.0.1:3306)/chat?charset=utf8"
	DbEngin,err := xorm.NewEngine(drivename,DsName)
	if nil != err{
		log.Fatal(err.Error())
	}
	// 是否显示sql
	DbEngin.ShowSQL(true)
	// 设置最大打开连接数
	DbEngin.SetMaxOpenConns(2)

	fmt.Println("init data base ok")
}

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

func RegisterView()  {
	tpl,err := template.ParseGlob("view/**/*")
	if(err != nil){
		// 打印错误信息并退出
		log.Fatal(err.Error())
	}
	for _,v := range tpl.Templates(){
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer,tplname,nil)
		})
	}
}

func main() {

	// 绑定请求和处理函数
	http.HandleFunc("/user/login", userLogin)

	// 提供静态资源目录支持
	//http.Handle("/",http.FileServer(http.Dir(".")))
	http.Handle("/asset/",http.FileServer(http.Dir(".")))

	//// 登录页面
	//http.HandleFunc("/user/login.shtml", func(writer http.ResponseWriter, request *http.Request) {
	//	// 解析
	//	tpl,err := template.ParseFiles("view/user/login.html")
	//	if(err != nil){
	//		// 打印错误信息并退出
	//		log.Fatal(err.Error())
	//	}
	//	tpl.ExecuteTemplate(writer,"/user/login.shtml",nil)
	//})
	//
	//// 注册页面
	//http.HandleFunc("/user/register.shtml", func(writer http.ResponseWriter, request *http.Request) {
	//	// 解析
	//	tpl,err := template.ParseFiles("view/user/register.html")
	//	if(err != nil){
	//		// 打印错误信息并退出
	//		log.Fatal(err.Error())
	//	}
	//	tpl.ExecuteTemplate(writer,"/user/register.shtml",nil)
	//})

	RegisterView()
	// 启动web服务器
	http.ListenAndServe(":8080",nil)
}
