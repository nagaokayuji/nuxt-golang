<template lang="pug">
div
  .form-title
    input.form-title__input(type="text" v-model="title" placeholder="input title")
  .form-deadline
    input.form-deadline__input(type="date" v-model="deadline")
  .form-submit
    button.form-submit__button(@click="registerTodo") submit
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

.form {
  &-title {
    padding: 12px;
    height: 48px;

    &__input {
      font-size: 1.2rem;
      width: 45vw;
      height: 100%;
    }
  }

  &-deadline {
    padding: 12px;
    height: 48px;

    &__input {
      font-size: 1.2rem;
      width: 45vw;
      height: 100%;
    }
  }

  &-submit {
    padding: 12px;
    &__button {
      height: 48px;
      padding: 4px 24px;
      cursor: pointer;
      color: white;
      background-color: blue;
      border-radius: 8px;
      font-size: 1.2rem;
      border: none;
      &:hover {
        background-color: rgb(51, 51, 255);
      }
    }
  }
}
</style>
