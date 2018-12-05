<template>
<div>
  <project-add @successAdd="getProjectList"></project-add>
  <el-table
    :data="tableData"
    stripe>
    <el-table-column
      prop="name"
      label="项目">
    </el-table-column>
    <el-table-column
      prop="server_id"
      label="所在服务器">
    </el-table-column>
    <el-table-column
      prop="location"
      label="地址">
    </el-table-column>
    <el-table-column
      label="操作"
    >
      <template slot-scope="scope">
        <project-edit :projectId="scope.row.project_id"  @successEdit="getProjectList"></project-edit>
        <el-button circle type="danger" icon="el-icon-delete" size="mini" @click="confirmDel(scope.row.project_id)" @successDel="getProjectList"></el-button>
      </template>
    </el-table-column>
  </el-table>
</div>
</template>

<script>
import ProjectAdd from '@/components/ProjectAdd'
import ProjectEdit from '@/components/ProjectEdit'

  export default {
    data() {
      return {
        tableData: [],
      }
    },
    components: {
      ProjectAdd,
      ProjectEdit,
    },
    mounted: function() {
      this.getProjectList();
    },
    methods: {
      getProjectList: function() {
        this.$http.get('/project/list')
        .then(response => {
          if (response.status == true) {
            this.tableData = response.data || [];
          }
        });
      },
      confirmDel: function(project_id) {
        this.$confirm('是否删除?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.projectdel(project_id);
        });
      },
      projectdel: function(project_id) {
        this.$http.get('/project/del', {
          params: {
            project_id: project_id,
          }
        })
        .then(response => {
          if (response.status == false) {
            this.$message.error('删除失败');
          } else {
            this.getProjectList();
            this.$message.success('删除成功');
          }
        })
      },
      successDel: function () {
        this.getProjectList();
      },
      successAdd: function () {
        this.getProjectList();
      },
    },
  }
</script>
