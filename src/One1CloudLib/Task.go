package One1CloudLib

//=========================================================================================================================================================================
// 工作清單
type Task struct {

	TaskID      int `json:"task_id"` //工作的ID
	User_ID    int64 `json:"user_id"`     //玩家帳號編號
	NickName   string `json:"nickname"` 	  //玩家暱稱

	CreateTime string `json:"createtime"` //工作建立時間
	UpdateTime string `json:"updatetime"` //工作更新時間

	TaskName  string `json:"task_name"`  //工作名稱
	TaskDescribe string `json:"task_describe"` //工作描述
	
	Memo    string `json:"menmo"`   //工作備註
}

//=========================================================================================================================================================================
// 工作清單結構列表 ( data 內的資料 )
type ResponseInfo_TaskList struct {
	Data_Count  int             `json:"data_count"`  // 資料筆數
	Task_List map[int]Task 		`json:"task_list"` 	 // 工作
}

