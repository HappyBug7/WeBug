<template>
  <div>
    <input type="radio" id="option1" value="option1" v-model="selectedOption">
    <label for="option1">Option 1</label>

    <input type="radio" id="option2" value="option2" v-model="selectedOption">
    <label for="option2">Option 2</label>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedOption: 'option1', // 设置初始选中的选项
    };
  },
  watch: {
    selectedOption(newVal) {
      console.log('选中的选项是:', newVal);
    },
  },
};
</script>