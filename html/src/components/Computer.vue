<template>
<div>
  <addComputer @successAdd="successAdd"></addComputer>

  <el-table
    :data="tableData"
    stripe>
    <el-table-column
      prop="name"
      label="服务器">
    </el-table-column>
    <el-table-column
      prop="ip"
      label="IP">
    </el-table-column>
    <el-table-column
      prop="port"
      label="PORT">
    </el-table-column>
    <el-table-column
      prop="user"
      label="USER">
    </el-table-column>
    <el-table-column
      label="操作"
    >
      <template slot-scope="scope">
        <server-edit :server-id="scope.row.server_id"  @successEdit="getServerList"></server-edit>
        <el-button circle type="danger" icon="el-icon-delete" size="mini" @click="confirmDel(scope.row.server_id)" ></el-button>
      </template>
    </el-table-column>
  </el-table>
</div>

</template>

<script>
import addComputer from '@/components/AddComputer'
import ServerEdit from '@/components/ServerEdit'
import axios from 'axios'
import qs from 'qs'

  export default {
    data() {
      return {
        tableData: [],
      }
    },
    components: {
      addComputer,
      ServerEdit
    },
    mounted: function() {
      this.getServerList();
    },
    methods: {
      getServerList: function() {
        axios.get('/server/list')
        .then(response => {
          if (response.data.status == true) {
            this.tableData = response.data.data || [];
          }
        });
      },
      successAdd: function(server) {
        this.tableData.push(server)
      },
      confirmDel: function(server_id) {
        this.$confirm('是否删除?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.serverdel(server_id);
        });
      },
      serverdel: function(server_id) {
        axios.get('/server/del', {
          params: {
            server_id: server_id,
          }
        })
        .then(response => {
          console.log(response)
          if (response.data.status == false) {
            this.$message.error('删除失败');
          } else {
            this.tableData = this.tableData.filter(function (server){
              if (server.server_id != server_id) {
                return server;
              }
            });
            this.$message.success('删除成功');
          }
        })
      },
      
    }
  }
</script>
