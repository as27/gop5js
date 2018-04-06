var socket = new WebSocket("ws://" + ServerName + ServerPort + "/ws");

var p5Data = {
  mouseX: 0,
  mouseY: 0
};

function setup() {
  createCanvas(CanvasWidth, CanvasHeight);
  noLoop();
}

var sketch_draw = "";

function draw() {
  clear();
  eval(sketch_draw)
  socket.send(JSON.stringify(getParams()));
  console.log(mouseX, mouseY);
}

socket.onmessage = function (event) {
  newData = JSON.parse(event.data);
  sketch_draw = newData.sketch_draw;
  console.log(sketch_draw)
  draw();
}

function getParams() {
  return {
    mouse_x: mouseX,
    mouse_y: mouseY,
    p_mouse_x: pmouseX,
    p_mouse_y: pmouseY,
    win_mouse_x: winMouseX,
    win_mouse_y: winMouseY,
    p_win_mouse_x: winMouseX,
    p_win_mouse_y: winMouseY,
    mouse_button: (mouseButton === 0) ? "" : mouseButton,
    mouse_is_pressed: mouseIsPressed
  }
}