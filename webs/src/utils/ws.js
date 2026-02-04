import { reactive } from "vue";

let instance = null;

export function createWs(userId,token) {
    if (instance) return instance;
    
    const ws = new WebSocket(`ws://localhost:8080/ws?userId=${userId}&token=${token}`);
    const state = reactive({
        isConnected: false,
        messages: [],
    });

    ws.onopen = () => {
        state.isConnected = true;
        console.log("WebSocket has been connected");
    }

    ws.onmessage = (event) => {
        const message = JSON.parse(event.data);
        state.messages.push(message);
        console.log("Received message:", message);
    }

    ws.onclose = () => {
        state.isConnected = false;
        console.log("WebSocket has been disconnected");
    }

    ws.onerror = (error) => {
        console.error("WebSocket has an error:", error);
        ws.close();
    }

    const sendMessage = (toUserId, message) => {
        if (state.isConnected) return
        ws.send(JSON.stringify({
            from: userId,
            to: toUserId,
            message,
            timestamp: Date.now()
        }));
    }
       
    instance = {
        ws,
        state,
        sendMessage
    };

    return instance;
}

export function getWsInstance() {
    return instance;
}