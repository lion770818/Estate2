package One1CloudLib

import (
	//"io"
	//"io/ioutil"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// 參考1 http://legendtkl.com/2016/03/11/go-log/
// 參考2 http://gotaly.blog.51cto.com/8861157/1405754

var (
	Info    *log.Logger // 一般資訊輸出
	Warning *log.Logger // 警告資訊輸出
	Error   *log.Logger // 錯誤資訊輸出

)

//var logFile	*File
var logFile *os.File = nil      // 開啟的檔案
var logFile_back *os.File = nil // 備份
var timeStr_back string = ""    // 備份
/*
func CommonLog_Println(a ...interface{}) (n int, err error) {
//	return Info(a...)
}
*/

//=========================================================================================================================================================================
// 開啟檔案或重開新檔
func inside_fileopen() {

	var ret bool = false
	var err error
	now := time.Now()
	var timeStr string = fmt.Sprintf("./log/%d_%d_%d_%d.log", now.Year(), now.Month(), now.Day(), now.Hour())

	//log.Println(timeStr_back, timeStr)
	ret = strings.Contains(timeStr_back, timeStr)
	if ret == true {

	} else {
		// 不相等 關閉舊檔案 開新檔案

		if logFile == nil {

			// 第一次寫檔案
			logFile, err = os.OpenFile(timeStr, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			//defer logFile.Close()	// 應該晚一點關閉
			if err != nil {
				log.Fatalln("open file error !")
			}

		} else {

			logFile.Close()
			Info = nil
			Warning = nil
			Error = nil
			logFile, err = os.OpenFile(timeStr, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			//defer logFile.Close()	// 應該晚一點關閉
			if err != nil {
				log.Fatalln("open file error !")
			}
		}

		//log.Println(timeStr_back, timeStr)
		timeStr_back = timeStr // 儲存檔名
		logFile_back = logFile
	}

}

//=========================================================================================================================================================================
//
func CommonLog_INFO_Printf(format string, v ...interface{}) {

	// 開啟檔案或重開新檔
	inside_fileopen()

	if Info == nil {
		Info = log.New(logFile, "[Debug]", log.Ldate|log.Ltime)
	}


	Info.Printf(format, v...)		// 寫檔 ( 理論上也要輸出到螢幕, 但卻不行 )
	log.Printf(format, v...)		// 因為無法輸出到螢幕 所以加這一行
}

//=========================================================================================================================================================================
//
func CommonLog_INFO_Println(v ...interface{}) {

	// 開啟檔案或重開新檔
	inside_fileopen()

	if Info == nil {
		Info = log.New(logFile, "[Debug]", log.Ldate|log.Ltime)
	}
	Info.Println(v...)
	log.Println(v...)		// 因為無法輸出到螢幕 所以加這一行
}

//=========================================================================================================================================================================
//
func CommonLog_WARNING_Printf(format string, v ...interface{}) {
	
		// 開啟檔案或重開新檔
		inside_fileopen()
	
		if Warning == nil {
			Warning = log.New(logFile, "[WARNING]", log.Ldate|log.Ltime)
		}
	
	
		Warning.Printf(format, v...)	// 寫檔 ( 理論上也要輸出到螢幕, 但卻不行 )
		log.Printf(format, v...)		// 因為無法輸出到螢幕 所以加這一行
	}
	
//=========================================================================================================================================================================
//
func CommonLog_WARNING_Println(v ...interface{}) {

	// 開啟檔案或重開新檔
	inside_fileopen()

	if Warning == nil {
		Warning = log.New(logFile, "[WARNING]", log.Ldate|log.Ltime)
	}
	Warning.Println(v...)
	log.Println(v...)		// 因為無法輸出到螢幕 所以加這一行
}

//=========================================================================================================================================================================
//
func CommonLog_ERROR_Fatal(v ...interface{}) {

	// 開啟檔案或重開新檔
	inside_fileopen()

	if Error == nil {
		Error = log.New(logFile, "[ERROR]", log.Ldate|log.Ltime)
	}
	Error.Fatal(v...)
	log.Print(v...)		// 因為無法輸出到螢幕 所以加這一行
	os.Exit(1)
}

//=========================================================================================================================================================================
//
func CommonLog_ERROR_Fatalf(format string, v ...interface{}) {

	// 開啟檔案或重開新檔
	inside_fileopen()

	if Error == nil {
		Error = log.New(logFile, "[ERROR]", log.Ldate|log.Ltime)
	}
	Error.Fatalf(format, v...)
	log.Printf(format, v...)		// 因為無法輸出到螢幕 所以加這一行
	os.Exit(1)
}

//=========================================================================================================================================================================
//
// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func CommonLog_ERROR_Fatalln(v ...interface{}) {

	// 開啟檔案或重開新檔
	inside_fileopen()

	if Error == nil {
		Error = log.New(logFile, "[ERROR]", log.Ldate|log.Ltime)
	}
	Error.Fatalln(v...)
	log.Print(v...)		// 因為無法輸出到螢幕 所以加這一行
	os.Exit(1)
}
