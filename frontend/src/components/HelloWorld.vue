<template>
  <div class="hello">
    <h1>{{ msg }}</h1>

    <div
      style="margin-top: 8px;white-space: pre-line;padding: 8px; background-color: black; font-size:14px; color: white;height: 580px;overflow-y: scroll; width:960px;word-wrap:break-word;text-align:left"
    >{{stdout}}</div>

    <p>stderr:</p>
    {{stderr}}
  </div>
</template>

<script>
import FFmpeg from "ffmpeg.js/ffmpeg-mp4.js";

export default {
  name: "HelloWorld",
  data() {
    return {
      msg: "Welcome to Your Vue.js App",
      stdout: "",
      stderr: ""
    };
  },
  methods: {
    ffmpegVersion() {
      // var ffmpeg = require("ffmpeg.js/ffmpeg-mp4.js");

      var _this = this;
      // Print FFmpeg's version.
      FFmpeg({
        arguments: ["-encoders"],
        print: function(data) {
          _this.stdout += data + "\n";
        },
        printErr: function(data) {
          _this.stderr += data + "\n";
        },
        onExit: function(code) {
          console.log("Process exited with code " + code);
          console.log(_this.stdout);
        }
      });
    }
  },
  mounted: function() {
    console.log("挂载到dom后");
    this.ffmpegVersion();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
