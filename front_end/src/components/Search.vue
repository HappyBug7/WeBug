<template>
  <div class="container">
    <div id="backBlur" :style="{display:state}" @click="blurFunction"></div>
    <div class="searchArea" :style="{zIndex:index}">
      <div class="searchBox" id="searchBox" :style="{zIndex:index,boxShadow:shadow,backdropFilter:filter}">
        <button class="searchButton" @click="searchFnuction">
          <i class='bx bx-search' id="searchIcon"></i>
        </button>
        <input type="text" id="mainSearch" @focus="focusFunction" v-model="q" :style="{zIndex:index}" @keyup.enter="searchFnuction" @input="autoCompelete">
      </div>
    </div>
    <div class="suggestion" :style="{display:isOpen? 'block':'none'}">
      <div class="suggests" v-for="(suggest,index) in this.s_data" :style="getStyle(index)" @click="suggestFunc(suggest)">{{ suggest }}</div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name:"Search",
  data(){
    return{
      q : "",
      index : 10,
      state : "none",
      shadow : "0px 0px 40px 1px rgba(00,00,00,0.2)",
      filter : "blur(5px)",
      s_data : [],
      last : "",
      isOpen:false
    }
  },
  methods:{
    searchFnuction(){
      window.location.href="https://cn.bing.com/search?q="+this.q
    },
    focusFunction(){
      this.state="block"
      this.index = 100
      // this.shadow = "none"
      this.shadow = "0px 0px 40px 1px rgba(0,0,0,0.2)"
      this.filter = "none"
      if(this.q != ""){
        var url = "/api/search?q="+this.q
        axios.get(url)
        .then(res => {
          // console.log(res.data.data.Data)
          this.s_data = res.data.data.Data
        })
        .catch(err => console.log(err))
      }
      this.isOpen = true
    },
    blurFunction(){
      this.state="none"
      this.index = 10
      this.shadow = "0px 0px 40px 1px rgba(0,0,0,0.2)"
      this.filter = "blur(5px)"
      this.isOpen = false
    },
    autoCompelete(){
      if(this.q != this.last){
        var url = "/api/search?q="+this.q
        axios.get(url)
        .then(res => {
          // console.log(res.data.data.Data)
          this.s_data = res.data.data.Data
        })
        .catch(err => console.log(err))
        this.last = this.q
      }
    },
    getStyle(index){
      // if(index == 0){
      //   return{
      //     borderRadius : '10px 10px 0 0'
      //   }
      // }
      // if(index == this.s_data.length -1){
      //   return{
      //     borderRadius : '0 0 10px 10px'
      //   }
      // }
      return{
        borderRadius : '0'
      }
    },
    suggestFunc(quest){
      window.location.href="https://cn.bing.com/search?q="+quest
    }
  }
}
</script>

<style scoped>
  .searchArea
  {
    background-attachment:scroll;
    width: 100vw;
    display: flex;
    align-items: center;
    flex-direction: column;
    background-repeat: no-repeat;
  }
  .searchBox
  {
    width: 400px;
    height: 40px;
    border-radius: 20px;
    background-color: transparent;
    border: 1px solid white;
    display: flex;
    align-items: center;
    flex-direction: row-reverse;
    transition: 0.5s;
  }
  #mainSearch
  {
    width: 325px;
    height: 30px;
    background-color: transparent;
    border-style: none;
    outline: none;
  }
  #backBlur
  {
    display: none;
    position: fixed;
    top: 0;
    width: 100vw;
    height: 100vh;
    margin: 0;
    padding: 0;
    backdrop-filter: blur(5px);
    z-index: 20;
  }
  .searchButton
  {
    border-style: none;
    background-color: transparent;
    font-size: 20px;
    text-align: center;
    width: 30px;
    height: 30px;
    border-radius: 50%;
    margin: 10px;
  }
  .searchButton:hover
  {
    background-color: rgba(255, 255, 255, 0.2);
  }
  .container{
    position: relative;
    margin: 20px;
  }
  #searchIcon{
    
  }
  #compelete{
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .suggestion{
    left: 50%;
    transform: translateX(-200px);
    width: 400px;
    position: absolute;
    display: flex;
    flex-direction: column;
    z-index: 200;
    backdrop-filter: blur(15px);
    border-radius: 10px;
    overflow: hidden;
    box-shadow: 2px 2px 14px 1px rgba(00,00,00,0.2);
  }
  .suggests{
    background-color: transparent;
    color: whitesmoke;
    transition: all 0.2s;
    text-decoration: none;
    user-select: none;
    z-index: 201;
    width: 420px;
    transform: translateX(-10px);
  }
  .suggests:hover{
    background-color: rgba(240, 248, 255, 0.209);
    transform: translateX(0);
  }
</style>