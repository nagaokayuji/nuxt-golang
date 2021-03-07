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

<style lang="scss">
.container {
  min-height: 100vh;
}
</style>
