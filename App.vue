<template>
  <b-container>
    <b-alert variant="danger" :show="alertVisible" dismissible="">{{alertText}}</b-alert>
    <b-card title="Входные данные" class="mt-2">
      <b-row>
        <b-col cols="8">
          <b-form-input v-model="inputData" @keyup="drawChart()" placeholder="Числа через пробел (0.004 43.23 0.34)"></b-form-input>
        </b-col>
        <b-col>
          <b-button-group>
          <b-form-file v-model="inputFile" @change="loadTextFromFile" placeholder="Выберите файл с данными..."></b-form-file>
          <b-button @click="clearFile()" variant="danger">X</b-button>
          </b-button-group>
        </b-col>
      </b-row>
    </b-card>

    <b-card title="Выбор метода" class="mt-2">
      <b-row>
        <b-col cols="4">
          <b-form-checkbox @change="doForecastExp">Экспоненциальное сглаживание</b-form-checkbox>
        </b-col>
        <b-col cols="3">
           <b-form-input v-model="inputExpAlpha" @change="doForecastExp" trim></b-form-input>
           <b-form-text>Коэффициент альфа</b-form-text>
        </b-col>
        <b-col cols="3">
           <p>Результат: {{resultExp}}</p>
        </b-col>
        <b-col cols="2">
           <b-form-input v-model="w1"></b-form-input>
           <b-form-text>Коэффициент доверия</b-form-text>
        </b-col>
      </b-row>

      <b-row>
        <b-col cols="4">
          <b-form-checkbox @change="doForecastSlice">Метод скользящего среднего</b-form-checkbox>
        </b-col>
        <b-col cols="3">
           <b-form-input v-model="inputSliceSmooth" @change="doForecastSlice" trim></b-form-input>
            <b-form-text>Коэффициент сглаживания</b-form-text>
        </b-col>
        <b-col cols="3">
            <p>Результат: {{resultSlice}}</p>
        </b-col>
        <b-col cols="2">
           <b-form-input v-model="w2"></b-form-input>
           <b-form-text>Коэффициент доверия</b-form-text>
        </b-col>
      </b-row>

      <b-row class="my-2">
        <b-col cols="4">
          <b-form-checkbox @change="doForecastRegression">Регрессионная модель</b-form-checkbox>
        </b-col>
        <b-col cols="3">
        </b-col>
        <b-col cols="3">
           <p>Результат: {{resultRegression}}</p>
        </b-col>
        <b-col cols="2">
           <b-form-input v-model="w3"></b-form-input>
           <b-form-text>Коэффициент доверия</b-form-text>
        </b-col>
      </b-row>
      <b-row>
        <b-col cols="10">
        </b-col>
        <b-col cols="2">
           <p>Результат P: {{resultP}}</p>
        </b-col>
      </b-row>
    </b-card>

    <b-card title="График результатов" class="mt-2">
      <line-chart :chartData="datacollection" :options="options"></line-chart>
    </b-card>

  </b-container>
</template>

<script>
import LineChart from './LineChart.js'

