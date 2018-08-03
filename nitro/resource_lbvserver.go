package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Lbvserver struct {
	Name                               string `json:"name"`
	State                              string `json:"state,omitempty"`
	Appflowlog                         string `json:"appflowlog,omitempty"`
	Authentication                     string `json:"authentication,omitempty"`
	Authenticationhost                 string `json:"authenticationhost,omitempty"`
	Authn401                           string `json:"authn401,omitempty"`
	Authnprofile                       string `json:"authnprofile,omitempty"`
	Authnvsname                        string `json:"authnvsname,omitempty"`
	Backuplbmethod                     string `json:"backuplbmethod,omitempty"`
	Backuppersistencetimeout           int    `json:"backuppersistencetimeout,omitempty"`
	Backupvserver                      string `json:"backupvserver,omitempty"`
	Bypassaaaa                         string `json:"bypassaaaa,omitempty"`
	Cacheable                          string `json:"cacheable,omitempty"`
	Clttimeout                         int    `json:"clttimeout,string,omitempty"`
	Comment                            string `json:"comment,omitempty"`
	Connfailover                       string `json:"connfailover,omitempty"`
	Cookiename                         string `json:"cookiename,omitempty"`
	Datalength                         int    `json:"datalength,string,omitempty"`
	Dataoffset                         int    `json:"dataoffset,string,omitempty"`
	Dbprofilename                      string `json:"dbprofilename,omitempty"`
	Dbslb                              string `json:"dbslb,omitempty"`
	Disableprimaryondown               string `json:"disableprimaryondown,omitempty"`
	Dns64                              string `json:"dns64,omitempty"`
	Dnsprofilename                     string `json:"dnsprofilename,omitempty"`
	Downstateflush                     string `json:"downstateflush,omitempty"`
	Hashlength                         int    `json:"hashlength,string,omitempty"`
	Healththreshold                    int    `json:"healththreshold,string,omitempty"`
	Httpprofilename                    string `json:"httpprofilename,omitempty"`
	Httpsredirecturl                   string `json:"httpsredirecturl,omitempty"`
	Icmpvsrresponse                    string `json:"icmpvsrresponse,omitempty"`
	Insertvserveripport                string `json:"insertvserveripport,omitempty"`
	Ipmask                             string `json:"ipmask,omitempty"`
	Ippattern                          string `json:"ippattern,omitempty"`
	Ipv46                              string `json:"ipv46,omitempty"`
	L2conn                             string `json:"l2conn,omitempty"`
	Lbmethod                           string `json:"lbmethod,omitempty"`
	Lbprofilename                      string `json:"lbprofilename,omitempty"`
	Listenpolicy                       string `json:"listenpolicy,omitempty"`
	Listenpriority                     int    `json:"listenpriority,string,omitempty"`
	M                                  string `json:"m,omitempty"`
	Macmoderetainvlan                  string `json:"macmoderetainvlan,omitempty"`
	Maxautoscalemembers                int    `json:"maxautoscalemembers,string,omitempty"`
	Minautoscalemembers                int    `json:"minautoscalemembers,string,omitempty"`
	Mssqlserverversion                 string `json:"mssqlserverversion,omitempty"`
	Mysqlcharacterset                  int    `json:"mysqlcharacterset,string,omitempty"`
	Mysqlprotocolversion               int    `json:"mysqlprotocolversion,string,omitempty"`
	Mysqlservercapabilities            int    `json:"mysqlservercapabilities,string,omitempty"`
	Mysqlserverversion                 string `json:"mysqlserverversion,omitempty"`
	Netmask                            string `json:"netmask,omitempty"`
	Netprofile                         string `json:"netprofile,omitempty"`
	Newservicerequest                  int    `json:"newservicerequest,string,omitempty"`
	Newservicerequestincrementinterval int    `json:"newservicerequestincrementinterval,string,omitempty"`
	Newservicerequestunit              string `json:"newservicerequestunit,omitempty"`
	Oracleserverversion                string `json:"oracleserverversion,omitempty"`
	Persistavpno                       []int  `json:"persistavpno,omitempty"`
	Persistencebackup                  string `json:"persistencebackup,omitempty"`
	Persistencetype                    string `json:"persistencetype,omitempty"`
	Persistmask                        string `json:"persistmask,omitempty"`
	Port                               int    `json:"port,omitempty"`
	Pq                                 string `json:"pq,omitempty"`
	Processlocal                       string `json:"processlocal,omitempty"`
	Push                               string `json:"push,omitempty"`
	Pushlabel                          string `json:"pushlabel,omitempty"`
	Pushmulticlients                   string `json:"pushmulticlients,omitempty"`
	Pushvserver                        string `json:"pushvserver,omitempty"`
	Range                              int    `json:"range,string,omitempty"`
	Recursionavailable                 string `json:"recursionavailable,omitempty"`
	Redirectfromport                   int    `json:"redirectfromport,omitempty"`
	Redirectportrewrite                string `json:"redirectportrewrite,omitempty"`
	Redirurl                           string `json:"redirurl,omitempty"`
	Resrule                            string `json:"resrule,omitempty"`
	Retainconnectionsoncluster         string `json:"retainconnectionsoncluster,omitempty"`
	Rhistate                           string `json:"rhistate,omitempty"`
	Rtspnat                            string `json:"rtspnat,omitempty"`
	Rule                               string `json:"rule,omitempty"`
	Sc                                 string `json:"sc,omitempty"`
	Servicetype                        string `json:"servicetype,omitempty"`
	Sessionless                        string `json:"sessionless,omitempty"`
	Skippersistency                    string `json:"skippersistency,omitempty"`
	Sobackupaction                     string `json:"sobackupaction,omitempty"`
	Somethod                           string `json:"somethod,omitempty"`
	Sopersistence                      string `json:"sopersistence,omitempty"`
	Sopersistencetimeout               int    `json:"sopersistencetimeout,string,omitempty"`
	Sothreshold                        int    `json:"sothreshold,string,omitempty"`
	Tcpprofilename                     string `json:"tcpprofilename,omitempty"`
	Td                                 int    `json:"td,string,omitempty"`
	Timeout                            int    `json:"timeout,omitempty"`
	Tosid                              int    `json:"tosid,string,omitempty"`
	Trofspersistence                   string `json:"trofspersistence,omitempty"`
	V6netmasklen                       int    `json:"v6netmasklen,string,omitempty"`
	V6persistmasklen                   int    `json:"v6persistmasklen,string,omitempty"`
	Vipheader                          string `json:"vipheader,omitempty"`
}

