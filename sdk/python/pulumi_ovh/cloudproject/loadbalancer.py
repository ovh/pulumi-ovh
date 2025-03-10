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

__all__ = ['LoadbalancerArgs', 'Loadbalancer']

@pulumi.input_type
class LoadbalancerArgs:
    def __init__(__self__, *,
                 flavor_id: pulumi.Input[str],
                 network: pulumi.Input['LoadbalancerNetworkArgs'],
                 region_name: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 listeners: Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerListenerArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Loadbalancer resource.
        :param pulumi.Input[str] flavor_id: Loadbalancer flavor id
        :param pulumi.Input['LoadbalancerNetworkArgs'] network: Network information to create the loadbalancer
        :param pulumi.Input[str] region_name: Region name
        :param pulumi.Input[str] service_name: Service name
        :param pulumi.Input[str] description: Description of the loadbalancer
        :param pulumi.Input[Sequence[pulumi.Input['LoadbalancerListenerArgs']]] listeners: Listeners to create with the loadbalancer
        :param pulumi.Input[str] name: Name of the resource
        """
        pulumi.set(__self__, "flavor_id", flavor_id)
        pulumi.set(__self__, "network", network)
        pulumi.set(__self__, "region_name", region_name)
        pulumi.set(__self__, "service_name", service_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if listeners is not None:
            pulumi.set(__self__, "listeners", listeners)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> pulumi.Input[str]:
        """
        Loadbalancer flavor id
        """
        return pulumi.get(self, "flavor_id")

    @flavor_id.setter
    def flavor_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "flavor_id", value)

    @property
    @pulumi.getter
    def network(self) -> pulumi.Input['LoadbalancerNetworkArgs']:
        """
        Network information to create the loadbalancer
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: pulumi.Input['LoadbalancerNetworkArgs']):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Input[str]:
        """
        Region name
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "region_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the loadbalancer
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def listeners(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerListenerArgs']]]]:
        """
        Listeners to create with the loadbalancer
        """
        return pulumi.get(self, "listeners")

    @listeners.setter
    def listeners(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerListenerArgs']]]]):
        pulumi.set(self, "listeners", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _LoadbalancerState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 floating_ip: Optional[pulumi.Input['LoadbalancerFloatingIpArgs']] = None,
                 listeners: Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerListenerArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input['LoadbalancerNetworkArgs']] = None,
                 operating_status: Optional[pulumi.Input[str]] = None,
                 provisioning_status: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 vip_address: Optional[pulumi.Input[str]] = None,
                 vip_network_id: Optional[pulumi.Input[str]] = None,
                 vip_subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Loadbalancer resources.
        :param pulumi.Input[str] created_at: The UTC date and timestamp when the resource was created
        :param pulumi.Input[str] description: Description of the loadbalancer
        :param pulumi.Input[str] flavor_id: Loadbalancer flavor id
        :param pulumi.Input['LoadbalancerFloatingIpArgs'] floating_ip: Information about floating IP
        :param pulumi.Input[Sequence[pulumi.Input['LoadbalancerListenerArgs']]] listeners: Listeners to create with the loadbalancer
        :param pulumi.Input[str] name: Name of the resource
        :param pulumi.Input['LoadbalancerNetworkArgs'] network: Network information to create the loadbalancer
        :param pulumi.Input[str] operating_status: Operating status of the resource
        :param pulumi.Input[str] provisioning_status: Provisioning status of the resource
        :param pulumi.Input[str] region: Region of the resource
        :param pulumi.Input[str] region_name: Region name
        :param pulumi.Input[str] service_name: Service name
        :param pulumi.Input[str] updated_at: UTC date and timestamp when the resource was created
        :param pulumi.Input[str] vip_address: IP address of the Virtual IP
        :param pulumi.Input[str] vip_network_id: Openstack ID of the network for the Virtual IP
        :param pulumi.Input[str] vip_subnet_id: ID of the subnet for the Virtual IP
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if flavor_id is not None:
            pulumi.set(__self__, "flavor_id", flavor_id)
        if floating_ip is not None:
            pulumi.set(__self__, "floating_ip", floating_ip)
        if listeners is not None:
            pulumi.set(__self__, "listeners", listeners)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if operating_status is not None:
            pulumi.set(__self__, "operating_status", operating_status)
        if provisioning_status is not None:
            pulumi.set(__self__, "provisioning_status", provisioning_status)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if vip_address is not None:
            pulumi.set(__self__, "vip_address", vip_address)
        if vip_network_id is not None:
            pulumi.set(__self__, "vip_network_id", vip_network_id)
        if vip_subnet_id is not None:
            pulumi.set(__self__, "vip_subnet_id", vip_subnet_id)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The UTC date and timestamp when the resource was created
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the loadbalancer
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> Optional[pulumi.Input[str]]:
        """
        Loadbalancer flavor id
        """
        return pulumi.get(self, "flavor_id")

    @flavor_id.setter
    def flavor_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor_id", value)

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> Optional[pulumi.Input['LoadbalancerFloatingIpArgs']]:
        """
        Information about floating IP
        """
        return pulumi.get(self, "floating_ip")

    @floating_ip.setter
    def floating_ip(self, value: Optional[pulumi.Input['LoadbalancerFloatingIpArgs']]):
        pulumi.set(self, "floating_ip", value)

    @property
    @pulumi.getter
    def listeners(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerListenerArgs']]]]:
        """
        Listeners to create with the loadbalancer
        """
        return pulumi.get(self, "listeners")

    @listeners.setter
    def listeners(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerListenerArgs']]]]):
        pulumi.set(self, "listeners", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input['LoadbalancerNetworkArgs']]:
        """
        Network information to create the loadbalancer
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input['LoadbalancerNetworkArgs']]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="operatingStatus")
    def operating_status(self) -> Optional[pulumi.Input[str]]:
        """
        Operating status of the resource
        """
        return pulumi.get(self, "operating_status")

    @operating_status.setter
    def operating_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operating_status", value)

    @property
    @pulumi.getter(name="provisioningStatus")
    def provisioning_status(self) -> Optional[pulumi.Input[str]]:
        """
        Provisioning status of the resource
        """
        return pulumi.get(self, "provisioning_status")

    @provisioning_status.setter
    def provisioning_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provisioning_status", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region of the resource
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[pulumi.Input[str]]:
        """
        Region name
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        UTC date and timestamp when the resource was created
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="vipAddress")
    def vip_address(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the Virtual IP
        """
        return pulumi.get(self, "vip_address")

    @vip_address.setter
    def vip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip_address", value)

    @property
    @pulumi.getter(name="vipNetworkId")
    def vip_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        Openstack ID of the network for the Virtual IP
        """
        return pulumi.get(self, "vip_network_id")

    @vip_network_id.setter
    def vip_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip_network_id", value)

    @property
    @pulumi.getter(name="vipSubnetId")
    def vip_subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the subnet for the Virtual IP
        """
        return pulumi.get(self, "vip_subnet_id")

    @vip_subnet_id.setter
    def vip_subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip_subnet_id", value)


