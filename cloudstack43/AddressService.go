//
// Copyright 2014, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack43

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type AssociateIpAddressParams struct {
	p map[string]interface{}
}

func (p *AssociateIpAddressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["isportable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isportable", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AssociateIpAddressParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *AssociateIpAddressParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *AssociateIpAddressParams) SetIsportable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isportable"] = v
	return
}

func (p *AssociateIpAddressParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *AssociateIpAddressParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *AssociateIpAddressParams) SetRegionid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["regionid"] = v
	return
}

func (p *AssociateIpAddressParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

func (p *AssociateIpAddressParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AssociateIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewAssociateIpAddressParams() *AssociateIpAddressParams {
	p := &AssociateIpAddressParams{}
	p.p = make(map[string]interface{})
	return p
}

// Acquires and associates a public IP to an account.
func (s *AddressService) AssociateIpAddress(p *AssociateIpAddressParams) (*AssociateIpAddressResponse, error) {
	resp, err := s.cs.newRequest("associateIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssociateIpAddressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type AssociateIpAddressResponse struct {
	JobID                     string `json:"jobid,omitempty"`
	Account                   string `json:"account,omitempty"`
	Isstaticnat               bool   `json:"isstaticnat,omitempty"`
	Allocated                 string `json:"allocated,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Zoneid                    string `json:"zoneid,omitempty"`
	Associatednetworkid       string `json:"associatednetworkid,omitempty"`
	Vpcid                     string `json:"vpcid,omitempty"`
	Vlanname                  string `json:"vlanname,omitempty"`
	Virtualmachineid          string `json:"virtualmachineid,omitempty"`
	Vlanid                    string `json:"vlanid,omitempty"`
	Associatednetworkname     string `json:"associatednetworkname,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Zonename                  string `json:"zonename,omitempty"`
	Projectid                 string `json:"projectid,omitempty"`
	Id                        string `json:"id,omitempty"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname,omitempty"`
	Purpose                   string `json:"purpose,omitempty"`
	Networkid                 string `json:"networkid,omitempty"`
	Issourcenat               bool   `json:"issourcenat,omitempty"`
	Isportable                bool   `json:"isportable,omitempty"`
	Ipaddress                 string `json:"ipaddress,omitempty"`
	Virtualmachinename        string `json:"virtualmachinename,omitempty"`
	Physicalnetworkid         string `json:"physicalnetworkid,omitempty"`
	Vmipaddress               string `json:"vmipaddress,omitempty"`
	State                     string `json:"state,omitempty"`
	Project                   string `json:"project,omitempty"`
	Issystem                  bool   `json:"issystem,omitempty"`
	Tags                      []struct {
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Forvirtualnetwork bool `json:"forvirtualnetwork,omitempty"`
}

type DisassociateIpAddressParams struct {
	p map[string]interface{}
}

func (p *DisassociateIpAddressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DisassociateIpAddressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DisassociateIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewDisassociateIpAddressParams(id string) *DisassociateIpAddressParams {
	p := &DisassociateIpAddressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Disassociates an ip address from the account.
func (s *AddressService) DisassociateIpAddress(p *DisassociateIpAddressParams) (*DisassociateIpAddressResponse, error) {
	resp, err := s.cs.newRequest("disassociateIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisassociateIpAddressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DisassociateIpAddressResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListPublicIpAddressesParams struct {
	p map[string]interface{}
}

func (p *ListPublicIpAddressesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["allocatedonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("allocatedonly", vv)
	}
	if v, found := p.p["associatednetworkid"]; found {
		u.Set("associatednetworkid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["forloadbalancing"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forloadbalancing", vv)
	}
	if v, found := p.p["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["issourcenat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issourcenat", vv)
	}
	if v, found := p.p["isstaticnat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isstaticnat", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["vlanid"]; found {
		u.Set("vlanid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListPublicIpAddressesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetAllocatedonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocatedonly"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetAssociatednetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["associatednetworkid"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetForloadbalancing(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forloadbalancing"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetForvirtualnetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvirtualnetwork"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetIssourcenat(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issourcenat"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetIsstaticnat(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isstaticnat"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetVlanid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlanid"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

func (p *ListPublicIpAddressesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListPublicIpAddressesParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewListPublicIpAddressesParams() *ListPublicIpAddressesParams {
	p := &ListPublicIpAddressesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AddressService) GetPublicIpAddressByID(id string) (*PublicIpAddress, int, error) {
	p := &ListPublicIpAddressesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListPublicIpAddresses(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.PublicIpAddresses[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for PublicIpAddress UUID: %s!", id)
}

// Lists all public ip addresses
func (s *AddressService) ListPublicIpAddresses(p *ListPublicIpAddressesParams) (*ListPublicIpAddressesResponse, error) {
	resp, err := s.cs.newRequest("listPublicIpAddresses", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPublicIpAddressesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListPublicIpAddressesResponse struct {
	Count             int                `json:"count"`
	PublicIpAddresses []*PublicIpAddress `json:"publicipaddress"`
}

type PublicIpAddress struct {
	Isportable                bool   `json:"isportable,omitempty"`
	Account                   string `json:"account,omitempty"`
	Virtualmachinename        string `json:"virtualmachinename,omitempty"`
	State                     string `json:"state,omitempty"`
	Networkid                 string `json:"networkid,omitempty"`
	Physicalnetworkid         string `json:"physicalnetworkid,omitempty"`
	Zonename                  string `json:"zonename,omitempty"`
	Vmipaddress               string `json:"vmipaddress,omitempty"`
	Vpcid                     string `json:"vpcid,omitempty"`
	Vlanid                    string `json:"vlanid,omitempty"`
	Allocated                 string `json:"allocated,omitempty"`
	Virtualmachineid          string `json:"virtualmachineid,omitempty"`
	Id                        string `json:"id,omitempty"`
	Projectid                 string `json:"projectid,omitempty"`
	Ipaddress                 string `json:"ipaddress,omitempty"`
	Issourcenat               bool   `json:"issourcenat,omitempty"`
	Project                   string `json:"project,omitempty"`
	Vlanname                  string `json:"vlanname,omitempty"`
	Issystem                  bool   `json:"issystem,omitempty"`
	Associatednetworkname     string `json:"associatednetworkname,omitempty"`
	Isstaticnat               bool   `json:"isstaticnat,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Associatednetworkid       string `json:"associatednetworkid,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Purpose                   string `json:"purpose,omitempty"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork,omitempty"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname,omitempty"`
	Tags                      []struct {
		Projectid    string `json:"projectid,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Zoneid string `json:"zoneid,omitempty"`
}