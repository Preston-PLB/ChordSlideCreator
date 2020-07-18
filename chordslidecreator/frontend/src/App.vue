<template>
  <div>
  <div class="header">
    <div class="title-bar">
      <h1>Cho Slide Creator</h1>
    </div>
  </div>

  <div class="content">
    <div class="files">
      <label class="option vertical-flex">
        <p>Select your files</p>
        <input type="text" name="files" placeholder="~/Documents/Songs/overcome.cho" v-model="path"/>
      </label>

    </div>
    <div class="fonts">
      <select>
        <option>Ubuntu</option>
      </select>

      <label class="option">
        <p>Text Color:</p>
        <input type="color" />
      </label>

    </div>
    <div class="settings">
      <form>
        <label class="option">
          <p>Height</p>
          <input type="text" name="height" placeholder="1920" v-model="height">
        </label>
        <label class="option">
          <p>Width</p>
          <input type="text" name="width" placeholder="1080" v-model="width">
        </label>
        <label class="option">
          <p>Nashville Numbers</p>
          <input type="checkbox" name="nashville">
        </label>
      </form>
    </div>
  </div>
  <div class="content">
    <button class="submit" @click="render()"> Render </button>
  </div>
  </div>
</template>

<script>

import "./assets/css/main.css";
//import FilePath from "./components/FilePath";
//import Settings from "./components/Settings";
//import Font from "./components/Font"

export default {
  name: "app",
  // components: {
  //   Settings,
  //   FilePath,
  //   Font
  // },

  props: {
    path: String,
    height: String,
    width: String,
  },
  methods: {
    render(){

      let songObject = {
        Name: "",
        PathToFile: this.$props.path,
        FontPath: "/usr/share/fonts/ubuntu/Ubuntu-M.ttf",

        height: this.$props.height,
        width: this.$props.width,
      }
      console.log("object " + JSON.stringify(songObject))

      window.backend.initSong(songObject)

    }
  }
};
</script>

<style>
  body {
    margin: 0;
    color: aliceblue;
    font-family: Ubuntu;
    background: #555555;
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

  .content div {
    /*border: black 2px solid;*/
  }

  .files{
    flex: 1;
    padding: 5px;
  }

  .fonts{
    flex: 1;
    padding: 5px;
  }

  .fonts select{
    width: 100%;
  }

  .settings{
    flex: 1;
    padding: 5px;
  }

  .option{
    display: flex;
    align-items: center;
    width: 100%;
    margin: 0;
    padding: 0;
  }
  .vertical-flex{
    flex-direction: column;
  }


  .option p {
    flex: 2;
    padding: 0 10px 0 0;
    margin: 0
  }

  .option input{
    flex: 1;
    margin: 0;
  }
</style>

