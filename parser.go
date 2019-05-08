package main

import (
	"encoding/csv"
	"io"
	"strings"
	"time"
)

type stat struct {
	TweetID                  string // ツイートID
	TweetURL                 string // ツイートの固定リンク
	Tweet                    string // ツイート本文
	TweetedAt                string // 時間
	Impression               string // インプレッション
	Engatement               string // エンゲージメント
	EngatementRate           string // エンゲージメント率
	Retweet                  string // リツイート
	Reply                    string // 返信
	Like                     string // いいね
	ProfileClick             string // ユーザープロフィールクリック
	URLClick                 string // URLクリック数
	HashTagClick             string // ハッシュタグクリック
	DetailClick              string // 詳細クリック
	TweetURLClick            string // 固定リンクのクリック数
	AppDisplay               string // アプリ表示
	AppInstall               string // アプリインストール
	Followers                string // フォローしている
	EmailTweet               string // ツイートをメール送信
	Tel                      string // ダイアル式電話
	PlayMedia                string // メディアの再生数
	MediaEngagement          string // メディアのエンゲージメント
	PromotionImpression      string // プロモのインプレッション
	PromotionEngagement      string // プロモのエンゲージメント
	PromotionEngagementRate  string // プロモのエンゲージメント率
	PromotionRetweet         string // プロモのリツイート
	PromotionReply           string // プロモの返信
	PromotionLike            string // プロモのいいね
	PromotionProfileClick    string // プロモのユーザープロフィールクリック
	PromotionURLClick        string // プロモのURLクリック数
	PromotionHashTagClick    string // プロモのハッシュタグクリック
	PromotionDetailClick     string // プロモの詳細クリック
	PromotionTweetURLClick   string // プロモの固定リンクのクリック数
	PromotionAppDisplay      string // プロモのアプリ表示
	PromotionAppInstall      string // プロモのアプリインストール
	PromotionFollowers       string // プロモのフォローしている
	PromotionEmailTweet      string // プロモのツイートをメール送信
	PromotionTel             string // プロモのダイアル式電話
	PromotionMediaPlay       string // プロモのメディアの再生数
	PromotionMediaEngagement string // プロモのメディアのエンゲージメント
}

func parse(r []string) (*stat, error) {
	var d stat
	d.TweetID = r[0]
	d.TweetURL = r[1]
	d.Tweet = strings.Replace(r[2], "\n", " ", -1)
	// format=2019-05-08 08:55 +0000
	tt, err := time.Parse("2006-01-02 15:04 +0000", r[3])
	if err != nil {
		return nil, err
	}
	// TweetedAt at local timezone
	d.TweetedAt = tt.Local().Format("2006-01-02 15:04")
	d.Impression = r[4]
	d.Engatement = r[5]
	d.EngatementRate = r[6]
	d.Retweet = r[7]
	d.Reply = r[8]
	d.Like = r[9]
	d.ProfileClick = r[10]
	d.URLClick = r[11]
	d.HashTagClick = r[12]
	d.DetailClick = r[13]
	d.TweetURLClick = r[14]
	d.AppDisplay = r[15]
	d.AppInstall = r[16]
	d.Followers = r[17]
	d.EmailTweet = r[18]
	d.Tel = r[19]
	d.PlayMedia = r[20]
	d.MediaEngagement = r[21]
	d.PromotionImpression = r[22]
	d.PromotionEngagement = r[23]
	d.PromotionEngagementRate = r[24]
	d.PromotionRetweet = r[25]
	d.PromotionReply = r[26]
	d.PromotionLike = r[27]
	d.PromotionProfileClick = r[28]
	d.PromotionURLClick = r[29]
	d.PromotionHashTagClick = r[30]
	d.PromotionDetailClick = r[31]
	d.PromotionTweetURLClick = r[32]
	d.PromotionAppDisplay = r[33]
	d.PromotionAppInstall = r[34]
	d.PromotionFollowers = r[35]
	d.PromotionEmailTweet = r[36]
	d.PromotionTel = r[37]
	d.PromotionMediaPlay = r[38]
	d.PromotionMediaEngagement = r[39]
	return &d, nil
}

func parser(r io.Reader) ([]*stat, error) {
	reader := csv.NewReader(r)
	recs, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var stats []*stat
	for i, rec := range recs {
		if i == 0 {
			continue
		}
		s, err := parse(rec)
		if err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}
	return stats, nil
}
