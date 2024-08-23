package data

import (
	"github.com/mitchellh/mapstructure"
	"testing"
)

func TestNewData(t *testing.T) {
	/*cfg := &conf.Bootstrap{}

	cfg.Data = &conf.Data{
		Database: &conf.Data_Database{
			Driver: "mysql",
			Source: "root:root1234@tcp(127.0.0.1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
	db, err := NewDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	mllmGrpcClient := NewMLLMClient()

	baseDB := mysql.NewBaseDBData(db)
	data, _, err := NewData(cfg, baseDB, mllmGrpcClient)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	resp, err := data.mllmClient.Ping(ctx, &mllm.PingReq{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)*/

	// 普通的操作
	//user, err := data.DB(ctx).User.Create().SetName("zhangsan").SetEmail("qq.com").Save(ctx)
	//if err != nil {
	//	t.Fatal(err)
	//}

	// // 事务操作
	// err = data.ExecTx(ctx, func(ctx context.Context) error {
	// 	user2, err := data.DB(ctx).User.Create().SetName("zhangsan4").SetEmail("qq2.com").Save(ctx)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	t.Log(user2)

	// 	user3, err := data.DB(ctx).User.Create().SetName("zhangsan5").SetEmail("qq3.com").Save(ctx)
	// 	t.Log(user3)

	// 	return nil
	// })

	// t.Log(err)
}

func TestMapStructure(t *testing.T) {
	type Stu struct {
		Id   string
		Name string
	}
	type Stu2 struct {
		Id   string
		Name string
	}
	s1 := []*Stu{
		&Stu{
			Id:   "1",
			Name: "111",
		},
		&Stu{
			Id:   "2",
			Name: "222",
		},
	}

	s2 := make([]*Stu2, 0)
	err := mapstructure.Decode(s1, &s2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(s2[0].Name)

	ss1 := &Stu{
		Id:   "1",
		Name: "111",
	}
	ss2 := &Stu2{}
	err = mapstructure.Decode(ss1, ss2)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss2.Id)
}
