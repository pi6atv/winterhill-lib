<template>
  <line-chart v-bind:options="options" v-bind:chartData="datasets" style="height: 300px"></line-chart>
</template>
<script>
import LineChartComponent from "@/components/LineChartComponent";

// let index = 0

export default {
  components: {LineChart: LineChartComponent},
  props: ['signal', 'calls'],
  data () {
    return {
      options: {
        elements: {
          point: {
            radius: 3
          }
        },
        scales: {
          yAxes: [{
            ticks: {
              min: 0,
              max: 100,
              beginAtZero: true
            }
          }],
          xAxes: [
              {
                id: 'time',
                type: 'time',
                position: 'bottom',
                distribution: 'linear',
                ticks: {
                  minRotation: 45,
                },
                time: {
                  stepSize: 30,
                  displayFormats: {
                    second: "HH:mm:ss"
                  },
                },
              },
              // {
              //   id: 'calls',
              //   type: 'category',
              //   position: 'top',
              //   labels: this.signal?.map(item => { // going through time
              //         const now = new Date(item.Time)
              //         try {
              //           if (new Date(this.calls[index].start) <= now && !(new Date(this.calls[index].end) < now)) { // we passed a call event
              //             index += 1
              //             return this.calls[index-1].value
              //           }
              //         } catch (e) {
              //           return ""
              //         }
              //         return ""
              //       }),
              //   ticks: {
              //     autoSkip: false,
              //   },
              // },
          ],
          title: {
            display: false,
          }
        },
        legend: {
          display: false,
        },
        responsive: true,
        maintainAspectRatio: false,
        // aspectRatio: 1,
      }
    }
  },

  computed: {
    datasets: function () {
      return {
        datasets: [
            {
              backgroundColor: '#2c4df5',
              borderColor: '#2c4df5',
              fill: false,
              xAxisID: 'time',
              data: this.signal.filter(item => {return item !== null}).map(item => {
                return {y: item.value, x: item.time}
              })
            },
        ],
      }
    }
  },
}
</script>