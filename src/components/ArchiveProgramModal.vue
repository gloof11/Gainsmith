<template>
  <div class="mr-4 mt-4 mb-4">
    <v-btn
      color="blue"
      variant="outlined"
      text="Archive"
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
      prepend-icon="mdi-archive-alert"
      title="Archive Program"
    >
      <v-form @submit.prevent="loading = true; submitArchive()">
        <v-container>
          <p class="mb-2">
            WARNING! This will remove the current program.
          </p>
          <v-text-field
            v-model="archiveFields.name"
            variant="outlined"
            label="Program Name"
            persistent-placeholder
            :rules="rules"
          />

          <v-date-input
            v-model="archiveFields.from"
            prepend-icon=""
            variant="outlined"
            label="From Date"
            persistent-placeholder
            :rules="rules"
          />

          <v-date-input
            v-model="archiveFields.to"
            prepend-icon=""
            variant="outlined"
            label="To Date"
            persistent-placeholder
            :rules="rules"
          />

          <v-textarea
            v-model="archiveFields.notes"
            variant="outlined"
            label="Notes"
            persistent-placeholder
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

  <!-- loading dialog -->
  <LoadingModal v-if="loading" title="Archiving..." icon="mdi-archive-alert" :status="archiveStatus" />
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue'
  import { useDisplay } from 'vuetify'
  import { pb } from '../lib/pocketbaseClient'
  import { rules } from '../lib/helpers'

  const { width } = useDisplay()

  const dialog = ref(false)
  const loading = ref(false)
  const archiveStatus = ref("")
  const archiveFields = ref({
    name: null,
    from: null,
    to: null,
    data: {},
    notes: null
  })

  const submitArchive = async () => {
    try {
      archiveStatus.value = "Gathering weight information"
      let weightInformation = await getWeightInformation()
      archiveStatus.value = "Gathering workout information"
      let workoutInformation = await getWorkoutInformation()
      archiveStatus.value = "Building JSON document"
      archiveFields.value.data = {
        weightInfo: weightInformation,
        workoutInfo: workoutInformation
      }

      await pb.collection('archived_programs').create(archiveFields.value)
    } catch (err) {
      console.log(err)
    } finally {
      loading.value = false
    }

    location.assign("/archived-programs")
  }

  async function getWeightInformation() {
    let weightInfo: [[string, number]?] = []

    const records = await pb.collection('weight_entries').getFullList({
      sort: 'date'
    })

    for (let i=0; i<records.length; i++) {
      weightInfo.push([records[i].date.toString().split(' ')[0], records[i].weight])
    }

    return weightInfo
  }

  async function getWorkoutInformation() {
    var workoutInformation: [{
      workout: string;
      type: string;
      metric: string;
      data: [{
        entryDate: string;
        value: number;
      }?]
    }?] = []

    // populate the workouts
    let records = await pb.collection('get_workout_entries_table').getFullList({
      sort: '-workout_type'
    })

    await Promise.all(records.map(async (value, index) =>  {
      archiveStatus.value = "Getting general workout information for " + value.workout
      let workout: string = value.workout
      let workout_type: string = value.workout_type
      let metric: string = value.metric
      let data: [{
        entryDate: string;
        value: number;
      }?] = []

      archiveStatus.value = "Getting workout data for " + value.workout
      // get workout data
      let dataRecords = await pb.collection('get_workout_entries_graph').getFullList({
        filter: `workout = '${workout}'`
      })

      dataRecords.map((item) => (
        data.push(
          {
            entryDate: String(item.created.toString().split(' ')[0]),
            value: Number(item.value)
          }
        )
      ))

      // add this to the spcific workout
      workoutInformation.push({
        workout: workout,
        type: workout_type,
        metric: metric,
        data: data
      })
    }))

    return workoutInformation
  }
</script>
