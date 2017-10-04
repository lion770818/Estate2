package One1CloudLib

import (
	"fmt"
	//"log"
	//"math/rand"
	//"time"

	//"database/sql"

	_ "github.com/go-sql-driver/mysql"
	//"mysql-master"
	//"encoding/json"
)

//=========================================================================================================================================================================
// 位子新增 (這邊嘗試過 用? 去組指令, 但是出錯時候 列印指令拿到一堆 ??? 所以還是用傳統的 組好mysql字串, 方便除錯 )
func Mysql_GameLosFish_Insert(gamelog GameLos_Fish) bool {
	
	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_GameLosFish_Insert (位子新增)")

	SqlQuery = fmt.Sprintf("INSERT INTO gamelog_fish(PlatformID, LobbyID, GameID, TableID, Seat_ID, GameMode, CreateTime, User_ID, Account, NickName, Round, Before_Balance_ci, Before_Balance_win, Balance_ci, Balance_win, BetLevel, Bet_Win, Process_Status ) values(%d,%d,%d,'%s',%d,%d,'%s',%d,'%s','%s',%d,%f,%f,%f,%f,%f,%f,%d)",
		gamelog.PlatformID, gamelog.LobbyID, gamelog.GameID, gamelog.TableID, gamelog.Seat_ID, gamelog.GameMode,
		Common_NowTimeGet(),
		gamelog.User_ID, gamelog.Account, gamelog.NickName,
		gamelog.Round, 
		gamelog.Before_Balance_ci, gamelog.Before_Balance_win,
		gamelog.Balance_ci, gamelog.Balance_win, 
		gamelog.BetLevel,
		gamelog.Bet_Win, gamelog.Process_Status)

	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)

	result, err := database.Exec(SqlQuery)
	if err != nil {
		ErrorStr := err.Error()
		fmt.Println(ErrorStr)
		//panic(err.Error()) // proper error handling instead of panic in your app
		ret = false
	} else {
		ret = true

		num, err := result.RowsAffected()
		if err != nil {
			CommonLog_INFO_Printf("fetch row affected failed:", err.Error())
			ret = false
		}

		if num > 0 {
			ret = true
		} else {
			ret = false
		}

		CommonLog_INFO_Printf("DB資料變動數 record number=%d ret=%v", num, ret)
	}
	
	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}