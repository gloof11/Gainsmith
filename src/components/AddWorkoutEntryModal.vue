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
  import { pb } from '../lib/pocketbaseClient'
  import { rules } from '../lib/helpers'

  const { width } = useDisplay()

  const dialog = ref(false)
  const workouts = ref()

  const workoutEntryObject = ref({
    workout: {workout: ""},
    value: null
  })

  const submitWorkoutEntry = async () => {
    try {
      let id = workouts.value.find((workout: any) => workout.workout === workoutEntryObject.value.workout.workout).id
      await pb.collection('workout_entries').create({
        workout: id,
        value: workoutEntryObject.value.value,
        remarks: ""
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
      console.log(typeof workouts.value[0])
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
