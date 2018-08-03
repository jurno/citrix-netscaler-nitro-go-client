package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Videooptimizationpolicy struct {
	Name        string `json:"name"`
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}

type VideooptimizationpolicyKey struct {
	Name string `json:"name"`
}

type VideooptimizationpolicyUnset struct {
	Name        string `json:"name"`
	Rule        bool   `json:"rule,string,omitempty"`
	Action      bool   `json:"action,string,omitempty"`
	Undefaction bool   `json:"undefaction,string,omitempty"`
	Comment     bool   `json:"comment,string,omitempty"`
	Logaction   bool   `json:"logaction,string,omitempty"`
}

type update_videooptimizationpolicy struct {
	Name        string `json:"name"`
	Rule        string `json:"rule,omitempty"`
	Action      string `json:"action,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
}

type rename_videooptimizationpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_videooptimizationpolicy_payload struct {
	Resource Videooptimizationpolicy `json:"videooptimizationpolicy"`
}

type rename_videooptimizationpolicy_payload struct {
	Rename rename_videooptimizationpolicy `json:"videooptimizationpolicy"`
}

type unset_videooptimizationpolicy_payload struct {
	Unset VideooptimizationpolicyUnset `json:"videooptimizationpolicy"`
}

type update_videooptimizationpolicy_payload struct {
	Update update_videooptimizationpolicy `json:"videooptimizationpolicy"`
}

type get_videooptimizationpolicy_result struct {
	Results []Videooptimizationpolicy `json:"videooptimizationpolicy"`
}

type count_videooptimizationpolicy_result struct {
	Results []Count `json:"videooptimizationpolicy"`
}

func videooptimizationpolicy_key_to_id_args(key VideooptimizationpolicyKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddVideooptimizationpolicy(resource Videooptimizationpolicy) error {
	payload := add_videooptimizationpolicy_payload{
		resource,
	}

	return c.post("videooptimizationpolicy", "", nil, payload)
}

func (c *NitroClient) RenameVideooptimizationpolicy(name string, newName string) error {
	payload := rename_videooptimizationpolicy_payload{
		rename_videooptimizationpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("videooptimizationpolicy", "", qs, payload)
}

func (c *NitroClient) CountVideooptimizationpolicy() (int, error) {
	var results count_videooptimizationpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("videooptimizationpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsVideooptimizationpolicy(key VideooptimizationpolicyKey) (bool, error) {
	var results count_videooptimizationpolicy_result

	id, qs := videooptimizationpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("videooptimizationpolicy", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListVideooptimizationpolicy() ([]Videooptimizationpolicy, error) {
	var results get_videooptimizationpolicy_result

	if err := c.get("videooptimizationpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetVideooptimizationpolicy(key VideooptimizationpolicyKey) (*Videooptimizationpolicy, error) {
	var results get_videooptimizationpolicy_result

	id, qs := videooptimizationpolicy_key_to_id_args(key)

	if err := c.get("videooptimizationpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one videooptimizationpolicy element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("videooptimizationpolicy element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteVideooptimizationpolicy(key VideooptimizationpolicyKey) error {
	id, qs := videooptimizationpolicy_key_to_id_args(key)

	return c.delete("videooptimizationpolicy", id, qs)
}

func (c *NitroClient) UnsetVideooptimizationpolicy(unset VideooptimizationpolicyUnset) error {
	payload := unset_videooptimizationpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("videooptimizationpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateVideooptimizationpolicy(resource Videooptimizationpolicy) error {
	payload := update_videooptimizationpolicy_payload{
		update_videooptimizationpolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
			resource.Undefaction,
			resource.Comment,
			resource.Logaction,
		},
	}

	return c.put("videooptimizationpolicy", "", nil, payload)
}
