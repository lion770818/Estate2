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

const (
	MEMBER_NEW_EMPLOYEE               Base = iota // 0 新進員工
	MEMBER_NOEMAL_EMPLOYEE                        // 1 一般員工
	MEMBER_SUPERVISOR                             // 2 組長
	MEMBER_ACCOUNT_MANAGER                        // 3 主任
	MEMBER_MANAGER                                // 4 經理
	MEMBER_HUMAN_RESOURCES_SUPERVISOR             // 5 人資
	MEMBER_ACCOUNTANT                             // 6 會計
	MEMBER_MAIN_SUPERVISOR                        // 7 最高管理者
	MEMBER_PRESIDENT                              // 8 董事長

)

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
	UpdateTime string `json:"updatetime"` //帳號更新時間

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
	Vip_rank int `json:"vip_rank"` //會員等級 0:實習生 1:一般員工 2:組長 3:主任 4:經理 5:人資 6:會計 7:最高管理者 8:董事長
}

//=========================================================================================================================================================================
// 會員結構列表 ( data 內的資料 )
type ResponseInfo_MemberInfoList struct {
	Data_Count  int                `json:"data_count"`  // 資料筆數
	Member_List map[int]MemberInfo `json:"member_list"` // 會員
}

//=========================================================================================================================================================================
// 顧客結構
type CustomerInfo struct {

	//================= db =============================
	CustomerID int    `json:"customer_id"` //顧客id
	User_ID    int64  `json:"user_id"`     //玩家帳號編號(哪個員工增加的顧客)
	NickName   string `json:"nickname"`    //玩家名稱(員工的名稱)

	CreateTime string `json:"createtime"` //資料建立時間
	UpdateTime string `json:"updatetime"` //資料更新時間

	CustomerName           string `json:"customer_name"`           //顧客名稱
	CustomerAge            int    `json:"customer_age"`            //顧客年紀
	CustomerGender         string `json:"customer_gender"`         //顧客性別 M:男 F:女
	CustomerIdentityNumber string `json:"customer_identityNumber"` //身分證字號

	CustomerPhoneNumber string `json:"customer_phone_number"` //顧客電話號碼
	CustomerAddress     string `json:"customer_address"`      //顧客地址

	CustomerHomeID      int     `json:"customer_home_id"`      //顧客房屋物件id編號
	CustomerHomeAge     int     `json:"customer_home_age"`     //顧客房屋屋齡
	CustomerHomeFootage float32 `json:"customer_home_footage"` //顧客房屋坪數
	CustomerHomePrice   int     `json:"customer_home_price"`   //顧客房屋價格

	Vip_rank int `json:"vip_rank"` //顧客等級 0:一般 1:熟客 2:VIP

}

//=========================================================================================================================================================================
// 顧客結構列表 ( data 內的資料 )
type ResponseInfo_CustomerInfoList struct {
	Data_Count    int                  `json:"data_count"`    // 資料筆數
	Customer_List map[int]CustomerInfo `json:"customer_list"` // 顧客
}
