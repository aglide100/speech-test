package txt

import (
	"errors"
	"strings"

	"github.com/jdkato/prose/v2"
)

func DivideText(text string) ([]string, error) {
	doc, _ := prose.NewDocument(text)

	var texts []string
	sents := doc.Sentences()

	for _, sent := range sents {
		if len(sent.Text) > 255 {
			ok := true
			sentences := strings.Split(sent.Text, ",")
			for _, val := range sentences {
				if len(val) > 255 {
					ok = false
					break
				}
				texts = append(texts, val)
			}

			if !ok {
				return nil, errors.New("Sent too long : /" + sent.Text)
			}

			continue
		}
		texts = append(texts, sent.Text)
	}

	return texts, nil
}