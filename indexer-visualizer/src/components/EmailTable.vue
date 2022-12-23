<script setup>
import SearchBar from './SearchBar.vue';
import Pagination from './Pagination.vue';
import { useMails } from '../composables/useMails';
import { onMounted } from 'vue';

const { mails, getMails, showMailDetail } = useMails();
onMounted(getMails);
</script>

<template>
  <div
    class="h-100 bg-gray-200 w-[600px] p-5 overflow-y-auto pb-[75px] font-light overflow-x-hidden relative">
    <SearchBar />
    <div
      class="w-full bg-white mb-5 cursor-pointer rounded-md h-auto p-3 flex gap-3"
      v-for="mail in mails"
      @click="() => showMailDetail(mail['_source'])"
      :key="mail['_id']">
      <div
        class="h-9 w-9 rounded-full bg-blue-600 flex items-center justify-center text-white font-bold">
        @
      </div>
      <div class="w-full p-1">
        <p>
          <span class="font-bold">From: </span> <span>{{ mail['_source']['from'] }}</span>
        </p>
        <p>
          <span class="font-bold">To: </span> <span>{{ mail['_source']['to'] }}</span>
        </p>
        <p>
          <span class="font-bold">Subject: </span> <span>{{ mail['_source']['subject'] }}</span>
        </p>
      </div>
    </div>
    <Pagination />
  </div>
</template>
