var golang = require('./main.js');
var global = Function('return this')();

window.golang = golang;
console.log("window.golang", window.golang);

global.document.addEventListener("DOMContentLoaded", function(){
  for( var i = 0; i < 10; i ++) {
      var li = golang.list.New(i, "#" + i);
      var node = li.Node();
      window.document.body.appendChild(node);
  }
});
