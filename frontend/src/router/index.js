import Vue from "vue";
import Router from "vue-router";
import HelloWorld from "@/components/HelloWorld";
import Snapshot from "@/components/Snapshot";
import Transcode from "@/components/Transcode";
import Video2gif from "@/components/Video2gif";
import Test from "@/components/Test";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "HelloWorld",
      component: HelloWorld
    },
    {
      path: "/snapshot",
      name: "Snapshot",
      component: Snapshot
    },
    {
      path: "/transcode",
      name: "Transcode",
      component: Transcode
    },
    {
      path: "/video2gif",
      name: "Video2gif",
      component: Video2gif
    },
    {
      path: "/test",
      name: "Test",
      component: Test
    }
  ]
});
