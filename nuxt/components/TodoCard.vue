<template lang="pug">
div.card
  .card-title {{todo.title}}
  .card-deadline {{todo.deadline | dateFormat}}
  .card-state(:class="{'-done': todo.state}") {{todo.state | stateFormat}}
  .card-delete(@click.prevent="deleteTodo(todo.uuid)")
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
    stateFormat(state: boolean): string {
      return state ? "Done" : "NOT Done";
    },
  },
  methods: {
    deleteTodo(uuid: string): void {
      this.$emit("deleteTodo", uuid);
    },
  },
});
</script>
<style lang="scss">
.card {
  padding: 12px;
  width: 65vw;
  display: flex;
  flex-direction: row;
  border: 1px solid;
  border-style: dashed;
  align-items: center; /* 縦方向中央揃え */
  margin: 12px 0;

  &-title {
    width: 50%;
    word-wrap: break-word;
    word-break: break-all;
  }

  &-deadline {
    width: 30%;
    color: gray;
    &::before {
      font-size: 1.6rem;
      color: black;
      content: "| ";
    }
  }

  &-state {
    width: 20%;
    color: blue;
    &::before {
      font-size: 1.6rem;
      color: black;
      content: "| ";
    }
    &.-done {
      color: green;
    }
  }
  &-delete {
    &::after {
      position: relative;
      left: 36px;
      width: 20px;
      height: 20px;
      content: "✗";
      cursor: pointer;
    }
  }
}
</style>
