<template>
<section style="display: inline-block">
  <el-button circle type="primary" icon="el-icon-edit" size="mini" @click="dialogFormVisible = true"></el-button>

  <el-dialog title="编辑项目" :visible.sync="dialogFormVisible" @open="modalOpen" v-loading="loading" >
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
      <el-button type="primary" @click="submitEditProject">确 定</el-button>
    </div>
  </el-dialog>
</section>

</template>

<script>

export default {
  name: 'addSubject',
  props: [
      'projectId',
  ],
  data: function() {
    return {
        loading: false,
        dialogFormVisible: false,
        serverList: [],
        form: {
          project_id: this.projectId,
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
    getProjectInfo: function (projectId) {
      this.$http.get('/project/info', {
          params: {
              project_id: projectId,
          }
      }).then(response => {
        if (response.status == true) {
          this.form = response.data || [];
        }
      });
    },
    modalOpen: function() {
      this.form.server_id  = '';
      this.form.name  = '';
      this.form.location  = '';
      this.serverList = this.getServerList();
      this.getProjectInfo(this.projectId);
    },
    submitEditProject: function() {
      this.loading = true;
      this.$http.post('/project/edit', this.form
      ).then(response => {
        this.loading=false;
        let data = response;
        if (data.status == false) {
          this.$message.error('编辑失败');
        } else {
          this.$emit('successEdit');
          this.dialogFormVisible = false,
          this.$message.success('编辑成功');
        }
      });
    },
  },

}
</script>

<style scoped>

</style>
