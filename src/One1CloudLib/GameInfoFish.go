package One1CloudLib

import (
	"encoding/json"
	//"math"
)

const (
	FISH_PROCESS_STATUS_UNKNOW        Base = iota // 0 未知的狀態
	FISH_PROCESS_STATUS_SHOOT                     // 1 魚機-射擊
	FISH_PROCESS_STATUS_HIT                       // 2 魚機-子彈擊中魚
	FISH_PROCESS_STATUS_FEATURE_SHOOT             // 3 魚機-特殊射擊(道具)
	FISH_PROCESS_STATUS_FEATURE_HIT               // 4 魚機-特殊子彈擊中魚(道具)
)

//=========================================================================================================================================================================
//       這邊都是魚機的回應封包定義區
//=========================================================================================================================================================================
// 魚機的座位資訊 結構 (記錄細節)
type SeatInfo_Fish struct {
	Seat_ID  int     `json:"seat_id"`  // 座位號碼
	User_ID  int64   `json:"user_id"`  // 加入玩家的 userid
	NickName string  `json:"nickname"` // 玩家暱稱
	Balance  int64   `json:"balance"`  // 客戶遊戲中餘額

	Bullet_type string  `json:"bullet_type"` // 選擇砲台 "1"=集中砲,"2"=鎖定砲
	Last_bet    int64   `json:"last_bet"`    // 當前押注值
}

//=========================================================================================================================================================================
// 魚機的進入遊戲 ( data 內的資料 )
type ResponseInfo_Fish_Enter_Game struct {
	Status          int     `json:"status"`          // 錯誤代碼
	Shoot_gap       float64 `json:"shoot_gap"`       // Auto射擊間隔時間
	Dlg_show_chance float64 `json:"dlg_show_chance"` // 魚出現隨機對話的機率
	Dlg_hit_chance  float64 `json:"dlg_hit_chance"`  // 擊殺後魚語音的機率
	TableID         string  `json:"table_id"`        // 桌號
	User_ID         int64   `json:"user_id"`         // 加入玩家的 userid
	Nickname        string  `json:"nickname"`        // 玩家暱稱
	Background      int     `json:"background"`      // 遊戲當前場景

	Balance_ci int64  `json:"balance_ci"` // 玩家想帶進來的錢
	Jp_prize   int64   `json:"jp_prize"`   // 當前JP累積金額
	Jp_bet     int64   `json:"jp_bet"`     // 獲取JP的最低押注

	Bet_options       [5]int                  `json:"bet_options"`       // 下注選項	[1,2,3,5,7,10]
	Bullet_options    [10]string              `json:"bullet_options"`    // 砲台選項	"1"=集中砲,"2"=鎖定砲
	Balance           int64                 `json:"balance"`           // 客戶遊戲中餘額
	Table_player_info [SEAT_MAX]SeatInfo_Fish `json:"table_player_info"` // 座位資訊
}

//=========================================================================================================================================================================
// 魚機的加入遊戲 ( data 內的資料 )
type ResponseInfo_Fish_Join_Game struct {
	Status        int    `json:"status"`          // 錯誤代碼
	TableID       string `json:"table_id"`        // 桌號
	TableArrayIdx int    `json:"table_array_idx"` // TableInfoList 的索引陣列 方便快速搜尋資料 ( 感覺綁在 member上可以省點時間 )

	NickName   string  `json:"nickname"`   // 玩家暱稱
	User_ID    int64   `json:"user_id"`    // 離開玩家的 userid
	Seat_ID    int     `json:"seat_id"`    // 座位
	Balance_ci int64   `json:"balance_ci"` // 玩家想帶進來的錢
}

//=========================================================================================================================================================================
// 魚機的離開遊戲 ( data 內的資料 )
type ResponseInfo_Fish_Exit_Game struct {
	Status  int    `json:"status"`   // 錯誤代碼
	TableID string `json:"table_id"` // 桌號

	NickName string `json:"nickname"` // 玩家暱稱
	User_ID  int64  `json:"user_id"`  // 離開玩家的 userid
	Seat_ID  int    `json:"seat_id"`  // 座位
}

