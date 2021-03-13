<template lang="pug">
  .container
    div
      button(@click="push") push here to send request
    div ↓↓ response ↓↓
    p {{sampleData}}
    TodoApp
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
      axios.get("/api/").then((response) => (this.sampleData = response.data));
    },
  },
});
</script>

<style lang="scss">
.container {
  text-align: center;
}
</style>
