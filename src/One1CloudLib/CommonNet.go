package One1CloudLib

import (

	//"fmt"
	"log"
	//"time"
	"encoding/json"
	"io/ioutil"
	"sync"

	"golang.org/x/net/websocket"
)

const VERSION string = "1.0.0" // Server 版本
const CLIENT_MAX int = 5000    // 最大的client連線數, 超過此連線束後, 一律往最後面塞, 但是找空位時, 只會找 0~5000, 所以先不要超過 CLIENT_MAX 值阿 XD

type ClientConn struct {
	IsUse     bool            // 是否有使用
	websocket *websocket.Conn // websocket 物件 (送封包+斷線+...)
	ClientIP  string          // client端的IP位置
	ClientID  int             // client端分配到的 陣列 idx
	Sn        int64           // 封包序號
	Member    MemberInfo      // 玩家資料

	DataMutex *sync.RWMutex // 資料同步用的鎖
}

/*
var (
	ClientList = make(map[int]ClientConn) // map Containing Clients   @__@ map 不太會用阿 取記憶體部分...

)*/
var ClientCount = 0
var ClientList [CLIENT_MAX]ClientConn

//=========================================================================================================================================================================
// server 環境控制檔案
type Config struct {
	Server_Port           int    // server port
	Mysql_IP              string // mysql ip
	Mysql_Account         string // 帳號
	Mysql_Password        string // 密碼
	Mysql_DB              string // 連線db
	Mysql_SetMaxOpenConns int    // 執行緒用
	Mysql_SetMaxIdleConns int
}

var ServerConfig Config // server 環境控制檔案

func init() {

	CommonLog_INFO_Printf("#CommonFunction init... ClientCount=%d", ClientCount)
	ClientCount = 0

	// 讀取Server文字檔
	ServerConfig = Config_ReadFile("./config.txt")
}

var DEBUG_TEST_MUTEX_1 bool = false

var DEBUG_TEST_MUTEX_0 bool = true // 最大 最沒效率的鎖

var DataMutex *sync.RWMutex // 資料同步用的鎖

//=========================================================================================================================================================================
// 增加一個client
func Common_ClientAdd(ws *websocket.Conn) (ClientConn, bool) {

	var ret bool = false
	var memberInfo MemberInfo               // 空的會員資料
	var ClientID = Common_IdleClientIDGet() // 找到空閒的 ClientID
	CommonLog_INFO_Printf("#Common_ClientAdd 增加玩家連線... ClientCount=%d 找到的ClientID=%d", ClientCount, ClientID)

	client := ws.Request().RemoteAddr
	log.Println("Client connected:", client)
	socketClientInfo := ClientConn{true, ws, client, ClientID, 0, memberInfo, new(sync.RWMutex)}

	if ClientID != -1 {

		// 將Client 連線資訊塞入ClientList內
		ClientList[ClientID] = socketClientInfo
		ret = true

		ClientCount = ClientCount + 1
		CommonLog_INFO_Printf("#Common_ClientAdd Connect Success 增加玩家連線成功... ClientCount=%d, ClientIP=%s, ClientID=%d", ClientCount, socketClientInfo.ClientIP, ClientID)
	} else {
		ret = false
		CommonLog_WARNING_Printf("#Common_ClientAdd Connect Fail 增加玩家連線失敗...  CLIENT_MAX=%d, ClientCount=%d", CLIENT_MAX, ClientCount)
	}

	return socketClientInfo, ret
}

//=========================================================================================================================================================================
// 刪除一個client
func Common_ClientDelete(ClientID int) bool {

	var ret bool = false
	var member MemberInfo // 空的會員資料
	CommonLog_INFO_Printf("#Common_ClientDelete 刪除玩家連線... ClientID=%d", ClientID)

	// 取出記憶體位置
	//Client := &ClientList[ClientID]
	pClient := Common_ClientInfoGet(ClientID)

	if pClient.IsUse == true {

		pClient.IsUse = false
		pClient.Member = member // 清除會員資料
		pClient.ClientID = -1
		pClient.ClientIP = ""

		ClientCount = ClientCount - 1
		pClient.Sn = 0
		ret = true

		CommonLog_INFO_Printf("#Common_ClientDelet 目前連線數目=%d, 玩家登出=%s, ClientID=%d", ClientCount, pClient.ClientIP, pClient.ClientID)
	} else {

		CommonLog_WARNING_Printf("#Common_ClientDelet Warning... 查無此ClientID=%d", ClientID)
		ret = false
	}

	return ret
}

