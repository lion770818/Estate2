package One1CloudLib

//"fmt"
//"log"
//"time"

//=========================================================================================================================================================================
// 網路封包
// 測試網址     http://www.websocket.org/echo.html
// server ip = ws://192.168.43.75:3000/One1CloudGameCmd		//my
// server ip = ws://192.168.0.105:3000/One1CloudGameCmd     //home
// server ip = ws://192.168.50.143:3000/One1CloudGameCmd	//onmyhome

const (
	NET_CMD_ACCOUNT_CREATE = "account_create" // 帳號建立
	NET_CMD_ACCOUNT_DELETE = "account_delete" // 帳號刪除
	NET_CMD_LOGIN          = "login"          // 登入								{"sys":"system", "cmd":"login", "sn":12345, "isEncode":false,"data":"{\"account\":\"cat111\",\"password\":\"1234\"}"}
	NET_CMD_LOGOUT         = "logout"         // 登出							    {"sys":"system", "cmd":"logout", "sn":12345, "isEncode":false,"data":"{\"account\":\"cat111\",\"password\":\"1234\"}"}

	NET_CMD_MEMBER_INSERT   = "member_insert"   // 新增會員							   {"sys":"system", "cmd":"member_insert", "sn":12345, "isEncode":false,"data":"{\"account\":\"cat111\",\"password\":\"1234\", \"NickName\":\"111\", \"IdentityNumber\":\"F124180631\" }"}
	NET_CMD_MEMBER_UPDATE   = "member_update"   // 更新會員							   {"sys":"system", "cmd":"member_update", "sn":12345, "isEncode":false,"data":"{\"account\":\"cat222\",\"password\":\"1234\", \"NickName\":\"更新會員002\", \"IdentityNumber\":\"F123456789\" }"}
	NET_CMD_MEMBER_DELETE   = "member_delete"   // 刪除會員                            {"sys":"system", "cmd":"member_delete", "sn":12345, "isEncode":false,"data":"{\"user_id\":1,\"account\":\"cat222\",\"password\":\"1234\", \"NickName\":\"更新會員002\", \"IdentityNumber\":\"F123456784\" }"}
	NET_CMD_MEMBER_LIST_GET = "member_list_get" // 會員清單取得						 {"sys":"system", "cmd":"member_list_get", "sn":12345, "isEncode":false,"data":"{\"platform_id\":1,\"account\":\"cat111\",\"password\":\"1234\" }"}

	NEW_CMD_CUSTOMER_INSERT   = "customer_insert"   // 新增顧客						{ "sys":"system","cmd":"customer_insert","sn":0,"isEncode":false,"data":"{\"user_id\":0,\"nickname\":\"\",\"createtime\":\"\",\"updatetime\":\"\",\"customer_name\":\"測試顧客\",\"customer_age\":0,\"customer_gender\":\"\",\"customer_identityNumber\":\"\",\"customer_phone_number\":\"\",\"customer_address\":\"\",\"customer_home_id\":0,\"customer_home_age\":0,\"customer_home_footage\":0,\"customer_home_price\":0,\"vip_rank\":0}"}
	NEW_CMD_CUSTOMER_UPDATE   = "customer_update"   // 更新顧客						{ "sys":"system","cmd":"customer_update","sn":0,"isEncode":false,"data":"{\"user_id\":0,\"nickname\":\"\",\"createtime\":\"\",\"updatetime\":\"\",\"customer_name\":\"測試顧客\",\"customer_age\":0,\"customer_gender\":\"\",\"customer_identityNumber\":\"\",\"customer_phone_number\":\"\",\"customer_address\":\"\",\"customer_home_id\":0,\"customer_home_age\":0,\"customer_home_footage\":0,\"customer_home_price\":0,\"vip_rank\":7}"}
	NET_CMD_CUSTOMER_DELETE   = "customer_delete"   // 刪除顧客                     { "sys":"system","cmd":"customer_delete","sn":0,"isEncode":false,"data":"{\"user_id\":0,\"nickname\":\"\",\"customer_name\":\"測試顧客\",\"customer_identityNumber\":\"\"}"}
	NET_CMD_CUSTOMER_LIST_GET = "customer_list_get" // 顧客清單取得				   {"sys":"system", "cmd":"customer_list_get", "sn":12345, "isEncode":false,"data":"{\"platform_id\":1,\"account\":\"cat111\",\"password\":\"1234\" }"}

	NEW_CMD_TASK_INSERT   = "task_insert"   // 新增工作						{"sys":"system","cmd":"task_insert","sn":0,"isEncode":false,"data":"{\"task_id\":0,\"user_id\":0,\"nickname\":\"\",\"createtime\":\"\",\"updatetime\":\"\",\"task_name\":\"工作清單01\",\"task_describe\":\"工作描述\",\"memo\":\"測試的memo\"}"}
	NEW_CMD_TASK_UPDATE   = "task_update"   // 更新工作						{"sys":"system","cmd":"task_update","sn":0,"isEncode":false,"data":"{\"task_id\":2,\"user_id\":1,\"nickname\":\"\",\"createtime\":\"\",\"updatetime\":\"\",\"task_name\":\"工作清單02\",\"task_describe\":\"工作描述2\",\"memo\":\"更新的memo\"}"}
	NET_CMD_TASK_DELETE   = "task_delete"   // 刪除工作                     {"sys":"system","cmd":"task_delete","sn":0,"isEncode":false,"data":"{\"task_id\":2,\"user_id\":1,\"nickname\":\"\",\"createtime\":\"\",\"updatetime\":\"\",\"task_name\":\"工作清單02\",\"task_describe\":\"工作描述2\",\"memo\":\"更新的memo\"}"}
	NET_CMD_TASK_LIST_GET = "task_list_get" // 工作清單取得				   {"sys":"system", "cmd":"task_list_get", "sn":12345, "isEncode":false,"data":"{\"platform_id\":1,\"account\":\"cat111\",\"password\":\"1234\" }"} 	

	NEW_CMD_HOME_INSERT   = "home_insert"   // 新增房屋						{"sys":"system","cmd":"home_insert","sn":0,"isEncode":false,"data":"{\"user_id\":0,\"nickname\":\"\",\"createtime\":\"\",\"updatetime\":\"\",\"home_name\":\"工作清單01\",\"home_address\":\"我住在地球\",\"home_age\":1,\"home_footage\":2,\"home_price\":200,\"vip_rank\":1,\"memo\":\"測試的memo\"}"}
	NEW_CMD_HOME_UPDATE   = "home_update"   // 更新房屋						{"sys":"system","cmd":"home_update","sn":0,"isEncode":false,"data":"{\"user_id\":1,\"nickname\":\"cat111\",\"home_id\":1,\"createtime\":\"\",\"updatetime\":\"\",\"home_name\":\"鼎藏1\",\"home_address\":\"我住在地球\",\"home_age\":1,\"home_footage\":2,\"home_price\":200,\"vip_rank\":1,\"memo\":\"測試的memo\"}"}
	NET_CMD_HOME_DELETE   = "home_delete"   // 刪除房屋                     {"sys":"system","cmd":"home_delete","sn":0,"isEncode":false,"data":"{\"task_id\":2,\"user_id\":1,\"nickname\":\"\",\"createtime\":\"\",\"updatetime\":\"\",\"task_name\":\"工作清單02\",\"task_describe\":\"工作描述2\",\"memo\":\"更新的memo\"}"}
	NET_CMD_HMOE_LIST_GET = "home_list_get" // 房屋清單取得				   {"sys":"system","cmd":"home_list_get", "sn":12345, "isEncode":false,"data":"{\"platform_id\":1,\"account\":\"cat111\",\"password\":\"1234\" }"} 	

	NET_CMD_MESSAGE_BOARD_INSERT = "message_board_insert" // 留言板新增
	//====== game ==========================================
	NET_CMD_LOBBYINFO_GET = "lobbyInfoGet" // 取得大廳資訊			  		     {"sys":"system", "cmd":"lobbyInfoGet", "sn":12345, "isEncode":false,"data":"{\"platform_id\":1}"}
	NET_CMD_ENTER_GAME    = "enter_game"   // 進入遊戲(autoMatch)				  {"sys":"game", "cmd":"enter_game", "sn":12345, "isEncode":false,"data":"{\"platform_id\":1,\"lobby_id\":1,\"game_id\":1001,\"udid\":1,\"user_id\":1,\"channel\":\"123\",\"publish_ver\":\"1.0.0\",\"refresh\":\"0\",\"balance_ci\":3000 }"}
	//{"sys":"game", "cmd":"enter_game", "sn":12345, "isEncode":false,"data":"{\"platform_id\":1,\"lobby_id\":8,\"game_id\":2001,\"udid\":1,\"user_id\":1,\"channel\":\"123\",\"publish_ver\":\"1.0.0\",\"refresh\":\"0\",\"balance_ci\":3000 }"}
	NET_CMD_JOIN_GAME = "join_game" // 加入遊戲(joinMatch)				  {"sys":"game", "cmd":"join_game", "sn":12345, "isEncode":false,"data":"{\"table_id\":\"FH11001-0000001\",\"table_array_idx\":0,\"user_id\":3,\"balance_ci\":2000}"}
	NET_CMD_EXIT_GAME = "exit_game" // 離開遊戲							  {"sys":"game", "cmd":"exit_game", "sn":12345, "isEncode":false,"data":"{\"table_id\":\"FH11001-0000001\",\"user_id\":1,\"seat_id\":1}"}
	//NET_CMD_SHUTDOWN				= "shutdown"				// 幾分鐘後關機

	NET_CMD_SLOT_SPIN = "slot_spin" // slot spin 開始玩				   {"sys":"game", "cmd":"slot_spin", "sn":12345, "isEncode":false,"data":"{\"table_id\":\"HG12001-0000001\",\"user_id\":1,\"seat_id\":1,\"bet\":100 }"}

	NET_CMD_FISH_SHOOT = "shoot" // 魚機-射擊			   	  	      {"sys":"game", "cmd":"shoot", "sn":12345, "isEncode":false,"data":"{\"table_id\":\"FH11001-0000001\",\"user_id\":1,\"x\":100,\"y\":200,\"bet\":100,\"bullet_type\":\"0\", \"bullet_id\":\"1234\" }"}

	NET_CMD_FISH_NEW_FISH = "new_fish" // 魚機-伺服器主動廣播，通知該桌產生新的魚

)

