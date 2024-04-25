<template>
  <div class="container">
    <div id="add_shortcut">
      <div id="filter" :style="{display:style}" @mouseleave="FilterCloseFunc" @click="SwitchFunc">
        <i class='bx bx-x' id="X"></i>
      </div>
      <i class='bx bxs-plane-take-off' id="shortcut_icon" @click="SwitchFunc" @mouseover="FilterOpenFunc"></i>
    </div>
    <draggable-resizable class="drag" :resizable="false" :x=-700 :y=-350 :z=11 :w=0 :h=0>
      <div class="input_container" :style="{display : isOpen ? 'block' : 'none'}">
        <div class="cross" @click="SwitchFunc">
          <i class='bx bx-x'></i>
        </div>
        <div class="inputDiv">
          <input type="text" placeholder="Enter the URL" style="margin: 5px; border-radius: 5px; border-style: none; height: 60px; padding: 5px; background-color: #ffeba5be;" v-model="url">
          <input type="text" placeholder="Enter the name" style="margin: 5px; border-radius: 5px; border-style: none; height: 60px; padding: 5px; background-color: #ffeba5be;" v-model="name">
          <button id="submitButton" @click="SubmitFunc">Add</button>
        </div>
        
      </div>
    </draggable-resizable>
  </div>
</template>

<script>
import axios from 'axios';
import "../../assets/style.css";
import VueDraggableResizable from 'vue-draggable-resizable';
export default {
  data(){
    return{
      isOpen: false,
      style: 'none',
      switch_ : 1,
      url: '',
      name: '',
    }
  },
  props:{
    position:{
      type: Array,
      default: function () {
        return []
      }
    }
  },
  components: {
    'draggable-resizable': VueDraggableResizable
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
    },
    SwitchFunc_(){
      this.switch_ == 1? this.switch_ = 2 : this.switch_ = 1
      this.url = ''
      this.name = ''
    },
    SubmitFunc(){
      const shortcut = {
        Link: this.url,
        Title: this.name
      }
      axios.post('/api/shortcut_add?', shortcut)
      .then(res => {
        this.data = res.data.data
      })
      .catch(err => console.log(err))
      alert('Shortcut added successfully')
      this.SwitchFunc()
      // console.log(shortcut)
    }
  },
}
</script>

<style scoped>
  *{
    font-family: 'comic sans ms';
  }
  .container{
    position: relative;
  }

  #add_shortcut{
    margin: 20px;
    width: 50px;
    height: 50px;
    background-color: #fff5f39d;
    border-radius: 50%;
    position: fixed;
    top: v-bind(position[0]);
    left: v-bind(position[1]);
    box-shadow: 2px 2px 14px 1px rgba(00,00,00,0.2);
    transition: all 0.3s;
    display: flex;
    flex-direction: column;
    z-index: 10;
  }
  #add_shortcut:hover{
    transform: scale(1.1);
    box-shadow: none;
    cursor: pointer;
  }

  #shortcut_icon{
    font-size: 40px;
    color: #1b9b74;
    margin: auto;
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

  .input_container{
    display: flex;
    flex-direction: column;
    align-items: center;
    border-radius: 10px;
    position: fixed;
    top: 10px;
    left: 80px;
  }

  .inputDiv{
    width: 300px;
    height: 180px;
    /* background-color: #ffd334be; */
    backdrop-filter: blur(10px);
    border-radius: 10px;
    display: flex;
    flex-direction: column;
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
    right: 0;
    top: -20px;
  }

  .cross:hover{
    box-shadow: 2px 2px 14px 1px rgba(0,0,0,0.2);
    backdrop-filter: blur(10px);
    color: red;
    cursor: pointer;
  }

  #submitButton{
    margin: 5px;
    border-radius: 5px;
    border-style: none;
    height: 60px;
    padding: 5px;
    background-color: #ffd334be;
    color: #CDFBE4;
    box-shadow: 0 0 14px 1px rgba(00,00,00,0.2);
  }

  #submitButton:hover{
    color: #2f61ea;
    cursor: pointer;
  }
</style>