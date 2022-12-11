<script setup lang="ts">
import VueJsonPretty from 'vue-json-pretty';
import 'vue-json-pretty/lib/styles.css';
import {ref} from "vue";

let bagSize: number = 100
let nbOfItems: number = 5

let generatedKnapSac = ref(undefined)
let displayGeneratedComp = ref(false)

let solution = ref(undefined)
let displaySolution = ref(false)

let generatorURL: string = "localhost/generator"
let naiveSolverURL: string = "localhost/naive-solver"


function getGeneratorURL() {
  return "http://" + generatorURL + "/generate?bagSize=" + bagSize + "&numberOfItems=" + nbOfItems
}

function getNaiveSolverURL() {
  return "http://" + naiveSolverURL + "/solve"
}

async function generate() {
  const response = await fetch(getGeneratorURL(), {
    method: 'POST',
    mode: "cors",
    headers: {'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'}
  });

  if (!response.ok) {
    alert("error while generating")
  }

  generatedKnapSac.value = await response.json()
  displayGeneratedComp.value = true
  displaySolution.value = false

}

async function solve() {
  const response = await fetch(getNaiveSolverURL(), {
    method: 'POST',
    mode: "cors",
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify(generatedKnapSac.value)
  });

  if (!response.ok) {
    alert("error while generating")
  }

  solution.value = await response.json()
  displaySolution.value = true
}
</script>

<template>
  <div class="app">
    <div class="form">
      <div>
        <p>Generator URL :</p>
        <input v-model.number="generatorURL"/>
      </div>
      <br>
      <div>
        <p>Naive solver URL :</p>
        <input v-model.number="naiveSolverURL"/>
      </div>

      <br><hr><br>

      <div>
        <p>Bag Size :</p>
        <input v-model.number="bagSize" placeholder="eg: 100"/>
      </div>
      <br>
      <div>
        <p>Number of items :</p>
        <input v-model.number="nbOfItems" placeholder="eg: 12"/>
      </div>

      <br><br>
      <button @click="generate">Generate Knap Sac</button>

      <br><br>
      <button @click="solve">Solve Knap Sac</button>
    </div>

    <div class="sac" :class="{ 'display-none': !displayGeneratedComp }">
      <h2>Generated Sac :</h2>
      <br>
      <vue-json-pretty :data=generatedKnapSac />
    </div>

    <div class="solution" :class="{ 'display-none': !displaySolution }">
      <h2>Solution :</h2>
      <br>
      <vue-json-pretty :data=solution />
    </div>
  </div>

</template>

<style scoped>
.app {
  max-width: 70%;
  width: fit-content;
  margin: 10% auto 0 auto;
  padding: 2rem;

  font-weight: normal;
}

.form {
  display: inline-block;
  margin-right: 50px;
  padding-right: 50px;
  border-right: 2px solid #2c3e50;
}

.sac {
  display: inline-block;
  max-height: 600px;
  overflow-y: auto;
  margin-right: 100px;
  padding-right: 25px;
  vertical-align: top;
}

.solution {
  display: inline-block;
  max-height: 600px;
  overflow-y: auto;
  margin-right: 50px;
  padding-right: 25px;
  vertical-align: top;
}

.display-none {
  display: none;
}

::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
  color: darkgrey;
  opacity: 1; /* Firefox */
}
</style>
