<template>
  <div>
    <div>
      <button v-on:click="askServer">Validate script</button> 
    </div>
    <span>{{serverAnswer}} </span>
  <div>
    <span>{{serverError}} </span>
  </div>
  </div>
</template>

<script>
import axios from "axios"
const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_API
});
export default {
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
    askServer() {
      axios_instance.get().then(result => {
      console.log(result.data)
      
      this.serverAnswer = "server: " + result.data
    }, error => {
      console.error(error); 
      this.serverAnswer = "Can not connect to server :( "
      this.serverError = error.message
    });
    }
  }
}
</script>
