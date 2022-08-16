/**
 * @Author: dexukong
 * @Description:
 * @File:  test_case
 * @Version: 1.0.0
 * @Date: 2022/08/09 16:10
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

type TestCase struct {
	ID              uint           `gorm:"primarykey;" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	Identity        string         `gorm:"column:identity;type:varchar(255);" json:"identity"`
	ProblemIdentity string         `gorm:"column:problem_identity;type:varchar(255);" json:"problem_identity"`
	Input           string         `gorm:"column:input;type:text;" json:"input"`
	Output          string         `gorm:"column:output;type:text;" json:"output"`
}

func (table *TestCase) TableName() string {
	return "test_case"
}