//=========================================================================================================================================================================
// 遊戲記錄 結構 (記錄細節)
type GameLos_Fish struct {
	PlatformID int    // 第三方平台編號
	LobbyID    int    // 大廳編號
	GameID     int    // 遊戲編號
	TableID    string // 桌號
	Seat_ID    int    // 座位號碼
	GameMode   int8   // 遊戲模式 1:魚機 2:SLOT 3:撲克 4:麻將

	CreateTime string // 建立時間

	User_ID  int64  // 玩家帳號編號
	Account  string // 魚機帳號
	NickName string // 魚機暱稱

	Round int64 // 在遊戲內的第幾局

	Before_Balance_ci  int64 //之前的銭
	Before_Balance_win int64 //之前的銭
	Balance_ci         int64 // 玩家分數_投
	Balance_win        int64 // 玩家贏的錢   win 先扣,在扣 ci  隨遊戲不斷變動

	BetLevel int64 			// 單一押注

	Bet_Win        int64 	// 玩家贏分
	Process_Status int     // 玩家處理的狀態 0:unknow 1:shoot 2:hit 3:feature_shoot 4:feature_hit

	//================================== 沒使用 ===========================
	Interval_bet    int64 // 區間押注紀錄
	Interval_bet_pt int64 // 區間押注指標

	Avg_bet          int64  // 平均押注
	Progress_water   int64  // 獎項水池[Y]
	Progress_odds    int64  // 獎項倍數[Y]
	Progress_support int64  // 獎項貢獻[Y]
	Wait_item_id     string // 待中獎項[Y]

	Wait_item_seat  string // 待中座位[Y]
	Wait_item_odds  int64  // 待中倍數[Y]
	Wait_item_value int64  // 待中分數[Y]

	Win_item_id    string // 座位獎項名稱
	Win_item_bet   int64  // 座位獎項押注
	Win_item_value int64  // 座位獎項分數
	Win_item_win   int64  // 座位獎項贏分

}

/*
Fish_id
0, // 迦魶魚
1,// 小丑魚
2,// 鰈魚
3,// 河魨
4,// 獅子魚
5,// 比目魚
6,// 龍蝦
7,// 旗魚
8,// 章魚
9,// 燈籠魚
10,// 海龜
11,// 鋸齒鯊
12,// 蝠魟
13,// 巨大化--小丑魚
14,// 巨大化--鰈魚
15,// 巨大化--河魨
16,// 鯊魚
17,// 殺人鯨
18,// 座頭鯨
19,// 海王--火龍
20,// 海王--鱷魚
21,// 炸彈蟹
22,// 電磁蟹
23,// 電鑽蟹

Feature
0: 一般魚
1: 旋風魚
2: 閃電魚
3: 電磁砲 Type 22
4: 鑽頭砲Type 23
5: 炸彈Type 21

*/
//=========================================================================================================================================================================
// 產魚腳本 ( data 內的資料 )
type Script_Fish_info struct {
	Fish_id          int `json:"fish_id"`   // 魚的id
	Fish_type        int `json:"fish_type"` // 魚的類型
	Feature          int `json:"feature"`   // 特殊事件type
	Feature_position int `json:"feature_position"`
	X                int `json:"x"`          // 魚進場位置x軸
	Y                int `json:"y"`          // 魚進場位置y軸
	O                int `json:"o"`          // 魚進場角度
	Fish_num         int `json:"fish_num"`   // 生成魚的數量
	Fish_array       int `json:"fish_array"` // 魚陣id，需要定義魚陣型的腳本，事先放在client端    ??????????????????
}

