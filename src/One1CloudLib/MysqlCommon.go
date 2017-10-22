package One1CloudLib

import (
	"fmt"
	"log"

	//"math/rand"
	//"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	//"mysql-master"
	//"encoding/json"
)

var database = &sql.DB{}

func Mysql_Init() {

	var err error

	CommonLog_INFO_Printf("#Mysql_Init")

	mysql_config := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", ServerConfig.Mysql_Account, ServerConfig.Mysql_Password, ServerConfig.Mysql_IP, ServerConfig.Mysql_DB)
	CommonLog_INFO_Printf("#Mysql_Init mysql_config=%s, Mysql_SetMaxIdleConns=%d, Mysql_SetMaxOpenConns=%d", mysql_config, ServerConfig.Mysql_SetMaxIdleConns, ServerConfig.Mysql_SetMaxOpenConns)
	//database, err = sql.Open("mysql", "root:1234@/one1cloud")						 	// 帳號:root 密碼:1234
	//database, err = sql.Open("mysql", "root:1234@tcp(localhost:3306)/one1cloud") 		//指定IP和端口
	//database, err = sql.Open("mysql", "root:1234@tcp(192.168.1.21:3306)/one1cloud") 	//指定IP和端口
	database, err = sql.Open("mysql", mysql_config)
	database.SetMaxOpenConns(ServerConfig.Mysql_SetMaxOpenConns)
	database.SetMaxIdleConns(ServerConfig.Mysql_SetMaxIdleConns)

	if err != nil {
		ErrorStr := err.Error()
		CommonLog_ERROR_Fatal(ErrorStr)
		panic(err.Error())
	}

	//defer database.Close()
}

