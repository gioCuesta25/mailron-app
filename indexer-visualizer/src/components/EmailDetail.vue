<script setup>
import { useMails } from '../composables/useMails';

const { selectedMail, showMailDetail } = useMails();
</script>

<template>
  <div
    v-if="!selectedMail"
    class="w-4/5 px-12 pb-24 pt-10 font-light overflow-y-auto flex items-center justify-center">
    <span class="font-light text-3xl text-neutral-400"> Select an email from the left </span>
  </div>
  <div v-if="selectedMail" class="w-4/5 px-12 pb-24 pt-10 font-light overflow-y-auto">
    <div class="w-full flex justify-end">
      <span
        @click="() => showMailDetail(undefined)"
        class="text-2xl font-bold cursor-pointer text-red-700 underline"
        >close</span
      >
    </div>
    <div>
      <p class="text-2xl border-b border-gray-300 pb-3">
        <span class="font-bold">From: </span>
        <span>{{ selectedMail['from'] }}</span>
      </p>
      <p class="text-2xl py-3 border-b border-gray-300">
        <span class="font-bold">To: </span>
        <span>{{ selectedMail['to'] }}</span>
      </p>
      <p class="text-2xl py-3 border-b border-gray-300">
        <span class="font-bold">Subject: </span>
        <span>{{ selectedMail['subject'] }}</span>
      </p>
      <p class="text-2xl py-3">
        {{ selectedMail['content'] }}
      </p>
    </div>
  </div>
</template>
