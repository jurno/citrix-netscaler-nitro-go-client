package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbmonitor struct {
	Monitorname                    string `json:"monitorname"`
	State                          string `json:"state,omitempty"`
	Action                         string `json:"action,omitempty"`
	Alertretries                   int    `json:"alertretries,omitempty"`
	Application                    string `json:"application,omitempty"`
	Attribute                      string `json:"attribute,omitempty"`
	Basedn                         string `json:"basedn,omitempty"`
	Binddn                         string `json:"binddn,omitempty"`
	Customheaders                  string `json:"customheaders,omitempty"`
	Database                       string `json:"database,omitempty"`
	Destip                         string `json:"destip,omitempty"`
	Destport                       int    `json:"destport,omitempty"`
	Deviation                      int    `json:"deviation,string,omitempty"`
	Dispatcherip                   string `json:"dispatcherip,omitempty"`
	Dispatcherport                 int    `json:"dispatcherport,omitempty"`
	Domain                         string `json:"domain,omitempty"`
	Downtime                       int    `json:"downtime,omitempty"`
	Evalrule                       string `json:"evalrule,omitempty"`
	Failureretries                 int    `json:"failureretries,omitempty"`
	Filename                       string `json:"filename,omitempty"`
	Filter                         string `json:"filter,omitempty"`
	Firmwarerevision               int    `json:"firmwarerevision,string,omitempty"`
	Group                          string `json:"group,omitempty"`
	Hostipaddress                  string `json:"hostipaddress,omitempty"`
	Hostname                       string `json:"hostname,omitempty"`
	Httprequest                    string `json:"httprequest,omitempty"`
	Inbandsecurityid               string `json:"inbandsecurityid,omitempty"`
	Interval                       int    `json:"interval,omitempty"`
	Iptunnel                       string `json:"iptunnel,omitempty"`
	Kcdaccount                     string `json:"kcdaccount,omitempty"`
	Lasversion                     string `json:"lasversion,omitempty"`
	Logonpointname                 string `json:"logonpointname,omitempty"`
	Lrtm                           string `json:"lrtm,omitempty"`
	Maxforwards                    int    `json:"maxforwards,string,omitempty"`
	Metrictable                    string `json:"metrictable,omitempty"`
	Mssqlprotocolversion           string `json:"mssqlprotocolversion,omitempty"`
	Netprofile                     string `json:"netprofile,omitempty"`
	Oraclesid                      string `json:"oraclesid,omitempty"`
	Originhost                     string `json:"originhost,omitempty"`
	Originrealm                    string `json:"originrealm,omitempty"`
	Password                       string `json:"password,omitempty"`
	Productname                    string `json:"productname,omitempty"`
	Query                          string `json:"query,omitempty"`
	Querytype                      string `json:"querytype,omitempty"`
	Radaccountsession              string `json:"radaccountsession,omitempty"`
	Radaccounttype                 int    `json:"radaccounttype,string,omitempty"`
	Radapn                         string `json:"radapn,omitempty"`
	Radframedip                    string `json:"radframedip,omitempty"`
	Radkey                         string `json:"radkey,omitempty"`
	Radmsisdn                      string `json:"radmsisdn,omitempty"`
	Radnasid                       string `json:"radnasid,omitempty"`
	Radnasip                       string `json:"radnasip,omitempty"`
	Recv                           string `json:"recv,omitempty"`
	Resptimeout                    int    `json:"resptimeout,omitempty"`
	Resptimeoutthresh              int    `json:"resptimeoutthresh,string,omitempty"`
	Retries                        int    `json:"retries,omitempty"`
	Reverse                        string `json:"reverse,omitempty"`
	Rtsprequest                    string `json:"rtsprequest,omitempty"`
	Scriptargs                     string `json:"scriptargs,omitempty"`
	Scriptname                     string `json:"scriptname,omitempty"`
	Secondarypassword              string `json:"secondarypassword,omitempty"`
	Secure                         string `json:"secure,omitempty"`
	Send                           string `json:"send,omitempty"`
	Sipmethod                      string `json:"sipmethod,omitempty"`
	Sipreguri                      string `json:"sipreguri,omitempty"`
	Sipuri                         string `json:"sipuri,omitempty"`
	Sitepath                       string `json:"sitepath,omitempty"`
	Snmpcommunity                  string `json:"snmpcommunity,omitempty"`
	Snmpoid                        string `json:"snmpoid,omitempty"`
	Snmpthreshold                  string `json:"snmpthreshold,omitempty"`
	Snmpversion                    string `json:"snmpversion,omitempty"`
	Sqlquery                       string `json:"sqlquery,omitempty"`
	Sslprofile                     string `json:"sslprofile,omitempty"`
	Storedb                        string `json:"storedb,omitempty"`
	Storefrontacctservice          string `json:"storefrontacctservice,omitempty"`
	Storefrontcheckbackendservices string `json:"storefrontcheckbackendservices,omitempty"`
	Storename                      string `json:"storename,omitempty"`
	Successretries                 int    `json:"successretries,omitempty"`
	Tos                            string `json:"tos,omitempty"`
	Tosid                          int    `json:"tosid,string,omitempty"`
	Transparent                    string `json:"transparent,omitempty"`
	Trofscode                      int    `json:"trofscode,string,omitempty"`
	Trofsstring                    string `json:"trofsstring,omitempty"`
	Type                           string `json:"type,omitempty"`
	Units1                         string `json:"units1,omitempty"`
	Units2                         string `json:"units2,omitempty"`
	Units3                         string `json:"units3,omitempty"`
	Units4                         string `json:"units4,omitempty"`
	Username                       string `json:"username,omitempty"`
	Validatecred                   string `json:"validatecred,omitempty"`
	Vendorid                       int    `json:"vendorid,string,omitempty"`
	Vendorspecificvendorid         int    `json:"vendorspecificvendorid,string,omitempty"`
}

