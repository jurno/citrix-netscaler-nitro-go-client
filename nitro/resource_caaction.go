package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Caaction struct {
	Name         string `json:"name"`
	Accumressize int    `json:"accumressize,string,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Lbvserver    string `json:"lbvserver,omitempty"`
	Type         string `json:"type,omitempty"`
}

type CaactionKey struct {
	Name string `json:"name"`
}

type CaactionUnset struct {
	Name         string `json:"name"`
	Accumressize bool   `json:"accumressize,string,omitempty"`
	Lbvserver    bool   `json:"lbvserver,string,omitempty"`
	Comment      bool   `json:"comment,string,omitempty"`
	Type         bool   `json:"type,string,omitempty"`
}

type update_caaction struct {
	Name         string `json:"name"`
	Accumressize int    `json:"accumressize,string,omitempty"`
	Lbvserver    string `json:"lbvserver,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Type         string `json:"type,omitempty"`
}

type rename_caaction struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_caaction_payload struct {
	Resource Caaction `json:"caaction"`
}

type rename_caaction_payload struct {
	Rename rename_caaction `json:"caaction"`
}

type unset_caaction_payload struct {
	Unset CaactionUnset `json:"caaction"`
}

type update_caaction_payload struct {
	Update update_caaction `json:"caaction"`
}

type get_caaction_result struct {
	Results []Caaction `json:"caaction"`
}

type count_caaction_result struct {
	Results []Count `json:"caaction"`
}

func caaction_key_to_id_args(key CaactionKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddCaaction(resource Caaction) error {
	payload := add_caaction_payload{
		resource,
	}

	return c.post("caaction", "", nil, payload)
}

func (c *NitroClient) RenameCaaction(name string, newName string) error {
	payload := rename_caaction_payload{
		rename_caaction{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("caaction", "", qs, payload)
}

func (c *NitroClient) CountCaaction() (int, error) {
	var results count_caaction_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("caaction", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsCaaction(key CaactionKey) (bool, error) {
	var results count_caaction_result

	id, qs := caaction_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("caaction", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListCaaction() ([]Caaction, error) {
	var results get_caaction_result

	if err := c.get("caaction", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetCaaction(key CaactionKey) (*Caaction, error) {
	var results get_caaction_result

	id, qs := caaction_key_to_id_args(key)

	if err := c.get("caaction", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one caaction element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("caaction element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteCaaction(key CaactionKey) error {
	id, qs := caaction_key_to_id_args(key)

	return c.delete("caaction", id, qs)
}

func (c *NitroClient) UnsetCaaction(unset CaactionUnset) error {
	payload := unset_caaction_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("caaction", "", qs, payload)
}

func (c *NitroClient) UpdateCaaction(resource Caaction) error {
	payload := update_caaction_payload{
		update_caaction{
			resource.Name,
			resource.Accumressize,
			resource.Lbvserver,
			resource.Comment,
			resource.Type,
		},
	}

	return c.put("caaction", "", nil, payload)
}
