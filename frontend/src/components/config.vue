<script setup>
import { reactive } from 'vue';
import { AddConfiguration, DowloadConfig } from '../../wailsjs/go/main/App';

const data = reactive({
  isOpen: false,
  resultDowloaText: '',
});

function upploadFileConf(file) {
  let reader = new FileReader();
  reader.onload = (event) => {
    const result = JSON.parse(event.target.result);
    const formatted = JSON.stringify(result, null, 2);
    AddConfiguration(formatted);
  };
  reader.readAsText(file.target.files[0]);
}

function toggleConfig() {
  data.isOpen = !data.isOpen;
}

function dowloadConfig() {
  DowloadConfig().then((result) => {
    data.resultDowloaText = result;
  });
}
</script>

<template>
  <p class="openConfig" @click="toggleConfig">Open config</p>
  <div class="config" v-if="data.isOpen">
    <p class="closing_icon" @click="toggleConfig">X</p>
    <p>Upload new config</p>
    <input
      type="file"
      id="configFile"
      @change="upploadFileConf"
      name="config file"
      accept=".json"
    />
    <button class="button" @click="dowloadConfig">
      DOWLOAD CURRENT CONFIG
    </button>
    <p>{{ data.resultDowloaText }}</p>
  </div>
</template>

<style scoped>
.config {
  background-color: rgb(33, 34, 34);
  width: 100%;
  height: 50%;
  position: absolute;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
#configFile {
  background-color: rgba(255, 255, 255, 1);
  color: black;
  border-radius: 5px;
  padding: 10px;
}
.closing_icon {
  position: absolute;
  right: 20px;
  top: 0;
  cursor: pointer;
}
.button {
  padding: 10px;
  border-radius: 5px;
  border: none;
  background-color: rgba(255, 255, 255, 1);
  margin-top: 20px;
  cursor: pointer;
}
.openConfig {
  position: absolute;
  bottom: 0;
  cursor: pointer;
  left: 20px;
  cursor: pointer;
}
</style>
