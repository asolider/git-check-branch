<template>
<el-menu
  :default-active="activeIndex"
  class=""
  mode="horizontal"
  @select="handleSelect"
  background-color="#545c64"
  text-color="#fff"
  active-text-color="#ffd04b">
  <el-menu-item index="1"><router-link :to="{name:'server-list'}">机器管理</router-link> </el-menu-item>
  
  <el-menu-item index="2"><router-link :to="{name:'subject-list'}">项目管理</router-link> </el-menu-item>
</el-menu>
</template>

<script>
export default {
  data: function () {
    return {
      activeIndex: "1"
    }
  },
  name: "Header",
  methods: {
    handleSelect: function () {}  
  }
}
</script>

<style scoped>

</style>

