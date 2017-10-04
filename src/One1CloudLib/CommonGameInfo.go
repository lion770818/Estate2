package One1CloudLib

import (
	"encoding/json"
	"fmt"
	"time"
)

//=========================================================================================================================================================================
// 定義遊戲模式
const (
	GAME_MODE_NULL    Base = iota // 0 沒有定義
	GAME_MODE_FISH                // 1 魚機
	GAME_MODE_SLOT                // 2 Slot
	GAME_MODE_POKER               // 3 撲克牌
	GAME_MODE_MAHJONG             // 4 麻將
)

const (
	New_Fish_Time int64 = 20 // 產魚的周期時間 單位秒數
)

//=========================================================================================================================================================================
// 共用的遊戲結構, 如果是專屬某個遊戲專用
//=========================================================================================================================================================================
// 遊戲 結構
type GameInfo struct {
	PlatformID       int    // 第三方平台編號
	GameID           int    // 遊戲編號
	GameName         string // 遊戲中文名稱
	GameEnName       string // 遊戲英文名稱
	GameMode         int8   // 遊戲模式 1:魚機 2:SLOT 3:撲克 4:麻將
	TableDestoryMode int8   // 0: unknow 1:散桌後刪除此桌資訊  2: 散桌後保留此桌資訊,等待玩家重新入桌
	TablePlayerMax   int8   // 桌內人數上限
}

//=========================================================================================================================================================================
// 所有Game的List
var (
	GameInfoList = make(map[int]GameInfo) // map GameInfo
)

//=========================================================================================================================================================================
// 桌子 結構 (只記錄桌號和散桌行為)
type TableInfo struct {
	//================ db 內的資料 ================
	PlatformID    int    // 第三方平台編號
	LobbyID       int    // 大廳編號
	GameID        int    // 遊戲編號
	TableID       string // 桌號 桌子編號規則 =  英文縮寫(2碼) + 平台編號(1碼) + GameID(4碼) + '-' +  流水編號(7碼), 例如 FH11001-0000001 代表魚機 0000001 桌 ( 最大可百萬桌)
	TableArrayIdx int    // TableInfoList 的索引陣列 方便快速搜尋資料

	TablePlayerMax   int8   // 房間上限人數
	TablePlayerNow   int8   // 房間目前人數
	GameMode         int8   // 遊戲模式 1:魚機 2:SLOT 3:撲克 4:麻將
	TableDestoryMode int8   // 0: unknow 1:散桌後刪除此桌資訊  2: 散桌後保留此桌資訊,等待玩家重新入桌
	CreateTime       string // 紀錄開桌時間
	UpdateTime       string // 開桌, 玩家入桌, 玩家離桌, 都更新此時間
	LobbyMatchID     int8   // 大廳配桌編號, 編號相同的廳館, 才可以配桌在一起, if( LobbyMatchID 相同  && GameID 相同 && 人數未滿 ) 則加入桌()

	//============= 只存在 記憶體的資料  =======
	SeatInfo [SEAT_MAX]SeatInfo // 每個座位的資料
	//SeatInfo []SeatInfo 		// 每個座位的資料
	bUse            bool      // 桌子是否有使用
	UpdateTimeObj   time.Time // 每桌的更新時間 (產魚 or dosomething)
	StartTimeCount  int64     // 每桌的更新時間 (產魚 or dosomething)
	UpdateTimeCount int64     // 每桌的更新時間 (產魚 or dosomething)

	value interface{}
}

const TABLEINFO_MAX int = 5000 // 桌子最大數量
const SEAT_MAX int = 4         // 一桌內的座位數量

var TableInfoCount = 0                     // 目前桌子數量
var TableInfoList [TABLEINFO_MAX]TableInfo // 桌子陣列

//=========================================================================================================================================================================
// 座位資訊 結構 (記錄細節)
type SeatInfo struct {
	id int64 // 流水號的id  自動遞增, 作主KEY使用 的索引陣列 方便快速搜尋資料

	PlatformID int    // 第三方平台編號
	LobbyID    int    // 大廳編號
	GameID     int    // 遊戲編號
	TableID    string // 桌號
	Seat_ID    int    // 座位號碼
	GameMode   int8   // 遊戲模式 1:魚機 2:SLOT 3:撲克 4:麻將

	CreateTime string // 建立時間
	UpdateTime string // 更新時間

	User_ID  int64  // 玩家帳號編號
	Account  string // 魚機帳號
	NickName string // 魚機暱稱

	Balance_ci  int64 // 玩家分數_投
	Balance_win int64 // 玩家贏的錢   win 先扣,在扣 ci  隨遊戲不斷變動
	BetLevel    int64 // 單一押注

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

	//============= 只存在 記憶體的資料  =======
	bUse     bool  // 位子是否有使用
	ClientID int   //Client端的ID
	Round    int64 //遊戲玩了第幾局了
}

//=========================================================================================================================================================================
// 抓取GameID
func Common_GameInfoGet(PlatformID int, GameID int) GameInfo {

	var pGame GameInfo
	CommonLog_INFO_Printf("#Common_GameInfoGet 取得一個 GameInfo 的記憶體位置... PlatformID=%d, GameID=%d", PlatformID, GameID)

	// 取出記憶體位置
	for i := 0; i < len(GameInfoList); i++ {
		pGameTmp := GameInfoList[i]
		if pGameTmp.PlatformID == PlatformID && pGameTmp.GameID == GameID {
			CommonLog_INFO_Printf("有找到 i=%d, PlatformID=%d, GameID=%d", i, PlatformID, GameID)
			pGame = pGameTmp
			break
		}
	}
	CommonLog_INFO_Printf("#Common_GameInfoGet 取得一個 GameInfo 的記憶體位置... PlatformID=%d, GameID=%d, GameName=%s, GameMode=%d", PlatformID, GameID, pGame.GameName, pGame.GameMode)
	return pGame
}

//=========================================================================================================================================================================
// 取得一個 TableInfo 的記憶體位置
func Match_TableInfoGet(TableArrayIdx int) *TableInfo {

	var pTable *TableInfo = nil
	//CommonLog_INFO_Printf("#Common_TableInfoGet 取得一個 桌子 的記憶體位置... TableArrayIdx=%d", TableArrayIdx)

	if TableArrayIdx >= 0 && TableArrayIdx < TABLEINFO_MAX {
		// 取出記憶體位置
		pTable = &TableInfoList[TableArrayIdx]
	} else {
		pTable = nil
		CommonLog_WARNING_Printf("#Match_TableInfoGet 錯誤的TableArrayIdx=%d", TableArrayIdx)
	}

	//CommonLog_INFO_Printf("#Common_TableInfoGet 取得一個 桌子 的記憶體位置... TableArrayIdx=%d pTable=%p", TableArrayIdx, pTable)
	return pTable
}

