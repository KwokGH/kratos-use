package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-use/api/common"
	"kratos-use/ent"
	"kratos-use/ent/diary"
	"kratos-use/ent/predicate"
	"kratos-use/internal/biz"
	"kratos-use/pkg/dbtool"
	"kratos-use/pkg/middleware"
	"kratos-use/pkg/stringx"
)

type diaryRepo struct {
	data *Data
	log  *log.Helper

	userRepo biz.UserRepo
}

func (d *diaryRepo) Create(ctx context.Context, diary *biz.Diary) (string, error) {
	userId := middleware.GetUserId(ctx)

	save, err := d.data.DB(ctx).Diary.Create().
		SetBelongAt(diary.BelongAt).
		SetTitle(diary.Title).
		SetContent(diary.Content).
		SetUserID(userId).
		SetTag(diary.Tag).
		Save(ctx)
	if err != nil {
		return "", err
	}

	return save.ID, nil
}

func (d *diaryRepo) Update(ctx context.Context, id string, input *biz.Diary) (string, error) {
	userId := middleware.GetUserId(ctx)

	save, err := d.data.DB(ctx).Diary.UpdateOneID(id).Where(diary.UserIDEQ(userId)).
		SetTitle(input.Title).
		SetContent(input.Content).
		SetBelongAt(input.BelongAt).
		SetTag(input.Tag).
		Save(ctx)
	if err != nil {
		return "", err
	}

	return save.ID, nil
}

func (d *diaryRepo) Get(ctx context.Context, id string) (*biz.Diary, error) {
	userId := middleware.GetUserId(ctx)

	diaryItem, err := d.data.DB(ctx).Diary.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if diaryItem.UserID != userId {
		return nil, common.ErrorNotForbidden("")
	}

	return &biz.Diary{
		Id:       diaryItem.ID,
		Title:    diaryItem.Title,
		Content:  diaryItem.Content,
		BelongAt: diaryItem.BelongAt,
		CreateAt: diaryItem.CreatedAt,
		UpdateAt: diaryItem.UpdatedAt,
		Tag:      diaryItem.Tag,
	}, nil
}

func (d *diaryRepo) Delete(ctx context.Context, id string) error {
	userId := middleware.GetUserId(ctx)
	err := d.data.DB(ctx).Diary.DeleteOneID(id).Where(diary.UserIDEQ(userId)).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	}

	return nil
}

func (d *diaryRepo) List(ctx context.Context, input *biz.ListInput) (int32, []*biz.Diary, error) {
	userId := middleware.GetUserId(ctx)

	offset, limit := dbtool.NewPager(input.Page, input.PageSize).Offset()

	cond := []predicate.Diary{
		diary.UserIDEQ(userId),
	}
	if input.Keyword != "" {
		cond = append(cond, diary.Or(diary.TagContains(input.Keyword), diary.TitleContains(input.Keyword)))
	}
	if input.EndAt > 0 || input.StartAt > 0 {
		cond = append(cond, diary.BelongAtGTE(input.StartAt))
		cond = append(cond, diary.BelongAtLTE(input.EndAt))
	}

	diaries, err := d.data.DB(ctx).Diary.Query().Where(cond...).Order(diary.ByCreatedAt(sql.OrderDesc())).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return 0, nil, err
	}

	//total, err := d.data.DB(ctx).Diary.Query().Where(cond...).Count(ctx)
	//if err != nil {
	//	return 0, nil, err
	//}

	list := make([]*biz.Diary, 0, len(diaries))
	for _, diaryItem := range diaries {
		list = append(list, &biz.Diary{
			Id:           diaryItem.ID,
			Title:        diaryItem.Title,
			Content:      diaryItem.Content,
			BelongAt:     diaryItem.BelongAt,
			CreateAt:     diaryItem.CreatedAt,
			UpdateAt:     diaryItem.UpdatedAt,
			ShortContent: stringx.GetShortContent(diaryItem.Content, 200),
			Tag:          diaryItem.Tag,
		})
	}

	return int32(0), list, nil
}

// NewDiaryRepo .
func NewDiaryRepo(data *Data, logger log.Logger, userRepo biz.UserRepo) biz.DiaryRepo {
	return &diaryRepo{
		data:     data,
		log:      log.NewHelper(logger),
		userRepo: userRepo,
	}
}
