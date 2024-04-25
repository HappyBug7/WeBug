<template>
  <div>
    <draggable-resizable class="drag" :resizable="false" :x=49 :y=0 :z=11 :w=0 :h=0>
      <div class="calendar-box">
        <div class="controls">
          <div class="prev" @click="prevMonth">
            <i class='bx bxs-chevron-up'></i>
          </div>
          <div class="next" @click="nextMonth">
            <i class='bx bxs-chevron-down'></i>
          </div>
          <div class="cross" @click="SwitchFunc">
            <i class='bx bx-x'></i>
          </div>
        </div>
        <div class="calendar-wrapper">
          <div class="calendar-toolbar">
            <div class="current">{{ currentDateStr }}</div>
          </div>
          <div class="calendar-week">
            <div
              class="week-item calendarBorder"
              v-for="item of weekList"
              :key="item">
              {{ item }}
            </div>
          </div>
          <div class="calendar-inner">
            <div
              class="calendar-item calendarBorder"
              v-for="(item, index) of calendarList"
              :key="index"
              :class="{
                'calendar-item': true,
                calendarBorder: true,
                'calendar-item-hover': !item.disable,
                'calendar-item-disabled': item.disable,
                'calendar-item-checked':
                  dayChecked && dayChecked.value == item.value,
              }"
              @click="handleClickDay(item)"
            >
              {{ item.date }}
            </div>
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
  data() {
    return {
      showYearMonth: {}, // 显示的年月
      calendarList: [], // 用于遍历显示
      shareDate: new Date(), // 享元模式，用来做 日期数据转换 优化
      dayChecked: {}, // 当前选择的天
      weekList: ["一", "二", "三", "四", "五", "六", "日"], // 周
      isOpen: ['isopen']
    };
  },
  created() {
    this.initDataFun(); // 初始化数据
  },
  computed: {
    // 显示当前时间
    currentDateStr() {
      let { year, month } = this.showYearMonth;
      return `${year}年${this.pad(month + 1)}月`;
    },
  },
  methods: {
    //#region 计算日历数据
    // 初始化数据
    initDataFun() {
      // 初始化当前时间
      this.setCurrentYearMonth(); // 设置日历显示的日期（年-月）
      this.createCalendar(); // 创建当前月对应日历的日期数据
      this.getCurrentDay(); // 获取今天
    },
    // 设置日历显示的日期（年-月）
    setCurrentYearMonth(d = new Date()) {
      let year = d.getFullYear();
      let month = d.getMonth();
      let date = d.getDate();
      this.showYearMonth = {
        year,
        month,
        date,
      };
    },
    getCurrentDay(d = new Date()) {
      let year = d.getFullYear();
      let month = d.getMonth();
      let date = d.getDate();
      this.dayChecked = {
        year,
        month,
        date,
        value: this.stringify(year, month, date),
        disable: false
      };
    },
    // 创建当前月对应日历的日期数据
    createCalendar() {
      // 一天有多少毫秒
      const oneDayMS = 24 * 60 * 60 * 1000;
      let list = [];
      let { year, month } = this.showYearMonth;
      
      // #region
      // ---------------仅仅只算某月的天---------------
      //   // 当前月，第一天和最后一天的毫秒数
      //   let begin = new Date(year, month, 1).getTime();
      //   let end = new Date(year, month + 1, 0).getTime();

      // ---------------计算某月前后需要填补的天---------------
      // 当前月份第一天是星期几, 0-6
      let firstDay = this.getFirstDayByMonths(year, month);
      // 填充多少天，因为我将星期日放到最后了，所以需要另外调整下
      let prefixDaysLen = firstDay === 0 ? 6 : firstDay - 1;
      // 向前移动之后的毫秒数
      let begin = new Date(year, month, 1).getTime() - oneDayMS * prefixDaysLen;
      // 当前月份最后一天是星期几, 0-6
      let lastDay = this.getLastDayByMonth(year, month);
      // 填充多少天，因为我将星期日放到最后了，所以需要另外调整下
      let suffixDaysLen = lastDay === 0 ? 0 : 7 - lastDay;
      // 向后移动之后的毫秒数
      let end = new Date(year, month + 1, 0).getTime() + oneDayMS * suffixDaysLen;
      // // 计算左侧时间段的循环数
      // let rowNum = Math.ceil((end - begin) / oneDayMS / 7);
      // let newPeriod = [];
      // for (let i = 0; i < rowNum; i++) {
      //   newPeriod.push({});
      // }
      // #endregion

      // 填充天
      while (begin <= end) {
        // 享元模式，避免重复 new Date
        this.shareDate.setTime(begin);
        let year = this.shareDate.getFullYear();
        let curMonth = this.shareDate.getMonth();
        let date = this.shareDate.getDate();
        list.push({
          year: year,
          month: curMonth + 1, // 月是从0开始的
          date: date,
          value: this.stringify(year, curMonth, date),
          disable: curMonth !== month,
        });
        begin += oneDayMS;
      }

      this.calendarList = list;
    },
    // 格式化时间
    stringify(year, month, date) {
      let str = [year, this.pad(month + 1), this.pad(date)].join("-");
      return str;
    },
    // 对小于 10 的数字，前面补 0
    pad(str) {
      return str < 10 ? `0${str}` : str;
    },
    // 点击上一月
    prevMonth() {
      this.showYearMonth.month--;
      this.recalculateYearMonth(); // 因为 month的变化 会超出 0-11 的范围， 所以需要重新计算
      this.createCalendar(); // 创建当前月对应日历的日期数据
    },
    // 点击下一月
    nextMonth() {
      this.showYearMonth.month++;
      this.recalculateYearMonth(); // 因为 month的变化 会超出 0-11 的范围， 所以需要重新计算
      this.createCalendar(); // 创建当前月对应日历的日期数据
    },
    // 重算：显示的某年某月
    recalculateYearMonth() {
      let { year, month, date } = this.showYearMonth;

      let maxDate = this.getDaysByMonth(year, month);
      // 预防其他月跳转到2月，2月最多只有29天，没有30-31
      date = Math.min(maxDate, date);

      let instance = new Date(year, month, date);
      this.setCurrentYearMonth(instance);
    },
    // 判断当前月有多少天
    getDaysByMonth(year, month) {
      return new Date(year, month + 1, 0).getDate();
    },
    // 当前月的第一天是星期几
    getFirstDayByMonths(year, month) {
      return new Date(year, month, 1).getDay();
    },
    // 当前月的最后一天是星期几
    getLastDayByMonth(year, month) {
      return new Date(year, month + 1, 0).getDay();
    },
    // #endregion 计算日历数据

    // 操作：点击了某天
    handleClickDay(item) {
      if (!item || item.disable) return;
      console.log(item);
      this.dayChecked = item;
    },
    SwitchFunc(){
      this.$emit('isopen', !this.isOpen)
    },
  },
  components: {
    'draggable-resizable': VueDraggableResizable
  }
};
</script>

