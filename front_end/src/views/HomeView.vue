<template>
  <div class="home">
    <simple-time id="time"></simple-time>
    <search></search>
    <web-short-cut></web-short-cut>
    <tool-bar></tool-bar>
    <!-- <setting></setting> -->
    <!-- <location></location> -->
  </div>
</template>

<script>
import Setting from '@/components/Setting.vue'
import Time from '@/components/Time.vue';
import WebShortCut from '@/components/WebShortCut.vue';
import Search from '@/components/Search.vue';
import SimpleTime from '@/components/SimpleTime.vue';
import Location from '@/components/Location.vue';
import ToolBar from '@/components/ToolBar.vue';


export default {
  name: 'HomeView',
  components: {
    Setting,
    Time,
    WebShortCut,
    Search,
    SimpleTime,
    Location,
    ToolBar,
  },
}
</script>

<style scoped>
  #time{
    margin-top: 100px;
  }
  .home{
    height: 100vh;
    background-color: #0F612F;
    background-image: url(../assets/bg_bench.jpg);
    background-size: 100vw 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>