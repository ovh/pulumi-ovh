# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['NetworkPrivateArgs', 'NetworkPrivate']

@pulumi.input_type
class NetworkPrivateArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 vlan_id: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a NetworkPrivate resource.
        :param pulumi.Input[builtins.str] name: The name of the network.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] regions: an array of valid OVHcloud public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.int] vlan_id: a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if vlan_id is not None:
            pulumi.set(__self__, "vlan_id", vlan_id)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the network.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        an array of valid OVHcloud public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.
        """
        return pulumi.get(self, "vlan_id")

    @vlan_id.setter
    def vlan_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "vlan_id", value)


@pulumi.input_type
class _NetworkPrivateState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 regions_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateRegionsAttributeArgs']]]] = None,
                 regions_openstack_ids: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 regions_statuses: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateRegionsStatusArgs']]]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 vlan_id: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering NetworkPrivate resources.
        :param pulumi.Input[builtins.str] name: The name of the network.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] regions: an array of valid OVHcloud public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkPrivateRegionsAttributeArgs']]] regions_attributes: A map representing information about the region.
               * `regions_attributes/region` - The id of the region.
               * `regions_attributes/status` - The status of the network in the region.
               * `regions_attributes/openstackid` - The private network id in the region.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] regions_openstack_ids: A map with region name as key, and region-specific openstack id as value
        :param pulumi.Input[Sequence[pulumi.Input['NetworkPrivateRegionsStatusArgs']]] regions_statuses: (Deprecated) A map representing the status of the network per region.
               * `regions_status/region` - (Deprecated) The id of the region.
               * `regions_status/status` - (Deprecated) The status of the network in the region.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] status: the status of the network. should be normally set to 'ACTIVE'.
        :param pulumi.Input[builtins.str] type: the type of the network. Either 'private' or 'public'.
        :param pulumi.Input[builtins.int] vlan_id: a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if regions_attributes is not None:
            pulumi.set(__self__, "regions_attributes", regions_attributes)
        if regions_openstack_ids is not None:
            pulumi.set(__self__, "regions_openstack_ids", regions_openstack_ids)
        if regions_statuses is not None:
            warnings.warn("""use the regions_attributes field instead""", DeprecationWarning)
            pulumi.log.warn("""regions_statuses is deprecated: use the regions_attributes field instead""")
        if regions_statuses is not None:
            pulumi.set(__self__, "regions_statuses", regions_statuses)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vlan_id is not None:
            pulumi.set(__self__, "vlan_id", vlan_id)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the network.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        an array of valid OVHcloud public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="regionsAttributes")
    def regions_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateRegionsAttributeArgs']]]]:
        """
        A map representing information about the region.
        * `regions_attributes/region` - The id of the region.
        * `regions_attributes/status` - The status of the network in the region.
        * `regions_attributes/openstackid` - The private network id in the region.
        """
        return pulumi.get(self, "regions_attributes")

    @regions_attributes.setter
    def regions_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateRegionsAttributeArgs']]]]):
        pulumi.set(self, "regions_attributes", value)

    @property
    @pulumi.getter(name="regionsOpenstackIds")
    def regions_openstack_ids(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A map with region name as key, and region-specific openstack id as value
        """
        return pulumi.get(self, "regions_openstack_ids")

    @regions_openstack_ids.setter
    def regions_openstack_ids(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "regions_openstack_ids", value)

    @property
    @pulumi.getter(name="regionsStatuses")
    @_utilities.deprecated("""use the regions_attributes field instead""")
    def regions_statuses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateRegionsStatusArgs']]]]:
        """
        (Deprecated) A map representing the status of the network per region.
        * `regions_status/region` - (Deprecated) The id of the region.
        * `regions_status/status` - (Deprecated) The status of the network in the region.
        """
        return pulumi.get(self, "regions_statuses")

    @regions_statuses.setter
    def regions_statuses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateRegionsStatusArgs']]]]):
        pulumi.set(self, "regions_statuses", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        the status of the network. should be normally set to 'ACTIVE'.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        the type of the network. Either 'private' or 'public'.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.
        """
        return pulumi.get(self, "vlan_id")

    @vlan_id.setter
    def vlan_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "vlan_id", value)


@pulumi.type_token("ovh:CloudProject/networkPrivate:NetworkPrivate")
class NetworkPrivate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 vlan_id: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        Creates a private network in a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        net = ovh.cloud_project.NetworkPrivate("net",
            service_name="XXXXXX",
            name="admin_network",
            regions=[
                "GRA1",
                "BHS1",
            ])
        ```

        ## Import

        Private network in a public cloud project can be imported using the `service_name` and the `network_id` as `pn-xxxx` format, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/networkPrivate:NetworkPrivate mynet service_name/network_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: The name of the network.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] regions: an array of valid OVHcloud public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.int] vlan_id: a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NetworkPrivateArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a private network in a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        net = ovh.cloud_project.NetworkPrivate("net",
            service_name="XXXXXX",
            name="admin_network",
            regions=[
                "GRA1",
                "BHS1",
            ])
        ```

        ## Import

        Private network in a public cloud project can be imported using the `service_name` and the `network_id` as `pn-xxxx` format, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/networkPrivate:NetworkPrivate mynet service_name/network_id
        ```

        :param str resource_name: The name of the resource.
        :param NetworkPrivateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkPrivateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 vlan_id: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkPrivateArgs.__new__(NetworkPrivateArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["regions"] = regions
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["vlan_id"] = vlan_id
            __props__.__dict__["regions_attributes"] = None
            __props__.__dict__["regions_openstack_ids"] = None
            __props__.__dict__["regions_statuses"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
        super(NetworkPrivate, __self__).__init__(
            'ovh:CloudProject/networkPrivate:NetworkPrivate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            regions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            regions_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateRegionsAttributeArgs', 'NetworkPrivateRegionsAttributeArgsDict']]]]] = None,
            regions_openstack_ids: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            regions_statuses: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateRegionsStatusArgs', 'NetworkPrivateRegionsStatusArgsDict']]]]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None,
            vlan_id: Optional[pulumi.Input[builtins.int]] = None) -> 'NetworkPrivate':
        """
        Get an existing NetworkPrivate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: The name of the network.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] regions: an array of valid OVHcloud public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        :param pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateRegionsAttributeArgs', 'NetworkPrivateRegionsAttributeArgsDict']]]] regions_attributes: A map representing information about the region.
               * `regions_attributes/region` - The id of the region.
               * `regions_attributes/status` - The status of the network in the region.
               * `regions_attributes/openstackid` - The private network id in the region.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] regions_openstack_ids: A map with region name as key, and region-specific openstack id as value
        :param pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateRegionsStatusArgs', 'NetworkPrivateRegionsStatusArgsDict']]]] regions_statuses: (Deprecated) A map representing the status of the network per region.
               * `regions_status/region` - (Deprecated) The id of the region.
               * `regions_status/status` - (Deprecated) The status of the network in the region.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] status: the status of the network. should be normally set to 'ACTIVE'.
        :param pulumi.Input[builtins.str] type: the type of the network. Either 'private' or 'public'.
        :param pulumi.Input[builtins.int] vlan_id: a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkPrivateState.__new__(_NetworkPrivateState)

        __props__.__dict__["name"] = name
        __props__.__dict__["regions"] = regions
        __props__.__dict__["regions_attributes"] = regions_attributes
        __props__.__dict__["regions_openstack_ids"] = regions_openstack_ids
        __props__.__dict__["regions_statuses"] = regions_statuses
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        __props__.__dict__["vlan_id"] = vlan_id
        return NetworkPrivate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the network.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        an array of valid OVHcloud public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="regionsAttributes")
    def regions_attributes(self) -> pulumi.Output[Sequence['outputs.NetworkPrivateRegionsAttribute']]:
        """
        A map representing information about the region.
        * `regions_attributes/region` - The id of the region.
        * `regions_attributes/status` - The status of the network in the region.
        * `regions_attributes/openstackid` - The private network id in the region.
        """
        return pulumi.get(self, "regions_attributes")

    @property
    @pulumi.getter(name="regionsOpenstackIds")
    def regions_openstack_ids(self) -> pulumi.Output[Mapping[str, builtins.str]]:
        """
        A map with region name as key, and region-specific openstack id as value
        """
        return pulumi.get(self, "regions_openstack_ids")

    @property
    @pulumi.getter(name="regionsStatuses")
    @_utilities.deprecated("""use the regions_attributes field instead""")
    def regions_statuses(self) -> pulumi.Output[Sequence['outputs.NetworkPrivateRegionsStatus']]:
        """
        (Deprecated) A map representing the status of the network per region.
        * `regions_status/region` - (Deprecated) The id of the region.
        * `regions_status/status` - (Deprecated) The status of the network in the region.
        """
        return pulumi.get(self, "regions_statuses")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        the status of the network. should be normally set to 'ACTIVE'.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        the type of the network. Either 'private' or 'public'.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.
        """
        return pulumi.get(self, "vlan_id")

