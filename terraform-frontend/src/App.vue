<template>
  <div id="app">
    <h1>Here you can run terraform scripts! (actually only validate now...)</h1>
    <Script/>
  </div>
</template>

<script>
import Script from './components/Script.vue'
export default {
  components: {
    Script
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
</style>