//=========================================================================================================================================================================
// 魚機的離開遊戲 ( data 內的資料 )
type ResponseInfo_New_Fish struct {
	Status  int    `json:"status"`   // 錯誤代碼
	TableID string `json:"table_id"` // 桌號

	Background  int    `json:"background"`  // 背景id
	Script_name string `json:"script_name"` // 腳本名稱

	Fish_info []Script_Fish_info `json:"fish_info"` // 魚資訊(複合資料)
}

//=========================================================================================================================================================================
// Fish_Shoot ( data 內的資料 )
type PacketCmd_Fish_Shoot struct {
	TableID string `json:"table_id"` // 桌號
	User_ID int64  `json:"user_id"`  // 玩家的 userid
	Seat_ID int8   `json:"seat_id"`  // 座位

	X           int     `json:"x"`           // 魚進場位置x軸
	Y           int     `json:"y"`           // 魚進場位置y軸
	Bet         int64   `json:"bet"`         // 壓注額
	Bullet_type string  `json:"bullet_type"` // 子彈類型 1~2
	Bullet_id   string  `json:"bullet_id"`   // 子彈序號
}

//=========================================================================================================================================================================
// 魚機的Fish_Shoot ( data 內的資料 )
type ResponseInfo_Fish_Shoot struct {
	Status  int    `json:"status"`   // 錯誤代碼
	TableID string `json:"table_id"` // 桌號

	NickName string `json:"nickname"` // 玩家暱稱
	User_ID  int64  `json:"user_id"`  // 玩家的 userid
	Seat_ID  int    `json:"seat_id"`  // 座位

	X           int     `json:"x"`           // 魚進場位置x軸
	Y           int     `json:"y"`           // 魚進場位置y軸
	Bet         int64   `json:"bet"`         // 壓注額
	Bullet_type string  `json:"bullet_type"` // 子彈類型 1~2
	Bullet_id   string  `json:"bullet_id"`   // 子彈序號 (client端產生)

}

//=========================================================================================================================================================================
// Fish_Hit ( data 內的資料 )
type PacketCmd_Fish_Hit struct {
	TableID string `json:"table_id"` // 桌號
	User_ID int64  `json:"user_id"`  // 玩家的 userid
	Seat_ID int8   `json:"seat_id"`  // 座位

	Bet         int64   `json:"bet"`         // 押注額
	Bullet_type string  `json:"bullet_type"` // 子彈類型 1~2
	Bullet_id   string  `json:"bullet_id"`   // 子彈序號 (client端產生)

	Fish_id      []int `json:"fish_id"`      // 魚的id, int陣列
	Test_kill    int   `json:"test_kill"`    // 非正式環境才有的功能 必殺flag. 1=必殺, 0 = 必槓龜, 沒key=正常打擊
	Test_jackpot int   `json:"test_jackpot"` // 非正式環境才有的功能 1 = 必拉JP. (但bet必須大於minimal JP bet)
}

//=========================================================================================================================================================================
// 魚機的 kill_info 給獎的魚id陣列
type Fish_Kill_Info struct {
	Fish_id      int     `json:"fish_id"`      // 魚的編id
	Win          int64   `json:"win"`          // 得分金額
	Feature_type string  `json:"feature_type"` // 特殊事件類別
	Feature_id   string  `json:"feature_id"`   // 特殊事件id
}

//=========================================================================================================================================================================
// 魚機的Fish_Hit ( data 內的資料 )  待跟貓王確認格式 看不懂文件
type ResponseInfo_Fish_Hit struct {
	Status  int    `json:"status"`   // 錯誤代碼
	TableID string `json:"table_id"` // 桌號

	NickName string `json:"nickname"` // 玩家暱稱
	User_ID  int64  `json:"user_id"`  // 玩家的 userid
	Seat_ID  int    `json:"seat_id"`  // 座位

	Bet_Win int64   `json:"bet_win"` // 總共贏多少  (我的測試變數)

	Bet       int64   `json:"bet"`       // 押注額
	Balance   int64    `json:"balance"`   // 玩家座位上的銭  seatinfo.Balance_ci + seatinfo.Balance_win
	Bullet_id string  `json:"bullet_id"` // 子彈序號 (client端產生)

	Kill_info []Fish_Kill_Info `json:"kill_info"` // 給獎的魚id陣列

}

