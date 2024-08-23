package templatex

import (
	"context"
	"html/template"
	"strings"
)

const (
	StoryTemplate = `
请根据以下信息，给我讲一个故事，字数要求：{{.WordCount}}：
主题：{{.Subject}}；
相关内容：{{.Relevant}}`
)

type StoryInput struct {
	Subject   string
	Relevant  string
	WordCount int32
}

func ParsePrompt(ctx context.Context, templateString string, data interface{}) (string, error) {
	contentSb := new(strings.Builder)
	tm, err := template.New("").Parse(templateString)
	tm = template.Must(tm, err)

	if err = tm.Execute(contentSb, data); err != nil {
		return "", err
	}

	templateString = contentSb.String()

	limit := 5
	for strings.Contains(templateString, "{{.") && limit > 0 {
		limit--
		tm, err = template.New("").Parse(templateString)
		tm = template.Must(tm, err)

		contentSb = new(strings.Builder)
		if err = tm.Execute(contentSb, data); err != nil {
			return "", err
		}
		templateString = contentSb.String()
	}

	return templateString, nil
}
