<template>
  <UpperNavBar value="3" />
  <v-container>
    <WeightGraph />
    <br>
    <WorkoutGraph workoutType="Push" />
    <br>
    <WorkoutGraph workout-type="Pull" />
    <br>
    <WorkoutGraph workout-type="Legs" />
  </v-container>
</template>

<script lang="ts" setup>
</script>
