<template>
  <div id="app">
    <div id="content">
      <mheader></mheader>
      <router-view/>
    </div>
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
#content {
  width: 60%;
  margin: 0 auto;
}
</style>
