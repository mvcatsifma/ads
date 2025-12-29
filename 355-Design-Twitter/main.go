package p355

import (
	"container/heap"
	"time"
)

type Twitter struct {
	followees map[int]map[int]bool
	h         *TweetHeap
}

func Constructor() Twitter {
	h := &TweetHeap{}
	heap.Init(h)

	return Twitter{
		followees: make(map[int]map[int]bool),
		h:         h,
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	heap.Push(t.h, Tweet{
		id:        tweetId,
		userID:    userId,
		timestamp: time.Now(),
	})
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	var retval []int
	var tweets []Tweet
	for t.h.Len() > 0 {
		tw := heap.Pop(t.h).(Tweet)
		if tw.userID == userId || t.isFollowee(userId, tw.userID) {
			retval = append(retval, tw.id)
		}
		tweets = append(tweets, tw)
		if len(retval) == 10 {
			break
		}
	}

	for _, tw := range tweets {
		heap.Push(t.h, tw)
	}

	return retval
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := t.followees[followerId]; !ok {
		t.followees[followerId] = make(map[int]bool)
	}
	t.followees[followerId][followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := t.followees[followerId]; !ok {
		return
	}
	t.followees[followerId][followeeId] = false
}

func (t *Twitter) isFollowee(userId int, followeeId int) bool {
	if _, ok := t.followees[userId]; !ok {
		return false
	}
	return t.followees[userId][followeeId]
}

type Tweet struct {
	id        int
	userID    int
	timestamp time.Time
}

// An TweetHeap is a max-heap of ints ordered by timestamp
type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].timestamp.After(h[j].timestamp) }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TweetHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Tweet))
}

func (h *TweetHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
