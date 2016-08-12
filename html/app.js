var serverURL = "http://raspi.fritz.box:4558/folder/2016"

new Vue({
  el: '#app',
  data: {
    path: "",
    content: null
  },
  created: function () {
    this.fetchData()
  },
  methods: {
    openFolder: function(f){
      this.path = this.path + "/" + f
      this.fetchData()
    },
    fetchData: function () {
      var xhr = new XMLHttpRequest()
      var self = this
      xhr.open('GET', serverURL+self.path)
      xhr.onload = function () {
        self.content = JSON.parse(xhr.responseText)
      }
      xhr.send()
    }
  }
})