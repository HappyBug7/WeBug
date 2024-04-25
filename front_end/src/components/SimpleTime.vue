<template>
  <div id="time">
    <div id="container">
      <div id="timeDisplay" @click="SwitchFunc">{{ time }}</div>
      <weather id="weather"></weather>
      <calendar :isopen=this.isOpen :style="{display : isOpen ? 'block' : 'none'}" @passtofather="myfn"></calendar>
    </div>
  </div>
</template>

<script>
import Weather from './Weather.vue';
import Calendar from '@/components/Calendar.vue';

export default {
  name:"SimpleTime",
  methods:{
    getTime(){
      const date = new Date();
      let hourNow = date.getHours();
      let minuteNow = date.getMinutes();
      if(minuteNow < 10)
      {
        minuteNow = "0" + String(minuteNow);
      }
      let timeNow = String(hourNow) + ":" + String(minuteNow);
      this.time = timeNow;
    },
    myfn(value){
      this.isOpen = value
    },
    SwitchFunc(){
      this.isOpen = !this.isOpen
    }
  },
  mounted(){
    this.timer = setInterval(() => {
      this.getTime()
    }, 1);
  },
  data(){
    return {
      time:"",
      date:"",
      isOpen:false,
    }
  },
  components: {
    Weather,
    Calendar,
  }
}
</script>

<style scoped>
  #time{
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 20vw;
    background-color: transparent;
  }

  #timeDisplay{
    user-select: none;
    font-size: 40px;
    transition: all 0.35s;
  }

  #timeDisplay:hover{
    transform: scale(1.2);
  }

  #dateDisplay{
    user-select: none;
    font-size: 30px;
  }

  #container{
    display: flex;
    flex-direction: column;
    margin: auto;
  }

  #weather{
    font-size: 30px;
  }
</style>