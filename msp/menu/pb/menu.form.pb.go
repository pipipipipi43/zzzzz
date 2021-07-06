// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: menu.proto

package pb

import (
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetMenuRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetMenuResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*MenuItem)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetSettingRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetSettingResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EngineSetting)(nil)

// GetMenuRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetMenuRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tenantGroup":
				m.TenantGroup = vals[0]
			case "tenantId":
				m.TenantId = vals[0]
			}
		}
	}
	return nil
}

// GetMenuResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetMenuResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// MenuItem implement urlenc.URLValuesUnmarshaler.
func (m *MenuItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clusterName":
				m.ClusterName = vals[0]
			case "clusterType":
				m.ClusterType = vals[0]
			case "key":
				m.Key = vals[0]
			case "cnName":
				m.CnName = vals[0]
			case "enName":
				m.EnName = vals[0]
			case "href":
				m.Href = vals[0]
			case "exists":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Exists = val
			case "mustExists":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.MustExists = val
			case "onlyK8S":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.OnlyK8S = val
			case "onlyNotK8S":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.OnlyNotK8S = val
			}
		}
	}
	return nil
}

// GetSettingRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetSettingRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tenantGroup":
				m.TenantGroup = vals[0]
			case "tenantId":
				m.TenantId = vals[0]
			}
		}
	}
	return nil
}

// GetSettingResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetSettingResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// EngineSetting implement urlenc.URLValuesUnmarshaler.
func (m *EngineSetting) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "addonName":
				m.AddonName = vals[0]
			case "cnName":
				m.CnName = vals[0]
			case "enName":
				m.EnName = vals[0]
			}
		}
	}
	return nil
}