type LbmonitorKey struct {
	Monitorname string `json:"monitorname"`
	Type        string `json:"type,omitempty"`
}

func lbmonitor_key_to_id_args(key LbmonitorKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "type:"+key.Type)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Monitorname, qs
}

type LbmonitorUnset struct {
	Monitorname string `json:"monitorname"`
	// TODO
	// Type string `json:"type,omitempty"`
	Type                           bool `json:"type,string,omitempty"`
	Action                         bool `json:"action,string,omitempty"`
	Httprequest                    bool `json:"httprequest,string,omitempty"`
	Rtsprequest                    bool `json:"rtsprequest,string,omitempty"`
	Customheaders                  bool `json:"customheaders,string,omitempty"`
	Maxforwards                    bool `json:"maxforwards,string,omitempty"`
	Sipmethod                      bool `json:"sipmethod,string,omitempty"`
	Sipreguri                      bool `json:"sipreguri,string,omitempty"`
	Sipuri                         bool `json:"sipuri,string,omitempty"`
	Send                           bool `json:"send,string,omitempty"`
	Recv                           bool `json:"recv,string,omitempty"`
	Query                          bool `json:"query,string,omitempty"`
	Querytype                      bool `json:"querytype,string,omitempty"`
	Username                       bool `json:"username,string,omitempty"`
	Password                       bool `json:"password,string,omitempty"`
	Secondarypassword              bool `json:"secondarypassword,string,omitempty"`
	Logonpointname                 bool `json:"logonpointname,string,omitempty"`
	Lasversion                     bool `json:"lasversion,string,omitempty"`
	Radkey                         bool `json:"radkey,string,omitempty"`
	Radnasid                       bool `json:"radnasid,string,omitempty"`
	Radnasip                       bool `json:"radnasip,string,omitempty"`
	Radaccounttype                 bool `json:"radaccounttype,string,omitempty"`
	Radframedip                    bool `json:"radframedip,string,omitempty"`
	Radapn                         bool `json:"radapn,string,omitempty"`
	Radmsisdn                      bool `json:"radmsisdn,string,omitempty"`
	Radaccountsession              bool `json:"radaccountsession,string,omitempty"`
	Lrtm                           bool `json:"lrtm,string,omitempty"`
	Deviation                      bool `json:"deviation,string,omitempty"`
	Units1                         bool `json:"units1,string,omitempty"`
	Scriptname                     bool `json:"scriptname,string,omitempty"`
	Scriptargs                     bool `json:"scriptargs,string,omitempty"`
	Validatecred                   bool `json:"validatecred,string,omitempty"`
	Domain                         bool `json:"domain,string,omitempty"`
	Dispatcherip                   bool `json:"dispatcherip,string,omitempty"`
	Dispatcherport                 bool `json:"dispatcherport,string,omitempty"`
	Interval                       bool `json:"interval,string,omitempty"`
	Units3                         bool `json:"units3,string,omitempty"`
	Resptimeout                    bool `json:"resptimeout,string,omitempty"`
	Units4                         bool `json:"units4,string,omitempty"`
	Resptimeoutthresh              bool `json:"resptimeoutthresh,string,omitempty"`
	Retries                        bool `json:"retries,string,omitempty"`
	Failureretries                 bool `json:"failureretries,string,omitempty"`
	Alertretries                   bool `json:"alertretries,string,omitempty"`
	Successretries                 bool `json:"successretries,string,omitempty"`
	Downtime                       bool `json:"downtime,string,omitempty"`
	Units2                         bool `json:"units2,string,omitempty"`
	Destip                         bool `json:"destip,string,omitempty"`
	Destport                       bool `json:"destport,string,omitempty"`
	State                          bool `json:"state,string,omitempty"`
	Reverse                        bool `json:"reverse,string,omitempty"`
	Transparent                    bool `json:"transparent,string,omitempty"`
	Iptunnel                       bool `json:"iptunnel,string,omitempty"`
	Tos                            bool `json:"tos,string,omitempty"`
	Tosid                          bool `json:"tosid,string,omitempty"`
	Secure                         bool `json:"secure,string,omitempty"`
	Group                          bool `json:"group,string,omitempty"`
	Filename                       bool `json:"filename,string,omitempty"`
	Basedn                         bool `json:"basedn,string,omitempty"`
	Binddn                         bool `json:"binddn,string,omitempty"`
	Filter                         bool `json:"filter,string,omitempty"`
	Attribute                      bool `json:"attribute,string,omitempty"`
	Database                       bool `json:"database,string,omitempty"`
	Oraclesid                      bool `json:"oraclesid,string,omitempty"`
	Sqlquery                       bool `json:"sqlquery,string,omitempty"`
	Evalrule                       bool `json:"evalrule,string,omitempty"`
	Snmpoid                        bool `json:"snmpoid,string,omitempty"`
	Snmpcommunity                  bool `json:"snmpcommunity,string,omitempty"`
	Snmpthreshold                  bool `json:"snmpthreshold,string,omitempty"`
	Snmpversion                    bool `json:"snmpversion,string,omitempty"`
	Metrictable                    bool `json:"metrictable,string,omitempty"`
	Application                    bool `json:"application,string,omitempty"`
	Sitepath                       bool `json:"sitepath,string,omitempty"`
	Storename                      bool `json:"storename,string,omitempty"`
	Storefrontacctservice          bool `json:"storefrontacctservice,string,omitempty"`
	Storefrontcheckbackendservices bool `json:"storefrontcheckbackendservices,string,omitempty"`
	Hostname                       bool `json:"hostname,string,omitempty"`
	Netprofile                     bool `json:"netprofile,string,omitempty"`
	Mssqlprotocolversion           bool `json:"mssqlprotocolversion,string,omitempty"`
	Originhost                     bool `json:"originhost,string,omitempty"`
	Originrealm                    bool `json:"originrealm,string,omitempty"`
	Hostipaddress                  bool `json:"hostipaddress,string,omitempty"`
	Vendorid                       bool `json:"vendorid,string,omitempty"`
	Productname                    bool `json:"productname,string,omitempty"`
	Firmwarerevision               bool `json:"firmwarerevision,string,omitempty"`
	Vendorspecificvendorid         bool `json:"vendorspecificvendorid,string,omitempty"`
	Kcdaccount                     bool `json:"kcdaccount,string,omitempty"`
	Storedb                        bool `json:"storedb,string,omitempty"`
	Trofscode                      bool `json:"trofscode,string,omitempty"`
	Trofsstring                    bool `json:"trofsstring,string,omitempty"`
	Sslprofile                     bool `json:"sslprofile,string,omitempty"`
}

