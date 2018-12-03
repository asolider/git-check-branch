<template>
<div>
  <el-button type="text" @click="dialogFormVisible = true">添加服务器</el-button>

  <el-dialog title="添加服务器" :visible.sync="dialogFormVisible" v-loading="loading" @closed="modalClosed">
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
      <el-button type="primary" @click="submitAddServer">确 定</el-button>
    </div>
  </el-dialog>
</div>

</template>

<script>
import axios from 'axios'
import qs from 'qs'
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
export default {
  name: 'addComputer',
  data: function() {
    return {
        loading: false,
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
  },
  methods: {
    submitAddServer: function() {
      this.loading = true;
      axios.post('/server/add', qs.stringify({server_name: this.form.serverName,
          server_ip : this.form.serverIp,
          server_port: this.form.serverPort,
          server_user: this.form.loginUser,
          server_passwd: this.form.loginPasswd,
        })
      ).then(response => {
        this.loading=false;
        let data = response.data;
        if (data.status == false) {
          this.$message.error('添加失败');
        } else {
          let serverId = data.data.insert_id
          this.$emit('successAdd', {server_id: serverId, name: this.form.serverName, ip : this.form.serverIp, port: this.form.serverPort, user: this.form.loginUser});
          this.dialogFormVisible = false,
          this.$message.success('添加成功');
        }
      });
    },
    modalClosed: function() {
      this.form.serverName = '';
      this.form.serverIp = '';
      this.form.loginUser = '';
      this.form.loginPasswd = '';
    },
  }
}
</script>

<style scoped>

</style>
