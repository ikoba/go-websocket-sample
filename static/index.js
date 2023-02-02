let socket = null;

window.addEventListener("load", function(event) {
  socket = new WebSocket("ws://localhost:3000/ws");
  socket.onopen = () => {
    appendMessage("system: socket opened");
  };
  socket.onmessage = (event) => {
    appendMessage(`server: ${event.data}`);
  };
});

function send() {
  const msg = document.querySelector("#message").value;
  if (msg === "")
    return;
  socket.send(msg);
  appendMessage(`client: "${msg}"`);
}

function appendMessage(message) {
  const p = document.createElement("p");
  p.textContent = message;
  document.querySelector("#messages").appendChild(p);
}