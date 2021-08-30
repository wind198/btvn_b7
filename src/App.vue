<template>
  <div class="calculator">
    <div class="monitor">{{ expression }}</div>
    <div class="buttonList">
      <Button
        v-for="buttonText in buttonTextArray"
        :key="buttonText.index"
        :text="buttonText"
        @add-expression="handleAdd"
        @evaluate-expression="handleEval"
        @clear-expression="handleClear"
      />
    </div>
  </div>
</template>

<script>
import Button from "./components/Button.vue";
import axios from "axios";
export default {
  name: "App",
  components: {
    Button,
  },
  data() {
    return {
      expression: "",
    };
  },
  computed: {
    buttonTextArray() {
      return [
        "c",
        "/",
        "*",
        "-",
        "7",
        "8",
        "9",
        "+",
        "4",
        "5",
        "6",
        "1",
        "2",
        "3",
        "=",
        "0",
        ".",
      ];
    },
  },
  methods: {
    handleAdd(text) {
      this.expression += text;
    },
    async handleEval() {
      console.log("sending");
      try {
        const response = await axios.post("/evaluate", {
          expression: this.expression,
        });
        console.log(response.data);
        this.expression = response.data;
      } catch(error) {
        console.log(error);
        
      }
    },
    handleClear() {
      console.log("clearing");
      this.expression = "";
    },
  },
};
</script>

<style lang="scss">
* {
  box-sizing: border-box;
}
html {
  font-family: Arial, Helvetica, sans-serif;
  font-size: 20px;
}
.calculator {
  width: 350px;
  padding: 0.5rem;
  background: rgb(182, 182, 182);
  border-radius: 10px;
  border: 1px inset gray;
}
.monitor {
  width: 100%;
  height: 70px;
  line-height: 70px;
  text-align: right;
  background: white;
  border: 1px inset gray;
  padding-right: 0.5rem;
  font-size: 125%;
  border-radius: 5px;
  margin-bottom: 0.3rem;
}
.buttonList {
  height: 400px;
  display: grid;
  grid-template-columns: 25% 25% 25% 25%;
  button {
    margin: 0.1rem;
    background-color: white;
    border: 1px inset gray;
    border-radius: 5px;
    // height: 25px;
    font-size: 20px;
    text-transform: uppercase;
    font-weight: bold;
    cursor: pointer;
  }
  .plus,
  .eval {
    grid-row: span 2;
  }
  .zero {
    grid-column: span 2;
  }
}
</style>
