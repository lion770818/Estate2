package One1CloudLib

import (
	"fmt"

	//"log"
	"time"

	"encoding/json"
	//"golang.org/x/net/websocket"
)

//=========================================================================================================================================================================
// 回傳目前時間 格式 2017-07-21 14:28:39
func Common_NowTimeGet() string {
	now := time.Now()
	var timeStr string = fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	return timeStr
}

//=========================================================================================================================================================================
// 回傳目前時間物件
func Common_NowTimeObjGet() time.Time {

	var now time.Time
	now = time.Now()
	return now
}

//=========================================================================================================================================================================
// 取得 登入資訊
func Common_Login(ClientID int, DecodeData string) (string, int) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值

	var pCinent = Common_ClientInfoGet(ClientID)
	CommonLog_INFO_Printf("#收到封包 Common_Login ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	Auth := PacketCmd_AuthInfo{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &Auth)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		//panic(err)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code
	} else {
		CommonLog_INFO_Printf("#收到封包 Common_Login Account=%s, Password=%s", Auth.Account, Auth.Password)

		// 讀取會員資料
		ret, member := Mysql_CommonMemberInfoGet(Auth.Account, Auth.Password)

		if ret == true {
			CommonLog_INFO_Printf("# 找到會員資料 PlatformID=%d, User_ID=%d, Account=%s, NickName=%s", member.PlatformID, member.User_ID, member.Account, member.NickName)
			CommonLog_INFO_Printf("#ret=%t", ret)

			// 設定拿到的ClientID
			member.ClientID = ClientID

			// 設定 Status = 1 登入狀態
			member.Status = 1
			ret2 := Mysql_CommonMemberInfo_UpdateStatus(member.Status, Auth)
			if ret2 == true {
				CommonLog_INFO_Printf("# 修改會員 Status 資料 PlatformID=%d, User_ID=%d, Account=%s, NickName=%s, Status=%d", member.PlatformID, member.User_ID, member.Account, member.NickName, member.Status)
			}

			// 儲存在 client array list 中
			if pCinent.IsUse == true {
				CommonLog_INFO_Printf("ClientID=%d, ClientIP=%s", ClientID, pCinent.ClientIP)
				pCinent.Member = member
				pCinent.Member.ClientID = ClientID

				Code = int(ERROR_CODE_SUCCESS)
			} else {
				CommonLog_WARNING_Printf("Warning!!! 找不到 ClientID=%d", ClientID)

				// 錯誤流程待補
				Code = int(ERROR_CODE_DATA_UPDATE_FAIL)
			}

			// 回傳給 client
			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(member)
			if err != nil {
				CommonLog_WARNING_Println("json err:", err)
			}

			DataMsg = string(DataMsgByte)
		} else {
			Code = int(ERROR_CODE_NO_FIND_ACCOUNT)
			DataMsg = fmt.Sprintf("查無此帳號 Account=%s", Auth.Account)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, DataMsg=%s", Code, DataMsg)
	return DataMsg, Code
}

//=========================================================================================================================================================================
// 取得 登出資訊
func Common_Logout(ClientID int, DecodeData string) (string, int) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值

	var pCinent = Common_ClientInfoGet(ClientID)
	CommonLog_INFO_Printf("#收到封包 Common_Logout ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	Auth := PacketCmd_AuthInfo{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &Auth)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code
	} else {
		CommonLog_INFO_Printf("#收到封包 Common_Logout Account=%s, Password=%s", Auth.Account, Auth.Password)

		// 讀取會員資料
		ret, member := Mysql_CommonMemberInfoGet(Auth.Account, Auth.Password)

		if ret == true {
			CommonLog_INFO_Printf("# 找到會員資料 PlatformID=%d, User_ID=%d, Account=%s, NickName=%s", member.PlatformID, member.User_ID, member.Account, member.NickName)
			CommonLog_INFO_Printf("#ret=%d", ret)

			// 設定 Status = 0 登入狀態
			member.Status = 0
			// 修改會員登入狀態
			ret2 := Mysql_CommonMemberInfo_UpdateStatus(member.Status, Auth)
			if ret2 == true {
				CommonLog_INFO_Printf("#修改會員 Status 資料 PlatformID=%d, User_ID=%d, Account=%s, NickName=%s, Status=%d", member.PlatformID, member.User_ID, member.Account, member.NickName, member.Status)
			} else {

				CommonLog_WARNING_Printf("Warning!!! 找不到 ClientID=%d", ClientID)

				// 錯誤流程待補
				Code = int(ERROR_CODE_DATA_UPDATE_FAIL)
			}

			// 清除client記憶體資料
			//pCinent.IsUse = false
			//pCinent.ClientIP = ""
			//pCinent.ClientID = -1
			member := MemberInfo{}
			pCinent.Member = member

			// 回傳給 client
			// 物件轉成json字串
			//DataMsgByte, err := json.Marshal(MemberInfo)
			DataMsgByte, err := json.Marshal("")
			if err != nil {
				CommonLog_WARNING_Printf("json err:", err)
				Code = int(ERROR_CODE_DATA_UPDATE_FAIL)
			}

			DataMsg = string(DataMsgByte)
		}

	}

	CommonLog_INFO_Printf("#Code=%d, DataMsg=%s", Code, DataMsg)
	return DataMsg, Code
}

