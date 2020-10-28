<template>
  <div class="Video2gif">
    
    

    <el-row>
      <el-col :span="24">
        
          <video v-if="sourceVideoUrl"
                id="myVideo"
                class="video-js"
                controls
                preload="auto"
                
                data-setup='{}'>
              <source :src="sourceVideoUrl" type="video/mp4"></source>
              <p class="vjs-no-js">
                To view this video please enable JavaScript, and consider upgrading to a
                web browser that
                <a href="http://videojs.com/html5-video-support/" target="_blank">
                  supports HTML5 video
                </a>
              </p>
            </video>
        
      </el-col>
    </el-row>

    <br>
    <el-row>
      <el-col :span="24">
      <el-upload
        class="avatar-uploader"
        action="/"
        :on-success="handleUploadSuccess"
        :before-upload="handleBeforeUpload"
        :on-remove="handleRemove"
        :on-preview="handlePreview"
        :auto-upload="aotoUploadFlag"
        :http-request="uploadFile"
        :before-remove="handleBeforeRemove"
        :on-exceed="handleExceed"
        :limit="1">
        <el-button size="small" type="primary">点击上传</el-button>
        <!-- <el-button type="danger" v-on:click="handleRemove">移除视频</el-button> -->
      </el-upload>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="8">
        <el-time-picker
          is-range
          v-model="duration"
          range-separator="至"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
          placeholder="选择时间范围">
        </el-time-picker>
      </el-col>

      <el-col :span="8">
      <el-input-number v-model="speed" :precision="2" :step="0.1" :max="5"></el-input-number>
      </el-col>
    </el-row>

    
    <el-row>
      <el-col :span="24">
        <el-button size="small" type="primary">生成GIF</el-button>
      </el-col>

      
    </el-row>


    <div v-if="debug"
      style="margin-top: 8px;white-space: pre-line;padding: 8px; background-color: black; font-size:14px; color: white;height: 580px;overflow-y: scroll; width:960px;word-wrap:break-word;text-align:left"
    >{{stdout}}</div>

    <div v-if="debug"
      style="margin-top: 8px;white-space: pre-line;padding: 8px; background-color: black; font-size:14px; color: white;height: 580px;overflow-y: scroll; width:960px;word-wrap:break-word;text-align:left"
    >{{stderr}}</div>

  </div>
</template>

<script>
import { vol, writeFileSync, readFileSync } from "memfs";
import FFmpeg from "ffmpeg.js/ffmpeg-mp4.js";

export default {
  name: "Video2gif",
  data: function() {
    return {
      msg: "Video2gif",
      debug:false,
      aotoUploadFlag: true,
      fileList:[],
      sourceVideoUrl: "",
      sourceName: "",
      sourceData: "",
      outputPictureUrl: "",
      outputName: "",
      stdout: "",
      stderr: "",
      duration:[],
      speed:1
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
      this.sourceVideoUrl = "";
      return true;
    },
    handleAddVideo() {
      console.log("handleAddVideo");
      this.sourceVideoUrl = "";
    },
    handleRemove(file, fileList) {
      console.log("handleRemove");
      this.sourceVideoUrl = "";
    },
    handlePreview(file) {
      console.log(file);
    },
    handleBeforeRemove() {
      console.log("handleBeforeRemove");
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
    },
    initVideo() {
        //初始化视频方法
        let myPlayer = this.$video(myVideo, {
            //确定播放器是否具有用户可以与之交互的控件。没有控件，启动视频播放的唯一方法是使用autoplay属性或通过Player API。
            controls: true,
            //自动播放属性,muted:静音播放
            autoplay: "muted",
            //建议浏览器是否应在<video>加载元素后立即开始下载视频数据。
            preload: "auto",
            //设置视频播放器的显示宽度（以像素为单位）
            width: "800px",
            //设置视频播放器的显示高度（以像素为单位）
            height: "400px"
        });
    }

  },
  mounted: function() {
    console.log("mounted");
    this.initVideo();
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
.video-js {
  width: 800px;
  height: 400px;
}
</style>
