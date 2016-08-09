var serverURL = "http://raspi.fritz.box:4558/folder/2016"

new Vue({
  el: '#app',
  data: {
    path: "/2016-08",
    folder: null
  },
  created: function () {
    this.fetchData()
  },
  methods: {
    fetchData: function () {
      var xhr = new XMLHttpRequest()
      var self = this
      xhr.open('GET', serverURL+self.path)
      xhr.onload = function () {
        self.folder = JSON.parse(xhr.responseText)
      }
      xhr.send()
    }
  }
})