//=========================================================================================================================================================================
// 關閉一個client
func Common_ClientClose(ClientID int, CleanData bool) bool {

	var Code int = int(ERROR_CODE_SUCCESS)   // 回應值
	var ResponseFinishData string = "unknow" // 回應的資料格式(加工後的最終版)
	var ResponseTmpData string = "unknow"    // 回應的資料格(暫時)
	var BroadcastTable TableInfo             // 要廣播的桌子
	var BroadcastMember MemberInfo           // 發起廣播的玩家
	var ClientIP string = ""                 // client的ip
	var ret bool = false
	CommonLog_INFO_Printf("#Common_ClientClose 開始 關閉一個client... ClientID=%d, CleanData=%t", ClientID, CleanData)

	// 取出Client端相關資訊
	pCliet := Common_ClientInfoGet(ClientID)
	ClientIP = pCliet.ClientIP

	// 抓出桌子
	pTable := Match_TableInfoGet(pCliet.Member.TableArrayIdx)

	// 判斷如果在桌內 就離開遊戲 並通知其它桌內玩家
	ResponseTmpData, Code, BroadcastTable, BroadcastMember = Match_ExitTable(ClientID, pTable, &pCliet.Member)
	CommonLog_INFO_Printf("#Common_ClientClose 離開遊戲-結果 Code=%d, ResponseTmpData:%s", Code, ResponseTmpData)
	// 廣播
	if Code == int(ERROR_CODE_SUCCESS) {
		ResponseFinishData, Code = Message_Broadcast(NET_CMD_EXIT_GAME, BroadcastTable, BroadcastMember, ResponseTmpData)
		ret = true
		CommonLog_INFO_Printf("#Common_ClientClose 廣播的 Code=%d, ResponseFinishData=%s", Code, ResponseFinishData)
	}

	if CleanData == true {
		// 清除會員資料
		ret = Common_ClientDelete(ClientID)
	}

	CommonLog_INFO_Printf("#Common_ClientClose 結束 關閉一個client... ClientID=%d, ClientIP=%s, ret=%t", ClientID, ClientIP, ret)

	return ret
}

//=========================================================================================================================================================================
// 取得一個 client 的記憶體位置
func Common_ClientInfoGet(ClientID int) *ClientConn {

	CommonLog_INFO_Printf("#Common_ClientInfoGet 取得一個 client 的記憶體位置... ClientID=%d", ClientID)

	// 取出記憶體位置
	pClient := &ClientList[ClientID]

	//CommonLog_INFO_Printf("#Common_ClientInfoGet 取得一個 client 的記憶體位置... ClientID=%d pClient=%p", ClientID, pClient)
	return pClient
}

//=========================================================================================================================================================================
// 找出空閒的 map idx
func Common_IdleClientIDGet() int {

	var bFind bool = false
	var ClientID int = -1

	for i := 0; i < CLIENT_MAX; i++ {

		// 取出記憶體位置
		var Client = &ClientList[i]

		if Client.IsUse == true {
			continue
		}

		CommonLog_INFO_Printf("i=%d 找到空位", i)
		ClientID = i
		Client.IsUse = true
		Client.Sn = 0
		bFind = true
		break
	}

	if bFind == false {
		CommonLog_WARNING_Printf("#Warning!!!! 都沒找到空位 CLIENT_MAX=%d, ClientCount=%d", CLIENT_MAX, ClientCount)
	}

	return ClientID
}

//=========================================================================================================================================================================
// 對client發送訊息
func Common_ClientSendMessage(ClientID int, SendDataMsg string) bool {

	var err error
	var ret bool = true

	CommonLog_INFO_Printf("#Common_ClientSendMessage(對client發送訊息) ClientID=%d, SendDataMsg=%s", ClientID, SendDataMsg)

	pCliet := Common_ClientInfoGet(ClientID)

	CommonLog_INFO_Printf("#Common_ClientSendMessage(對client發送訊息) ClientID=%d, ClientIP=%s, Account=%s, User_ID=%d, SendDataMsg=%s",
		ClientID, pCliet.ClientIP, pCliet.Member.Account, pCliet.Member.User_ID, SendDataMsg)

	if err = websocket.Message.Send(pCliet.websocket, SendDataMsg); err != nil {
		CommonLog_WARNING_Println("Can't Send... ClientID=%d, ClientIP=%s, Account=%s, User_ID=%d, SendDataMsg=%s",
			ClientID, pCliet.ClientIP, pCliet.Member.Account, pCliet.Member.User_ID, SendDataMsg)
		ret = false
	}

	return ret
}

