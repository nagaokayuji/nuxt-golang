<template lang="pug">
div
  input(type="text" v-model="title" placeholder="input title")
  input(type="date" v-model="deadline")
  button(@click="registerTodo") submit
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import moment from "moment";

interface IData {
  title: string;
  deadline: string;
}

export default Vue.extend({
  data(): IData {
    return {
      title: "",
      deadline: moment().format("YYYY-MM-DD"),
    };
  },
  computed: {
    _deadline(): string {
      return moment(this.deadline, "YYYY-MM-DD").toDate().toISOString();
    },
  },
  methods: {
    registerTodo() {
      axios
        .post("/api/todos", {
          title: this.title,
          deadline: this._deadline,
        })
        .then(this.$emit("refleshTodos"));
    },
  },
});
</script>
<style>
.body {
  color: red;
}
</style>
