<template>
  <div class="log-viewer">
    <div class="sidebar">
      <h3>日志文件</h3>
      <ul>
        <li 
          v-for="file in logFiles" 
          :key="file"
          @click="selectFile(file)"
          :class="{ active: selectedFile === file }"
        >
          {{ file }}
        </li>
      </ul>
    </div>
    <div class="content">
      <div class="search-box">
        <input 
          v-model="searchText" 
          placeholder="搜索日志..."
          @input="highlightText"
        />
      </div>
      <div class="log-content" ref="logContent">
        <pre v-html="highlightedContent"></pre>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      logFiles: [],
      selectedFile: '',
      logContent: '',
      searchText: '',
      highlightedContent: '',
      socket: null
    }
  },
  mounted() {
    //this.fetchLogFiles()
    //this.initWebSocket()
  },
  methods: {
    async fetchLogFiles() {
      try {
        const response = await fetch('http://localhost:8080/api/logs')
        const data = await response.json()
        this.logFiles = data.files
      } catch (error) {
        console.error('获取日志文件失败:', error)
      }
    },
    initWebSocket() {
      this.socket = new WebSocket('ws://localhost:8080/ws')
      
      this.socket.onmessage = (event) => {
        this.logContent += event.data + '\n'
        this.highlightText()
      }
      
      this.socket.onclose = () => {
        console.log('WebSocket连接关闭')
      }
    },
    selectFile(file) {
      this.selectedFile = file
      this.logContent = ''
      if (this.socket) {
        this.socket.send(JSON.stringify({ file }))
      }
    },
    highlightText() {
      if (!this.searchText) {
        this.highlightedContent = this.logContent
        return
      }
      
      const regex = new RegExp(this.searchText, 'gi')
      this.highlightedContent = this.logContent.replace(
        regex, 
        match => `<span class="highlight">${match}</span>`
      )
    }
  },
  beforeUnmount() {
    if (this.socket) {
      this.socket.close()
    }
  }
}
</script>

<style scoped>
.log-viewer {
  display: flex;
  height: 100vh;
}

.sidebar {
  width: 250px;
  background: #f5f5f5;
  padding: 20px;
  border-right: 1px solid #ddd;
  overflow-y: auto;
}

.sidebar ul {
  list-style: none;
  padding: 0;
}

.sidebar li {
  padding: 8px 12px;
  cursor: pointer;
  border-radius: 4px;
  margin-bottom: 4px;
}

.sidebar li:hover {
  background: #e9e9e9;
}

.sidebar li.active {
  background: #e0e0e0;
  font-weight: bold;
}

.content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.search-box {
  padding: 10px;
  background: #fff;
  border-bottom: 1px solid #ddd;
}

.search-box input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.log-content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  background: #fff;
}

pre {
  margin: 0;
  white-space: pre-wrap;
  font-family: monospace;
}

.highlight {
  background-color: yellow;
  font-weight: bold;
}
</style>