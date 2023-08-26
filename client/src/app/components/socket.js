class Socket {
  constructor(path) {
    this.URI = process.env.NEXT_PUBLIC_SOCKET_URI + path + "/";
    this.socket = null;
    this.isOpen = false;
  }

  open() {
    this.socket = new WebSocket(this.URI);
    this.socket.onopen = () => {
      this.isOpen = true;
    };
    this.socket.onclose = () => {
      this.isOpen = false;
    };
  }

  reconnect(timeout = 5000, maxAttempts = 10) {
    var attempts = 1;
    while (this.socket.readyState !== WebSocket.OPEN) {
      setTimeout(() => {
        console.log("Attempting to reconnect", attempts);
        this.socket = new WebSocket(this.URI);
      }, timeout * attempts);
      attempts++;
      if (this.socket && this.socket.readyState === WebSocket.OPEN) {
        break;
      }
      if (maxAttempts && attempts > maxAttempts) {
        console.log("Max attempts reached");
        break;
      }
    }
  }

  receive(callback) {
    this.socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      callback(data);
    };
  }

  send(message) {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify(message));
    }
  }

  close() {
    this.socket.close();
  }
}

export default Socket;
