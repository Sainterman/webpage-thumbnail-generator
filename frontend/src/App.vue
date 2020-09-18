<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Generate a thumbnail of a website</h1>

        <form v-on:submit.prevent="makeWebsiteThumbnail">
          <div class="form-group">
            <input v-model="websiteUrl" type="text" id="website-input" placeholder="Enter a website" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Generate!</button>
          </div>
        </form>

        <img :src="thumbnailUrl"/>
      </div>
    </div>
  </div>
</template>
<script>
  import axios from 'axios';
  export default {
    name: 'App',

    data() { return {
      websiteUrl: '',
      thumbnailUrl: '',
    }},

    methods: {
      makeWebsiteThumbnail() {
        console.log(`I should create a website thumbnail of ${this.websiteUrl}`);
        axios.post("https://screenshotapi.net/api/v1/screenshot", {
          token: "NIBMXXMR3522DJ8ZBZ4FXUMLMSEWJNJX",
          url: this.websiteUrl,
          width: 1920,
          height: 1080,
          output: 'json',
          thumbnail_width: 720
        })
        .then((response) => {
          this.thumbnailUrl= response.data.screenshot;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        })
      }
    }
  }
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
