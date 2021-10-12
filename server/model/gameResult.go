package model

type GameResult struct {
	P1Move string `json:"p1Move"`
	P2Move string `json:"p2Move"`
	Winner string `json:"winner"`
}
