<template>
  <div class="container">
    <div id="todo_list">
      <div id="filter" :style="{ display: style, backgroundColor: filterBackgroundColor }" @mouseleave="FilterCloseFunc" @click="SwitchFunc">
        <i class='bx bx-x' id="X"></i>
      </div>
      <i class='bx bx-list-ul' id="todo_list_icon" @click="SwitchFunc" @mouseover="FilterOpenFunc"></i>
    </div>
  </div>
  <draggable-resizable class="drag" :resizable="false" :x=-700 :y=-350 :z=11 :w=0 :h=0>
    <div class="list_container" :style="{ display: isOpen ? 'block' : 'none', backgroundColor: listContainerBackgroundColor }">
      <div id="head" :style="{ backgroundColor: headBackgroundColor }">TODO List!!!</div>
    </div>
  </draggable-resizable>
</template>

<script>
import "../../assets/style.css";
import VueDraggableResizable from 'vue-draggable-resizable';

export default {
  name: 'todo-list',
  data() {
    return {
      isOpen: false,
      style: 'none',
    }
  },
  props: {
    position: {
      type: Array,
      default: function () {
        return [];
      }
    },
    filterBackgroundColor: {
      type: String,
      default: '#fff5f39d' // Default color for filter background
    },
    listContainerBackgroundColor: {
      type: String,
      default: '#249aa7' // Default color for list container background
    },
    headBackgroundColor: {
      type: String,
      default: '#ffd334be' // Default color for head background
    },
    todoListBackgroundColor: {
      type: String,
      default: '#fff5f39d' // Default color for todo list background
    }
  },
  components: {
    'draggable-resizable': VueDraggableResizable
  },
  methods: {
    SwitchFunc() {
      this.isOpen = !this.isOpen;
    },
    FilterOpenFunc() {
      if (this.isOpen) {
        this.style = 'block';
      }
    },
    FilterCloseFunc() {
      this.style = 'none';
    },
  },
}
</script>

<style scoped>
.container {
  position: relative;
}

#todo_list {
  margin: 20px;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  position: fixed;
  top: v-bind(position[0]);
  left: v-bind(position[1]);
  box-shadow: 2px 2px 14px 1px rgba(00, 00, 00, 0.2);
  transition: all 0.3s;
  display: flex;
  flex-direction: column;
  z-index: 10;
  background-color: v-bind(todoListBackgroundColor);
}

#todo_list:hover {
  transform: scale(1.1);
  box-shadow: none;
  cursor: pointer;
}

#todo_list_icon {
  font-size: 40px;
  color: #249aa7;
  margin: auto;
}

#filter {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

#X {
  font-size: 50px;
  margin: auto;
}

.input_container {
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 10px;
  position: fixed;
  top: 10px;
  left: 80px;
}

.inputDiv {
  width: 300px;
  height: 180px;
  backdrop-filter: blur(10px);
  border-radius: 10px;
  display: flex;
  flex-direction: column;
}

.cross {
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

.cross:hover {
  box-shadow: 2px 2px 14px 1px rgba(0, 0, 0, 0.2);
  backdrop-filter: blur(10px);
  color: red;
  cursor: pointer;
}

.list_container {
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 20px;
  position: fixed;
  top: 10px;
  left: 80px;
  width: 300px;
  height: 100px;
  backdrop-filter: blur(10px);
}

#head {
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  margin: auto;
  border-radius: 20px;
  width: 100%;
  height: 50px;
  font-size: 30px;
}
</style>
