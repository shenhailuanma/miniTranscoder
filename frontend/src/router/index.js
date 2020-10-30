import Vue from "vue";
import Router from "vue-router";
import Home from "@/components/Home";
// import Snapshot from "@/components/Snapshot";
import Transcode from "@/components/Transcode";
// import Video2gif from "@/components/Video2gif";
// import Test from "@/components/Test";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "Home",
      component: Home
    },
    {
      path: "/video",
      name: "video",
      component: Transcode
    }
  ]
});
