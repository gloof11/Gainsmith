<template>
  <div class="ml-4 mt-4 mb-4">
    <v-btn
      color="green"
      text="Add A Workout"
      @click="dialog = true"
    />
  </div>

  <v-dialog
    v-model="dialog"
    width="auto"
  >
    <v-card
      :width="`${width/1.2}`"
      max-width="400"
      prepend-icon="mdi-plus"
      title="Add a Workout"
    >
      <v-form @submit.prevent="submitWorkout">
        <v-container>
          <v-text-field
            v-model="workoutObject.workout"
            variant="outlined"
            label="Workout"
            :rules="rules"
          />

          <v-select
            v-model="workoutObject.workout_type"
            :items="['Push', 'Pull', 'Legs', 'Other']"
            label="Workout Type"
            :rules="rules"
          />

          <v-text-field
            v-model="workoutObject.metric"
            variant="outlined"
            label="Metric"
            :rules="rules"
          />
        </v-container>

        <div class="d-flex justify-end mr-8 mb-4">
          <v-btn
            color="green"
            variant="outlined"
            type="submit"
            text="Ok"
          />
        </div>
      </v-form>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue'
  import { useDisplay } from 'vuetify'
  import { pb } from '../lib/pocketbaseClient'
  import { rules } from '../lib/helpers'

  const { width } = useDisplay()

  const dialog = ref(false)
  const workoutObject = ref({
    workout: null,
    workout_type: null,
    metric: null
  })

  const submitWorkout = async () => {
    try {
      const records = await pb.collection('workouts').create({
        "workout": workoutObject.value.workout,
        "workout_type": workoutObject.value.workout_type,
        "metric": workoutObject.value.metric
      })
      console.log(records)
    } catch (err) {
      console.log(err)
    }

    location.reload()
  }
</script>
