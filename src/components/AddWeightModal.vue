<template>
  <div>
    <v-btn
      density="compact"
      icon="mdi-plus"
      size="large"
      color="green-darken-3"
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
      title="Add a Weight"
    >
      <v-form @submit.prevent="submitWeight">
        <v-container>
          <v-date-input
            v-model="weightObject.date"
            prepend-icon=""
            variant="outlined"
            label="Date"
            persistent-placeholder
            :rules="rules"
          />

          <v-text-field
            v-model="weightObject.weight"
            variant="outlined"
            label="Weight (lb)"
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
  const weightObject = ref({
    date: null,
    weight: 0
  })

  const submitWeight = async () => {
    try {
      const records = await pb.collection('weight_entries').create({
        "date": weightObject.value.date,
        "weight": weightObject.value.weight
      })
      console.log(records)
    } catch (err) {
      console.log(err)
    }

    location.reload()
  }
</script>