class Loadbalancer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 listeners: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LoadbalancerListenerArgs', 'LoadbalancerListenerArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[Union['LoadbalancerNetworkArgs', 'LoadbalancerNetworkArgsDict']]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Loadbalancer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the loadbalancer
        :param pulumi.Input[str] flavor_id: Loadbalancer flavor id
        :param pulumi.Input[Sequence[pulumi.Input[Union['LoadbalancerListenerArgs', 'LoadbalancerListenerArgsDict']]]] listeners: Listeners to create with the loadbalancer
        :param pulumi.Input[str] name: Name of the resource
        :param pulumi.Input[Union['LoadbalancerNetworkArgs', 'LoadbalancerNetworkArgsDict']] network: Network information to create the loadbalancer
        :param pulumi.Input[str] region_name: Region name
        :param pulumi.Input[str] service_name: Service name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoadbalancerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Loadbalancer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LoadbalancerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoadbalancerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 listeners: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LoadbalancerListenerArgs', 'LoadbalancerListenerArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[Union['LoadbalancerNetworkArgs', 'LoadbalancerNetworkArgsDict']]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoadbalancerArgs.__new__(LoadbalancerArgs)

            __props__.__dict__["description"] = description
            if flavor_id is None and not opts.urn:
                raise TypeError("Missing required property 'flavor_id'")
            __props__.__dict__["flavor_id"] = flavor_id
            __props__.__dict__["listeners"] = listeners
            __props__.__dict__["name"] = name
            if network is None and not opts.urn:
                raise TypeError("Missing required property 'network'")
            __props__.__dict__["network"] = network
            if region_name is None and not opts.urn:
                raise TypeError("Missing required property 'region_name'")
            __props__.__dict__["region_name"] = region_name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["created_at"] = None
            __props__.__dict__["floating_ip"] = None
            __props__.__dict__["operating_status"] = None
            __props__.__dict__["provisioning_status"] = None
            __props__.__dict__["region"] = None
            __props__.__dict__["updated_at"] = None
            __props__.__dict__["vip_address"] = None
            __props__.__dict__["vip_network_id"] = None
            __props__.__dict__["vip_subnet_id"] = None
        super(Loadbalancer, __self__).__init__(
            'ovh:CloudProject/loadbalancer:Loadbalancer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            flavor_id: Optional[pulumi.Input[str]] = None,
            floating_ip: Optional[pulumi.Input[Union['LoadbalancerFloatingIpArgs', 'LoadbalancerFloatingIpArgsDict']]] = None,
            listeners: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LoadbalancerListenerArgs', 'LoadbalancerListenerArgsDict']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[Union['LoadbalancerNetworkArgs', 'LoadbalancerNetworkArgsDict']]] = None,
            operating_status: Optional[pulumi.Input[str]] = None,
            provisioning_status: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            region_name: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            vip_address: Optional[pulumi.Input[str]] = None,
            vip_network_id: Optional[pulumi.Input[str]] = None,
            vip_subnet_id: Optional[pulumi.Input[str]] = None) -> 'Loadbalancer':
        """
        Get an existing Loadbalancer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The UTC date and timestamp when the resource was created
        :param pulumi.Input[str] description: Description of the loadbalancer
        :param pulumi.Input[str] flavor_id: Loadbalancer flavor id
        :param pulumi.Input[Union['LoadbalancerFloatingIpArgs', 'LoadbalancerFloatingIpArgsDict']] floating_ip: Information about floating IP
        :param pulumi.Input[Sequence[pulumi.Input[Union['LoadbalancerListenerArgs', 'LoadbalancerListenerArgsDict']]]] listeners: Listeners to create with the loadbalancer
        :param pulumi.Input[str] name: Name of the resource
        :param pulumi.Input[Union['LoadbalancerNetworkArgs', 'LoadbalancerNetworkArgsDict']] network: Network information to create the loadbalancer
        :param pulumi.Input[str] operating_status: Operating status of the resource
        :param pulumi.Input[str] provisioning_status: Provisioning status of the resource
        :param pulumi.Input[str] region: Region of the resource
        :param pulumi.Input[str] region_name: Region name
        :param pulumi.Input[str] service_name: Service name
        :param pulumi.Input[str] updated_at: UTC date and timestamp when the resource was created
        :param pulumi.Input[str] vip_address: IP address of the Virtual IP
        :param pulumi.Input[str] vip_network_id: Openstack ID of the network for the Virtual IP
        :param pulumi.Input[str] vip_subnet_id: ID of the subnet for the Virtual IP
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LoadbalancerState.__new__(_LoadbalancerState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["flavor_id"] = flavor_id
        __props__.__dict__["floating_ip"] = floating_ip
        __props__.__dict__["listeners"] = listeners
        __props__.__dict__["name"] = name
        __props__.__dict__["network"] = network
        __props__.__dict__["operating_status"] = operating_status
        __props__.__dict__["provisioning_status"] = provisioning_status
        __props__.__dict__["region"] = region
        __props__.__dict__["region_name"] = region_name
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["vip_address"] = vip_address
        __props__.__dict__["vip_network_id"] = vip_network_id
        __props__.__dict__["vip_subnet_id"] = vip_subnet_id
        return Loadbalancer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The UTC date and timestamp when the resource was created
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the loadbalancer
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> pulumi.Output[str]:
        """
        Loadbalancer flavor id
        """
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> pulumi.Output['outputs.LoadbalancerFloatingIp']:
        """
        Information about floating IP
        """
        return pulumi.get(self, "floating_ip")

    @property
    @pulumi.getter
    def listeners(self) -> pulumi.Output[Optional[Sequence['outputs.LoadbalancerListener']]]:
        """
        Listeners to create with the loadbalancer
        """
        return pulumi.get(self, "listeners")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output['outputs.LoadbalancerNetwork']:
        """
        Network information to create the loadbalancer
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="operatingStatus")
    def operating_status(self) -> pulumi.Output[str]:
        """
        Operating status of the resource
        """
        return pulumi.get(self, "operating_status")

    @property
    @pulumi.getter(name="provisioningStatus")
    def provisioning_status(self) -> pulumi.Output[str]:
        """
        Provisioning status of the resource
        """
        return pulumi.get(self, "provisioning_status")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Region of the resource
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Output[str]:
        """
        Region name
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        UTC date and timestamp when the resource was created
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="vipAddress")
    def vip_address(self) -> pulumi.Output[str]:
        """
        IP address of the Virtual IP
        """
        return pulumi.get(self, "vip_address")

    @property
    @pulumi.getter(name="vipNetworkId")
    def vip_network_id(self) -> pulumi.Output[str]:
        """
        Openstack ID of the network for the Virtual IP
        """
        return pulumi.get(self, "vip_network_id")

    @property
    @pulumi.getter(name="vipSubnetId")
    def vip_subnet_id(self) -> pulumi.Output[str]:
        """
        ID of the subnet for the Virtual IP
        """
        return pulumi.get(self, "vip_subnet_id")

