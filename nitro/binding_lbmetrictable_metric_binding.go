package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type LbmetrictableMetricBinding struct {
	Metric      string `json:"metric,omitempty"`
	Metrictable string `json:"metrictable,omitempty"`
	Snmpoid     string `json:"snmpoid,omitempty"`
}

type LbmetrictableMetricBindingKey struct {
	Metrictable string
	Metric      string
}

type add_lbmetrictable_metric_binding_payload struct {
	Resource LbmetrictableMetricBinding `json:"lbmetrictable_metric_binding"`
}

type get_lbmetrictable_metric_binding_result struct {
	Results []LbmetrictableMetricBinding `json:"lbmetrictable_metric_binding"`
}

type count_lbmetrictable_metric_binding_result struct {
	Results []Count `json:"lbmetrictable_metric_binding"`
}

func lbmetrictable_metric_binding_key_to_id_args(key LbmetrictableMetricBindingKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "metrictable:"+key.Metrictable)
	args = append(args, "metric:"+key.Metric)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return "", qs
}

func (c *NitroClient) AddLbmetrictableMetricBinding(binding LbmetrictableMetricBinding) error {
	payload := add_lbmetrictable_metric_binding_payload{
		binding,
	}

	return c.put("lbmetrictable_metric_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountLbmetrictableMetricBinding() (int, error) {
	var results count_lbmetrictable_metric_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("lbmetrictable_metric_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountLbmetrictableMetricBinding(id string) (int, error) {
	var results count_lbmetrictable_metric_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbmetrictable_metric_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbmetrictableMetricBinding(id string) (bool, error) {
	// TODO : wrong implementation
	if count, err := c.CountLbmetrictableMetricBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListLbmetrictableMetricBinding() ([]LbmetrictableMetricBinding, error) {
	var results get_lbmetrictable_metric_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("lbmetrictable_metric_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListLbmetrictableMetricBinding(id string) ([]LbmetrictableMetricBinding, error) {
	var results get_lbmetrictable_metric_binding_result

	if err := c.get("lbmetrictable_metric_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) (*LbmetrictableMetricBinding, error) {
	var results get_lbmetrictable_metric_binding_result

	id, qs := lbmetrictable_metric_binding_key_to_id_args(key)

	if err := c.get("lbmetrictable_metric_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmetrictable_metric_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbmetrictable_metric_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbmetrictableMetricBinding(key LbmetrictableMetricBindingKey) error {
	id, qs := lbmetrictable_metric_binding_key_to_id_args(key)

	return c.delete("lbmetrictable_metric_binding", id, qs)
}