//=========================================================================================================================================================================
// 找尋可以加入的桌子 (簡單,但是桌子數量多就效能慢)
func Match_TableInfoTableJoinSeatch(PlatformID int, GameID int, Member MemberInfo) (*TableInfo, bool) {

	var pTable *TableInfo = nil
	var bfind bool = false
	var findCount int = 0
	var bsameAccount bool = false
	CommonLog_INFO_Printf("#Match_TableInfoTableJoinSeatch PlatformID=%d, GameID=%d, TableInfoCount=%d", PlatformID, GameID, TableInfoCount)

	// 掃描整個array
	for i := 0; i < TABLEINFO_MAX; i++ {
		pTable = Match_TableInfoGet(i)

		if pTable.bUse == false {
			// 沒使用的桌子
			continue
		}

		findCount++
		if findCount > TableInfoCount {
			// 已經快速搜尋完全部資料了
			bfind = false
			CommonLog_INFO_Printf("#Match_TableInfoTableJoinSeatch 已經快速搜尋完全部資料了, 離開迴圈 findCount=%d, TableInfoCount=%d", findCount, TableInfoCount)
			break
		}

		if pTable.TablePlayerMax-pTable.TablePlayerNow <= 0 {
			// 人滿
			continue
		}

		if pTable.PlatformID == PlatformID && pTable.GameID == GameID {

			// 	檢查是否有同樣ID
			bsameAccount = false
			for j := 0; j < SEAT_MAX; j++ {
				Seat := pTable.SeatInfo[j]
				if Seat.User_ID == Member.User_ID {
					bsameAccount = true
					CommonLog_INFO_Printf("此桌有相同玩家 TableID=%s, Account=%s, User_ID=%d", pTable.TableID, Member.Account, Member.User_ID)
					break
				}
			}

			if bsameAccount == false {
				pTable.TableArrayIdx = i
				bfind = true
				CommonLog_INFO_Printf("找到空桌 PlatformID=%d, GameID=%d, TableID=%s, GameMode=%d, TableDestoryMode=%d, 目前人數[%d/%d]",
					PlatformID, GameID, pTable.TableID, pTable.GameMode, pTable.TableDestoryMode, pTable.TablePlayerNow, pTable.TablePlayerMax)
				break
			}

		}
	}

	return pTable, bfind
}

//=========================================================================================================================================================================
// 找尋可以加入的全新桌子
func Match_TableIdleSeatch() *TableInfo {

	var pTable *TableInfo = nil
	var bfind bool = false
	CommonLog_INFO_Printf("#Match_TableIdleSeatch(找尋可以加入的全新桌子) TableInfoCount=%d", TableInfoCount)

	for i := 0; i < TABLEINFO_MAX; i++ {
		pTable = Match_TableInfoGet(i)

		if pTable.bUse == true {
			continue
		}
		if pTable.TablePlayerNow != 0 {
			// 有人
			CommonLog_WARNING_Printf("#Match_TableIdleSeatch 怎會有幽靈人口? i=%d, TableID=%s, TablePlayerNow=%d", i, pTable.TableID, pTable.TablePlayerNow)
			continue
		}

		// 將桌子設定成已使用
		pTable.bUse = true
		pTable.TableArrayIdx = i
		bfind = true
		break
	}

	if bfind == true {
		CommonLog_INFO_Printf("#Match_TableIdleSeatch 有找到 TableArrayIdx=%d, TableID=%s", pTable.TableArrayIdx, pTable.TableID)
	} else {
		pTable = nil
		CommonLog_INFO_Printf("#Match_TableIdleSeatch 沒有找到 TableInfoCount=%d", TableInfoCount)
	}

	return pTable
}

//=========================================================================================================================================================================
// 開新桌
func Match_OpenTable(ClientID int, lobbyinfo LobbyInfo, game GameInfo) *TableInfo {

	CommonLog_INFO_Printf("#Match_OpenTable(開新桌)---開始 PlatformID=%d, LobbyID=%d, GameID=%d, GameMode=%d, GameName=%s, TableInfoCount=%d", lobbyinfo.PlatformID, lobbyinfo.LobbyID, game.GameID, game.GameMode, game.GameName, TableInfoCount)

	pClient := Common_ClientInfoGet(ClientID)
	if DEBUG_TEST_MUTEX_1 == true {
		CommonLog_INFO_Printf("#Match_OpenTable Lock ClientID=%d", ClientID)
		pClient.DataMutex.Lock()
	}

	//1. 用PlatformID, GameID 抓取GameInfo

	// 找到空人的桌子
	pNewTable := Match_TableIdleSeatch()
	if pNewTable == nil {
		CommonLog_WARNING_Printf("#Match_OpenTable(警告, 找不到桌子) ClientID=%d, PlatformID=%d, LobbyID=%d, GameID=%d", ClientID, game.PlatformID, lobbyinfo.LobbyID, game.GameID)

		if DEBUG_TEST_MUTEX_1 == true {
			CommonLog_INFO_Printf("#Match_OpenTable Unlock1 ClientID=%d", ClientID)
			pClient.DataMutex.Unlock()
		}
		return nil
	}

	//2. 寫入tableinfo db
	pNewTable.bUse = true
	pNewTable.PlatformID = game.PlatformID
	pNewTable.LobbyID = lobbyinfo.LobbyID
	pNewTable.GameID = game.GameID
	pNewTable.TableID = Match_TableIdMake(game) //make 桌號
	//pNewTable.TableArrayIdx 	= TableArrayIdx  // 已經設定過

	pNewTable.TablePlayerMax = game.TablePlayerMax
	pNewTable.TablePlayerNow = 0
	pNewTable.GameMode = game.GameMode
	pNewTable.TableDestoryMode = game.TableDestoryMode
	pNewTable.CreateTime = Common_NowTimeGet()
	pNewTable.UpdateTime = Common_NowTimeGet()
	pNewTable.UpdateTimeObj = Common_NowTimeObjGet()
	pNewTable.StartTimeCount = 0
	pNewTable.UpdateTimeCount = 0
	pNewTable.LobbyMatchID = lobbyinfo.LobbyMatchID

	Mysql_CommonTableInfo_Insert(*pNewTable)

	//3. 更新 tableinfo 記憶體
	TableInfoCount++

	//4. 呼叫 jointable()
	//jointable()
	if DEBUG_TEST_MUTEX_1 == true {
		CommonLog_INFO_Printf("#Match_OpenTable Unlock2 ClientID=%d", ClientID)
		pClient.DataMutex.Unlock()
	}

	CommonLog_INFO_Printf("#Match_OpenTable(開新桌)---結束 PlatformID=%d, LobbyID=%d, GameID=%d, TableID=%s, GameMode=%d, GameName=%s, TableInfoCount=%d", lobbyinfo.PlatformID, lobbyinfo.LobbyID, game.GameID, pNewTable.TableID, game.GameMode, game.GameName, TableInfoCount)

	return pNewTable
}

