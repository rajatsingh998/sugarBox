package Models

import (
	"sync"
	"time"
)

type UserEventCounter struct {
	likeCounts    int64
	commentCounts int64
	postCounts    int64
	date          time.Time
	rwMutex       sync.RWMutex
}

func (uc *UserEventCounter) GetLikeCount() int64 {
	defer uc.rwMutex.RUnlock()
	uc.rwMutex.RLock()
	return uc.likeCounts
}

func (uc *UserEventCounter) GetCommentCount() int64 {
	defer uc.rwMutex.RUnlock()
	uc.rwMutex.RLock()
	return uc.commentCounts
}

func (uc *UserEventCounter) GetPostCount() int64 {
	defer uc.rwMutex.RUnlock()
	uc.rwMutex.RLock()
	return uc.postCounts
}

func (uc *UserEventCounter) GetDate() time.Time {
	defer uc.rwMutex.RUnlock()
	uc.rwMutex.RLock()
	return uc.date
}

func (uc *UserEventCounter) SetDate(date time.Time) {
	defer uc.rwMutex.RUnlock()
	uc.rwMutex.RLock()
	uc.date = date
}

func (uc *UserEventCounter) IncreaseLikeCountByOne() {
	defer uc.rwMutex.RUnlock()
	uc.rwMutex.RLock()
	oldLikeCount := uc.likeCounts
	newLikeCount := oldLikeCount + 1
	uc.likeCounts = newLikeCount
}

func (uc *UserEventCounter) IncreaseCommentCountByOne() {
	defer uc.rwMutex.RUnlock()
	uc.rwMutex.RLock()
	oldCommentCount := uc.commentCounts
	newCommentCount := oldCommentCount + 1
	uc.commentCounts = newCommentCount
}

func (uc *UserEventCounter) IncreasePostCountByOne() {
	defer uc.rwMutex.RUnlock()
	uc.rwMutex.RLock()
	oldPostCount := uc.postCounts
	newPostCount := oldPostCount + 1
	uc.postCounts = newPostCount
}
