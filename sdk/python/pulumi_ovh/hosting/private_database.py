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

__all__ = ['PrivateDatabaseArgs', 'PrivateDatabase']

@pulumi.input_type
class PrivateDatabaseArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabaseOrderArgs']]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['PrivateDatabasePlanArgs']] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabasePlanOptionArgs']]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivateDatabase resource.
        :param pulumi.Input[str] display_name: Name displayed in customer panel for your private database
        :param pulumi.Input[Sequence[pulumi.Input['PrivateDatabaseOrderArgs']]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: Ovh Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input['PrivateDatabasePlanArgs'] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input['PrivateDatabasePlanOptionArgs']]] plan_options: Product Plan to order
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if orders is not None:
            pulumi.set(__self__, "orders", orders)
        if ovh_subsidiary is not None:
            pulumi.set(__self__, "ovh_subsidiary", ovh_subsidiary)
        if payment_mean is not None:
            warnings.warn("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""", DeprecationWarning)
            pulumi.log.warn("""payment_mean is deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
        if payment_mean is not None:
            pulumi.set(__self__, "payment_mean", payment_mean)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if plan_options is not None:
            pulumi.set(__self__, "plan_options", plan_options)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name displayed in customer panel for your private database
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def orders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabaseOrderArgs']]]]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @orders.setter
    def orders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabaseOrderArgs']]]]):
        pulumi.set(self, "orders", value)

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh Subsidiary
        """
        return pulumi.get(self, "ovh_subsidiary")

    @ovh_subsidiary.setter
    def ovh_subsidiary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ovh_subsidiary", value)

    @property
    @pulumi.getter(name="paymentMean")
    @_utilities.deprecated("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
    def payment_mean(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @payment_mean.setter
    def payment_mean(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_mean", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['PrivateDatabasePlanArgs']]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['PrivateDatabasePlanArgs']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabasePlanOptionArgs']]]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @plan_options.setter
    def plan_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabasePlanOptionArgs']]]]):
        pulumi.set(self, "plan_options", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _PrivateDatabaseState:
    def __init__(__self__, *,
                 database_urn: Optional[pulumi.Input[str]] = None,
                 cpu: Optional[pulumi.Input[int]] = None,
                 datacenter: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 hostname_ftp: Optional[pulumi.Input[str]] = None,
                 infrastructure: Optional[pulumi.Input[str]] = None,
                 offer: Optional[pulumi.Input[str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabaseOrderArgs']]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['PrivateDatabasePlanArgs']] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabasePlanOptionArgs']]]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 port_ftp: Optional[pulumi.Input[int]] = None,
                 quota_size: Optional[pulumi.Input[int]] = None,
                 quota_used: Optional[pulumi.Input[int]] = None,
                 ram: Optional[pulumi.Input[int]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 version_label: Optional[pulumi.Input[str]] = None,
                 version_number: Optional[pulumi.Input[float]] = None):
        """
        Input properties used for looking up and filtering PrivateDatabase resources.
        :param pulumi.Input[int] cpu: Number of CPU on your private database
        :param pulumi.Input[str] datacenter: Datacenter where this private database is located
        :param pulumi.Input[str] display_name: Name displayed in customer panel for your private database
        :param pulumi.Input[str] hostname: Private database hostname
        :param pulumi.Input[str] hostname_ftp: Private database FTP hostname
        :param pulumi.Input[str] infrastructure: Infrastructure where service was stored
        :param pulumi.Input[str] offer: Type of the private database offer
        :param pulumi.Input[Sequence[pulumi.Input['PrivateDatabaseOrderArgs']]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: Ovh Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input['PrivateDatabasePlanArgs'] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input['PrivateDatabasePlanOptionArgs']]] plan_options: Product Plan to order
        :param pulumi.Input[int] port: Private database service port
        :param pulumi.Input[int] port_ftp: Private database FTP port
        :param pulumi.Input[int] quota_size: Space allowed (in MB) on your private database
        :param pulumi.Input[int] quota_used: Sapce used (in MB) on your private database
        :param pulumi.Input[int] ram: Amount of ram (in MB) on your private database
        :param pulumi.Input[str] server: Private database server name
        :param pulumi.Input[str] state: Private database state
        :param pulumi.Input[str] type: Private database type
        :param pulumi.Input[str] version: Private database available versions
        :param pulumi.Input[str] version_label: Private database version label
        :param pulumi.Input[float] version_number: Private database version number
        """
        if database_urn is not None:
            pulumi.set(__self__, "database_urn", database_urn)
        if cpu is not None:
            pulumi.set(__self__, "cpu", cpu)
        if datacenter is not None:
            pulumi.set(__self__, "datacenter", datacenter)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if hostname_ftp is not None:
            pulumi.set(__self__, "hostname_ftp", hostname_ftp)
        if infrastructure is not None:
            pulumi.set(__self__, "infrastructure", infrastructure)
        if offer is not None:
            pulumi.set(__self__, "offer", offer)
        if orders is not None:
            pulumi.set(__self__, "orders", orders)
        if ovh_subsidiary is not None:
            pulumi.set(__self__, "ovh_subsidiary", ovh_subsidiary)
        if payment_mean is not None:
            warnings.warn("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""", DeprecationWarning)
            pulumi.log.warn("""payment_mean is deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
        if payment_mean is not None:
            pulumi.set(__self__, "payment_mean", payment_mean)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if plan_options is not None:
            pulumi.set(__self__, "plan_options", plan_options)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if port_ftp is not None:
            pulumi.set(__self__, "port_ftp", port_ftp)
        if quota_size is not None:
            pulumi.set(__self__, "quota_size", quota_size)
        if quota_used is not None:
            pulumi.set(__self__, "quota_used", quota_used)
        if ram is not None:
            pulumi.set(__self__, "ram", ram)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if version_label is not None:
            pulumi.set(__self__, "version_label", version_label)
        if version_number is not None:
            pulumi.set(__self__, "version_number", version_number)

    @property
    @pulumi.getter(name="DatabaseURN")
    def database_urn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_urn")

    @database_urn.setter
    def database_urn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_urn", value)

    @property
    @pulumi.getter
    def cpu(self) -> Optional[pulumi.Input[int]]:
        """
        Number of CPU on your private database
        """
        return pulumi.get(self, "cpu")

    @cpu.setter
    def cpu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cpu", value)

    @property
    @pulumi.getter
    def datacenter(self) -> Optional[pulumi.Input[str]]:
        """
        Datacenter where this private database is located
        """
        return pulumi.get(self, "datacenter")

    @datacenter.setter
    def datacenter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datacenter", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name displayed in customer panel for your private database
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Private database hostname
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="hostnameFtp")
    def hostname_ftp(self) -> Optional[pulumi.Input[str]]:
        """
        Private database FTP hostname
        """
        return pulumi.get(self, "hostname_ftp")

    @hostname_ftp.setter
    def hostname_ftp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname_ftp", value)

    @property
    @pulumi.getter
    def infrastructure(self) -> Optional[pulumi.Input[str]]:
        """
        Infrastructure where service was stored
        """
        return pulumi.get(self, "infrastructure")

    @infrastructure.setter
    def infrastructure(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "infrastructure", value)

    @property
    @pulumi.getter
    def offer(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the private database offer
        """
        return pulumi.get(self, "offer")

    @offer.setter
    def offer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "offer", value)

    @property
    @pulumi.getter
    def orders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabaseOrderArgs']]]]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @orders.setter
    def orders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabaseOrderArgs']]]]):
        pulumi.set(self, "orders", value)

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh Subsidiary
        """
        return pulumi.get(self, "ovh_subsidiary")

    @ovh_subsidiary.setter
    def ovh_subsidiary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ovh_subsidiary", value)

    @property
    @pulumi.getter(name="paymentMean")
    @_utilities.deprecated("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
    def payment_mean(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @payment_mean.setter
    def payment_mean(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_mean", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['PrivateDatabasePlanArgs']]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['PrivateDatabasePlanArgs']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabasePlanOptionArgs']]]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @plan_options.setter
    def plan_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateDatabasePlanOptionArgs']]]]):
        pulumi.set(self, "plan_options", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Private database service port
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="portFtp")
    def port_ftp(self) -> Optional[pulumi.Input[int]]:
        """
        Private database FTP port
        """
        return pulumi.get(self, "port_ftp")

    @port_ftp.setter
    def port_ftp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port_ftp", value)

    @property
    @pulumi.getter(name="quotaSize")
    def quota_size(self) -> Optional[pulumi.Input[int]]:
        """
        Space allowed (in MB) on your private database
        """
        return pulumi.get(self, "quota_size")

    @quota_size.setter
    def quota_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "quota_size", value)

    @property
    @pulumi.getter(name="quotaUsed")
    def quota_used(self) -> Optional[pulumi.Input[int]]:
        """
        Sapce used (in MB) on your private database
        """
        return pulumi.get(self, "quota_used")

    @quota_used.setter
    def quota_used(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "quota_used", value)

    @property
    @pulumi.getter
    def ram(self) -> Optional[pulumi.Input[int]]:
        """
        Amount of ram (in MB) on your private database
        """
        return pulumi.get(self, "ram")

    @ram.setter
    def ram(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ram", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        Private database server name
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Private database state
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Private database type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Private database available versions
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="versionLabel")
    def version_label(self) -> Optional[pulumi.Input[str]]:
        """
        Private database version label
        """
        return pulumi.get(self, "version_label")

    @version_label.setter
    def version_label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_label", value)

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> Optional[pulumi.Input[float]]:
        """
        Private database version number
        """
        return pulumi.get(self, "version_number")

    @version_number.setter
    def version_number(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "version_number", value)


class PrivateDatabase(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabaseOrderArgs', 'PrivateDatabaseOrderArgsDict']]]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[Union['PrivateDatabasePlanArgs', 'PrivateDatabasePlanArgsDict']]] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabasePlanOptionArgs', 'PrivateDatabasePlanOptionArgsDict']]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PrivateDatabase resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Name displayed in customer panel for your private database
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabaseOrderArgs', 'PrivateDatabaseOrderArgsDict']]]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: Ovh Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input[Union['PrivateDatabasePlanArgs', 'PrivateDatabasePlanArgsDict']] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabasePlanOptionArgs', 'PrivateDatabasePlanOptionArgsDict']]]] plan_options: Product Plan to order
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PrivateDatabaseArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PrivateDatabase resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PrivateDatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateDatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabaseOrderArgs', 'PrivateDatabaseOrderArgsDict']]]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[Union['PrivateDatabasePlanArgs', 'PrivateDatabasePlanArgsDict']]] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabasePlanOptionArgs', 'PrivateDatabasePlanOptionArgsDict']]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateDatabaseArgs.__new__(PrivateDatabaseArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["orders"] = orders
            __props__.__dict__["ovh_subsidiary"] = ovh_subsidiary
            __props__.__dict__["payment_mean"] = payment_mean
            __props__.__dict__["plan"] = plan
            __props__.__dict__["plan_options"] = plan_options
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["database_urn"] = None
            __props__.__dict__["cpu"] = None
            __props__.__dict__["datacenter"] = None
            __props__.__dict__["hostname"] = None
            __props__.__dict__["hostname_ftp"] = None
            __props__.__dict__["infrastructure"] = None
            __props__.__dict__["offer"] = None
            __props__.__dict__["port"] = None
            __props__.__dict__["port_ftp"] = None
            __props__.__dict__["quota_size"] = None
            __props__.__dict__["quota_used"] = None
            __props__.__dict__["ram"] = None
            __props__.__dict__["server"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["version"] = None
            __props__.__dict__["version_label"] = None
            __props__.__dict__["version_number"] = None
        super(PrivateDatabase, __self__).__init__(
            'ovh:Hosting/privateDatabase:PrivateDatabase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database_urn: Optional[pulumi.Input[str]] = None,
            cpu: Optional[pulumi.Input[int]] = None,
            datacenter: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            hostname_ftp: Optional[pulumi.Input[str]] = None,
            infrastructure: Optional[pulumi.Input[str]] = None,
            offer: Optional[pulumi.Input[str]] = None,
            orders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabaseOrderArgs', 'PrivateDatabaseOrderArgsDict']]]]] = None,
            ovh_subsidiary: Optional[pulumi.Input[str]] = None,
            payment_mean: Optional[pulumi.Input[str]] = None,
            plan: Optional[pulumi.Input[Union['PrivateDatabasePlanArgs', 'PrivateDatabasePlanArgsDict']]] = None,
            plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabasePlanOptionArgs', 'PrivateDatabasePlanOptionArgsDict']]]]] = None,
            port: Optional[pulumi.Input[int]] = None,
            port_ftp: Optional[pulumi.Input[int]] = None,
            quota_size: Optional[pulumi.Input[int]] = None,
            quota_used: Optional[pulumi.Input[int]] = None,
            ram: Optional[pulumi.Input[int]] = None,
            server: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None,
            version_label: Optional[pulumi.Input[str]] = None,
            version_number: Optional[pulumi.Input[float]] = None) -> 'PrivateDatabase':
        """
        Get an existing PrivateDatabase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cpu: Number of CPU on your private database
        :param pulumi.Input[str] datacenter: Datacenter where this private database is located
        :param pulumi.Input[str] display_name: Name displayed in customer panel for your private database
        :param pulumi.Input[str] hostname: Private database hostname
        :param pulumi.Input[str] hostname_ftp: Private database FTP hostname
        :param pulumi.Input[str] infrastructure: Infrastructure where service was stored
        :param pulumi.Input[str] offer: Type of the private database offer
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabaseOrderArgs', 'PrivateDatabaseOrderArgsDict']]]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: Ovh Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input[Union['PrivateDatabasePlanArgs', 'PrivateDatabasePlanArgsDict']] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrivateDatabasePlanOptionArgs', 'PrivateDatabasePlanOptionArgsDict']]]] plan_options: Product Plan to order
        :param pulumi.Input[int] port: Private database service port
        :param pulumi.Input[int] port_ftp: Private database FTP port
        :param pulumi.Input[int] quota_size: Space allowed (in MB) on your private database
        :param pulumi.Input[int] quota_used: Sapce used (in MB) on your private database
        :param pulumi.Input[int] ram: Amount of ram (in MB) on your private database
        :param pulumi.Input[str] server: Private database server name
        :param pulumi.Input[str] state: Private database state
        :param pulumi.Input[str] type: Private database type
        :param pulumi.Input[str] version: Private database available versions
        :param pulumi.Input[str] version_label: Private database version label
        :param pulumi.Input[float] version_number: Private database version number
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateDatabaseState.__new__(_PrivateDatabaseState)

        __props__.__dict__["database_urn"] = database_urn
        __props__.__dict__["cpu"] = cpu
        __props__.__dict__["datacenter"] = datacenter
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["hostname_ftp"] = hostname_ftp
        __props__.__dict__["infrastructure"] = infrastructure
        __props__.__dict__["offer"] = offer
        __props__.__dict__["orders"] = orders
        __props__.__dict__["ovh_subsidiary"] = ovh_subsidiary
        __props__.__dict__["payment_mean"] = payment_mean
        __props__.__dict__["plan"] = plan
        __props__.__dict__["plan_options"] = plan_options
        __props__.__dict__["port"] = port
        __props__.__dict__["port_ftp"] = port_ftp
        __props__.__dict__["quota_size"] = quota_size
        __props__.__dict__["quota_used"] = quota_used
        __props__.__dict__["ram"] = ram
        __props__.__dict__["server"] = server
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["state"] = state
        __props__.__dict__["type"] = type
        __props__.__dict__["version"] = version
        __props__.__dict__["version_label"] = version_label
        __props__.__dict__["version_number"] = version_number
        return PrivateDatabase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="DatabaseURN")
    def database_urn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "database_urn")

    @property
    @pulumi.getter
    def cpu(self) -> pulumi.Output[int]:
        """
        Number of CPU on your private database
        """
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter
    def datacenter(self) -> pulumi.Output[str]:
        """
        Datacenter where this private database is located
        """
        return pulumi.get(self, "datacenter")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Name displayed in customer panel for your private database
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        Private database hostname
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="hostnameFtp")
    def hostname_ftp(self) -> pulumi.Output[str]:
        """
        Private database FTP hostname
        """
        return pulumi.get(self, "hostname_ftp")

    @property
    @pulumi.getter
    def infrastructure(self) -> pulumi.Output[str]:
        """
        Infrastructure where service was stored
        """
        return pulumi.get(self, "infrastructure")

    @property
    @pulumi.getter
    def offer(self) -> pulumi.Output[str]:
        """
        Type of the private database offer
        """
        return pulumi.get(self, "offer")

    @property
    @pulumi.getter
    def orders(self) -> pulumi.Output[Sequence['outputs.PrivateDatabaseOrder']]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> pulumi.Output[str]:
        """
        Ovh Subsidiary
        """
        return pulumi.get(self, "ovh_subsidiary")

    @property
    @pulumi.getter(name="paymentMean")
    @_utilities.deprecated("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
    def payment_mean(self) -> pulumi.Output[Optional[str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output['outputs.PrivateDatabasePlan']:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> pulumi.Output[Optional[Sequence['outputs.PrivateDatabasePlanOption']]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Private database service port
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="portFtp")
    def port_ftp(self) -> pulumi.Output[int]:
        """
        Private database FTP port
        """
        return pulumi.get(self, "port_ftp")

    @property
    @pulumi.getter(name="quotaSize")
    def quota_size(self) -> pulumi.Output[int]:
        """
        Space allowed (in MB) on your private database
        """
        return pulumi.get(self, "quota_size")

    @property
    @pulumi.getter(name="quotaUsed")
    def quota_used(self) -> pulumi.Output[int]:
        """
        Sapce used (in MB) on your private database
        """
        return pulumi.get(self, "quota_used")

    @property
    @pulumi.getter
    def ram(self) -> pulumi.Output[int]:
        """
        Amount of ram (in MB) on your private database
        """
        return pulumi.get(self, "ram")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        Private database server name
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Private database state
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Private database type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Private database available versions
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="versionLabel")
    def version_label(self) -> pulumi.Output[str]:
        """
        Private database version label
        """
        return pulumi.get(self, "version_label")

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> pulumi.Output[float]:
        """
        Private database version number
        """
        return pulumi.get(self, "version_number")

