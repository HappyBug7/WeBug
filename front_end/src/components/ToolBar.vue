<template>
  <div class="container">
    <!-- <music></music> -->
    <add-short-cut :position="['0px', '0px']"></add-short-cut>
    <schedule :position="['70px', '0px']"></schedule>
    <setting :position="['140px', '0px']"></setting>
    <todo-list :position="['210px', '0px']"></todo-list>
  </div>
</template>

<script>
import AddShortCut from './tools/AddShortCut.vue'
import Music from './tools/Music.vue'
import Schedule from './tools/Schedule.vue'
import Setting from './tools/Setting.vue'
import TodoList from './tools/TodoList.vue'
export default {
  components: { Schedule, Music, Setting, AddShortCut, TodoList },
  data(){
    return{
    }
  }
}
</script>

<style scoped>

</style>