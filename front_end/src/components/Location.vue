<template>
  <button @click="Update" id="update">Weather Update</button>
</template>

<script>
import axios from 'axios';

export default {
  data(){
    return{
    }
  },
  methods:{
    Update(){
      if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(function(position) {
          var latitude = position.coords.latitude;
          var longitude = position.coords.longitude;
          console.log("lon: " + longitude);
          console.log("lat: " + latitude);

          let url = '/api/weather_get?lat=' + latitude.toString() + '&lon=' + longitude.toString()
          axios.post(url)
          .then(res => {
            // console.log(res)
          })
          .catch(err => console.log(err))
        });
      } else {
        console.log("unsupported");
      }
    }
  },
  mounted(){
  }
}
</script>

<style scoped>

#update{
  border-style: none;
  border-radius: 5px;
  width: 120px;
  height: 30px;
  background-color: #d1f5ffc8;
  color: #65deff;
  box-shadow: 0 0 14px 1px rgba(00,00,00,0.2);
}
#update:hover{
  box-shadow: none;
  transition: 0.5s;
}

</style>