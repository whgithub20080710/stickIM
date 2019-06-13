package main

import (
	"net/http"

	"./ctrl"
	"log"
	"html/template"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterView(){
	//一次解析出全部模板
	tpl,err := template.ParseGlob("view/**/*")
	if nil!=err{
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _,v := range tpl.Templates(){
		//
		tplname := v.Name();
		fmt.Println("HandleFunc     "+v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			//
			fmt.Println("parse     "+v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w,tplname,nil)
			if err!=nil{
				log.Fatal(err.Error())
			}
		})
	}

}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)

	//1 提供静态资源目录支持
	//http.Handle("/", http.FileServer(http.Dir(".")))

	//2 指定目录的静态文件
	http.Handle("/asset/",http.FileServer(http.Dir(".")))

	RegisterView()

	http.ListenAndServe(":8080",nil)
}