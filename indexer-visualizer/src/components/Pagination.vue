<script setup>
import { watch } from 'vue';
import { useMails } from '../composables/useMails';

const { from, nextPage, getMails, previousPage, mails, totalMails } = useMails();

watch(from, getMails);
</script>

<template>
  <div
    class="h-16 w-[481px] fixed bottom-0 bg-blue-600 text-white mx-[-20px] px-8 font-bold underline underline-offset-1 flex justify-between items-center">
    <span @click="() => previousPage()" class="cursor-pointer">{{
      from > 0 ? 'Previous' : ''
    }}</span>
    <span @click="() => nextPage()" class="cursor-pointer">Next</span>
  </div>
</template>
