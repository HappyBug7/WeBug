<template>
  <div class="container">
    <div id="schedule_container">
      <div id="filter" :style="{display:style}" @mouseleave="FilterCloseFunc" @click="SwitchFunc">
        <i class='bx bx-x' id="X"></i>
      </div>
      <i class='bx bxs-book-bookmark' id="schedule_icon" @click="SwitchFunc" @mouseover="FilterOpenFunc"></i>
    </div>
    <draggable-resizable class="drag" :resizable="false" :x=-700 :y=-350 :z=11 :w=0 :h=0>
      <div class="classes_ontainer" :style="{display : isOpen ? 'block' : 'none'}">
        <div class="switch_bar" @click="SwitchFunc_">
          <div class="dot" :style="{backgroundColor: switch_ == 1 ? 'whitesmoke' : 'rgb(120, 120, 120)'}"></div>
          <div class="dot" :style="{backgroundColor: switch_ == 2 ? 'whitesmoke' : 'rgb(120, 120, 120)'}"></div>
        </div>
        <div class="cross" @click="SwitchFunc">
          <i class='bx bx-x'></i>
        </div>
        <div class="Schedule">
          <div :style="{display : switch_ == 1 ? 'block' : 'none'}">
            <div v-for="class_ in this.data" class="classes">
              <div class="flex_row">
                <div class="flex_column left">
                  <div>{{ Time[0][class_.class] }}</div>
                  <div>{{ class_.duration }}</div>
                </div>
                <div class="flex_column right">
                  <div>{{ class_.name }}</div>
                  <div>{{ class_.place }}</div>
                </div>
              </div>
            </div>
          </div>
          <div :style="{display : switch_ == 2 ? 'block' : 'none'}">
            <div class="flex_row">
              <div class="flex_column">
                <div v-for="time_ in this.time" class="time">
                  <div>{{ Time[0][time_] }}</div>
                </div>
              </div>
              <div v-for="day in this.data_apd" class="flex_column">
                <div v-for="class_ in day" class="flex_column simp_class">
                  <div>{{ class_.name }}</div>
                  <div>{{ class_.duration }}</div>
                </div>
              </div>
            </div>
          </div>
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
      Time:[{
        12:"8:00-9:50",
        34:"10:10-12:00",
        56:"14:00-15:50",
        78:"16:10-18:00",
        910:"19:10-21:00"
      },{
        12:"8:00-9:50",
        34:"10:10-12:00",
        56:"14:30-16:20",
        78:"16:40-18:30",
        910:"19:40-21:30"
      }],
      Day:0,
      data:[],
      data_apd:[],
      time:[12,34,56,78,910],
      switch_ : 1,
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
    }
  },
  mounted(){
    let now = new Date();
    this.Day = now.getDay();
    // this.Day = 1
    axios.get('/api/class?day='+this.Day)
    .then(res => {
      this.data = res.data.data
    })
    .catch(err => console.log(err))

    for(let i = 0; i < 7; i++){
      let data_temp = []
      for(let j = 0; j < 5; j++){
        axios.get('/api/classes?day='+(i+1)+'&time='+this.time[j])
        .then(res => {
          data_temp.push(res.data.data)
        })
        .catch(err => console.log(err))
      }
      this.data_apd.push(data_temp)
    }
    console.log(this.data_apd)
  },
  components: {
    'draggable-resizable': VueDraggableResizable
  }
  // mounted(){
  //   axios.get('/api/weather')
  //   .then(res => {
  //     // console.log(res)
  //     this.code = res.data.data
  //   })
  //   .catch(err => console.log(err))
  // }
}
</script>

<style scoped>
  *{
    user-select: none;
  }
  .container{
    position: relative;
  }
  #schedule_container{
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
  #schedule_container:hover{
    transform: scale(1.1);
    box-shadow: none;
    cursor: pointer;
  }

  #schedule_icon{
    font-size: 40px;
    color: #1b6983;
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
  .classes_ontainer{
    display: flex;
    flex-direction: column;
    align-items: center;
    border-radius: 10px;
    position: fixed;
    top: 10px;
    left: 80px;
  }
  .Schedule{
    backdrop-filter: blur(10px);
    border-radius: 10px;
    box-shadow: 2px 2px 14px 1px rgba(0,0,0,0.2);
  }
  .flex_row{
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  .flex_column{
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .classes{
    width: 300px;
  }
  .left{
    color: #4B4737;
    width: 100px;
    background-color: #ffeba5be;
    border-radius: 10px;
    margin: 3px;
  }
  .right{
    color: #4B4737;
    width: 200px;
    background-color: #fef3cabe;
    border-radius: 10px;
    margin: 3px;
  }
  .simp_class{
    width: 120px;
    height: 62px;
    background-color: #fef3cabe;
    border-radius: 10px;
    margin: 3px;
    color: #4B4737;
    div{
      margin: auto;
    }
  }
  .time{
    width: 100px;
    height: 62px;
    background-color: #ffeba5be;
    color: #4B4737;
    border-radius: 10px;
    margin: 3px;
    display: flex;
    div{
      margin: auto;
    }
  }
  .switch_bar{
    width: 25px;
    height: 16px;
    border-radius: 8px;
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-bottom: 3px;
  }
  .switch_bar:hover{
    box-shadow: 2px 2px 14px 1px rgba(0,0,0,0.2);
    backdrop-filter: blur(10px);
    cursor: pointer;
  }
  .dot{
    background-color: whitesmoke;
    border-radius: 50%;
    height: 6px;
    width: 6px;
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
    right: 0;
    top: 0;
  }
  .cross:hover{
    box-shadow: 2px 2px 14px 1px rgba(0,0,0,0.2);
    backdrop-filter: blur(10px);
    color: red;
    cursor: pointer;
  }
</style>