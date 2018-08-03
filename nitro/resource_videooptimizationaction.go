package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Videooptimizationaction struct {
	Name    string `json:"name"`
	Comment string `json:"comment,omitempty"`
	Rate    int    `json:"rate,omitempty"`
	Type    string `json:"type,omitempty"`
}

type VideooptimizationactionKey struct {
	Name string `json:"name"`
}

type VideooptimizationactionUnset struct {
	Name    string `json:"name"`
	Type    bool   `json:"type,string,omitempty"`
	Rate    bool   `json:"rate,string,omitempty"`
	Comment bool   `json:"comment,string,omitempty"`
}

type update_videooptimizationaction struct {
	Name    string `json:"name"`
	Type    string `json:"type,omitempty"`
	Rate    int    `json:"rate,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type rename_videooptimizationaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_videooptimizationaction_payload struct {
	Resource Videooptimizationaction `json:"videooptimizationaction"`
}

type rename_videooptimizationaction_payload struct {
	Rename rename_videooptimizationaction `json:"videooptimizationaction"`
}

type unset_videooptimizationaction_payload struct {
	Unset VideooptimizationactionUnset `json:"videooptimizationaction"`
}

type update_videooptimizationaction_payload struct {
	Update update_videooptimizationaction `json:"videooptimizationaction"`
}

type get_videooptimizationaction_result struct {
	Results []Videooptimizationaction `json:"videooptimizationaction"`
}

type count_videooptimizationaction_result struct {
	Results []Count `json:"videooptimizationaction"`
}

func videooptimizationaction_key_to_id_args(key VideooptimizationactionKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddVideooptimizationaction(resource Videooptimizationaction) error {
	payload := add_videooptimizationaction_payload{
		resource,
	}

	return c.post("videooptimizationaction", "", nil, payload)
}

func (c *NitroClient) RenameVideooptimizationaction(name string, newName string) error {
	payload := rename_videooptimizationaction_payload{
		rename_videooptimizationaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("videooptimizationaction", "", qs, payload)
}

func (c *NitroClient) CountVideooptimizationaction() (int, error) {
	var results count_videooptimizationaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("videooptimizationaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsVideooptimizationaction(key VideooptimizationactionKey) (bool, error) {
	var results count_videooptimizationaction_result

	id, qs := videooptimizationaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("videooptimizationaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListVideooptimizationaction() ([]Videooptimizationaction, error) {
	var results get_videooptimizationaction_result

	if err := c.get("videooptimizationaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetVideooptimizationaction(key VideooptimizationactionKey) (*Videooptimizationaction, error) {
	var results get_videooptimizationaction_result

	id, qs := videooptimizationaction_key_to_id_args(key)

	if err := c.get("videooptimizationaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one videooptimizationaction element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("videooptimizationaction element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteVideooptimizationaction(key VideooptimizationactionKey) error {
	id, qs := videooptimizationaction_key_to_id_args(key)

	return c.delete("videooptimizationaction", id, qs)
}

func (c *NitroClient) UnsetVideooptimizationaction(unset VideooptimizationactionUnset) error {
	payload := unset_videooptimizationaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("videooptimizationaction", "", qs, payload)
}

func (c *NitroClient) UpdateVideooptimizationaction(resource Videooptimizationaction) error {
	payload := update_videooptimizationaction_payload{
		update_videooptimizationaction{
			resource.Name,
			resource.Type,
			resource.Rate,
			resource.Comment,
		},
	}

	return c.put("videooptimizationaction", "", nil, payload)
}
