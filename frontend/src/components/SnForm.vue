<template>
  <div v-if="!isLoggedIn" class="auth-form">
    <div class="form-container">
      <div class="form-group">
        <label>用户名:</label>
        <input v-model="username" type="text" class="form-input" />
      </div>
      <div class="form-group">
        <label>密码:</label>
        <input v-model="password" type="password" class="form-input" />
      </div>
      <button @click="login" class="submit-btn">登录</button>
    </div>
  </div>
  
  <div v-else>
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
  </div>
</template>

<script>
export default {
  data() {
    return {
      isLoggedIn: false,
      username: '',
      password: '',
      token: '',
      sn: '',
      pwd: '',
      showModal: false,
      responseData: {}
    }
  },
  methods: {
    async login() {
      try {
        const response = await fetch('http://localhost:8888/users/login', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            username: this.username,
            password: this.password
          })
        });
        
        const result = await response.json();
        if (result.code === 200) {
          this.token = result.data.token;
          localStorage.setItem('token', this.token);
          this.isLoggedIn = true;
        } else {
          alert('登录失败: ' + result.msg);
        }
      } catch (error) {
        alert('登录失败: ' + error.message);
      }
    },
    
    async submitForm() {
      try {
        const response = await fetch('http://localhost:8888/users/device', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${this.token || localStorage.getItem('token')}`
          },
          body: JSON.stringify({
            sn: this.sn,
            pwd: this.pwd
          })
        });
        
        const result = await response.json();
        if (result.code === 200) {
          this.responseData = result;
          this.showModal = true;
        } else {
          alert('请求失败: ' + result.msg);
        }
      } catch (error) {
        alert('请求失败: ' + error.message);
      }
    }
  },
  created() {
    const token = localStorage.getItem('token');
    if (token) {
      this.token = token;
      this.isLoggedIn = true;
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
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  z-index: 1000;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0,0,0,0.4);
}

.modal-content {
  background-color: #fefefe;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
  max-width: 800px;
  max-height: 80vh;
  overflow-y: auto;
  border-radius: 8px;
  position: relative;
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