//=========================================================================================================================================================================
// 分析cmd 並分配和處理
//func Common_Analysis(PacketCmd CommonPacketCmd) CommonResponseInfo {
func Common_Analysis(ClientID int, ClientIP string, receiveMsg string) CommonResponseInfo {

	var DecodeData string = ""         // 解密後的資料
	var ResponseData string = "unknow" // 回應的資料格式
	var Response CommonResponseInfo    // 共用回應結構
	var Code int = 0

	pClient := Common_ClientInfoGet(ClientID)

	if DEBUG_TEST_MUTEX_0 == true {

		CommonLog_INFO_Printf("#Common_Analysis DEBUG_TEST_MUTEX_0 資料同步鎖開啟 ClientID=%d, ClientIP=%s, receiveMsg=%s", ClientID, ClientIP, receiveMsg)
		DataMutex.Lock()

		defer DataMutex.Unlock()
	}

	CommonLog_INFO_Printf("#Common_Analysis 開始================================ ClientID=%d, ClientIP=%s, 收到訊息receiveMsg:%s", ClientID, ClientIP, receiveMsg)

	//------------------------------------------------
	// 收到的資料 json轉換
	receiveMsgByte := []byte(receiveMsg)
	//var PacketCmd = CommonPacketCmd{} // 共用的接收物件
	PacketCmd := CommonPacketCmd{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &PacketCmd)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 receiveMsg=%s", receiveMsg)

		Response.Code = -1
		Response.Message = "json解析錯誤"
		return Response
	}

	CommonLog_INFO_Printf("Cmd(命令種類):%s", PacketCmd.Cmd)
	CommonLog_INFO_Printf("sys(是否是 system cmd):%s", PacketCmd.Sys)
	CommonLog_INFO_Printf("sn:%d", PacketCmd.Sn)
	CommonLog_INFO_Printf("IsEncode(是否加密):%t", PacketCmd.IsEncode)
	CommonLog_INFO_Printf("data(封包資料):%s", PacketCmd.Data)

	if PacketCmd.IsEncode {
		CommonLog_INFO_Printf("缺一個加解密的功能,日後補")
		DecodeData = "缺解密資料,日後補"

	} else {
		DecodeData = PacketCmd.Data
	}

	CommonLog_INFO_Printf("DecodeData:%s", DecodeData)

	// cmd 拷貝一份 回傳時候可以使用
	Response.Ret = PacketCmd.Cmd
	Response.Sys = PacketCmd.Sys

	// 預設的回應訊息
	Response.Code = 0
	Response.Message = "正確執行"

	// 檢查sn
	//pClient := Common_ClientInfoGet(ClientID)
	pClient.Sn++
	Response.Sn = pClient.Sn

	// 訊息分配
	if PacketCmd.Sys == "system" {

		//系統的cmd
		ResponseData, Code = Common_DispatchSystem(ClientID, PacketCmd.Cmd, DecodeData)
	} else if PacketCmd.Sys == "game" {
		//遊戲的cmd
		ResponseData, Code = Common_DispatchGame(ClientID, PacketCmd.Cmd, DecodeData)
	} else {

		Code = int(ERROR_CODE_NO_FIND_CMD)
		CommonLog_INFO_Printf("#Common_Analysis ClientID=%d, ClientIP=%s, 錯誤的sys=%s", ClientID, ClientIP, PacketCmd.Sys)
	}

	CommonLog_INFO_Printf("#Common_Analysis ClientID=%d, ClientIP=%s, Cmd=%s, Code:%d", ClientID, ClientIP, PacketCmd.Cmd, Code)

	if Code != 0 {
		// 轉換錯誤代碼
		Response.Code = ErrorCode[Code].Code
		Response.Message = ErrorCode[Code].Message
		ResponseData = ""
		CommonLog_WARNING_Printf("#Common_Analysis 發生錯誤 ClientID=%d, ClientIP=%s, Code=%d, Message=%s", ClientID, ClientIP, Response.Code, Response.Message)

		// Code, Response 如果有連續異常N次 直接斷線 並列入黑名單
	}
	Response.Data = ResponseData
	Response.ClientID = ClientID

	CommonLog_INFO_Printf("#Common_Analysis 結束================================ ClientID=%d, ClientIP=%s, SN=%d, ResponseData:%s", ClientID, ClientIP, pClient.Sn, ResponseData)

	return Response
}