//=========================================================================================================================================================================
// 新增會員
func Common_MemberInsert(ClientID int, DecodeData string) (string, int) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值

	CommonLog_INFO_Printf("#收到封包 Common_MemberInsert ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	NewMember := MemberInfo{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &NewMember)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code
	} else {
		CommonLog_INFO_Printf("#收到封包 Common_MemberInsert Account=%s, Password=%s, NickName=%s, IdentityNumber=%s",
			NewMember.Account, NewMember.Password, NewMember.NickName, NewMember.IdentityNumber)

		// 檢查玩家是否有登入
		pClient := Common_ClientInfoGet(ClientID)
		if pClient.IsUse == true && pClient.Member.Status == 1 {

			// 讀取玩家的會員資料
			pMember := &pClient.Member // 取址

			if pMember.Vip_rank < 6 {
				Code = int(ERROR_CODE_ERROR_PERMISSION_DENIED)
				CommonLog_WARNING_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code
			}

			// 新增會員
			NewMember.CreateTime = Common_NowTimeGet()
			NewMember.LoginTime = Common_NowTimeGet()
			NewMember.UpdateTime = Common_NowTimeGet()
			ret := Mysql_CommonMemberInfo_Insert(NewMember)
			if ret == false {
				Code = int(ERROR_CODE_ERROR_CREATE_MEMBER)
				CommonLog_WARNING_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code
			}

			// 回傳給 client
			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(NewMember)
			//DataMsgByte, err := json.Marshal("")
			if err != nil {
				CommonLog_WARNING_Printf("json err:", err)
				Code = int(ERROR_CODE_DATA_UPDATE_FAIL)
			}

			DataMsg = string(DataMsgByte)

		} else {
			Code = int(ERROR_CODE_NO_LOGIN)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, DataMsg=%s", Code, DataMsg)
	return DataMsg, Code
}

//=========================================================================================================================================================================
// 更新會員資料
func Common_MemberUpdate(ClientID int, DecodeData string) (string, int) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值

	CommonLog_INFO_Printf("#收到封包 Common_MemberUpdate(更新會員資料) ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	UpdateMember := MemberInfo{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &UpdateMember)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code
	} else {
		CommonLog_INFO_Printf("#收到封包 Common_MemberUpdate Account=%s, Password=%s, NickName=%s, IdentityNumber=%s",
			UpdateMember.Account, UpdateMember.Password, UpdateMember.NickName, UpdateMember.IdentityNumber)

		// 檢查玩家是否有登入
		pClient := Common_ClientInfoGet(ClientID)
		if pClient.IsUse == true && pClient.Member.Status == 1 {

			// 讀取玩家的會員資料
			pMember := &pClient.Member // 取址

			if pMember.Vip_rank < 6 {
				Code = int(ERROR_CODE_ERROR_PERMISSION_DENIED)
				CommonLog_WARNING_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code
			}

			// 更新會員
			UpdateMember.UpdateTime = Common_NowTimeGet()
			ret := Mysql_CommonMemberInfo_UpdateData(UpdateMember)
			if ret == false {
				Code = int(ERROR_CODE_ERROR_UPDATE_MEMBER)
				CommonLog_WARNING_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code
			}

			ResponseData := fmt.Sprintf("更新會員成功 Account=%s, Password=%s, NickName=%s, IdentityNumber=%s", UpdateMember.Account, UpdateMember.Password, UpdateMember.NickName, UpdateMember.IdentityNumber)

			// 回傳給 client
			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(ResponseData)
			//DataMsgByte, err := json.Marshal("")
			if err != nil {
				CommonLog_WARNING_Printf("json err:", err)
				Code = int(ERROR_CODE_DATA_UPDATE_FAIL)
			}

			DataMsg = string(DataMsgByte)

		} else {
			Code = int(ERROR_CODE_NO_LOGIN)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, DataMsg=%s", Code, DataMsg)
	return DataMsg, Code
}

//=========================================================================================================================================================================
// 刪除會員資料
func Common_MemberDelete(ClientID int, DecodeData string) (string, int) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值

	CommonLog_INFO_Printf("#收到封包 Common_MemberDelete(更新會員資料) ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	DeleteMember := MemberInfo{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &DeleteMember)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code
	} else {
		CommonLog_INFO_Printf("#收到封包 Common_MemberDelete Account=%s, Password=%s, NickName=%s, IdentityNumber=%s",
			DeleteMember.Account, DeleteMember.Password, DeleteMember.NickName, DeleteMember.IdentityNumber)

		// 檢查玩家是否有登入
		pClient := Common_ClientInfoGet(ClientID)
		if pClient.IsUse == true && pClient.Member.Status == 1 {

			// 讀取玩家的會員資料
			pMember := &pClient.Member // 取址

			if pMember.Vip_rank < 6 {
				Code = int(ERROR_CODE_ERROR_PERMISSION_DENIED)
				CommonLog_WARNING_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code
			}

			// 刪除會員
			ret := Mysql_CommonMemberInfo_Delete(DeleteMember)
			if ret == false {
				Code = int(ERROR_CODE_NO_FIND_ACCOUNT)
				CommonLog_WARNING_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code
			}

			ResponseData := fmt.Sprintf("刪除會員成功 Account=%s, Password=%s, NickName=%s, IdentityNumber=%s", DeleteMember.Account, DeleteMember.Password, DeleteMember.NickName, DeleteMember.IdentityNumber)

			// 回傳給 client
			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(ResponseData)
			//DataMsgByte, err := json.Marshal("")
			if err != nil {
				CommonLog_WARNING_Printf("json err:", err)
				Code = int(ERROR_CODE_DATA_UPDATE_FAIL)
			}

			DataMsg = string(DataMsgByte)

		} else {
			Code = int(ERROR_CODE_NO_LOGIN)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, DataMsg=%s", Code, DataMsg)
	return DataMsg, Code
}

//=========================================================================================================================================================================
// 取得會員列表資料
func Common_MemberListGet(ClientID int, DecodeData string) (string, int) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值

	CommonLog_INFO_Printf("#收到封包 Common_MemberListGet(取得會員列表資料) ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	Auth := PacketCmd_AuthInfo{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &Auth)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code
	} else {
		CommonLog_INFO_Printf("#收到封包 Common_MemberListGet Account=%s, Password=%s",
			Auth.Account, Auth.Password)

		// 檢查玩家是否有登入
		pClient := Common_ClientInfoGet(ClientID)
		if pClient.IsUse == true && pClient.Member.Status == 1 {

			// 讀取玩家的會員資料
			pMember := &pClient.Member // 取址

			if pMember.Vip_rank < 6 {
				Code = int(ERROR_CODE_ERROR_PERMISSION_DENIED)
				CommonLog_WARNING_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code
			}

			// 取得會員清單
			MemberList, ret := Mysql_CommonMemberListGet2(Auth)
			if ret == false {
				Code = int(ERROR_CODE_NO_FIND_ACCOUNT)
				CommonLog_WARNING_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code
			}

			CommonLog_INFO_Printf("取得會員清單成功 Account=%s, Password=%s, MemberList Count=%d",
				Auth.Account, Auth.Password, MemberList.Data_Count)

			// 回傳給 client
			// 物件轉成json字串
			if ret == true {
				
				// 物件轉成json字串
				DataMsgByte, err := json.Marshal(MemberList)
				if err != nil {
					CommonLog_WARNING_Printf("json err:", err)
				}

				DataMsg = string(DataMsgByte)
			}

		} else {
			Code = int(ERROR_CODE_NO_LOGIN)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, DataMsg=%s", Code, DataMsg)
	return DataMsg, Code
}

//=========================================================================================================================================================================
// 分析cmd 並分配和處理
// 取得 大廳資訊
func Common_LobbyInfoGet(ClientID int, DecodeData string) (string, int) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值

	CommonLog_INFO_Printf("#收到封包 Common_LobbyInfoGet ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	LobbyinfoGet := PacketCmd_LobbyinfoGet{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &LobbyinfoGet)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code
	} else {
		CommonLog_INFO_Printf("#收到封包 Common_LobbyInfoGet PlatformID=%d", LobbyinfoGet.PlatformID)

		// 取得大廳資訊
		LobbyInfo, ret := Mysql_CommonLobbyInfoGet2(LobbyinfoGet.PlatformID)

		if ret == true {
			CommonLog_INFO_Printf("#LobbyName=%s", LobbyInfo[0].LobbyName)
			CommonLog_INFO_Printf("#ret=%d", ret)

			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(LobbyInfo)
			if err != nil {
				CommonLog_WARNING_Printf("json err:", err)
			}

			DataMsg = string(DataMsgByte)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, DataMsg=%s", Code, DataMsg)
	return DataMsg, Code
}
