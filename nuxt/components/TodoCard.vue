<template lang="pug">
div.container
  .card-title {{todo.title}}
  .card-deadline {{todo.deadline | dateFormat}}
  .card-state {{todo.state}}
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
<style lang="scss">
.container {
  padding: 12px;
  width: 65vw;
  display: flex;
  flex-direction: row;
}

.card {
  &-title {
    width: 50%;
  }
  &-deadline {
    width: 30%;
    color: gray;
  }
  &-state {
    width: 20%;
    color: blue;
    .-done {
      color: green;
    }
  }
}
</style>
