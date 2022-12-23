<script setup>
import { watch } from 'vue';
import { useMails } from '../composables/useMails';

const { searchTerm, getMails, setSearchTerm } = useMails();
watch(searchTerm, getMails);
</script>

<template>
  <div class="w-full">
    <div class="mb-3 w-full mt-2">
      <input
        type="search"
        :value="searchTerm"
        @input="(event) => setSearchTerm(event.target.value)"
        class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
        placeholder="Search" />
    </div>
  </div>
</template>
