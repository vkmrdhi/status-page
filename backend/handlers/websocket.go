package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket for status updates
func StatusUpdates(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		// Simulating real-time status updates
		// In a real scenario, this could be triggered by status changes in the system
		statusUpdate := map[string]string{
			"service": "API",
			"status":  "Operational",
			"message": "Everything is running smoothly!",
		}
		err := conn.WriteJSON(statusUpdate)
		if err != nil {
			log.Println("Failed to send WebSocket message:", err)
			return
		}
		time.Sleep(10 * time.Second) // Simulating a delay for updates
	}
}