type LbvserverKey struct {
	Name string `json:"name"`
}

type LbvserverUnset struct {
	Name                               string `json:"name"`
	Ipv46                              bool   `json:"ipv46,string,omitempty"`
	Ippattern                          bool   `json:"ippattern,string,omitempty"`
	Ipmask                             bool   `json:"ipmask,string,omitempty"`
	Persistencetype                    bool   `json:"persistencetype,string,omitempty"`
	Timeout                            bool   `json:"timeout,string,omitempty"`
	Persistencebackup                  bool   `json:"persistencebackup,string,omitempty"`
	Backuppersistencetimeout           bool   `json:"backuppersistencetimeout,string,omitempty"`
	Lbmethod                           bool   `json:"lbmethod,string,omitempty"`
	Hashlength                         bool   `json:"hashlength,string,omitempty"`
	Netmask                            bool   `json:"netmask,string,omitempty"`
	V6netmasklen                       bool   `json:"v6netmasklen,string,omitempty"`
	Backuplbmethod                     bool   `json:"backuplbmethod,string,omitempty"`
	Rule                               bool   `json:"rule,string,omitempty"`
	Cookiename                         bool   `json:"cookiename,string,omitempty"`
	Resrule                            bool   `json:"resrule,string,omitempty"`
	Persistmask                        bool   `json:"persistmask,string,omitempty"`
	V6persistmasklen                   bool   `json:"v6persistmasklen,string,omitempty"`
	Pq                                 bool   `json:"pq,string,omitempty"`
	Sc                                 bool   `json:"sc,string,omitempty"`
	Rtspnat                            bool   `json:"rtspnat,string,omitempty"`
	M                                  bool   `json:"m,string,omitempty"`
	Tosid                              bool   `json:"tosid,string,omitempty"`
	Datalength                         bool   `json:"datalength,string,omitempty"`
	Dataoffset                         bool   `json:"dataoffset,string,omitempty"`
	Sessionless                        bool   `json:"sessionless,string,omitempty"`
	Trofspersistence                   bool   `json:"trofspersistence,string,omitempty"`
	Connfailover                       bool   `json:"connfailover,string,omitempty"`
	Backupvserver                      bool   `json:"backupvserver,string,omitempty"`
	Redirurl                           bool   `json:"redirurl,string,omitempty"`
	Cacheable                          bool   `json:"cacheable,string,omitempty"`
	Clttimeout                         bool   `json:"clttimeout,string,omitempty"`
	Somethod                           bool   `json:"somethod,string,omitempty"`
	Sothreshold                        bool   `json:"sothreshold,string,omitempty"`
	Sopersistence                      bool   `json:"sopersistence,string,omitempty"`
	Sopersistencetimeout               bool   `json:"sopersistencetimeout,string,omitempty"`
	Healththreshold                    bool   `json:"healththreshold,string,omitempty"`
	Sobackupaction                     bool   `json:"sobackupaction,string,omitempty"`
	Redirectportrewrite                bool   `json:"redirectportrewrite,string,omitempty"`
	Downstateflush                     bool   `json:"downstateflush,string,omitempty"`
	Insertvserveripport                bool   `json:"insertvserveripport,string,omitempty"`
	Vipheader                          bool   `json:"vipheader,string,omitempty"`
	Disableprimaryondown               bool   `json:"disableprimaryondown,string,omitempty"`
	Authenticationhost                 bool   `json:"authenticationhost,string,omitempty"`
	Authentication                     bool   `json:"authentication,string,omitempty"`
	Authn401                           bool   `json:"authn401,string,omitempty"`
	Authnvsname                        bool   `json:"authnvsname,string,omitempty"`
	Push                               bool   `json:"push,string,omitempty"`
	Pushvserver                        bool   `json:"pushvserver,string,omitempty"`
	Pushlabel                          bool   `json:"pushlabel,string,omitempty"`
	Pushmulticlients                   bool   `json:"pushmulticlients,string,omitempty"`
	Listenpolicy                       bool   `json:"listenpolicy,string,omitempty"`
	Listenpriority                     bool   `json:"listenpriority,string,omitempty"`
	Tcpprofilename                     bool   `json:"tcpprofilename,string,omitempty"`
	Httpprofilename                    bool   `json:"httpprofilename,string,omitempty"`
	Dbprofilename                      bool   `json:"dbprofilename,string,omitempty"`
	Comment                            bool   `json:"comment,string,omitempty"`
	L2conn                             bool   `json:"l2conn,string,omitempty"`
	Oracleserverversion                bool   `json:"oracleserverversion,string,omitempty"`
	Mssqlserverversion                 bool   `json:"mssqlserverversion,string,omitempty"`
	Mysqlprotocolversion               bool   `json:"mysqlprotocolversion,string,omitempty"`
	Mysqlserverversion                 bool   `json:"mysqlserverversion,string,omitempty"`
	Mysqlcharacterset                  bool   `json:"mysqlcharacterset,string,omitempty"`
	Mysqlservercapabilities            bool   `json:"mysqlservercapabilities,string,omitempty"`
	Appflowlog                         bool   `json:"appflowlog,string,omitempty"`
	Netprofile                         bool   `json:"netprofile,string,omitempty"`
	Icmpvsrresponse                    bool   `json:"icmpvsrresponse,string,omitempty"`
	Rhistate                           bool   `json:"rhistate,string,omitempty"`
	Newservicerequest                  bool   `json:"newservicerequest,string,omitempty"`
	Newservicerequestunit              bool   `json:"newservicerequestunit,string,omitempty"`
	Newservicerequestincrementinterval bool   `json:"newservicerequestincrementinterval,string,omitempty"`
	Minautoscalemembers                bool   `json:"minautoscalemembers,string,omitempty"`
	Maxautoscalemembers                bool   `json:"maxautoscalemembers,string,omitempty"`
	Persistavpno                       bool   `json:"persistavpno,string,omitempty"`
	Skippersistency                    bool   `json:"skippersistency,string,omitempty"`
	Authnprofile                       bool   `json:"authnprofile,string,omitempty"`
	Macmoderetainvlan                  bool   `json:"macmoderetainvlan,string,omitempty"`
	Dbslb                              bool   `json:"dbslb,string,omitempty"`
	Dns64                              bool   `json:"dns64,string,omitempty"`
	Bypassaaaa                         bool   `json:"bypassaaaa,string,omitempty"`
	Recursionavailable                 bool   `json:"recursionavailable,string,omitempty"`
	Processlocal                       bool   `json:"processlocal,string,omitempty"`
	Dnsprofilename                     bool   `json:"dnsprofilename,string,omitempty"`
	Lbprofilename                      bool   `json:"lbprofilename,string,omitempty"`
	Redirectfromport                   bool   `json:"redirectfromport,string,omitempty"`
	Httpsredirecturl                   bool   `json:"httpsredirecturl,string,omitempty"`
	Retainconnectionsoncluster         bool   `json:"retainconnectionsoncluster,string,omitempty"`
}

