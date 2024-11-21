# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['UdpFarmArgs', 'UdpFarm']

@pulumi.input_type
class UdpFarmArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[float],
                 service_name: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 vrack_network_id: Optional[pulumi.Input[float]] = None):
        """
        The set of arguments for constructing a UdpFarm resource.
        :param pulumi.Input[float] port: Port attached to your farm ([1..49151]). Inherited from frontend if null
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] zone: Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        :param pulumi.Input[str] display_name: Readable label for loadbalancer farm
        :param pulumi.Input[float] vrack_network_id: Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "zone", zone)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if vrack_network_id is not None:
            pulumi.set(__self__, "vrack_network_id", vrack_network_id)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[float]:
        """
        Port attached to your farm ([1..49151]). Inherited from frontend if null
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[float]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Readable label for loadbalancer farm
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="vrackNetworkId")
    def vrack_network_id(self) -> Optional[pulumi.Input[float]]:
        """
        Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        """
        return pulumi.get(self, "vrack_network_id")

    @vrack_network_id.setter
    def vrack_network_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "vrack_network_id", value)


@pulumi.input_type
class _UdpFarmState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_id: Optional[pulumi.Input[float]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 vrack_network_id: Optional[pulumi.Input[float]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UdpFarm resources.
        :param pulumi.Input[str] display_name: Readable label for loadbalancer farm
        :param pulumi.Input[float] farm_id: Id of your farm.
        :param pulumi.Input[float] port: Port attached to your farm ([1..49151]). Inherited from frontend if null
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[float] vrack_network_id: Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        :param pulumi.Input[str] zone: Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if farm_id is not None:
            pulumi.set(__self__, "farm_id", farm_id)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if vrack_network_id is not None:
            pulumi.set(__self__, "vrack_network_id", vrack_network_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Readable label for loadbalancer farm
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="farmId")
    def farm_id(self) -> Optional[pulumi.Input[float]]:
        """
        Id of your farm.
        """
        return pulumi.get(self, "farm_id")

    @farm_id.setter
    def farm_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "farm_id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[float]]:
        """
        Port attached to your farm ([1..49151]). Inherited from frontend if null
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="vrackNetworkId")
    def vrack_network_id(self) -> Optional[pulumi.Input[float]]:
        """
        Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        """
        return pulumi.get(self, "vrack_network_id")

    @vrack_network_id.setter
    def vrack_network_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "vrack_network_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class UdpFarm(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 vrack_network_id: Optional[pulumi.Input[float]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a backend server group (farm) to be used by loadbalancing frontend(s)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farm_name = ovh.ip_load_balancing.UdpFarm("farmName",
            display_name="ingress-8080-gra",
            port=80,
            service_name=lb.service_name,
            zone="gra")
        ```

        ## Import

        UDP Farm can be imported using the following format `service_name` and the `id` of the farm, separated by "/" e.g.

        bash

        ```sh
        $ pulumi import ovh:IpLoadBalancing/udpFarm:UdpFarm farmname service_name/farm_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Readable label for loadbalancer farm
        :param pulumi.Input[float] port: Port attached to your farm ([1..49151]). Inherited from frontend if null
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[float] vrack_network_id: Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        :param pulumi.Input[str] zone: Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UdpFarmArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a backend server group (farm) to be used by loadbalancing frontend(s)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farm_name = ovh.ip_load_balancing.UdpFarm("farmName",
            display_name="ingress-8080-gra",
            port=80,
            service_name=lb.service_name,
            zone="gra")
        ```

        ## Import

        UDP Farm can be imported using the following format `service_name` and the `id` of the farm, separated by "/" e.g.

        bash

        ```sh
        $ pulumi import ovh:IpLoadBalancing/udpFarm:UdpFarm farmname service_name/farm_id
        ```

        :param str resource_name: The name of the resource.
        :param UdpFarmArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UdpFarmArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 vrack_network_id: Optional[pulumi.Input[float]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UdpFarmArgs.__new__(UdpFarmArgs)

            __props__.__dict__["display_name"] = display_name
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["vrack_network_id"] = vrack_network_id
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["farm_id"] = None
        super(UdpFarm, __self__).__init__(
            'ovh:IpLoadBalancing/udpFarm:UdpFarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            farm_id: Optional[pulumi.Input[float]] = None,
            port: Optional[pulumi.Input[float]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            vrack_network_id: Optional[pulumi.Input[float]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'UdpFarm':
        """
        Get an existing UdpFarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Readable label for loadbalancer farm
        :param pulumi.Input[float] farm_id: Id of your farm.
        :param pulumi.Input[float] port: Port attached to your farm ([1..49151]). Inherited from frontend if null
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[float] vrack_network_id: Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        :param pulumi.Input[str] zone: Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UdpFarmState.__new__(_UdpFarmState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["farm_id"] = farm_id
        __props__.__dict__["port"] = port
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["vrack_network_id"] = vrack_network_id
        __props__.__dict__["zone"] = zone
        return UdpFarm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Readable label for loadbalancer farm
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="farmId")
    def farm_id(self) -> pulumi.Output[float]:
        """
        Id of your farm.
        """
        return pulumi.get(self, "farm_id")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[float]:
        """
        Port attached to your farm ([1..49151]). Inherited from frontend if null
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="vrackNetworkId")
    def vrack_network_id(self) -> pulumi.Output[Optional[float]]:
        """
        Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        """
        return pulumi.get(self, "vrack_network_id")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        """
        return pulumi.get(self, "zone")

