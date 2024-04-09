<template>
  <a class="container" href="https://weather.com/zh-CN/weather/today/l/CHXX0008:1:CH?Goto=Redirected" target="_blank">
    <i :class="icon_list[this.code]"></i>
  </a>
</template>

<script>
import axios from 'axios';

export default {
  name:"Weather",
  data(){
    return{
      icon_list : {
        200:"bx bx-cloud-light-rain",
        201:"bx bx-cloud-rain",
        202:"bx bx-cloud-rain",
        230:"bx bx-cloud-lightning",
        231:"bx bx-cloud-lightning",
        232:"bx bx-cloud-lightning",
        300:"bx bx-cloud-drizzle",
        301:"bx bx-cloud-drizzle",
        302:"bx bx-cloud-drizzle",
        500:"bx bx-cloud-light-rain",
        501:"bx bx-cloud-light-rain",
        502:"bx bx-cloud-light-rain",
        520:"bx bx-cloud-light-rain",
        521:"bx bx-cloud-light-rain",
        522:"bx bx-cloud-light-rain",
        600:"bx bx-cloud-snow",
        601:"bx bx-cloud-snow",
        602:"bx bx-cloud-snow",
        610:"bx bx-cloud-snow",
        611:"bx bx-cloud-snow",
        612:"bx bx-cloud-snow",
        621:"bx bx-cloud-snow",
        622:"bx bx-cloud-snow",
        623:"bx bx-cloud-snow",
        700:"bx bx-cloud",
        711:"bx bx-cloud",
        721:"bx bx-cloud",
        731:"bx bx-cloud",
        741:"bx bx-cloud",
        751:"bx bx-cloud",
        800:"bx bx-sun",
        801:"bx bx-sun",
        802:"bx bx-sun",
        803:"bx bx-sun",
        804:"bx bx-cloud",
        900:"bx bx-cloud"
      },
      code : 0
    }
  },
  mounted() {
    axios.get('/api/weather')
    .then(res => {
      // console.log(res)
      this.code = res.data.data
    })
    .catch(err => console.log(err))
  }
}
</script>

<style scoped>
  .container{
    transition: all 0.3s;
    color: #CDFBE4
  }
  .container:hover{
    transform: scale(1.2);
  }
</style>