<script setup>
import { reactive } from 'vue';
import { AddTranslationHandler } from '../../wailsjs/go/main/App';
import Config from './config.vue';

const data = reactive({
  translations: '',
  folderPath: '',
  translationKey: '',

  resultText: 'Enter your translations here 👇',
});


function submitTranslations() {
  AddTranslationHandler(data.translations, data.folderPath, data.translationKey).then(
    (result) => {
      data.resultText = result;
    }
  );
}
</script>

<template>
  <main class="main">
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="inputs" class="input-box">
      <input
        id="translations"
        v-model="data.translations"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Add Translation"
      />
      <input
        id="translationKey"
        v-model="data.translationKey"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Add Translation Key"
      />
      <input
        id="folderPath"
        v-model="data.folderPath"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Add Folder Path"
      />
      <button class="btn" @click="submitTranslations">Add Translatins</button>
    </div>
    <Config></Config>
  </main>
</template>

<style scoped>
#inputs {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  gap: 10px;
  margin: 0;
  padding: 0;
}

#configFile {
  position: absolute;
  bottom: 30px;
  width: 100%;
  background-color: #cfd9df;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  width: 70%;
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
