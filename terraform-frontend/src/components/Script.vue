<template>
  <div class="d-flex flex-column">
    <div class="mb-5"> 
      <button type="button" class="btn btn-primary btn-lg px-4  "  v-on:click="validateScript">Validate script</button> 
      </div>
    <p class="text-left"> {{serverAnswer}} </p>
    <span>{{serverError}} </span>
  </div>
</template>

<script>
import http from "../http-common";

export default {
  name : "validate-script",
  components: {
    
  },
  data () {
    return {
      serverAnswer: '',
      serverError:''
    }
  },
  created() {
  },
  methods: {
    validateScript() {
      http.get("/v1/validate").then(result => {
      console.log(result.data)
      
      this.serverAnswer = "server: " + result.data
    }, error => {
      console.error(error); 
      this.serverAnswer = "Can not connect to server :( "
      this.serverError = error.message + " " + process.env.VUE_APP_BACKEND_IP
    });
    }
  }
}
</script>

<style> 

html, body {
  margin: 0;
  padding: 0;
}

</style>