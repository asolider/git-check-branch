<template>
<div>
  <addComputer></addComputer>

  <el-table
    :data="tableData"
    stripe>
    <el-table-column
      prop="server_name"
      label="日期">
    </el-table-column>
    <el-table-column
      prop="ip"
      label="姓名">
    </el-table-column>
    <el-table-column
      prop="address"
      label="地址">
    </el-table-column>
  </el-table>
</div>

</template>

<script>
import addComputer from '@/components/AddComputer'
  export default {
    data() {
      return {
        tableData: [{
          server_name: '2016-05-02',
          ip: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄'
        }, {
          server_name: '2016-05-04',
          ip: '王小虎',
          address: '上海市普陀区金沙江路 1517 弄'
        }]
      }
    },
    components: {
      addComputer,
    }
  }
</script>
