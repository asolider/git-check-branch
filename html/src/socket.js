const socket = new WebSocket("wss://echo.websocket.org")

let count = 0

setInterval(() => {
  socket.send(`Message Number ${++count}`)
}, 5000)

export default socket