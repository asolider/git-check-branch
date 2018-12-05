<template>
<div>
  <el-button type="text" @click="dialogFormVisible = true">添加项目</el-button>

  <el-dialog title="添加项目" :visible.sync="dialogFormVisible" @open="modalOpen" v-loading="loading" >
    <el-form :model="form" :label-width="formLabelWidth">
      <el-form-item label="所在服务器">
        <el-select v-model="form.server_id" placeholder="请选择服务器">
          <el-option v-for="(server) in serverList" :label="server.name" :value="server.server_id" :key="server.server_id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="项目名" >
        <el-input v-model="form.name" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="根目录" >
        <el-input v-model="form.location" autocomplete="off"></el-input>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" @click="submitAddProject">确 定</el-button>
    </div>
  </el-dialog>
</div>

</template>

<script>

export default {
  name: 'addSubject',
  data: function() {
    return {
        loading: false,
        dialogFormVisible: false,
        serverList: [],
        form: {
          server_id: '',
          name: '',
          location: '',
        },
        formLabelWidth: '100px'
      };
  },
  mounted: function() {

  },
  methods: {
    getServerList: function() {
      this.$http.get('/server/list')
      .then(response => {
        if (response.status == true) {
          this.serverList = response.data || [];
        }
      });
    },
    modalOpen: function() {
      this.form.server_id  = '';
      this.form.name  = '';
      this.form.location  = '';
      this.serverList = this.getServerList();
    },
    submitAddProject: function() {
      this.loading = true;
      this.$http.post('/project/add', this.form
      ).then(response => {
        this.loading=false;
        let data = response;
        if (data.status == false) {
          this.$message.error('添加失败');
        } else {
          this.$emit('successAdd');
          this.dialogFormVisible = false,
          this.$message.success('添加成功');
        }
      });
    },
  },

}
</script>

<style scoped>

</style>
