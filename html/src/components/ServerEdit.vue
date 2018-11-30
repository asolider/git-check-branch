<template>
<section id="edit-server">
  <el-button circle type="primary" icon="el-icon-edit" size="mini" @click="dialogFormVisible = true"></el-button>

  <el-dialog title="编辑服务器" :visible.sync="dialogFormVisible" v-loading="loading" @closed="modalClosed" @open='modalOpen'>
    <el-form :model="form">
      <el-form-item label="服务器名" :label-width="formLabelWidth">
        <el-input v-model="form.name" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="IP" :label-width="formLabelWidth">
        <el-input v-model="form.ip" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="PORT" :label-width="formLabelWidth">
        <el-input v-model="form.port" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="USER" :label-width="formLabelWidth">
        <el-input v-model="form.user" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="PASSWD" :label-width="formLabelWidth">
        <el-input v-model="form.passwd" autocomplete="off"></el-input>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" @click="submitEditServer">确 定</el-button>
    </div>
  </el-dialog>
</section>

</template>

<script>
import axios from 'axios'
import qs from 'qs'
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
export default {
  name: 'ServerEdit',
  props: [
    'serverId'
  ],
  data: function() {
    return {
        loading: false,
        dialogFormVisible: false,
        form: [],
        formLabelWidth: '100px'
      };
  },
  methods: {
    getServerInfo: function() {
        axios.get('server/info', {
            params: {
                server_id: this.serverId,
            }
        }).then(response => {
            if (response.data.status == true) {
                this.form = response.data.data
            } else {
                this.$message.error("找不到该信息啊");
                this.form = [];
            }
        });
    },
    submitEditServer: function() {
      this.loading = true;
      axios.post('/server/edit', qs.stringify({
          server_id: this.serverId,
          server_name: this.form.name,
          server_ip : this.form.ip,
          server_port: this.form.port,
          server_user: this.form.user,
          server_passwd: this.form.passwd,
        })
      ).then(response => {
        this.loading=false;
        let data = response.data;
        if (data.status == false) {
          this.$message.success('编辑失败：' + data.msg);
        } else {
          this.$emit('successEdit');
          this.dialogFormVisible = false,
          this.$message.success('编辑成功');
        }
      });
    },
    modalClosed: function() {
      this.form.serverName = '';
      this.form.serverIp = '';
      this.form.loginUser = '';
      this.form.loginPasswd = '';
      this.form.PORT = '';
    },
    modalOpen: function() {
        this.getServerInfo();
    }
  }
}
</script>

<style scoped>
    #edit-server{
        display: inline-block
    }
</style>
