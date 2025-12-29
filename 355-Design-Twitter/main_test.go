package p355

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwitter(t *testing.T) {
	t.Run("LeetCode example test case", func(t *testing.T) {
		// Input operations:
		// ["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
		// [[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]

		// Expected output:
		// [null, null, [5], null, null, [6, 5], null, [5]]

		// Initialize Twitter
		twitter := Constructor()

		// Operation 1: postTweet(1, 5) → null
		twitter.PostTweet(1, 5)

		// Operation 2: getNewsFeed(1) → [5]
		feed := twitter.GetNewsFeed(1)
		assert.Equal(t, []int{5}, feed, "User 1 should see their own tweet")

		// Operation 3: follow(1, 2) → null
		twitter.Follow(1, 2)

		// Operation 4: postTweet(2, 6) → null
		twitter.PostTweet(2, 6)

		// Operation 5: getNewsFeed(1) → [6, 5]
		feed = twitter.GetNewsFeed(1)
		assert.Equal(t, []int{6, 5}, feed, "User 1 should see followee's tweet first (most recent)")

		// Operation 6: unfollow(1, 2) → null
		twitter.Unfollow(1, 2)

		// Operation 7: getNewsFeed(1) → [5]
		feed = twitter.GetNewsFeed(1)
		assert.Equal(t, []int{5}, feed, "User 1 should only see their own tweet after unfollowing")
	})

	t.Run("complex follow/unfollow test case", func(t *testing.T) {
		// Input operations:
		// ["Twitter","postTweet","follow","follow","getNewsFeed","postTweet","getNewsFeed","getNewsFeed","unfollow","getNewsFeed","getNewsFeed","unfollow","getNewsFeed","getNewsFeed"]
		// [[],[1,5],[1,2],[2,1],[2],[2,6],[1],[2],[2,1],[1],[2],[1,2],[1],[2]]

		// Expected output:
		// [null,null,null,null,[5],null,[6,5],[6,5],null,[6,5],[6],null,[5],[6]]

		twitter := Constructor()

		// postTweet(1,5) → null
		twitter.PostTweet(1, 5)

		// follow(1,2) → null
		twitter.Follow(1, 2)

		// follow(2,1) → null
		twitter.Follow(2, 1)

		// getNewsFeed(2) → [5]
		feed := twitter.GetNewsFeed(2)
		assert.Equal(t, []int{5}, feed, "User 2 should see user 1's tweet after following")

		// postTweet(2,6) → null
		twitter.PostTweet(2, 6)

		// getNewsFeed(1) → [6,5]
		feed = twitter.GetNewsFeed(1)
		assert.Equal(t, []int{6, 5}, feed, "User 1 should see user 2's tweet first (most recent), then own tweet")

		// getNewsFeed(2) → [6,5]
		feed = twitter.GetNewsFeed(2)
		assert.Equal(t, []int{6, 5}, feed, "User 2 should see own tweet first, then user 1's tweet")

		// unfollow(2,1) → null
		twitter.Unfollow(2, 1)

		// getNewsFeed(1) → [6,5]
		feed = twitter.GetNewsFeed(1)
		assert.Equal(t, []int{6, 5}, feed, "User 1 still follows user 2, should see both tweets")

		// getNewsFeed(2) → [6]
		feed = twitter.GetNewsFeed(2)
		assert.Equal(t, []int{6}, feed, "User 2 unfollowed user 1, should only see own tweet")

		// unfollow(1,2) → null
		twitter.Unfollow(1, 2)

		// getNewsFeed(1) → [5]
		feed = twitter.GetNewsFeed(1)
		assert.Equal(t, []int{5}, feed, "User 1 unfollowed user 2, should only see own tweet")

		// getNewsFeed(2) → [6]
		feed = twitter.GetNewsFeed(2)
		assert.Equal(t, []int{6}, feed, "User 2 should see own tweet")
	})

	t.Run("self-follow prevention", func(t *testing.T) {
		twitter := Constructor()

		// Test empty news feed
		feed := twitter.GetNewsFeed(1)
		assert.Empty(t, feed, "Empty news feed for user with no tweets or follows")

		// User posts a tweet
		twitter.PostTweet(1, 100)

		// Attempt self-follow (should be ignored)
		twitter.Follow(1, 1)

		// User should see their own tweet (not because of following, but because it's their own)
		feed = twitter.GetNewsFeed(1)
		assert.Equal(t, []int{100}, feed, "User sees own tweet regardless of self-follow attempt")

		// Post another tweet
		twitter.PostTweet(1, 101)

		// Attempt self-unfollow (should have no effect since self-follow was ignored)
		twitter.Unfollow(1, 1)

		// Should still see both tweets
		feed = twitter.GetNewsFeed(1)
		assert.Contains(t, feed, 100, "Should still see own tweets")
		assert.Contains(t, feed, 101, "Should still see own tweets")

		// Test unfollow non-existent relationship
		twitter.Unfollow(1, 999) // Should not crash
	})

	t.Run("news feed ordering test", func(t *testing.T) {
		twitter := Constructor()

		// Create multiple users and tweets
		twitter.PostTweet(1, 1) // User 1 posts tweet 1
		twitter.PostTweet(2, 2) // User 2 posts tweet 2
		twitter.PostTweet(1, 3) // User 1 posts tweet 3

		twitter.Follow(1, 2) // User 1 follows user 2

		feed := twitter.GetNewsFeed(1)
		// Should be ordered by recency and include both own and followee's tweets
		assert.NotEmpty(t, feed, "Feed should not be empty")
		assert.Contains(t, feed, 1, "Should contain user's own tweets")
		assert.Contains(t, feed, 2, "Should contain followee's tweets")
		assert.Contains(t, feed, 3, "Should contain user's recent tweets")
	})

	t.Run("follow and unfollow operations", func(t *testing.T) {
		twitter := Constructor()

		twitter.PostTweet(1, 100)
		twitter.PostTweet(2, 200)

		// User 1 follows user 2
		twitter.Follow(1, 2)
		feed := twitter.GetNewsFeed(1)
		assert.Len(t, feed, 2, "Should see both own and followee's tweets")

		// User 1 unfollows user 2
		twitter.Unfollow(1, 2)
		feed = twitter.GetNewsFeed(1)
		assert.Equal(t, []int{100}, feed, "Should only see own tweets after unfollowing")
	})
}
