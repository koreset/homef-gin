package utils

import (
	"time"
	"html/template"
	"github.com/koreset/homefnew/app/utils"
)

func UnsafeHtml(s string) template.HTML {
	return template.HTML(s)

}

func StripSummaryTags(s string) string {
	return utils.RemoveAllTags(s)
}

func DisplayDateString(s time.Time) string {
	return utils.DisplayDateWithTime(s)
}

func DisplayDateV2(s int32) string {
	return utils.DisplayDate(int64(s))
}
