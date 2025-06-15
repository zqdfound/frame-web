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
  </div>
</template>

  
  <script>
  export default {
    data() {
      return {
        sn: '',
        pwd: ''
      }
    },
    methods: {
      async submitForm() {
        try {
          const response = await fetch('http://localhost:8080/api/sn', {
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
          alert(JSON.stringify(result));
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
  </style>