<template>
  <div class="d-flex align-center">
    <h1 v-if="props.workoutType != ''">
      {{ props.workoutType }} Day Progression
    </h1>
  </div>

  <div v-if="graphData !== undefined">
    <div class="bg-grey-lighten-2 rounded">
      <Line :data="graphData" :options="options" />
    </div>
    <div v-if="inTable" class="d-flex justify-center">
      <AddWorkoutEntryModal :workout="workout" />
    </div>
  </div>
  <div v-else>
    <v-progress-circular indeterminate :size="89" :width="8" class="mb-8" />
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import pb from '../lib/pocketbaseClient'
import type { RecordModel } from 'pocketbase'

import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Colors
} from 'chart.js'
import { Line } from 'vue-chartjs'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Colors
)

const graphData = ref()

const options = {
  responsive: true,
}

const props = defineProps({
  workoutType: {
    type: String,
    default: ""
  },
  workout: {
    type: String,
    default: ""
  },
  data: {
    type: Array,
    default: undefined
  },
  inTable: {
    type: Boolean,
    default: false
  }
});

interface DataPoint {
  label: string;
  data: {
    x: string | number;
    y: string | number;
  }[] | undefined;
  tension: number;
}

async function getWorkouts(workout_type: string | undefined): Promise<string[] | undefined> {
  try {
    let records: RecordModel[] | any

    records = await pb.collection('workouts').getFullList({
      filter: `workout_type = '${workout_type}'`
    })
    return records.map((item: RecordModel) => item.workout)
  } catch (error) {
    console.log(error)
  }
}

async function getData(workout: string) {
  try {
    type Results = {
      id: number;
      workout: string;
      workout_type: string;
      value: number;
      created: string;
    }

    let records: Results[];

    records = await pb.collection('get_workout_entries_graph').getFullList({
      filter: `workout = '${workout}'`
    })

    return records.map((item) => (
      {
        x: item.created.toString().split(' ')[0],
        y: item.value
      }
    ))
  } catch (error) {
    console.log(error)
  }
}

async function buildDataSets(workouts: Array<string> | undefined): Promise<DataPoint[]> {
  const dataSet: DataPoint[] = []

  if (workouts !== undefined) {
    for (let i = 0; i < workouts.length; i++) {
      const data = await getData(workouts[i])
      dataSet.push({
        label: workouts[i],
        data: data,
        tension: 0.1
      })
    }
  }

  return dataSet
}

async function getEntries() {
  try {

    let data: DataPoint[]

    const dates: string[] = []

    if (props.data != undefined && props.workout != undefined) {
      data = []
      const tempData: { x: string; y: string }[] | undefined = []
      props.data.forEach((value: any) => {
        tempData.push({
          x: value.entryDate.toString().split(' ')[0],
          y: value.value
        })
      })

      data = [{
        label: props.workout.toString(),
        data: tempData,
        tension: 0.1
      }]

      // get the dates
      data?.[0]?.data?.forEach((value) => {
        dates.push(value.x.toString())
      })

      graphData.value = {
        labels: dates,
        datasets: data
      }
    } else {
      if (props.workout == '') {
        const workouts = await getWorkouts(props.workoutType)
        data = await buildDataSets(workouts)

        interface Date {
          id: number;
          workout_type: string;
          created: string;
        }

        let apiDates: Date[] = []

        try {
          apiDates = await pb.collection('get_workout_entries_graph_dates').getFullList({
            filter: `workout_type = '${props.workoutType}'`
          })
        } catch (error) {
          console.log(error)
        }

        apiDates.forEach((value) => {
          dates.push(value.created)
        })

        graphData.value = {
          labels: dates,
          datasets: data,
        }
      } else {
        data = await buildDataSets([props.workout])

        // get the dates
        data?.[0]?.data?.forEach((value) => {
          dates.push(value.x.toString())
        })

        graphData.value = {
          labels: dates,
          datasets: data
        }
      }
    }


  } catch (err) {
    console.log(err)
  }
}

onMounted(() => {
  getEntries()
})

</script>