//=========================================================================================================================================================================
// 加入桌
func Match_JoinTable(ClientID int, lobbyinfo LobbyInfo, gameinfo GameInfo, pMember *MemberInfo, pTable *TableInfo, joinTableID string, Balance_ci int64) (string, int, int) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值

	CommonLog_INFO_Printf("#Match_JoinTable(加入桌) ClientID=%d, joinTableID=%s, Account=%s, User_ID=%d, Balance_ci(玩家帶入桌的錢)=%d", ClientID, joinTableID, pMember.Account, pMember.User_ID, Balance_ci)
	// 檢查桌號是否存在
	if len(joinTableID) > 0 {
		if pTable.TableID != joinTableID {
			Code = int(ERROR_CODE_NO_FIND_TABLE)
			CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
			return DataMsg, Code, 0
		}
	}

	// 檢查桌子使否使用中
	if pTable.bUse == false {
		Code = int(ERROR_CODE_NO_USE)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s, joinTableID=%s", Code, ErrorCode[Code].Message, DataMsg, joinTableID)
		return DataMsg, Code, 0
	}
	CommonLog_INFO_Printf("有找到 TableArrayIdx=%d, PlatformID=%d, LobbyID=%d, GameID=%d, TableID=%s, GameMode=%d, TableDestoryMode=%d, 桌內人數=[%d/%d]",
		pTable.TableArrayIdx, pTable.PlatformID, pTable.LobbyID, pTable.GameID, pTable.TableID, pTable.GameMode, pTable.TableDestoryMode, pTable.TablePlayerNow, pTable.TablePlayerMax)

	// 檢查桌子是否滿人
	if pTable.TablePlayerMax-pTable.TablePlayerNow <= 0 {
		Code = int(ERROR_CODE_FULL_PLAYER)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, 桌內人數=[%d/%d]", Code, ErrorCode[Code].Message, pTable.TablePlayerNow, pTable.TablePlayerMax)
		return DataMsg, Code, 0
	}

	// 複驗檢查錢是否足夠
	if pMember.Balance-Balance_ci < 0 {
		Code = int(ERROR_CODE_CARRY_BALANCE_NOT_ENOUGH)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, 桌內人數=[%d/%d]", Code, ErrorCode[Code].Message, pTable.TablePlayerNow, pTable.TablePlayerMax)
		return DataMsg, Code, 0
	}

	// 找到可加入的位子
	var bFind bool = false     // 是否有發現
	var bIsSameAccount = false // 相同帳號檢測
	var pSeat *SeatInfo = nil
	var Seat_ID int = 0
	for i := 0; i < SEAT_MAX; i++ {

		pSeat = &pTable.SeatInfo[i]
		CommonLog_INFO_Printf("開始找位子 i=%d Account=%s, User_ID=%d, Seat_ID=%d", i, pSeat.Account, pSeat.User_ID, pSeat.Seat_ID)
		if pSeat.bUse == true {
			continue
		}

		bFind = true
		Seat_ID = i
		CommonLog_INFO_Printf("找到位子")
		break
	}
	if bFind == false {
		Code = int(ERROR_CODE_NO_FIND_SEAT)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, 人數=[%d/%d]", Code, ErrorCode[Code].Message, pTable.TablePlayerNow, pTable.TablePlayerMax)
		return DataMsg, Code, 0
	}

	//檢查是否有重複帳號入桌
	for i := 0; i < SEAT_MAX; i++ {

		pSeatcheck := pTable.SeatInfo[i]
		CommonLog_INFO_Printf("檢查是否有重複帳號入桌 i=%d Account=%s, User_ID=%d, Seat_ID=%d", i, pSeatcheck.Account, pSeatcheck.User_ID, pSeatcheck.Seat_ID)
		if pSeatcheck.bUse == false {
			continue
		}

		if pSeatcheck.User_ID == pMember.User_ID {
			bIsSameAccount = true
			CommonLog_INFO_Printf("重複入桌 TableID=%s, User_ID=%d", pTable.TableID, pMember.User_ID)
			break
		}
	}
	if bIsSameAccount == true {
		Code = int(ERROR_CODE_RE_JOIN_TABLE)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
		return DataMsg, Code, 0
	}

	// 更新記憶體資料 tableinfo  and seatinfo
	pTable.TablePlayerNow++

	// 新位子資訊
	pNewSeatInfo := &pTable.SeatInfo[Seat_ID]
	pNewSeatInfo.bUse = true
	pNewSeatInfo.User_ID = pMember.User_ID
	pNewSeatInfo.Account = pMember.Account
	pNewSeatInfo.NickName = pMember.NickName
	pNewSeatInfo.Balance_ci = Balance_ci // pMember.Balance_ci
	pNewSeatInfo.Balance_win = 0         // 每次進桌都是0嗎, 再找志清check一下
	pNewSeatInfo.PlatformID = gameinfo.PlatformID
	pNewSeatInfo.LobbyID = lobbyinfo.LobbyID
	pNewSeatInfo.GameID = gameinfo.GameID
	pNewSeatInfo.GameMode = gameinfo.GameMode
	pNewSeatInfo.TableID = pTable.TableID
	pNewSeatInfo.Seat_ID = Seat_ID
	pNewSeatInfo.ClientID = ClientID

	// 等機率東西出來, 在一併寫入db

	pNewSeatInfo.CreateTime = Common_NowTimeGet()
	pNewSeatInfo.UpdateTime = Common_NowTimeGet()

	// 更新DB資料 更新桌
	Mysql_CommonTableInfo_Update(*pTable)
	// 更新DB資料 加入位子
	Mysql_CommonSeatInfo_Insert(*pNewSeatInfo)

	// 更新DB資料 更新玩家的錢
	// 1. 讀取DB 玩家的錢
	// 2. 寫入DB
	// 2. 檢查 DB 跟記憶體的錢是否一致
	var before_Balance int64 = 0
	var after_Balance int64 = 0
	// 檢查 DB 跟記憶體的錢是否一致
	// 重新讀取DB 玩家的錢
	ret, LoadMember := Mysql_CommonMemberInfoGet(pMember.Account, pMember.Password)
	if ret == true {

		// 因為多開, 所以加入桌時候, 要重新跟DB要一次新的資料
		pMember.Balance = LoadMember.Balance
		before_Balance = pMember.Balance

		// 2.寫入DB ( 用DB的錢 - Balance_ci )
		ret = Mysql_CommonMemberInfo_SubBalance(LoadMember, Balance_ci)
		if ret == false {
			Code = int(ERROR_CODE_BALANCE_UPDATE_FAIL)
			CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
			return DataMsg, Code, 0
		}

		after_Balance = pMember.Balance - Balance_ci

		ret, LoadMember2 := Mysql_CommonMemberInfoGet(pMember.Account, pMember.Password)
		if ret == true {

			// 比對錢是否正確
			if LoadMember2.Balance != after_Balance {
				// 發生錯誤了
				CommonLog_INFO_Printf("#Match_JoinTable(加入桌) 警告! 錢對不起來, ClientID=%d, joinTableID=%s, Account=%s, User_ID=%d, Balance_ci(玩家帶入桌的錢)=%d Balance[%d!=%d]",
					ClientID, joinTableID, pMember.Account, pMember.User_ID, Balance_ci, LoadMember.Balance, pMember.Balance)

				Code = int(ERROR_CODE_BALANCE_UPDATE_FAIL)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
				return DataMsg, Code, 0
			} else {

				pMember.Balance = pMember.Balance - Balance_ci
				pMember.TableArrayIdx = pTable.TableArrayIdx // member 存一份幫助搜尋
				CommonLog_INFO_Printf("#Match_JoinTable(加入位子完成) 快速搜尋 ClientID=%d, TableID=%s, TableArrayIdx=%d, Account=%s, User_ID=%d, Balance_ci(帶入桌的錢)=%d, before_Balance=%d, after_Balance=%d",
					ClientID, pTable.TableID, pMember.TableArrayIdx, pMember.Account, pMember.User_ID, Balance_ci, before_Balance, pMember.Balance)
			}

			/*
				//pMember.Balance = pMember.Balance - Balance_ci
				pMember.Balance = LoadMember2.Balance
				pMember.TableArrayIdx = pTable.TableArrayIdx // member 存一份幫助搜尋
				CommonLog_INFO_Printf("#Match_JoinTable(加入位子完成) 快速搜尋 ClientID=%d, TableID=%s, TableArrayIdx=%d, Account=%s, User_ID=%d, Balance_ci(帶入桌的錢)=%f, before_Balance=%f, pMember.Balance=%f,after_Balance=%f",
					ClientID, pTable.TableID, pMember.TableArrayIdx, pMember.Account, pMember.User_ID, Balance_ci, before_Balance, pMember.Balance, after_Balance )
			*/
		} else {
			// 找不到玩家
			Code = int(ERROR_CODE_NO_FIND_ACCOUNT)
			CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
			return DataMsg, Code, 0
		}

	} else {
		Code = int(ERROR_CODE_NO_FIND_ACCOUNT)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
		return DataMsg, Code, 0
	}

	return DataMsg, Code, Seat_ID
}

