<template>
  <div>
    <div v-if="progressInfos">
      <div class="mb-2"
        v-for="(progressInfo, index) in progressInfos"
        :key="index"
      >
        <span>{{progressInfo.fileName}}</span>
        <div class="progress">
          <div class="progress-bar progress-bar-info"
            role="progressbar"
            :aria-valuenow="progressInfo.percentage"
            aria-valuemin="0"
            aria-valuemax="100"
            :style="{ width: progressInfo.percentage + '%' }"
          >
            {{progressInfo.percentage}}%
          </div>
        </div>
      </div>
    </div>

      <input type="file" multiple @change="selectFile" />

    <button class="btn btn-success"
      :disabled="!selectFile"
      @click="uploadFiles"
    >
      Upload
    </button>

    <div v-if="message" class="alert alert-light" role="alert">
      <ul>
        <li v-for="(ms, i) in message.split('\n')" :key="i">
          {{ ms }}
        </li>
      </ul>
    </div>

      <div class="mb-5"> 
      <button type="button" class="btn btn-warning btn-lg px-4  "  v-on:click="uploadInfo"> What files are there on the server? </button> 
      </div>
    <p class="text-left"> {{fileNames}} </p>
  </div>
</template>


<script>
import UploadService from "../services/UploadScriptsService";


export default {
  name: "upload-scripts",
  data() {
    return {
      selectedScripts: undefined,
      progressInfos: [],
      message: "",
      scriptInfos: [],
      fileNames: ""
    };
  },
  methods: {
   uploadInfo() {
     UploadService.uploadInfo()
     .then((response) => {
       this.fileNames = response.data
     })
     .catch((error) => {
          this.fileNames = "sorry! could not upload file names... Reason: " + error.message;
        });
   },
   selectFile() {
      this.progressInfos = [];
      this.selectedScripts = event.target.files;
    },
    uploadFiles() {
      this.message = "";

      for (let i = 0; i < this.selectedScripts.length; i++) {
        this.upload(i, this.selectedScripts[i]);
      }
    },
     upload(idx, script) {
      this.progressInfos[idx] = { percentage: 0, scriptname: script.name };

      UploadService.upload(script, (event) => {
        this.progressInfos[idx].percentage = Math.round(100 * event.loaded / event.total);
      })
        .then(() => {
          if (idx != 0) {
          this.message += script.name + " ";

          } else {
          this.message = "succesfull load files on server! c: -- ";
          }
          return
        })
        .catch((error) => {
          this.progressInfos[idx].percentage = 0;
          this.message = "sorry! could not upload the file... Reason: " + error.message;
        });
    },
    // calling when component adding to DOM
    mounted() {
    UploadService.getFiles().then((response) => {
      this.fileInfos = response.data;
    });
  }
  }
};
</script>
