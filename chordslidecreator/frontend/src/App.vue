<template>
  <div>
  <div class="header">
    <div class="title-bar">
      <h1>Cho Slide Creator</h1>
    </div>
  </div>

  <div class="content">
    <FilePath />
    <Font />
    <Settings />
  </div>
  <div class="content">
    <button class="submit" @click="render()"> Render </button>
  </div>
  </div>
</template>

<script>

import "./assets/css/main.css";
import FilePath from "./components/FilePath";
import Settings from "./components/Settings";
import Font from "./components/Font"

export default {
  name: "app",
  components: {
    Settings,
    FilePath,
    Font
  },
  methods: {
    render(){
      //let fontColor = Font.data()
      //let filePath = FilePath.data()
      let settings = Settings.data()

      let songObject = {
        Name: "",
        PathToFile: FilePath.data(),
        FontPath: "/usr/share/fonts/ubuntu/Ubuntu-M.ttf",

        height: settings.h,
        width: settings.w,
      }
      console.log("object " + songObject)

      window.backend.initSong(songObject)

    }
  }
};
</script>

<style>
  body {
    margin: 0;
    color: white;
    font-family: Ubuntu;
  }

  .header{
    display: flex;
  }

  .title-bar{
    width: 100%;
    text-align: center;
  }

  .content {
    display: flex;
  }
</style>