<style lang="less" scoped>
@calendarWidth: 672px; // 90 * 7  + 2 * 3 * 7
.calendar-box {
  color: #4B4737;
  display: flex;
  justify-content: center;
  align-items: center;
  .controls{
    display: flex;
    flex-direction: row;
    align-items: center;
    position: absolute;
    top: 6px;
    right: -336px;
    .prev,
    .next {
      display: flex;
      flex-direction: row;
      i{
        margin: auto;
      }
      width: 16px;
      height: 16px;
      // font-size: 30px;
      border-radius: 50%;
    }
    .prev:hover,.next:hover{
      box-shadow: 2px 2px 14px 1px rgba(0,0,0,0.2);
      backdrop-filter: blur(10px);
      color: rgb(111, 233, 255);
      cursor: pointer;
    }
    .cross{
      width: 16px;
      height: 16px;
      border-radius: 50%;
      display: flex;
      flex-direction: row;
      align-items: center;
      // margin-bottom: 3px;
      // position: absolute;
      // right: 0;
      // top: 0;
    }
    .cross:hover{
      box-shadow: 2px 2px 14px 1px rgba(0,0,0,0.2);
      backdrop-filter: blur(10px);
      color: red;
      cursor: pointer;
    }
  }
  .calendar-wrapper {
    .calendar-toolbar {
      width: @calendarWidth;
      display: flex;
      justify-content: space-between;
      align-items: center;
      .current {
        cursor: pointer;
        &:hover {
          color: #438bef;
        }
      }
    }
    .calendar-week {
      width: @calendarWidth;
      display: flex;
      flex-wrap: wrap;
      .week-item {
        width: 90px;
        height: 50px;
        background-color: #fef3cabe;
        border-radius: 10px;
        margin: 3px;
      }
    }
    .calendar-inner {
      width: @calendarWidth;
      display: flex;
      flex-wrap: wrap;
      .calendar-item {
        width: 90px;
        height: 60px;
        background-color: #fef3cabe;
        border-radius: 10px;
        margin: 3px;
      }
      .calendar-item-hover {
        cursor: pointer;
        &:hover {
          color: #fff;
          background-color: #438bef;
        }
      }
      .calendar-item-disabled {
        color: #acacac;
        cursor: not-allowed;
      }
      .calendar-item-checked {
        color: #fff;
        background-color: #438bef;
      }
    }
    .calendarBorder {
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }
}
*{
  user-select: none;
}
</style>
