package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Diary struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	BelongAt     int64  `json:"belong_at"`
	UserId       string `json:"user_id"`
	CreateAt     int64  `json:"create_at"`
	UpdateAt     int64  `json:"update_at"`
	ShortContent string `json:"short_content"`
	Tag          string `json:"tag"`
}

// DiaryRepo is a Greater repo.
type DiaryRepo interface {
	Create(ctx context.Context, diary *Diary) (string, error)
	Update(ctx context.Context, id string, diary *Diary) (string, error)
	Get(ctx context.Context, id string) (*Diary, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, input *ListInput) (int32, []*Diary, error)
}

type ListInput struct {
	Keyword  string
	StartAt  int64
	EndAt    int64
	Page     int32
	PageSize int32
}

// DiaryUsecase is a Diary usecase.
type DiaryUsecase struct {
	repo DiaryRepo
	log  *log.Helper
	tx   Transaction
}

func (d *DiaryUsecase) Create(ctx context.Context, diary *Diary) (string, error) {
	return d.repo.Create(ctx, diary)
}

func (d *DiaryUsecase) Update(ctx context.Context, id string, diary *Diary) (string, error) {
	return d.repo.Update(ctx, id, diary)
}

func (d *DiaryUsecase) Get(ctx context.Context, id string) (*Diary, error) {
	return d.repo.Get(ctx, id)
}

func (d *DiaryUsecase) Delete(ctx context.Context, id string) error {
	return d.repo.Delete(ctx, id)
}

func (d *DiaryUsecase) List(ctx context.Context, input *ListInput) (int32, []*Diary, error) {
	return d.repo.List(ctx, input)
}

// NewDiaryUsecase new a Diary usecase.
func NewDiaryUsecase(repo DiaryRepo, tx Transaction, logger log.Logger) *DiaryUsecase {
	return &DiaryUsecase{repo: repo, tx: tx, log: log.NewHelper(logger)}
}
