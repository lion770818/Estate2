package One1CloudLib

//"fmt"
//"log"
//"time"
/*
// 會員結構 ( slot 使用 )
type MemberInfoSlots struct {
	UID        int
	Account    string
	Password   string
	Game_Money int
	Bonus      int
}
*/

// 會員結構
type MemberInfo struct {
	//================= 記憶體的資料 =============================
	ClientID      int `json:"client_id"` //Client端的ID
	TableArrayIdx int `json:"table_array_idx"`

	//================= db =============================
	User_ID    int64 `json:"user_id"`     //玩家帳號編號
	PlatformID int   `json:"platform_id"` //第三方平台編號
	DeviceID   int   `json:"device_id"`   //玩家裝置編號

	IP         string `json:"ip"`         //玩家IP Address
	MacAddress string `json:"macaddress"` //玩家網卡 Address

	CreateTime string `json:"createtime"` //帳號建立時間
	LoginTime  string `json:"logintime"`  //帳號登入時間
	UpdateTime string `json:"updatetime"` //帳號登入時間

	Account  string `json:"account"`  //帳號
	Password string `json:"password"` //密碼
	NickName string `json:"nickname"` //暱稱

	IdentityNumber string `json:"identityNumber"` //身分證字號
	Address        string `json:"address"`        //玩家地址
	PhoneNumber    string `json:"phone_number"`   //玩家電話號碼

	Balance int64 `json:"balance"` //玩家的錢
	Bonus   int64 `json:"bonus"`   //玩家的紅利
	Salary  int64 `json:"salary"`  //玩家的月薪(每月)

	Status   int `json:"status"`   //玩家狀態 0: 登出 1:登入 2:斷線中 3:斷線連回中
	Vip_rank int `json:"vip_rank"` //會員等級
}

//=========================================================================================================================================================================
// 會員結構列表 ( data 內的資料 )
type ResponseInfo_MemberInfoList struct {
	Data_Count int					`json:"data_count"`		// 資料筆數
	Member_List map[int]MemberInfo 	`json:"member_list"`  	// 會員
	//Member_List []MemberInfo    `json:"member_list"`  	// 會員
}