type update_lbmonitor struct {
	Monitorname string `json:"monitorname"`
	// TODO
	// Type string `json:"type,omitempty"`
	Type                           string `json:"type,omitempty"`
	Action                         string `json:"action,omitempty"`
	Httprequest                    string `json:"httprequest,omitempty"`
	Rtsprequest                    string `json:"rtsprequest,omitempty"`
	Customheaders                  string `json:"customheaders,omitempty"`
	Maxforwards                    int    `json:"maxforwards,string,omitempty"`
	Sipmethod                      string `json:"sipmethod,omitempty"`
	Sipreguri                      string `json:"sipreguri,omitempty"`
	Sipuri                         string `json:"sipuri,omitempty"`
	Send                           string `json:"send,omitempty"`
	Recv                           string `json:"recv,omitempty"`
	Query                          string `json:"query,omitempty"`
	Querytype                      string `json:"querytype,omitempty"`
	Username                       string `json:"username,omitempty"`
	Password                       string `json:"password,omitempty"`
	Secondarypassword              string `json:"secondarypassword,omitempty"`
	Logonpointname                 string `json:"logonpointname,omitempty"`
	Lasversion                     string `json:"lasversion,omitempty"`
	Radkey                         string `json:"radkey,omitempty"`
	Radnasid                       string `json:"radnasid,omitempty"`
	Radnasip                       string `json:"radnasip,omitempty"`
	Radaccounttype                 int    `json:"radaccounttype,string,omitempty"`
	Radframedip                    string `json:"radframedip,omitempty"`
	Radapn                         string `json:"radapn,omitempty"`
	Radmsisdn                      string `json:"radmsisdn,omitempty"`
	Radaccountsession              string `json:"radaccountsession,omitempty"`
	Lrtm                           string `json:"lrtm,omitempty"`
	Deviation                      int    `json:"deviation,string,omitempty"`
	Units1                         string `json:"units1,omitempty"`
	Scriptname                     string `json:"scriptname,omitempty"`
	Scriptargs                     string `json:"scriptargs,omitempty"`
	Validatecred                   string `json:"validatecred,omitempty"`
	Domain                         string `json:"domain,omitempty"`
	Dispatcherip                   string `json:"dispatcherip,omitempty"`
	Dispatcherport                 int    `json:"dispatcherport,omitempty"`
	Interval                       int    `json:"interval,omitempty"`
	Units3                         string `json:"units3,omitempty"`
	Resptimeout                    int    `json:"resptimeout,omitempty"`
	Units4                         string `json:"units4,omitempty"`
	Resptimeoutthresh              int    `json:"resptimeoutthresh,string,omitempty"`
	Retries                        int    `json:"retries,omitempty"`
	Failureretries                 int    `json:"failureretries,omitempty"`
	Alertretries                   int    `json:"alertretries,omitempty"`
	Successretries                 int    `json:"successretries,omitempty"`
	Downtime                       int    `json:"downtime,omitempty"`
	Units2                         string `json:"units2,omitempty"`
	Destip                         string `json:"destip,omitempty"`
	Destport                       int    `json:"destport,omitempty"`
	State                          string `json:"state,omitempty"`
	Reverse                        string `json:"reverse,omitempty"`
	Transparent                    string `json:"transparent,omitempty"`
	Iptunnel                       string `json:"iptunnel,omitempty"`
	Tos                            string `json:"tos,omitempty"`
	Tosid                          int    `json:"tosid,string,omitempty"`
	Secure                         string `json:"secure,omitempty"`
	Group                          string `json:"group,omitempty"`
	Filename                       string `json:"filename,omitempty"`
	Basedn                         string `json:"basedn,omitempty"`
	Binddn                         string `json:"binddn,omitempty"`
	Filter                         string `json:"filter,omitempty"`
	Attribute                      string `json:"attribute,omitempty"`
	Database                       string `json:"database,omitempty"`
	Oraclesid                      string `json:"oraclesid,omitempty"`
	Sqlquery                       string `json:"sqlquery,omitempty"`
	Evalrule                       string `json:"evalrule,omitempty"`
	Snmpoid                        string `json:"snmpoid,omitempty"`
	Snmpcommunity                  string `json:"snmpcommunity,omitempty"`
	Snmpthreshold                  string `json:"snmpthreshold,omitempty"`
	Snmpversion                    string `json:"snmpversion,omitempty"`
	Metrictable                    string `json:"metrictable,omitempty"`
	Application                    string `json:"application,omitempty"`
	Sitepath                       string `json:"sitepath,omitempty"`
	Storename                      string `json:"storename,omitempty"`
	Storefrontacctservice          string `json:"storefrontacctservice,omitempty"`
	Storefrontcheckbackendservices string `json:"storefrontcheckbackendservices,omitempty"`
	Hostname                       string `json:"hostname,omitempty"`
	Netprofile                     string `json:"netprofile,omitempty"`
	Mssqlprotocolversion           string `json:"mssqlprotocolversion,omitempty"`
	Originhost                     string `json:"originhost,omitempty"`
	Originrealm                    string `json:"originrealm,omitempty"`
	Hostipaddress                  string `json:"hostipaddress,omitempty"`
	Vendorid                       int    `json:"vendorid,string,omitempty"`
	Productname                    string `json:"productname,omitempty"`
	Firmwarerevision               int    `json:"firmwarerevision,string,omitempty"`
	Vendorspecificvendorid         int    `json:"vendorspecificvendorid,string,omitempty"`
	Kcdaccount                     string `json:"kcdaccount,omitempty"`
	Storedb                        string `json:"storedb,omitempty"`
	Trofscode                      int    `json:"trofscode,string,omitempty"`
	Trofsstring                    string `json:"trofsstring,omitempty"`
	Sslprofile                     string `json:"sslprofile,omitempty"`
}

