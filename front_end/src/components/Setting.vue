<template>
  <div class="container">
    <div id="setting_container">
      <div id="filter" :style="{display:style}" @mouseleave="FilterCloseFunc" @click="SwitchFunc">
        <i class='bx bx-x' id="X"></i>
      </div>
      <i class='bx bxs-cog' id="setting_icon" @click="SwitchFunc" @mouseover="FilterOpenFunc"></i>
    </div>
    <draggable-resizable class="drag" :resizable="false" :x=-630 :y=-350 :z=11 :w=0 :h=0>
      <div :style="{display : isOpen ? 'block' : 'none'}">
        <div id="main">
          <div class="cross" @click="SwitchFunc">
            <i class='bx bx-x'></i>
          </div>
          <div id="setting">
            <div class="tabs">
              <input type="radio" id="test1" name="ground" checked>
              <input type="radio" id="test2" name="ground">
              <div class="icon">
                <label for="test1">
                  <i class='bx bx-cloud'></i>
                </label>
                <label for="test2">
                  <i class='bx bx-message-dots'></i>
                </label>
                <div class="indicator"></div>
              </div>
            </div>
            <div class="content"></div>
          </div>
        </div>
      </div>
    </draggable-resizable>
  </div>
</template>

<script>
import "../assets/style.css";
import VueDraggableResizable from 'vue-draggable-resizable';
export default {
  name:"Setting",
  data(){
    return{
      isOpen: false,
      style: 'none'
    }
  },
  methods:{
    SwitchFunc(){
      this.isOpen = !this.isOpen
    },
    FilterOpenFunc(){
      if(this.isOpen){
        this.style = 'block'
      }
    },
    FilterCloseFunc(){
      this.style = 'none'
    }
  },
  components: {
    'draggable-resizable': VueDraggableResizable
  }
}
</script>

<style scoped>
  #main{
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  #setting{
    margin: auto;
    display: flex;
    flex-direction: row;
    align-items: center;
    background-color: transparent;
    height: 350px;
    width: 520px;
    backdrop-filter: blur(10px);
    border-radius: 10px;
    box-shadow: 2px 2px 14px 1px rgba(0,0,0,0.2);
  }
  .tabs{
    display: flex;
    height: 100%;
    align-items: center;
    position: relative;
    input{
      display: none;
      &:nth-child(1):checked ~ .icon label:nth-child(1){
        opacity: 1;
      }
      &:nth-child(2):checked ~ .icon label:nth-child(2){
        opacity: 1;
      }

      &:nth-child(1):checked ~ .icon .indicator{
        top:0;
      }
      &:nth-child(2):checked ~ .icon .indicator{
        top:50%;
      }
    }
  }
  .icon{
    position: relative;
    display: flex;
    flex-direction: column;
    height: 100%;
    label{
      margin: auto;
      flex: 1;
      width: 70px;
      height: 70px;
      color: #D6FADC;
      font-size: 48px;
      line-height: 70px;
      text-align: center;
      opacity: 0.5;
      z-index: 2;
      display: flex;
      flex-direction: column;
      align-items: center;
      label{
        height: 100%;
        margin: auto;
        display: flex;
        align-items: center;
        flex-direction: column;
      }
    }
  }

  .indicator{
    position: absolute;
    left: 0;
    top: 0;
    width: 70px;
    height: 70px;
    background-color: #faebe791;
    border-top-left-radius: 25px;
    border-bottom-left-radius: 25px;
    transition: 0.5s;

    &::before {
      content: " ";
      position: absolute;
      right: 0px;
      top: -30px;
      width: 30px;
      height: 30px;
      background-image: radial-gradient(circle 30px at 0 0, transparent 30px, #faebe791 50%);
    }
    &::after {
      content: " ";
      position: absolute;
      right: 0;
      bottom: -30px;
      width: 30px;
      height: 30px;
      background-image: radial-gradient(circle 30px at 0 100%, transparent 30px, #faebe791 50%);
    }
  }
  .content{
    width: 450px;
    height: 350px;
    background-color: #faebe791;
    border-radius: 10px;
  }
  .container{
    position: relative;
  }

  #setting_icon{
    font-size: 40px;
    color: #919191;
    margin: auto;
  }

  #setting_container{
    margin: 20px;
    width: 50px;
    height: 50px;
    background-color: #fff5f39d;
    border-radius: 50%;
    position: fixed;
    top: 140px;
    left: 0;
    box-shadow: 2px 2px 14px 1px rgba(00,00,00,0.2);
    transition: all 0.3s;
    display: flex;
    flex-direction: column;
    z-index: 10;
  }
  #setting_container:hover{
    transform: scale(1.1);
    box-shadow: none;
  }

  #filter{
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.389);
    border-radius: 50%;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  #X{
    font-size: 50px;
    margin: auto;
  }
  .cross{
    width: 16px;
    height: 16px;
    border-radius: 50%;
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-bottom: 3px;
    position: absolute;
    right: -520px;
    top: -18px;
  }
  .cross:hover{
    box-shadow: 2px 2px 14px 1px rgba(0,0,0,0.2);
    backdrop-filter: blur(10px);
    color: red;
  }
</style>