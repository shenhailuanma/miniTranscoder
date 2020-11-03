<template>
  <div>
    <h2 style="text-align: center;">{{ msg }}</h2>

    <!--  upload button-->
    <div style="width:100px;padding:10px">
      <el-button type="primary" size="small" @click="showUploadDialog" style="width:100px;border:20px;">Upload
      </el-button>
    </div>


    <el-card class="box-card">
      <!--  uploaded file job list -->
      <el-table
        :data="jobList"
        style="width: 100%"
        size="small"
        @selection-change="handleJobListSelectionChange">
        <el-table-column
          type="selection"
          width="50">
        </el-table-column>
        <el-table-column
          type="index"
          width="40">
        </el-table-column>
        <el-table-column
          prop="SourceName"
          label="Name">
        </el-table-column>
        <el-table-column
          prop="SourceSize"
          label="Size"
          width="100">
        </el-table-column>
        <el-table-column
          prop="Progress"
          label="Progress"
          width="240">
          <template slot-scope="scope">
            <el-progress v-if="scope.row.Status === 'success'" :percentage="scope.row.Progress" :stroke-width="12"
                         status="success"></el-progress>
            <el-progress v-else-if="scope.row.Status === 'failed'" :percentage="scope.row.Progress" :stroke-width="12"
                         status="exception"></el-progress>
            <el-progress v-else :percentage="scope.row.Progress" :stroke-width="12"></el-progress>

          </template>
        </el-table-column>
        <el-table-column
          prop="action"
          label="Action"
          width="180">
          <template slot-scope="scope">
            <!-- <el-button circle size="mini" type="warning" icon="el-icon-tickets"></el-button>-->
            <el-button circle size="mini" type="primary" icon="el-icon-download"
                       @click="downloadFile(scope.row)"></el-button>
            <el-button circle size="mini" type="warning" icon="el-icon-caret-right"
                       @click="handlePlayVideoSelect(scope.row)"></el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        background
        layout="prev, pager, next"
        :page-size="pageSize"
        :total="jobCount" style="text-align: center;" @current-change="handelCurrentPageChange">
      </el-pagination>
    </el-card>

    <el-dialog
      :visible.sync="dialogUploadVisible"
      width="60%"
      :before-close="handleUploadDialogClose">
      <el-upload
        class="upload-file"
        action="/api/file/upload"
        accept="video/*"
        :on-preview="handleOnPreview"
        :on-remove="handleOnRemove"
        :on-change="handleOnChange"
        :before-remove="beforeRemove"
        multiple
        :file-list="fileList">
        <el-button size="small" type="primary">Select File</el-button>
      </el-upload>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleUploadDialogClose">Cannel</el-button>
        <el-button type="primary" @click="handleUploadJobSubmit">Submit</el-button>
      </span>
    </el-dialog>

    <el-dialog :visible.sync="dialogPlayVideoVisible"
               width="60%"
               :before-close="handlePlayVideoDialogClose">
      <video v-if="playVideoUrl" class="avatar" :src="playVideoUrl" controls="controls"></video>
    </el-dialog>

  </div>
</template>

<script>
import FileSaver from "file-saver";
import {apiGetJobList, apiCreateJobTranscode, apiGetJobsCount} from '@/api/job';
import {objectDeepCopy} from '@/utils/utils';

export default {
  name: "Home",
  data() {
    return {
      msg: "Welcome to Mini Transcoder",
      jobList: [],
      dialogUploadVisible: false,
      dialogPlayVideoVisible: false,
      fileList: [],
      heartbeatTimer: null,
      playVideoUrl: "",
      pageSize: 10,
      jobCount: 100,
      currentPage: 0
    };
  },
  methods: {
    prepareData() {
      // get job list
      apiGetJobList(this.pageSize, this.currentPage - 1).then(response => {
        console.log("prepareData, apiGetJobList response:", response);
        this.jobList = objectDeepCopy(response.data);
        this.outputVideoUrl = this.jobList[0].Output;
      }).catch(err => {
        console.log("prepareData, apiGetJobList err:", err);
      })

      // get jobs count
      apiGetJobsCount().then(response => {
        this.jobCount = response.data;
      }).catch(err => {
        console.log("prepareData, apiGetJobsCount err:", err);
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

      for (let i = 0; i < this.fileList.length; i++) {
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
    },
    handlePlayVideoSelect(row) {
      this.dialogPlayVideoVisible = true;
      this.playVideoUrl = row.Output;
    },
    handlePlayVideoDialogClose() {
      this.dialogPlayVideoVisible = false;
      this.playVideoUrl = "";
    },
    downloadFile(row) {
      console.log("downloadFile, name:", row.SourceName);
      FileSaver.saveAs(row.Output, row.SourceName);
    },
    handelCurrentPageChange(page) {
      console.log("handelCurrentPageChange, page:", page);
      this.currentPage = page;
    }
  },
  mounted() {
    this.prepareData();

    // init cron job
    this.heartbeatTimer = setInterval(() => {
      // console.log("Heart click");
      this.prepareData();
    }, 1500);
  },
  beforeDestroy: function () {
    clearInterval(this.heartbeatTimer);
  },
}
</script>

<style scoped>
.avatar {
  width: 100%;
  display: block;
}
</style>
