<template>
  <div class="container">
    <music></music>
    <schedule></schedule>
    <setting></setting>
  </div>
</template>

<script>
import Music from './Music.vue'
import Schedule from './Schedule.vue'
import Setting from './Setting.vue'
export default {
  components: { Schedule, Music, Setting },

}
</script>

<style scoped>

</style>