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
from . import outputs
from ._inputs import *

__all__ = ['TcpFarmArgs', 'TcpFarm']

@pulumi.input_type
class TcpFarmArgs:
    def __init__(__self__, *,
                 service_name: pulumi.Input[builtins.str],
                 zone: pulumi.Input[builtins.str],
                 balance: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 probe: Optional[pulumi.Input['TcpFarmProbeArgs']] = None,
                 stickiness: Optional[pulumi.Input[builtins.str]] = None,
                 vrack_network_id: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a TcpFarm resource.
        :param pulumi.Input[builtins.str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[builtins.str] zone: Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        :param pulumi.Input[builtins.str] balance: Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        :param pulumi.Input[builtins.str] display_name: Readable label for loadbalancer farm
        :param pulumi.Input[builtins.int] port: Port attached to your farm ([1..49151]). Inherited from frontend if null
        :param pulumi.Input['TcpFarmProbeArgs'] probe: define a backend healthcheck probe
        :param pulumi.Input[builtins.str] stickiness: Stickiness type. No stickiness if null (`sourceIp`)
        :param pulumi.Input[builtins.int] vrack_network_id: Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        """
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "zone", zone)
        if balance is not None:
            pulumi.set(__self__, "balance", balance)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if probe is not None:
            pulumi.set(__self__, "probe", probe)
        if stickiness is not None:
            pulumi.set(__self__, "stickiness", stickiness)
        if vrack_network_id is not None:
            pulumi.set(__self__, "vrack_network_id", vrack_network_id)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[builtins.str]:
        """
        Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter
    def balance(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        """
        return pulumi.get(self, "balance")

    @balance.setter
    def balance(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "balance", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Readable label for loadbalancer farm
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Port attached to your farm ([1..49151]). Inherited from frontend if null
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def probe(self) -> Optional[pulumi.Input['TcpFarmProbeArgs']]:
        """
        define a backend healthcheck probe
        """
        return pulumi.get(self, "probe")

    @probe.setter
    def probe(self, value: Optional[pulumi.Input['TcpFarmProbeArgs']]):
        pulumi.set(self, "probe", value)

    @property
    @pulumi.getter
    def stickiness(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Stickiness type. No stickiness if null (`sourceIp`)
        """
        return pulumi.get(self, "stickiness")

    @stickiness.setter
    def stickiness(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "stickiness", value)

    @property
    @pulumi.getter(name="vrackNetworkId")
    def vrack_network_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        """
        return pulumi.get(self, "vrack_network_id")

    @vrack_network_id.setter
    def vrack_network_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "vrack_network_id", value)


@pulumi.input_type
class _TcpFarmState:
    def __init__(__self__, *,
                 balance: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 probe: Optional[pulumi.Input['TcpFarmProbeArgs']] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 stickiness: Optional[pulumi.Input[builtins.str]] = None,
                 vrack_network_id: Optional[pulumi.Input[builtins.int]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering TcpFarm resources.
        :param pulumi.Input[builtins.str] balance: Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        :param pulumi.Input[builtins.str] display_name: Readable label for loadbalancer farm
        :param pulumi.Input[builtins.int] port: Port attached to your farm ([1..49151]). Inherited from frontend if null
        :param pulumi.Input['TcpFarmProbeArgs'] probe: define a backend healthcheck probe
        :param pulumi.Input[builtins.str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[builtins.str] stickiness: Stickiness type. No stickiness if null (`sourceIp`)
        :param pulumi.Input[builtins.int] vrack_network_id: Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        :param pulumi.Input[builtins.str] zone: Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        """
        if balance is not None:
            pulumi.set(__self__, "balance", balance)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if probe is not None:
            pulumi.set(__self__, "probe", probe)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if stickiness is not None:
            pulumi.set(__self__, "stickiness", stickiness)
        if vrack_network_id is not None:
            pulumi.set(__self__, "vrack_network_id", vrack_network_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def balance(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        """
        return pulumi.get(self, "balance")

    @balance.setter
    def balance(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "balance", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Readable label for loadbalancer farm
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Port attached to your farm ([1..49151]). Inherited from frontend if null
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def probe(self) -> Optional[pulumi.Input['TcpFarmProbeArgs']]:
        """
        define a backend healthcheck probe
        """
        return pulumi.get(self, "probe")

    @probe.setter
    def probe(self, value: Optional[pulumi.Input['TcpFarmProbeArgs']]):
        pulumi.set(self, "probe", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def stickiness(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Stickiness type. No stickiness if null (`sourceIp`)
        """
        return pulumi.get(self, "stickiness")

    @stickiness.setter
    def stickiness(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "stickiness", value)

    @property
    @pulumi.getter(name="vrackNetworkId")
    def vrack_network_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        """
        return pulumi.get(self, "vrack_network_id")

    @vrack_network_id.setter
    def vrack_network_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "vrack_network_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone", value)


class TcpFarm(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 balance: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 probe: Optional[pulumi.Input[Union['TcpFarmProbeArgs', 'TcpFarmProbeArgsDict']]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 stickiness: Optional[pulumi.Input[builtins.str]] = None,
                 vrack_network_id: Optional[pulumi.Input[builtins.int]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates a backend server group (farm) to be used by loadbalancing frontend(s)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farm_name = ovh.ip_load_balancing.TcpFarm("farmName",
            display_name="ingress-8080-gra",
            service_name=lb.service_name,
            zone="GRA")
        ```

        ## Import

        TCP Farm can be imported using the following format `service_name` and the `id` of the farm, separated by "/" e.g.

        bash

        ```sh
        $ pulumi import ovh:IpLoadBalancing/tcpFarm:TcpFarm farmname service_name/farm_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] balance: Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        :param pulumi.Input[builtins.str] display_name: Readable label for loadbalancer farm
        :param pulumi.Input[builtins.int] port: Port attached to your farm ([1..49151]). Inherited from frontend if null
        :param pulumi.Input[Union['TcpFarmProbeArgs', 'TcpFarmProbeArgsDict']] probe: define a backend healthcheck probe
        :param pulumi.Input[builtins.str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[builtins.str] stickiness: Stickiness type. No stickiness if null (`sourceIp`)
        :param pulumi.Input[builtins.int] vrack_network_id: Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        :param pulumi.Input[builtins.str] zone: Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TcpFarmArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a backend server group (farm) to be used by loadbalancing frontend(s)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farm_name = ovh.ip_load_balancing.TcpFarm("farmName",
            display_name="ingress-8080-gra",
            service_name=lb.service_name,
            zone="GRA")
        ```

        ## Import

        TCP Farm can be imported using the following format `service_name` and the `id` of the farm, separated by "/" e.g.

        bash

        ```sh
        $ pulumi import ovh:IpLoadBalancing/tcpFarm:TcpFarm farmname service_name/farm_id
        ```

        :param str resource_name: The name of the resource.
        :param TcpFarmArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TcpFarmArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 balance: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 probe: Optional[pulumi.Input[Union['TcpFarmProbeArgs', 'TcpFarmProbeArgsDict']]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 stickiness: Optional[pulumi.Input[builtins.str]] = None,
                 vrack_network_id: Optional[pulumi.Input[builtins.int]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TcpFarmArgs.__new__(TcpFarmArgs)

            __props__.__dict__["balance"] = balance
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["port"] = port
            __props__.__dict__["probe"] = probe
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["stickiness"] = stickiness
            __props__.__dict__["vrack_network_id"] = vrack_network_id
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(TcpFarm, __self__).__init__(
            'ovh:IpLoadBalancing/tcpFarm:TcpFarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            balance: Optional[pulumi.Input[builtins.str]] = None,
            display_name: Optional[pulumi.Input[builtins.str]] = None,
            port: Optional[pulumi.Input[builtins.int]] = None,
            probe: Optional[pulumi.Input[Union['TcpFarmProbeArgs', 'TcpFarmProbeArgsDict']]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            stickiness: Optional[pulumi.Input[builtins.str]] = None,
            vrack_network_id: Optional[pulumi.Input[builtins.int]] = None,
            zone: Optional[pulumi.Input[builtins.str]] = None) -> 'TcpFarm':
        """
        Get an existing TcpFarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] balance: Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        :param pulumi.Input[builtins.str] display_name: Readable label for loadbalancer farm
        :param pulumi.Input[builtins.int] port: Port attached to your farm ([1..49151]). Inherited from frontend if null
        :param pulumi.Input[Union['TcpFarmProbeArgs', 'TcpFarmProbeArgsDict']] probe: define a backend healthcheck probe
        :param pulumi.Input[builtins.str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[builtins.str] stickiness: Stickiness type. No stickiness if null (`sourceIp`)
        :param pulumi.Input[builtins.int] vrack_network_id: Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        :param pulumi.Input[builtins.str] zone: Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TcpFarmState.__new__(_TcpFarmState)

        __props__.__dict__["balance"] = balance
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["port"] = port
        __props__.__dict__["probe"] = probe
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["stickiness"] = stickiness
        __props__.__dict__["vrack_network_id"] = vrack_network_id
        __props__.__dict__["zone"] = zone
        return TcpFarm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def balance(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        """
        return pulumi.get(self, "balance")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Readable label for loadbalancer farm
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Port attached to your farm ([1..49151]). Inherited from frontend if null
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def probe(self) -> pulumi.Output[Optional['outputs.TcpFarmProbe']]:
        """
        define a backend healthcheck probe
        """
        return pulumi.get(self, "probe")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def stickiness(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Stickiness type. No stickiness if null (`sourceIp`)
        """
        return pulumi.get(self, "stickiness")

    @property
    @pulumi.getter(name="vrackNetworkId")
    def vrack_network_id(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        """
        return pulumi.get(self, "vrack_network_id")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[builtins.str]:
        """
        Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        """
        return pulumi.get(self, "zone")

