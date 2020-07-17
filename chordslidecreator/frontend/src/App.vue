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
  }
};
</script>

<style>
  body {
    margin: 0;
    color: aliceblue;
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

