<script setup>
import { reactive } from 'vue';
import { DivideTabs } from '../../wailsjs/go/main/App';

const data = reactive({
  translations: '',
  folderPath: '',
  translationKey: '',

  resultText: 'Enter your translations here 👇',

  languages: [
    'en_US',
    'en_GB',
    'en_AU',
    'en_NZ',
    'en_CA',
    'de_DE',
    'es_MX',
    'es_ES',
    'fr_CA',
    'fr_FR',
    'nl_NL',
    'pt_BR',
  ],
});

function greet() {
  DivideTabs(data.translations, data.folderPath, data.translationKey).then(
    (result) => {
      data.resultText = result;
    }
  );
}
</script>
<!-- 'ciao ciao  components  stai  bene'-->
<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input
        id="translations"
        v-model="data.translations"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Add Translation"
      />
      <br />
      <br />
      <input
        id="translationKey"
        v-model="data.translationKey"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Add Translation Key"
      />
      <br />
      <br />
      <input
        id="folderPath"
        v-model="data.folderPath"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Add Folder Path"
      />
      <br />
      <br />
      <button class="btn" @click="greet">Add Translatins</button>
    </div>
  </main>
</template>

<style scoped>
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
