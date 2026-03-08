<template>
  <div class="ma-4">
    <h1>{{ props.workoutType }}</h1>
    <div class="rounded-xl">
      <v-table v-if="loading == false" :hover="true">
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
              Workout Value
            </th>
            <th class="text-left">
              Progression
            </th>
          </tr>
        </thead>
        <tbody>
            <tr
              v-for="(item) in entries"
              :key="item.workout"
              @click="dialog = true; currentSelectedWorkout = item.workout"
            >
              <td>
                <v-row
                  class="ml-1"
                  align="center"
                >
                  {{ item.workout }}
                </v-row>
              </td>
              <td>
                <v-row
                  class="ml-1"
                  align="center"
                  justify="space-between"
                >
                {{ item.current_value }} {{ item.metric }}
                </v-row>
              </td>
              <td>
                <v-icon v-if="item.progress === 'higher_weight'" icon="mdi-arrow-up-bold" color="green" size="x-large" />
                <v-icon v-else-if="item.progress === 'lower_weight'" icon="mdi-arrow-down-bold" color="red" size="x-large" />
                <v-icon v-else icon="mdi-minus" color="grey" size="x-large" />
              </td>
            </tr>
        </tbody>
      </v-table>
      <LoadingModal v-else-if="loading == true" title="Loading..." />

      <!-- Workout Details Modal -->
      <v-dialog
        v-model="dialog"
        width="auto"
      >
        <v-card
          :width="`${width/1.2}`"
          max-width="400"
          prepend-icon="mdi-information-outline"
          :title="`${currentSelectedWorkout} Progression`"
        >
          <WorkoutGraph :workout="currentSelectedWorkout" />
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue'
  import { useDisplay } from 'vuetify'
  import { pb } from '../lib/pocketbaseClient'

  const { width } = useDisplay()
  const entries = ref()
  const dialog = ref(false)
  const currentSelectedWorkout = ref("")
  const loading = ref(false)
  
  const props = defineProps({
    workoutType: {
      type: String
    }
  })

  async function getWorkoutEntries() {
    try {
      loading.value = true
      const records = await pb.collection('get_workout_entries_table').getFullList({
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
    getWorkoutEntries()
  })
</script>
