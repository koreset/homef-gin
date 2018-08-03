package services

type MenuItem struct {
	Type  string
	Count uint
}

func GetContentCounts() []MenuItem {
	rows, err := GetDB().Table("contents").Select("count(type) as count, type").Group("type").Rows()
	if err != nil {
		panic(err)
	}
	var items []MenuItem
	for rows.Next() {
		var item MenuItem
		GetDB().ScanRows(rows, &item)
		items = append(items, item)
	}
	return items
}
