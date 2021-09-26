package opensea

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type NFTResponse struct {
	Assets []*Asset `json:"assets"`
}

type Asset struct {
	TokenID           *string        `json:"token_id"`
	Name              *string        `json:"name"`
	Description       *string        `json:"description"`
	ImageURL          *string        `json:"image_url"`
	ImagePreviewURL   *string        `json:"image_preview_url"`
	ImageThumbnailURL *string        `json:"image_thumbnail_url"`
	OwnerData         *Owner         `json:"owner"`
	AssetContractData *AssetContract `json:"asset_contract"`
}

type Owner struct {
	UserData *User   `json:"user"`
}

type User struct {
	Username *string `json:"username"`
}

type AssetContract struct {
	Address     *string `json:"address"`
	Name        *string `json:"name"`
	Symbol      *string `json:"symbol"`
	Description *string `json:"description"`
}

type Client struct {
	BaseURL    string
	httpClient *http.Client
}

func New(baseURL string, timeout time.Duration) *Client {
	return &Client{
		BaseURL: baseURL,
		httpClient: &http.Client{
			Transport: &http.Transport{
				ResponseHeaderTimeout: timeout,
			},
			Timeout: timeout,
		},
	}
}

func (c *Client) GetNFTInfo(owner string) (*NFTResponse, error) {
	resBody := &NFTResponse{}

	query := url.Values{}
	query.Add("owner", owner)
	query.Add("limit", "50")

	err := c.makeGetRequest(query.Encode(), "/api/v1/assets", resBody)
	return resBody, err
}

// TODO: add api key parameter
func (c *Client) makeGetRequest(queryString string, requestURL string, result interface{}) error {
	req, err := http.NewRequest("GET", c.BaseURL+requestURL, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = queryString
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)
	return err
}
