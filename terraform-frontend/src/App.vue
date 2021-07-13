<template>
  <div class="d-flex flex-column">
    <h1 class="display-1">Here you can load scripts and run them remotely</h1>
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
