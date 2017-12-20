package One1CloudLib

// 會員結構
type HomeInfo struct {

	//================= db =============================
	HomeID   int    `json:"home_id"`  //房屋物件id編號
	User_ID  int64  `json:"user_id"`  //玩家帳號編號(哪個員工增加的房屋)
	NickName string `json:"nickname"` //玩家名稱(員工的名稱)

	CreateTime string `json:"createtime"` //資料建立時間
	UpdateTime string `json:"updatetime"` //資料更新時間

	HomeName    string  `json:"home_name"`    //房屋名稱
	HomeArea    int     `json:"home_area"`    //房屋名稱
	HomeAddress string  `json:"home_address"` //地址
	HomeAge     int     `json:"home_age"`     //房屋屋齡
	HomeFootage float32 `json:"home_footage"` //房屋坪數
	HomePrice   int     `json:"home_price"`   //房屋價格

	Vip_rank int `json:"vip_rank"` //房屋等級 0:雅房 1:套房 2:兩房一廳 3:三房兩廳 4:工廠 5:辦公室 6:透天厝 7:豪宅

	Memo string `json:"memo"` //房屋備註
}

//=========================================================================================================================================================================
// 房屋結構列表 ( data 內的資料 )
type ResponseInfo_HomeInfoList struct {
	Data_Count int              `json:"data_count"` // 資料筆數
	Home_List  map[int]HomeInfo `json:"home_list"`  // 房屋
}
