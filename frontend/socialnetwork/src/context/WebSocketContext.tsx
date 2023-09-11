// WebSocketContext.js
import React, { createContext, ReactNode, useContext } from 'react';
import { WebSocketWriteMessage, WebSocketReadMessage, useWebSocket } from '../Socket';

export type WebSocketContextType = {
    message: WebSocketWriteMessage | null;
    sendMessage: ((messageToSend: WebSocketReadMessage) => void);
};

const WebSocketContext = createContext<WebSocketContextType>({
    message: null,
    sendMessage: () => { },
});

export function useWebSocketContext() {
    return useContext(WebSocketContext);
}

type WebSocketProviderProps = {
    children: ReactNode;
    ws: ReturnType<typeof useWebSocket>; // Import the useWebSocket hook here
};

export function WebSocketProvider({ children, ws }: WebSocketProviderProps) {
    const { message, sendMessage } = ws;

    return (
        <WebSocketContext.Provider value={{ message, sendMessage }}>
            {children}
        </WebSocketContext.Provider>
    );
}
