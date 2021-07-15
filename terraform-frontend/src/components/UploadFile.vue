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
<!-- 
    <div class="card">
      <div class="card-header">List of Files</div>
      <ul class="list-group list-group-flush">
        <li class="list-group-item"
          v-for="(file, index) in fileInfos"
          :key="index"
        >
          <a :href="file.url">{{ file.name }}</a>
        </li>
      </ul>
    </div> -->
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
    };
  },
  methods: {
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
