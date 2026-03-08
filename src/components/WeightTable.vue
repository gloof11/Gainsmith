<template>
  <div class="ma-4">
    <div class="rounded-xl">
      <v-table v-if="loading == false">
        <thead>
          <tr>
            <th class="text-left">
              <v-row
                class="ml-1"
                align="center"
              >
                Date
                <v-btn
                  v-if="direction == 'ascending'"
                  density="compact"
                  class="ml-3"
                  icon="mdi-sort-ascending"
                  size="medium"
                  @click="changeDirection()"
                />
                <v-btn
                  v-else-if="direction == 'descending'"
                  density="compact"
                  class="ml-3"
                  icon="mdi-sort-descending"
                  size="medium"
                  @click="changeDirection()"
                />
              </v-row>
            </th>
            <th class="text-left">
              Body Weight (lb)
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(item, index) in entries"
            :key="item.date"
          >
            <td>{{ item.date.split(' ')[0] }}</td>
            <td>
              <v-row
                class="ml-1"
                align="center"
                justify="space-between"
              >
                {{ item.weight }}
                <v-btn
                  density="compact"
                  icon="mdi-minus"
                  variant="text"
                  size="large"
                  class="mr-16"
                  color="red"
                  @click="deleteEntry(item.date, index)"
                />
              </v-row>
            </td>
          </tr>
          <tr>
            <td>Add a weight!</td>
            <td>
              <v-row
                class="ml-1"
                align="center"
                justify="space-between"
              >
                <AddWeightModal />
              </v-row>
            </td>
          </tr>
        </tbody>
      </v-table>
      <LoadingModal v-else-if="loading == true" title="Loading..." />
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue'
  import { pb }  from '../lib/pocketbaseClient'

  const loading = ref(false)
  const entries = ref()
  const direction = ref('ascending')

  const deleteEntry = async (date: Date, index: number) => {
    entries.value.splice(index, 1)

    try {
      const record = await pb.collection('weight_entries').getFirstListItem(`date="${date}"`)
      const response = await pb.collection('weight_entries').delete(record.id)
      console.log(response)
    } catch (err) {
      console.log(err)
    }
  }

  async function getEntries() {
    try {
      loading.value = true
      const records = await pb.collection('weight_entries').getFullList({
        sort: 'date'
      })
      entries.value = records
    } catch (err) {
      console.log(err)
    } finally {
      loading.value = false
    }
  }

  onMounted(() => {
    getEntries()
  })

  const changeDirection = () => {
    if (direction.value == 'ascending') {
      entries.value.reverse()
      direction.value = 'descending'
    } else if (direction.value == 'descending') {
      entries.value.reverse()
      direction.value = 'ascending'
    }
  }
</script>
