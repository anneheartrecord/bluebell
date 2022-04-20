package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("voteforpost", zap.Int64("userID", userID), zap.String("postID", p.PostID))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
