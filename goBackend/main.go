// =================================================================
// 檔案路徑: main.go
// 說明: 註冊新的 API 路由
// =================================================================
package main

import (
	"log"
	"net/http"

    "iotDashboard/goBackend/controllers"
    "iotDashboard/goBackend/services"
)

func main() {
	const serialPortName = "COM3" // <--- 請確認這裡仍然是您正確的序列埠名稱

	go services.StartSerialReader(serialPortName)

	// 註冊原本的數據 API
	http.HandleFunc("/api/sensor-data", controllers.SensorDataHandler)
	// 註冊新的指令 API
	http.HandleFunc("/api/command", controllers.CommandHandler)

	log.Println("伺服器已啟動於 http://localhost:8080")
	log.Println("GET /api/sensor-data - 獲取感測器資料")
	log.Println("POST /api/command - 發送控制指令")

	log.Fatal(http.ListenAndServe(":8080", nil))
}