export default {
  name: 'app',

  components: {
    LineChart
  },

  data () {
    return {
      alertVisible: false,
      alertText: '',

      inputFile: null,
      inputData: '',
      inputFileData: '',
      graphXCount: 0,

      inputExpAlpha: 1,
      resultExp: 0,

      inputSliceSmooth: 0.5,
      resultSlice: 0,

      resultRegression: 0,

      w1: 0.4,
      w2: 0.1,
      w3: 0.5,

      datacollection: {},
      options: {
        responsive: true,
        aspectRatio: 5,
        maintainAspectRatio: false,
        scales: {
          xAxes: [{
            type: 'linear',
            display: false
          }]
        }
      }
    }
  },

  mounted () {
    this.drawChart()
  },

  computed: {
    rA () {
      return this.w1 * this.resultExp
    },
    rB () {
      return this.w2 * this.resultSlice
    },
    rC () {
      return this.w3 * this.resultRegression
    },
    resultP () {
      return this.rA + this.rB + this.rC
    }
  },

  methods: {
    showError (text) {
      this.alertText = text
      this.alertVisible = true
    },

    drawChart () {
      var datasets = []

      var inputDataset = []
      this.graphXCount = 0
      this.getInputData().split(' ').forEach(element => {
        if (element) {
          inputDataset.push({ x: this.graphXCount++, y: parseFloat(element) })
        }
      })

      datasets = [{
        label: 'Входные данные',
        backgroundColor: 'green',
        borderColor: 'green',
        fill: false,
        data: inputDataset
      },
      {
        label: 'Эксп. сглаживание',
        backgroundColor: 'red',
        borderColor: 'red',
        fill: false,
        data: []
      },
      {
        label: 'Метод скольз. среднего',
        backgroundColor: 'brown',
        borderColor: 'brown',
        fill: false,
        data: []
      },
      {
        label: 'Регрессионная модель',
        backgroundColor: 'blue',
        borderColor: 'blue',
        fill: false,
        data: []
      }]

      this.datacollection = {
        datasets: datasets
      }
    },

    drawResult (i, result) {
      var dc = this.datacollection.datasets
      dc[i].data = [{ x: this.graphXCount, y: result }]
      this.datacollection = {
        datasets: dc
      }
    },

    drawArrayOfResults (i, result) {
      var dc = this.datacollection.datasets
      dc[i].data = result
      this.datacollection = {
        datasets: dc
      }
    },

    hideResultFromGraph (i) {
      var dc = this.datacollection.datasets
      dc[i].data = []
      this.datacollection = {
        datasets: dc
      }
    },

    loadTextFromFile (ev) {
      const file = ev.target.files[0]
      const reader = new FileReader()

      reader.onload = e => {
        this.inputFileData = e.target.result
        this.drawChart()
      }
      reader.readAsText(file)
    },

    clearFile () {
      this.inputFile = null
      this.inputFileData = ''
      this.drawChart()
    },

    getInputData () {
      if (this.inputFileData) {
        return this.inputFileData
      }
      return this.inputData
    },

    doForecastExp (checked) {
      var self = this
      if (checked) {
        this.axios.post('forecast/exp/', { 'numbers': this.getInputData(), 'coef': parseFloat(this.inputExpAlpha) })
          .then((r) => {
            let graphData = []
            r.data.split('\n').forEach((element, index, arr) => {
              if (element) {
                if (index === 0) {
                  this.resultExp = element
                } else {
                  let point = element.split(' ')
                  graphData.push({ x: parseFloat(point[0]), y: parseFloat(point[1]) })
                }
              }
            })
            self.drawArrayOfResults(1, graphData)
          })
          .catch((err) => {
            self.showError(err.response.data)
          })
      } else {
        this.hideResultFromGraph(1)
      }
    },

    doForecastSlice (checked) {
      var self = this
      if (checked) {
        this.axios.post('forecast/slice/', { 'numbers': this.getInputData(), 'coef': parseFloat(this.inputExpAlpha) })
          .then((r) => {
            this.resultSlice = parseFloat(r.data)
            this.drawResult(2, this.resultSlice)
          })
          .catch((err) => {
            self.showError(err.response.data)
          })
      } else {
        this.hideResultFromGraph(2)
      }
    },

    doForecastRegression (checked) {
      var self = this
      if (checked) {
        this.axios.post('forecast/regression/', { 'numbers': this.getInputData(), 'coef': 0 })
          .then((r) => {
            let graphData = []
            r.data.split('\n').forEach((element, index, arr) => {
              if (element) {
                if (index === 0) {
                  this.resultRegression = element
                } else {
                  let point = element.split(' ')
                  graphData.push({ x: parseFloat(point[0]), y: parseFloat(point[1]) })
                }
              }
            })
            self.drawArrayOfResults(3, graphData)
          })
          .catch((err) => {
            self.showError(err.response.data)
          })
      } else {
        this.hideResultFromGraph(3)
      }
    }
  }

}
</script>

<style>

</style>
