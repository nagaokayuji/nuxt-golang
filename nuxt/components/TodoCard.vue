<template lang="pug">
div
  p {{todo.title}} {{todo.deadline | dateFormat}} {{todo.state}}

</template>

<script lang="ts">
import Vue, { PropType } from "vue";
import axios from "axios";
import moment from "moment";
interface TodoResponse {
  uuid: string;
  title: string;
  deadline: Date;
  status: boolean;
}

export default Vue.extend({
  props: {
    todo: Object as PropType<TodoResponse>,
  },
  filters: {
    dateFormat(date: Date): string {
      return moment(date).format("YYYY/MM/DD");
    },
  },
  methods: {},
});
</script>
<style>
.body {
  color: red;
}
</style>
