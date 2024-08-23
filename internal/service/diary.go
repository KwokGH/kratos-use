package service

import (
	"context"
	"kratos-use/api/mini/diary/v1"
	"kratos-use/internal/biz"
)

type DiaryService struct {
	//v1.UnimplementedDiaryServer

	diaryUseCase *biz.DiaryUsecase
	userUseCase  *biz.UserUsecase
}

func NewDiaryService(diaryUseCase *biz.DiaryUsecase) *DiaryService {
	return &DiaryService{
		diaryUseCase: diaryUseCase,
	}
}

func (s *DiaryService) ListDiary(ctx context.Context, req *v1.ListDiaryReq) (*v1.ListDiaryResp, error) {
	total, list, err := s.diaryUseCase.List(ctx, &biz.ListInput{
		Keyword:  req.Keywords,
		StartAt:  req.StartAt,
		EndAt:    req.EndAt,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	resp := &v1.ListDiaryResp{
		Total: total,
		List:  make([]*v1.ListItemDiaryResp, 0, len(list)),
	}

	for _, item := range list {
		respItem := &v1.ListItemDiaryResp{
			Id:        item.Id,
			Title:     item.Title,
			CreatedAt: item.CreateAt,
			UpdatedAt: item.UpdateAt,
			BelongAt:  item.BelongAt,
			Desc:      item.ShortContent,
			Tag:       item.Tag,
		}
		resp.List = append(resp.List, respItem)
	}

	return resp, nil
}
func (s *DiaryService) GetDiary(ctx context.Context, req *v1.GetDiaryReq) (*v1.GetDiaryResp, error) {
	out, err := s.diaryUseCase.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetDiaryResp{
		Id:        out.Id,
		Title:     out.Title,
		CreatedAt: out.CreateAt,
		UpdatedAt: out.UpdateAt,
		Content:   out.Content,
		BelongAt:  out.BelongAt,
		Tag:       out.Tag,
	}, nil
}
func (s *DiaryService) CreateDiary(ctx context.Context, req *v1.CreateDiaryReq) (*v1.CreateDiaryResp, error) {
	out, err := s.diaryUseCase.Create(ctx, &biz.Diary{
		Title:    req.Title,
		Content:  req.Content,
		BelongAt: req.BelongAt,
		Tag:      req.Tag,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateDiaryResp{
		Id: out,
	}, nil
}
func (s *DiaryService) UpdateDiary(ctx context.Context, req *v1.UpdateDiaryReq) (*v1.UpdateDiaryResp, error) {
	out, err := s.diaryUseCase.Update(ctx, req.Id, &biz.Diary{
		Title:    req.Title,
		Content:  req.Content,
		BelongAt: req.BelongAt,
		Tag:      req.Tag,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UpdateDiaryResp{
		Id: out,
	}, nil
}

func (s *DiaryService) DeleteDiary(ctx context.Context, req *v1.DeleteDiaryReq) (*v1.DeleteDiaryResp, error) {
	err := s.diaryUseCase.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteDiaryResp{
		Id: req.Id,
	}, nil
}