//=========================================================================================================================================================================
// mysql 取得大廳資訊
// PlatformID 第三方平台編號
// GameID 遊戲編號
func Mysql_CommonLobbyInfoGet(PlatformID int, GameID int) (bool, LobbyInfo) {

	var ret bool = false
	var SqlQuery string = ""

	CommonLog_INFO_Printf("#Mysql_CommonLobbyInfoGet (取得大廳資訊) PlatformID=%d, GameID=%d", PlatformID, GameID)

	SqlQuery = fmt.Sprintf("SELECT PlatformID, LobbyID, GameID, LobbyName, LobbyMatchID Total_water1, Total_water2, BetLevel FROM lobbyinfo where PlatformID=%d AND GameID=%d;", PlatformID, GameID)
	log.Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	lobbyinfo := LobbyInfo{} // 用來接的物件
	for stmtIns.Next() {

		if err2 := stmtIns.Scan(&lobbyinfo.PlatformID, &lobbyinfo.LobbyID, &lobbyinfo.GameID, &lobbyinfo.LobbyName, &lobbyinfo.LobbyMatchID, &lobbyinfo.Total_water1, &lobbyinfo.Total_water2, &lobbyinfo.BetLevel); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("PlatformID=%d is GameID=%d, LobbyID=%d LobbyName=%s BetLevel=%d", lobbyinfo.PlatformID, lobbyinfo.GameID, lobbyinfo.LobbyID, lobbyinfo.LobbyName, lobbyinfo.BetLevel)
	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	return ret, lobbyinfo
}

//=========================================================================================================================================================================
// mysql 取得大廳資訊
// PlatformID 第三方平台編號
func Mysql_CommonLobbyInfoGet2(PlatformID int) (map[int]LobbyInfo, bool) {

	var ret bool = false
	var (
		LobbyInfoListTmp = make(map[int]LobbyInfo) // map LobbyInfo
	)

	CommonLog_INFO_Printf("#Mysql_CommonLobbyInfoGet2 (取得大廳資訊) PlatformID=%d", PlatformID)

	var SqlQuery string = ""
	SqlQuery = fmt.Sprintf("SELECT PlatformID, LobbyID, GameID, LobbyName, LobbyMatchID, Total_water1, Total_water2, BetLevel FROM lobbyinfo  where PlatformID=%d", PlatformID)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	var Count int = 0
	for stmtIns.Next() {

		lobbyinfo := LobbyInfo{} // 用來接的物件
		if err2 := stmtIns.Scan(&lobbyinfo.PlatformID, &lobbyinfo.LobbyID, &lobbyinfo.GameID, &lobbyinfo.LobbyName, &lobbyinfo.LobbyMatchID, &lobbyinfo.Total_water1, &lobbyinfo.Total_water2, &lobbyinfo.BetLevel); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("PlatformID=%d is GameID=%d, LobbyID=%d LobbyName=%s BetLevel=%d", lobbyinfo.PlatformID, lobbyinfo.GameID, lobbyinfo.LobbyID, lobbyinfo.LobbyName, lobbyinfo.BetLevel)
		LobbyInfoListTmp[Count] = lobbyinfo
		Count++

	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	return LobbyInfoListTmp, ret
}

//=========================================================================================================================================================================
// mysql 取得大廳資訊
// PlatformID 第三方平台編號
// GameID 遊戲編號
func Mysql_CommonLobbyInfoGetAll() bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonLobbyInfoGetAll (取得大廳資訊)")

	SqlQuery = fmt.Sprintf("SELECT PlatformID, LobbyID, GameID, LobbyName, LobbyMatchID, Total_water1, Total_water2, BetLevel FROM lobbyinfo")
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	var Count int = 0
	for stmtIns.Next() {

		lobbyinfo := LobbyInfo{} // 用來接的物件
		if err2 := stmtIns.Scan(&lobbyinfo.PlatformID, &lobbyinfo.LobbyID, &lobbyinfo.GameID, &lobbyinfo.LobbyName, &lobbyinfo.LobbyMatchID, &lobbyinfo.Total_water1, &lobbyinfo.Total_water2, &lobbyinfo.BetLevel); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("PlatformID=%d is LobbyID=%d, GameID=%d, LobbyName=%s, LobbyMatchID=%d, Total_water1=%f Total_water2=%f BetLevel=%d",
			lobbyinfo.PlatformID, lobbyinfo.LobbyID, lobbyinfo.GameID, lobbyinfo.LobbyName, lobbyinfo.LobbyMatchID, lobbyinfo.Total_water1, lobbyinfo.Total_water2, lobbyinfo.BetLevel)
		LobbyInfoList[Count] = lobbyinfo
		Count++

	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	CommonLog_INFO_Printf("大廳數量 GameinfoCount=%d", len(LobbyInfoList))

	return ret
}

//=========================================================================================================================================================================
// mysql 取得會員資訊
// Account  	遊戲帳號
// Password 	遊戲編號
func Mysql_CommonMemberInfoGet(Account string, Password string) (bool, MemberInfo) {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonMemberInfoGet (取得會員資訊) Account=%s, Password=%s", Account, Password)

	SqlQuery = fmt.Sprintf("SELECT User_ID, PlatformID, DeviceID, IP, MacAddress, CreateTime, LoginTime, UpdateTime, Account, Password, NickName,  IdentityNumber, Address, PhoneNumber, Balance, Bonus, Salary, Status, Vip_rank FROM memberinfo where  Account='%s' AND Password='%s';", Account, Password)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	memberInfo := MemberInfo{} // 用來接的物件

	for stmtIns.Next() {

		if err2 := stmtIns.Scan(&memberInfo.User_ID, &memberInfo.PlatformID, &memberInfo.DeviceID, &memberInfo.IP, &memberInfo.MacAddress,
			&memberInfo.CreateTime, &memberInfo.LoginTime, &memberInfo.UpdateTime, &memberInfo.Account, &memberInfo.Password, &memberInfo.NickName,
			&memberInfo.IdentityNumber, &memberInfo.Address,
			&memberInfo.PhoneNumber, &memberInfo.Balance, &memberInfo.Bonus, &memberInfo.Salary,
			&memberInfo.Status, &memberInfo.Vip_rank); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true
		CommonLog_INFO_Printf("PlatformID=%d is User_ID=%d, Account=%s, NickName=%s, Balance(玩家的錢)=%d", memberInfo.PlatformID, memberInfo.User_ID, memberInfo.Account, memberInfo.NickName, memberInfo.Balance)
	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	return ret, memberInfo
}

//=========================================================================================================================================================================
// 更新會員登入狀態
// status 狀態
// auth 驗證物件
func Mysql_CommonMemberInfo_UpdateStatus(status int, Auth PacketCmd_AuthInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonMemberInfo_UpdateStatus (更新會員登入狀態)")

	if status == 1 {
		// 登入
		SqlQuery = fmt.Sprintf("update memberinfo set status=%d, LoginTime='%s', UpdateTime='%s' where Account='%s' AND PassWord='%s';",
			status, Common_NowTimeGet(), Common_NowTimeGet(), Auth.Account, Auth.Password)
	} else {
		// 登出
		SqlQuery = fmt.Sprintf("update memberinfo set status=%d, UpdateTime='%s' where  Account='%s' AND PassWord='%s';",
			status, Common_NowTimeGet(), Auth.Account, Auth.Password)
	}

	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	result, err2 := database.Exec(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	} else {
		ret = true
	}

	if ret == true {
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

		CommonLog_INFO_Printf("update record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 更新會員資料
// member 會員物件
func Mysql_CommonMemberInfo_UpdateData(member MemberInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonMemberInfo_UpdateData (更新會員登入狀態)")

	SqlQuery = fmt.Sprintf("update memberinfo set PlatformID=%d, DeviceID=%d, IP='%s',  MacAddress='%s', UpdateTime='%s', Password='%s', NickName='%s', IdentityNumber='%s', Address='%s', PhoneNumber='%s', Balance=%d, Bonus=%d, Salary=%d, Status=%d, Vip_rank=%d where Account='%s' AND PassWord='%s';",
		member.PlatformID, member.DeviceID, member.IP, member.MacAddress,
		member.UpdateTime,
		member.Password, member.NickName, member.IdentityNumber, member.Address, member.PhoneNumber,
		member.Balance, member.Bonus, member.Salary, member.Status, member.Vip_rank,
		member.Account, member.Password)

	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	result, err2 := database.Exec(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	} else {
		ret = true
	}

	if ret == true {
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

		CommonLog_INFO_Printf("update record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 會員資料新增 (這邊嘗試過 用? 去組指令, 但是出錯時候 列印指令拿到一堆 ??? 所以還是用傳統的 組好mysql字串, 方便除錯 )
func Mysql_CommonMemberInfo_Insert(member MemberInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonMemberInfo_Insert (會員資料)")

	SqlQuery = fmt.Sprintf("INSERT INTO memberinfo(User_ID, PlatformID, DeviceID, IP, MacAddress, CreateTime, LoginTime, UpdateTime, Account, Password, NickName, IdentityNumber, Address, PhoneNumber, Balance, Bonus, Salary, Status, Vip_rank) values(%d,%d,%d,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',%d,%d,%d,%d,%d)",
		member.User_ID, member.PlatformID, member.DeviceID, member.IP, member.MacAddress,
		member.CreateTime, member.LoginTime, member.UpdateTime, member.Account, member.Password, member.NickName, member.IdentityNumber, member.Address, member.PhoneNumber,
		member.Balance, member.Bonus, member.Salary, member.Status, member.Vip_rank)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)

	result, err := database.Exec(SqlQuery)
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("#Mysql_CommonMemberInfo_Insert record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 會員刪除
func Mysql_CommonMemberInfo_Delete(member MemberInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonMemberInfo_Delete (會員刪除) User_ID=%d, Account=%s, NickName=%s",
		member.User_ID, member.Account, member.NickName)

	SqlQuery = fmt.Sprintf("delete from memberinfo where User_ID=%d AND Account='%s' AND IdentityNumber='%s';",
		member.User_ID, member.Account, member.IdentityNumber)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	result, err := database.Exec(SqlQuery) // ? = placeholder
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("delete record number=%d ret=%v", num, ret)
	}

	//defer result.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 會員清單取得
func Mysql_CommonMemberListGet(Auth PacketCmd_AuthInfo) (map[int]MemberInfo, bool) {

	var ret bool = false
	var (
		MemberInfoListTmp = make(map[int]MemberInfo) // map MemberInfo
	)

	CommonLog_INFO_Printf("#Mysql_CommonMemberListGet (取得大廳資訊) PlatformID=%d, Account=%s, Password=%s",
		Auth.PlatformID, Auth.Account, Auth.Password)

	var SqlQuery string = ""

	SqlQuery = fmt.Sprintf("SELECT User_ID, PlatformID, DeviceID, IP, MacAddress, CreateTime, LoginTime, UpdateTime, Account, Password, NickName, IdentityNumber, Address, PhoneNumber, Balance, Bonus, Salary, Status, Vip_rank FROM memberinfo where PlatformID=%d order by User_ID DESC;", Auth.PlatformID)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	var Count int = 0
	for stmtIns.Next() {

		member := MemberInfo{} // 用來接的物件
		if err2 := stmtIns.Scan(&member.User_ID, &member.PlatformID, &member.DeviceID, &member.IP, &member.MacAddress, &member.CreateTime, &member.LoginTime, &member.UpdateTime, &member.Account, &member.Password, &member.NickName, &member.IdentityNumber, &member.Address, &member.PhoneNumber, &member.Balance, &member.Bonus, &member.Salary, &member.Status, &member.Vip_rank); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("User_ID=%d, PlatformID=%d is DeviceID=%d, Account=%s NickName=%s IdentityNumber=%s", member.User_ID, member.PlatformID, member.DeviceID, member.Account, member.NickName, member.IdentityNumber)
		MemberInfoListTmp[Count] = member
		Count++

	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	return MemberInfoListTmp, ret
}

//=========================================================================================================================================================================
// 會員清單取得
func Mysql_CommonMemberListGet2(Auth PacketCmd_AuthInfo) (ResponseInfo_MemberInfoList, bool) {

	var ret bool = false
	var (
		MemberInfoList ResponseInfo_MemberInfoList
	)

	CommonLog_INFO_Printf("#Mysql_CommonMemberListGet(會員清單取得) PlatformID=%d, Account=%s, Password=%s",
		Auth.PlatformID, Auth.Account, Auth.Password)

	var SqlQuery string = ""

	SqlQuery = fmt.Sprintf("SELECT User_ID, PlatformID, DeviceID, IP, MacAddress, CreateTime, LoginTime, UpdateTime, Account, Password, NickName, IdentityNumber, Address, PhoneNumber, Balance, Bonus, Salary, Status, Vip_rank FROM memberinfo where PlatformID=%d order by User_ID DESC;", Auth.PlatformID)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	MemberInfoList.Data_Count = 0
	MemberInfoList.Member_List = make(map[int]MemberInfo)
	for stmtIns.Next() {

		member := MemberInfo{} // 用來接的物件
		if err2 := stmtIns.Scan(&member.User_ID, &member.PlatformID, &member.DeviceID, &member.IP, &member.MacAddress, &member.CreateTime, &member.LoginTime, &member.UpdateTime, &member.Account, &member.Password, &member.NickName, &member.IdentityNumber, &member.Address, &member.PhoneNumber, &member.Balance, &member.Bonus, &member.Salary, &member.Status, &member.Vip_rank); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("User_ID=%d, PlatformID=%d is DeviceID=%d, Account=%s NickName=%s IdentityNumber=%s", member.User_ID, member.PlatformID, member.DeviceID, member.Account, member.NickName, member.IdentityNumber)
		MemberInfoList.Member_List[MemberInfoList.Data_Count] = member
		MemberInfoList.Data_Count++

	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	return MemberInfoList, ret
}

//=========================================================================================================================================================================
// 更新會員的錢  減少
func Mysql_CommonMemberInfo_SubBalance(member MemberInfo, ChangeBalance int64) bool {

	var ret bool = false
	var SqlQuery string = ""
	var ChangeBalanceValue int64 = 0

	CommonLog_INFO_Printf("#Mysql_CommonMemberInfo_SubBalance (更新會員的錢) Account=%s, User_ID=%d, NickName=%s, Balance(玩家的錢)=%d, ChangeBalance(變化的錢)=%d",
		member.Account, member.User_ID, member.NickName, member.Balance, ChangeBalance)

	// 錢的變化量
	ChangeBalanceValue = member.Balance - ChangeBalance
	if ChangeBalanceValue < 0 {
		CommonLog_WARNING_Printf("#Mysql_CommonMemberInfo_SubBalance 失敗 餘額不足 Account=%s, User_ID=%d, NickName=%s, Balance(玩家的錢)=%d, ChangeBalance(變化的錢)=%d",
			member.Account, member.User_ID, member.NickName, member.Balance, ChangeBalance)
		return false
	}

	// 是否要檢查變化量的上限?

	SqlQuery = fmt.Sprintf("update memberinfo set Balance=%d, UpdateTime='%s'  where  Account='%s' AND PassWord='%s' AND User_ID=%d;",
		ChangeBalanceValue, Common_NowTimeGet(), member.Account, member.Password, member.User_ID)

	CommonLog_INFO_Printf("#Mysql_CommonMemberInfo_SubBalance SqlQuery=%s", SqlQuery)
	result, err := database.Exec(SqlQuery) // ? = placeholder
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("delete record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 更新會員的錢  增加
func Mysql_CommonMemberInfo_AddBalance(member MemberInfo, ChangeBalance int64) bool {

	var ret bool = false
	var SqlQuery string = ""
	var ChangeBalanceValue int64 = 0

	CommonLog_INFO_Printf("#Mysql_CommonMemberInfo_AddBalance (更新會員的錢) Account=%s, User_ID=%d, NickName=%s, Balance(玩家的錢)=%d, ChangeBalance(變化的錢)=%d",
		member.Account, member.User_ID, member.NickName, member.Balance, ChangeBalance)

	// 錢的變化量
	ChangeBalanceValue = member.Balance + ChangeBalance

	// 是否要檢查變化量的上限?

	SqlQuery = fmt.Sprintf("update memberinfo set Balance=%d, UpdateTime='%s'  where  Account='%s' AND PassWord='%s' AND User_ID=%d;",
		ChangeBalanceValue, Common_NowTimeGet(), member.Account, member.Password, member.User_ID)

	CommonLog_INFO_Printf("#Mysql_CommonMemberInfo_AddBalance SqlQuery=%s", SqlQuery)
	result, err := database.Exec(SqlQuery) // ? = placeholder
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("delete record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 更新座位上會員的錢
// seat 要更新的座位資訊
// ChangeBalance_ci  Balance_ci 的變化量
// ChangeBalance_win  Balance_win 的變化量
func Mysql_SeatInfo_ChangeBalance(seat SeatInfo, ChangeBalance_ci int64, ChangeBalance_win int64) bool {

	var ret bool = false
	var SqlQuery string = ""
	var ChangeBalance_ci_Value int64 = 0
	var ChangeBalance_win_Value int64 = 0

	CommonLog_INFO_Printf("#Mysql_SeatInfo_ChangeBalance (更新座位上會員的錢) Account=%s, User_ID=%d, NickName=%s, Balance_ci(玩家的錢)=%d, Balance_win(玩家的錢)=%d, ChangeBalance_ci(變化的錢)=%d, ChangeBalance_win(變化的錢)=%d",
		seat.Account, seat.User_ID, seat.NickName, seat.Balance_ci, seat.Balance_win, ChangeBalance_ci, ChangeBalance_win)

	// 是否要檢查變化量的上限?

	// 錢的變化量檢查
	ChangeBalance_ci_Value = seat.Balance_ci + ChangeBalance_ci
	if ChangeBalance_ci_Value < 0 {
		CommonLog_WARNING_Printf("#Mysql_SeatInfo_ChangeBalance 失敗 餘額不足 Account=%s, User_ID=%d, NickName=%s, Balance_ci(玩家的錢)=%d, Balance_win(玩家的錢)=%d, ChangeBalance_ci(變化的錢)=%d, ChangeBalance_win(變化的錢)=%d",
			seat.Account, seat.User_ID, seat.NickName, seat.Balance_ci, seat.Balance_win, ChangeBalance_ci, ChangeBalance_win)
		return false
	}
	// 錢的變化量檢查
	ChangeBalance_win_Value = seat.Balance_win + ChangeBalance_win
	if ChangeBalance_win_Value < 0 {
		CommonLog_WARNING_Printf("#Mysql_SeatInfo_ChangeBalance 失敗 餘額不足 Account=%s, User_ID=%d, NickName=%s, Balance_ci(玩家的錢)=%d, Balance_win(玩家的錢)=%d, ChangeBalance_ci(變化的錢)=%d, ChangeBalance_win(變化的錢)=%d",
			seat.Account, seat.User_ID, seat.NickName, seat.Balance_ci, seat.Balance_win, ChangeBalance_ci, ChangeBalance_win)
		return false
	}

	SqlQuery = fmt.Sprintf("update seatinfo set Balance_ci=%d, Balance_win=%d, UpdateTime='%s'  where PlatformID=%d AND GameID=%d AND TableID='%s' AND Seat_ID=%d AND Account='%s' AND User_ID=%d;",
		ChangeBalance_ci_Value, ChangeBalance_win_Value, Common_NowTimeGet(), seat.PlatformID, seat.GameID, seat.TableID, seat.Seat_ID, seat.Account, seat.User_ID)

	CommonLog_INFO_Printf("#Mysql_SeatInfo_ChangeBalance SqlQuery=%s", SqlQuery)
	result, err := database.Exec(SqlQuery) // ? = placeholder
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
		ret = false
	} else {
		ret = true

		num, err := result.RowsAffected()
		if err != nil {
			CommonLog_INFO_Printf("fetch row affected failed:", err.Error())
			ret = false
		}

		/*
			// 因為 ChangeBalance_ci ChangeBalance_win 都是0的時候 就會沒更新任何記錄
			// 何時等於0呢 當押注 跟贏分 一樣的時候
			// 所以不能檢查太細
			if num > 0 {
				ret = true
			} else {
				ret = false
			}*/

		CommonLog_INFO_Printf("影響的記錄 record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// mysql 取得遊戲資訊
func Mysql_CommonGameInfoGetAll() bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonGameInfoGetAll (取得遊戲資訊)")

	SqlQuery = fmt.Sprintf("SELECT PlatformID, GameID, GameName, GameEnName, GameMode, TableDestoryMode, TablePlayerMax FROM gameinfo")
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	var Count int = 0
	for stmtIns.Next() {

		gameInfo := GameInfo{} // 用來接的物件
		if err2 := stmtIns.Scan(&gameInfo.PlatformID, &gameInfo.GameID, &gameInfo.GameName, &gameInfo.GameEnName, &gameInfo.GameMode, &gameInfo.TableDestoryMode, &gameInfo.TablePlayerMax); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("PlatformID=%d is GameID=%d, GameName=%s GameEnName=%s GameMode=%d TableDestoryMode=%d, TablePlayerMax=%d",
			gameInfo.PlatformID, gameInfo.GameID, gameInfo.GameName, gameInfo.GameEnName, gameInfo.GameMode, gameInfo.TableDestoryMode, gameInfo.TablePlayerMax)
		GameInfoList[Count] = gameInfo
		Count++

	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	CommonLog_INFO_Printf("遊戲數量 GameinfoCount=%d", len(GameInfoList))

	return ret
}

//=========================================================================================================================================================================
// mysql 取得桌子資訊
func Mysql_CommonTableInfoGetAll() bool {

	var ret bool = false
	var SqlQuery string = ""

	CommonLog_INFO_Printf("#Mysql_CommonTableInfoGetAll (取得桌子資訊)")

	SqlQuery = fmt.Sprintf("SELECT PlatformID, LobbyID, GameID, TableID, TableArrayIdx, TablePlayerMax, TablePlayerNow, GameMode, TableDestoryMode, CreateTime, UpdateTime FROM tableinfo order by TableArrayIdx")
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	TableInfoCount = 0
	for stmtIns.Next() {

		tableInfo := TableInfo{} // 用來接的物件
		if err2 := stmtIns.Scan(&tableInfo.PlatformID, &tableInfo.LobbyID, &tableInfo.GameID, &tableInfo.TableID, &tableInfo.TableArrayIdx, &tableInfo.TablePlayerMax, &tableInfo.TablePlayerNow, &tableInfo.GameMode, &tableInfo.TableDestoryMode, &tableInfo.CreateTime, &tableInfo.UpdateTime); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("PlatformID=%d is LobbyID=%d, GameID=%d TableID=%s TablePlayerMax=%d, TablePlayerNow=%d, GameMode=%d, TableDestoryMode=%d, CreateTime=%s, UpdateTime=%s",
			tableInfo.PlatformID, tableInfo.LobbyID, tableInfo.GameID, tableInfo.TableID, tableInfo.TablePlayerMax, tableInfo.TablePlayerNow, tableInfo.GameMode,
			tableInfo.TableDestoryMode, &tableInfo.CreateTime, &tableInfo.UpdateTime)

		if tableInfo.TableArrayIdx >= 0 && tableInfo.TableArrayIdx < TABLEINFO_MAX {
			tableInfo.bUse = true
			tableInfo.TablePlayerNow = 0                       // 強制設定成 0 等到 Mysql_CommonSeatInfoGetAll 再依照實際算到的值去更新
			TableInfoList[tableInfo.TableArrayIdx] = tableInfo // 根據 TableArrayIdx 塞入 相對應的 array[] 內
			CommonLog_INFO_Printf("TableID=%s, TablePlayerNow=%d", TableInfoList[tableInfo.TableArrayIdx].TableID, TableInfoList[tableInfo.TableArrayIdx].TablePlayerNow)
			TableInfoCount++
		} else {
			// 跑錯誤流程
			ret = false
			CommonLog_INFO_Printf("錯誤的TableArrayIdx=", tableInfo.TableArrayIdx)
		}

	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	CommonLog_INFO_Printf("桌子 TableInfoCount=%d", TableInfoCount)

	return ret
}

//=========================================================================================================================================================================
// mysql 顯示桌子資訊
func Mysql_CommonTableInfoShow() {

	CommonLog_INFO_Printf("#Mysql_CommonTableInfoShow 顯示桌子資訊 TableInfoCount=%d", TableInfoCount)
	var Count int = 0

	for i := 0; i < TABLEINFO_MAX; i++ {
		pTable := Match_TableInfoGet(i)

		// 防呆檢查
		if pTable.bUse == false {
			continue
		}

		if Count >= TableInfoCount {
			CommonLog_INFO_Printf("#Mysql_CommonTableInfoShow 快速搜尋完畢 Count=%d, TableInfoCount=%d", Count, TableInfoCount)
			break
		}
		Count++

		CommonLog_INFO_Printf("列印開始 i=%d, PlatformID=%d, LobbyID=%d, GameID=%d, TableID=%s, GameMode=%d, TableDestoryMode=%d, 人數[%d/%d]",
			i, pTable.PlatformID, pTable.LobbyID, pTable.GameID, pTable.TableID, pTable.GameMode, pTable.TableDestoryMode, pTable.TablePlayerNow, pTable.TablePlayerMax)

		// 列印位子資訊
		for j := 0; j < SEAT_MAX; j++ {
			seatInfo := pTable.SeatInfo[j]
			if seatInfo.bUse == false {
				continue
			}
			CommonLog_INFO_Printf("有找到位子 j=%d, TableID=%s seat_id=%d, Account=%s, NickName=%s, User_ID=%d, Balance_ci=%d, Balance_win=%d",
				j, seatInfo.TableID, seatInfo.Seat_ID, seatInfo.Account, seatInfo.NickName, seatInfo.User_ID, seatInfo.Balance_ci, seatInfo.Balance_win)
		}
	}
}

//=========================================================================================================================================================================
// mysql 取得位子資訊
func Mysql_CommonSeatInfoGetAll() bool {

	var ret bool = false
	var SqlQuery string = ""

	// 區域變數 組合完桌子資訊就可捨棄
	const SEATINFO_MAX int = TABLEINFO_MAX * SEAT_MAX
	var SeatInfoCount = 0
	var SeatInfoList [SEATINFO_MAX]SeatInfo

	CommonLog_INFO_Printf("#Mysql_CommonSeatInfoGetAll (取得位子資訊)")

	SqlQuery = fmt.Sprintf("SELECT id, PlatformID, LobbyID, GameID, TableID, Seat_ID, GameMode, CreateTime, UpdateTime, User_ID, Account, NickName, Balance_ci, Balance_win, BetLevel, Interval_bet, Interval_bet_pt, Avg_bet, Progress_water, Progress_odds, Progress_support, Wait_item_id, Wait_item_seat, Wait_item_odds, Wait_item_value, Win_item_id, Win_item_bet, Win_item_value, Win_item_win FROM seatinfo order by id")
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	SeatInfoCount = 0
	for stmtIns.Next() {

		seatInfo := SeatInfo{} // 用來接的物件
		if err2 := stmtIns.Scan(&seatInfo.id, &seatInfo.PlatformID, &seatInfo.LobbyID, &seatInfo.GameID, &seatInfo.TableID, &seatInfo.Seat_ID, &seatInfo.GameMode, &seatInfo.CreateTime, &seatInfo.UpdateTime, &seatInfo.User_ID, &seatInfo.Account, &seatInfo.NickName, &seatInfo.Balance_ci, &seatInfo.Balance_win, &seatInfo.BetLevel, &seatInfo.Interval_bet, &seatInfo.Interval_bet_pt, &seatInfo.Avg_bet, &seatInfo.Progress_water, &seatInfo.Progress_odds, &seatInfo.Progress_support, &seatInfo.Wait_item_id, &seatInfo.Wait_item_seat, &seatInfo.Wait_item_odds, &seatInfo.Wait_item_value, &seatInfo.Win_item_id, &seatInfo.Win_item_bet, &seatInfo.Win_item_value, &seatInfo.Win_item_win); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("id=%d, PlatformID=%d is LobbyID=%d, GameID=%d, TableID=%s, Seat_ID=%d, GameMode=%d, User_ID=%d, Account=%s, NickName=%s, Balance_ci=%d, Balance_win=%d", seatInfo.id, seatInfo.PlatformID, seatInfo.LobbyID, seatInfo.GameID, seatInfo.TableID, seatInfo.Seat_ID, seatInfo.GameMode, seatInfo.User_ID, seatInfo.Account, seatInfo.NickName, seatInfo.Balance_ci, seatInfo.Balance_win)
		SeatInfoList[SeatInfoCount] = seatInfo // 塞入 相對應的 array[] 內
		SeatInfoCount++
	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	CommonLog_INFO_Printf("位子 SeatInfoCount=%d", SeatInfoCount)

	// 掃描所有的位子,幫他們找到所屬的桌子
	var bfind bool = false
	for i := 0; i < SeatInfoCount; i++ {

		bfind = false
		seatInfo := SeatInfoList[i]
		CommonLog_INFO_Printf("搜尋開始 i=%d, TableID=%s, Account=%s", i, seatInfo.TableID, seatInfo.Account)
		for j := 0; j < TABLEINFO_MAX; j++ {
			pTable := Match_TableInfoGet(j)

			CommonLog_INFO_Printf("搜尋開始 j=%d, TableID=%s", j, pTable.TableID)

			// 防呆檢查
			if pTable.bUse == false {
				CommonLog_WARNING_Printf("錯誤的bUse=%v, TableID=%s", pTable.bUse, pTable.TableID)
				continue
			}

			if pTable.TableID == seatInfo.TableID {
				CommonLog_INFO_Printf("有找到桌子 TableID=%s", seatInfo.TableID)

				// 防呆檢查
				if seatInfo.Seat_ID < 0 || seatInfo.Seat_ID >= SEAT_MAX {
					CommonLog_WARNING_Printf("錯誤的 Seat_ID=%d(數值異常), TableID=%s, Account=%s", seatInfo.Seat_ID, seatInfo.TableID, seatInfo.Account)
					continue
				}

				// 防呆檢查
				seatInfoCheck := pTable.SeatInfo[seatInfo.Seat_ID]
				if seatInfoCheck.bUse == true {
					CommonLog_WARNING_Printf("錯誤的 Seat_ID=%d(已有玩家), TableID=%s, Account=%s", seatInfo.Seat_ID, seatInfo.TableID, seatInfo.Account)
					continue
				}

				// 找到位子坐下
				seatInfo.bUse = true
				pTable.SeatInfo[seatInfo.Seat_ID] = seatInfo
				pTable.TablePlayerNow++
				Mysql_CommonTableInfo_Update(*pTable)
				CommonLog_INFO_Printf("有找到位子 i=%d, TableID=%s seat_id=%d, Account=%s", i, seatInfo.TableID, seatInfo.Seat_ID, seatInfo.Account)
				bfind = true
				break
			}
		}

		if bfind == false {
			CommonLog_WARNING_Printf("沒找到位子2 i=%d, TableID=%s, Account=%s, Seat_ID=%d", i, seatInfo.TableID, seatInfo.Account, seatInfo.Seat_ID)
		}
	}

	return ret
}

//=========================================================================================================================================================================
// 位子刪除
func Mysql_CommonSeatInfo_Delete(PlatformID int, LobbyID int, GameID int, TableID string, Seat_ID int, User_ID int64) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonSeatInfo_Delete (位子刪除)")

	SqlQuery = fmt.Sprintf("delete from seatinfo where PlatformID=%d AND LobbyID=%d AND GameID=%d AND TableID='%s' AND Seat_ID=%d AND User_ID=%d;", PlatformID, LobbyID, GameID, TableID, Seat_ID, User_ID)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	result, err := database.Exec(SqlQuery) // ? = placeholder
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("delete record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// mysql 重新取得會員座位資訊
// Account  	遊戲帳號
// Password 	遊戲編號
func Mysql_CommonSeatInfo_Get(TableID string, seat SeatInfo) (bool, SeatInfo) {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonSeatInfo_Get (重新取得會員座位資訊) TableID=%s, Account=%s, User_ID=%d, NickName=%s", TableID, seat.Account, seat.User_ID, seat.NickName)

	SqlQuery = fmt.Sprintf("SELECT id, PlatformID, LobbyID, GameID, TableID, Seat_ID, GameMode, CreateTime, UpdateTime, User_ID, Account, NickName, Balance_ci, Balance_win, BetLevel, Interval_bet, Interval_bet_pt, Avg_bet, Progress_water, Progress_odds, Progress_support, Wait_item_id, Wait_item_seat, Wait_item_odds, Wait_item_value, Win_item_id, Win_item_bet, Win_item_value, Win_item_win FROM seatinfo where TableID='%s' AND Seat_ID=%d AND Account='%s' AND User_ID=%d",
		TableID, seat.Seat_ID, seat.Account, seat.User_ID)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	seatInfo := SeatInfo{} // 用來接的物件
	for stmtIns.Next() {

		if err2 := stmtIns.Scan(&seatInfo.id, &seatInfo.PlatformID, &seatInfo.LobbyID, &seatInfo.GameID, &seatInfo.TableID, &seatInfo.Seat_ID, &seatInfo.GameMode, &seatInfo.CreateTime, &seatInfo.UpdateTime, &seatInfo.User_ID, &seatInfo.Account, &seatInfo.NickName, &seatInfo.Balance_ci, &seatInfo.Balance_win, &seatInfo.BetLevel, &seatInfo.Interval_bet, &seatInfo.Interval_bet_pt, &seatInfo.Avg_bet, &seatInfo.Progress_water, &seatInfo.Progress_odds, &seatInfo.Progress_support, &seatInfo.Wait_item_id, &seatInfo.Wait_item_seat, &seatInfo.Wait_item_odds, &seatInfo.Wait_item_value, &seatInfo.Win_item_id, &seatInfo.Win_item_bet, &seatInfo.Win_item_value, &seatInfo.Win_item_win); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("id=%d, PlatformID=%d is LobbyID=%d, GameID=%d, TableID=%s, Seat_ID=%d, GameMode=%d, User_ID=%d, Account=%s, NickName=%s, Balance_ci=%d, Balance_win=%d", seatInfo.id, seatInfo.PlatformID, seatInfo.LobbyID, seatInfo.GameID, seatInfo.TableID, seatInfo.Seat_ID, seatInfo.GameMode, seatInfo.User_ID, seatInfo.Account, seatInfo.NickName, seatInfo.Balance_ci, seatInfo.Balance_win)
	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	return ret, seatInfo
}

//=========================================================================================================================================================================
// 位子新增 (這邊嘗試過 用? 去組指令, 但是出錯時候 列印指令拿到一堆 ??? 所以還是用傳統的 組好mysql字串, 方便除錯 )
func Mysql_CommonSeatInfo_Insert(seat SeatInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonSeatInfo_Insert (位子新增)")

	//SqlQuery = fmt.Sprintf("INSERT INTO seatinfo(PlatformID, LobbyID, GameID, TableID, Seat_ID, SeatArrayIdx, GameMode, CreateTime, UpdateTime, User_ID, Account, Balance_ci, BetLevel, Interval_bet, Interval_bet_pt, Avg_bet, Progress_water, Progress_odds, Progress_support, Wait_item_id, Wait_item_seat, Wait_item_odds, Wait_item_value, Win_item_id, Win_item_bet, Win_item_value, Win_item_win) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",)
	SqlQuery = fmt.Sprintf("INSERT INTO seatinfo(PlatformID, LobbyID, GameID, TableID, Seat_ID, GameMode, CreateTime, UpdateTime, User_ID, Account, NickName, Balance_ci, Balance_win, BetLevel, Interval_bet, Interval_bet_pt, Avg_bet, Progress_water, Progress_odds, Progress_support, Wait_item_id, Wait_item_seat, Wait_item_odds, Wait_item_value, Win_item_id, Win_item_bet, Win_item_value, Win_item_win) values(%d,%d,%d,'%s',%d,%d,'%s','%s',%d,'%s','%s',%d,%d,%d,%d,%d,%d,%d,%d,%d,'%s','%s',%d,%d,'%s',%d,%d,%d)",
		seat.PlatformID, seat.LobbyID, seat.GameID, seat.TableID, seat.Seat_ID, seat.GameMode,
		Common_NowTimeGet(), Common_NowTimeGet(),
		seat.User_ID, seat.Account, seat.NickName,
		seat.Balance_ci, seat.Balance_win, seat.BetLevel,
		seat.Interval_bet, seat.Interval_bet_pt,
		seat.Avg_bet, seat.Progress_water, seat.Progress_odds, seat.Progress_support, seat.Wait_item_id,
		seat.Wait_item_seat, seat.Wait_item_odds, seat.Wait_item_value,
		seat.Win_item_id, seat.Win_item_bet, seat.Win_item_value, seat.Win_item_win)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)

	result, err := database.Exec(SqlQuery)
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("delete record number=%d ret=%v", num, ret)
	}
	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 桌子新增
func Mysql_CommonTableInfo_Insert(table TableInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonTableInfo_Insert (桌子新增)")

	SqlQuery = fmt.Sprintf("INSERT INTO tableinfo(PlatformID, LobbyID, GameID, TableID,TableArrayIdx,TablePlayerMax,TablePlayerNow,GameMode,TableDestoryMode,CreateTime,UpdateTime,LobbyMatchID) values(?,?,?,?,?,?,?,?,?,?,?,?)")
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)

	result, err := database.Exec(SqlQuery,
		table.PlatformID,
		table.LobbyID,
		table.GameID,
		table.TableID,
		table.TableArrayIdx,
		table.TablePlayerMax,
		table.TablePlayerNow,
		table.GameMode,
		table.TableDestoryMode,
		Common_NowTimeGet(),
		Common_NowTimeGet(),
		table.LobbyMatchID)
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("delete record number=%d ret=%v", num, ret)
	}
	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 桌子刪除
func Mysql_CommonTableInfo_Delete(PlatformID int, LobbyID int, GameID int, TableID string, TableArrayIdx int) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonTableInfo_Delete (桌子刪除)")

	SqlQuery = fmt.Sprintf("delete from tableinfo where PlatformID=%d AND LobbyID=%d AND GameID=%d AND TableID='%s' AND TableArrayIdx=%d;", PlatformID, LobbyID, GameID, TableID, TableArrayIdx)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	result, err := database.Exec(SqlQuery) // ? = placeholder
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("delete record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 桌子更新
func Mysql_CommonTableInfo_Update(table TableInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonTableInfo_Delete (桌子更新)")

	SqlQuery = fmt.Sprintf("update tableinfo set TablePlayerNow=%d,UpdateTime='%s' where PlatformID=%d AND LobbyID=%d AND GameID=%d AND TableID='%s' AND TableArrayIdx=%d;",
		table.TablePlayerNow, Common_NowTimeGet(), table.PlatformID, table.LobbyID, table.GameID, table.TableID, table.TableArrayIdx)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	result, err := database.Exec(SqlQuery) // ? = placeholder
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("delete record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// mysql 取得桌子流水編號
func Mysql_CommonTableInfo_IdGet() int64 {

	var SqlQuery string = ""
	var tableName string = ""
	var id sql.NullInt64
	var return_id int64 = 0
	CommonLog_INFO_Printf("#Mysql_CommonTableInfo_IdGet (mysql 取得桌子流水編號)")

	//SqlQuery = fmt.Sprintf("SELECT id FROM tableinfo")
	SqlQuery = fmt.Sprintf("select table_name, AUTO_INCREMENT from information_schema.tables where table_name='tableinfo';")
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
	}

	for stmtIns.Next() {

		//if err2 := stmtIns.Scan(&id); err2 != nil {
		if err2 := stmtIns.Scan(&tableName, &id); err2 != nil {
			CommonLog_WARNING_Println(err2)
			break
		}
	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	return_id = id.Int64
	CommonLog_INFO_Printf("桌子 tableName=%s, id=%d", tableName, return_id)

	return return_id
}

//=========================================================================================================================================================================
// 顧客資料新增 (這邊嘗試過 用? 去組指令, 但是出錯時候 列印指令拿到一堆 ??? 所以還是用傳統的 組好mysql字串, 方便除錯 )
func Mysql_CommonCustomerInfo_Insert(customer CustomerInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonCustomerInfo_Insert (會員資料)")

	customer.CreateTime = Common_NowTimeGet()
	customer.UpdateTime = Common_NowTimeGet()

	SqlQuery = fmt.Sprintf("INSERT INTO customer(User_ID, NickName, CreateTime, UpdateTime, CustomerName, CustomerAge, CustomerGender, CustomerIdentityNumber, CustomerPhoneNumber, CustomerAddress, CustomerHomeID, CustomerHomeAge, CustomerHomeFootage, CustomerHomePrice, Vip_rank ) values(%d,'%s','%s','%s','%s',%d,'%s','%s','%s','%s','%d',%d,%f,%d,%d)",
		customer.User_ID, customer.NickName, customer.CreateTime, customer.UpdateTime, customer.CustomerName,
		customer.CustomerAge, customer.CustomerGender, customer.CustomerIdentityNumber,
		customer.CustomerPhoneNumber, customer.CustomerAddress,
		customer.CustomerHomeID, customer.CustomerHomeAge, customer.CustomerHomeFootage, customer.CustomerHomePrice, customer.Vip_rank)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)

	result, err := database.Exec(SqlQuery)
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("#Mysql_CommonCustomerInfo_Insert record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 顧客資料更新 (這邊嘗試過 用? 去組指令, 但是出錯時候 列印指令拿到一堆 ??? 所以還是用傳統的 組好mysql字串, 方便除錯 )
func Mysql_CommonCustomerInfo_Update(User_ID int64, customer CustomerInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonCustomerInfo_Update (顧客資料更新)")

	customer.UpdateTime = Common_NowTimeGet()

	SqlQuery = fmt.Sprintf("update customer set UpdateTime='%s', CustomerName='%s', CustomerAge=%d, CustomerGender='%s', CustomerIdentityNumber='%s', CustomerPhoneNumber='%s', CustomerAddress='%s', CustomerHomeID=%d, CustomerHomeAge=%d, CustomerHomeFootage=%f, CustomerHomePrice=%d, Vip_rank=%d where User_ID=%d AND CustomerName='%s' AND CustomerIdentityNumber='%s'; ",
		customer.UpdateTime,
		customer.CustomerName, customer.CustomerAge, customer.CustomerGender, customer.CustomerIdentityNumber,
		customer.CustomerPhoneNumber, customer.CustomerAddress,
		customer.CustomerHomeID, customer.CustomerHomeAge, customer.CustomerHomeFootage, customer.CustomerHomePrice, customer.Vip_rank,
		User_ID, customer.CustomerName, customer.CustomerIdentityNumber)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)

	result, err := database.Exec(SqlQuery)
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("#Mysql_CommonCustomerInfo_Update record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 顧客刪除
func Mysql_CommonCustomerInfo_Delete(User_ID int64, customer CustomerInfo) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonCustomerInfo_Delete (顧客刪除) User_ID=%d, CustomerName=%s, CustomerIdentityNumber=%s",
		User_ID, customer.CustomerName, customer.CustomerIdentityNumber)

	customer.CreateTime = Common_NowTimeGet()
	customer.UpdateTime = Common_NowTimeGet()

	SqlQuery = fmt.Sprintf("delete from customer where User_ID=%d AND CustomerName='%s' AND CustomerIdentityNumber='%s';",
		User_ID, customer.CustomerName, customer.CustomerIdentityNumber)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	result, err := database.Exec(SqlQuery) // ? = placeholder
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("Mysql_CommonCustomerInfo_Delete record number=%d ret=%v", num, ret)
	}

	//defer result.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 顧客清單取得 Customer
func Mysql_CommonCustomerListGet(member MemberInfo) (ResponseInfo_CustomerInfoList, bool) {

	var ret bool = false
	var (
		CustomerInfoList ResponseInfo_CustomerInfoList
	)

	CommonLog_INFO_Printf("#Mysql_CommonMemberListGet (顧客清單取得) User_ID=%d", member.User_ID)

	var SqlQuery string = ""

	SqlQuery = fmt.Sprintf("SELECT User_ID, NickName , CreateTime, UpdateTime, CustomerName, CustomerAge, CustomerGender, CustomerIdentityNumber, CustomerPhoneNumber, CustomerAddress, CustomerHomeID, CustomerHomeAge, CustomerHomeFootage, CustomerHomePrice, Vip_rank FROM customer where User_ID=%d order by User_ID DESC;", member.User_ID)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	CustomerInfoList.Data_Count = 0
	CustomerInfoList.Customer_List = make(map[int]CustomerInfo)
	for stmtIns.Next() {

		customer := CustomerInfo{} // 用來接的物件
		if err2 := stmtIns.Scan(&customer.User_ID, &customer.NickName, &customer.CreateTime,
			&customer.UpdateTime, &customer.CustomerName, &customer.CustomerAge, &customer.CustomerGender,
			&customer.CustomerIdentityNumber, &customer.CustomerPhoneNumber, &customer.CustomerAddress, &customer.CustomerHomeID,
			&customer.CustomerHomeAge, &customer.CustomerHomeFootage, &customer.CustomerHomePrice, &customer.Vip_rank); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("User_ID=%d, NickName=%s is CustomerName=%s, CustomerIdentityNumber=%s CustomerPhoneNumber=%s CustomerAddress=%s",
			customer.User_ID, customer.NickName, customer.CustomerName, customer.CustomerIdentityNumber,
			customer.CustomerPhoneNumber, customer.CustomerAddress)
		CustomerInfoList.Customer_List[CustomerInfoList.Data_Count] = customer
		CustomerInfoList.Data_Count++

	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	return CustomerInfoList, ret
}

//=========================================================================================================================================================================
// 工作資料新增 (這邊嘗試過 用? 去組指令, 但是出錯時候 列印指令拿到一堆 ??? 所以還是用傳統的 組好mysql字串, 方便除錯 )
func Mysql_CommonTask_Insert(task Task) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonTask_Insert (工作資料)")

	task.CreateTime = Common_NowTimeGet()
	task.UpdateTime = Common_NowTimeGet()

	SqlQuery = fmt.Sprintf("INSERT INTO task(User_ID, NickName, CreateTime, UpdateTime, TaskName, TaskDescribe, Memo) values(%d,'%s','%s','%s','%s','%s','%s')",
		task.User_ID, task.NickName, task.CreateTime, task.UpdateTime, task.TaskName, task.TaskDescribe, task.Memo)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)

	result, err := database.Exec(SqlQuery)
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("#Mysql_CommonTask_Insert record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 工作資料更新 (這邊嘗試過 用? 去組指令, 但是出錯時候 列印指令拿到一堆 ??? 所以還是用傳統的 組好mysql字串, 方便除錯 )
func Mysql_CommonTask_Update(User_ID int64, task Task) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonTask_Update (工作資料更新)")

	task.UpdateTime = Common_NowTimeGet()

	SqlQuery = fmt.Sprintf("update task set UpdateTime='%s', TaskName='%s', TaskDescribe='%s', Memo='%s' where User_ID=%d AND Task_ID=%d; ",
		task.UpdateTime, task.TaskName, task.TaskDescribe, task.Memo, task.User_ID, task.TaskID)

	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)

	result, err := database.Exec(SqlQuery)
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("#Mysql_CommonTask_Update record number=%d ret=%v", num, ret)
	}

	//defer database.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 工作刪除
func Mysql_CommonTask_Delete(User_ID int64, task Task) bool {

	var ret bool = false
	var SqlQuery string = ""
	CommonLog_INFO_Printf("#Mysql_CommonTask_Delete (工作刪除) User_ID=%d, NickName=%s, TaskID=%d, TaskName=%s",
		User_ID, task.NickName, task.TaskID, task.TaskName)

	SqlQuery = fmt.Sprintf("delete from task where User_ID=%d AND Task_ID=%d;", User_ID, task.TaskID)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	result, err := database.Exec(SqlQuery) // ? = placeholder
	if err != nil {
		ErrorStr := err.Error()
		CommonLog_WARNING_Println(ErrorStr)
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

		CommonLog_INFO_Printf("Mysql_CommonTask_Delete record number=%d ret=%v", num, ret)
	}

	//defer result.Close() // Close the statement when we leave main() / the program

	return ret
}

//=========================================================================================================================================================================
// 工作清單取得 task
func Mysql_CommonTaskListGet(member MemberInfo) (ResponseInfo_TaskList, bool) {

	var ret bool = false
	var (
		TaskList ResponseInfo_TaskList
	)

	CommonLog_INFO_Printf("#Mysql_CommonMemberListGet (工作清單取得) User_ID=%d", member.User_ID)

	var SqlQuery string = ""

	SqlQuery = fmt.Sprintf("SELECT User_ID, NickName, CreateTime, UpdateTime, Task_ID, TaskName, TaskDescribe, Memo FROM task where User_ID=%d order by Task_ID DESC;", member.User_ID)
	CommonLog_INFO_Printf("SqlQuery=%s", SqlQuery)
	stmtIns, err2 := database.Query(SqlQuery) // ? = placeholder
	if err2 != nil {
		ErrorStr2 := err2.Error()
		CommonLog_WARNING_Println(ErrorStr2)
		ret = false
	}

	TaskList.Data_Count = 0
	TaskList.Task_List = make(map[int]Task)
	for stmtIns.Next() {

		task := Task{} // 用來接的物件
		if err2 := stmtIns.Scan(&task.User_ID, &task.NickName, &task.CreateTime, &task.UpdateTime, &task.TaskID, &task.TaskName,
			&task.TaskDescribe, &task.Memo); err2 != nil {
			CommonLog_WARNING_Println(err2)
			ret = false
			break
		}
		ret = true

		CommonLog_INFO_Printf("User_ID=%d, NickName=%s is CreateTime=%s, UpdateTime=%s TaskID=%d TaskName=%s TaskDescribe=%s, Memo=%s",
			task.User_ID, task.NickName, task.CreateTime, task.UpdateTime, task.TaskID, task.TaskName, task.TaskDescribe, task.Memo)
		TaskList.Task_List[TaskList.Data_Count] = task
		TaskList.Data_Count++

	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program

	return TaskList, ret
}
