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
from . import outputs
from ._inputs import *

__all__ = ['NetworkPrivateSubnetV2Args', 'NetworkPrivateSubnetV2']

@pulumi.input_type
class NetworkPrivateSubnetV2Args:
    def __init__(__self__, *,
                 cidr: pulumi.Input[str],
                 network_id: pulumi.Input[str],
                 region: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 allocation_pools: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2AllocationPoolArgs']]]] = None,
                 dhcp: Optional[pulumi.Input[bool]] = None,
                 dns_nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_gateway_ip: Optional[pulumi.Input[bool]] = None,
                 gateway_ip: Optional[pulumi.Input[str]] = None,
                 host_routes: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2HostRouteArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 use_default_public_dns_resolver: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a NetworkPrivateSubnetV2 resource.
        :param pulumi.Input[str] cidr: IP range of the subnet
               Changing this value recreates the subnet.
        :param pulumi.Input[str] network_id: The id of the network.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] region: The region in which the network subnet will be created.
               Ex.: "GRA1". Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2AllocationPoolArgs']]] allocation_pools: DHCP allocation pools of subnet
        :param pulumi.Input[bool] dhcp: Enable DHCP.
               Changing this forces a new resource to be created. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_nameservers: DNS nameservers used by DHCP
               Changing this value recreates the resource. Defaults to OVH default DNS nameserver.
        :param pulumi.Input[bool] enable_gateway_ip: Set to true if you want to set a default gateway IP.
               Changing this value recreates the resource. Defaults to true.
        :param pulumi.Input[str] gateway_ip: See Argument Reference above.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2HostRouteArgs']]] host_routes: Static host routes of subnet
        :param pulumi.Input[str] name: Name of the subnet
               Changing this value recreates the subnet.
        :param pulumi.Input[bool] use_default_public_dns_resolver: Set to false if you want to use your DNS resolver.
               Changing this value recreates the resource.
        """
        pulumi.set(__self__, "cidr", cidr)
        pulumi.set(__self__, "network_id", network_id)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "service_name", service_name)
        if allocation_pools is not None:
            pulumi.set(__self__, "allocation_pools", allocation_pools)
        if dhcp is not None:
            pulumi.set(__self__, "dhcp", dhcp)
        if dns_nameservers is not None:
            pulumi.set(__self__, "dns_nameservers", dns_nameservers)
        if enable_gateway_ip is not None:
            pulumi.set(__self__, "enable_gateway_ip", enable_gateway_ip)
        if gateway_ip is not None:
            pulumi.set(__self__, "gateway_ip", gateway_ip)
        if host_routes is not None:
            pulumi.set(__self__, "host_routes", host_routes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if use_default_public_dns_resolver is not None:
            pulumi.set(__self__, "use_default_public_dns_resolver", use_default_public_dns_resolver)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Input[str]:
        """
        IP range of the subnet
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[str]:
        """
        The id of the network.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The region in which the network subnet will be created.
        Ex.: "GRA1". Changing this value recreates the resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="allocationPools")
    def allocation_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2AllocationPoolArgs']]]]:
        """
        DHCP allocation pools of subnet
        """
        return pulumi.get(self, "allocation_pools")

    @allocation_pools.setter
    def allocation_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2AllocationPoolArgs']]]]):
        pulumi.set(self, "allocation_pools", value)

    @property
    @pulumi.getter
    def dhcp(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable DHCP.
        Changing this forces a new resource to be created. Defaults to true.
        """
        return pulumi.get(self, "dhcp")

    @dhcp.setter
    def dhcp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dhcp", value)

    @property
    @pulumi.getter(name="dnsNameservers")
    def dns_nameservers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        DNS nameservers used by DHCP
        Changing this value recreates the resource. Defaults to OVH default DNS nameserver.
        """
        return pulumi.get(self, "dns_nameservers")

    @dns_nameservers.setter
    def dns_nameservers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_nameservers", value)

    @property
    @pulumi.getter(name="enableGatewayIp")
    def enable_gateway_ip(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true if you want to set a default gateway IP.
        Changing this value recreates the resource. Defaults to true.
        """
        return pulumi.get(self, "enable_gateway_ip")

    @enable_gateway_ip.setter
    def enable_gateway_ip(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_gateway_ip", value)

    @property
    @pulumi.getter(name="gatewayIp")
    def gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "gateway_ip")

    @gateway_ip.setter
    def gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_ip", value)

    @property
    @pulumi.getter(name="hostRoutes")
    def host_routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2HostRouteArgs']]]]:
        """
        Static host routes of subnet
        """
        return pulumi.get(self, "host_routes")

    @host_routes.setter
    def host_routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2HostRouteArgs']]]]):
        pulumi.set(self, "host_routes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the subnet
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="useDefaultPublicDnsResolver")
    def use_default_public_dns_resolver(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to false if you want to use your DNS resolver.
        Changing this value recreates the resource.
        """
        return pulumi.get(self, "use_default_public_dns_resolver")

    @use_default_public_dns_resolver.setter
    def use_default_public_dns_resolver(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_default_public_dns_resolver", value)


@pulumi.input_type
class _NetworkPrivateSubnetV2State:
    def __init__(__self__, *,
                 allocation_pools: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2AllocationPoolArgs']]]] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 dhcp: Optional[pulumi.Input[bool]] = None,
                 dns_nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_gateway_ip: Optional[pulumi.Input[bool]] = None,
                 gateway_ip: Optional[pulumi.Input[str]] = None,
                 host_routes: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2HostRouteArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 use_default_public_dns_resolver: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering NetworkPrivateSubnetV2 resources.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2AllocationPoolArgs']]] allocation_pools: DHCP allocation pools of subnet
        :param pulumi.Input[str] cidr: IP range of the subnet
               Changing this value recreates the subnet.
        :param pulumi.Input[bool] dhcp: Enable DHCP.
               Changing this forces a new resource to be created. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_nameservers: DNS nameservers used by DHCP
               Changing this value recreates the resource. Defaults to OVH default DNS nameserver.
        :param pulumi.Input[bool] enable_gateway_ip: Set to true if you want to set a default gateway IP.
               Changing this value recreates the resource. Defaults to true.
        :param pulumi.Input[str] gateway_ip: See Argument Reference above.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2HostRouteArgs']]] host_routes: Static host routes of subnet
        :param pulumi.Input[str] name: Name of the subnet
               Changing this value recreates the subnet.
        :param pulumi.Input[str] network_id: The id of the network.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] region: The region in which the network subnet will be created.
               Ex.: "GRA1". Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[bool] use_default_public_dns_resolver: Set to false if you want to use your DNS resolver.
               Changing this value recreates the resource.
        """
        if allocation_pools is not None:
            pulumi.set(__self__, "allocation_pools", allocation_pools)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if dhcp is not None:
            pulumi.set(__self__, "dhcp", dhcp)
        if dns_nameservers is not None:
            pulumi.set(__self__, "dns_nameservers", dns_nameservers)
        if enable_gateway_ip is not None:
            pulumi.set(__self__, "enable_gateway_ip", enable_gateway_ip)
        if gateway_ip is not None:
            pulumi.set(__self__, "gateway_ip", gateway_ip)
        if host_routes is not None:
            pulumi.set(__self__, "host_routes", host_routes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if use_default_public_dns_resolver is not None:
            pulumi.set(__self__, "use_default_public_dns_resolver", use_default_public_dns_resolver)

    @property
    @pulumi.getter(name="allocationPools")
    def allocation_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2AllocationPoolArgs']]]]:
        """
        DHCP allocation pools of subnet
        """
        return pulumi.get(self, "allocation_pools")

    @allocation_pools.setter
    def allocation_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2AllocationPoolArgs']]]]):
        pulumi.set(self, "allocation_pools", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        IP range of the subnet
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def dhcp(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable DHCP.
        Changing this forces a new resource to be created. Defaults to true.
        """
        return pulumi.get(self, "dhcp")

    @dhcp.setter
    def dhcp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dhcp", value)

    @property
    @pulumi.getter(name="dnsNameservers")
    def dns_nameservers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        DNS nameservers used by DHCP
        Changing this value recreates the resource. Defaults to OVH default DNS nameserver.
        """
        return pulumi.get(self, "dns_nameservers")

    @dns_nameservers.setter
    def dns_nameservers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_nameservers", value)

    @property
    @pulumi.getter(name="enableGatewayIp")
    def enable_gateway_ip(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true if you want to set a default gateway IP.
        Changing this value recreates the resource. Defaults to true.
        """
        return pulumi.get(self, "enable_gateway_ip")

    @enable_gateway_ip.setter
    def enable_gateway_ip(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_gateway_ip", value)

    @property
    @pulumi.getter(name="gatewayIp")
    def gateway_ip(self) -> Optional[pulumi.Input[str]]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "gateway_ip")

    @gateway_ip.setter
    def gateway_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_ip", value)

    @property
    @pulumi.getter(name="hostRoutes")
    def host_routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2HostRouteArgs']]]]:
        """
        Static host routes of subnet
        """
        return pulumi.get(self, "host_routes")

    @host_routes.setter
    def host_routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkPrivateSubnetV2HostRouteArgs']]]]):
        pulumi.set(self, "host_routes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the subnet
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the network.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which the network subnet will be created.
        Ex.: "GRA1". Changing this value recreates the resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="useDefaultPublicDnsResolver")
    def use_default_public_dns_resolver(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to false if you want to use your DNS resolver.
        Changing this value recreates the resource.
        """
        return pulumi.get(self, "use_default_public_dns_resolver")

    @use_default_public_dns_resolver.setter
    def use_default_public_dns_resolver(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_default_public_dns_resolver", value)


class NetworkPrivateSubnetV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_pools: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2AllocationPoolArgs', 'NetworkPrivateSubnetV2AllocationPoolArgsDict']]]]] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 dhcp: Optional[pulumi.Input[bool]] = None,
                 dns_nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_gateway_ip: Optional[pulumi.Input[bool]] = None,
                 gateway_ip: Optional[pulumi.Input[str]] = None,
                 host_routes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2HostRouteArgs', 'NetworkPrivateSubnetV2HostRouteArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 use_default_public_dns_resolver: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Creates a subnet in a private network of a public cloud region.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        subnet = ovh.cloud_project.NetworkPrivateSubnetV2("subnet",
            cidr="192.168.168.0/24",
            dhcp=True,
            dns_nameservers=["1.1.1.1"],
            enable_gateway_ip=True,
            network_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
            region="XXX1",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            use_default_public_dns_resolver=False)
        ```

        ## Import

        Subnet in a private network of a public cloud project can be imported using the `service_name`, `region`, `network_id` and `subnet_id`, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/networkPrivateSubnetV2:NetworkPrivateSubnetV2 mysubnet 5ceb661434891538b54a4f2c66fc4b746e/BHS5/25807101-8aaa-4ea5-b507-61f0d661b101/0f0b73a4-403b-45e4-86d0-b438f1291909
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2AllocationPoolArgs', 'NetworkPrivateSubnetV2AllocationPoolArgsDict']]]] allocation_pools: DHCP allocation pools of subnet
        :param pulumi.Input[str] cidr: IP range of the subnet
               Changing this value recreates the subnet.
        :param pulumi.Input[bool] dhcp: Enable DHCP.
               Changing this forces a new resource to be created. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_nameservers: DNS nameservers used by DHCP
               Changing this value recreates the resource. Defaults to OVH default DNS nameserver.
        :param pulumi.Input[bool] enable_gateway_ip: Set to true if you want to set a default gateway IP.
               Changing this value recreates the resource. Defaults to true.
        :param pulumi.Input[str] gateway_ip: See Argument Reference above.
        :param pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2HostRouteArgs', 'NetworkPrivateSubnetV2HostRouteArgsDict']]]] host_routes: Static host routes of subnet
        :param pulumi.Input[str] name: Name of the subnet
               Changing this value recreates the subnet.
        :param pulumi.Input[str] network_id: The id of the network.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] region: The region in which the network subnet will be created.
               Ex.: "GRA1". Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[bool] use_default_public_dns_resolver: Set to false if you want to use your DNS resolver.
               Changing this value recreates the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkPrivateSubnetV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a subnet in a private network of a public cloud region.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        subnet = ovh.cloud_project.NetworkPrivateSubnetV2("subnet",
            cidr="192.168.168.0/24",
            dhcp=True,
            dns_nameservers=["1.1.1.1"],
            enable_gateway_ip=True,
            network_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
            region="XXX1",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            use_default_public_dns_resolver=False)
        ```

        ## Import

        Subnet in a private network of a public cloud project can be imported using the `service_name`, `region`, `network_id` and `subnet_id`, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/networkPrivateSubnetV2:NetworkPrivateSubnetV2 mysubnet 5ceb661434891538b54a4f2c66fc4b746e/BHS5/25807101-8aaa-4ea5-b507-61f0d661b101/0f0b73a4-403b-45e4-86d0-b438f1291909
        ```

        :param str resource_name: The name of the resource.
        :param NetworkPrivateSubnetV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkPrivateSubnetV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_pools: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2AllocationPoolArgs', 'NetworkPrivateSubnetV2AllocationPoolArgsDict']]]]] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 dhcp: Optional[pulumi.Input[bool]] = None,
                 dns_nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_gateway_ip: Optional[pulumi.Input[bool]] = None,
                 gateway_ip: Optional[pulumi.Input[str]] = None,
                 host_routes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2HostRouteArgs', 'NetworkPrivateSubnetV2HostRouteArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 use_default_public_dns_resolver: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkPrivateSubnetV2Args.__new__(NetworkPrivateSubnetV2Args)

            __props__.__dict__["allocation_pools"] = allocation_pools
            if cidr is None and not opts.urn:
                raise TypeError("Missing required property 'cidr'")
            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["dhcp"] = dhcp
            __props__.__dict__["dns_nameservers"] = dns_nameservers
            __props__.__dict__["enable_gateway_ip"] = enable_gateway_ip
            __props__.__dict__["gateway_ip"] = gateway_ip
            __props__.__dict__["host_routes"] = host_routes
            __props__.__dict__["name"] = name
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["use_default_public_dns_resolver"] = use_default_public_dns_resolver
        super(NetworkPrivateSubnetV2, __self__).__init__(
            'ovh:CloudProject/networkPrivateSubnetV2:NetworkPrivateSubnetV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_pools: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2AllocationPoolArgs', 'NetworkPrivateSubnetV2AllocationPoolArgsDict']]]]] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            dhcp: Optional[pulumi.Input[bool]] = None,
            dns_nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            enable_gateway_ip: Optional[pulumi.Input[bool]] = None,
            gateway_ip: Optional[pulumi.Input[str]] = None,
            host_routes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2HostRouteArgs', 'NetworkPrivateSubnetV2HostRouteArgsDict']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            use_default_public_dns_resolver: Optional[pulumi.Input[bool]] = None) -> 'NetworkPrivateSubnetV2':
        """
        Get an existing NetworkPrivateSubnetV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2AllocationPoolArgs', 'NetworkPrivateSubnetV2AllocationPoolArgsDict']]]] allocation_pools: DHCP allocation pools of subnet
        :param pulumi.Input[str] cidr: IP range of the subnet
               Changing this value recreates the subnet.
        :param pulumi.Input[bool] dhcp: Enable DHCP.
               Changing this forces a new resource to be created. Defaults to true.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_nameservers: DNS nameservers used by DHCP
               Changing this value recreates the resource. Defaults to OVH default DNS nameserver.
        :param pulumi.Input[bool] enable_gateway_ip: Set to true if you want to set a default gateway IP.
               Changing this value recreates the resource. Defaults to true.
        :param pulumi.Input[str] gateway_ip: See Argument Reference above.
        :param pulumi.Input[Sequence[pulumi.Input[Union['NetworkPrivateSubnetV2HostRouteArgs', 'NetworkPrivateSubnetV2HostRouteArgsDict']]]] host_routes: Static host routes of subnet
        :param pulumi.Input[str] name: Name of the subnet
               Changing this value recreates the subnet.
        :param pulumi.Input[str] network_id: The id of the network.
               Changing this forces a new resource to be created.
        :param pulumi.Input[str] region: The region in which the network subnet will be created.
               Ex.: "GRA1". Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[bool] use_default_public_dns_resolver: Set to false if you want to use your DNS resolver.
               Changing this value recreates the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkPrivateSubnetV2State.__new__(_NetworkPrivateSubnetV2State)

        __props__.__dict__["allocation_pools"] = allocation_pools
        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["dhcp"] = dhcp
        __props__.__dict__["dns_nameservers"] = dns_nameservers
        __props__.__dict__["enable_gateway_ip"] = enable_gateway_ip
        __props__.__dict__["gateway_ip"] = gateway_ip
        __props__.__dict__["host_routes"] = host_routes
        __props__.__dict__["name"] = name
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["region"] = region
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["use_default_public_dns_resolver"] = use_default_public_dns_resolver
        return NetworkPrivateSubnetV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationPools")
    def allocation_pools(self) -> pulumi.Output[Sequence['outputs.NetworkPrivateSubnetV2AllocationPool']]:
        """
        DHCP allocation pools of subnet
        """
        return pulumi.get(self, "allocation_pools")

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        IP range of the subnet
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def dhcp(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable DHCP.
        Changing this forces a new resource to be created. Defaults to true.
        """
        return pulumi.get(self, "dhcp")

    @property
    @pulumi.getter(name="dnsNameservers")
    def dns_nameservers(self) -> pulumi.Output[Sequence[str]]:
        """
        DNS nameservers used by DHCP
        Changing this value recreates the resource. Defaults to OVH default DNS nameserver.
        """
        return pulumi.get(self, "dns_nameservers")

    @property
    @pulumi.getter(name="enableGatewayIp")
    def enable_gateway_ip(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to true if you want to set a default gateway IP.
        Changing this value recreates the resource. Defaults to true.
        """
        return pulumi.get(self, "enable_gateway_ip")

    @property
    @pulumi.getter(name="gatewayIp")
    def gateway_ip(self) -> pulumi.Output[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "gateway_ip")

    @property
    @pulumi.getter(name="hostRoutes")
    def host_routes(self) -> pulumi.Output[Optional[Sequence['outputs.NetworkPrivateSubnetV2HostRoute']]]:
        """
        Static host routes of subnet
        """
        return pulumi.get(self, "host_routes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the subnet
        Changing this value recreates the subnet.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        The id of the network.
        Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which the network subnet will be created.
        Ex.: "GRA1". Changing this value recreates the resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="useDefaultPublicDnsResolver")
    def use_default_public_dns_resolver(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to false if you want to use your DNS resolver.
        Changing this value recreates the resource.
        """
        return pulumi.get(self, "use_default_public_dns_resolver")

