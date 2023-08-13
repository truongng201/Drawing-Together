class Socket {
    constructor(path) {
        this.URI = process.env.SOCKET_URI + '/' + path;
        this.socket = new WebSocket(this.URI);
    }


    openSocket() {
        this.socket.onopen = () => {
            console.log("Socket opened")
        }
    }

    sendMessage(message) {
        this.socket.send(message);
    }

    receiveMessage(handleData) {
        this.socket.onmessage = (event) => {
            handleData(event.data)
        }
    }

    closeSocket() {
        this.socket.onclose = () => {
            console.log('Socket closed');
        }
    }

    error() {
        this.socket.onerror = (event) => {
            console.log(event);
        }
    }

    reconnectSocket() {
        this.socket.onclose = () => {
            console.log('Socket closed. Reconnecting...');
            setTimeout(() => {
                this.socket = new WebSocket(this.URI)
                this.openSocket();
            }, 1000);
        }
    }
}

export default Socket;
