// Package p355 implements a simplified Twitter-like social media system.
package p355

import (
	"container/heap"
	"time"
)

// Twitter represents a social media platform supporting tweets, follows, and news feeds.
// Users can post tweets, follow other users, and retrieve a personalized news feed.
type Twitter struct {
	followees map[int]map[int]bool // Maps userId to set of users they follow (userId -> followeeId -> isFollowed)
	tweets    map[int][]Tweet      // Maps userId to their tweets in chronological order (oldest to newest)
}

// Constructor creates and returns a new Twitter instance with empty follow relationships and tweets.
func Constructor() Twitter {
	return Twitter{
		followees: make(map[int]map[int]bool),
		tweets:    make(map[int][]Tweet),
	}
}

// PostTweet adds a new tweet for the specified user with the current timestamp.
// Tweets are stored in chronological order for efficient news feed generation.
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.tweets[userId] = append(t.tweets[userId], Tweet{
		id:        tweetId,
		userID:    userId,
		timestamp: time.Now(),
	})
}

// GetNewsFeed returns up to 10 most recent tweets from the user and their followees.
// Uses a max-heap to efficiently merge tweets from multiple users in timestamp order.
// Automatically includes the user's own tweets by temporarily adding self-follow relationship.
func (t *Twitter) GetNewsFeed(userId int) []int {
	var retval []int
	tweetHeap := &TweetHeap{}
	heap.Init(tweetHeap)

	// Ensure user can see their own tweets by adding self-follow
	if t.followees[userId] == nil {
		t.followees[userId] = make(map[int]bool)
	}
	t.followees[userId][userId] = true

	// Add most recent tweet from each followed user (including self) to heap
	for followeeId, isFollowed := range t.followees[userId] {
		if isFollowed {
			if tweets, ok := t.tweets[followeeId]; ok && len(tweets) > 0 {
				index := len(tweets) - 1 // Start from most recent tweet
				tweet := tweets[index]
				tweet.index = index // Track position for lazy loading of older tweets
				heap.Push(tweetHeap, tweet)
			}
		}
	}

	// Extract up to 10 most recent tweets using heap
	for tweetHeap.Len() > 0 && len(retval) < 10 {
		tweet := heap.Pop(tweetHeap).(Tweet)
		retval = append(retval, tweet.id)

		// Lazy loading: add next older tweet from same user if available
		if tweet.index > 0 {
			nextIndex := tweet.index - 1
			next := t.tweets[tweet.userID][nextIndex]
			next.index = nextIndex
			heap.Push(tweetHeap, next)
		}
	}

	return retval
}

// Follow creates a follow relationship from followerId to followeeId.
// Initializes the follower's followee map if it doesn't exist.
func (t *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := t.followees[followerId]; !ok {
		t.followees[followerId] = make(map[int]bool)
	}
	t.followees[followerId][followeeId] = true
}

// Unfollow removes the follow relationship from followerId to followeeId.
// Sets the relationship to false rather than deleting to maintain map structure.
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := t.followees[followerId]; !ok {
		return
	}
	t.followees[followerId][followeeId] = false
}

// isFollowee checks if userId follows followeeId.
// Returns false if the user has no follow relationships or doesn't follow the specified user.
func (t *Twitter) isFollowee(userId int, followeeId int) bool {
	if _, ok := t.followees[userId]; !ok {
		return false
	}
	return t.followees[userId][followeeId]
}

// Tweet represents a single tweet with metadata for efficient news feed processing.
type Tweet struct {
	id        int       // Unique tweet identifier
	userID    int       // ID of user who posted the tweet
	timestamp time.Time // When the tweet was posted (for chronological ordering)
	index     int       // Position in user's tweet slice (for lazy loading of older tweets)
}

// TweetHeap implements a max-heap of Tweets ordered by timestamp (most recent first).
// Used to efficiently merge tweets from multiple users for news feed generation.
type TweetHeap []Tweet

// Len returns the number of tweets in the heap.
func (h TweetHeap) Len() int { return len(h) }

// Less compares tweets by timestamp for max-heap ordering (newer tweets have higher priority).
func (h TweetHeap) Less(i, j int) bool { return h[i].timestamp.After(h[j].timestamp) }

// Swap exchanges two tweets in the heap.
func (h TweetHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds a tweet to the heap. Uses pointer receiver because it modifies slice length.
func (h *TweetHeap) Push(x any) {
	*h = append(*h, x.(Tweet))
}

// Pop removes and returns the most recent tweet from the heap. Uses pointer receiver.
func (h *TweetHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
