import Vue from 'vue'
import Router from 'vue-router'
import Computer from '@/components/Computer'
import Subject from '@/components/Subject'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'server-list',
      component: Computer
    },
    {
      path: '/subject',
      name: 'subject-list',
      component: Subject
    }
  ]
})
