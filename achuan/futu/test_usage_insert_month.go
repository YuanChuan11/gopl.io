package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock             sync.Mutex
	usageTableSuffix = ""
)

func main() {
	usage1 := Usage{
		UsageDate: time.Now(),
	}
	err := UpsertUsagesWithCreateTable(usage1)
	if err != nil {
		fmt.Println(err)
	}

	usage2 := Usage{
		UsageDate: time.Now().AddDate(0, -1, 0),
	}
	err = UpsertUsagesWithCreateTable(usage2)
	if err != nil {
		fmt.Println(err)
	}
	err = UpsertUsagesWithCreateTable(usage2)
	if err != nil {
		fmt.Println(err)
	}
}

type Usage struct {
	UsageDate time.Time `gorm:"column:usage_date;type:date;not null;uniqueIndex:uk_combined,priority:2;comment:用量日期" json:"usage_date"` // 用量日期
}

func (usage *Usage) TableName() string {
	if usageTableSuffix != "" {
		return "usage" + "_" + usageTableSuffix
	}
	return "usage"
}

func SetUsageTableSuffix(suffix string) {
	usageTableSuffix = suffix
}

func GetUsageTableSuffix() string {
	return usageTableSuffix
}

// CreateUsageTable month type: 200601
func CreateUsageTable(month string) error {
	lock.Lock()
	defer lock.Unlock()
	oldSuffix := GetUsageTableSuffix()
	if month == oldSuffix {
		fmt.Printf("table usage_%v already init", month)
		fmt.Println()
		return nil
	}
	SetUsageTableSuffix(month)
	return nil
}

func upsertUsages(usage Usage) error {
	batchSize := 1
	total := 5
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}
		tableName := usage.TableName()
		fmt.Printf(usage.UsageDate.String() + ":" + tableName)
		fmt.Println()
	}
	return nil
}

func UpsertUsagesWithCreateTable(usage Usage) error {
	month := usage.UsageDate.Format("200601")
	err := CreateUsageTable(month)
	if err != nil {
		return err
	}

	return upsertUsages(usage)
}
