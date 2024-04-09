<template>
  <div class="web">
    <div id="shortcuts">
      <a v-for="shortcut in this.s_data" :href="shortcut.LinkOri" :title="shortcut.LinkTitle" class="shortcut" target="_blank">
        <div class="container">
          <img :src="shortcut.LinkIcon" class="imgs"/>
        </div>
        <span>{{ this.calculateLength(shortcut.LinkTitle) > this.maxlen ? shortcut.LinkTitle.slice(0, this.getMaxRange(shortcut.LinkTitle))+"..." : shortcut.LinkTitle }}</span>
      </a>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name:"WebShortCut",
  data() {
    return {
      // link1_icon : '',
      // link1: '',
      // tit1: '',
      // maxlen: 9,
      // curr: 1
      link_icon : [],
      link : [],
      tit : [],
      s_data : [],
      maxlen : 9
    };
  },
  methods: {
    calculateLength(str) {
      const pattern = /[\u4e00-\u9fa5]/g;
      const replacedStr = str.replace(pattern, 'aa');

      return replacedStr.length;
    },
    getMaxRange(str) {
      for(let i = 0; i < str.length; i++)
      {
        if(this.calculateLength(str.slice(0,i)) > this.maxlen){
          return i
        }
      }
      return str.length
    }
  },
  mounted() {
    axios.get('/api/shortcuts')
    .then(res => {
      this.s_data = res.data.data
      // console.log(this.s_data)
    })
    .catch(err => console.log(err))
  }
}
</script>

<style scoped>

  #shortcuts{
    display: flex;
    user-select: none;
  }
  .shortcut{
    width: 90px;
    height: 90px;
    display: flex;
    flex-direction: column;
    align-items: center;
    text-decoration: none;
    color: #8FD4B3;
    font-size: 13px;
    transition: all 0.2s;
  }
  .container{
    width: 50px;
    height: 50px;
    display: flex;
    flex-direction: row;
    align-items: center;
    background-color: #cdfbe4a5;
    border-radius: 10px;
    box-shadow: 0 0 14px 1px rgba(00,00,00,0.2);
  }

  .shortcut:hover{
    transform: scale(1.2);
    width: 99px;
  }
  .shortcut:hover > div{
    box-shadow: none;
    transition: 0.5s;
  }
  .imgs{
    width: 20px;
    height: 20px;
    margin: auto;
  }

</style>