package acorn

import (
	"strings"

	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type timeline struct {
	Params []*acorntypes.TimelineParams
}

// Generate a new timeline html element
func (h HTML) NewTimeline(params []*acorntypes.TimelineParams) string {
	t := timeline{Params: params}
	return t.getTimelines()
}

func (t timeline) getTimeline(params *acorntypes.TimelineParams) string {
	timeLen := len(params.Time)
	titleLen := len(params.Title)
	contentLen := len(params.Content)

	if timeLen > 11 {
		timeLen = 11
	}
	if titleLen > 37 {
		titleLen = 37
	}
	if contentLen > 78 {
		contentLen = 78
	}

	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
        <tr>
          <td class="col" width="112">
            <table cellpadding="0" cellspacing="0" role="presentation" width="100%">
              <tr>
                <td style="font-size: 14px; padding: 32px 8px 0 8px;">` + params.Time[0:timeLen] + `</td>
              </tr>
            </table>
          </td>
          <td class="hide-sm" align="center" width="48" style="vertical-align: top;">
            <img src="` + IMAGE_TIMELINE + `" alt="" width="35">
          </td>
          <td class="col" width="392">
            <table cellpadding="0" cellspacing="0" role="presentation" width="100%">
              <tr>
                <td class="py-sm-8" style="padding: 32px 8px 0 8px;">
                  <h4 style="margin: 0 0 15px;">` + params.Title[0:titleLen] + `</h4>
                  <p style="margin: 0;">` + params.Content[0:contentLen] + `</p>
                </td>
              </tr>
            </table>
          </td>
        </tr>
    </table>
	`
}

func (t timeline) getTimelines() string {
	var tLine []string
	if len(t.Params) > 0 {
		for _, params := range t.Params {
			timeLine := t.getTimeline(params)
			tLine = append(tLine, timeLine)
		}
	}
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td style="padding: 32px 24px;">
				` + strings.Join(tLine, " ") + `
			</td>
		</tr>
	</table>
	`
}
