package utils

import (
	"time"
	"html/template"
)

func UnsafeHtml(s string) template.HTML {
	return template.HTML(s)

}

func StripSummaryTags(s string) string {
	return RemoveAllTags(s)
}

func DisplayDateString(s time.Time) string {
	return DisplayDateWithTime(s)
}

func DisplayDateV2(s int32) string {
	return DisplayDate(int64(s))
}

