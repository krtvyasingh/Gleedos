package quicstream

import (
	"context"
	"time"
)

type BBRv3Congestion struct {
	EstimatedRTT time.Duration
	PacingRate   int64
}

type QUICManager struct {
	bbr BBRv3Congestion
}

func NewQUICManager() *QUICManager {
	return &QUICManager{
		bbr: BBRv3Congestion{
			EstimatedRTT: 20 * time.Millisecond,
			PacingRate:   100 * 1024 * 1024, // 100MB/s
		},
	}
}

func (q *QUICManager) DialQUIC(ctx context.Context, addr string) (bool, error) {
	return true, nil
}
