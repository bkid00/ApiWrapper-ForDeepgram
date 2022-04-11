package deepgram


type Deepgram struct {
	ApiKey string
}

func (dg *Deepgram) Host() string {
	return "http://api.deepgram.com"
}

func (dg *Deepgram) GroupSearchHost() string {
	return "http://groupsearch.api.deepgram.com"
}


type ResponseError struct {
	Error string `json:"error"`
}


type getObjectInfoRequest struct {
	Action    string `json:"action"`
	UserId    string `json:"userID"`
	ContentId string `json:"contentID"`
}

type FilterParameters struct {
	Nmax int32   `json:"Nmax"`
	Pmin float32 `json:"Pmin"`
}

type GroupFilterParameters struct {
	Nmax int32 `json:"Nmax"`
}


type checkBalanceRequest struct {
	Action string `json:"action"`
	UserId string `json:"userID"`
}


type CheckBalanceResponse struct {
	Balance float32 `json:"balance"`
	UserId  string  `json:"userID"`
}


type CheckStatusResponse struct {
	Status string `json:"status"`
}

type uploadRequest struct {
	Action  string   `json:"action"`
	UserId  string   `json:"userID"`
	DataUrl string   `json:"data_url"`
	Tags    []string `json:"tags"`
}


type UploadResponse struct {
	ContentId string `json:"contentID"`
}


type uploadListRequest struct {
	Action  string   `json:"action"`
	UserId  string   `json:"userID"`
	DataUrl []string `json:"data_url"`
}


type UploadListResponse struct {
	ContentId []string `json:"contentID"`
}


type tagRequest struct {
	Action    string `json:"action"`
	UserId    string `json:"userID"`
	ContentId string `json:"contentID"`
	Tag       string `json:"tag"`
}

type TagResponse struct {
	Result string `json:"result"`
}


type GetTagsResponse struct {
	ContentId string   `json:"contentID"`
	Tags      []string `json:"tags"`
}

type TranscriptResponse struct {
	ContentId           string    `json:"contentID"`
	Paragraphs          []string  `json:"paragraphs"`
	ParagraphStartTimes []float32 `json:"paragraphStartTimes"`
}

type QueryRequestParameters struct {
	Snippet *bool
	Nmax    *int32
	Pmin    *float32
	Sort    *string
}

type querySearchRequest struct {
	Action    string           `json:"action"`
	UserId    string           `json:"userID"`
	ContentId string           `json:"contentID"`
	Query     string           `json:"query"`
	Sort      string           `json:"sort"`
	Snippet   bool             `json:"snippet"`
	Filter    FilterParameters `json:"filter"`
}


type QueryResponse struct {
	Snippet   []string  `json:"snippet"`
	StartTime []float32 `json:"startTime"`
	EndTime   []float32 `json:"endTime"`
	P         []float32 `json:"P"`
	N         []int32   `json:"N"`
}


type groupSearchRequest struct {
	Action string `json:"action"`
	UserId string `json:"userID"`
	Tag    string `json:"tag"`
	Query  string `json:"query"`
}


type GroupSearchResponse struct {
	ContentId []string  `json:"contentID"`
	P         []float32 `json:"P"`
	N         []int32   `json:"N"`
}

type ParallelSearchParameters struct {
	Snippet    *bool
	Tag        *string
	GroupNmax  *int32
	ObjectNmax *int32
	ObjectPmin *float32
	Sort       *string
}


type parallelSearchRequest struct {
	Action       string                `json:"action"`
	UserId       string                `json:"userID"`
	Query        string                `json:"query"`
	Tag          string                `json:"tag"`
	GroupFilter  GroupFilterParameters `json:"group_filter"`
	ObjectFilter FilterParameters      `json:"object_filter"`
	Snippet      bool                  `json:"snippet"`
	Sort         string                `json:"sort"`
}


type ObjectResult struct {
	ContentId string    `json:"contentID"`
	Snippet   []string  `json:"snippet"`
	StartTime []float32 `json:"startTime"`
	EndTime   []float32 `json:"endTime"`
	N         []int32   `json:"N"`
	P         []float32 `json:"P"`
}

type ParallelSearchResponse struct {
	ObjectResult []ObjectResult `json:"object_result"`
}
