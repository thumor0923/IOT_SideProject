// =================================================================
// 檔案路徑: models/sensor_data.go
// 說明: 新增一個 Command 結構體
// =================================================================
package models

// SensorData 結構體保持不變
type SensorData struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

// Command 結構體用來對應從前端傳來的指令 JSON
type Command struct {
	Command string `json:"command"`
}
