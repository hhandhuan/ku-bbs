package frontend

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/hhandhuan/ku-bbs/internal/consts"
	fe "github.com/hhandhuan/ku-bbs/internal/entity/frontend"
	remindSub "github.com/hhandhuan/ku-bbs/internal/subject/remind"
	"github.com/hhandhuan/ku-bbs/pkg/utils/page"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/o1egl/govatar"
	"gorm.io/gorm"

	"github.com/hhandhuan/ku-bbs/internal/model"
	"github.com/hhandhuan/ku-bbs/internal/service"
	"github.com/hhandhuan/ku-bbs/pkg/config"
	"github.com/hhandhuan/ku-bbs/pkg/utils/encrypt"
)

func UserService(ctx *gin.Context) *sUser {
	return &sUser{ctx: service.Context(ctx)}
}

type sUser struct {
	ctx *service.BaseContext
}

// Register 用户登录
func (s *sUser) Register(req *fe.RegisterReq) error {
	var user *model.Users
	err := model.User().M.Where("name = ?", req.Name).Find(&user).Error
	if err != nil {
		return errors.New("服务内部错误")
	}
	if user.ID > 0 {
		return errors.New("用户名已被注册，请更换用户名继续尝试")
	}

	avatar, err := s.genAvatar(req.Name, req.Gender)
	if err != nil {
		return errors.New("用户默认头像生成失败")
	}

	res := model.User().M.Create(&model.Users{
		Name:     req.Name,
		Avatar:   avatar,
		Password: encrypt.GenerateFromPassword(req.Password),
		Gender:   uint8(req.Gender),
		State:    consts.EnableState,
	})
	if res.Error != nil || res.RowsAffected <= 0 {
		return errors.New("用户注册失败，请稍后在试")
	}

	return nil
}

// genAvatar 生成用户默认头像
func (s *sUser) genAvatar(name string, gender uint) (string, error) {
	path := fmt.Sprintf("%s/users/", config.Conf.Upload.Path)

	// 检查目录是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
		_ = os.Chmod(path, os.ModePerm)
	}

	avatarName := encrypt.Md5(gconv.String(time.Now().UnixMicro()))
	avatarPath := fmt.Sprintf("users/%s.png", avatarName)
	uploadPath := fmt.Sprintf("%s/%s", config.Conf.Upload.Path, avatarPath)

	if err := govatar.GenerateFileForUsername(govatar.Gender(gender-1), name, uploadPath); err != nil {
		log.Println(err)
		return "", err
	} else {
		return "/assets/upload/" + avatarPath, nil
	}
}

// Login 处理用户登录
func (s *sUser) Login(req *fe.LoginReq) error {
	var user model.Users
	err := model.User().M.Where("name = ?", req.Name).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("服务内部错误")
	}

	if user.ID <= 0 || !encrypt.CompareHashAndPassword(user.Password, req.Password) {
		return errors.New("用户名或密码错误")
	}

	// 账户是否禁用
	if user.State == consts.DisableState {
		return errors.New("账户已被禁用")
	}

	data := map[string]interface{}{
		"last_login_at": time.Now(),
		"last_login_ip": s.ctx.Ctx.ClientIP(),
	}

	u := model.User().M.Where("id", user.ID).Updates(data)
	if u.Error != nil || u.RowsAffected <= 0 {
		return fmt.Errorf("登录失败，服务内部错误: %v", u.Error)
	}

	s.ctx.SetAuth(user)

	return nil

}

// Logout 用户登出
func (s *sUser) Logout() {
	s.ctx.Forget()
}

// Edit 编辑用户
func (s *sUser) Edit(req *fe.EditUserReq) error {
	var user model.Users
	f := model.User().M.Where("name", req.Name).Find(&user)
	if f.Error != nil {
		log.Println(f.Error)
		return fmt.Errorf("修改信息失败: %v", f.Error)
	}
	currUser := s.ctx.Auth()
	if user.ID > 0 && currUser.ID != user.ID {
		return errors.New("用户名已存在")
	}
	data := &model.Users{
		Name:    req.Name,
		Email:   req.Email,
		Gender:  uint8(req.Gender),
		City:    req.City,
		Site:    req.Site,
		Job:     req.Job,
		Desc:    req.Desc,
		State:   currUser.State,
		IsAdmin: currUser.IsAdmin,
	}
	u := model.User().M.Where("id", currUser.ID).Updates(data)
	if u.Error != nil {
		return fmt.Errorf("修改信息失败: %v", f.Error)
	}
	if u.RowsAffected <= 0 {
		return errors.New("修改信息失败")
	}

	model.User().M.Where("id", currUser.ID).Find(&user)

	s.ctx.SetAuth(user)

	return nil
}

// EditPassword 修改密码
func (s *sUser) EditPassword(req *fe.EditPasswordReq) error {
	currUser := s.ctx.Auth()

	if !encrypt.CompareHashAndPassword(currUser.Password, req.OldPassword) {
		return errors.New("旧密码错误")
	}

	u := model.User().M.Where("id", currUser.ID).Update("password", encrypt.GenerateFromPassword(req.Password))
	if u.Error != nil || u.RowsAffected <= 0 {
		log.Println(u.Error)
		return errors.New("修改密码失败")
	}

	s.ctx.Forget()

	return nil
}

