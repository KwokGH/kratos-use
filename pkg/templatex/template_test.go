package templatex

import (
	"context"
	"fmt"
	"testing"
)

func TestParsePromptTemplate(t *testing.T) {
	input := &StoryInput{
		Subject:   "中国历史人物",
		Relevant:  "秦始皇",
		WordCount: 1000,
	}
	//m := make(map[string]interface{})
	//b, _ := json.Marshal(input)
	//err := json.Unmarshal(b, &m)
	//if err != nil {
	//	t.Fatal(err)
	//}
	str, err := ParsePrompt(context.Background(), StoryTemplate, input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(str)
}
