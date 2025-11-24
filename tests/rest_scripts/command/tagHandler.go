package command

type TagHandler interface {
	SetTags(tags *Tagging)
	GetTags() *Tagging
	SetPayload(payload string)
}
