package templ

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFormatting(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "there are two line breaks after the package statement",
			input: ` // first line removed to make indentation clear in Go code
{% package test %}
{% import "strings" %}`,
			expected: `// first line removed to make indentation clear in Go code
{% package test %}

{% import "strings" %}

`,
		},
		{
			name: "import statements are all placed next to each other, and there are two lines after",
			input: ` // first line removed to make indentation clear in Go code
{% package test %}
{% import "strings" %}

{% import "net/url" %}


{% import "rand" %}
`,
			expected: `// first line removed to make indentation clear in Go code
{% package test %}

{% import "strings" %}
{% import "net/url" %}
{% import "rand" %}

`,
		},
		{
			name: "import statements don't have whitespace before or after them",
			input: ` // first line removed to make indentation clear in Go code
{% package test %}

  {% import "net/url" %}  
`,
			expected: `// first line removed to make indentation clear in Go code
{% package test %}

{% import "net/url" %}

`,
		},
		{
			name: "void elements are converted to self-closing elements",
			input: ` // first line removed to make indentation clear in Go code
{% package test %}

{% templ input(value, validation string) %}
<area></area>
<base></base>
<br></br>
<col></col>
<command></command>
<embed></embed>
<hr></hr>
<img></img>
<input></input>
<keygen></keygen>
<link></link>
<meta></meta>
<param></param>
<source></source>
<track></track>
<wbr></wbr>

{% endtempl %}
`,
			expected: `// first line removed to make indentation clear in Go code
{% package test %}

{% templ input(value, validation string) %}
	<area/>
	<base/>
	<br/>
	<col/>
	<command/>
	<embed/>
	<hr/>
	<img/>
	<input/>
	<keygen/>
	<link/>
	<meta/>
	<param/>
	<source/>
	<track/>
	<wbr/>
{% endtempl %}

`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// Remove the first line of the test data.
			input := strings.SplitN(tt.input, "\n", 2)[1]
			expected := strings.SplitN(tt.expected, "\n", 2)[1]

			// Execute the test.
			template, err := ParseString(input)
			if err != nil {
				t.Fatalf("failed to parse template: %v", err)
			}
			w := new(strings.Builder)
			err = template.Write(w)
			if err != nil {
				t.Fatalf("failed to write template: %v", err)
			}
			if diff := cmp.Diff(expected, w.String()); diff != "" {
				t.Error(diff)
			}
		})
	}
}
