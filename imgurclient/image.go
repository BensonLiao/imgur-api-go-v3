package imgurclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// ImageBase const, imgur's image related api endpoint
const ImageBase = APIBase + "/image"

// AnonymousUpload func, do anonymous upload without authorize
// by access file resource path
func (cl Client) AnonymousUpload(path string) (Response, error) {
	var err error
	ir := Response{}
	req, err := cl.newFileUploadRequest(
		ImageBase,
		nil,
		"image",
		"./test.png",
	)
	authHeader := []string{"Client-ID " + cl.ClientID}
	req.Header.Add("Authorization", strings.Join(authHeader, " "))
	response, err := cl.Do(req)
	if err != nil {
		return ir, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &ir)
	if err != nil {
		return ir, err
	}
	return ir, err
}

// AnonymousUploadByImgMessage func, do anonymous upload without authorize
// by read binary file content with format base64
func (cl Client) AnonymousUploadByImgMessage(content []byte) (Response, error) {
	var err error
	ir := Response{}
	req, err := cl.newImgContentUploadRequest(
		ImageBase,
		nil,
		content,
		"image",
	)
	authHeader := []string{"Client-ID " + cl.ClientID}
	req.Header.Add("Authorization", strings.Join(authHeader, " "))
	response, err := cl.Do(req)
	if err != nil {
		return ir, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &ir)
	if err != nil {
		return ir, err
	}
	return ir, err
}

// DeleteAnonymousUploadedImg func, delete uploaded imgur image
func (cl Client) DeleteAnonymousUploadedImg(deleteHash string) (DeleteResponse, error) {
	var err error
	ir := DeleteResponse{}
	req, err := http.NewRequest("DELETE", ImageBase+"/"+deleteHash, nil)
	authHeader := []string{"Client-ID " + cl.ClientID}
	req.Header.Add("Authorization", strings.Join(authHeader, " "))
	response, err := cl.Do(req)
	if err != nil {
		return ir, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &ir)
	if err != nil {
		return ir, err
	}
	return ir, err
}