//============================================================================================================
// 讀取Config文字檔
func Config_ReadFile(FilePath string) Config {

	CommonLog_INFO_Printf("#FishScript_ReadFile(讀取腳本文字檔) FilePath=%s", FilePath)

	config := Config{} // 用來接的物件

	CommonLog_INFO_Printf("#FishScript_ReadFile FilePath=%s", FilePath)

	data, error := ioutil.ReadFile(FilePath)
	if error != nil {
		CommonLog_WARNING_Println("錯誤的讀檔 error=", error, "FilePath=", FilePath)
		return config
	}

	var fileStr string = string(data)
	//CommonLog_INFO_Printf("讀取到的文字 fileStr=%s", fileStr)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(fileStr)

	err := json.Unmarshal(receiveMsgByte, &config)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", fileStr)
	} else {

		// 簡單列印一下
		CommonLog_INFO_Printf("讀取Config文字檔 FilePath=%s, Server_Port=%d, Mysql_IP=%s, Mysql_Account=%s, Mysql_Password=%s, Mysql_SetMaxOpenConns=%d, Mysql_SetMaxIdleConns=%d",
			FilePath, config.Server_Port, config.Mysql_IP, config.Mysql_Account, config.Mysql_Password, config.Mysql_SetMaxOpenConns, config.Mysql_SetMaxIdleConns)
	}

	return config
}

//=========================================================================================================================================================================
// 系統 init
func Common_Init() {
	CommonLog_INFO_Printf("#Common_Init VERSION(Server版本)=%s, CLIENT_MAX(連線數)=%d, TABLEINFO_MAX(配桌桌數)=%d", VERSION, CLIENT_MAX, TABLEINFO_MAX)

	// 讀取大廳資訊
	Common_LoadData()

	// 執行緒的鎖
	DataMutex = new(sync.RWMutex)
}

//=========================================================================================================================================================================
// 系統 資訊讀取
func Common_LoadData() {
	CommonLog_INFO_Printf("#Common_LoadData")

	// mysql 初始化
	Mysql_Init()

	// 讀取所有大廳資訊
	//Mysql_CommonLobbyInfoGetAll()

	// 讀取所有遊戲資訊
	//Mysql_CommonGameInfoGetAll()

	// 取得所有桌子資訊 ( 等到 桌子可以跟 ws 綁在一起 才可以打開, 不然DB內幽靈桌的問題需要討論, 但是這邊可以做的是 把錢從房間拿回帳號 )
	//Mysql_CommonTableInfoGetAll()

	// 取得所有位子資訊 ( 等到 桌子可以跟 ws 綁在一起 才可以打開, 不然DB內幽靈桌的問題需要討論 )
	//Mysql_CommonSeatInfoGetAll()

	// 依照桌子, 去把錢從房間拿回帳號, 再刪除桌子資訊

	// 顯示桌子資訊 (原則上剛開機都是空的)
	//Mysql_CommonTableInfoShow()

	//Mysql_CommonTableInfo_IdGet()

	// 桌子的Process ( 產魚腳本運行 or dosomething )
	go Table_Process()

	//obj := CustomerInfo{}
	//obj.CustomerName = "測試顧客"
	obj := Task{}
	obj.TaskName = "工作清單01"
	DataMsgByte, err := json.Marshal(obj)
	if err != nil {
		CommonLog_WARNING_Println("json err:", err)
	}

	DataMsg := string(DataMsgByte)
	CommonLog_INFO_Printf("DataMsg=%s", DataMsg)

	obj0 := CommonPacketCmd{}
	obj0.Cmd = "task_insert"
	obj0.Sys = "system"
	obj0.Data = DataMsg
	DataMsgByte0, err0 := json.Marshal(obj0)
	if err0 != nil {
		CommonLog_WARNING_Println("json err:", err0)
	}
	DataMsg0 := string(DataMsgByte0)
	CommonLog_INFO_Printf("組合出來的字串=%s", DataMsg0)
	CommonLog_INFO_Printf("===========")

}

