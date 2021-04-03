<template lang="pug">
div
  .title
    input.title__input(type="text" v-model="title" placeholder="input title")
  .deadline
    input.deadline__input(type="date" v-model="deadline")
  .submit
    button.submit__button(@click="registerTodo") submit
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
    async registerTodo(): Promise<void> {
      await axios.post("/api/todos", {
        title: this.title,
        deadline: this._deadline,
      });
      this.$emit("refleshTodos");
    },
  },
});
</script>
<style lang="scss">
.body {
  color: red;
}

.title {
  padding: 12px;
  height: 48px;

  &__input {
    width: 45vw;
    height: 100%;
  }
}

.deadline {
  padding: 12px;
  height: 48px;

  &__input {
    width: 45vw;
    height: 100%;
  }
}

.submit {
  padding: 12px;
  &__button {
    font-size: 1.2rem;
  }
}
</style>
