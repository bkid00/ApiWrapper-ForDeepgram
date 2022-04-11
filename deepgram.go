package deepgram

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func makeRequest(url string, payload interface{}) ([]byte, error) {
	reqJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	request.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}


	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}


func parseResponse(response []byte, t interface{}) error {
	respErr := new(ResponseError)
	err := json.Unmarshal(response, respErr)
	if err != nil {
		return err
	}
	if respErr.Error != "" {
		return errors.New(respErr.Error)
	}
	err = json.Unmarshal(response, t)
	if err != nil {
		return err
	}
	return nil
}


func (dg *Deepgram) CheckBalance() (*CheckBalanceResponse, error) {
	req := checkBalanceRequest{
		Action: "get_balance",
		UserId: dg.ApiKey,
	}
	resp, err := makeRequest(dg.Host(), req)
	if err != nil {
		return nil, err
	}
	result := new(CheckBalanceResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dg *Deepgram) CheckStatus(obj string) (*CheckStatusResponse, error) {
	req := getObjectInfoRequest{
		Action:    "get_object_status",
		UserId:    dg.ApiKey,
		ContentId: obj,
	}
	resp, err := makeRequest(dg.Host(), req)
	if err != nil {
		return nil, err
	}
	result := new(CheckStatusResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dg *Deepgram) Upload(mediaUrl string, tags []string) (*UploadResponse, error) {
	req := uploadRequest{
		Action:  "index_content",
		UserId:  dg.ApiKey,
		DataUrl: mediaUrl,
		Tags:    tags,
	}
	resp, err := makeRequest(dg.Host(), req)
	if err != nil {
		return nil, err
	}
	result := new(UploadResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func (dg *Deepgram) UploadList(mediaUrls []string) (*UploadListResponse, error) {
	req := uploadListRequest{
		Action:  "index_content_list",
		UserId:  dg.ApiKey,
		DataUrl: mediaUrls,
	}
	resp, err := makeRequest(dg.Host(), req)
	if err != nil {
		return nil, err
	}
	result := new(UploadListResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dg *Deepgram) Query(obj, query string, options *QueryRequestParameters) (*QueryResponse, error) {
	if options == nil {
		options = new(QueryRequestParameters)
	}
	if options.Snippet == nil {
		snippet := true
		options.Snippet = &snippet
	}
	if options.Nmax == nil {
		var nmax int32 = 10
		options.Nmax = &nmax
	}
	if options.Pmin == nil {
		var pmin float32 = 0.55
		options.Pmin = &pmin
	}
	if options.Sort == nil {
		sort := "time"
		options.Sort = &sort
	}
	req := querySearchRequest{
		Action:    "object_search",
		UserId:    dg.ApiKey,
		ContentId: obj,
		Query:     query,
		Sort:      *options.Sort,
		Snippet:   *options.Snippet,
		Filter: FilterParameters{
			Nmax: *options.Nmax,
			Pmin: *options.Pmin,
		},
	}
	resp, err := makeRequest(dg.Host(), req)
	if err != nil {
		return nil, err
	}
	result := new(QueryResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func (dg *Deepgram) GroupSearch(query, tag string) (*GroupSearchResponse, error) {
	req := groupSearchRequest{
		Action: "group_search",
		UserId: dg.ApiKey,
		Tag:    tag,
		Query:  query,
	}
	resp, err := makeRequest(dg.GroupSearchHost(), req)
	if err != nil {
		return nil, err
	}
	result := new(GroupSearchResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dg *Deepgram) ParallelSearch(query string, options *ParallelSearchParameters) (*ParallelSearchResponse, error) {
	if options == nil {
		options = new(ParallelSearchParameters)
	}
	if options.Snippet == nil {
		snippet := true
		options.Snippet = &snippet
	}
	if options.Tag == nil {
		tag := ""
		options.Tag = &tag
	}
	if options.GroupNmax == nil {
		var nmax int32 = 10
		options.GroupNmax = &nmax
	}
	if options.ObjectNmax == nil {
		var nmax int32 = 10
		options.ObjectNmax = &nmax
	}
	if options.ObjectPmin == nil {
		var pmin float32 = 0.55
		options.ObjectPmin = &pmin
	}
	if options.Sort == nil {
		sort := "time"
		options.Sort = &sort
	}
	req := parallelSearchRequest{
		Action:  "parallel_search",
		UserId:  dg.ApiKey,
		Query:   query,
		Tag:     *options.Tag,
		Sort:    *options.Sort,
		Snippet: *options.Snippet,
		GroupFilter: GroupFilterParameters{
			Nmax: *options.GroupNmax,
		},
		ObjectFilter: FilterParameters{
			Nmax: *options.ObjectNmax,
			Pmin: *options.ObjectPmin,
		},
	}
	resp, err := makeRequest(dg.GroupSearchHost(), req)
	if err != nil {
		return nil, err
	}
	result := new(ParallelSearchResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func (dg *Deepgram) Tag(obj, tag string) (*TagResponse, error) {
	req := tagRequest{
		Action:    "tag_object",
		UserId:    dg.ApiKey,
		ContentId: obj,
		Tag:       tag,
	}
	resp, err := makeRequest(dg.Host(), req)
	if err != nil {
		return nil, err
	}
	result := new(TagResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func (dg *Deepgram) GetTags(obj string) (*GetTagsResponse, error) {
	req := getObjectInfoRequest{
		Action:    "get_object_tags",
		UserId:    dg.ApiKey,
		ContentId: obj,
	}
	resp, err := makeRequest(dg.Host(), req)
	if err != nil {
		return nil, err
	}
	result := new(GetTagsResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func (dg *Deepgram) Transcript(obj string) (*TranscriptResponse, error) {
	req := getObjectInfoRequest{
		Action:    "get_object_transcript",
		UserId:    dg.ApiKey,
		ContentId: obj,
	}
	resp, err := makeRequest(dg.Host(), req)
	if err != nil {
		return nil, err
	}
	result := new(TranscriptResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