type rename_lbmonitor struct {
	Name    string `json:"monitorname"`
	Newname string `json:"newname"`
}

type add_lbmonitor_payload struct {
	Resource Lbmonitor `json:"lbmonitor"`
}

type rename_lbmonitor_payload struct {
	Rename rename_lbmonitor `json:"lbmonitor"`
}

type state_lbmonitor struct {
	Key LbmonitorKey `json:"monitorname"`
}

type state_lbmonitor_payload struct {
	Sate state_lbmonitor `json:"lbmonitor"`
}

type unset_lbmonitor_payload struct {
	Unset LbmonitorUnset `json:"lbmonitor"`
}

type update_lbmonitor_payload struct {
	Update update_lbmonitor `json:"lbmonitor"`
}

type get_lbmonitor_result struct {
	Results []Lbmonitor `json:"lbmonitor"`
}

type count_lbmonitor_result struct {
	Results []Count `json:"lbmonitor"`
}

func (c *NitroClient) AddLbmonitor(resource Lbmonitor) error {
	payload := add_lbmonitor_payload{
		resource,
	}

	return c.post("lbmonitor", "", nil, payload)
}

func (c *NitroClient) RenameLbmonitor(name string, newName string) error {
	payload := rename_lbmonitor_payload{
		rename_lbmonitor{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("lbmonitor", "", qs, payload)
}

func (c *NitroClient) CountLbmonitor() (int, error) {
	var results count_lbmonitor_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbmonitor", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbmonitor(key LbmonitorKey) (bool, error) {
	var results count_lbmonitor_result

	id, qs := lbmonitor_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("lbmonitor", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListLbmonitor() ([]Lbmonitor, error) {
	results := get_lbmonitor_result{}

	if err := c.get("lbmonitor", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbmonitor(key LbmonitorKey) (*Lbmonitor, error) {
	var results get_lbmonitor_result

	id, qs := lbmonitor_key_to_id_args(key)

	if err := c.get("lbmonitor", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbmonitor element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbmonitor element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbmonitor(key LbmonitorKey) error {
	id, qs := lbmonitor_key_to_id_args(key)

	return c.delete("lbmonitor", id, qs)
}

func (c *NitroClient) UnsetLbmonitor(unset LbmonitorUnset) error {
	payload := unset_lbmonitor_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("lbmonitor", "", qs, payload)
}

func (c *NitroClient) UpdateLbmonitor(resource Lbmonitor) error {
	payload := update_lbmonitor_payload{
		update_lbmonitor{
			resource.Monitorname,
			// TODO
			// resource.Type,
			resource.Type,
			resource.Action,
			resource.Httprequest,
			resource.Rtsprequest,
			resource.Customheaders,
			resource.Maxforwards,
			resource.Sipmethod,
			resource.Sipreguri,
			resource.Sipuri,
			resource.Send,
			resource.Recv,
			resource.Query,
			resource.Querytype,
			resource.Username,
			resource.Password,
			resource.Secondarypassword,
			resource.Logonpointname,
			resource.Lasversion,
			resource.Radkey,
			resource.Radnasid,
			resource.Radnasip,
			resource.Radaccounttype,
			resource.Radframedip,
			resource.Radapn,
			resource.Radmsisdn,
			resource.Radaccountsession,
			resource.Lrtm,
			resource.Deviation,
			resource.Units1,
			resource.Scriptname,
			resource.Scriptargs,
			resource.Validatecred,
			resource.Domain,
			resource.Dispatcherip,
			resource.Dispatcherport,
			resource.Interval,
			resource.Units3,
			resource.Resptimeout,
			resource.Units4,
			resource.Resptimeoutthresh,
			resource.Retries,
			resource.Failureretries,
			resource.Alertretries,
			resource.Successretries,
			resource.Downtime,
			resource.Units2,
			resource.Destip,
			resource.Destport,
			resource.State,
			resource.Reverse,
			resource.Transparent,
			resource.Iptunnel,
			resource.Tos,
			resource.Tosid,
			resource.Secure,
			resource.Group,
			resource.Filename,
			resource.Basedn,
			resource.Binddn,
			resource.Filter,
			resource.Attribute,
			resource.Database,
			resource.Oraclesid,
			resource.Sqlquery,
			resource.Evalrule,
			resource.Snmpoid,
			resource.Snmpcommunity,
			resource.Snmpthreshold,
			resource.Snmpversion,
			resource.Metrictable,
			resource.Application,
			resource.Sitepath,
			resource.Storename,
			resource.Storefrontacctservice,
			resource.Storefrontcheckbackendservices,
			resource.Hostname,
			resource.Netprofile,
			resource.Mssqlprotocolversion,
			resource.Originhost,
			resource.Originrealm,
			resource.Hostipaddress,
			resource.Vendorid,
			resource.Productname,
			resource.Firmwarerevision,
			resource.Vendorspecificvendorid,
			resource.Kcdaccount,
			resource.Storedb,
			resource.Trofscode,
			resource.Trofsstring,
			resource.Sslprofile,
		},
	}

	return c.put("lbmonitor", "", nil, payload)
}

func (c *NitroClient) EnableLbmonitor(key LbmonitorKey) error {
	payload := state_lbmonitor_payload{
		state_lbmonitor{
			key,
		},
	}

	qs := map[string]string{
		"action": "enable",
	}

	return c.post("lbmonitor", "", qs, payload)
}

func (c *NitroClient) DisableLbmonitor(key LbmonitorKey) error {
	payload := state_lbmonitor_payload{
		state_lbmonitor{
			key,
		},
	}

	qs := map[string]string{
		"action": "disable",
	}

	return c.post("lbmonitor", "", qs, payload)
}
