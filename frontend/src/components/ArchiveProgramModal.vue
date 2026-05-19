<template>
  <div class="mr-4 mt-4 mb-4">
    <v-btn color="blue" variant="outlined" text="Archive" @click="dialog = true" />
  </div>

  <v-dialog v-model="dialog" width="auto">
    <v-card :width="`${width / 1.2}`" max-width="400" prepend-icon="mdi-archive-alert" title="Archive Program">
      <v-form @submit.prevent="loading = true; submitArchive()">
        <v-container>
          <p class="mb-2">
            WARNING! This will remove the current program.
          </p>
          <v-text-field v-model="archiveFields.name" variant="outlined" label="Program Name" persistent-placeholder
            :rules="rules" />

          <v-date-input v-model="archiveFields.from" prepend-icon="" variant="outlined" label="From Date"
            persistent-placeholder :rules="rules" />

          <v-date-input v-model="archiveFields.to" prepend-icon="" variant="outlined" label="To Date"
            persistent-placeholder :rules="rules" />

          <v-textarea v-model="archiveFields.notes" variant="outlined" label="Notes" persistent-placeholder />

          <v-card v-if="errorMessage" color="red" :text="errorMessage"></v-card>
        </v-container>

        <div class="d-flex justify-end mr-8 mb-4">
          <v-btn color="green" variant="outlined" type="submit" text="Ok" />
        </div>
      </v-form>
    </v-card>
  </v-dialog>
</template>


<script>
import { useDisplay } from 'vuetify'
import pb from '../lib/pocketbaseClient'
import { rules } from '../lib/helpers'
export default {
  data() {
    return {
      dialog: false,
      id: "",
      archiveStatus: "",
      archiveFields: {
        name: "",
        from: "",
        to: "",
        notes: ""
      },
      width: useDisplay().width,
      rules: rules,
      errorMessage: ""
    }
  },
  methods: {
    async submitArchive() {
      try {
        this.id = await pb.send('/api/archive/create', {
          method: 'POST',
          body: this.archiveFields,
        })
        location.assign("/program/?id=" + this.id.id)
      } catch (err) {
        this.errorMessage = err.message
      }

    }
  }
}
</script>
