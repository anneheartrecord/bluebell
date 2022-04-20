package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" `
}
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"'`
	Direction int    `json:"direction"  binding:"oneof=1 0 -1"`
}
type ParamPostList struct {
	Page        int64  `form:"page" json:"page"`
	Size        int64  `form:"size" json:"size"`
	Order       string `form:"order" json:"order"`
	CommunityID int64  `json:"community_id" form:"community_id"`
}
