<template>
  <div v-if="loading == false">
    <div class="d-flex align-center">
      <h1>Weight</h1>
      <div v-if="entries != undefined">
        <div v-if="entries[entries.length-1].weight > entries[entries.length-2].weight">
          <v-icon  icon="mdi-arrow-up-bold" color="green" size="x-large" />
          Weight has gone up!
        </div>
        <div v-else-if="entries[entries.length-1].weight < entries[entries.length-2].weight">
          <v-icon icon="mdi-arrow-down-bold" color="red" size="x-large" />
          Weight has gone down!
        </div>
      </div>
    </div>

    <div class="bg-grey-lighten-2 rounded">
      <Line :data="graphData" :options="options" />
    </div>
  </div>
  <div v-else>
    <v-progress-circular indeterminate :size="89" :width="8" class="mb-8" />
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue'
  import { pb } from '../lib/pocketbaseClient'

  import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
  } from 'chart.js'
  import { Line } from 'vue-chartjs'

  ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
  )

  const entries = ref()
  const loading = ref(false)

  const dates: Array<string> = []
  const weights: Array<number> = []

  const graphData = ref({
    labels: [''],
    datasets: [
      {
        label: 'Data One',
        backgroundColor: '#f87979',
        data: [0],
        tension: 0.1
      },
    ],
  })

  const options = {
    responsive: true,
  }

  async function getEntries() {
    try {
      loading.value = true
      const records = await pb.collection('weight_entries').getFullList({
        sort: 'date'
      })
      entries.value = records

      for (let i=0; i<records.length; i++) {
        dates.push(records[i].date.toString().split(' ')[0])
        weights.push(records[i].weight)
      }

      graphData.value = {
        labels: dates,
        datasets: [
          {
            label: 'Weight',
            backgroundColor: '#f87979',
            data: weights,
            tension: 0.1
          },
        ],
      }
    } catch (err) {
      console.log(err)
    } finally {
      loading.value = false
    }
  }

  onMounted(() => {
    getEntries()
  })
</script>
