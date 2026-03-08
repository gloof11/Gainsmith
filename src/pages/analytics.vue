<template>
  <UpperNavBar value=3 />
  <v-container>
    <WeightGraph />
    <br>
    <WorkoutGraph workoutType="Push" />
    <br>
    <WorkoutGraph workoutType="Pull" />
    <br>
    <WorkoutGraph workoutType="Legs" />
  </v-container>
</template>

<script lang="ts" setup>
</script>
