<script setup>
import { ref, onMounted } from 'vue';
defineProps(['mails']);

const tableHeaders = ['From', 'To', 'Subject'];

const mails = ref({});

onMounted(async () => {
  const res = await fetch('http://localhost:4000/documents/match?term=allen&from=0');
  const data = await res.json();
  mails.value = data['hits']['hits'];
  console.log(mails.value);
});
</script>

<template>
  <div class="flex flex-col">
    <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
      <div class="py-2 inline-block min-w-full sm:px-6 lg:px-8">
        <div class="overflow-hidden">
          <table class="min-w-full">
            <thead class="bg-white border-b">
              <tr>
                <th
                  scope="col"
                  class="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                  v-for="(header, index) in tableHeaders"
                  :key="index">
                  {{ header }}
                </th>
              </tr>
            </thead>
            <tbody>
              <tr class="border-b" v-for="mail in mails">
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ mail['_source']['from'] }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ mail['_source']['to'] }}
                </td>
                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                  {{ mail['_source']['subject'] }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
