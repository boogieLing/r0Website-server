// Package models
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 时间模型
 * @File:  time
 * @Version: 1.0.0
 * @Date: 2022/7/5 19:09
 */
package models

type TimeSort struct {
	SortDirection int64 `json:"sort_direction"`
	SortFlag      bool  `json:"sort_flag"`
}
