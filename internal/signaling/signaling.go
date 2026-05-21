package signaling

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type signalingMsg struct {
	Type string `json:"type"`
	SDP  string `json:"sdp,omitempty"`
	// Input forwarding fields (Phase 1: via WebSocket)
	T string  `json:"t,omitempty"` // "m"=mouse, "k"=keyboard
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	C string  `json:"c,omitempty"` // key code
	D bool    `json:"d,omitempty"` // key down
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("[signaling] WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()
	log.Println("[signaling] New client connected:", r.RemoteAddr)

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{URLs: []string{"stun:stun.l.google.com:19302"}},
		},
	}

	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Fatal("[signaling] PeerConnection error:", err)
	}
	defer pc.Close()

	videoTrack, err := webrtc.NewTrackLocalStaticSample(
		webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8},
		"video", "gostream",
	)
	if err != nil {
		log.Fatal("[signaling] Track error:", err)
	}
	pc.AddTrack(videoTrack)

	pc.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		log.Printf("[signaling] ICE state: %s", state)
	})

	// TODO Phase 3: collegare capture.Capturer + encoder qui
	// go encoder.StreamFrames(capturer.Start(), videoTrack)

	for {
		var msg signalingMsg
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("[signaling] Read error:", err)
			return
		}

		switch msg.Type {
		case "offer":
			offer := webrtc.SessionDescription{Type: webrtc.SDPTypeOffer, SDP: msg.SDP}
			if err := pc.SetRemoteDescription(offer); err != nil {
				log.Println("[signaling] SetRemoteDescription error:", err)
				continue
			}
			answer, err := pc.CreateAnswer(nil)
			if err != nil {
				log.Println("[signaling] CreateAnswer error:", err)
				continue
			}
			if err := pc.SetLocalDescription(answer); err != nil {
				log.Println("[signaling] SetLocalDescription error:", err)
				continue
			}
			conn.WriteJSON(signalingMsg{Type: "answer", SDP: pc.LocalDescription().SDP})
			log.Println("[signaling] Offer/Answer exchange complete")
		}

		// Input injection (Phase 1: WebSocket) — da migrare a DataChannel in Phase 4
		switch msg.T {
		case "m":
			// TODO: inputInjector.InjectMouse(input.MouseEvent{X: msg.X, Y: msg.Y})
			_ = msg.X
			_ = msg.Y
		case "k":
			// TODO: inputInjector.InjectKey(input.KeyEvent{Code: msg.C, Pressed: msg.D})
			_ = msg.C
			_ = msg.D
		}
	}
}

// encodeJSON is a helper used in tests
func encodeJSON(v any) []byte {
	b, _ := json.Marshal(v)
	return b
}
