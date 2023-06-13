<script setup lang="ts">
import VueJsonPretty from 'vue-json-pretty';
import {ref} from "vue";
import type {KnapSackSolution} from "@/type/KnapSackSolutionType";

let solution = ref(undefined)
let loading = ref(false)
let displaySolution = ref(false)

let maxValue = ref(0)
let durationTime = ref(0)

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

  solution.value = await response.json()

  console.log(props.solverName, " solution: ", solution.value)

  const knapsackSolution: KnapSackSolution = solution.value as unknown as KnapSackSolution
  maxValue.value = knapsackSolution.max_value
  durationTime.value = knapsackSolution.duration

  loading.value = false
  displaySolution.value = true
}

</script>

<script lang="ts">
export default {
  name: "Solver",
  // @ts-ignore
  props: ['generatedKnapSac', 'solverName', 'solverURL']
}
</script>

<template>
  <div v-if="props.generatedKnapSac !== undefined"  class="solution">
    <h2>{{ solverName }} Solver :</h2>
    <br>

    <p>{{ solverName }} solver URL :</p>
    <input v-model.number="solverURL"/>
    <br>
    <br>

    <button id="solve-btn" @click="solve">Solve Knap Sac</button>

    <hr>

    <div v-if="loading">
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

    <div v-if="displaySolution && !loading" id="solution">
      <p>Solution max value: <span class="bold">{{ maxValue }}</span></p>
      <p>Solution solving took: <span class="bold">{{ durationTime.toFixed(2) }}</span>s</p>
      <br>
      <vue-json-pretty :data=solution />
    </div>

  </div>
</template>

<style scoped>

#solve-btn {
  margin-bottom: 12%;
}

#solution {
  margin-top: 12%;
}

.bold {
  font-weight: bold !important;
  color: aqua;
}

.solution {
  display: inline-block;
  max-height: 55em;
  overflow-y: auto;
  padding-right: 25px;
  vertical-align: top;

  /* Left blue bar */
  border-left: 2px solid #2c3e50;
  margin-left: 50px;
  padding-left: 50px;
}

</style>