type update_lbvserver struct {
	Name                               string `json:"name"`
	Ipv46                              string `json:"ipv46,omitempty"`
	Ippattern                          string `json:"ippattern,omitempty"`
	Ipmask                             string `json:"ipmask,omitempty"`
	Persistencetype                    string `json:"persistencetype,omitempty"`
	Timeout                            int    `json:"timeout,omitempty"`
	Persistencebackup                  string `json:"persistencebackup,omitempty"`
	Backuppersistencetimeout           int    `json:"backuppersistencetimeout,omitempty"`
	Lbmethod                           string `json:"lbmethod,omitempty"`
	Hashlength                         int    `json:"hashlength,string,omitempty"`
	Netmask                            string `json:"netmask,omitempty"`
	V6netmasklen                       int    `json:"v6netmasklen,string,omitempty"`
	Backuplbmethod                     string `json:"backuplbmethod,omitempty"`
	Rule                               string `json:"rule,omitempty"`
	Cookiename                         string `json:"cookiename,omitempty"`
	Resrule                            string `json:"resrule,omitempty"`
	Persistmask                        string `json:"persistmask,omitempty"`
	V6persistmasklen                   int    `json:"v6persistmasklen,string,omitempty"`
	Pq                                 string `json:"pq,omitempty"`
	Sc                                 string `json:"sc,omitempty"`
	Rtspnat                            string `json:"rtspnat,omitempty"`
	M                                  string `json:"m,omitempty"`
	Tosid                              int    `json:"tosid,string,omitempty"`
	Datalength                         int    `json:"datalength,string,omitempty"`
	Dataoffset                         int    `json:"dataoffset,string,omitempty"`
	Sessionless                        string `json:"sessionless,omitempty"`
	Trofspersistence                   string `json:"trofspersistence,omitempty"`
	Connfailover                       string `json:"connfailover,omitempty"`
	Backupvserver                      string `json:"backupvserver,omitempty"`
	Redirurl                           string `json:"redirurl,omitempty"`
	Cacheable                          string `json:"cacheable,omitempty"`
	Clttimeout                         int    `json:"clttimeout,string,omitempty"`
	Somethod                           string `json:"somethod,omitempty"`
	Sothreshold                        int    `json:"sothreshold,string,omitempty"`
	Sopersistence                      string `json:"sopersistence,omitempty"`
	Sopersistencetimeout               int    `json:"sopersistencetimeout,string,omitempty"`
	Healththreshold                    int    `json:"healththreshold,string,omitempty"`
	Sobackupaction                     string `json:"sobackupaction,omitempty"`
	Redirectportrewrite                string `json:"redirectportrewrite,omitempty"`
	Downstateflush                     string `json:"downstateflush,omitempty"`
	Insertvserveripport                string `json:"insertvserveripport,omitempty"`
	Vipheader                          string `json:"vipheader,omitempty"`
	Disableprimaryondown               string `json:"disableprimaryondown,omitempty"`
	Authenticationhost                 string `json:"authenticationhost,omitempty"`
	Authentication                     string `json:"authentication,omitempty"`
	Authn401                           string `json:"authn401,omitempty"`
	Authnvsname                        string `json:"authnvsname,omitempty"`
	Push                               string `json:"push,omitempty"`
	Pushvserver                        string `json:"pushvserver,omitempty"`
	Pushlabel                          string `json:"pushlabel,omitempty"`
	Pushmulticlients                   string `json:"pushmulticlients,omitempty"`
	Listenpolicy                       string `json:"listenpolicy,omitempty"`
	Listenpriority                     int    `json:"listenpriority,string,omitempty"`
	Tcpprofilename                     string `json:"tcpprofilename,omitempty"`
	Httpprofilename                    string `json:"httpprofilename,omitempty"`
	Dbprofilename                      string `json:"dbprofilename,omitempty"`
	Comment                            string `json:"comment,omitempty"`
	L2conn                             string `json:"l2conn,omitempty"`
	Oracleserverversion                string `json:"oracleserverversion,omitempty"`
	Mssqlserverversion                 string `json:"mssqlserverversion,omitempty"`
	Mysqlprotocolversion               int    `json:"mysqlprotocolversion,string,omitempty"`
	Mysqlserverversion                 string `json:"mysqlserverversion,omitempty"`
	Mysqlcharacterset                  int    `json:"mysqlcharacterset,string,omitempty"`
	Mysqlservercapabilities            int    `json:"mysqlservercapabilities,string,omitempty"`
	Appflowlog                         string `json:"appflowlog,omitempty"`
	Netprofile                         string `json:"netprofile,omitempty"`
	Icmpvsrresponse                    string `json:"icmpvsrresponse,omitempty"`
	Rhistate                           string `json:"rhistate,omitempty"`
	Newservicerequest                  int    `json:"newservicerequest,string,omitempty"`
	Newservicerequestunit              string `json:"newservicerequestunit,omitempty"`
	Newservicerequestincrementinterval int    `json:"newservicerequestincrementinterval,string,omitempty"`
	Minautoscalemembers                int    `json:"minautoscalemembers,string,omitempty"`
	Maxautoscalemembers                int    `json:"maxautoscalemembers,string,omitempty"`
	Persistavpno                       []int  `json:"persistavpno,omitempty"`
	Skippersistency                    string `json:"skippersistency,omitempty"`
	Authnprofile                       string `json:"authnprofile,omitempty"`
	Macmoderetainvlan                  string `json:"macmoderetainvlan,omitempty"`
	Dbslb                              string `json:"dbslb,omitempty"`
	Dns64                              string `json:"dns64,omitempty"`
	Bypassaaaa                         string `json:"bypassaaaa,omitempty"`
	Recursionavailable                 string `json:"recursionavailable,omitempty"`
	Processlocal                       string `json:"processlocal,omitempty"`
	Dnsprofilename                     string `json:"dnsprofilename,omitempty"`
	Lbprofilename                      string `json:"lbprofilename,omitempty"`
	Redirectfromport                   int    `json:"redirectfromport,omitempty"`
	Httpsredirecturl                   string `json:"httpsredirecturl,omitempty"`
	Retainconnectionsoncluster         string `json:"retainconnectionsoncluster,omitempty"`
}

