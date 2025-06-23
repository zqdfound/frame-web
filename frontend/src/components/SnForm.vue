<template>
  <div class="auth-form">
    <div class="form-container">
      <div class="form-group">
        <label>SN:</label>
        <input v-model="sn" type="text" class="form-input" />
      </div>
      <div class="form-group">
        <label>密码:</label>
        <input v-model="pwd" type="password" class="form-input" />
      </div>
      <button @click="submitForm" class="submit-btn">确定</button>
    </div>
    
    <!-- 新增弹窗 -->
    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="showModal = false">&times;</span>
        <h3>设备信息</h3>
        <table class="data-table">
          <tr v-for="(value, key) in responseData.data" :key="key">
            <td class="key-column">{{ key }}</td>
            <td>{{ value }}</td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      sn: '',
      pwd: '',
      showModal: false,
      responseData: {}
    }
  },
  methods: {
    async submitForm() {
      try {
        const response = await fetch('http://localhost:8888/device', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            sn: this.sn,
            pwd: this.pwd
          })
        });
        
        const result = await response.json();
        if (result.code === 0) {
          this.responseData = result;
          this.showModal = true;
        } else {
          alert('请求失败: ' + result.msg);
        }
      } catch (error) {
        alert('请求失败: ' + error.message);
      }
    }
  }
}
</script>

<style scoped>
.auth-form {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f5f5;
}

.form-container {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  width: 300px;
}

.form-group {
  margin-bottom: 1.5rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
}

.form-input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

.submit-btn {
  width: 100%;
  padding: 0.75rem;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
}

.submit-btn:hover {
  background-color: #45a049;
}

/* 新增弹窗样式 */
.modal {
  display: block;
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0,0,0,0.4);
}

.modal-content {
  background-color: #fefefe;
  margin: 10% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
  max-width: 800px;
  border-radius: 8px;
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
  cursor: pointer;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.data-table tr {
  border-bottom: 1px solid #ddd;
}

.data-table td {
  padding: 8px;
  text-align: left;
}

.key-column {
  font-weight: bold;
  width: 30%;
}
</style>