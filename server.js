const WebSocket = require('ws');

const server = new WebSocket.Server({ port: 3000 });

server.on('connection', (socket) => {
  console.log('Client connected');

  socket.on('message', (message) => {
    console.log(`Received: ${message}`);

    if (message == 'hello') {
      // 如果客户端发送 'hello'，则回复 'yes'
      socket.send('yes');
    }
  });

  socket.on('close', () => {
    console.log('Client disconnected');
  });
});
