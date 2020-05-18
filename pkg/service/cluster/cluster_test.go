package cluster

import (
	"github.com/spf13/viper"
	"ko3-gin/pkg/config"
	"ko3-gin/pkg/db"
	clusterModel "ko3-gin/pkg/model/cluster"
	"ko3-gin/pkg/model/common"
	"log"
	"testing"
)

func Init() {
	config.Init()
	phase := db.InitDBPhase{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Name:     viper.GetString("db.name"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
	}
	err := phase.Init()
	if err != nil {
		log.Fatalf("can not init db,%s", err)
	}
}

func TestSave(t *testing.T) {
	Init()
	item := clusterModel.Cluster{
		BaseModel: common.BaseModel{
			Name: "test",
		},
	}
	err := Save(&item)
	if err != nil {
		t.Fatalf("can not create item,%s", err)
	}
}

func TestList(t *testing.T) {
	Init()
	items, err := List()
	if err != nil {
		t.Fatalf("can not list item,%s", err)
	}
	t.Log(items)
}

func TestPage(t *testing.T) {
	Init()
	items, total, err := Page(1, 10)
	if err != nil {
		t.Fatalf("can not page item,%s", err)
	}
	t.Log(items)
	t.Log(total)
}


func TestDelete(t *testing.T) {
	Init()
	err := Delete("test")
	if err != nil {
		t.Fatalf("can not delete item,%s", err)
	}
}

