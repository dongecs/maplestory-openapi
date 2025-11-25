package kms

// Notice list/detail payloads.
type NoticeList struct {
	Notice []NoticeListItem `json:"notice"`
}

type NoticeListItem struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	NoticeID int    `json:"notice_id"`
	Date     string `json:"date"`
}

type NoticeDetail struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Contents string `json:"contents"`
	Date     string `json:"date"`
}

type UpdateNoticeList = NoticeList
type UpdateNoticeDetail = NoticeDetail
type EventNoticeList = NoticeList
type EventNoticeDetail = NoticeDetail
type CashshopNoticeList = NoticeList
type CashshopNoticeDetail = NoticeDetail
