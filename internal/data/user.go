package data

import (
	"context"
	"github.com/mitchellh/mapstructure"
	api_common "kratos-use/api/common"
	"kratos-use/ent"
	"kratos-use/ent/user"
	"kratos-use/internal/biz"
	"kratos-use/internal/entity"
	"kratos-use/pkg/randx"
	"kratos-use/pkg/unique"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) CreateForRegister(ctx context.Context, input *entity.RegisterInput) (string, error) {
	passwordSalt := randx.GetRandomCharacter(8)
	md5Password := unique.GetMd5(input.Password + passwordSalt)

	newUser, err := r.data.DB(ctx).User.Create().
		SetAccount(input.Account).SetPassword(md5Password).SetPasswordSalt(passwordSalt).Save(ctx)

	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "创建用户失败", "err", err)
		if ent.IsConstraintError(err) {
			return "", entity.ErrUserConflict
		}
		return "", err
	}

	return newUser.ID, nil
}

func (r *userRepo) GetUserById(ctx context.Context, id string) (*entity.UserDTO, error) {
	out, err := r.data.DB(ctx).User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, api_common.ErrorNotFound("")
		}
		return nil, err
	}

	result := &entity.UserDTO{}
	err = mapstructure.Decode(out, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepo) GetUserByAccount(ctx context.Context, account string) (*entity.UserDTO, error) {
	out, err := r.data.DB(ctx).User.Query().Where(user.AccountEQ(account)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, api_common.ErrorNotFound("")
		}
		return nil, err
	}

	result := &entity.UserDTO{}
	err = mapstructure.Decode(out, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//func (r *userRepo) LoginByWechat(ctx context.Context, input *biz.LoginByWechatInput) (*biz.LoginByWechatOutput, error) {
//	wechatCfg := r.data.bootstrap.App.Wechat
//	// 根据登录凭证获取session
//	code2SessionInput := &wechatx.Code2SessionInput{
//		Appid:     wechatCfg.AppId,
//		AppSecret: wechatCfg.Secret,
//		Code:      input.Code,
//	}
//	code2SessionOutput, err := wechatx.Code2Session(ctx, code2SessionInput)
//	if err != nil {
//		return nil, biz.ErrInternalServer
//	}
//
//	if code2SessionOutput.Errcode != 0 {
//		return nil, biz.ErrInternalServer
//	}
//
//	userId := ""
//	openId := code2SessionOutput.Openid
//
//	// 从数据库中查询凭证信息
//	firstAuth, err := r.data.DB(ctx).ThirdAuth.Query().Where(thirdauth.Openid(openId)).First(ctx)
//	if err != nil {
//		if ent.IsNotFound(err) {
//			// 凭证不存在，新增凭证和用户信息
//			err := r.data.ExecTx(ctx, func(ctx context.Context) error {
//				newUser, err := r.data.DB(ctx).User.Create().
//					SetName("灵沐" + randx.GetValidateCode(4)).Save(ctx)
//				if err != nil {
//					return err
//				}
//
//				userId = newUser.ID
//
//				_, err = r.data.DB(ctx).ThirdAuth.Create().
//					SetOpenid(openId).
//					SetUnionid(code2SessionOutput.Unionid).
//					SetUserID(userId).
//					Save(ctx)
//				if err != nil {
//					return err
//				}
//
//				return nil
//			})
//			if err != nil {
//				return nil, err
//			}
//		} else {
//			return nil, err
//		}
//	} else {
//		// 凭证存在，直接返回token
//		userId = firstAuth.UserID
//	}
//
//	// 构造token
//	auth := r.data.bootstrap.App.Auth
//	token, err := jwtx.CreateJwtToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, map[string]interface{}{
//		"userId": userId,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &biz.LoginByWechatOutput{
//		AuthToken: token,
//		OpenId:    openId,
//	}, nil
//}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