//=========================================================================================================================================================================
//=========================================================================================================================================================================
//=========================================================================================================================================================================
// 系統的CMD (此函式放到最下面)
func Common_DispatchSystem(ClientID int, Cmd string, DecodeData string) (string, int) {

	var Code int = int(ERROR_CODE_SUCCESS) // 回應值
	var ResponseData string = "unknow"     // 回應的資料格式

	CommonLog_INFO_Printf("#Common_DispatchSystem(系統的CMD) ClientID=%d, Cmd=%s, DecodeData=%s", ClientID, Cmd, DecodeData)

	// 分析 cmd 並且分派和處理
	switch Cmd {

	case NET_CMD_ACCOUNT_CREATE:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_ACCOUNT_CREAT")
		}
	case NET_CMD_LOGIN:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_LOGIN")

			// 取得登入資訊
			ResponseData, Code = Common_Login(ClientID, DecodeData)
			CommonLog_INFO_Printf("#取得登入資訊 Code:%d, ResponseData:%s", Code, ResponseData)

		}
	case NET_CMD_LOGOUT:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_LOGOUT")

			// 取得登出資訊
			ResponseData, Code = Common_Logout(ClientID, DecodeData)
			CommonLog_INFO_Printf("#取得登出資訊 Code:%d, ResponseData:%s", Code, ResponseData)

			// 刪除 array client info
			//Common_ClientDelete(ClientID)
			Common_ClientClose(ClientID, false) // 避免卡桌, 就都透過這個函式處理吧
		}

	case NET_CMD_MEMBER_INSERT:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_MEMBER_INSERT")

			// 新增會員
			ResponseData, Code = Common_MemberInsert(ClientID, DecodeData)
			CommonLog_INFO_Printf("#新增會員 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NET_CMD_MEMBER_UPDATE:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_MEMBER_UPDATE")

			// 修改會員資料
			ResponseData, Code = Common_MemberUpdate(ClientID, DecodeData)
			CommonLog_INFO_Printf("#修改會員資料 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NET_CMD_MEMBER_DELETE:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_MEMBER_DELETE")

			// 刪除會員
			ResponseData, Code = Common_MemberDelete(ClientID, DecodeData)
			CommonLog_INFO_Printf("#刪除會員資料 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NET_CMD_MEMBER_LIST_GET:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_MEMBER_LIST_GET")

			// 會員清單取得
			ResponseData, Code = Common_MemberListGet(ClientID, DecodeData)
			CommonLog_INFO_Printf("#會員清單取得 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NEW_CMD_CUSTOMER_INSERT:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NEW_CMD_CUSTOMER_INSERT")

			// 新增顧客
			ResponseData, Code = Common_CustomerInsert(ClientID, DecodeData)
			CommonLog_INFO_Printf("#新增顧客 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NEW_CMD_CUSTOMER_UPDATE:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NEW_CMD_CUSTOMER_UPDATE")

			// 更新顧客
			ResponseData, Code = Common_CustomerUpdate(ClientID, DecodeData)
			CommonLog_INFO_Printf("#更新顧客 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NET_CMD_CUSTOMER_DELETE:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_CUSTOMER_DELETE")

			// 刪除顧客
			ResponseData, Code = Common_CustomerDelete(ClientID, DecodeData)
			CommonLog_INFO_Printf("#刪除顧客 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NET_CMD_CUSTOMER_LIST_GET:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_CUSTOMER_LIST_GET")

			// 顧客清單取得
			ResponseData, Code = Common_CustomerListGet(ClientID, DecodeData)
			CommonLog_INFO_Printf("#顧客清單取得 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NEW_CMD_TASK_INSERT:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NEW_CMD_TASK_INSERT")

			// 新增工作
			ResponseData, Code = Common_TaskInsert(ClientID, DecodeData)
			CommonLog_INFO_Printf("#新增工作 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NEW_CMD_TASK_UPDATE:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NEW_CMD_TASK_UPDATE")

			// 更新工作
			ResponseData, Code = Common_TaskUpdate(ClientID, DecodeData)
			CommonLog_INFO_Printf("#更新工作 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NET_CMD_TASK_DELETE:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_TASK_DELETE")

			// 刪除工作
			ResponseData, Code = Common_TaskDelete(ClientID, DecodeData)
			CommonLog_INFO_Printf("#刪除工作 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NET_CMD_TASK_LIST_GET:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_TASK_LIST_GET")

			// 工作清單取得
			ResponseData, Code = Common_TaskListGet(ClientID, DecodeData)
			CommonLog_INFO_Printf("工作清單取得 Code:%d, ResponseData:%s", Code, ResponseData)
		}

	case NET_CMD_LOBBYINFO_GET:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_LOBBYINFO_GET")

			// 取得大廳資訊
			ResponseData, Code = Common_LobbyInfoGet(ClientID, DecodeData)
			CommonLog_INFO_Printf("#取得大廳資訊 ResponseData:%s", ResponseData)
		}

	default:
		CommonLog_WARNING_Printf("warning 未處理的cmd=%s", Cmd)
		Code = int(ERROR_CODE_NO_FIND_CMD)
		ResponseData = ErrorCode[ERROR_CODE_NO_FIND_CMD].Message
		//break
	}

	return ResponseData, Code
}
