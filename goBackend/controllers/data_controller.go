// =================================================================
// 檔案路徑: controllers/data_controller.go
// 說明: 處理 HTTP 請求的控制器 (Controller)
// =================================================================
package controllers

import (
	"encoding/json" // 用於將 Go 結構體轉換為 JSON
	"log"
	"net/http" // Go 內建的 HTTP 函式庫

	"iotDashboard/goBackend/services" // 引入我們的服務
)

// SensorDataHandler 是一個處理函式，負責回應對 API 的請求
func SensorDataHandler(w http.ResponseWriter, r *http.Request) {
	// --- 讀取全域變數 ---
	// 在讀取資料前，先鎖定 Mutex，確保我們讀到的是一筆完整的資料
	services.GlobalDataStore.Mu.Lock()
	// 複製一份最新的資料出來
	latestData := services.GlobalDataStore.Data
	// 讀取完成後，立刻解鎖 Mutex
	services.GlobalDataStore.Mu.Unlock()

	// 設定 HTTP 回應的標頭 (Header)，告訴瀏覽器我們回傳的是 JSON 格式
	w.Header().Set("Content-Type", "application/json")
	// 解決跨域問題 (CORS)，允許任何來源的前端網頁來請求我們的 API
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 使用 json.NewEncoder 將 Go 的結構體 (latestData) 編碼成 JSON 格式，
	// 並直接寫入到 HTTP 的回應中 (w)
	err := json.NewEncoder(w).Encode(latestData)
	if err != nil {
		// 如果編碼失敗，在伺服器端印出錯誤
		log.Printf("無法將資料編碼為 JSON: %v", err)
	}
}
