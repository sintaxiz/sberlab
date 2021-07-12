<template>
  <div id="app">
    <h1>Here you can run terraform scripts! (actually only validate now...)</h1>
    <upload-file/>
    <Script/>
  </div>
</template>

<script>
import Script from './components/Script.vue'
import UploadFile from './components/UploadFile.vue'
export default {
  components: {
    Script,
    UploadFile
  }
}
</script>

<style>
*, *::before, *::after {
  box-sizing: border-box;
}
#app {
  max-width: 400px;
  margin: 0 auto;
  line-height: 1.4;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color:  black;
}
h1 {
  text-align: center;
}


#all {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  color: #e42a2a;
}
</style>