<template>
  <UpperNavBar value="4" />
  <div
    v-if="loading == false"
    class="ma-4 rounded-xl"
  >
    <v-expansion-panels>
      <v-expansion-panel
        v-for="(item) in records"
        :key="item.id"
      >
        <v-expansion-panel-title>
          {{ item.name }}
        </v-expansion-panel-title>
        <v-expansion-panel-text>
          <label for="from">From:</label>
          {{ item.from.toString().split(' ')[0] }}

          <br>
          <br>

          <label for="to">To:</label>
          {{ item.to.toString().split(' ')[0] }}

          <br>
          <br>

          <label for="notes">Notes:</label>
          <p>
            {{ item.notes }}
          </p>

          <br>

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
  </div>
  <LoadingModal
    v-else-if="loading == true"
    title="Loading..."
  />
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue'
  import pb from '../lib/pocketbaseClient'
  import type { RecordModel } from 'pocketbase'

  const loading = ref(false)

  let records: RecordModel[]

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
