package main

import "fmt"

func format(s *stat) string {
	return fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t",
		s.TweetID, s.TweetURL, s.Tweet, s.TweetedAt, s.Impression, s.Engatement, s.EngatementRate, s.Retweet, s.Reply, s.Like, s.ProfileClick, s.URLClick, s.HashTagClick, s.DetailClick, s.TweetURLClick, s.AppDisplay, s.AppInstall, s.Followers, s.EmailTweet, s.Tel, s.PlayMedia, s.MediaEngagement, s.PromotionImpression, s.PromotionEngagement, s.PromotionEngagementRate, s.PromotionRetweet, s.PromotionReply, s.PromotionLike, s.PromotionProfileClick, s.PromotionURLClick, s.PromotionHashTagClick, s.PromotionDetailClick, s.PromotionTweetURLClick, s.PromotionAppDisplay, s.PromotionAppInstall, s.PromotionFollowers, s.PromotionEmailTweet, s.PromotionTel, s.PromotionMediaPlay, s.PromotionMediaEngagement)
}
