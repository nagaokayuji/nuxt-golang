<template lang="pug">
div
  TodoForm(@refleshTodos="getAllTodos")
  button(@click="getAllTodos") getAll
  ul
    li(v-for="todo in todos" :key="todo.uuid")
      TodoCard(:todo="todo")

</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";

// form
interface TodoRequest {
  title: string;
  deadline: Date;
}

export interface TodoResponse {
  uuid: string;
  title: string;
  deadline: Date;
  status: boolean;
}

interface IData {
  todos: TodoResponse[];
}

export default Vue.extend({
  data(): IData {
    return {
      todos: [],
    };
  },
  mounted() {
    // initial view
    this.getAllTodos();
  },
  methods: {
    getAllTodos() {
      axios
        .get<TodoResponse[]>("/api/todos")
        .then((data) => (this.todos = data.data));
    },
  },
});
</script>
<style>
.body {
  color: red;
}
ul {
  list-style: none;
}
</style>
