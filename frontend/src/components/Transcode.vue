<template>
  <div class="Snapshot">
    <el-row>
      <el-col :span="11">
        <el-card class="box-card">
          <el-upload
            class="avatar-uploader"
            action="/"
            :show-file-list="false"
            :on-success="handleUploadSuccess"
            :before-upload="handleBeforeUpload"
            :on-remove="handleRemove"
            :on-preview="handlePreview"
            :auto-upload="aotoUploadFlag"
            :http-request="uploadFile"
          >
            <!-- <img v-if="sourceVideoUrl" :src="sourceVideoUrl" class="avatar"> -->
            <video v-if="sourceVideoUrl" :src="sourceVideoUrl" controls="controls" class="avatar"></video>
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-card>
      </el-col>
      <el-col :span="2">
        <div class="grid-content bg-purple-light">
          <br>
          <el-button v-on:click="doTranscodeGif" size="small" type="primary" round>gif</el-button>
          <br>
          <el-button size="small" type="primary" round>截图</el-button>
        </div>
      </el-col>
      <el-col :span="11">
        <el-card class="box-card">
          <img v-if="outputPictureUrl" :src="outputPictureUrl" class="avatar">
          <i v-else class="el-icon-loading avatar-uploader-icon"></i>
        </el-card>
      </el-col>
    </el-row>

    <br>
    <el-button v-on:click="handleButton" type="warning" icon="el-icon-star-off" circle></el-button>

    <div
      style="margin-top: 8px;white-space: pre-line;padding: 8px; background-color: black; font-size:14px; color: white;height: 580px;overflow-y: scroll; width:960px;word-wrap:break-word;text-align:left"
    >{{stdout}}</div>

    <div
      style="margin-top: 8px;white-space: pre-line;padding: 8px; background-color: black; font-size:14px; color: white;height: 580px;overflow-y: scroll; width:960px;word-wrap:break-word;text-align:left"
    >{{stderr}}</div>
  </div>
</template>

<script>
import { vol, writeFileSync, readFileSync } from "memfs";
import FFmpeg from "ffmpeg.js/ffmpeg-mp4.js";

