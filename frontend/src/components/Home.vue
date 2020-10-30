<template>
  <div>
    <h1 style="text-align: center;">{{ msg }}</h1>
    <el-card class="box-card" shadow="never">
      <!--  upload button-->
      <el-button type="primary" size="small" @click="dialogUploadVisible = true">Upload</el-button>
    </el-card>


    <br>
    <el-card class="box-card">
      <!--  uploaded file job list -->
      <el-table
        :data="jobList"
        style="width: 100%"
        @selection-change="handleJobListSelectionChange">
        <el-table-column
          type="selection"
          width="55">
        </el-table-column>
        <el-table-column
          prop="name"
          label="Name">
        </el-table-column>
        <el-table-column
          prop="status"
          label="Status"
          width="160">
        </el-table-column>
        <el-table-column
          prop="action"
          label="Action"
          width="180">
          <template slot-scope="scope">
            <el-button circle size="mini" type="warning" icon="el-icon-tickets"></el-button>
            <el-button circle size="mini" type="primary" icon="el-icon-download"></el-button>
            <el-button circle size="mini" type="success" icon="el-icon-caret-right"></el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog
      title="Upload"
      :visible.sync="dialogUploadVisible"
      width="60%"
      :before-close="handleUploadDialogClose">
      <el-upload
        class="upload-file"
        action="http://localhost:8080/api/file/upload"
        accept="video/*"
        :on-preview="handlePreview"
        :on-remove="handleRemove"
        :before-remove="beforeRemove"
        multiple
        :file-list="fileList">
        <el-button size="small" type="primary">点击上传</el-button>
      </el-upload>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogUploadVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogUploadVisible = false">确 定</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      msg: "Welcome to Mini Transcoder",
      jobList: [
        {
          name: "test1",
          status: "done",
        },
        {
          name: "test1"
        }, {
          name: "test1"
        }, {
          name: "test1"
        }
      ],
      dialogUploadVisible: false,
      fileList:[]
    };
  },
  methods: {
    handleJobListSelectionChange(value) {
      console.log("handleJobListSelectionChange, value:", value);
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview(file) {
      console.log(file);
    },
    handleExceed(files, fileList) {
      this.$message.warning(`当前限制选择 3 个文件，本次选择了 ${files.length} 个文件，共选择了 ${files.length + fileList.length} 个文件`);
    },
    beforeRemove(file, fileList) {
      return this.$confirm(`确定移除 ${ file.name }？`);
    },
    handleUploadDialogClose() {
      console.log("handleUploadDialogClose");
    }
  }
}
</script>

<style scoped>

</style>
