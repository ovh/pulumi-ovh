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
from . import _utilities

__all__ = ['OvhcloudConnectPopDatacenterConfigArgs', 'OvhcloudConnectPopDatacenterConfig']

@pulumi.input_type
class OvhcloudConnectPopDatacenterConfigArgs:
    def __init__(__self__, *,
                 config_pop_id: pulumi.Input[builtins.float],
                 datacenter_id: pulumi.Input[builtins.float],
                 service_name: pulumi.Input[builtins.str],
                 ovh_bgp_area: Optional[pulumi.Input[builtins.float]] = None,
                 subnet: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a OvhcloudConnectPopDatacenterConfig resource.
        :param pulumi.Input[builtins.float] config_pop_id: ID of the pop configuration
        :param pulumi.Input[builtins.float] datacenter_id: ID of the datacenter linked
        :param pulumi.Input[builtins.str] service_name: Service name
        :param pulumi.Input[builtins.float] ovh_bgp_area: OVH Private AS
        :param pulumi.Input[builtins.str] subnet: Subnet should be a /28 min
        """
        pulumi.set(__self__, "config_pop_id", config_pop_id)
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        pulumi.set(__self__, "service_name", service_name)
        if ovh_bgp_area is not None:
            pulumi.set(__self__, "ovh_bgp_area", ovh_bgp_area)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)

    @property
    @pulumi.getter(name="configPopId")
    def config_pop_id(self) -> pulumi.Input[builtins.float]:
        """
        ID of the pop configuration
        """
        return pulumi.get(self, "config_pop_id")

    @config_pop_id.setter
    def config_pop_id(self, value: pulumi.Input[builtins.float]):
        pulumi.set(self, "config_pop_id", value)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Input[builtins.float]:
        """
        ID of the datacenter linked
        """
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: pulumi.Input[builtins.float]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="ovhBgpArea")
    def ovh_bgp_area(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        OVH Private AS
        """
        return pulumi.get(self, "ovh_bgp_area")

    @ovh_bgp_area.setter
    def ovh_bgp_area(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "ovh_bgp_area", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Subnet should be a /28 min
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subnet", value)


@pulumi.input_type
class _OvhcloudConnectPopDatacenterConfigState:
    def __init__(__self__, *,
                 config_pop_id: Optional[pulumi.Input[builtins.float]] = None,
                 datacenter_id: Optional[pulumi.Input[builtins.float]] = None,
                 ovh_bgp_area: Optional[pulumi.Input[builtins.float]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 subnet: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering OvhcloudConnectPopDatacenterConfig resources.
        :param pulumi.Input[builtins.float] config_pop_id: ID of the pop configuration
        :param pulumi.Input[builtins.float] datacenter_id: ID of the datacenter linked
        :param pulumi.Input[builtins.float] ovh_bgp_area: OVH Private AS
        :param pulumi.Input[builtins.str] service_name: Service name
        :param pulumi.Input[builtins.str] status: Status of the pop configuration
        :param pulumi.Input[builtins.str] subnet: Subnet should be a /28 min
        """
        if config_pop_id is not None:
            pulumi.set(__self__, "config_pop_id", config_pop_id)
        if datacenter_id is not None:
            pulumi.set(__self__, "datacenter_id", datacenter_id)
        if ovh_bgp_area is not None:
            pulumi.set(__self__, "ovh_bgp_area", ovh_bgp_area)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)

    @property
    @pulumi.getter(name="configPopId")
    def config_pop_id(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        ID of the pop configuration
        """
        return pulumi.get(self, "config_pop_id")

    @config_pop_id.setter
    def config_pop_id(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "config_pop_id", value)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        ID of the datacenter linked
        """
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter(name="ovhBgpArea")
    def ovh_bgp_area(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        OVH Private AS
        """
        return pulumi.get(self, "ovh_bgp_area")

    @ovh_bgp_area.setter
    def ovh_bgp_area(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "ovh_bgp_area", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Status of the pop configuration
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Subnet should be a /28 min
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subnet", value)


@pulumi.type_token("ovh:index/ovhcloudConnectPopDatacenterConfig:OvhcloudConnectPopDatacenterConfig")
class OvhcloudConnectPopDatacenterConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_pop_id: Optional[pulumi.Input[builtins.float]] = None,
                 datacenter_id: Optional[pulumi.Input[builtins.float]] = None,
                 ovh_bgp_area: Optional[pulumi.Input[builtins.float]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 subnet: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates a datacenter configuration for an Ovhcloud Connect product.

        For the `datacenter_id` in the `Required` section, you will need to choose an available datacenter from the data-source `get_ovhcloud_connect_datacenters`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.float] config_pop_id: ID of the pop configuration
        :param pulumi.Input[builtins.float] datacenter_id: ID of the datacenter linked
        :param pulumi.Input[builtins.float] ovh_bgp_area: OVH Private AS
        :param pulumi.Input[builtins.str] service_name: Service name
        :param pulumi.Input[builtins.str] subnet: Subnet should be a /28 min
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OvhcloudConnectPopDatacenterConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a datacenter configuration for an Ovhcloud Connect product.

        For the `datacenter_id` in the `Required` section, you will need to choose an available datacenter from the data-source `get_ovhcloud_connect_datacenters`.

        :param str resource_name: The name of the resource.
        :param OvhcloudConnectPopDatacenterConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OvhcloudConnectPopDatacenterConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_pop_id: Optional[pulumi.Input[builtins.float]] = None,
                 datacenter_id: Optional[pulumi.Input[builtins.float]] = None,
                 ovh_bgp_area: Optional[pulumi.Input[builtins.float]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 subnet: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OvhcloudConnectPopDatacenterConfigArgs.__new__(OvhcloudConnectPopDatacenterConfigArgs)

            if config_pop_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_pop_id'")
            __props__.__dict__["config_pop_id"] = config_pop_id
            if datacenter_id is None and not opts.urn:
                raise TypeError("Missing required property 'datacenter_id'")
            __props__.__dict__["datacenter_id"] = datacenter_id
            __props__.__dict__["ovh_bgp_area"] = ovh_bgp_area
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["subnet"] = subnet
            __props__.__dict__["status"] = None
        super(OvhcloudConnectPopDatacenterConfig, __self__).__init__(
            'ovh:index/ovhcloudConnectPopDatacenterConfig:OvhcloudConnectPopDatacenterConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_pop_id: Optional[pulumi.Input[builtins.float]] = None,
            datacenter_id: Optional[pulumi.Input[builtins.float]] = None,
            ovh_bgp_area: Optional[pulumi.Input[builtins.float]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            subnet: Optional[pulumi.Input[builtins.str]] = None) -> 'OvhcloudConnectPopDatacenterConfig':
        """
        Get an existing OvhcloudConnectPopDatacenterConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.float] config_pop_id: ID of the pop configuration
        :param pulumi.Input[builtins.float] datacenter_id: ID of the datacenter linked
        :param pulumi.Input[builtins.float] ovh_bgp_area: OVH Private AS
        :param pulumi.Input[builtins.str] service_name: Service name
        :param pulumi.Input[builtins.str] status: Status of the pop configuration
        :param pulumi.Input[builtins.str] subnet: Subnet should be a /28 min
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OvhcloudConnectPopDatacenterConfigState.__new__(_OvhcloudConnectPopDatacenterConfigState)

        __props__.__dict__["config_pop_id"] = config_pop_id
        __props__.__dict__["datacenter_id"] = datacenter_id
        __props__.__dict__["ovh_bgp_area"] = ovh_bgp_area
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        __props__.__dict__["subnet"] = subnet
        return OvhcloudConnectPopDatacenterConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configPopId")
    def config_pop_id(self) -> pulumi.Output[builtins.float]:
        """
        ID of the pop configuration
        """
        return pulumi.get(self, "config_pop_id")

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Output[builtins.float]:
        """
        ID of the datacenter linked
        """
        return pulumi.get(self, "datacenter_id")

    @property
    @pulumi.getter(name="ovhBgpArea")
    def ovh_bgp_area(self) -> pulumi.Output[builtins.float]:
        """
        OVH Private AS
        """
        return pulumi.get(self, "ovh_bgp_area")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Status of the pop configuration
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[builtins.str]:
        """
        Subnet should be a /28 min
        """
        return pulumi.get(self, "subnet")

