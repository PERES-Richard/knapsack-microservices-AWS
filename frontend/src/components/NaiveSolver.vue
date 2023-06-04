<script setup lang="ts">
import VueJsonPretty from 'vue-json-pretty';
import {ref} from "vue";

let solution = ref(undefined)
let loading = ref(false)
let displaySolution = ref(false)

const props = defineProps({
  generatedKnapSac: Object,
  solverName: String
})

let solverURL: string = "localhost/" + props.solverName?.toLocaleLowerCase() + "-solver/solve"

async function solve() {
  loading.value = true

  const response = await fetch("http://" + solverURL, {
    method: 'POST',
    mode: "cors",
    keepalive: true,
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify(props.generatedKnapSac)
  });

  if (!response.ok) {
    alert("error while solving")
  }

  console.log(response)
  solution.value = await response.json()
  loading.value = false
  displaySolution.value = true
}

</script>

<script lang="ts">
export default {
  name: "Solver",
  props: ['generatedKnapSac', 'solverName', 'solverURL']
}
</script>

<template>
  <div class="solution">
    <h2>{{ solverName }} Solver :</h2>
    <br>

    <p>{{ solverName }} solver URL :</p>
    <input v-model.number="solverURL"/>
    <br>
    <br>

    <button @click="solve">Solve Knap Sac</button>

    <br>
    <br>
    <hr>
    <br>


    <div :class="{ 'display-none': !loading }">
      <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
           style="margin: auto; display: block; shape-rendering: auto;" width="200px" height="200px"
           viewBox="0 0 100 100" preserveAspectRatio="xMidYMid">
        <circle cx="50" cy="50" r="32" stroke-width="8" stroke="#93dbe9"
                stroke-dasharray="50.26548245743669 50.26548245743669" fill="none" stroke-linecap="round">
          <animateTransform attributeName="transform" type="rotate" dur="1s" repeatCount="indefinite" keyTimes="0;1"
                            values="0 50 50;360 50 50"></animateTransform>
        </circle>
        <circle cx="50" cy="50" r="23" stroke-width="8" stroke="#689cc5"
                stroke-dasharray="36.12831551628262 36.12831551628262" stroke-dashoffset="36.12831551628262" fill="none"
                stroke-linecap="round">
          <animateTransform attributeName="transform" type="rotate" dur="1s" repeatCount="indefinite" keyTimes="0;1"
                            values="0 50 50;-360 50 50"></animateTransform>
        </circle>
      </svg>
    </div>

    <div :class="{ 'display-none': !displaySolution && loading }">
      <br>
      <vue-json-pretty :data=solution />
    </div>

  </div>
</template>

<style scoped>

.display-none {
  display: none;
}

.solution {
  display: inline-block;
  max-height: 55em;
  overflow-y: auto;
  margin-right: 50px;
  padding-right: 25px;
  vertical-align: top;

  /* Left blue bar */
  border-left: 2px solid #2c3e50;
  margin-left: 50px;
  padding-left: 50px;
}

</style>
