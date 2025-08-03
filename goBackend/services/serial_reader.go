// =================================================================
// 檔案路徑: services/serial_reader.go
// 說明: 負責從序列埠讀取和解析資料的服務
// =================================================================
package services

import (
	"bufio"         // 用於帶緩衝的讀取
	"encoding/json" // 用於解析 JSON
	"log"           // 用於印出日誌
	"sync"          // 用於處理並行讀寫的安全問題 (Mutex)

	"iotDashboard/goBackend/models" // 引入我們自己定義的 Model

	"go.bug.st/serial" // 第三方序列埠函式庫
)

// GlobalDataStore 是一個全域變數，用來儲存最新的一筆感測器資料。
// 在真實的大型專案中，可能會使用資料庫，但對於這個專案，一個全域變數就足夠了。
var GlobalDataStore struct {
	// Mutex (互斥鎖) 是解決並行問題的關鍵。
	// 當一個 Goroutine (我們的背景讀取任務) 正在寫入最新資料時，
	// 另一個 Goroutine (處理網頁請求的任務) 可能會同時來讀取。
	// Mutex 確保同一時間只有一個 Goroutine 能存取 Data，避免資料錯亂。
	Mu   sync.Mutex
	Data models.SensorData
}

// StartSerialReader 是一個會在背景持續運行的函式
// 它接收序列埠名稱 (例如 "COM3") 作為參數
func StartSerialReader(portName string) {
	// 設定序列埠的參數
	mode := &serial.Mode{
		BaudRate: 9600, // 鮑率必須和 Arduino 程式中的 Serial.begin(9600) 一致
	}

	// 嘗試打開指定的序列埠
	port, err := serial.Open(portName, mode)
	if err != nil {
		// 如果打開失敗 (例如 COM Port 名稱錯誤或被佔用)，印出錯誤並結束程式
		log.Fatalf("無法打開序列埠 %s: %v", portName, err)
	}
	log.Printf("成功打開序列埠 %s，開始監聽...", portName)

	// 使用 bufio.NewScanner 來幫助我們一次讀取一行資料
	// Arduino 每次傳送一筆 JSON 資料後都會換行，所以這個方法非常適合
	scanner := bufio.NewScanner(port)
	for scanner.Scan() {
		// 讀取到一行文字 (也就是我們的 JSON 字串)
		line := scanner.Text()

		// 建立一個空的 SensorData 結構體變數，用來存放解析後的資料
		var data models.SensorData

		// 嘗試將讀取到的 JSON 字串解析到 data 變數中
		err := json.Unmarshal([]byte(line), &data)
		if err != nil {
			// 如果解析失敗 (可能資料不完整或格式錯誤)，印出錯誤訊息並繼續下一輪讀取
			log.Printf("解析 JSON 失敗: %v, 原始資料: %s", err, line)
			continue
		}

		// --- 寫入全域變數 ---
		// 在寫入資料前，先鎖定 Mutex，防止其他 Goroutine 同時讀取
		GlobalDataStore.Mu.Lock()
		// 更新全域變數中的資料
		GlobalDataStore.Data = data
		// 寫入完成後，解鎖 Mutex，讓其他 Goroutine 可以讀取
		GlobalDataStore.Mu.Unlock()

		// 在後端控制台印出收到的資料，方便除錯
		// log.Printf("收到資料: 溫度=%.2f°C, 濕度=%.2f%%", data.Temperature, data.Humidity)
	}

	// 如果 scanner 迴圈因為錯誤而結束，印出錯誤訊息
	if err := scanner.Err(); err != nil {
		log.Printf("從序列埠讀取時發生錯誤: %v", err)
	}
}
