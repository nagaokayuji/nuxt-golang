<template lang="pug">
div
  TodoForm
  button(@click="getAllTodos") getAll
  div(v-for="todo in todos" :key="todo.uuid")
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
    console.log("mounted");
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
</style>