//=========================================================================================================================================================================
// Feature_Shoot ( data 內的資料 )
type PacketCmd_Fish_Feature_Shoot struct {
	TableID string `json:"table_id"` // 桌號
	User_ID int64  `json:"user_id"`  // 玩家的 userid
	Seat_ID int8   `json:"seat_id"`  // 座位

	X   int     `json:"x"`   // 魚進場位置x軸
	Y   int     `json:"y"`   // 魚進場位置y軸
	Bet float64 `json:"bet"` // 押注額

	Feature_type string `json:"feature_type"` // 特殊事件類別 "1"=旋風,"2"=閃電,"3"=雷射,"4"=鑽頭,"5"=炸彈
	Feature_id   string `json:"feature_id"`   // 特殊事件id
}

//=========================================================================================================================================================================
// 魚機的Feature_Shoot ( data 內的資料 )
type ResponseInfo_Fish_Feature_Shoot struct {
	Status  int    `json:"status"`   // 錯誤代碼
	TableID string `json:"table_id"` // 桌號

	NickName string `json:"nickname"` // 玩家暱稱
	User_ID  int64  `json:"user_id"`  // 玩家的 userid
	Seat_ID  int    `json:"seat_id"`  // 座位

	X            int     `json:"x"`            // 魚進場位置x軸
	Y            int     `json:"y"`            // 魚進場位置y軸
	Bet          float64 `json:"bet"`          // 押注額
	Feature_type string  `json:"feature_type"` // 特殊事件類別 "1"=旋風,"2"=閃電,"3"=雷射,"4"=鑽頭,"5"=炸彈
	Feature_id   string  `json:"feature_id"`   // 特殊事件id

	Odds float64 `json:"odds"` // 賠率 需參考賠率表
}

//=========================================================================================================================================================================
// Fish_Feature_Hit ( data 內的資料 )
type PacketCmd_Fish_Feature_Hit struct {
	TableID string `json:"table_id"` // 桌號
	User_ID int64  `json:"user_id"`  // 玩家的 userid
	Seat_ID int8   `json:"seat_id"`  // 座位

	Bet          int64   `json:"bet"`          // 押注額
	Feature_type string  `json:"feature_type"` // 特殊事件類別 "1"=旋風,"2"=閃電,"3"=雷射,"4"=鑽頭,"5"=炸彈
	Feature_id   string  `json:"feature_id"`   // 特殊事件id

	Fish_id []int  `json:"fish_id"` // 魚的id, int陣列
	Stage   string `json:"stage"`   // "0"：只要odds沒用完都直接死，鑽頭砲鑽到魚時會用。 "1"：正規化死亡率，讓傳上來的魚都有機會給獎，閃電、雷射、鑽頭的結尾、炸彈。
	//Test_kill    int   `json:"test_kill"`    // 非正式環境才有的功能 必殺flag. 1=必殺, 0 = 必槓龜, 沒key=正常打擊
	//Test_jackpot int   `json:"test_jackpot"` // 非正式環境才有的功能 1 = 必拉JP. (但bet必須大於minimal JP bet)
}

//=========================================================================================================================================================================
// 魚機的Fish_Feature_Hit ( data 內的資料 )  待跟貓王確認格式 看不懂文件
type ResponseInfo_Fish_Feature_Hit struct {
	Status  int    `json:"status"`   // 錯誤代碼
	TableID string `json:"table_id"` // 桌號

	NickName string `json:"nickname"` // 玩家暱稱
	User_ID  int64  `json:"user_id"`  // 玩家的 userid
	Seat_ID  int    `json:"seat_id"`  // 座位

	Bet_Win float64 `json:"bet_win"` // 總共贏多少  (我的測試變數)

	Bet         int64   `json:"bet"`         // 押注額
	Balance     int64   `json:"balance"`     // 玩家座位上的銭  seatinfo.Balance_ci + seatinfo.Balance_win
	Remain_odds int64   `json:"remain_odds"` // 沒用完的odds, 應該可以不用回傳(client目前沒用到)
	Feature_id  string  `json:"feature_id"`  // 特殊事件id

	Kill_info []Fish_Kill_Info `json:"kill_info"` // 給獎的魚id陣列

}

