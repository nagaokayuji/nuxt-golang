<template lang="pug">
  .container
    Logo
    div
      button(@click="push") push here to send request
    div ↓↓ response ↓↓
    span {{sampleData}}
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

export default Vue.extend({
  data() {
    return {
      sampleData: "",
    };
  },
  mounted() {
    console.log("mounted");
  },
  methods: {
    push(): void {
      console.log("pushed");
      axios.get("/api/").then((response) => (this.sampleData = response.data));
    },
  },
});
</script>

<style>
.container {
  min-height: 100vh;
}

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
