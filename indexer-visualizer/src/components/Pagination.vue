<script setup>
import { watch } from 'vue';
import { useMails } from '../composables/useMails';

const { from, nextPage, getMails, previousPage } = useMails();

watch(from, getMails);
</script>

<template>
  <div
    class="h-16 sticky bottom-0 bg-blue-600 text-white mx-[-20px] px-8 font-bold underline underline-offset-1 flex justify-between items-center">
    <span @click="() => previousPage()" class="cursor-pointer">{{
      from > 0 ? 'Anterior' : ''
    }}</span>
    <span @click="() => nextPage()" class="cursor-pointer">Siguiente</span>
  </div>
</template>
