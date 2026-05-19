<template>
  <v-layout
    class="ma-2"
    style="height: 56px;"
  >
    <v-bottom-navigation
      class="rounded-xl overflow-auto pl-16 pl-sm-0"
      style="scrollbar-width: none;"
      active
    >
      <v-btn
        to="/"
        :active="false"
      >
        <v-icon>mdi-home-variant</v-icon>
        <span>Home</span>
      </v-btn>

      <v-btn
        to="/workouts"
        :active="false"
      >
        <v-icon>mdi-dumbbell</v-icon>
        <span>Workouts</span>
      </v-btn>

      <v-btn
        to="/workout-entries"
        :active="false"
      >
        <v-icon>mdi-weight</v-icon>
        <span>Weight</span>
      </v-btn>

      <v-btn
        to="/analytics"
        :active="false"
      >
        <v-icon>mdi-google-analytics</v-icon>
        <span>Analytics</span>
      </v-btn>

      <v-btn
        to="/archived-programs"
        :active="false"
      >
        <v-icon>mdi-archive-alert</v-icon>
        <span>Archives</span>
      </v-btn>
    </v-bottom-navigation>
  </v-layout>
</template>

<script lang="ts" setup>
  const value = defineModel("value")
</script>
