/**
*支持aes128、aes256加解密
*参考文档
*https://blog.csdn.net/a_lzq/article/details/108204967?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromBaidu-19.control&dist_request_id=1328626.22722.16154492627822725&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromBaidu-19.control
*
 */
package common

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/shopspring/decimal"
)

// 全局 transport
var globalTransport *http.Transport

func init() {
	globalTransport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func HttpPost(api string, data map[string]interface{}) (code int, body string, err error) {
	req := make(url.Values)
	for key, item := range data {
		req[key] = []string{fmt.Sprintf("%v", item)}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//把post表单发送给目标服务器
	res, err := client.PostForm(api, req)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()
	bytess, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}
	return res.StatusCode, string(bytess), nil
}

func HttpGet(api string) (code int, body string, err error) {
	res, err := http.Get(api)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()
	bytess, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}
	return res.StatusCode, string(bytess), nil
}

func Md5Encode(str string) string {
	data := []byte(str)
	h := md5.New()
	h.Write(data)
	output := h.Sum(nil)
	return fmt.Sprintf("%x", output)
}

func FormatCollectTaskData(task map[string]interface{}) map[string]interface{} {
	add_time := int64(0)
	if _, ok := task["add_time"].(float64); ok {
		add_time = int64(task["add_time"].(float64))
	} else {
		add_time = task["add_time"].(int64)
	}
	settlement_id := ""
	switch item := task["settlement_id"].(type) {
	case string:
		_settlement_id, _ := decimal.NewFromString(item)
		__settlement_id, _ := _settlement_id.Float64()
		settlement_id = decimal.NewFromFloat(__settlement_id).String()
	case int64:
		settlement_id = fmt.Sprintf("%d", item)
	case float64:
		settlement_id = decimal.NewFromFloat(item).String()
	}

	shop_id := ""
	switch item := task["shop_id"].(type) {
	case string:
		_shop_id, _ := decimal.NewFromString(item)
		__shop_id, _ := _shop_id.Float64()
		shop_id = decimal.NewFromFloat(__shop_id).String()
	case int64:
		shop_id = fmt.Sprintf("%d", item)
	case float64:
		shop_id = decimal.NewFromFloat(item).String()
	}

	total := ""
	switch item := task["total"].(type) {
	case string:
		_total, _ := decimal.NewFromString(item)
		__total, _ := _total.Float64()
		total = decimal.NewFromFloat(__total).String()
	case float64:
		total = decimal.NewFromFloat(item).String()
	case int64:
		total = fmt.Sprintf("%d", item)
	}

	typeStr := ""
	switch item := task["type"].(type) {
	case string:
		_type, _ := decimal.NewFromString(item)
		__type, _ := _type.Float64()
		typeStr = decimal.NewFromFloat(__type).String()
	case float64:
		typeStr = decimal.NewFromFloat(item).String()
	case int64:
		typeStr = fmt.Sprintf("%d", item)
	}

	data := map[string]interface{}{
		"settlement_id": settlement_id,
		"shop_id":       shop_id,
		"coin":          task["coin"].(string),
		"contract":      task["contract"].(string),
		"address":       task["address"].(string),
		"to_address":    task["to_address"].(string),
		"total":         total,
		"type":          typeStr,
		"add_time":      add_time,
		"salt":          task["salt"].(string),
	}
	return data
}

func FormatRefundTaskData(task map[string]interface{}) map[string]interface{} {
	add_time := int64(0)
	if _, ok := task["add_time"].(float64); ok {
		add_time = int64(task["add_time"].(float64))
	} else {
		add_time = task["add_time"].(int64)
	}
	order_id := ""
	switch item := task["settlement_id"].(type) {
	case string:
		_order_id, _ := decimal.NewFromString(item)
		__order_id, _ := _order_id.Float64()
		order_id = decimal.NewFromFloat(__order_id).String()
	case int64:
		order_id = fmt.Sprintf("%d", item)
	case float64:
		order_id = decimal.NewFromFloat(item).String()
	}

	shop_id := ""
	switch item := task["shop_id"].(type) {
	case string:
		_shop_id, _ := decimal.NewFromString(item)
		__shop_id, _ := _shop_id.Float64()
		shop_id = decimal.NewFromFloat(__shop_id).String()
	case int64:
		shop_id = fmt.Sprintf("%d", item)
	case float64:
		shop_id = decimal.NewFromFloat(item).String()
	}

	total := ""
	switch item := task["total"].(type) {
	case string:
		_total, _ := decimal.NewFromString(item)
		__total, _ := _total.Float64()
		total = decimal.NewFromFloat(__total).String()
	case float64:
		total = decimal.NewFromFloat(item).String()
	case int64:
		total = fmt.Sprintf("%d", item)
	}

	typeStr := ""
	fmt.Printf("task type:  %v \n", task["type"])
	switch item := task["type"].(type) {
	case string:
		_type, _ := decimal.NewFromString(item)
		__type, _ := _type.Float64()
		typeStr = decimal.NewFromFloat(__type).String()
	case float64:
		typeStr = decimal.NewFromFloat(item).String()
	case int64:
		typeStr = fmt.Sprintf("%d", item)
	}

	data := map[string]interface{}{
		"order_id":     order_id,
		"shop_id":      shop_id,
		"coin":         task["coin"].(string),
		"contract":     task["contract"].(string),
		"from_address": task["from_address"].(string),
		"to_address":   task["to_address"].(string),
		"total":        total,
		"type":         typeStr,
		"add_time":     add_time,
		"salt":         task["salt"].(string),
	}
	return data
}

func Request(params map[string]interface{}, headers map[string]string, url string, res interface{}) error {
	ctx := context.Background()

	var err error
	req := &http.Request{}
	if len(params) > 0 {
		postData, err := json.Marshal(params)
		if err != nil {
			return err
		}
		req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(postData))
	} else {
		req, err = http.NewRequestWithContext(ctx, "GET", url, nil)
	}

	if err != nil {
		return err
	}

	// headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	// http client and send request
	httpclient := &http.Client{
		Transport: globalTransport,
		Timeout:   10 * time.Second,
	}
	resp, err := httpclient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// parse body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(body) != 0 {
		if err := json.Unmarshal(body, &res); err != nil {
			return err
		}
	}

	// return result
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return fmt.Errorf("get status code: %d", resp.StatusCode)
	}
	return nil
}

func Upload(url string, headers map[string]string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	httpclient := &http.Client{
		Transport: globalTransport,
		Timeout:   10 * time.Second,
	}
	// Submit the request
	res, err := httpclient.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK || res.StatusCode != http.StatusCreated {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return
}
