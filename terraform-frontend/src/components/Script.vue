<template>
  <div class="d-flex flex-column">
    <button type="button" class="btn btn-primary"  v-on:click="validateScript">Validate script</button>
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
      http.get().then(result => {
      console.log(result.data)
      
      this.serverAnswer = "server: " + result.data + process.env.VUE_APP_BACKEND_IP
    }, error => {
      console.error(error); 
      this.serverAnswer = "Can not connect to server :( "
      this.serverError = error.message + process.env.VUE_APP_BACKEND_IP
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