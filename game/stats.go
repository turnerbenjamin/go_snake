package game

import "time"

type stats struct {
	start         time.Time
	frames        int
	fps           float64
	fpsSampleSize int
}

func newStats(fpsSampleSize int, startingFPS int) *stats {
	return &stats{
		start:         time.Now(),
		fpsSampleSize: fpsSampleSize,
		fps:           float64(startingFPS),
	}
}

func (s *stats) update() {
	s.frames++
	if s.frames == s.fpsSampleSize {
		s.fps = float64(s.frames) / time.Since(s.start).Seconds()
		s.frames = 0
		s.start = time.Now()
	}
}

func (s *stats) startTracking(){
		s.frames = 0
		s.start = time.Now()
}