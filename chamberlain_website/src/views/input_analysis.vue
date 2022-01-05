<template>
  <div id="YearStatistics" class="panel-chart">
    <span>
      年度收入统计:<br>
      <Dropdown v-model="currentYearObject" :options="yearOptions" optionLabel="name" :placeholder="currentYearObject.name" @change="changeYear"/>
    </span>
    <Chart type="pie" :data="yearTypeStatisticsData" class="inner-panel-chart" :height="394" :width="394"/>&nbsp;&nbsp;
    <Chart type="bar" :data="yearMonthStatisticsData" class="inner-panel-chart" :height="394" :width="1200"/>
  </div>
  <div id="YearsTrending" class="panel-chart">
    <span>年度收入趋势: </span>
    <Chart type="line" :data="yearsTrendingData" :height="398" :width="1630"/>
  </div>
</template>

<script>
import InputService from '../api/input_mgmt.ts';

export default {
  name: "input_analysis",
  data() {
    return {
      currentYear: new Date().getFullYear(),
      currentYearObject: null,
      yearOptions: [],
      yearsTrendingData: {},
      yearMonthStatisticsData: {},
      yearTypeStatisticsData: {}
    }
  },
  inputService: null,
  created() {
    this.inputService = new InputService();
    this.getYearsTrending()
    this.getYearMonthStatistics()
    this.getYearTypeStatistics()
    this.currentYearObject = {"name": this.currentYear + "年", "code": String(this.currentYear)}
  },
  methods: {
    getYearsTrending() {
      this.inputService.getYearsTrendingData().then(queryData => {
        let labels = []
        let allInputData = []
        let actualData = []
        let taxData = []
        let tmpYearOptions = []
        queryData.forEach(function (aYearData) {
          let yearDisplay = aYearData.Year + "年"
          labels.push(yearDisplay)
          allInputData.push(aYearData.AllInput)
          actualData.push(aYearData.Actual)
          taxData.push(aYearData.Tax)
          tmpYearOptions.push({"name": yearDisplay, "code": String(aYearData.Year)})
        })

        let datasets = [{
          label: '税前总收入',
          fill: false,
          borderColor: '#0000ff',
          backgroundColor: '#0000ff',
          data: allInputData
        }, {
          label: '税后总收入',
          fill: false,
          borderColor: '#ff0000',
          backgroundColor: '#ff0000',
          data: actualData
        }, {
          label: '纳税总额',
          fill: false,
          borderColor: '#00ffff',
          backgroundColor: '#00ffff',
          data: taxData
        }]
        this.yearOptions = tmpYearOptions
        this.yearOptions.sort(function(a, b) {
          return b.code - a.code
        })
        this.yearsTrendingData = {"labels": labels, "datasets": datasets}
      })
    },
    getYearTypeStatistics() {
      this.inputService.getYearTypeStatisticsData(this.currentYear).then(queryData => {
        let labels = []
        let displayData = []
        queryData.forEach(function (aTypeData) {
          labels.push(aTypeData.Type)
          displayData.push(aTypeData.AllInput)
        })

        this.yearTypeStatisticsData = {
          "labels": labels, "datasets": [{
            data: displayData,
            backgroundColor: [
              "#42A5F5",
              "#66BB6A",
              "#FFA726",
              "#EE82EE"
            ],
            hoverBackgroundColor: [
              "#64B5F6",
              "#81C784",
              "#FFB74D",
              "#DDA0DD"
            ]
          }
          ]
        }
      })
    },
    getYearMonthStatistics() {
      this.inputService.getYearMonthStatisticsData(this.currentYear).then(queryData => {
        let labels = []
        let allInputData = []
        let actualData = []
        let taxData = []
        queryData.forEach(function (aMonthData) {
          labels.push(aMonthData.Month + "月")
          allInputData.push(aMonthData.AllInput)
          actualData.push(aMonthData.Actual)
          taxData.push(aMonthData.Tax)
        })

        let datasets = [{
          label: '税前总收入',
          fill: false,
          backgroundColor: '#0000ff',
          data: allInputData
        }, {
          label: '税后总收入',
          fill: false,
          backgroundColor: '#ff0000',
          data: actualData
        }, {
          label: '纳税总额',
          fill: false,
          backgroundColor: '#00ffff',
          data: taxData
        }]
        this.yearMonthStatisticsData = {"labels": labels, "datasets": datasets}
      })
    },
    changeYear() {
      this.currentYear = this.currentYearObject.code
      this.getYearMonthStatistics()
      this.getYearTypeStatistics()
    }
  }
}
</script>

<style scoped>
.panel-chart {
  display: flex;
  height: 50%;
  width: 92%;
  text-align: center;
  margin: 0 auto;
  overflow: auto
}

.inner-panel-chart {
  display: inline-block;
}
</style>