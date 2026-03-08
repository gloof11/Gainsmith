<template>
  <UpperNavBar value=2 />
  <WorkoutEntriesTable workoutType="Push" />
  <WorkoutEntriesTable workoutType="Pull" />
  <WorkoutEntriesTable workoutType="Legs" />
  <WorkoutEntriesTable workoutType="Other" />
  <div class="d-flex justify-space-between">
    <AddWorkoutEntryModal />
    <ArchiveProgramModal />
  </div>
</template>

<script lang="ts" setup>
</script>
