import { createRouter, createWebHistory } from 'vue-router'
import LogViewer from '../components/LogViewer.vue'
import SnForm from '../components/SnForm.vue'

const routes = [
  {
    path: '/logs',
    name: 'Logs',
    component: LogViewer
  },
  {
    path: '/sn',
    name: 'SnForm',
    component: SnForm
  }
]

const router = createRouter({
  history: createWebHistory('/static/'),
  routes
})

export default router