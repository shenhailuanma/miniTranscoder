<template>
  <div>
    <h1 style="text-align: center;">{{ msg }}</h1>
    <el-card class="box-card" shadow="never">
      <!--  upload button-->
      <el-button type="primary" size="small" @click="showUploadDialog">Upload</el-button>
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
          prop="size"
          label="Size"
          width="100">
        </el-table-column>
        <el-table-column
          prop="progress"
          label="Progress"
          width="240">
          <template slot-scope="scope">
            <el-progress :percentage="100" :stroke-width="12" status="success"></el-progress>
          </template>
        </el-table-column>
        <!--        <el-table-column-->
        <!--          prop="status"-->
        <!--          label="Status"-->
        <!--          width="160">-->
        <!--        </el-table-column>-->
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
      :visible.sync="dialogUploadVisible"
      width="60%"
      :before-close="handleUploadDialogClose">
      <el-upload
        class="upload-file"
        action="http://localhost:8080/api/file/upload"
        accept="video/*"
        :on-preview="handleOnPreview"
        :on-remove="handleOnRemove"
        :on-change="handleOnChange"
        :before-remove="beforeRemove"
        multiple
        :file-list="fileList">
        <el-button size="small" type="primary">点击上传</el-button>
      </el-upload>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleUploadDialogClose">Cannel</el-button>
        <el-button type="primary" @click="handleUploadJobSubmit">Submit</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>

import {apiGetJobList, apiCreateJobTranscode} from '@/api/job';

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
      fileList: []
    };
  },
  methods: {
    prepareData() {
      apiGetJobList().then(response => {
        console.log("prepareData, apiGetJobList response:", response);
        this.jobList = response.data;
      }).catch(err => {
        console.log("prepareData, apiGetJobList err:", err);
      })
    },
    handleJobListSelectionChange(value) {
      console.log("handleJobListSelectionChange, value:", value);
    },
    showUploadDialog() {
      this.dialogUploadVisible = true;
      this.fileList = [];
    },
    handleOnRemove(file, fileList) {
      console.log("handleOnRemove, file:", file);
      console.log("handleOnRemove, fileList:", fileList);
      this.fileList = fileList;
    },
    handleOnPreview(file) {
      console.log("handleOnPreview, file:", file);
    },
    handleOnChange(file, fileList) {
      console.log("handleOnChange, file:", file);
      console.log("handleOnChange, fileList:", fileList);
      this.fileList = fileList;
    },
    beforeRemove(file, fileList) {
      return this.$confirm(`Are you sure to remove ${file.name}？`);
    },
    handleUploadDialogClose() {
      console.log("handleUploadDialogClose");
      this.dialogUploadVisible = false;
      this.fileList = [];
    },
    handleUploadJobSubmit() {
      console.log("handleUploadJobSubmit, fileList:", this.fileList);

      for (let i=0; i < this.fileList.length; i++) {
        let jobData = {};
        jobData.inputs = [];
        jobData.inputs.push(this.fileList[i].name);

        // create job
        apiCreateJobTranscode(jobData).then(response => {
          console.log("handleUploadJobSubmit, apiCreateJobTranscode response:", response);
          // reload jobs
          this.prepareData();
        }).catch(err => {
          console.log("handleUploadJobSubmit, apiCreateJobTranscode err:", err);
        })
      }

      this.dialogUploadVisible = false;
    }
  },
  mounted() {
    this.prepareData();
  }
}
</script>

<style scoped>

</style>
