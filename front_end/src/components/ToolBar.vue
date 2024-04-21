<template>
  <div class="container">
    <!-- <music></music> -->
    <add-short-cut></add-short-cut>
    <schedule></schedule>
    <setting></setting>
  </div>
</template>

<script>
import AddShortCut from './AddShortCut.vue'
import Music from './Music.vue'
import Schedule from './Schedule.vue'
import Setting from './Setting.vue'
export default {
  components: { Schedule, Music, Setting, AddShortCut },

}
</script>

<style scoped>

</style>