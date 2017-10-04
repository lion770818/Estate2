package One1CloudLib

//"fmt"
//"log"
//"time"

// 大廳 結構
type LobbyInfo struct {
	PlatformID int `json:"platform_id"` //第三方平台編號
	LobbyID    int `json:"lobby_id"`    //大廳規則編號
	GameID     int `json:"game_id"`     //遊戲編號

	LobbyName    string `json:"lobby_name"`     // 大聽中文名稱
	LobbyMatchID int8   `json:"lobby_match_id"` // 大廳配桌編號, 編號相同的廳館, 才可以配桌在一起   ( LobbyMatchID 相同  GameID 相同 )

	Total_water1 float64 `json:"total_water1"` // 總水池1
	Total_water2 float64 `json:"total_water2"` // 總水池2
	BetLevel     float64 `json:"bet_level"`    // 單一押注金額
}

var (
	LobbyInfoList = make(map[int]LobbyInfo) // map LobbyInfo
)

//=========================================================================================================================================================================
// 抓取 LobbyInfo
func Common_LobbyInfoGet2(PlatformID int, GameID int) LobbyInfo {

	var lobbyInfo LobbyInfo
	CommonLog_INFO_Printf("#Common_LobbyInfoGet2 取得一個 LobbyInfo 的記憶體位置... PlatformID=%d, GameID=%d", PlatformID, GameID)

	// 取出記憶體位置
	for i := 0; i < len(LobbyInfoList); i++ {
		lobbyInfoTmp := LobbyInfoList[i]
		if lobbyInfoTmp.PlatformID == PlatformID && lobbyInfoTmp.GameID == GameID {
			CommonLog_INFO_Printf("有找到 i=%d, PlatformID=%d, GameID=%d", i, PlatformID, GameID)
			lobbyInfo = lobbyInfoTmp
			break
		}
	}
	CommonLog_INFO_Printf("#Common_LobbyInfoGet2 取得一個 LobbyInfo 的記憶體位置... PlatformID=%d, GameID=%d, LobbyName=%s, LobbyID=%d, BetLevel=%f, Total_water=[%f,%f]",
		PlatformID, GameID, lobbyInfo.LobbyName, lobbyInfo.LobbyID, lobbyInfo.BetLevel, lobbyInfo.Total_water1, lobbyInfo.Total_water2)
	return lobbyInfo
}
