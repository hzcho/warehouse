package topichandler

import (
	topichandler "audit/internal/domain/topic_handler"
	"audit/internal/usecase"
)

type TopicHandlers struct {
	topichandler.Operation
}

func NewTopicHandlers(usecases *usecase.Usecases) *TopicHandlers {
	return &TopicHandlers{
		Operation: NewOperation(usecases.Operation),
	}
}
