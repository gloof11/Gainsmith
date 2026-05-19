<template>
  <div class="ml-4 mt-4 mb-4">
    <v-btn
      color="green"
      text="Add Weight Entry"
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
      <v-form @submit.prevent="submitWorkoutEntry">
        <v-container>
          <v-select
            v-model="workoutEntryObject.workout"
            :items="workouts"
            :item-props="itemProps"
            label="Workout"
            :rules="rules"
            :disabled="props.workout != ''"
          />

          <v-text-field
            v-model="workoutEntryObject.value"
            variant="outlined"
            label="Value"
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
  import { ref, watch } from 'vue'
  import { useDisplay } from 'vuetify'
  import { rules } from '../lib/helpers'
  import pb from '../lib/pocketbaseClient'

  const { width } = useDisplay()

  const dialog = ref(false)
  const workouts = ref()

  const props = defineProps({
    workout: {
      type: String,
      default: ""
    },
  });

  const workoutEntryObject = ref({
    workout: {workout: props.workout},
    value: null
  })

  const submitWorkoutEntry = async () => {
    try {
      const id = workouts.value.find((workout: any) => workout.workout === workoutEntryObject.value.workout.workout).id

      await pb.collection('workout_entries').create({
        "workout": id,
        "value": workoutEntryObject.value.value,
      })
    } catch (err) {
      console.log(err)
    }

    location.reload()
  }

  const loadWorkouts = async () => {
    try {
      const records = await pb.collection('workouts').getFullList({
        sort: '-workout_type'
      })
      workouts.value = records
    } catch (err) {
      console.log(err)
    }
  }

  function itemProps(item: any) {
    return {
      title: item.workout,
      subtitle: item.workout_type
    }
  }

  watch(dialog, () => {
    if (dialog.value == true)
      loadWorkouts()
  })

</script>
