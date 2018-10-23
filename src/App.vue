<template>
  <div id="app">
    <el-row>
      <el-col :span="16" :offset="4">
        <mheader></mheader>
        <router-view/>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Mheader from '@/components/Header'

export default {
  name: 'App',
  components: {
    Mheader,
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
