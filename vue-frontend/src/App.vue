<script setup>
import { ref, onMounted } from 'vue'

// 使用 ref() 建立「響應式變數」。
// 這代表當這些變數的值改變時，畫面上有用到它們的地方會自動更新。
const temperature = ref(0)
const humidity = ref(0)
const lastUpdated = ref(null)
const error = ref(null)

// 定義一個非同步函式來獲取後端 API 的資料
const fetchData = async () => {
  try {
    // 使用瀏覽器內建的 fetch API 向我們的 Go 後端發送請求
    const response = await fetch('http://localhost:8080/api/sensor-data')
    
    // 如果請求失敗 (例如後端伺服器沒開)，就拋出錯誤
    if (!response.ok) {
      throw new Error('後端伺服器回應錯誤')
    }
    
    // 將回應的內容解析為 JSON
    const data = await response.json()
    
    // 更新我們的響應式變數
    // .value 是用來存取 ref() 變數的實際值
    temperature.value = data.temperature.toFixed(2) // 取到小數點後兩位
    humidity.value = data.humidity.toFixed(2)
    lastUpdated.value = new Date().toLocaleTimeString() // 記錄更新時間
    error.value = null // 清除之前的錯誤訊息

  } catch (e) {
    // 如果在 try 區塊中發生任何錯誤，就在這裡捕捉
    console.error(e) // 在瀏覽器開發者工具中印出詳細錯誤
    error.value = '無法獲取資料。請確認後端伺服器是否已啟動，且 Arduino 已連接。' // 在畫面上顯示錯誤訊息
  }
}

// onMounted() 是一個 Vue 的「生命週期鉤子 (Lifecycle Hook)」。
// 它裡面的程式碼，會在元件第一次被掛載到畫面上時執行。
onMounted(() => {
  // 1. 立即執行一次，讓使用者能馬上看到資料
  fetchData()

  // 2. 設定一個計時器 (setInterval)，每 3 秒鐘自動重複執行 fetchData 函式
  // 這樣就能實現儀表板的「即時」更新效果
  setInterval(fetchData, 3000) 
})
</script>

<template>
  <div id="dashboard">
    <header>
      <h1>房間溫濕度儀表板</h1>
      <p v-if="lastUpdated" class="last-updated">最後更新於: {{ lastUpdated }}</p>
    </header>
    
    <main>
      <div v-if="error" class="error-box">
        <p>{{ error }}</p>
      </div>
      
      <div v-else class="cards-container">
        <div class="card temperature-card">
          <h2>溫度</h2>
          <div class="value">
            <span>{{ temperature }}</span>
            <span class="unit">°C</span>
          </div>
        </div>
        
        <div class="card humidity-card">
          <h2>濕度</h2>
          <div class="value">
            <span>{{ humidity }}</span>
            <span class="unit">%</span>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<style scoped>
/* 在這裡我們定義這個元件的 CSS 樣式 */
#dashboard {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  max-width: 800px;
  margin: 40px auto;
  padding: 20px;
  background-color: #f4f7f6;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  color: #333;
}

header {
  text-align: center;
  border-bottom: 1px solid #e0e0e0;
  padding-bottom: 20px;
  margin-bottom: 20px;
}

h1 {
  font-size: 2.5em;
  color: #2c3e50;
  font-weight: 600;
}

.last-updated {
  color: #7f8c8d;
  font-size: 0.9em;
}

.cards-container {
  display: flex;
  justify-content: space-around;
  gap: 20px;
}

.card {
  background-color: #ffffff;
  padding: 30px;
  border-radius: 10px;
  text-align: center;
  flex: 1;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 15px rgba(0, 0, 0, 0.1);
}

h2 {
  margin-top: 0;
  font-size: 1.2em;
  color: #34495e;
  font-weight: 500;
}

.value {
  font-size: 4em;
  font-weight: bold;
}

.temperature-card .value {
  color: #e74c3c;
}

.humidity-card .value {
  color: #3498db;
}

.unit {
  font-size: 0.4em;
  vertical-align: super;
  margin-left: 5px;
}

.error-box {
  background-color: #ffebee;
  color: #c62828;
  border: 1px solid #ef9a9a;
  padding: 15px;
  border-radius: 8px;
  text-align: center;
}
</style>