export default {
  name: "Snapshot",
  data: function() {
    return {
      msg: "Snapshot",
      aotoUploadFlag: true,
      sourceVideoUrl: "",
      sourceName: "",
      sourceData: "",
      outputPictureUrl: "",
      outputName: "",
      stdout: "",
      stderr: ""
    };
  },
  methods: {
    handleUploadSuccess(res, file) {
      console.log("handleUploadSuccess");
      console.log(res);
      console.log(file);
      this.sourceVideoUrl = URL.createObjectURL(file.raw);
      console.log("sourceVideoUrl:" + this.sourceVideoUrl);
    },
    handleBeforeUpload() {
      console.log("handleBeforeUpload");
      return true;
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
      this.sourceVideoUrl = "";
    },
    handlePreview(file) {
      console.log(file);
    },
    uploadFile(params) {
      console.log("uploadFile", params);
      this.sourceVideoUrl = URL.createObjectURL(params.file);
      console.log("sourceVideoUrl:" + this.sourceVideoUrl);
      console.log("file:");

      var reader = new FileReader();
      var _this = this;

      reader.onloadstart = function() {
        console.log("FileReader onloadstart");
      };
      reader.onprogress = function() {
        console.log("FileReader onprogress");
      };
      reader.onloadend = function() {
        console.log("FileReader onloadend");
      };
      reader.onprogress = function() {
        console.log("FileReader onprogress");
      };
      reader.onabort = function() {
        console.log("FileReader onabort");
      };

      reader.onload = function(evt) {
        console.log("FileReader onload");
        console.log(evt.target);
        // console.log(evt.target.result.byteLength);

        console.log(evt.target.result);

        // let sourceData = new Uint8Array(evt.target.result);
        _this.sourceData = new Uint8Array(evt.target.result);
        // writeFileSync("source.mp4", evt.target.result);
      };
      //   reader.readAsBinaryString(params.file);
      //   reader.readAsText(params.file, "UTF-8");
      reader.readAsArrayBuffer(params.file);
    },
    handleButton() {
      console.log("handleButton");
      //   writeFileSync("/text.txt", "Hello world!");
      //   console.log(readFileSync("/text.txt", "utf8"));
      //   console.log(vol.toJSON());

      var _this = this;
      this.stdout = "";
      this.stderr = "";

      console.log(_this.sourceData);

      //   var sourceData = new Uint8Array(readFileSync("source.mp4"));
      //   var sourceData = new Uint8Array(this.sourceData);

      //   arguments: ["-i", "source.mp4", "-c", "copy", "output.mp4"],

      var result = FFmpeg({
        MEMFS: [{ name: "source", data: _this.sourceData }],
        arguments: [
          "-nostdin",
          "-i",
          "source",
          "-c:v",
          "h264",
          "-f",
          "mp4",
          "-y",
          "output.mp4"
        ],
        print: function(data) {
          console.log("stdio:" + data);
          _this.stdout += data + "\n";
        },
        printErr: function(data) {
          console.log("stderr:" + data);
          _this.stderr += data + "\n";
        },
        onExit: function(code) {
          console.log("Process exited with code " + code);
        },
        stdin: function() {
          console.log("stdin:");
        }
      });

      console.log("result:");
      console.log(result);

      // Write out.webm to disk.
      //   var output = result.MEMFS[0];
      //   console.log("output.name:" + output.name);
      //   writeFileSync(output.name, Buffer(output.data));
    },
    doTranscodeGif() {
      console.log("handleButton");

      var _this = this;
      this.stdout = "";
      this.stderr = "";

      var result = FFmpeg({
        MEMFS: [{ name: "source", data: _this.sourceData }],
        arguments: ["-nostdin", "-i", "source", "-r", "5", "-y", "output.gif"],
        print: function(data) {
          // console.log("stdio:" + data);
          _this.stdout += data + "\n";
        },
        printErr: function(data) {
          console.log(data);
          _this.stderr += data + "\n";
        },
        onExit: function(code) {
          console.log("Process exited with code " + code);
        },
        stdin: function() {
          console.log("stdin:");
        }
      });

      console.log("result:");
      console.log(result);

      // Write to disk.
      if (result.MEMFS.length > 0) {
        var output = result.MEMFS[0];
        console.log("output.name:" + output.name);
        _this.outputPictureUrl =
          "data:image/png;base64," + _this.arrayBufferToBase64(output.data);
        writeFileSync(_this.outputPictureUrl, Buffer(output.data));
      }
      //   var output = result.MEMFS[0];
      //   console.log("output.name:" + output.name);
      //   writeFileSync(output.name, Buffer(output.data));
    },
    doSnapshot() {
      console.log("doSnapshot");

      var _this = this;
      this.stdout = "";
      this.stderr = "";

      var result = FFmpeg({
        MEMFS: [{ name: "source", data: _this.sourceData }],
        arguments: [
          "-nostdin",
          "-i",
          "source",
          "-c:v",
          "h264",
          "-f",
          "mp4",
          "-y",
          "output.mp4"
        ],
        print: function(data) {
          console.log("stdio:" + data);
          _this.stdout += data + "\n";
        },
        printErr: function(data) {
          console.log("stderr:" + data);
          _this.stderr += data + "\n";
        },
        onExit: function(code) {
          console.log("Process exited with code " + code);
        },
        stdin: function() {
          console.log("stdin:");
        }
      });
    },
    arrayBufferToBase64(buffer) {
      let binary = "";
      let bytes = new Uint8Array(buffer);
      let len = bytes.byteLength;
      for (let i = 0; i < len; i++) {
        binary += String.fromCharCode(bytes[i]);
      }
      return window.btoa(binary);
    }
  },
  mounted: function() {
    console.log("mounted");
  },
  activated: function() {
    console.log("activated");
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