//=========================================================================================================================================================================
// 錯誤代碼
//=========================================================================================================================================================================
// 共用的回應結構
type Base int
type CommonCodeInfo struct {
	Code    int    // 回應的代碼
	Message string // 回應的訊息
}

const (
	ERROR_CODE_SUCCESS                  Base = iota // 0沒有錯誤
	ERROR_CODE_NO_FIND_CMD                          // 1找不到Cmd
	ERROR_CODE_NO_FIND_ACCOUNT                      // 2找不到帳號
	ERROR_CODE_NO_LOGIN                             // 3帳號未登入
	ERROR_CODE_CLIENT_TOO_MATCH                     // 4服務器上限滿額
	ERROR_CODE_NO_FIND_TABLE                        // 5找不到桌子
	ERROR_CODE_NO_FIND_SEAT                         // 6找不到位子
	ERROR_CODE_NO_USE                               // 7資源未使用
	ERROR_CODE_FULL_PLAYER                          // 8人滿了
	ERROR_CODE_CARRY_BALANCE_NOT_ENOUGH             // 9想帶進來的錢不夠
	ERROR_CODE_RE_JOIN_TABLE                        // 10重複入桌
	ERROR_CODE_ERROR_DATA                           // 11資料錯誤
	ERROR_CODE_ERROR_PARAMETER                      // 12參數錯誤
	ERROR_CODE_ERROR_OPEN_TATBL                     // 13開桌失敗
	ERROR_CODE_ERROR_USER_ID                        // 14錯誤的User_ID
	ERROR_CODE_BALANCE_UPDATE_FAIL                  // 15更新錢錯誤
	ERROR_CODE_TABLE_BALANCE_NOT_ENOUGH             // 16桌內的錢不夠
	ERROR_CODE_BALANCE_CHECK_FAIL                   // 17檢查錢錯誤
	ERROR_CODE_DATA_UPDATE_FAIL                     // 18更新資料錯誤
	ERROR_CODE_ERROR_JSON_MARSHAL                   // 19Json解析錯誤
	ERROR_CODE_ERROR_GAME_MODE                      // 20錯誤的GameMode
	ERROR_CODE_ERROR_JOIN_TABLE                     // 21加入桌失敗
	ERROR_CODE_ERROR_PERMISSION_DENIED              // 22權限不足

	ERROR_CODE_ERROR_CREATE_MEMBER                  // 23新增會員失敗
	ERROR_CODE_ERROR_UPDATE_MEMBER                  // 24更新會員失敗
	ERROR_CODE_ERROR_DELETE_MEMBER                  // 25刪除會員失敗

	ERROR_CODE_ERROR_CREATE_CUSTOMER                // 26新增顧客失敗
	ERROR_CODE_ERROR_UPDATE_CUSTOMER                // 27更新顧客失敗
	ERROR_CODE_ERROR_DELETE_CUSTOMER                // 28刪除顧客失敗	

	ERROR_CODE_ERROR_CREATE_TASK                  	// 29新增工作失敗
	ERROR_CODE_ERROR_UPDATE_TASK                  	// 30更新工作失敗
	ERROR_CODE_ERROR_DELETE_TASK                  	// 31刪除工作失敗	

	ERROR_CODE_ERROR_CREATE_HOME                  	// 32新增房屋失敗
	ERROR_CODE_ERROR_UPDATE_HOME                  	// 33更新房屋失敗
	ERROR_CODE_ERROR_DELETE_HOME                  	// 34刪除房屋失敗	
	ERROR_CODE_MAX
)

