<template>
<div>
  <el-button type="text" @click="dialogFormVisible = true">添加服务器</el-button>

  <el-dialog title="添加服务器" :visible.sync="dialogFormVisible" >
    <el-form :model="form">
      <el-form-item label="服务器名" :label-width="formLabelWidth">
        <el-input v-model="form.serverName" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="IP" :label-width="formLabelWidth">
        <el-input v-model="form.serverIp" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="PORT" :label-width="formLabelWidth">
        <el-input v-model="form.serverPort" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="USER" :label-width="formLabelWidth">
        <el-input v-model="form.loginUser" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="PASSWD" :label-width="formLabelWidth">
        <el-input v-model="form.loginPasswd" autocomplete="off"></el-input>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
    </div>
  </el-dialog>
</div>

</template>

<script>
export default {
  name: 'addComputer',
  data: function() {
    return {
        dialogFormVisible: false,
        form: {
          serverName: '',
          serverIp: '',
          serverPort: '22',
          loginUser: '',
          loginPasswd: ''
        },
        formLabelWidth: '100px'
      };
  }
}
</script>

<style scoped>

</style>
