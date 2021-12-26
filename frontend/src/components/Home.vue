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
          label="SourceName">
        </el-table-column>
        <el-table-column
          prop="SourceSize"
          label="SourceSize"
          width="100">
          <template slot-scope="scope">
            {{ transFilesize(scope.row.SourceSize) }}
          </template>
        </el-table-column>
        <el-table-column
          prop="OutputSize"
          label="OutputSize"
          width="100">
          <template slot-scope="scope">
            {{ transFilesize(scope.row.OutputSize) }}
          </template>
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
            <el-button circle size="mini" type="danger" icon="el-icon-delete"
                       @click="handlePlayVideoSelect(scope.row)"></el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog
      :visible.sync="dialogUploadVisible"
      width="60%"
      top="2vh"
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

<!--      <div>-->
<!--        <el-divider content-position="left">Configurations</el-divider>-->
<!--        <el-radio-group v-model="userTransParams.preset">-->
<!--          <el-radio label="veryslow">High</el-radio>-->
<!--          <el-radio label="medium">Normal</el-radio>-->
<!--        </el-radio-group>-->
<!--      </div>-->

      <span slot="footer" class="dialog-footer">
        <el-button @click="handleUploadDialogClose">Cannel</el-button>
        <el-button :disabled="fileList.length === 0" type="primary" @click="handleUploadJobSubmit">Submit</el-button>
      </span>
    </el-dialog>

    <el-dialog :visible.sync="dialogPlayVideoVisible"
               width="60%"
               top="2vh"
               :before-close="handlePlayVideoDialogClose">
      <video v-if="playVideoUrl" class="avatar" :src="playVideoUrl" controls="controls"></video>
    </el-dialog>

  </div>
</template>

<script>
import FileSaver from "file-saver";
import {apiGetJobList, apiCreateJobTranscode} from '@/api/job';
import {objectDeepCopy} from '@/utils/utils';

export default {
  name: "Home",
  data() {
    return {
      msg: "Mini Transcoder",
      jobList: [],
      dialogUploadVisible: false,
      dialogPlayVideoVisible: false,
      fileList: [],
      heartbeatTimer: null,
      playVideoUrl: "",
      jobCount: 0,
      userTransParams: {
        preset:"veryslow"
      },
      transParams: {
        inputs:[],
        outputs:[
          {
            format:"mp4",
            streams:[
              {
                kind: "video",
                video: {
                  codec: "h264",
                  preset: "veryslow"
                }
              },
              {
                kind: "audio",
                audio: {
                  codec: "aac",
                  channels:2
                }
              }
            ]
          }
        ]
      }
    };
  },
  methods: {
    prepareData() {
      // get job list
      apiGetJobList().then(response => {
        console.log("prepareData, apiGetJobList response:", response);
        this.jobList = objectDeepCopy(response.data);
        this.jobCount = this.jobList.length;
      }).catch(err => {
        console.log("prepareData, apiGetJobList err:", err);
      })
    },
    transFilesize(size) {
      let pointLength = 1
      let units = ["B", "K", "M", "G", "TB"];
      let unit = "B"
      while ((unit = units.shift()) && size > 1024) {
        size = size / 1024;
      }
      return (unit === 'B' ? size : size.toFixed( pointLength)) + unit;
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
        let jobData = objectDeepCopy(this.transParams);
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
      console.log("downloadFile, Output file:", row.Output);
      let name = row.Output.substring(row.Output.lastIndexOf("/")+1)

      name = `${row.ID}_${name}`;
      console.log("downloadFile, Output file save as name:", name);
      FileSaver.saveAs(row.Output, name);
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
