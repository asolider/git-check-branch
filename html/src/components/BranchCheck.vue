<template>
  <div class="inline">
    <el-button type="success" size="mini" @click="dialogVisible = true">切换分支</el-button>

    <el-dialog title="切换分支" :visible.sync="dialogVisible" @open="getProjectBranchList">
      <div class="output">
        <p v-for="msg in messages" class="log" :key="msg">{{ msg }}</p>
      </div>

      <div>
        <el-select v-model="selectbranch" placeholder="选择分支" size="mini">
          <el-option v-for="branch in branchList" :label="branch" :value="branch" :key="branch"></el-option>
        </el-select>

        <el-button
          type="primary"
          size="mini"
          :disabled="checking"
          :loading="checking"
          @click="checkBranch"
        >切换</el-button>
        <el-button type="info" size="mini">清屏</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
//import Socket from "../socket";

export default {
  props: ["projectId"],
  data: function() {
    return {
      dialogVisible: false,
      branchList: [],
      selectbranch: "",
      messages: ["12", "34"],
      checking: false
    };
  },
  mounted: function() {
    //this.getProjectBranchList();
    // Socket.onmessage = this.handleMessage;
  },
  methods: {
    handleMessage(msg) {
      console.log(msg.data);
      this.messages.push(msg.data);
    },
    getProjectBranchList() {
      this.$http
        .get("/branch/list/" + this.projectId)
        .then(response => {
          if (response.status == true) {
            this.branchList = response.data.branch_list;
          }
        })
        .catch(error => {
          console.log(error.message);
        });
    },
    checkBranch() {
      if (this.selectbranch == "") {
        this.$message.error("请先选择分支");
        return false;
      }
      this.checking = true;
      this.$http
        .get("/branch/check/" + this.projectId, {
          params: {
            selectbranch: this.selectbranch
          }
        })
        .then(response => {
          this.checking = false;
          this.messages = this.messages.concat(response.data.output);
          if (response.status == false) {
            this.$message.error(response.msg);
          }
          console.log(this.messages);
        })
        .catch(error => {
          this.checking = false;
          this.$message.error(error.message);
        });
    }
  }
};
</script>


<style scoped>
.inline {
  display: inline-block;
}
.output {
  margin-top: -40px;
  margin-bottom: 5px;
  height: 300px;
  background-color: blanchedalmond;
}
.log {
  height: 12px;
  line-height: 12px;
  font-size: 12px;
  margin: 0px;
}
</style>