type rename_lbvserver struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_lbvserver_payload struct {
	Resource Lbvserver `json:"lbvserver"`
}

type rename_lbvserver_payload struct {
	Rename rename_lbvserver `json:"lbvserver"`
}

type state_lbvserver_payload struct {
	Key LbvserverKey `json:"lbvserver"`
}

type unset_lbvserver_payload struct {
	Unset LbvserverUnset `json:"lbvserver"`
}

type update_lbvserver_payload struct {
	Update update_lbvserver `json:"lbvserver"`
}

type get_lbvserver_result struct {
	Results []Lbvserver `json:"lbvserver"`
}

type count_lbvserver_result struct {
	Results []Count `json:"lbvserver"`
}

func lbvserver_key_to_id_args(key LbvserverKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return key.Name, qs
}

func (c *NitroClient) AddLbvserver(resource Lbvserver) error {
	payload := add_lbvserver_payload{
		resource,
	}

	return c.post("lbvserver", "", nil, payload)
}

func (c *NitroClient) RenameLbvserver(name string, newName string) error {
	payload := rename_lbvserver_payload{
		rename_lbvserver{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("lbvserver", "", qs, payload)
}

func (c *NitroClient) CountLbvserver() (int, error) {
	var results count_lbvserver_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbvserver", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsLbvserver(key LbvserverKey) (bool, error) {
	var results count_lbvserver_result

	id, qs := lbvserver_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("lbvserver", id, qs, &results); err != nil {
		return false, err
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListLbvserver() ([]Lbvserver, error) {
	var results get_lbvserver_result

	if err := c.get("lbvserver", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetLbvserver(key LbvserverKey) (*Lbvserver, error) {
	var results get_lbvserver_result

	id, qs := lbvserver_key_to_id_args(key)

	if err := c.get("lbvserver", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbvserver element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("lbvserver element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteLbvserver(key LbvserverKey) error {
	id, qs := lbvserver_key_to_id_args(key)

	return c.delete("lbvserver", id, qs)
}

func (c *NitroClient) UnsetLbvserver(unset LbvserverUnset) error {
	payload := unset_lbvserver_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("lbvserver", "", qs, payload)
}

func (c *NitroClient) UpdateLbvserver(resource Lbvserver) error {
	payload := update_lbvserver_payload{
		update_lbvserver{
			resource.Name,
			resource.Ipv46,
			resource.Ippattern,
			resource.Ipmask,
			resource.Persistencetype,
			resource.Timeout,
			resource.Persistencebackup,
			resource.Backuppersistencetimeout,
			resource.Lbmethod,
			resource.Hashlength,
			resource.Netmask,
			resource.V6netmasklen,
			resource.Backuplbmethod,
			resource.Rule,
			resource.Cookiename,
			resource.Resrule,
			resource.Persistmask,
			resource.V6persistmasklen,
			resource.Pq,
			resource.Sc,
			resource.Rtspnat,
			resource.M,
			resource.Tosid,
			resource.Datalength,
			resource.Dataoffset,
			resource.Sessionless,
			resource.Trofspersistence,
			resource.Connfailover,
			resource.Backupvserver,
			resource.Redirurl,
			resource.Cacheable,
			resource.Clttimeout,
			resource.Somethod,
			resource.Sothreshold,
			resource.Sopersistence,
			resource.Sopersistencetimeout,
			resource.Healththreshold,
			resource.Sobackupaction,
			resource.Redirectportrewrite,
			resource.Downstateflush,
			resource.Insertvserveripport,
			resource.Vipheader,
			resource.Disableprimaryondown,
			resource.Authenticationhost,
			resource.Authentication,
			resource.Authn401,
			resource.Authnvsname,
			resource.Push,
			resource.Pushvserver,
			resource.Pushlabel,
			resource.Pushmulticlients,
			resource.Listenpolicy,
			resource.Listenpriority,
			resource.Tcpprofilename,
			resource.Httpprofilename,
			resource.Dbprofilename,
			resource.Comment,
			resource.L2conn,
			resource.Oracleserverversion,
			resource.Mssqlserverversion,
			resource.Mysqlprotocolversion,
			resource.Mysqlserverversion,
			resource.Mysqlcharacterset,
			resource.Mysqlservercapabilities,
			resource.Appflowlog,
			resource.Netprofile,
			resource.Icmpvsrresponse,
			resource.Rhistate,
			resource.Newservicerequest,
			resource.Newservicerequestunit,
			resource.Newservicerequestincrementinterval,
			resource.Minautoscalemembers,
			resource.Maxautoscalemembers,
			resource.Persistavpno,
			resource.Skippersistency,
			resource.Authnprofile,
			resource.Macmoderetainvlan,
			resource.Dbslb,
			resource.Dns64,
			resource.Bypassaaaa,
			resource.Recursionavailable,
			resource.Processlocal,
			resource.Dnsprofilename,
			resource.Lbprofilename,
			resource.Redirectfromport,
			resource.Httpsredirecturl,
			resource.Retainconnectionsoncluster,
		},
	}

	return c.put("lbvserver", "", nil, payload)
}

func (c *NitroClient) EnableLbvserver(key LbvserverKey) error {
	payload := state_lbvserver_payload{
		key,
	}

	qs := map[string]string{
		"action": "enable",
	}

	return c.post("lbvserver", "", qs, payload)
}

func (c *NitroClient) DisableLbvserver(key LbvserverKey) error {
	payload := state_lbvserver_payload{
		key,
	}

	qs := map[string]string{
		"action": "disable",
	}

	return c.post("lbvserver", "", qs, payload)
}
