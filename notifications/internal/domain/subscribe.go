package domain

import "context"

func (s *Service) Subscribe(ctx context.Context) {
	s.LogsReciver.Subscribe(ctx)
}
