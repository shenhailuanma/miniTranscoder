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
          prop="ID"
          label="ID"
          width="60">
        </el-table-column>
        <el-table-column
          prop="SourceName"
          label="SourceName">
        </el-table-column>
        <el-table-column
          prop="Description"
          label="Description">
        </el-table-column>
        <el-table-column
          prop="SourceSize"
          label="SourceSize"
          align="center"
          width="100">
          <template slot-scope="scope">
            {{ transFilesize(scope.row.SourceSize) }}
          </template>
        </el-table-column>
        <el-table-column
          prop="OutputSize"
          label="OutputSize"
          align="center"
          width="100">
          <template slot-scope="scope">
            {{ transFilesize(scope.row.OutputSize) }}
          </template>
        </el-table-column>
        <el-table-column
          prop="Progress"
          label="Progress"
          align="center"
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
          prop="Publish"
          label="Publish"
          align="center"
          width="80">
          <template slot-scope="scope">
            <i v-if="scope.row.Publish" class="el-icon-success" style="color: green; font-size: 18px"></i>
            <i v-else class="el-icon-circle-close" style="font-size: 18px"></i>
          </template>
        </el-table-column>
        `
        <el-table-column
          prop="action"
          label="Action"
          align="center"
          width="220">
          <template slot-scope="scope">
            <el-button circle size="mini" type="warning" icon="el-icon-notebook-2"
                       @click="handleShowLogSelect(scope.row)"></el-button>
            <el-button v-if="scope.row.Status === 'success'" circle size="mini" type="success" icon="el-icon-download"
                       @click="downloadFile(scope.row)"></el-button>
            <el-button v-if="scope.row.Status === 'success'" circle size="mini" type="success"
                       icon="el-icon-caret-right"
                       @click="handlePlayVideoSelect(scope.row)"></el-button>
            <el-button v-if="scope.row.Status === 'success'" circle size="mini" type="primary" icon="el-icon-setting"
                       @click="handleUpdateJob(scope.row)"></el-button>
            <el-button v-if="scope.row.Status === 'success' || scope.row.Status === 'failed'" circle size="mini" type="danger" icon="el-icon-delete"
                       @click="handleRemoveJob(scope.row)"></el-button>
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

      <div v-if="fileList.length > 0">
        <el-divider content-position="left">Configurations</el-divider>
        <el-form label-position="right" label-width="100px" :model="userTransParams">
          <el-form-item label="VideoSize">
            <el-radio-group v-model="userTransParams.height">
              <el-radio label="0">Origin</el-radio>
              <el-radio label="2160">4K(2160P)</el-radio>
              <el-radio label="1440">2K(1440P)</el-radio>
              <el-radio label="1080">1080P</el-radio>
              <el-radio label="720">720P</el-radio>
              <el-radio label="480">480P</el-radio>
              <el-radio label="360">360P</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="Compression">
            <el-radio-group v-model="userTransParams.preset">
              <el-radio label="veryslow">High</el-radio>
              <el-radio label="medium">Normal</el-radio>
              <el-radio label="veryfast">Low</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-form>
      </div>

      <span slot="footer" class="dialog-footer">
        <el-button @click="handleUploadDialogClose">Cancel</el-button>
        <el-button :disabled="fileList.length === 0" type="primary" @click="handleUploadJobSubmit">Submit</el-button>
      </span>
    </el-dialog>

    <el-dialog :visible.sync="dialogShowLogVisible"
               title="Log"
               width="80%"
               top="2vh"
               :before-close="handleShowLogDialogClose">
      <div>
        <codemirror v-model="jobLogData" :options="fileRawDataCmOptions"></codemirror>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="dialogPlayVideoVisible"
               width="60%"
               top="2vh"
               :before-close="handlePlayVideoDialogClose">
      <video v-if="playVideoUrl" class="avatar" :src="playVideoUrl" controls="controls"></video>
    </el-dialog>

    <el-dialog :visible.sync="dialogRemoveJobVisible"
               title="Remove"
               width="60%"
               top="2vh">
      <p>Are you sure to remove all files of this job ? </p>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogRemoveJobVisible = false">Cancel</el-button>
        <el-button type="primary" @click="doRemoveJob()">Submit</el-button>
      </span>
    </el-dialog>

    <el-dialog :visible.sync="dialogUpdateJobVisible"
               title="Update"
               width="50%"
               top="2vh">
      <el-form ref="updateJobForm" :model="updateJobForm" label-width="100px" size="mini">
        <el-form-item label="ID">
          <el-input v-model="updateJobForm.ID" disabled></el-input>
        </el-form-item>
        <el-form-item label="SourceName">
          <el-input v-model="updateJobForm.SourceName" disabled></el-input>
        </el-form-item>
        <el-form-item label="Description">
          <el-input v-model="updateJobForm.Description"></el-input>
        </el-form-item>
        <el-form-item label="File">
          <el-input v-model="updateJobForm.RelativePath"></el-input>
        </el-form-item>
        <el-form-item label="Snapshot">
          <el-input v-model="updateJobForm.Snapshot"></el-input>
        </el-form-item>
        <el-form-item label="Publish">
          <el-switch v-model="updateJobForm.Publish"></el-switch>
        </el-form-item>

      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogUpdateJobVisible = false">Cancel</el-button>
        <el-button type="primary" @click="doUpdateJob()">Submit</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { codemirror } from "vue-codemirror";
import "codemirror/lib/codemirror.css";

import FileSaver from "file-saver";
import {apiGetJobList, apiCreateJobTranscode, apiRemoveJob, apiUpdateJob, apiGetJobData} from '@/api/job';
import {objectDeepCopy} from '@/utils/utils';

export default {
  name: "Home",
  components: {
    codemirror,
  },
  data() {
    return {
      msg: "Mini Transcoder",
      jobList: [],
      dialogUploadVisible: false,
      dialogShowLogVisible: false,
      dialogPlayVideoVisible: false,
      dialogRemoveJobVisible: false,
      dialogUpdateJobVisible: false,
      fileList: [],
      heartbeatTimer: null,
      playVideoUrl: "",
      jobLogData: "",
      fileRawDataCmOptions: {
        tabSize: 4,
        mode: "text/x-go",
        theme: "default",
        styleActiveLine: true,
        lineNumbers: true,
        line: true,
        readOnly: true,
      },
      selectedRow: {},
      jobCount: 0,
      userTransParams: {
        preset: "veryslow",
        height: "0",
      },
      transParams: {
        inputs: [],
        outputs: [
          {
            format: "mp4",
            streams: [
              {
                kind: "video",
                video: {
                  codec: "h264",
                  preset: "veryslow",
                  fps: 25,
                  height: 0,
                }
              },
              {
                kind: "audio",
                audio: {
                  codec: "aac",
                  channels: 2,
                  sample_rate: 44100,
                }
              }
            ]
          }
        ]
      },
      updateJobForm: {},
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
      return (unit === 'B' ? size : size.toFixed(pointLength)) + unit;
    },
    handleJobListSelectionChange(value) {
      console.log("handleJobListSelectionChange, value:", value);
    },
    showUploadDialog() {
      this.dialogUploadVisible = true;
      this.fileList = [];
      this.userTransParams = {
        preset: "veryslow",
        height: "0",
      };
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

        jobData.outputs[0].streams[0].video.preset = this.userTransParams.preset;
        jobData.outputs[0].streams[0].video.height = Number(this.userTransParams.height);

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
    handleShowLogDialogClose() {
      this.dialogShowLogVisible = false;
      this.jobLogData = "";
    },
    handleShowLogSelect(row) {
      this.dialogShowLogVisible = true;
      this.jobLogData = "";
      // get log
      apiGetJobData(row.ID, "job.log").then(response => {
        console.log("handleShowLogSelect, apiGetJobData, response:", response)
        this.jobLogData = response;
      }).catch(err => {
        console.log("handleShowLogSelect, apiGetJobData, err:", err)
      })

    },
    handlePlayVideoSelect(row) {
      this.dialogPlayVideoVisible = true;
      this.playVideoUrl = row.RelativePath;
    },
    handlePlayVideoDialogClose() {
      this.dialogPlayVideoVisible = false;
      this.playVideoUrl = "";
    },
    handleRemoveJob(row) {
      this.selectedRow = row;
      this.dialogRemoveJobVisible = true;
    },
    doRemoveJob() {
      this.dialogRemoveJobVisible = false;
      console.log("doRemoveJob, ID:", this.selectedRow.ID);
      apiRemoveJob(this.selectedRow.ID).then(response => {
        console.log("doRemoveJob, ID:", this.selectedRow.ID, ", success");
      }).catch(err => {
        console.log("doRemoveJob, ID:", this.selectedRow.ID, ", error:", err);
      })
    },
    handleUpdateJob(row) {
      this.selectedRow = row;
      this.updateJobForm = objectDeepCopy(row);

      this.dialogUpdateJobVisible = true;
    },
    doUpdateJob() {
      this.dialogUpdateJobVisible = false;
      console.log("doUpdateJob, updateJobForm:", this.updateJobForm);
      apiUpdateJob(this.updateJobForm.ID, this.updateJobForm).then(response => {
        console.log("doUpdateJob, ID:", this.updateJobForm.ID, ", success");
      }).catch(err => {
        console.log("doUpdateJob, ID:", this.updateJobForm.ID, ", error:", err);
      })
    },
    downloadFile(row) {
      console.log("downloadFile, Output file:", row.Output);
      let name = row.Output.substring(row.Output.lastIndexOf("/") + 1)

      name = `${row.ID}_${name}`;
      console.log("downloadFile, Output file save as name:", name);
      FileSaver.saveAs(row.RelativePath, name);
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
