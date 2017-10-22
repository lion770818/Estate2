package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	//"time"

	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"

	"encoding/json"

	//"code.google.com/p/go.net/websocket"
	"golang.org/x/net/websocket"

	"./src/One1CloudLib"

	"html/template"
)

/*
// #cgo LDFLAGS: -L. -lDllTest -lstdc++
// #cgo CFLAGS: -I ./
// #include <stdio.h>
// #include <stdlib.h>
// #include "DllTest.h"
import "C"
*/

//
// golang调用c++文件 http://www.cnblogs.com/lavin/p/golang-call-cpp-or-cc-file.html

//=========================================================================================================================================================================
// 參考 Go语言用WebSocket的简单例子          http://www.cnblogs.com/ghj1976/archive/2013/04/22/3035592.html
// 且戰且走HTML5(2) 應用主軸：WebSocket      http://ithelp.ithome.com.tw/articles/10102394
// 指令									 http://www.wklken.me/posts/2014/03/02/01-intro.html
// package								   https://openhome.cc/Gossip/Go/Package.html

//=========
// html
// http://blog.jex.tw/blog/2014/03/08/golang-html/
const tmpl = `
<html>
    <head>
        <title></title>
    </head>
    <body>

    </body>
</html>`

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func tHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("ex").Parse(tmpl))
	v := map[string]interface{}{
		"Title": "Test <b>World</b>",
		"Body":  template.HTML("Hello <b>World</b>"),
	}
	t.Execute(w, v)
}

// http://127.0.0.1:5678/
// http://127.0.0.1:5678/hello
func Html_Test() {

	One1CloudLib.CommonLog_INFO_Println("#Html_Test start...")

	http.HandleFunc("/", tHandler)
	http.HandleFunc("/hello", hello)

	err := http.ListenAndServe(":5678", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//=========================================================================================================================================================================
// 網路層 訊息主要進入點
// server ip = ws://192.168.1.21:1234/One1CloudGameCmd
// 測試cmd   = {"sys":"system", "cmd":"login", "sn":12345, "isEncode":false,"data":"{\"PlatformID\":1,\"GameID\":0,\"Account\":\"cat111\",\"Password\":\"1234\"}"}
func WS_One1CloudGameCmd(ws *websocket.Conn) {
	var err error
	var ret bool                                 // 是否完成client的 connect
	var Response One1CloudLib.CommonResponseInfo // 共用回應結構
	var socketClientInfo One1CloudLib.ClientConn // 共用socketClient結構
	var receiveMsg string = ""                   // 底層接收到的字串
	var responseMsg string = "Unknow"            // 回應的結果

	One1CloudLib.CommonLog_INFO_Println("#websocket WS_One1GameCmd")

	// 增加client
	socketClientInfo, ret = One1CloudLib.Common_ClientAdd(ws)
	One1CloudLib.CommonLog_INFO_Printf("底層收到連線請求 Client connect ret=%t...", ret)

	if ret {
		for {

			One1CloudLib.CommonLog_INFO_Println("底層 Client Receive Data...")
			if err = websocket.Message.Receive(ws, &receiveMsg); err != nil {

				One1CloudLib.CommonLog_INFO_Printf("底層 有Client斷線或關閉瀏覽器了 ClientID=%d, ClientIP=%s", socketClientInfo.ClientID, socketClientInfo.ClientIP)
				//One1CloudLib.Common_ClientDelete(socketClientInfo.ClientID)
				One1CloudLib.Common_ClientClose(socketClientInfo.ClientID, true) // 玩家斷線或關閉瀏覽器
				break
			}

			One1CloudLib.CommonLog_INFO_Printf("底層收到的資料 Received back from client: " + receiveMsg)

			// 丟到 analysis 去 分析命令並執行, 再傳回結果
			Response = One1CloudLib.Common_Analysis(socketClientInfo.ClientID, socketClientInfo.ClientIP, receiveMsg)

			//------------------------------------------------
			// 組成回應格式
			DataMsgByte, err := json.Marshal(Response)
			if err != nil {
				One1CloudLib.CommonLog_WARNING_Println("底層解析json 失敗  err:", err)
			}
			responseMsg := string(DataMsgByte)
			//responseMsgTmp := string(DataMsgByte)
			//responseMsg = strings.Replace(responseMsgTmp, "\\", "", -1)

			One1CloudLib.CommonLog_INFO_Printf("底層回應格式 ClientID=%d, clientIP=%s, 回應訊息responseMsg=%s", socketClientInfo.ClientID, socketClientInfo.ClientIP, responseMsg)
			if err = websocket.Message.Send(ws, responseMsg); err != nil {
				One1CloudLib.CommonLog_WARNING_Println("底層 網路異常 Can't send")
				break
			}

		}
	} else {

		// 服務器上限滿額
		var Code int = int(One1CloudLib.ERROR_CODE_CLIENT_TOO_MATCH)
		Response.Code = One1CloudLib.ErrorCode[Code].Code
		Response.Message = One1CloudLib.ErrorCode[Code].Message
		//------------------------------------------------
		// 組成回應格式
		DataMsgByte, err := json.Marshal(Response)
		if err != nil {
			One1CloudLib.CommonLog_WARNING_Println("json2 err:", err)
		}
		responseMsg = string(DataMsgByte)

		One1CloudLib.CommonLog_INFO_Printf("回應格式 clientIP=%s, 回應訊息responseMsg=%s", socketClientInfo.ClientIP, responseMsg)
		if err = websocket.Message.Send(ws, responseMsg); err != nil {
			One1CloudLib.CommonLog_WARNING_Println("底層解析json2 失敗  err:", err)
		}

		// 主動斷線
		err = ws.Close()
		if err != nil {
			One1CloudLib.CommonLog_WARNING_Println("底層 網路異常 Can't Close")
		}
	}

}

//=========================================================================================================================================================================
// websocket 初始化
func websocket_Init() {

	One1CloudLib.CommonLog_INFO_Println("#websocket_Init begin")

	http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line

	// 有修改了交握檢查 checkOrigin
	http.Handle("/One1CloudGameCmd", websocket.Handler(WS_One1CloudGameCmd))

	// 監聽的port 日後變成 讀取txt 檔案
	server_port := fmt.Sprintf(":%d", One1CloudLib.ServerConfig.Server_Port)
	One1CloudLib.CommonLog_INFO_Printf("#websocket_Init server_port %s", server_port)
	if err := http.ListenAndServe(server_port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	One1CloudLib.CommonLog_INFO_Println("#websocket_Init end")
}

//============================================================================================================
// 建構子
func init() {
	One1CloudLib.CommonLog_INFO_Println("============================golang啟動! 建構子Init...")

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func Test() {

	log.Printf("============== http service Test 韓式  開始 ===============")
	http.HandleFunc("/aaa", sayhelloName)    //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Printf("============== http service Test 韓式  結束 ===============")
}

//============================================================================================================
// 主進入點
func main() {

	One1CloudLib.CommonLog_INFO_Println("============================golang啟動! Start...")

	//a := C.add(1, 2)
	//a := C.Add(1000, 2)
	//One1CloudLib.CommonLog_INFO_Printf("讀取 c++ a=%d", a)
	// 初始化 物件
	One1CloudLib.Common_Init()

	//go Test()
	//go Html_Test()

	// websocket 模組啟動
	websocket_Init()

	One1CloudLib.CommonLog_INFO_Println("============================golang啟動! End...")
}
