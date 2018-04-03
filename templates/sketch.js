var socket = new WebSocket("ws://localhost:2700/ws");

function setup() {
   createCanvas(500,500);
   background(127);
   noLoop();

  }
  
  function draw() {
    ellipse(10,10,50,50);
  }
 
socket.onmessage = function(event){
    console.log(event.data);
}