// EditAvatar 修改头像
func (s *sUser) EditAvatar(ctx *gin.Context) error {
	file, err := ctx.FormFile("avatar")
	if err != nil {
		return err
	}

	// 目前限制头像大小
	if file.Size > 1024*1024*config.Conf.Upload.AvatarFileSize {
		return errors.New("仅支持小于 1M 大小的图片")
	}

	arr := strings.Split(file.Filename, ".")
	ext := arr[len(arr)-1]

	// 检查图片格式
	if !gstr.InArray(config.Conf.Upload.ImageExt, ext) {
		return errors.New("file format not supported")
	}

	avatarName := encrypt.Md5(gconv.String(time.Now().UnixMicro()))
	avatarPath := fmt.Sprintf("users/%s.png", avatarName)
	uploadPath := fmt.Sprintf("%s/%s", config.Conf.Upload.Path, avatarPath)

	if err := ctx.SaveUploadedFile(file, uploadPath); err != nil {
		log.Println(err)
		return errors.New("修改头像失败")
	}

	userID := s.ctx.Auth().ID
	savePath := "/assets/upload/" + avatarPath
	u := model.User().M.Where("id", userID).Update("avatar", savePath)
	if u.Error != nil || u.RowsAffected <= 0 {
		log.Println(u.Error)
		return errors.New("修改头像失败")
	}

	var user model.Users
	model.User().M.Where("id", userID).Find(&user)

	s.ctx.SetAuth(user)

	return nil
}

// Home 用户主页
func (s *sUser) Home(req *fe.GetUserHomeReq) (gin.H, error) {
	var user *fe.User
	if req.Tab == "" {
		req.Tab = consts.UserTopicTab
	}

	query := model.User().M.Where("id", req.ID)
	if s.ctx.Check() {
		query = query.Preload("Follow", "user_id = ? AND state = ?", s.ctx.Auth().ID, consts.FollowedState)
	}

	if r := query.Limit(1).Find(&user); r.Error != nil {
		return nil, r.Error
	}
	if user.ID <= 0 {
		return nil, errors.New("用户不存在")
	}

	if req.Tab == consts.UserTopicTab {
		var (
			list   []*fe.Topic
			total  int64
			limit  = 20
			offset = (req.Page - 1) * limit
		)

		query = model.Topic().M.Where("user_id", req.ID)
		if c := query.Count(&total); c.Error != nil {
			return nil, c.Error
		}

		if f := query.Preload("Node").Limit(limit).Offset(offset).Find(&list); f.Error != nil {
			return nil, f.Error
		}

		baseUrl := s.ctx.Ctx.Request.RequestURI
		pageObj := page.New(int(total), limit, gconv.Int(req.Page), baseUrl)

		return gin.H{"user": user, "list": list, "req": req, "page": pageObj}, nil
	} else if req.Tab == consts.UserFollowTab {
		var (
			list   []*fe.Follow
			total  int64
			limit  = 20
			offset = (req.Page - 1) * limit
		)

		query = model.Follow().M.Where("user_id", req.ID).Where("state", consts.FollowedState)
		if c := query.Count(&total); c.Error != nil {
			return nil, c.Error
		}

		f := query.Preload("Fans").Limit(limit).Offset(offset).Find(&list)
		if f.Error != nil {
			return nil, f.Error
		}

		baseUrl := s.ctx.Ctx.Request.RequestURI
		pageObj := page.New(int(total), limit, gconv.Int(req.Page), baseUrl)

		return gin.H{"user": user, "list": list, "req": req, "page": pageObj}, nil
	} else {
		var (
			list   []*fe.Follow
			total  int64
			limit  = 20
			offset = (req.Page - 1) * limit
		)
		query = model.Follow().M.Where("target_id", req.ID).Where("state", consts.FollowedState)
		if c := query.Count(&total); c.Error != nil {
			return nil, c.Error
		}

		f := query.Preload("Follower").Limit(limit).Offset(offset).Find(&list)
		if f.Error != nil {
			return nil, f.Error
		}

		baseUrl := s.ctx.Ctx.Request.RequestURI
		pageObj := page.New(int(total), limit, gconv.Int(req.Page), baseUrl)

		return gin.H{"user": user, "list": list, "req": req, "page": pageObj}, nil
	}
}

// Follow 关注
func (s *sUser) Follow(req *fe.FollowUserReq) (int, error) {
	if req.UserID == s.ctx.Auth().ID {
		return 0, errors.New("无法关注自己")
	}

	var user *model.Users
	err := model.User().M.Where("id", req.UserID).Find(&user).Error
	if err != nil {
		return 0, err
	}
	if user == nil || user.ID <= 0 {
		return 0, errors.New("用户不存在")
	}

	var follow *model.Follows
	err = model.Follow().M.Where("user_id = ? AND target_id = ?", s.ctx.Auth().ID, req.UserID).Find(&follow).Error
	if err != nil {
		return 0, err
	}

	if follow.ID <= 0 {
		data := &model.Follows{UserId: s.ctx.Auth().ID, TargetId: req.UserID, State: 1}
		if c := model.Follow().M.Create(data); c.Error != nil || c.RowsAffected <= 0 {
			log.Println(c.Error)
			return 0, errors.New("关注失败")
		}

		sub := remindSub.New()
		sub.Attach(&remindSub.FollowObs{Sender: s.ctx.Auth().ID, Receiver: req.UserID})
		sub.Notify()

		return 1, nil
	}

	state := consts.UnFollowedState
	if follow.State == consts.UnFollowedState {
		state = consts.FollowedState
	}

	if err = model.Follow().M.Where("id", follow.ID).Update("state", state).Error; err != nil {
		return 0, err
	}

	return state, nil
}
