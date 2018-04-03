var socket = new WebSocket("ws://localhost:2700/ws");

function setup() {
   createCanvas(500,500);
   background(127);
   noLoop();

  }
  
  var data = {
    Ellipses: []
  };

  function draw() {
    clear()
    data.Ellipses.forEach(function(obj){
      console.log(obj);
      ellipse(obj.x,obj.y,obj.w,obj.h);
    });

    socket.send("ready");
  }
 
socket.onmessage = function(event){
    newData = JSON.parse(event.data);
    data = newData.shapes
    console.log(data)
    draw();
}