<template lang="pug">
div
  div.card
    .card-title {{todo.title}}
    .card-deadline {{todo.deadline | dateFormat}}
    .card-state(:class="{'-done': todo.state}" @click.prevent="toggleTodo(todo.uuid)")
    .card-delete
      .card-delete__button(@click.prevent="deleteTodo(todo.uuid)")
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
    toggleTodo(uuid: string): void {
      this.$emit("toggleTodo", uuid);
    },
  },
});
</script>
<style lang="scss">
.card {
  width: 65vw;
  display: flex;
  flex-direction: row;
  border: 1px solid;
  border-style: dashed;
  align-items: center; /* 縦方向中央揃え */
  margin: 12px 0;
  font-size: 1rem;

  &-title {
    width: calc(100% - 6rem - 20px);
    word-wrap: break-word;
    word-break: break-all;
    padding: 12px;
  }

  &-deadline {
    padding: 12px;
    width: 6rem;
    color: gray;
  }

  &-state {
    height: 18px;
    width: 18px;
    font-size: 18px;
    cursor: pointer;
    margin: 12px;
    position: relative;
    &::before {
      position: absolute;
      height: 18px;
      width: 18px;
      border: 1px solid;
      display: inline-block;
      content: "";
    }
    &.-done {
      &::after {
        position: absolute;
        content: "";
        height: 5px;
        width: 12px;
        display: inline-block;
        border-left: 1px solid;
        border-bottom: 1px solid;
        transform: rotate(-45deg);
        left: 4px;
        top: 4px;
      }
    }
  }

  &-delete {
    position: relative;
    right: -20px;
    &__button {
      cursor: pointer;
      display: inline-block;
      vertical-align: middle;
      color: rgb(73, 73, 73);
      width: 0.9rem;
      height: 0.05rem;
      background: currentColor;
      position: relative;
      transform: rotate(45deg);
      &::before {
        content: "";
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: inherit;
        transform: rotate(90deg);
      }
    }
  }
}
</style>
