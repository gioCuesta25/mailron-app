<script setup>
import Header from './components/Header.vue';
import EmailTable from './components/EmailTable.vue';
import EmailDetail from './components/EmailDetail.vue';
import SearchBar from './components/SearchBar.vue';

import { ref, onMounted, watch } from 'vue';

const mails = ref({});
const selectedMail = ref({});
const searchTerm = ref('');

onMounted(async () => {
  const res = await fetch(`http://localhost:4000/documents/match?term=${searchTerm.value}&from=0`);
  const data = await res.json();
  mails.value = data['hits']['hits'];
});

const selectMail = (mail) => {
  selectedMail.value = mail;
};

const search = (text) => {
  searchTerm.value = text;
};

watch(searchTerm, async (newValue, oldValue) => {
  console.log(`searching: ${newValue}`);
  try {
    const res = await fetch(`http://localhost:4000/documents/match?term=${newValue}&from=0`);
    const data = await res.json();
    mails.value = data['hits']['hits'];
  } catch (error) {
    console.log('Error: ', error);
  }
});
</script>

<template>
  <div class="flex flex-col w-screen h-screen overflow-hidden">
    <div class="flex gap-1 h-full">
      <EmailTable :mails="mails" @select-mail="selectMail" />
      <EmailDetail :mail="selectedMail" />
    </div>
  </div>
</template>
