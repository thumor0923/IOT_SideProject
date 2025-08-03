// =================================================================
// 檔案路徑: main.go
// 說明: 整個後端程式的進入點
// =================================================================
package main

import (
	"log"
	"net/http"

	"iotDashboard/goBackend/controllers"
	"iotDashboard/goBackend/services"
)

func main() {
	// =================================================================
	// !!! 非常重要：請在這裡填入您自己電腦的序列埠名稱 !!!
	// =================================================================
	// Windows 範例: "COM3"
	// macOS 範例: "/dev/tty.usbserial-1420"
	// Linux 範例: "/dev/ttyACM0"
	const serialPortName = "COM3" // <--- 請修改這裡

	// 使用 go 關鍵字，在一個新的 "Goroutine" (可以想成是輕量級的執行緒) 中
	// 啟動我們的序列埠讀取服務。
	// 這樣做的好處是，讀取序列埠的任務會在背景持續運行，
	// 完全不會影響到下方網頁伺服器的啟動與運作。
	// 這就是 Go 語言強大的並行 (Concurrency) 能力！
	go services.StartSerialReader(serialPortName)

	// 設定 API 的路由 (Route)
	// 當有人訪問 "http://localhost:8080/api/sensor-data" 時，
	// 就由 controllers.SensorDataHandler 這個函式來處理請求。
	http.HandleFunc("/api/sensor-data", controllers.SensorDataHandler)

	// 印出訊息，告訴使用者伺服器已經啟動
	log.Println("伺服器已啟動於 http://localhost:8080")
	log.Println("您可以透過訪問 http://localhost:8080/api/sensor-data 來獲取感測器資料")

	// 啟動 Web 伺服器，並監聽 8080 連接埠。
	// 如果啟動失敗，會印出致命錯誤。
	log.Fatal(http.ListenAndServe(":8080", nil))
}
