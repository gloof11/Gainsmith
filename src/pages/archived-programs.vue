<template>
  <UpperNavBar value=4 />
  <v-expansion-panels v-if="loading == false">
    <v-expansion-panel v-for="(item) in records" :key="item.id">

      <v-expansion-panel-title>
        {{ item.name }}
      </v-expansion-panel-title>
      <v-expansion-panel-text>
        <label for="from">From:</label>
        {{  item.from.toString().split(' ')[0]  }}

        <br />
        <br />

        <label for="to">To:</label>
        {{  item.to.toString().split(' ')[0]  }}

        <br />
        <br />

        <label for="notes">Notes:</label>
        <p>
          {{ item.notes }}
        </p>

        <br />

        <router-link :to="`/program/?id=${item.id}`">
          <v-btn
            color="blue"
            variant="outlined"
            text="View"
          />
        </router-link>

      </v-expansion-panel-text>

    </v-expansion-panel>
  </v-expansion-panels>
  <LoadingModal v-else-if="loading == true" title="Loading..." />
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue'
  import { pb } from '../lib/pocketbaseClient'

  const loading = ref(false)

  let records: any = []

  async function getPrograms() {
    try {
      loading.value = true
      records = await pb.collection('archived_programs').getFullList({
        sort: 'created'
      })
    } catch (error) {
      console.log(error)
    } finally {
      loading.value = false
    }
  }

  onMounted(() => {
    getPrograms()
  })
</script>
