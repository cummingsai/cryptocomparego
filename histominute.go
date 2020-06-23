package cryptocomparego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cummingsai/cryptocomparego/context"
	"github.com/pkg/errors"
)

const (
	histominuteBasePath = "data/histo/minute/daily"
)

// Get the history kline data of any cryptocurrency in any other currency that you need.
type HistominuteService interface {
	Get(context.Context, *HistominuteRequest) (*HistominuteResponse, *Response, error)
}

type HistominuteServiceOp struct {
	client *Client
}

var _ HistodayService = &HistodayServiceOp{}

type HistominuteResponse struct {
	Response          string         `json:"Response"`
	Message           string         `json:"Message"` // Error Message
	Type              int            `json:"Type"`
	Aggregated        bool           `json:"Aggregated"`
	Data              []Histominute  `json:"Data"`
	TimeTo            int64          `json:"TimeTo"`
	TimeFrom          int64          `json:"TimeFrom"`
	FirstValueInArray bool           `json:"FirstValueInArray"`
	ConversionType    conversionType `json:"ConversionType"`
}

type Histominute struct {
	Time       int64   `json:"time"`
	Close      float64 `json:"close"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Open       float64 `json:"open"`
	VolumeFrom float64 `json:"volumefrom"`
	VolumeTo   float64 `json:"volumeto"`
}

type HistominuteRequest struct {
	Fsym string
	Tsym string
	E    string
	Date string
}

func NewHistominuteRequest(fsym string, tsym string, date string) *HistominuteRequest {
	pr := HistominuteRequest{Fsym: fsym, Tsym: tsym}
	pr.E = "CCCAGG"
	pr.Date = date
	return &pr
}

func (hr *HistominuteRequest) FormattedQueryString(baseUrl string) string {
	values := url.Values{}

	if len(hr.Fsym) > 0 {
		values.Add("fsym", hr.Fsym)
	}

	if len(hr.Tsym) > 0 {
		values.Add("tsym", hr.Tsym)
	}

	if len(hr.E) > 0 {
		values.Add("e", hr.E)
	}
	if len(hr.Date) > 0 {
		values.Add("date", hr.Date)
	}

	return fmt.Sprintf("%s?%s", baseUrl, values.Encode())
}

func (s *HistominuteServiceOp) Get(ctx context.Context, histominuteRequest *HistominuteRequest) (*HistominuteResponse, *Response, error) {

	path := histodyBasePath

	if histominuteRequest != nil {
		path = histominuteRequest.FormattedQueryString(histominuteBasePath)
	}
	reqUrl := fmt.Sprintf("%s%s", s.client.MinURL.String(), path)
	fmt.Println(reqUrl + "&api_key=4a0867ab22e8806d04f9bf19a88c658d26d1fb7d4753a28536299a452f96f441")
	resp, err := http.Get(reqUrl + "&api_key=" + s.client.ApiKey)
	res := Response{}
	res.Response = resp
	if err != nil {
		return nil, &res, err
	}
	defer func() { resp.Body.Close() }()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &res, err
	}
	if len(buf) <= 0 {
		return nil, &res, errors.New("Empty response")
	}

	hr := HistominuteResponse{}
	err = json.Unmarshal(buf, &hr)
	if err != nil {
		return nil, &res, errors.Wrap(err, fmt.Sprintf("JSON Unmarshal error, raw string is '%s'", string(buf)))
	}
	if hr.Response == "Error" {
		return nil, &res, errors.New(hr.Message)
	}

	return &hr, &res, nil
}
