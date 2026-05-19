<template>
  <UpperNavBar value="2" />
  <WorkoutEntriesTable workout-type="Push" />
  <WorkoutEntriesTable workout-type="Pull" />
  <WorkoutEntriesTable workout-type="Legs" />
  <WorkoutEntriesTable workout-type="Other" />
  <div class="d-flex justify-space-between">
    <AddWorkoutEntryModal />

    <ArchiveProgramModal />
  </div>
</template>
