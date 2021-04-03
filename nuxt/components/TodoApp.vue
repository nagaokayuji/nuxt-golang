<template lang="pug">
div
  .todo-form
    TodoForm(@refleshTodos="getAllTodos")
  ul
    li(v-for="todo in todos" :key="todo.uuid")
      TodoCard(:todo="todo" @deleteTodo="deleteTodo")

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
    getAllTodos(): void {
      axios
        .get<TodoResponse[]>("/api/todos")
        .then((data) => (this.todos = data.data));
    },
    async deleteTodo(uuid: string): Promise<void> {
      await axios.delete<TodoResponse[]>("/api/todos/" + uuid);
      this.getAllTodos();
    },
  },
});
</script>
<style lang="scss">
.body {
  color: red;
}
ul {
  list-style: none;
  padding: 0;
}
</style>
