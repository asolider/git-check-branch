<template>
<div>
  <el-button type="text" @click="dialogFormVisible = true">添加项目</el-button>

  <el-dialog title="添加项目" :visible.sync="dialogFormVisible" >
    <el-form :model="form">
      <el-form-item label="所在服务器" :label-width="formLabelWidth">
        <el-input v-model="form.serverIp" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="项目名" :label-width="formLabelWidth">
        <el-input v-model="form.subjectName" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="根目录" :label-width="formLabelWidth">
        <el-input v-model="form.subjectRoot" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="当前分支" :label-width="formLabelWidth">
        <el-input v-model="form.currentBranch" autocomplete="off"></el-input>
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
  name: 'addSubject',
  data: function() {
    return {
        dialogFormVisible: false,
        form: {
          serverIp: '',
          subjectName: '',
          subjectRoot: '',
          currentBranch: '',
        },
        formLabelWidth: '100px'
      };
  }
}
</script>

<style scoped>

</style>
