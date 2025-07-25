// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domain

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "ovh:Domain/dSRecords:DSRecords":
		r = &DSRecords{}
	case "ovh:Domain/dynhostLogin:DynhostLogin":
		r = &DynhostLogin{}
	case "ovh:Domain/name:Name":
		r = &Name{}
	case "ovh:Domain/nameServers:NameServers":
		r = &NameServers{}
	case "ovh:Domain/zone:Zone":
		r = &Zone{}
	case "ovh:Domain/zoneDNSSec:ZoneDNSSec":
		r = &ZoneDNSSec{}
	case "ovh:Domain/zoneImport:ZoneImport":
		r = &ZoneImport{}
	case "ovh:Domain/zoneRecord:ZoneRecord":
		r = &ZoneRecord{}
	case "ovh:Domain/zoneRedirection:ZoneRedirection":
		r = &ZoneRedirection{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/dSRecords",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/dynhostLogin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/name",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/nameServers",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/zone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/zoneDNSSec",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/zoneImport",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/zoneRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/zoneRedirection",
		&module{version},
	)
}
