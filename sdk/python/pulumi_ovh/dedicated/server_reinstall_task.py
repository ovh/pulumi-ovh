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

__all__ = ['ServerReinstallTaskArgs', 'ServerReinstallTask']

@pulumi.input_type
class ServerReinstallTaskArgs:
    def __init__(__self__, *,
                 os: pulumi.Input[builtins.str],
                 service_name: pulumi.Input[builtins.str],
                 bootid_on_destroy: Optional[pulumi.Input[builtins.int]] = None,
                 customizations: Optional[pulumi.Input['ServerReinstallTaskCustomizationsArgs']] = None,
                 properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 storages: Optional[pulumi.Input[Sequence[pulumi.Input['ServerReinstallTaskStorageArgs']]]] = None):
        """
        The set of arguments for constructing a ServerReinstallTask resource.
        :param pulumi.Input[builtins.str] os: Operating system to install.
        :param pulumi.Input[builtins.str] service_name: The service_name of your dedicated server.
        :param pulumi.Input[builtins.int] bootid_on_destroy: If set, reboot the server on the specified boot id during destroy phase.
        :param pulumi.Input['ServerReinstallTaskCustomizationsArgs'] customizations: Available attributes and their types are OS-dependant. Example: `hostname`.
               
               > **WARNING** Some customizations may be required on some Operating Systems. [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] properties: Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
        :param pulumi.Input[Sequence[pulumi.Input['ServerReinstallTaskStorageArgs']]] storages: OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
        """
        pulumi.set(__self__, "os", os)
        pulumi.set(__self__, "service_name", service_name)
        if bootid_on_destroy is not None:
            pulumi.set(__self__, "bootid_on_destroy", bootid_on_destroy)
        if customizations is not None:
            pulumi.set(__self__, "customizations", customizations)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if storages is not None:
            pulumi.set(__self__, "storages", storages)

    @property
    @pulumi.getter
    def os(self) -> pulumi.Input[builtins.str]:
        """
        Operating system to install.
        """
        return pulumi.get(self, "os")

    @os.setter
    def os(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "os", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The service_name of your dedicated server.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="bootidOnDestroy")
    def bootid_on_destroy(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        If set, reboot the server on the specified boot id during destroy phase.
        """
        return pulumi.get(self, "bootid_on_destroy")

    @bootid_on_destroy.setter
    def bootid_on_destroy(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "bootid_on_destroy", value)

    @property
    @pulumi.getter
    def customizations(self) -> Optional[pulumi.Input['ServerReinstallTaskCustomizationsArgs']]:
        """
        Available attributes and their types are OS-dependant. Example: `hostname`.

        > **WARNING** Some customizations may be required on some Operating Systems. [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
        """
        return pulumi.get(self, "customizations")

    @customizations.setter
    def customizations(self, value: Optional[pulumi.Input['ServerReinstallTaskCustomizationsArgs']]):
        pulumi.set(self, "customizations", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def storages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerReinstallTaskStorageArgs']]]]:
        """
        OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
        """
        return pulumi.get(self, "storages")

    @storages.setter
    def storages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerReinstallTaskStorageArgs']]]]):
        pulumi.set(self, "storages", value)


@pulumi.input_type
class _ServerReinstallTaskState:
    def __init__(__self__, *,
                 bootid_on_destroy: Optional[pulumi.Input[builtins.int]] = None,
                 comment: Optional[pulumi.Input[builtins.str]] = None,
                 customizations: Optional[pulumi.Input['ServerReinstallTaskCustomizationsArgs']] = None,
                 done_date: Optional[pulumi.Input[builtins.str]] = None,
                 function: Optional[pulumi.Input[builtins.str]] = None,
                 last_update: Optional[pulumi.Input[builtins.str]] = None,
                 os: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 start_date: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 storages: Optional[pulumi.Input[Sequence[pulumi.Input['ServerReinstallTaskStorageArgs']]]] = None):
        """
        Input properties used for looking up and filtering ServerReinstallTask resources.
        :param pulumi.Input[builtins.int] bootid_on_destroy: If set, reboot the server on the specified boot id during destroy phase.
        :param pulumi.Input[builtins.str] comment: Details of this task. (should be `Install asked`)
        :param pulumi.Input['ServerReinstallTaskCustomizationsArgs'] customizations: Available attributes and their types are OS-dependant. Example: `hostname`.
               
               > **WARNING** Some customizations may be required on some Operating Systems. [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
        :param pulumi.Input[builtins.str] done_date: Completion date in RFC3339 format.
        :param pulumi.Input[builtins.str] function: Function name (should be `hardInstall`).
        :param pulumi.Input[builtins.str] last_update: Last update
        :param pulumi.Input[builtins.str] os: Operating system to install.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] properties: Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
        :param pulumi.Input[builtins.str] service_name: The service_name of your dedicated server.
        :param pulumi.Input[builtins.str] start_date: Task creation date in RFC3339 format.
        :param pulumi.Input[builtins.str] status: Task status (should be `done`)
        :param pulumi.Input[Sequence[pulumi.Input['ServerReinstallTaskStorageArgs']]] storages: OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
        """
        if bootid_on_destroy is not None:
            pulumi.set(__self__, "bootid_on_destroy", bootid_on_destroy)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if customizations is not None:
            pulumi.set(__self__, "customizations", customizations)
        if done_date is not None:
            pulumi.set(__self__, "done_date", done_date)
        if function is not None:
            pulumi.set(__self__, "function", function)
        if last_update is not None:
            pulumi.set(__self__, "last_update", last_update)
        if os is not None:
            pulumi.set(__self__, "os", os)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if storages is not None:
            pulumi.set(__self__, "storages", storages)

    @property
    @pulumi.getter(name="bootidOnDestroy")
    def bootid_on_destroy(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        If set, reboot the server on the specified boot id during destroy phase.
        """
        return pulumi.get(self, "bootid_on_destroy")

    @bootid_on_destroy.setter
    def bootid_on_destroy(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "bootid_on_destroy", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Details of this task. (should be `Install asked`)
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def customizations(self) -> Optional[pulumi.Input['ServerReinstallTaskCustomizationsArgs']]:
        """
        Available attributes and their types are OS-dependant. Example: `hostname`.

        > **WARNING** Some customizations may be required on some Operating Systems. [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
        """
        return pulumi.get(self, "customizations")

    @customizations.setter
    def customizations(self, value: Optional[pulumi.Input['ServerReinstallTaskCustomizationsArgs']]):
        pulumi.set(self, "customizations", value)

    @property
    @pulumi.getter(name="doneDate")
    def done_date(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Completion date in RFC3339 format.
        """
        return pulumi.get(self, "done_date")

    @done_date.setter
    def done_date(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "done_date", value)

    @property
    @pulumi.getter
    def function(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Function name (should be `hardInstall`).
        """
        return pulumi.get(self, "function")

    @function.setter
    def function(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "function", value)

    @property
    @pulumi.getter(name="lastUpdate")
    def last_update(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Last update
        """
        return pulumi.get(self, "last_update")

    @last_update.setter
    def last_update(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "last_update", value)

    @property
    @pulumi.getter
    def os(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Operating system to install.
        """
        return pulumi.get(self, "os")

    @os.setter
    def os(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "os", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The service_name of your dedicated server.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Task creation date in RFC3339 format.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Task status (should be `done`)
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def storages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServerReinstallTaskStorageArgs']]]]:
        """
        OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
        """
        return pulumi.get(self, "storages")

    @storages.setter
    def storages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServerReinstallTaskStorageArgs']]]]):
        pulumi.set(self, "storages", value)


@pulumi.type_token("ovh:Dedicated/serverReinstallTask:ServerReinstallTask")
class ServerReinstallTask(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bootid_on_destroy: Optional[pulumi.Input[builtins.int]] = None,
                 customizations: Optional[pulumi.Input[Union['ServerReinstallTaskCustomizationsArgs', 'ServerReinstallTaskCustomizationsArgsDict']]] = None,
                 os: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 storages: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServerReinstallTaskStorageArgs', 'ServerReinstallTaskStorageArgsDict']]]]] = None,
                 __props__=None):
        """
        Create a ServerReinstallTask resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] bootid_on_destroy: If set, reboot the server on the specified boot id during destroy phase.
        :param pulumi.Input[Union['ServerReinstallTaskCustomizationsArgs', 'ServerReinstallTaskCustomizationsArgsDict']] customizations: Available attributes and their types are OS-dependant. Example: `hostname`.
               
               > **WARNING** Some customizations may be required on some Operating Systems. [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
        :param pulumi.Input[builtins.str] os: Operating system to install.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] properties: Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
        :param pulumi.Input[builtins.str] service_name: The service_name of your dedicated server.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ServerReinstallTaskStorageArgs', 'ServerReinstallTaskStorageArgsDict']]]] storages: OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerReinstallTaskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ServerReinstallTask resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ServerReinstallTaskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerReinstallTaskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bootid_on_destroy: Optional[pulumi.Input[builtins.int]] = None,
                 customizations: Optional[pulumi.Input[Union['ServerReinstallTaskCustomizationsArgs', 'ServerReinstallTaskCustomizationsArgsDict']]] = None,
                 os: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 storages: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServerReinstallTaskStorageArgs', 'ServerReinstallTaskStorageArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerReinstallTaskArgs.__new__(ServerReinstallTaskArgs)

            __props__.__dict__["bootid_on_destroy"] = bootid_on_destroy
            __props__.__dict__["customizations"] = customizations
            if os is None and not opts.urn:
                raise TypeError("Missing required property 'os'")
            __props__.__dict__["os"] = os
            __props__.__dict__["properties"] = properties
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["storages"] = storages
            __props__.__dict__["comment"] = None
            __props__.__dict__["done_date"] = None
            __props__.__dict__["function"] = None
            __props__.__dict__["last_update"] = None
            __props__.__dict__["start_date"] = None
            __props__.__dict__["status"] = None
        super(ServerReinstallTask, __self__).__init__(
            'ovh:Dedicated/serverReinstallTask:ServerReinstallTask',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bootid_on_destroy: Optional[pulumi.Input[builtins.int]] = None,
            comment: Optional[pulumi.Input[builtins.str]] = None,
            customizations: Optional[pulumi.Input[Union['ServerReinstallTaskCustomizationsArgs', 'ServerReinstallTaskCustomizationsArgsDict']]] = None,
            done_date: Optional[pulumi.Input[builtins.str]] = None,
            function: Optional[pulumi.Input[builtins.str]] = None,
            last_update: Optional[pulumi.Input[builtins.str]] = None,
            os: Optional[pulumi.Input[builtins.str]] = None,
            properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            start_date: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            storages: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServerReinstallTaskStorageArgs', 'ServerReinstallTaskStorageArgsDict']]]]] = None) -> 'ServerReinstallTask':
        """
        Get an existing ServerReinstallTask resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] bootid_on_destroy: If set, reboot the server on the specified boot id during destroy phase.
        :param pulumi.Input[builtins.str] comment: Details of this task. (should be `Install asked`)
        :param pulumi.Input[Union['ServerReinstallTaskCustomizationsArgs', 'ServerReinstallTaskCustomizationsArgsDict']] customizations: Available attributes and their types are OS-dependant. Example: `hostname`.
               
               > **WARNING** Some customizations may be required on some Operating Systems. [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
        :param pulumi.Input[builtins.str] done_date: Completion date in RFC3339 format.
        :param pulumi.Input[builtins.str] function: Function name (should be `hardInstall`).
        :param pulumi.Input[builtins.str] last_update: Last update
        :param pulumi.Input[builtins.str] os: Operating system to install.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] properties: Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
        :param pulumi.Input[builtins.str] service_name: The service_name of your dedicated server.
        :param pulumi.Input[builtins.str] start_date: Task creation date in RFC3339 format.
        :param pulumi.Input[builtins.str] status: Task status (should be `done`)
        :param pulumi.Input[Sequence[pulumi.Input[Union['ServerReinstallTaskStorageArgs', 'ServerReinstallTaskStorageArgsDict']]]] storages: OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerReinstallTaskState.__new__(_ServerReinstallTaskState)

        __props__.__dict__["bootid_on_destroy"] = bootid_on_destroy
        __props__.__dict__["comment"] = comment
        __props__.__dict__["customizations"] = customizations
        __props__.__dict__["done_date"] = done_date
        __props__.__dict__["function"] = function
        __props__.__dict__["last_update"] = last_update
        __props__.__dict__["os"] = os
        __props__.__dict__["properties"] = properties
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["start_date"] = start_date
        __props__.__dict__["status"] = status
        __props__.__dict__["storages"] = storages
        return ServerReinstallTask(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bootidOnDestroy")
    def bootid_on_destroy(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        If set, reboot the server on the specified boot id during destroy phase.
        """
        return pulumi.get(self, "bootid_on_destroy")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[builtins.str]:
        """
        Details of this task. (should be `Install asked`)
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def customizations(self) -> pulumi.Output[Optional['outputs.ServerReinstallTaskCustomizations']]:
        """
        Available attributes and their types are OS-dependant. Example: `hostname`.

        > **WARNING** Some customizations may be required on some Operating Systems. [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
        """
        return pulumi.get(self, "customizations")

    @property
    @pulumi.getter(name="doneDate")
    def done_date(self) -> pulumi.Output[builtins.str]:
        """
        Completion date in RFC3339 format.
        """
        return pulumi.get(self, "done_date")

    @property
    @pulumi.getter
    def function(self) -> pulumi.Output[builtins.str]:
        """
        Function name (should be `hardInstall`).
        """
        return pulumi.get(self, "function")

    @property
    @pulumi.getter(name="lastUpdate")
    def last_update(self) -> pulumi.Output[builtins.str]:
        """
        Last update
        """
        return pulumi.get(self, "last_update")

    @property
    @pulumi.getter
    def os(self) -> pulumi.Output[builtins.str]:
        """
        Operating system to install.
        """
        return pulumi.get(self, "os")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The service_name of your dedicated server.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> pulumi.Output[builtins.str]:
        """
        Task creation date in RFC3339 format.
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Task status (should be `done`)
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def storages(self) -> pulumi.Output[Optional[Sequence['outputs.ServerReinstallTaskStorage']]]:
        """
        OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
        """
        return pulumi.get(self, "storages")

