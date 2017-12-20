package One1CloudLib

//=========================================================================================================================================================================
// 留言板
type MessageBoard struct {
	MessageBoardID   int    `json:"message_board_id"`   //留言板的ID
	MessageBoardName string `json:"message_board_name"` //留言板的名稱
	User_ID          int64  `json:"user_id"`            //玩家帳號編號
	NickName         string `json:"nickname"`           //玩家暱稱
	HomeName         string `json:"home_name"`          //房屋名稱

	CreateTime string `json:"createtime"` //留言板建立時間
	UpdateTime string `json:"updatetime"` //留言板更新時間

	PhoneNumber string `json:"phone_number"` //玩家電話號碼
	Gender      string `json:"gender"`       //性別 M:男 F:女

	Rent int `json:"rent"` //租金

	IsPet               int `json:"is_pet"`                //是否養寵物
	IsSmoke             int `json:"is_smoke"`              //是否抽煙
	IsHouseholdRegister int `json:"is_household_register"` //是否入戶籍
	IsTax               int `json:"is_tax"`                //租金
	Isfaith             int `json:"is_faith"`              //租金

	Memo string `json:"memo"` //備註
}

//=========================================================================================================================================================================
// 留言板結構列表 ( data 內的資料 )
type ResponseInfo_MessageBoardList struct {
	Data_Count        int                  `json:"data_count"`         // 資料筆數
	MessageBoard_List map[int]MessageBoard `json:"message_board_list"` // 留言板
}
