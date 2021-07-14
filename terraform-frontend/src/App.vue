<template>
  <div class="col-lg-8 mx-auto p-3 py-md-5">
    <h1>Here you can load scripts and run them remotely</h1>
    <p class="fs-5 col-md-8"> Just load some .tf configs to validate them! </p>
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
