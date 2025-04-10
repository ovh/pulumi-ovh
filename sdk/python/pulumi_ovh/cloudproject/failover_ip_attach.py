# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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

__all__ = ['FailoverIpAttachArgs', 'FailoverIpAttach']

@pulumi.input_type
class FailoverIpAttachArgs:
    def __init__(__self__, *,
                 service_name: pulumi.Input[builtins.str],
                 block: Optional[pulumi.Input[builtins.str]] = None,
                 continent_code: Optional[pulumi.Input[builtins.str]] = None,
                 geo_loc: Optional[pulumi.Input[builtins.str]] = None,
                 ip: Optional[pulumi.Input[builtins.str]] = None,
                 routed_to: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a FailoverIpAttach resource.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] block: The IP block
               * `continentCode` - The Ip continent
        :param pulumi.Input[builtins.str] continent_code: Ip continent
        :param pulumi.Input[builtins.str] geo_loc: Ip location
        :param pulumi.Input[builtins.str] ip: The failover ip address to attach
        :param pulumi.Input[builtins.str] routed_to: The GUID of an instance to which the failover IP address is be attached
        """
        pulumi.set(__self__, "service_name", service_name)
        if block is not None:
            pulumi.set(__self__, "block", block)
        if continent_code is not None:
            pulumi.set(__self__, "continent_code", continent_code)
        if geo_loc is not None:
            pulumi.set(__self__, "geo_loc", geo_loc)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if routed_to is not None:
            pulumi.set(__self__, "routed_to", routed_to)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def block(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The IP block
        * `continentCode` - The Ip continent
        """
        return pulumi.get(self, "block")

    @block.setter
    def block(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "block", value)

    @property
    @pulumi.getter(name="continentCode")
    def continent_code(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Ip continent
        """
        return pulumi.get(self, "continent_code")

    @continent_code.setter
    def continent_code(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "continent_code", value)

    @property
    @pulumi.getter(name="geoLoc")
    def geo_loc(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Ip location
        """
        return pulumi.get(self, "geo_loc")

    @geo_loc.setter
    def geo_loc(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "geo_loc", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The failover ip address to attach
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="routedTo")
    def routed_to(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The GUID of an instance to which the failover IP address is be attached
        """
        return pulumi.get(self, "routed_to")

    @routed_to.setter
    def routed_to(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "routed_to", value)


@pulumi.input_type
class _FailoverIpAttachState:
    def __init__(__self__, *,
                 block: Optional[pulumi.Input[builtins.str]] = None,
                 continent_code: Optional[pulumi.Input[builtins.str]] = None,
                 geo_loc: Optional[pulumi.Input[builtins.str]] = None,
                 ip: Optional[pulumi.Input[builtins.str]] = None,
                 progress: Optional[pulumi.Input[builtins.int]] = None,
                 routed_to: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 sub_type: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering FailoverIpAttach resources.
        :param pulumi.Input[builtins.str] block: The IP block
               * `continentCode` - The Ip continent
        :param pulumi.Input[builtins.str] continent_code: Ip continent
        :param pulumi.Input[builtins.str] geo_loc: Ip location
        :param pulumi.Input[builtins.str] ip: The failover ip address to attach
        :param pulumi.Input[builtins.int] progress: Current operation progress in percent
               * `routedTo` - Instance where ip is routed to
        :param pulumi.Input[builtins.str] routed_to: The GUID of an instance to which the failover IP address is be attached
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] status: Ip status, can be `ok` or `operationPending`
               * `subType` - IP sub type, can be `cloud` or `ovh`
        :param pulumi.Input[builtins.str] sub_type: IP sub type
        """
        if block is not None:
            pulumi.set(__self__, "block", block)
        if continent_code is not None:
            pulumi.set(__self__, "continent_code", continent_code)
        if geo_loc is not None:
            pulumi.set(__self__, "geo_loc", geo_loc)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if progress is not None:
            pulumi.set(__self__, "progress", progress)
        if routed_to is not None:
            pulumi.set(__self__, "routed_to", routed_to)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if sub_type is not None:
            pulumi.set(__self__, "sub_type", sub_type)

    @property
    @pulumi.getter
    def block(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The IP block
        * `continentCode` - The Ip continent
        """
        return pulumi.get(self, "block")

    @block.setter
    def block(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "block", value)

    @property
    @pulumi.getter(name="continentCode")
    def continent_code(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Ip continent
        """
        return pulumi.get(self, "continent_code")

    @continent_code.setter
    def continent_code(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "continent_code", value)

    @property
    @pulumi.getter(name="geoLoc")
    def geo_loc(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Ip location
        """
        return pulumi.get(self, "geo_loc")

    @geo_loc.setter
    def geo_loc(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "geo_loc", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The failover ip address to attach
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def progress(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Current operation progress in percent
        * `routedTo` - Instance where ip is routed to
        """
        return pulumi.get(self, "progress")

    @progress.setter
    def progress(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "progress", value)

    @property
    @pulumi.getter(name="routedTo")
    def routed_to(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The GUID of an instance to which the failover IP address is be attached
        """
        return pulumi.get(self, "routed_to")

    @routed_to.setter
    def routed_to(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "routed_to", value)

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
        Ip status, can be `ok` or `operationPending`
        * `subType` - IP sub type, can be `cloud` or `ovh`
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subType")
    def sub_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        IP sub type
        """
        return pulumi.get(self, "sub_type")

    @sub_type.setter
    def sub_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "sub_type", value)


class FailoverIpAttach(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block: Optional[pulumi.Input[builtins.str]] = None,
                 continent_code: Optional[pulumi.Input[builtins.str]] = None,
                 geo_loc: Optional[pulumi.Input[builtins.str]] = None,
                 ip: Optional[pulumi.Input[builtins.str]] = None,
                 routed_to: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Attaches a failover IP address to a compute instance

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_failover_ip = ovh.cloud_project.FailoverIpAttach("myFailoverIp",
            ip="XXXXXX",
            routed_to="XXXXXX",
            service_name="XXXXXX")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] block: The IP block
               * `continentCode` - The Ip continent
        :param pulumi.Input[builtins.str] continent_code: Ip continent
        :param pulumi.Input[builtins.str] geo_loc: Ip location
        :param pulumi.Input[builtins.str] ip: The failover ip address to attach
        :param pulumi.Input[builtins.str] routed_to: The GUID of an instance to which the failover IP address is be attached
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FailoverIpAttachArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attaches a failover IP address to a compute instance

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_failover_ip = ovh.cloud_project.FailoverIpAttach("myFailoverIp",
            ip="XXXXXX",
            routed_to="XXXXXX",
            service_name="XXXXXX")
        ```

        :param str resource_name: The name of the resource.
        :param FailoverIpAttachArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FailoverIpAttachArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block: Optional[pulumi.Input[builtins.str]] = None,
                 continent_code: Optional[pulumi.Input[builtins.str]] = None,
                 geo_loc: Optional[pulumi.Input[builtins.str]] = None,
                 ip: Optional[pulumi.Input[builtins.str]] = None,
                 routed_to: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FailoverIpAttachArgs.__new__(FailoverIpAttachArgs)

            __props__.__dict__["block"] = block
            __props__.__dict__["continent_code"] = continent_code
            __props__.__dict__["geo_loc"] = geo_loc
            __props__.__dict__["ip"] = ip
            __props__.__dict__["routed_to"] = routed_to
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["progress"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["sub_type"] = None
        super(FailoverIpAttach, __self__).__init__(
            'ovh:CloudProject/failoverIpAttach:FailoverIpAttach',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            block: Optional[pulumi.Input[builtins.str]] = None,
            continent_code: Optional[pulumi.Input[builtins.str]] = None,
            geo_loc: Optional[pulumi.Input[builtins.str]] = None,
            ip: Optional[pulumi.Input[builtins.str]] = None,
            progress: Optional[pulumi.Input[builtins.int]] = None,
            routed_to: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            sub_type: Optional[pulumi.Input[builtins.str]] = None) -> 'FailoverIpAttach':
        """
        Get an existing FailoverIpAttach resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] block: The IP block
               * `continentCode` - The Ip continent
        :param pulumi.Input[builtins.str] continent_code: Ip continent
        :param pulumi.Input[builtins.str] geo_loc: Ip location
        :param pulumi.Input[builtins.str] ip: The failover ip address to attach
        :param pulumi.Input[builtins.int] progress: Current operation progress in percent
               * `routedTo` - Instance where ip is routed to
        :param pulumi.Input[builtins.str] routed_to: The GUID of an instance to which the failover IP address is be attached
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] status: Ip status, can be `ok` or `operationPending`
               * `subType` - IP sub type, can be `cloud` or `ovh`
        :param pulumi.Input[builtins.str] sub_type: IP sub type
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FailoverIpAttachState.__new__(_FailoverIpAttachState)

        __props__.__dict__["block"] = block
        __props__.__dict__["continent_code"] = continent_code
        __props__.__dict__["geo_loc"] = geo_loc
        __props__.__dict__["ip"] = ip
        __props__.__dict__["progress"] = progress
        __props__.__dict__["routed_to"] = routed_to
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        __props__.__dict__["sub_type"] = sub_type
        return FailoverIpAttach(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def block(self) -> pulumi.Output[builtins.str]:
        """
        The IP block
        * `continentCode` - The Ip continent
        """
        return pulumi.get(self, "block")

    @property
    @pulumi.getter(name="continentCode")
    def continent_code(self) -> pulumi.Output[builtins.str]:
        """
        Ip continent
        """
        return pulumi.get(self, "continent_code")

    @property
    @pulumi.getter(name="geoLoc")
    def geo_loc(self) -> pulumi.Output[builtins.str]:
        """
        Ip location
        """
        return pulumi.get(self, "geo_loc")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[builtins.str]:
        """
        The failover ip address to attach
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def progress(self) -> pulumi.Output[builtins.int]:
        """
        Current operation progress in percent
        * `routedTo` - Instance where ip is routed to
        """
        return pulumi.get(self, "progress")

    @property
    @pulumi.getter(name="routedTo")
    def routed_to(self) -> pulumi.Output[builtins.str]:
        """
        The GUID of an instance to which the failover IP address is be attached
        """
        return pulumi.get(self, "routed_to")

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
        Ip status, can be `ok` or `operationPending`
        * `subType` - IP sub type, can be `cloud` or `ovh`
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subType")
    def sub_type(self) -> pulumi.Output[builtins.str]:
        """
        IP sub type
        """
        return pulumi.get(self, "sub_type")

