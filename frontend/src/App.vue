<script setup lang="ts">
import 'vue-json-pretty/lib/styles.css';
import {ref} from "vue";
import VueJsonPretty from 'vue-json-pretty';
import Solver from "@/components/SolverComp.vue";

let bagSize: number = 100
let nbOfItems: number = 5

let generatedKnapSac = ref(undefined)
let displayGeneratedComp = ref(false)

let generatorURL: string = "localhost/generator"

function getGeneratorURL() {
  return "http://" + generatorURL + "/generate?bagSize=" + bagSize + "&numberOfItems=" + nbOfItems
}

async function generate() {
  const response = await fetch(getGeneratorURL(), {
    method: 'POST',
    mode: "cors",
    keepalive: true,
    headers: {'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'}
  });

  if (!response.ok) {
    alert("error while generating")
  }

  generatedKnapSac.value = await response.json()
  console.log("Generated knapsack: ", generatedKnapSac.value)

  displayGeneratedComp.value = true
}

const solverList = ["Naive"]

</script>

<template>
  <div class="app">
    <div class="form" :class="{ 'blue-separator': displayGeneratedComp }">
      <div>
        <p>Generator URL :</p>
        <input v-model.number="generatorURL"/>
      </div>

      <br>
      <hr>
      <br>

      <div>
        <p>Bag Size :</p>
        <input v-model.number="bagSize" placeholder="eg: 100"/>
      </div>
      <br>
      <div>
        <p>Number of items :</p>
        <input v-model.number="nbOfItems" placeholder="eg: 12"/>
      </div>

      <button @click="generate" id="generate-btn">Generate Knap Sac</button>
    </div>

    <div v-if="displayGeneratedComp" class="sac">
      <h2>Generated Sac :</h2>
      <br>
      <vue-json-pretty :data=generatedKnapSac />
    </div>

    <Solver :generatedKnapSac=generatedKnapSac v-for="solverName in solverList" :key="solverName" :solver-name=solverName />

  </div>

</template>

<style scoped>
.app {
  max-width: 100%;
  width: fit-content;
  margin: 2% auto 0 auto;
  padding: 2rem;

  font-weight: normal;
}

.form {
  display: inline-block;
}

.blue-separator {
  /* Right blue bar */
  margin-right: 50px;
  padding-right: 50px;
  border-right: 2px solid #2c3e50;
}

#generate-btn {
  margin-top: 15%;
}

.sac {
  display: inline-block;
  max-height: 55em;
  overflow-y: auto;
  padding-right: 25px;
  vertical-align: top;
}

::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
  color: darkgrey;
  opacity: 1; /* Firefox */
}
</style>
