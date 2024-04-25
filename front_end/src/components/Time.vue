<template>
  <div id="time">
    <div id="container">
      <div id="timeDisplay">{{ time }}</div>
      <div id="dateDisplay">{{ date }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name:"Time",
  methods:{
    getTime(){
      const date = new Date();
      let dayToday = date.getDate();
      let monthToday = date.getMonth() + 1;
      let hourNow = date.getHours();
      let minuteNow = date.getMinutes();
      let secondNow = date.getSeconds();
      if(secondNow < 10)
      {
        secondNow = "0" + String(secondNow);
      }
      if(minuteNow < 10)
      {
        minuteNow = "0" + String(minuteNow);
      }
      let dateToday = String(monthToday) + "月" + String(dayToday) + "日";
      let timeNow = String(hourNow) + ":" + String(minuteNow) + ":" + String(secondNow);
      this.time = timeNow;
      this.date = dateToday;
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
      date:""
    }
  }
}
</script>

<style scoped>
  #time{
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 20vw;
    height: 20vh;
    border: 5px solid black;
    border-radius: 10px;
    background-color: transparent;
  }

  #timeDisplay{
    user-select: none;
    font-size: 30px;
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
</style>