import { useState, useRef, useEffect } from "react";

export interface WebSocketMessage {
    type: string
    data: any
}

export const useWebSocket = (url: string, options: { headers?: Record<string, string> } = {}) => {
    const [message, setMessage] = useState<WebSocketMessage | null>(null);
    const ws = useRef<WebSocket | null>(null);

    useEffect(() => {
        const socket = new WebSocket(url);

        socket.onopen = () => {
            console.log("WebSocket opened");

            if (options.headers) {
                for (const [header, value] of Object.entries(options.headers)) {
                    socket.send(JSON.stringify({ type: 'header', header, value }));
                }
            }
        };

        socket.onclose = () => {
            console.log("WebSocket closed");
        };

        socket.onmessage = (event) => {
            let message = JSON.parse(event.data);
            console.log(message);
            setMessage(message);
        };

        ws.current = socket;

        return () => {
            if (ws.current) {
                ws.current.close();
            }
        };
    }, [url]);

    const sendMessage = (messageToSend: string) => {
        if (ws.current) {
            ws.current.send(messageToSend);
        }
    };

    return { message, sendMessage };
};
