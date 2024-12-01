import React, { createContext, useContext, useState, useEffect } from 'react';

const WebSocketContext = createContext<WebSocket | null>(null);

export const WebSocketProvider: React.FC<{ url: string; children: React.ReactNode }> = ({ url, children }) => {
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    const ws = new WebSocket(url);
    setSocket(ws);

    return () => {
      ws.close();
    };
  }, [url]);

  return <WebSocketContext.Provider value={socket}>{children}</WebSocketContext.Provider>;
};

export const useWebSocket = () => {
  return useContext(WebSocketContext);
};
