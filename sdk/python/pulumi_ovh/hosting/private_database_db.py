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

__all__ = ['PrivateDatabaseDbArgs', 'PrivateDatabaseDb']

@pulumi.input_type
class PrivateDatabaseDbArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[builtins.str],
                 service_name: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a PrivateDatabaseDb resource.
        :param pulumi.Input[builtins.str] database_name: Name of your new database
        :param pulumi.Input[builtins.str] service_name: The internal name of your private database.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of your new database
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The internal name of your private database.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _PrivateDatabaseDbState:
    def __init__(__self__, *,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering PrivateDatabaseDb resources.
        :param pulumi.Input[builtins.str] database_name: Name of your new database
        :param pulumi.Input[builtins.str] service_name: The internal name of your private database.
        """
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of your new database
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The internal name of your private database.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)


class PrivateDatabaseDb(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a new database on your private cloud database service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        database = ovh.hosting.PrivateDatabaseDb("database",
            database_name="XXXXXX",
            service_name="XXXXXX")
        ```

        ## Import

        OVHcloud Webhosting database can be imported using the `service_name` and the `database_name`, separated by "/" E.g.,

        ```sh
        $ pulumi import ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb database service_name/database_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] database_name: Name of your new database
        :param pulumi.Input[builtins.str] service_name: The internal name of your private database.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateDatabaseDbArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a new database on your private cloud database service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        database = ovh.hosting.PrivateDatabaseDb("database",
            database_name="XXXXXX",
            service_name="XXXXXX")
        ```

        ## Import

        OVHcloud Webhosting database can be imported using the `service_name` and the `database_name`, separated by "/" E.g.,

        ```sh
        $ pulumi import ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb database service_name/database_name
        ```

        :param str resource_name: The name of the resource.
        :param PrivateDatabaseDbArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateDatabaseDbArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateDatabaseDbArgs.__new__(PrivateDatabaseDbArgs)

            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(PrivateDatabaseDb, __self__).__init__(
            'ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database_name: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None) -> 'PrivateDatabaseDb':
        """
        Get an existing PrivateDatabaseDb resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] database_name: Name of your new database
        :param pulumi.Input[builtins.str] service_name: The internal name of your private database.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateDatabaseDbState.__new__(_PrivateDatabaseDbState)

        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["service_name"] = service_name
        return PrivateDatabaseDb(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of your new database
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The internal name of your private database.
        """
        return pulumi.get(self, "service_name")

