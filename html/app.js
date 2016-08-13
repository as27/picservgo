var serverURL = "http://raspi.fritz.box:4558/folder/"

new Vue({
  el: '#app',
  data: {
    path: "",
    content: null,
    breadcrumps: []
  },
  created: function () {
    this.fetchData()
  },
  methods: {
    openFolder: function(f){
      this.breadcrumps.push(this.path)
      if (this.path == ""){
        this.path = this.path + f
      }else{
        this.path = this.path + "/" + f
      }
      this.fetchData()
    },
    levelUp: function(){
      this.path = this.breadcrumps.pop()
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