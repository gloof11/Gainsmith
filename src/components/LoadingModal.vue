<template>
  <v-dialog
    v-model="props.loading"
    width="auto"
    persistent
  >
    <v-card
      :width="`${width/1.2}`"
      max-width="300"
      :prepend-icon="props.icon"
      :title="props.title"
    >
      <div class="text-center mt-2 mb-8">
        <v-progress-circular indeterminate :size="89" :width="8" class="mb-8" />
        <p>
          {{ props.status }}
        </p>
      </div>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
  import { useDisplay } from 'vuetify'
  const { width } = useDisplay()

  const props = defineProps({
    title: {
      type: String,
      default: "Loading"
    },
    icon: {
      type: String,
      default: ""
    },
    status: {
      type: String,
      default: ""
    },
    loading: {
      type: Boolean,
      default: true
    }
  });
</script>
