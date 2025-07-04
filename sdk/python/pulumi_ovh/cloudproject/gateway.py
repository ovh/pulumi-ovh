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

__all__ = ['GatewayArgs', 'Gateway']

@pulumi.input_type
class GatewayArgs:
    def __init__(__self__, *,
                 model: pulumi.Input[builtins.str],
                 network_id: pulumi.Input[builtins.str],
                 region: pulumi.Input[builtins.str],
                 subnet_id: pulumi.Input[builtins.str],
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Gateway resource.
        :param pulumi.Input[builtins.str] model: Model of the gateway.
        :param pulumi.Input[builtins.str] network_id: ID of the private network.
        :param pulumi.Input[builtins.str] region: Region of the gateway.
        :param pulumi.Input[builtins.str] subnet_id: ID of the subnet.
        :param pulumi.Input[builtins.str] name: Name of the gateway.
        :param pulumi.Input[builtins.str] service_name: ID of the private network.
        """
        pulumi.set(__self__, "model", model)
        pulumi.set(__self__, "network_id", network_id)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def model(self) -> pulumi.Input[builtins.str]:
        """
        Model of the gateway.
        """
        return pulumi.get(self, "model")

    @model.setter
    def model(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "model", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[builtins.str]:
        """
        ID of the private network.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[builtins.str]:
        """
        Region of the gateway.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[builtins.str]:
        """
        ID of the subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the gateway.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the private network.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _GatewayState:
    def __init__(__self__, *,
                 external_informations: Optional[pulumi.Input[Sequence[pulumi.Input['GatewayExternalInformationArgs']]]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['GatewayInterfaceArgs']]]] = None,
                 model: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Gateway resources.
        :param pulumi.Input[Sequence[pulumi.Input['GatewayExternalInformationArgs']]] external_informations: List of External Information of the gateway.
        :param pulumi.Input[Sequence[pulumi.Input['GatewayInterfaceArgs']]] interfaces: Interfaces list of the gateway.
        :param pulumi.Input[builtins.str] model: Model of the gateway.
        :param pulumi.Input[builtins.str] name: Name of the gateway.
        :param pulumi.Input[builtins.str] network_id: ID of the private network.
        :param pulumi.Input[builtins.str] region: Region of the gateway.
        :param pulumi.Input[builtins.str] service_name: ID of the private network.
        :param pulumi.Input[builtins.str] status: Status of the gateway.
        :param pulumi.Input[builtins.str] subnet_id: ID of the subnet.
        """
        if external_informations is not None:
            pulumi.set(__self__, "external_informations", external_informations)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if model is not None:
            pulumi.set(__self__, "model", model)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="externalInformations")
    def external_informations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GatewayExternalInformationArgs']]]]:
        """
        List of External Information of the gateway.
        """
        return pulumi.get(self, "external_informations")

    @external_informations.setter
    def external_informations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GatewayExternalInformationArgs']]]]):
        pulumi.set(self, "external_informations", value)

    @property
    @pulumi.getter
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GatewayInterfaceArgs']]]]:
        """
        Interfaces list of the gateway.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GatewayInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter
    def model(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Model of the gateway.
        """
        return pulumi.get(self, "model")

    @model.setter
    def model(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "model", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the gateway.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the private network.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Region of the gateway.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the private network.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Status of the gateway.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.type_token("ovh:CloudProject/gateway:Gateway")
class Gateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 model: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a new Gateway for existing subnet in the specified public cloud project.

        ## Import

        A gateway can be imported using the `service_name`, `region` and `id` (identifier of the gateway) properties, separated by a `/`.

        However, please note that in the case of a gateway import, `network_id` and `subnet_id` values used at gateway creation are not injected back in the state.

        If you want to define these properties on your imported resource, you have to add an "ignore_changes" lifecycle argument in order not to trigger a recreation, as suggested in the following example.

        terraform

        resource "ovh_cloud_project_gateway" "imported_gateway" {

          service_name = ovh_cloud_project_network_private.mypriv.service_name

          name         = "<my-imported-gateway>"

          model        = "<my-model>"

          region       = "<my-region>"

          network_id   = "<my-imported-gateway-network-id>"

          subnet_id    = "<my-imported-gateway-subnet-id>"

          lifecycle {

            ignore_changes = [network_id, subnet_id]

          }

        }

        import {

          id = "<service-name>/<region>/<gateway-id>"

          to = ovh_cloud_project_gateway.imported_gateway

        }

        bash

        ```sh
        $ pulumi import ovh:CloudProject/gateway:Gateway gateway service_name/region/id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] model: Model of the gateway.
        :param pulumi.Input[builtins.str] name: Name of the gateway.
        :param pulumi.Input[builtins.str] network_id: ID of the private network.
        :param pulumi.Input[builtins.str] region: Region of the gateway.
        :param pulumi.Input[builtins.str] service_name: ID of the private network.
        :param pulumi.Input[builtins.str] subnet_id: ID of the subnet.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a new Gateway for existing subnet in the specified public cloud project.

        ## Import

        A gateway can be imported using the `service_name`, `region` and `id` (identifier of the gateway) properties, separated by a `/`.

        However, please note that in the case of a gateway import, `network_id` and `subnet_id` values used at gateway creation are not injected back in the state.

        If you want to define these properties on your imported resource, you have to add an "ignore_changes" lifecycle argument in order not to trigger a recreation, as suggested in the following example.

        terraform

        resource "ovh_cloud_project_gateway" "imported_gateway" {

          service_name = ovh_cloud_project_network_private.mypriv.service_name

          name         = "<my-imported-gateway>"

          model        = "<my-model>"

          region       = "<my-region>"

          network_id   = "<my-imported-gateway-network-id>"

          subnet_id    = "<my-imported-gateway-subnet-id>"

          lifecycle {

            ignore_changes = [network_id, subnet_id]

          }

        }

        import {

          id = "<service-name>/<region>/<gateway-id>"

          to = ovh_cloud_project_gateway.imported_gateway

        }

        bash

        ```sh
        $ pulumi import ovh:CloudProject/gateway:Gateway gateway service_name/region/id
        ```

        :param str resource_name: The name of the resource.
        :param GatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 model: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewayArgs.__new__(GatewayArgs)

            if model is None and not opts.urn:
                raise TypeError("Missing required property 'model'")
            __props__.__dict__["model"] = model
            __props__.__dict__["name"] = name
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["service_name"] = service_name
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["external_informations"] = None
            __props__.__dict__["interfaces"] = None
            __props__.__dict__["status"] = None
        super(Gateway, __self__).__init__(
            'ovh:CloudProject/gateway:Gateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            external_informations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GatewayExternalInformationArgs', 'GatewayExternalInformationArgsDict']]]]] = None,
            interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GatewayInterfaceArgs', 'GatewayInterfaceArgsDict']]]]] = None,
            model: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            network_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            subnet_id: Optional[pulumi.Input[builtins.str]] = None) -> 'Gateway':
        """
        Get an existing Gateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['GatewayExternalInformationArgs', 'GatewayExternalInformationArgsDict']]]] external_informations: List of External Information of the gateway.
        :param pulumi.Input[Sequence[pulumi.Input[Union['GatewayInterfaceArgs', 'GatewayInterfaceArgsDict']]]] interfaces: Interfaces list of the gateway.
        :param pulumi.Input[builtins.str] model: Model of the gateway.
        :param pulumi.Input[builtins.str] name: Name of the gateway.
        :param pulumi.Input[builtins.str] network_id: ID of the private network.
        :param pulumi.Input[builtins.str] region: Region of the gateway.
        :param pulumi.Input[builtins.str] service_name: ID of the private network.
        :param pulumi.Input[builtins.str] status: Status of the gateway.
        :param pulumi.Input[builtins.str] subnet_id: ID of the subnet.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayState.__new__(_GatewayState)

        __props__.__dict__["external_informations"] = external_informations
        __props__.__dict__["interfaces"] = interfaces
        __props__.__dict__["model"] = model
        __props__.__dict__["name"] = name
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["region"] = region
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        __props__.__dict__["subnet_id"] = subnet_id
        return Gateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="externalInformations")
    def external_informations(self) -> pulumi.Output[Sequence['outputs.GatewayExternalInformation']]:
        """
        List of External Information of the gateway.
        """
        return pulumi.get(self, "external_informations")

    @property
    @pulumi.getter
    def interfaces(self) -> pulumi.Output[Sequence['outputs.GatewayInterface']]:
        """
        Interfaces list of the gateway.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter
    def model(self) -> pulumi.Output[builtins.str]:
        """
        Model of the gateway.
        """
        return pulumi.get(self, "model")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the gateway.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the private network.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        Region of the gateway.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        ID of the private network.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Status of the gateway.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the subnet.
        """
        return pulumi.get(self, "subnet_id")

