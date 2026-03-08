<template>
  <div class="ma-4">
    <h1>{{ props.workoutType }}</h1>
    <div class="rounded-xl">
      <v-table v-if="loading == false">
        <thead>
          <tr>
            <th class="text-left">
              <v-row
                class="ml-1"
                align="center"
              >
                Workout
              </v-row>
            </th>
            <th class="text-left">
              Workout Type
            </th>
            <th class="text-left">
              Metrics
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(item, index) in entries"
            :key="item.workout"
          >
            <td>{{ item.workout }}</td>
            <td>{{ item.workout_type }}</td>
            <td>
              <v-row
                class="ml-1"
                align="center"
                justify="space-between"
              >
                {{ item.metric }}
                <v-btn
                  density="compact"
                  icon="mdi-minus"
                  variant="text"
                  size="large"
                  class="mr-16"
                  color="red"
                  @click="deleteEntry(item.workout, index)"
                />
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
  import { pb } from '../lib/pocketbaseClient'

  const loading = ref(false)
  const entries = ref()
  const direction = ref('ascending')
  
  const props = defineProps({
    workoutType: {
      type: String
    }
  })

  const deleteEntry = async (workout: string, index: number) => {
    entries.value.splice(index, 1)

    const record = await pb.collection('workouts').getFirstListItem(`workout="${workout}"`)
    const response = await pb.collection('workouts').delete(record.id)
    console.log(response)
  }

  async function getWorkouts() {
    try {
      loading.value = true
      const records = await pb.collection('workouts').getFullList({
        filter: `workout_type = '${props.workoutType}'`
      })
      entries.value = records
    } catch (error) {
      console.log(error)
    } finally {
      loading.value = false
    }
  }

  onMounted(() => {
    getWorkouts()
  })
</script>