//var ErrorCode = [...]CommonCodeInfo{
var ErrorCode = [ERROR_CODE_MAX]CommonCodeInfo{
	{0, "沒有錯誤"},
	{-1, "找不到Cmd"},
	{-2, "找不到帳號"},
	{-3, "帳號未登入"},
	{-4, "服務器上限滿額"},
	{-5, "找不到桌子"},
	{-6, "找不到位子"},
	{-7, "資源未使用"},
	{-8, "人滿了"},
	{-9, "想帶進來的錢不夠"},
	{-10, "重複入桌"},
	{-11, "資料錯誤"},
	{-12, "參數錯誤"},
	{-13, "開桌失敗"},
	{-14, "錯誤的User_ID"},
	{-15, "更新錢錯誤"},
	{-16, "桌內的錢不夠"},
	{-17, "檢查錢錯誤"},
	{-18, "更新資料錯誤"},
	{-19, "Json解析錯誤"},
	{-20, "錯誤的GameMode"},
	{-21, "加入桌失敗"},
	{-22, "權限不足"},
	{-23, "新增會員失敗"},
	{-24, "更新會員失敗"},
	{-25, "刪除會員失敗"},
	{-26, "新增顧客失敗"},
	{-27, "更新顧客失敗"},
	{-28, "刪除顧客失敗"},
	{-29, "新增工作失敗"},
	{-30, "更新工作失敗"},
	{-31, "刪除工作失敗"},	

	{-32, "新增房屋失敗"},	
	{-33, "更新房屋失敗"},	
	{-34, "刪除房屋失敗"},	
}

//=========================================================================================================================================================================
// 玩家狀態
const (
	STATUS_LOGOUT     Base = iota //0:登出
	STATUS_LOGIN                  //1:登入
	STATUS_DISCONNECT             //2:斷線中
	STATUS_RECONNECT              //3:斷線連回中
)