//●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●
//●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●
//●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●
// Fish_Shoot 開始玩
func Fish_Shoot(ClientID int, DecodeData string) (string, int, TableInfo, MemberInfo) {
	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值
	var BroadcastTable TableInfo           // 要廣播的桌子
	var BroadcastMember MemberInfo         // 發起廣播的玩家
	var seat_backup SeatInfo               // 備份的位子資訊
	CommonLog_INFO_Printf("#收到封包 Fish_Shoot(開始玩) ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	Fish_Shoot := PacketCmd_Fish_Shoot{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &Fish_Shoot)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		panic(err)
	} else {

		CommonLog_INFO_Printf("#TableID=%s, User_ID=%d, Seat_ID=%d, Bet(押注額)=%f, Bullet_type(子彈類型)=%s, Bullet_id(子彈序號)=%s",
			Fish_Shoot.TableID, Fish_Shoot.User_ID, Fish_Shoot.Seat_ID, Fish_Shoot.Bet, Fish_Shoot.Bullet_type, Fish_Shoot.Bullet_id)

		// 讀取玩家的socket資料
		pClient := Common_ClientInfoGet(ClientID)

		// 讀取玩家的會員資料
		pMember := &pClient.Member

		// 檢查玩家是否有登入
		//pClient.Member.Statu = 1
		if pClient.IsUse == true && pClient.Member.Status == 1 {

			// 加入 User_ID 檢查
			if Fish_Shoot.User_ID != pMember.User_ID {
				Code = int(ERROR_CODE_ERROR_USER_ID)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 檢查桌號是否存在
			CommonLog_INFO_Printf("檢查桌號是否存在 TableID=%s, TableArrayIdx=%d", Fish_Shoot.TableID, pMember.TableArrayIdx)
			pTable := Match_TableInfoGet(pMember.TableArrayIdx)
			if pTable.TableID != Fish_Shoot.TableID {
				Code = int(ERROR_CODE_NO_FIND_TABLE)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 檢查玩家是否在桌內
			var bFind bool = false
			var pSeat *SeatInfo // 取址
			//var SeatIdx int = 0
			for i := 0; i < SEAT_MAX; i++ {
				pSeat = &pTable.SeatInfo[i]

				CommonLog_INFO_Printf("i=%d, User_ID=%d, Account=%s, Seat_ID=%d", i, pSeat.User_ID, pSeat.Account, pSeat.Seat_ID)
				if pSeat.bUse == false {
					continue
				}

				if pSeat.User_ID == Fish_Shoot.User_ID {
					bFind = true
					//SeatIdx = i
					CommonLog_INFO_Printf("有找到 i=%d, TableID=%s, User_ID=%d, Account=%s", i, pTable.TableID, pSeat.User_ID, pSeat.Account)
					break
				}
			}
			if bFind == false {
				// 找不到位子
				Code = int(ERROR_CODE_NO_FIND_SEAT)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 檢查是否為 Fish 類型的遊戲
			if pTable.GameMode != int8(GAME_MODE_FISH) {
				Code = int(ERROR_CODE_ERROR_GAME_MODE)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 檢查桌內銭是否夠押注 (日後改絕對值)
			if pSeat.Balance_ci+pSeat.Balance_win-Fish_Shoot.Bet < 0 {
				Code = int(ERROR_CODE_TABLE_BALANCE_NOT_ENOUGH)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			CommonLog_INFO_Printf("#Fish_Shoot(開始玩) TableID=%s, User_ID=%d, Account=%s", Fish_Shoot.TableID, pSeat.User_ID, pSeat.Account)

			// 隨機產生結果
			//rand.Seed(time.Now().UnixNano())
			//var IsWin int = rand.Intn(2)        // 是否中獎
			//var Bet_Win float64 = 0             // 玩家贏分
			var FinishPlayerBalance int64 = 0 // 玩家此次金額的變化

			/*
				if IsWin > 0 {
					Bet_Win = (1 + float64(rand.Intn(10))) * Fish_Shoot.Bet // 玩家贏分
					FinishPlayerBalance = Bet_Win - Fish_Shoot.Bet          // 玩家此次金額的變化 ( 贏分 + 押注 )
					CommonLog_INFO_Printf("#Fish_Shoot(開始玩) 有中獎 TableID=%s, User_ID=%d, Account=%s, Bet_Win=%f, FinishPlayerBalance=%f",
						Fish_Shoot.TableID, pSeat.User_ID, pSeat.Account, Bet_Win, FinishPlayerBalance)
				} else {
					Bet_Win = 0

					CommonLog_INFO_Printf("#Fish_Shoot(開始玩) 沒中獎 TableID=%s, User_ID=%d, Account=%s, Bet_Win=%f, FinishPlayerBalance=%f",
						Fish_Shoot.TableID, pSeat.User_ID, pSeat.Account, Bet_Win, FinishPlayerBalance)
				}*/

			// 玩家射擊子彈的bet
			FinishPlayerBalance = -Fish_Shoot.Bet

			// 扣銭  改變 Balance_ci 或 Balance_win
			var before_Balance_ci int64 = pSeat.Balance_ci //之前的銭
			//var after_Balance_ci   float64 = 0					//spin之後的銭
			var before_Balance_win int64 = pSeat.Balance_win //之前的銭
			//var after_Balance_win  float64 = 0					//spin之後的銭

			var ChangeBalance_ci int64 = 0  // 銭的變化量
			var ChangeBalance_win int64 = 0 // 銭的變化量

			if pSeat.Balance_win+FinishPlayerBalance >= 0 {
				//pSeat.Balance_win -= FinishPlayerBalance
				ChangeBalance_win += FinishPlayerBalance
				CommonLog_INFO_Printf("#Fish_Shoot(開始玩) Balance_win 銭夠扣 before_Balance_win(之前)=%d, after_Balance_win(之後)=%d", before_Balance_win, (pSeat.Balance_win + ChangeBalance_win))
			} else {
				// Balance_ci 跟 Balance_win 一起扣
				if pSeat.Balance_win > 0 {
					var RemainBalance int64 = pSeat.Balance_win + FinishPlayerBalance // 這邊會是負數, 記綠欠的剩餘款項

					//pSeat.Balance_win += FinishPlayerBalance		// 先扣win
					//pSeat.Balance_ci += RemainBalance				// 剩餘款項再用 ci 去支付
					ChangeBalance_win += FinishPlayerBalance // 先扣win
					ChangeBalance_ci += RemainBalance        // 剩餘款項再用 ci 去支付
					CommonLog_INFO_Printf("#Fish_Shoot(開始玩) Balance_win 有銭, 分段扣 before_Balance_ci(之前)=%d, after_Balance_ci(之後)=%d <----> before_Balance_win(之前)=%d, after_Balance_win(之後)=%d",
						before_Balance_ci, (pSeat.Balance_ci + ChangeBalance_ci), before_Balance_win, (pSeat.Balance_win + ChangeBalance_win))
				} else {

					// 只扣 Balance_ci
					//pSeat.Balance_ci += FinishPlayerBalance
					ChangeBalance_ci += FinishPlayerBalance
					CommonLog_INFO_Printf("#Fish_Shoot(開始玩) Balance_ci 有銭, 直接扣 before_Balance_ci(之前)=%d, after_Balance_ci(之後)=%d", before_Balance_ci, (pSeat.Balance_ci + ChangeBalance_ci))
				}
			}

			// 檢查一下銭是否為負值
			if pSeat.Balance_ci < 0 {
				Code = int(ERROR_CODE_ERROR_DATA)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}
			if pSeat.Balance_win < 0 {
				Code = int(ERROR_CODE_ERROR_DATA)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}
			//after_Balance_ci  = pSeat.Balance_ci
			//after_Balance_win = pSeat.Balance_win

			// 先抓DB內的資料
			ret1, LoadSeat := Mysql_CommonSeatInfo_Get(pTable.TableID, *pSeat)

			// 檢查銭跟記憶體是否一致
			if ret1 == true {
				if LoadSeat.Balance_ci != pSeat.Balance_ci || LoadSeat.Balance_win != pSeat.Balance_win || LoadSeat.Account != pSeat.Account || LoadSeat.User_ID != pSeat.User_ID {
					Code = int(ERROR_CODE_BALANCE_CHECK_FAIL)
					CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
					return DataMsg, Code, BroadcastTable, BroadcastMember
				}
			} else {
				Code = int(ERROR_CODE_NO_FIND_ACCOUNT)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 更新DB內的錢 (座位)
			ret2 := Mysql_SeatInfo_ChangeBalance(*pSeat, ChangeBalance_ci, ChangeBalance_win)
			if ret2 == false {
				Code = int(ERROR_CODE_BALANCE_UPDATE_FAIL)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 更新記憶體
			pSeat.Balance_ci += ChangeBalance_ci
			pSeat.Balance_win += ChangeBalance_win
			CommonLog_INFO_Printf("#Fish_Shoot(開始玩) 更新銭的記憶體 before_Balance_ci(之前)=%d, after_Balance_ci(之後)=%d, ChangeBalance_ci(變化)=%d", before_Balance_ci, pSeat.Balance_ci, ChangeBalance_ci)
			CommonLog_INFO_Printf("#Fish_Shoot(開始玩) 更新銭的記憶體 before_Balance_win(之前)=%d, after_Balance_win(之後)=%d, ChangeBalance_win(變化)=%d", before_Balance_win, pSeat.Balance_win, ChangeBalance_win)

			// 廣播設定
			BroadcastTable = *pTable   // 拷貝廣播資訊
			BroadcastMember = *pMember // 取值(只是拿 Member.ClientID, 來廣播而已 )
			seat_backup = *pSeat       // 備份的位子資訊

			// 遊戲玩了第幾局了
			pSeat.Round++

			// 設定廣播給其它同桌玩家  "re"  or "fw"
			// 組合回傳字串
			ResponseFish_Shoot := ResponseInfo_Fish_Shoot{}
			ResponseFish_Shoot.Status = int(ERROR_CODE_SUCCESS)
			ResponseFish_Shoot.TableID = BroadcastTable.TableID
			ResponseFish_Shoot.Seat_ID = seat_backup.Seat_ID
			ResponseFish_Shoot.User_ID = seat_backup.User_ID
			ResponseFish_Shoot.NickName = seat_backup.NickName

			// 要廣播的子彈資訊
			ResponseFish_Shoot.Bet = Fish_Shoot.Bet
			ResponseFish_Shoot.Bullet_type = Fish_Shoot.Bullet_type
			ResponseFish_Shoot.Bullet_id = Fish_Shoot.Bullet_id
			ResponseFish_Shoot.X = Fish_Shoot.X
			ResponseFish_Shoot.Y = Fish_Shoot.Y

			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(ResponseFish_Shoot)
			if err != nil {
				CommonLog_WARNING_Println("json err:", err)
			}
			DataMsg = string(DataMsgByte)

			//===========================================================================================================================================
			// 寫入GameLog
			gamelog := GameLos_Fish{}
			gamelog.PlatformID = pTable.PlatformID
			gamelog.LobbyID = pTable.LobbyID
			gamelog.GameID = pTable.GameID
			gamelog.TableID = pTable.TableID
			gamelog.Seat_ID = pSeat.Seat_ID
			gamelog.GameMode = pTable.GameMode
			gamelog.CreateTime = Common_NowTimeGet()

			gamelog.User_ID = pSeat.User_ID
			gamelog.Account = pSeat.Account
			gamelog.NickName = pSeat.NickName

			gamelog.Round = pSeat.Round                             // 玩家在桌內的遊戲局數
			gamelog.Balance_ci = pSeat.Balance_ci                   // 目前實際的錢
			gamelog.Balance_win = pSeat.Balance_win                 // 目前實際的錢
			gamelog.BetLevel = ResponseFish_Shoot.Bet               // 玩家押注
			gamelog.Bet_Win = 0                                     // 單純贏分
			gamelog.Before_Balance_ci = before_Balance_ci           // 之前的錢
			gamelog.Before_Balance_win = before_Balance_win         // 之前的錢
			gamelog.Process_Status = int(FISH_PROCESS_STATUS_SHOOT) // 遊戲狀態
			// 寫入GameLog db
			Mysql_GameLosFish_Insert(gamelog)

		} else {
			Code = int(ERROR_CODE_NO_LOGIN)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
	return DataMsg, Code, BroadcastTable, BroadcastMember
}

//=========================================================================================================================================================================
// 處理 魚機的 產魚腳本
func Fish_Process(pTable *TableInfo) (bool, string, string, int, TableInfo, MemberInfo) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值
	//var TableID string = ""
	//var Seat_ID int = 0

	var BroadcastTable TableInfo   // 要廣播的桌子
	var BroadcastMember MemberInfo // 發起廣播的玩家 ( 空的 不指定要廣播給哪個玩家 )
	var Cmd string = ""            // 產魚的cmd
	var NeedBroadcast bool = false // 需要廣播時候

	// 判斷桌內有人
	if pTable.TablePlayerNow > 0 {

		// 產魚計數器
		pTable.UpdateTimeCount++

		// 計算時間差額
		var Diff int64 = pTable.UpdateTimeCount - pTable.StartTimeCount

		// 判斷是否到了產魚時間
		if Diff >= New_Fish_Time {
			CommonLog_INFO_Printf("#Fish_Process 時間相等 產魚開始 TableID:%s,  記數器count=[%d,%d]", pTable.TableID, pTable.StartTimeCount, pTable.UpdateTimeCount)
			pTable.UpdateTimeObj = Common_NowTimeObjGet()
			pTable.StartTimeCount = pTable.UpdateTimeCount

			// 廣播設定
			NeedBroadcast = true
			Cmd = NET_CMD_FISH_NEW_FISH
			Code = int(ERROR_CODE_SUCCESS)
			BroadcastTable = *pTable

			// 產魚腳本
			// 設定廣播給其它同桌玩家  "re"  or "fw"
			// 組合回傳字串
			New_Fish := ResponseInfo_New_Fish{}
			New_Fish.Status = int(ERROR_CODE_SUCCESS)
			New_Fish.TableID = BroadcastTable.TableID
			New_Fish.Background = 1
			New_Fish.Script_name = "test.json"

			ScriptInfo := Script_Fish_info{}
			ScriptInfo.Fish_id = 1
			ScriptInfo.Fish_num = 10
			ScriptInfo2 := Script_Fish_info{}
			ScriptInfo2.Fish_id = 2
			ScriptInfo2.Fish_num = 20

			New_Fish.Fish_info = append(New_Fish.Fish_info, ScriptInfo)
			New_Fish.Fish_info = append(New_Fish.Fish_info, ScriptInfo2)

			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(New_Fish)
			if err != nil {
				CommonLog_WARNING_Println("json err:", err)
			}
			DataMsg = string(DataMsgByte)

		}
	}

	return NeedBroadcast, Cmd, DataMsg, Code, BroadcastTable, BroadcastMember
}
