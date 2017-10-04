package One1CloudLib

//"fmt"
//"log"
//"time"

//=========================================================================================================================================================================
// 共用封包命令結構
type CommonPacketCmd struct {
	Cmd      string `json:"cmd"`      // 命令種類
	Sys      string `json:"sys"`      // 是否是 system cmd  ( sys:"game" 遊戲封包 sys:"system" 系統封包)
	Sn       int64  `json:"sn"`       // 封包序號
	IsEncode bool   `json:"isEncode"` // 是否加密
	Data     string `json:"data"`     // 封包資料
}

//=========================================================================================================================================================================
// 共用的回應結構
type CommonResponseInfo struct {
	Code     int    `json:"code"`      // 回應的代碼
	Message  string `json:"message"`   // 回應的訊息
	Sys      string `json:"sys"`       // 是否是 system cmd  ( sys:"game" 遊戲封包 sys:"system" 系統封包)
	ClientID int    `json:"client_id"` // 回應的Client, 是 ClientList 的idx值, 方便快速找到client
	Sn       int64  `json:"sn"`        // 封包序號

	Ret  string `json:"ret"`  // 回應的命令種類
	Data string `json:"data"` // 回應的資料
	//PacketCmd CommonPacketCmd 		// 共用的cmd
}

//=========================================================================================================================================================================
// 共用的回應結構
type CommonResponseInfo_Forward struct {
	Forward string `json:"fw"` // 廣播給其他人
}

//=========================================================================================================================================================================
// 共用的回應結構
type CommonResponseInfo_Reply struct {
	Reply string `json:"re"` // 回應給該玩家
}

//=========================================================================================================================================================================
// 驗證結構 ( data 內的資料 )
type PacketCmd_AuthInfo struct {
	PlatformID int `json:"platform_id"` //平台編號
	//GameID     int //遊戲編號

	Account  string `json:"account"`  //帳號
	Password string `json:"password"` //密碼
}

//=========================================================================================================================================================================
// 取得大廳封包 ( data 內的資料 )
type PacketCmd_LobbyinfoGet struct {
	PlatformID int `json:"platform_id"` //平台編號
	//GameID int						//遊戲編號
}

//=========================================================================================================================================================================
// 進入遊戲 ( data 內的資料 )
type PacketCmd_Enter_Game struct {
	PlatformID int `json:"platform_id"` //平台編號
	LobbyID    int `json:"lobby_id"`    //大廳編號
	GameID     int `json:"game_id"`     //遊戲編號

	SN          int    `json:"sn"`          // 魚機的SN
	udid        string `json:"udid"`        // 裝置識別碼  ( 他們的帳號 ?)
	User_ID     int64  `json:"user_id"`     // 加入玩家的 userid
	channel     string `json:"channel"`     // 裝置平台       Android IOS、、PC
	publish_ver string `json:"publish_ver"` // Client目前版本 x.y.z  e.g. 1.0.1
	refresh     string `json:"refresh"`     // 重新取得資訊 0

	Balance_ci int64 `json:"balance_ci"` // 玩家分數_投 (玩家帶進房間的錢)
}

//=========================================================================================================================================================================
// 加入遊戲 ( data 內的資料 )
type PacketCmd_Join_Game struct {
	TableID       string `json:"table_id"` // 桌號
	TableArrayIdx int    `json:"table_array_idx"`       // TableInfoList 的索引陣列 方便快速搜尋資料 ( 感覺綁在 member上可以省點時間 )

	User_ID    int64   `json:"user_id"`    // 加入玩家的 userid
	Balance_ci int64   `json:"balance_ci"` // 玩家想帶進來的錢
}

//=========================================================================================================================================================================
// 離開遊戲 ( data 內的資料 )
type PacketCmd_Exit_Game struct {
	TableID string `json:"table_id"` // 桌號
	User_ID int64  `json:"user_id"`  // 離開玩家的 userid
	Seat_ID int8   `json:"seat_id"`  // 座位
}
