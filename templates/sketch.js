var socket = new WebSocket("ws://localhost:2700/ws");

function setup() {
   createCanvas(500,500);
   background(127);
   noLoop();

  }
  
  var sketch_draw = "";

  function draw() {
    eval(sketch_draw)

    socket.send("ready");
  }
 
socket.onmessage = function(event){
    newData = JSON.parse(event.data);
    sketch_draw = newData.sketch_draw;
    console.log(sketch_draw)
    draw();
}