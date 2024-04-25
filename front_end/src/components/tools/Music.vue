<template>
  <div class="container">
    <div id="music_container">
      <div id="filter" :style="{display:style}" @mouseleave="FilterCloseFunc" @click="SwitchFunc">
        <i class='bx bx-x' id="X"></i>
      </div>
      <i class='bx bx-music' id="music_icon" @click="SwitchFunc" @mouseover="FilterOpenFunc"></i>
    </div>
    <iframe :style="{display : isOpen ? 'block' : 'none'}" class="music" frameborder="no" marginwidth="0" marginheight="0" src="//music.163.com/outchain/player?type=0&id=9306895894&auto=0&height=430">
    </iframe>
  </div>
</template>

<script>
export default {
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
  }
}
</script>

<style scoped>
  .container{
    position: relative;
  }
  .music{
    background-color: transparent;
    position: fixed;
    top: 0;
    left: 80px;
    width: 330px; 
    height: 450px;
    user-select: none;
  }

  #music_icon{
    font-size: 40px;
    color: #FF7171;
    margin: auto;
  }

  #music_container{
    margin: 20px;
    width: 50px;
    height: 50px;
    background-color: #fff5f39d;
    border-radius: 50%;
    position: fixed;
    top: 0;
    left: 0;
    box-shadow: 2px 2px 14px 1px rgba(00,00,00,0.2);
    transition: all 0.3s;
    display: flex;
    flex-direction: column;
    z-index: 10;
  }
  #music_container:hover{
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
</style>