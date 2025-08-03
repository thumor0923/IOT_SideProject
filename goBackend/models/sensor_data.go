// =================================================================
// 檔案路徑: models/sensor_data.go
// 說明: 定義我們資料的結構 (Model)
// =================================================================
package models

// SensorData 結構體用來對應從 Arduino 傳來的 JSON 資料
// `json:"..."` 這種標籤 (tag) 是用來告訴 Go 的 JSON 套件，
// 在進行 JSON 編碼或解碼時，這個欄位對應的 JSON key 是什麼。
type SensorData struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}