//=========================================================================================================================================================================
// 離開桌
func Match_ExitTable(ClientID int, pTable *TableInfo, pMember *MemberInfo) (string, int, TableInfo, MemberInfo) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值
	var BroadcastTable TableInfo           // 要廣播的桌子
	var BroadcastMember MemberInfo         // 發起廣播的玩家
	var seat_backup SeatInfo               // 備份的位子資訊

	CommonLog_INFO_Printf("#Match_ExitTable ClientID=%d, Account=%s, User_ID=%d", ClientID, pMember.Account, pMember.User_ID)

	// 讀取玩家的socket資料
	pClient := Common_ClientInfoGet(ClientID)

	// 檢查玩家是否有登入
	if pClient.IsUse == true && pClient.Member.Status == 1 {

		CommonLog_INFO_Printf("目前桌子資訊 TableID=%s, TableArrayIdx=%d, GameID=%d, GameMode=%d", pTable.TableID, pMember.TableArrayIdx, pTable.GameID, pTable.GameMode)

		// 檢查玩家是否在桌內
		var bFind bool = false
		var pSeat *SeatInfo
		//var SeatIdx int = 0
		for i := 0; i < SEAT_MAX; i++ {
			pSeat = &pTable.SeatInfo[i]

			CommonLog_INFO_Printf("i=%d, User_ID=%d, Account=%s, Seat_ID=%d", i, pSeat.User_ID, pSeat.Account, pSeat.Seat_ID)
			if pSeat.bUse == false {
				continue
			}

			if pSeat.User_ID == pMember.User_ID {
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

		CommonLog_INFO_Printf("#離開遊戲 TableID=%s, User_ID=%d, Account=%s", pTable.TableID, pSeat.User_ID, pSeat.Account)

		// 離開遊戲 設定
		BroadcastTable = *pTable   // 拷貝廣播資訊
		BroadcastMember = *pMember // 取值(只是拿 Member.ClientID, 來廣播而已 )
		seat_backup = *pSeat       // 備份的位子資訊

		// 因為同帳號多開, 所以重新讀取一下
		ret, LoadMember0 := Mysql_CommonMemberInfoGet(pMember.Account, pMember.Password)
		if ret == false {
			Code = int(ERROR_CODE_NO_FIND_ACCOUNT)
			CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
			return DataMsg, Code, BroadcastTable, BroadcastMember
		}
		// 重新更新
		pMember.Balance = LoadMember0.Balance

		//==========================================================================================================================================
		// 攜帶的錢返回 memberinfo 的 balance ( Balance_ci + Balance_win )
		//@@@
		// 重新讀取DB 玩家的錢
		var before_Balance int64 = pMember.Balance
		var after_Balance int64 = 0
		var table_Balance int64 = pSeat.Balance_ci + pSeat.Balance_win // 桌內目前的錢
		ret, LoadMember := Mysql_CommonMemberInfoGet(pMember.Account, pMember.Password)
		if ret == true {

			// 2.寫入DB ( 用DB的錢 - Balance_ci )
			ret = Mysql_CommonMemberInfo_AddBalance(LoadMember, table_Balance)
			if ret == false {
				Code = int(ERROR_CODE_BALANCE_UPDATE_FAIL)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 計算預期的金額
			after_Balance = pMember.Balance + table_Balance

			ret, LoadMember2 := Mysql_CommonMemberInfoGet(pMember.Account, pMember.Password)
			if ret == true {

				// 比對錢是否正確
				if LoadMember2.Balance != after_Balance {
					// 發生錯誤了
					CommonLog_INFO_Printf("#Match_ExitTable(離開桌) 警告! 錢對不起來, ClientID=%d, TableID=%s, Account=%s, User_ID=%d, table_Balance(桌內目前的錢)=%d Balance[%d!=%d]",
						ClientID, pTable.TableID, pMember.Account, pMember.User_ID, table_Balance, LoadMember.Balance, pMember.Balance)

					Code = int(ERROR_CODE_BALANCE_UPDATE_FAIL)
					CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
					return DataMsg, Code, BroadcastTable, BroadcastMember
				} else {

					// 更新玩家的錢
					pMember.Balance = pMember.Balance + table_Balance

					pMember.TableArrayIdx = pTable.TableArrayIdx // member 存一份幫助搜尋
					CommonLog_INFO_Printf("#Match_ExitTable(離開桌) 快速搜尋 ClientID=%d, TableID=%s, TableArrayIdx=%d, Account=%s, User_ID=%d, table_Balance(桌內目前的錢)=%d, before_Balance=%d, after_Balance=%d",
						ClientID, pTable.TableID, pMember.TableArrayIdx, pMember.Account, pMember.User_ID, table_Balance, before_Balance, pMember.Balance)
				}

			}

		} else {
			Code = int(ERROR_CODE_NO_FIND_ACCOUNT)
			CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
			return DataMsg, Code, BroadcastTable, BroadcastMember
		}

		// 刪除座位資訊DB
		Mysql_CommonSeatInfo_Delete(pTable.PlatformID, pTable.LobbyID, pTable.GameID, pTable.TableID, pSeat.Seat_ID, pSeat.User_ID)

		//===========================================================================================================================================

		// 通知機率系統把錢算完

		// 寫入GameLog

		// 刪除位子記憶體資料
		*pSeat = SeatInfo{}

		// 將玩家資料關連跟桌子關連解除
		pMember.TableArrayIdx = -1

		// 桌子人數減1
		pTable.TablePlayerNow--
		Mysql_CommonTableInfo_Update(*pTable)
		CommonLog_INFO_Printf("目前桌子人數 TableID=%s, TablePlayerNow=%d", pTable.TableID, pTable.TablePlayerNow)

		if pTable.TableDestoryMode == 1 {

			CommonLog_INFO_Printf("散桌後刪除此桌資訊 TableID=%s, TableDestoryMode=%d, 人數[%d/%d]", pTable.TableID, pTable.TableDestoryMode, pTable.TablePlayerNow, pTable.TablePlayerMax)

			// 判斷人數
			if pTable.TablePlayerNow <= 0 {

				// 刪除桌資訊DB  PlatformID int, LobbyID int, GameID int, TableID string, TableArrayIdx int
				Mysql_CommonTableInfo_Delete(pTable.PlatformID, pTable.LobbyID, pTable.GameID, pTable.TableID, pTable.TableArrayIdx)

				// 刪除桌資訊記憶體資料
				pTable.bUse = false
				*pTable = TableInfo{}
			} else {
				CommonLog_INFO_Printf("還有人在桌內, 所以不用刪除桌 TableID=%s, 人數[%d/%d]", pTable.TableID, pTable.TablePlayerNow, pTable.TablePlayerMax)
			}

		} else {
			CommonLog_INFO_Printf("散桌後保留此桌資訊,等待玩家重新入桌 TableID=%s, TableDestoryMode=%d", pTable.TableID, pTable.TableDestoryMode)
		}

		// 除錯用
		Mysql_CommonTableInfoShow()

		// 設定廣播給其它同桌玩家  "re"  or "fw"
		// 組合回傳字串
		fishExitGame := ResponseInfo_Fish_Exit_Game{}
		fishExitGame.Status = int(ERROR_CODE_SUCCESS)
		fishExitGame.TableID = BroadcastTable.TableID
		fishExitGame.Seat_ID = seat_backup.Seat_ID
		fishExitGame.User_ID = seat_backup.User_ID
		fishExitGame.NickName = seat_backup.NickName

		// 物件轉成json字串
		DataMsgByte, err := json.Marshal(fishExitGame)
		if err != nil {
			CommonLog_WARNING_Println("json err:", err)
		}
		DataMsg = string(DataMsgByte)

	} else {
		Code = int(ERROR_CODE_NO_LOGIN)
	}

	CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
	return DataMsg, Code, BroadcastTable, BroadcastMember
}

//=========================================================================================================================================================================
// 製造桌號  桌子編號規則 =  英文縮寫(2碼) + 平台編號(1碼) + GameID(4碼) + 流水編號(7碼) 例如 FH 1 1001 0000001 代表魚機 0000001 桌 ( 最大可百萬桌)
func Match_TableIdMake(game GameInfo) string {

	var id int64 = 0
	var TableID string = ""
	CommonLog_INFO_Printf("#Match_TableIdMake(製造桌號) GameEnName=%s, PlatformID=%dm GameID=%d", game.GameEnName, game.PlatformID, game.GameID)

	id = Mysql_CommonTableInfo_IdGet()
	if id == -1 {
		id = 1
		CommonLog_INFO_Printf("因為都沒桌子, 所以自己 new 1 一個編號")
	}
	TableID = fmt.Sprintf("%s%d%04d-%07d", game.GameEnName, game.PlatformID, game.GameID, id)

	CommonLog_INFO_Printf("#Match_TableIdMake(製造桌號) TableID=%s", TableID)
	return TableID
}

//=========================================================================================================================================================================
// 進入遊戲
func Common_EnterGame(ClientID int, DecodeData string) (string, int, TableInfo, MemberInfo) {

	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值
	var TableID string = ""
	var Seat_ID int = 0

	var BroadcastTable TableInfo   // 要廣播的桌子
	var BroadcastMember MemberInfo // 發起廣播的玩家

	CommonLog_INFO_Printf("#收到封包 Common_EnterGame ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	enter_Game := PacketCmd_Enter_Game{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &enter_Game)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code, BroadcastTable, BroadcastMember
	} else {

		CommonLog_INFO_Printf("#PlatformID=%d, LobbyID=%d, GameID=%d, User_ID=%d, Balance_ci=%d", enter_Game.PlatformID, enter_Game.LobbyID, enter_Game.GameID, enter_Game.User_ID, enter_Game.Balance_ci)

		// 讀取玩家的socket資料
		// 檢查玩家是否有登入
		pClient := Common_ClientInfoGet(ClientID)
		//pClient.MemberInfo.Status = 1 // 假資料
		if pClient.IsUse == true && pClient.Member.Status == 1 {

			// 讀取玩家的會員資料
			pMember := &pClient.Member // 取址

			// 檢查錢包, 想帶的錢是否足夠
			if pMember.Balance-enter_Game.Balance_ci < 0 {
				Code = int(ERROR_CODE_CARRY_BALANCE_NOT_ENOUGH)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 找到對應的LobbyInfo
			lobbyinfo := Common_LobbyInfoGet2(enter_Game.PlatformID, enter_Game.GameID)
			if lobbyinfo.LobbyID == 0 {
				Code = int(ERROR_CODE_ERROR_DATA)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, LobbyID=%d,LobbyName=%s", Code, ErrorCode[Code].Message, lobbyinfo.LobbyID, lobbyinfo.LobbyName)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 找到對應的GameInfo
			gameinfo := Common_GameInfoGet(enter_Game.PlatformID, enter_Game.GameID)
			if gameinfo.GameID == 0 {
				Code = int(ERROR_CODE_ERROR_DATA)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, GameID=%d,GameName=%s", Code, ErrorCode[Code].Message, gameinfo.GameID, gameinfo.GameName)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 找尋可以加入的桌子 (簡單,但是桌子數量多就效能慢)
			pTable, bfind := Match_TableInfoTableJoinSeatch(enter_Game.PlatformID, enter_Game.GameID, *pMember)
			if bfind == true {

				CommonLog_INFO_Printf("#Common_EnterGame(有找到可入桌的舊桌子) PlatformID=%d, LobbyID=%d, GameID=%d, TableID=%s, TableArrayIdx=%d, User_ID=%d, Account=%s",
					pTable.PlatformID, pTable.LobbyID, pTable.GameID, pTable.TableID, pTable.TableArrayIdx, pMember.User_ID, pMember.Account)

				// 有找到可入桌的舊桌子, 加入此桌( 也有攜帶錢入桌 )
				DataMsg, Code, Seat_ID = Match_JoinTable(ClientID, lobbyinfo, gameinfo, pMember, pTable, "", enter_Game.Balance_ci)
				if Code != 0 {
					Code = int(ERROR_CODE_ERROR_JOIN_TABLE)
					CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, GameID=%d,GameName=%s", Code, ErrorCode[Code].Message, gameinfo.GameID, gameinfo.GameName)
					return DataMsg, Code, BroadcastTable, BroadcastMember
				}

				BroadcastTable = *pTable   // 拷貝廣播資訊
				BroadcastMember = *pMember // 取值(只是拿 Member.ClientID, 來廣播而已 )
				TableID = pTable.TableID
				// 設定廣播給其它同桌玩家  "re"  or "fw"
				//Match_Broadcast(NET_CMD_ENTER_GAME, *pTable, *pMember)

				// 組合回傳字串

			} else {

				// 另開新桌
				pNewTable := Match_OpenTable(ClientID, lobbyinfo, gameinfo)
				if pNewTable == nil {
					Code = int(ERROR_CODE_ERROR_OPEN_TATBL)
					CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, GameID=%d,GameName=%s, 桌子使用率:[%d/%d]", Code, ErrorCode[Code].Message, gameinfo.GameID, gameinfo.GameName, TableInfoCount, TABLEINFO_MAX)
					return DataMsg, Code, BroadcastTable, BroadcastMember
				}

				CommonLog_INFO_Printf("#Common_EnterGame(另開新桌) PlatformID=%d, LobbyID=%d, GameID=%d, TableID=%s, TableArrayIdx=%d, User_ID=%d, Account=%s",
					pNewTable.PlatformID, pNewTable.LobbyID, pNewTable.GameID, pNewTable.TableID, pNewTable.TableArrayIdx, pMember.User_ID, pMember.Account)

				// 加入此桌
				DataMsg, Code, Seat_ID = Match_JoinTable(ClientID, lobbyinfo, gameinfo, pMember, pNewTable, pNewTable.TableID, enter_Game.Balance_ci)
				if Code != 0 {
					Code = int(ERROR_CODE_ERROR_JOIN_TABLE)
					CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, GameID=%d,GameName=%s", Code, ErrorCode[Code].Message, gameinfo.GameID, gameinfo.GameName)
					return DataMsg, Code, BroadcastTable, BroadcastMember
				}

				BroadcastTable = *pNewTable // 拷貝廣播資訊
				BroadcastMember = *pMember  // 取值(只是拿 Member.ClientID, 來廣播而已 )
				TableID = pNewTable.TableID
				// 組合回傳字串
			}

			// 組合回傳字串
			fishEnterGame := ResponseInfo_Fish_Enter_Game{}
			fishEnterGame.Status = int(ERROR_CODE_SUCCESS)
			fishEnterGame.Shoot_gap = 0.33
			fishEnterGame.Dlg_show_chance = 0.2
			fishEnterGame.Dlg_hit_chance = 0.2
			fishEnterGame.TableID = TableID
			fishEnterGame.User_ID = pMember.User_ID
			fishEnterGame.Nickname = pMember.NickName // 等我有暱稱在放在去
			fishEnterGame.Background = 0
			fishEnterGame.Jp_prize = 12345678
			fishEnterGame.Jp_bet = 5
			fishEnterGame.Bet_options[0] = 1
			fishEnterGame.Bet_options[1] = 2
			fishEnterGame.Bet_options[2] = 3
			fishEnterGame.Bet_options[3] = 5
			fishEnterGame.Bet_options[4] = 7
			fishEnterGame.Balance = enter_Game.Balance_ci //玩家帶進來的錢 //pMember.Balance_ci + pMember.Balance_wi //在調整

			for i := 0; i < SEAT_MAX; i++ {
				seatTmp := BroadcastTable.SeatInfo[i]
				if seatTmp.bUse == false {
					continue
				}

				fishEnterGame.Table_player_info[i].Seat_ID = seatTmp.Seat_ID
				fishEnterGame.Table_player_info[i].User_ID = seatTmp.User_ID
				fishEnterGame.Table_player_info[i].NickName = seatTmp.NickName
				fishEnterGame.Table_player_info[i].Balance = enter_Game.Balance_ci //玩家帶進來的錢 //pMember.Balance_ci + pMember.Balance_wi //在調整
				fishEnterGame.Table_player_info[i].Bullet_type = "1"
				fishEnterGame.Table_player_info[i].Last_bet = 10
			}

			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(fishEnterGame)
			if err != nil {
				CommonLog_WARNING_Println("json err:", err)
			}
			DataMsg = string(DataMsgByte)

		} else {
			Code = int(ERROR_CODE_NO_LOGIN)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, Seat_ID=%d, DataMsg=%s", Code, ErrorCode[Code].Message, Seat_ID, DataMsg)
	return DataMsg, Code, BroadcastTable, BroadcastMember
}

//=========================================================================================================================================================================
// 加入遊戲
func Common_JoinGame(ClientID int, DecodeData string) (string, int, TableInfo, MemberInfo) {
	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值
	var BroadcastTable TableInfo           // 要廣播的桌子
	var BroadcastMember MemberInfo         // 發起廣播的玩家
	var Seat_ID int = 0
	CommonLog_INFO_Printf("#收到封包 Common_JoinGame ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	join_Game := PacketCmd_Join_Game{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &join_Game)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code, BroadcastTable, BroadcastMember
	} else {

		CommonLog_INFO_Printf("#TableID=%s, TableArrayIdx=%d, User_ID=%d, Balance_ci=%f", join_Game.TableID, join_Game.TableArrayIdx, join_Game.User_ID, join_Game.Balance_ci)

		// 讀取玩家的socket資料
		// 檢查玩家是否有登入
		pClient := Common_ClientInfoGet(ClientID)
		if pClient.IsUse == true && pClient.Member.Status == 1 {

			// 讀取玩家的會員資料
			pMember := &pClient.Member // 取址

			// 找玩家錢是否夠玩 (想帶進來的錢跟身上的錢去比較)
			// 檢查錢包, 想帶的錢是否足夠
			if pMember.Balance-join_Game.Balance_ci < 0 {
				Code = int(ERROR_CODE_CARRY_BALANCE_NOT_ENOUGH)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 檢查桌號是否存在
			pTable := Match_TableInfoGet(join_Game.TableArrayIdx)
			if pTable.TableID != join_Game.TableID {
				Code = int(ERROR_CODE_NO_FIND_TABLE)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 檢查桌子使否使用中
			if pTable.bUse == false {
				Code = int(ERROR_CODE_NO_USE)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}
			CommonLog_INFO_Printf("有找到 TableArrayIdx=%d, PlatformID=%d, LobbyID=%d, GameID=%d, TableID=%s, 人數=[%d/%d]",
				join_Game.TableArrayIdx, pTable.PlatformID, pTable.LobbyID, pTable.GameID, pTable.TableID, pTable.TablePlayerNow, pTable.TablePlayerMax)

			// 取得LobbyInfo
			lobbyinfo := Common_LobbyInfoGet2(pTable.PlatformID, pTable.GameID)
			if lobbyinfo.LobbyID == 0 {
				Code = int(ERROR_CODE_ERROR_DATA)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, LobbyID=%d,LobbyName=%s", Code, ErrorCode[Code].Message, lobbyinfo.LobbyID, lobbyinfo.LobbyName)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 找到對應的GameInfo
			gameinfo := Common_GameInfoGet(pTable.PlatformID, pTable.GameID)
			if gameinfo.GameID == 0 {
				Code = int(ERROR_CODE_ERROR_DATA)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, GameID=%d,GameName=%s", Code, ErrorCode[Code].Message, gameinfo.GameID, gameinfo.GameName)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 加入 User_ID 檢查
			if join_Game.User_ID != pMember.User_ID {
				Code = int(ERROR_CODE_ERROR_USER_ID)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 加入桌
			DataMsg, Code, Seat_ID = Match_JoinTable(ClientID, lobbyinfo, gameinfo, pMember, pTable, pTable.TableID, join_Game.Balance_ci)
			if Code != 0 {
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s,Account=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.Account, pMember.User_ID)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			CommonLog_INFO_Printf("DataMsg=%s", DataMsg)

			// 設定廣播給其它同桌玩家  "re"  or "fw"
			// 組合回傳字串
			BroadcastTable = *pTable   // 拷貝廣播資訊
			BroadcastMember = *pMember // 取值(只是拿 Member.ClientID, 來廣播而已 )

			seat := pTable.SeatInfo[Seat_ID]
			fishJoinGame := ResponseInfo_Fish_Join_Game{}
			fishJoinGame.Status = int(ERROR_CODE_SUCCESS)
			fishJoinGame.TableID = pTable.TableID
			fishJoinGame.TableArrayIdx = pTable.TableArrayIdx
			fishJoinGame.User_ID = seat.User_ID
			fishJoinGame.Seat_ID = seat.Seat_ID
			fishJoinGame.NickName = seat.NickName
			fishJoinGame.Balance_ci = join_Game.Balance_ci

			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(fishJoinGame)
			if err != nil {
				CommonLog_WARNING_Println("json err:", err)
			}
			DataMsg = string(DataMsgByte)

		} else {
			Code = int(ERROR_CODE_NO_LOGIN)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", ErrorCode[Code].Code, ErrorCode[Code].Message, DataMsg)
	return DataMsg, Code, BroadcastTable, BroadcastMember
}

//=========================================================================================================================================================================
// 離開遊戲 (稍後優化 整進Match系統內)
func Common_ExitGame(ClientID int, DecodeData string) (string, int, TableInfo, MemberInfo) {
	var DataMsg string = "unknow"
	var Code int = int(ERROR_CODE_SUCCESS) // 回應值
	var BroadcastTable TableInfo           // 要廣播的桌子
	var BroadcastMember MemberInfo         // 發起廣播的玩家
	var seat_backup SeatInfo               // 備份的位子資訊
	CommonLog_INFO_Printf("#收到封包 Common_ExitGame ClientID=%d, DecodeData=%s", ClientID, DecodeData)

	// 收到的資料 json轉換
	receiveMsgByte := []byte(DecodeData)
	exit_Game := PacketCmd_Exit_Game{} // 用來接的物件
	err := json.Unmarshal(receiveMsgByte, &exit_Game)
	if err != nil {
		CommonLog_WARNING_Printf("錯誤的json格式 DecodeData=%s", DecodeData)
		Code = int(ERROR_CODE_ERROR_JSON_MARSHAL)
		CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
		return DataMsg, Code, BroadcastTable, BroadcastMember
	} else {

		CommonLog_INFO_Printf("#TableID=%s, User_ID=%d, Seat_ID=%d", exit_Game.TableID, exit_Game.User_ID, exit_Game.Seat_ID)

		// 讀取玩家的socket資料
		pClient := Common_ClientInfoGet(ClientID)

		// 讀取玩家的會員資料
		pMember := &pClient.Member

		// 檢查玩家是否有登入
		if pClient.IsUse == true && pClient.Member.Status == 1 {

			// 加入 User_ID 檢查
			if exit_Game.User_ID != pMember.User_ID {
				Code = int(ERROR_CODE_ERROR_USER_ID)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s", Code, ErrorCode[Code].Message)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 檢查桌號是否存在
			CommonLog_INFO_Printf("檢查桌號是否存在 TableID=%s, TableArrayIdx=%d", exit_Game.TableID, pMember.TableArrayIdx)
			pTable := Match_TableInfoGet(pMember.TableArrayIdx)
			if pTable.TableID != exit_Game.TableID {
				Code = int(ERROR_CODE_NO_FIND_TABLE)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			// 檢查玩家是否在桌內
			var bFind bool = false
			var pSeat *SeatInfo
			//var SeatIdx int = 0
			for i := 0; i < SEAT_MAX; i++ {
				pSeat = &pTable.SeatInfo[i]

				CommonLog_INFO_Printf("i=%d, User_ID=%d, Account=%s, Seat_ID=%d", i, pSeat.User_ID, pSeat.Account, pSeat.Seat_ID)
				if pSeat.bUse == false {
					continue
				}

				if pSeat.User_ID == exit_Game.User_ID {
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

			CommonLog_INFO_Printf("#離開遊戲 TableID=%s, User_ID=%d, Account=%s", exit_Game.TableID, pSeat.User_ID, pSeat.Account)

			// 離開遊戲 設定
			BroadcastTable = *pTable   // 拷貝廣播資訊
			BroadcastMember = *pMember // 取值(只是拿 Member.ClientID, 來廣播而已 )
			seat_backup = *pSeat       // 備份的位子資訊

			// 刪除座位資訊DB
			Mysql_CommonSeatInfo_Delete(pTable.PlatformID, pTable.LobbyID, pTable.GameID, pTable.TableID, pSeat.Seat_ID, pSeat.User_ID)

			//==========================================================================================================================================
			// 攜帶的錢返回 memberinfo 的 balance ( Balance_ci + Balance_win )
			//@@@
			// 重新讀取DB 玩家的錢
			var before_Balance int64 = pMember.Balance
			var after_Balance int64 = 0
			var table_Balance int64 = pSeat.Balance_ci + pSeat.Balance_win // 桌內目前的錢
			ret, LoadMember := Mysql_CommonMemberInfoGet(pMember.Account, pMember.Password)
			if ret == true {

				// 2.寫入DB ( 用DB的錢 - Balance_ci )
				ret = Mysql_CommonMemberInfo_AddBalance(LoadMember, table_Balance)
				if ret == false {
					Code = int(ERROR_CODE_BALANCE_UPDATE_FAIL)
					CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
					return DataMsg, Code, BroadcastTable, BroadcastMember
				}

				// 計算預期的金額
				after_Balance = LoadMember.Balance + table_Balance // 拿DB內的資料去筆對

				ret, LoadMember2 := Mysql_CommonMemberInfoGet(pMember.Account, pMember.Password)
				if ret == true {

					// 比對錢是否正確
					if LoadMember2.Balance != after_Balance {
						// 發生錯誤了
						CommonLog_INFO_Printf("#Common_ExitGame(離開遊戲) 警告! 錢對不起來, ClientID=%d, TableID=%s, Account=%s, User_ID=%d, table_Balance(桌內目前的錢)=%d Balance[%d!=%d]",
							ClientID, pTable.TableID, pMember.Account, pMember.User_ID, table_Balance, LoadMember2.Balance, after_Balance)

						Code = int(ERROR_CODE_BALANCE_UPDATE_FAIL)
						CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
						return DataMsg, Code, BroadcastTable, BroadcastMember
					} else {

						// 更新玩家的錢
						pMember.Balance = after_Balance

						pMember.TableArrayIdx = pTable.TableArrayIdx // member 存一份幫助搜尋
						CommonLog_INFO_Printf("#Common_ExitGame(離開遊戲) 快速搜尋 ClientID=%d, TableID=%s, TableArrayIdx=%d, Account=%s, User_ID=%d, table_Balance(桌內目前的錢)=%d, before_Balance=%d, after_Balance=%d",
							ClientID, pTable.TableID, pMember.TableArrayIdx, pMember.Account, pMember.User_ID, table_Balance, before_Balance, pMember.Balance)
					}

				}

			} else {
				Code = int(ERROR_CODE_NO_FIND_ACCOUNT)
				CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, TableID=%s, User_ID=%d", Code, ErrorCode[Code].Message, pTable.TableID, pMember.User_ID)
				return DataMsg, Code, BroadcastTable, BroadcastMember
			}

			//===========================================================================================================================================

			// 寫入GameLog

			// 刪除位子記憶體資料
			*pSeat = SeatInfo{}

			// 將玩家資料關連跟桌子關連解除
			pMember.TableArrayIdx = -1

			// 桌子人數減1
			pTable.TablePlayerNow--
			Mysql_CommonTableInfo_Update(*pTable)
			CommonLog_INFO_Printf("#Common_ExitGame目前桌子人數 TableID=%s, TablePlayerNow=%d", pTable.TableID, pTable.TablePlayerNow)

			if pTable.TableDestoryMode == 1 {

				CommonLog_INFO_Printf("#Common_ExitGame檢查桌內人數 TableID=%s, TableDestoryMode=%d, 人數[%d/%d]", pTable.TableID, pTable.TableDestoryMode, pTable.TablePlayerNow, pTable.TablePlayerMax)

				// 判斷人數
				if pTable.TablePlayerNow <= 0 {

					// 刪除桌資訊DB  PlatformID int, LobbyID int, GameID int, TableID string, TableArrayIdx int
					Mysql_CommonTableInfo_Delete(pTable.PlatformID, pTable.LobbyID, pTable.GameID, pTable.TableID, pTable.TableArrayIdx)

					// 桌子數量減少
					TableInfoCount--
					CommonLog_INFO_Printf("#Common_ExitGame(散桌後刪除此桌資訊) TableInfoCount=%d, TableID=%s, TableDestoryMode=%d, 人數[%d/%d]", TableInfoCount, pTable.TableID, pTable.TableDestoryMode, pTable.TablePlayerNow, pTable.TablePlayerMax)

					// 刪除桌資訊記憶體資料
					pTable.bUse = false
					*pTable = TableInfo{}

				} else {
					CommonLog_INFO_Printf("#Common_ExitGame(還有人在桌內), 所以不用刪除桌 TableID=%s, 人數[%d/%d]", pTable.TableID, pTable.TablePlayerNow, pTable.TablePlayerMax)
				}

			} else {
				CommonLog_INFO_Printf("#Common_ExitGame(散桌後保留此桌資訊),等待玩家重新入桌 TableInfoCount=%d, TableID=%s, TableDestoryMode=%d", TableInfoCount, pTable.TableID, pTable.TableDestoryMode)
			}

			// 除錯用
			Mysql_CommonTableInfoShow()

			// 設定廣播給其它同桌玩家  "re"  or "fw"
			// 組合回傳字串
			fishExitGame := ResponseInfo_Fish_Exit_Game{}
			fishExitGame.Status = int(ERROR_CODE_SUCCESS)
			fishExitGame.TableID = BroadcastTable.TableID
			fishExitGame.Seat_ID = seat_backup.Seat_ID
			fishExitGame.User_ID = seat_backup.User_ID
			fishExitGame.NickName = seat_backup.NickName

			// 物件轉成json字串
			DataMsgByte, err := json.Marshal(fishExitGame)
			if err != nil {
				CommonLog_WARNING_Println("json err:", err)
			}
			DataMsg = string(DataMsgByte)

		} else {
			Code = int(ERROR_CODE_NO_LOGIN)
		}
	}

	CommonLog_INFO_Printf("#Code=%d, errorMsg=%s, DataMsg=%s", Code, ErrorCode[Code].Message, DataMsg)
	return DataMsg, Code, BroadcastTable, BroadcastMember
}

//=========================================================================================================================================================================
// 廣播給其他同桌玩家 ( 只廣播 fw 的cmd, re自己的透過一般管道回覆)
func Message_Broadcast(Cmd string, Table TableInfo, BroadcastMember MemberInfo, BroadcastMessage string) (string, int) {

	var Code int = int(ERROR_CODE_SUCCESS) // 回應值
	var ResponseDataMsg string = "unknow"
	var Response CommonResponseInfo // 共用回應結構

	CommonLog_INFO_Printf("#Message_Broadcast(廣播給其他同桌玩家)-開始 Cmd=%s, TableID=%s, (發起廣播的人)Account=%s, User_ID=%d", Cmd, Table.TableID, BroadcastMember.Account, BroadcastMember.User_ID)

	for i := 0; i < SEAT_MAX; i++ {
		seat := Table.SeatInfo[i]
		ClientID := seat.ClientID

		if seat.bUse == false {
			continue
		}

		CommonLog_INFO_Printf("ClientID=%d, TableID=%s, Account=%s, User_ID=%d", ClientID, Table.TableID, seat.Account, seat.User_ID)

		// 檢查玩家是否有登入
		pClient := Common_ClientInfoGet(ClientID)
		Member := pClient.Member
		if pClient.IsUse == false || Member.Status == 0 {

			Code = int(ERROR_CODE_NO_LOGIN)
			CommonLog_WARNING_Printf("警告 玩家沒登入 ClientID=%d, TableID=%s, Account=%s, User_ID=%d", ClientID, Table.TableID, seat.Account, seat.User_ID)
			break
		}

		if seat.User_ID == BroadcastMember.User_ID {
			ResponseObj := CommonResponseInfo_Reply{}
			ResponseObj.Reply = BroadcastMessage

			DataMsgByte, err := json.Marshal(ResponseObj)
			if err != nil {
				Code = int(ERROR_CODE_ERROR_DATA)
				CommonLog_INFO_Println("json2 err:", err)
				break
			} else {

				ResponseDataMsg = string(DataMsgByte)
			}
			CommonLog_INFO_Printf("自己是廣播者, 所以遵循一般管道回去 (發起廣播的人)Account=%s, User_ID=%d", BroadcastMember.Account, BroadcastMember.User_ID)

		} else {

			// 廣播給其他玩家

			// cmd 拷貝一份 回傳時候可以使用
			Response.Ret = Cmd
			Response.Sys = "game"
			Response.ClientID = ClientID
			pClient.Sn++
			Response.Sn = pClient.Sn

			// 預設的回應訊息
			Response.Code = 0
			Response.Message = "正確執行"

			ResponseObj := CommonResponseInfo_Forward{}
			ResponseObj.Forward = BroadcastMessage

			DataMsgByte, err := json.Marshal(ResponseObj)
			if err != nil {
				Code = int(ERROR_CODE_ERROR_DATA)
				CommonLog_INFO_Println("json2 err:", err)
				break
			} else {

				// 送廣播封包嚕
				//var responseMsg string = string(DataMsgByte)
				Response.Data = string(DataMsgByte)

				//------------------------------------------------
				// 組成回應格式
				DataMsgByte2, err := json.Marshal(Response)
				if err != nil {
					CommonLog_WARNING_Println("json2 err:", err)
				}
				var responseMsg string = string(DataMsgByte2)

				Common_ClientSendMessage(ClientID, responseMsg)
				CommonLog_INFO_Printf("送廣播封包嚕 Cmd=%s, TableID=%s, (被廣播的人)Account=%s, User_ID=%d", Cmd, Table.TableID, Member.Account, Member.User_ID)
			}

		}
	}

	CommonLog_INFO_Printf("#Message_Broadcast(廣播給其他同桌玩家)-結束 Cmd=%s, TableID=%s, (發起廣播的人)Account=%s, User_ID=%d", Cmd, Table.TableID, BroadcastMember.Account, BroadcastMember.User_ID)
	return ResponseDataMsg, Code
}

//=========================================================================================================================================================================
// 桌子Process
func Table_Process() {

	for {

		var Code int = int(ERROR_CODE_SUCCESS)   // 回應值
		var ResponseFinishData string = "unknow" // 回應的資料格式(加工後的最終版)
		var ResponseTmpData string = "unknow"    // 回應的資料格(暫時)
		var BroadcastTable TableInfo             // 要廣播的桌子
		var BroadcastMember MemberInfo           // 發起廣播的玩家
		var NeedBroadcast bool = false           // 需要廣播時候
		var Cmd string = ""                      // 廣播CMD
		var Count int = 0

		// 掃描整個array
		for i := 0; i < TABLEINFO_MAX; i++ {
			pTable := Match_TableInfoGet(i)

			NeedBroadcast = false
			if pTable.bUse == false {
				// 沒使用的桌子
				continue
			}
			if Count > TableInfoCount {
				//
				break
			}

			// 依序分配桌子
			switch pTable.GameMode {
			case int8(GAME_MODE_FISH): // 魚機  NeedBroadcast, Cmd, DataMsg, Code, BroadcastTable, BroadcastMember
				NeedBroadcast, Cmd, ResponseTmpData, Code, BroadcastTable, BroadcastMember = Fish_Process(pTable)
			case int8(GAME_MODE_SLOT): // Slot
			case int8(GAME_MODE_POKER): // 撲克牌
			case int8(GAME_MODE_MAHJONG): // 麻將

			default:
				CommonLog_WARNING_Printf("#Table_Process(廣播給其他同桌玩家) warning 未處理的 i=%d, PlatformID=%d, GameID=%d, GameMode=%d, TableID=%s",
					i, pTable.PlatformID, pTable.GameID, pTable.GameMode, pTable.TableID)
			}

			// 找到有使用的桌子
			Count++

			// 廣播
			if NeedBroadcast == true {
				CommonLog_INFO_Printf("#Table_Process(廣播給其他同桌玩家) PlatformID=%d, GameID=%d, GameMode=%d, TableID=%s, 人數=[%d/%d]",
					pTable.PlatformID, pTable.GameID, pTable.GameMode, pTable.TableID, pTable.TablePlayerNow, pTable.TablePlayerMax)
				ResponseFinishData, Code = Message_Broadcast(Cmd, BroadcastTable, BroadcastMember, ResponseTmpData)

				CommonLog_INFO_Printf("#Table_Process(廣播給其他同桌玩家) TableID=%s, Code=%d, ResponseFinishData=%s", pTable.TableID, Code, ResponseFinishData)
			}

		}

		//time.Sleep(1)
		time.Sleep(1 * time.Second)
	}
}

//=========================================================================================================================================================================
//=========================================================================================================================================================================
//=========================================================================================================================================================================
// 遊戲的CMD
func Common_DispatchGame(ClientID int, Cmd string, DecodeData string) (string, int) {

	var Code int = int(ERROR_CODE_SUCCESS)   // 回應值
	var ResponseFinishData string = "unknow" // 回應的資料格式(加工後的最終版)
	var ResponseTmpData string = "unknow"    // 回應的資料格(暫時)
	var BroadcastTable TableInfo             // 要廣播的桌子
	var BroadcastMember MemberInfo           // 發起廣播的玩家

	CommonLog_INFO_Printf("#Common_DispatchGame(遊戲的CMD) ClientID=%d, Cmd=%s, DecodeData=%s", ClientID, Cmd, DecodeData)

	// 後面按照 GameMode 來分遊戲 再分配一次

	// 分析 cmd 並且分派和處理
	switch Cmd {

	case NET_CMD_ENTER_GAME:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_ENTER_GAME")

			// 進入遊戲
			ResponseTmpData, Code, BroadcastTable, BroadcastMember = Common_EnterGame(ClientID, DecodeData)
			CommonLog_INFO_Printf("#進入遊戲-結果 ResponseTmpData:%s", ResponseTmpData)
		}

	case NET_CMD_JOIN_GAME:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_JOIN_GAME")

			// 加入遊戲
			ResponseTmpData, Code, BroadcastTable, BroadcastMember = Common_JoinGame(ClientID, DecodeData)
			CommonLog_INFO_Printf("#加入遊戲-結果 ResponseTmpData:%s", ResponseTmpData)
		}

	case NET_CMD_EXIT_GAME:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_EXIT_GAME")

			// 離開遊戲
			ResponseTmpData, Code, BroadcastTable, BroadcastMember = Common_ExitGame(ClientID, DecodeData)
			CommonLog_INFO_Printf("#離開遊戲-結果 ResponseTmpData:%s", ResponseTmpData)
		}

	case NET_CMD_FISH_SHOOT:
		{
			CommonLog_INFO_Printf("#收到封包 CMD=NET_CMD_FISH_SHOOT")

			// 魚機-射擊
			ResponseTmpData, Code, BroadcastTable, BroadcastMember = Fish_Shoot(ClientID, DecodeData)
			CommonLog_INFO_Printf("#魚機射擊-結果 ResponseTmpData:%s", ResponseTmpData)
		}

	default:
		CommonLog_WARNING_Printf("warning 未處理的cmd=%s", Cmd)
		Code = int(ERROR_CODE_NO_FIND_CMD)
		ResponseFinishData = ""
		//break
	}

	// 廣播
	if Code == int(ERROR_CODE_SUCCESS) {
		ResponseFinishData, Code = Message_Broadcast(Cmd, BroadcastTable, BroadcastMember, ResponseTmpData)
	}

	return ResponseFinishData, Code
}
