class Socket {
    constructor() {
        this.socket = new WebSocket(process.env.SOCKET_URI);
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
                this.socket = new WebSocket(process.env.SOCKET_URI)
                this.openSocket();
            }, 1000);
        }
    }
}

export default Socket;
