<template>
  <button :class="buttonClass" @click="eventOnclick">
    {{ text }}
  </button>
</template>

<script>
export default {
  name: "Button",
  props: {
    text: String,
  },
  computed: {
    buttonClass() {
      switch (this.text) {
        case "0":
          return "zero";

        case "+":
          return "plus";

        case "=":
          return "eval";
      }
      return null;
    },
    eventOnclick(){
        switch (this.text) {
        case "c":
          return this.clearExpression;

        case "=":
          return this.evaluateExpression;

      }
      return this.addToExpression;
    }
  },
  methods:{
      addToExpression(){
          this.$emit('add-expression',this.text)
      },
      evaluateExpression(){
          this.$emit('evaluate-expression')
      },
      clearExpression(){
          this.$emit('clear-expression')
      }
  }
};
</script>

<style>
</style>