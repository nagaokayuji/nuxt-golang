<template lang="pug">
div
  p {{todo.uuid}} {{todo.title}}

</template>

<script lang="ts">
import Vue, { PropType } from "vue";
import axios from "axios";

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
  methods: {},
});
</script>
<style>
.body {
  color: red;
}
</style>
