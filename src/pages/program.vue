<template>
  <UpperNavBar value=4 />
  <div v-if="loading == false">
    <h1>
      Weight Overview
    </h1>

    <!-- weight table -->
    <v-table >
      <thead>
        <tr>
          <th class="text-left">
            <v-row
              class="ml-1"
              align="center"
            >
              Date
            </v-row>
          </th>
          <th class="text-left">
            Body Weight (lb)
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(item) in records.data.weightInfo"
          :key="item.date"
        >
          <td>{{ item[0] }}</td>
          <td>
            <v-row
              class="ml-1"
              align="center"
              justify="space-between"
            >
              {{ item[1] }}
            </v-row>
          </td>
        </tr>
      </tbody>
    </v-table>

    <br />
    <h1>
      Workout Overview
    </h1>
    <!-- workout table -->
    <v-table :hover="true">
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
        </tr>
      </thead>
      <tbody>
          <tr
            v-for="(item, index) in records.data.workoutInfo"
            :key="item.workout"
            @click="dialog = true; currentSelectedWorkout = item.workout; currentSelectedIndex = index"
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
              {{ item.data.pop().value }} {{ item.metric }}
              </v-row>
            </td>
          </tr>
      </tbody>
    </v-table>

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
        <WorkoutGraph :workout="currentSelectedWorkout" :data="records.data.workoutInfo[currentSelectedIndex].data" />
      </v-card>
    </v-dialog>
  </div>
  <LoadingModal v-else-if="loading == true" title="Loading..." />
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue'
  import { useRoute } from 'vue-router'
  import { useDisplay } from 'vuetify'
  import { pb } from '../lib/pocketbaseClient'

  const loading = ref(true)
  const dialog = ref(false)
  const records = ref()
  const currentSelectedWorkout = ref("")
  const currentSelectedIndex = ref(0)

  const { width } = useDisplay()

  async function getPrograms() {
    try {
      loading.value = true
      let id = useRoute().query.id

      if (id !== null)
        records.value = await pb.collection('archived_programs').getOne(id.toString(), {});
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
