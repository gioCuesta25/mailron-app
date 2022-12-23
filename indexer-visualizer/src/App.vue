<script setup>
import Header from './components/Header.vue';
import EmailTable from './components/EmailTable.vue';
import EmailDetail from './components/EmailDetail.vue';
</script>

<template>
  <div class="flex flex-col w-screen h-screen overflow-hidden">
    <Header />
    <div class="flex gap-3 h-full">
      <EmailTable />
      <EmailDetail />
    </div>
  </div>
</template>
