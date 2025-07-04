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

__all__ = ['CloudProjectContainerregistryIamArgs', 'CloudProjectContainerregistryIam']

@pulumi.input_type
class CloudProjectContainerregistryIamArgs:
    def __init__(__self__, *,
                 registry_id: pulumi.Input[builtins.str],
                 delete_users: Optional[pulumi.Input[builtins.bool]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a CloudProjectContainerregistryIam resource.
        :param pulumi.Input[builtins.str] registry_id: The ID of the Managed Private Registry. **Changing this value recreates the resource.**
        :param pulumi.Input[builtins.bool] delete_users: Delete existing users from Harbor. IAM can't be enabled if there is at least one user already created. This parameter is only used at IAM configuration creation. **Changing this value recreates the resource.**
        :param pulumi.Input[builtins.str] service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        pulumi.set(__self__, "registry_id", registry_id)
        if delete_users is not None:
            pulumi.set(__self__, "delete_users", delete_users)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the Managed Private Registry. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter(name="deleteUsers")
    def delete_users(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Delete existing users from Harbor. IAM can't be enabled if there is at least one user already created. This parameter is only used at IAM configuration creation. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "delete_users")

    @delete_users.setter
    def delete_users(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "delete_users", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _CloudProjectContainerregistryIamState:
    def __init__(__self__, *,
                 delete_users: Optional[pulumi.Input[builtins.bool]] = None,
                 iam_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 registry_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering CloudProjectContainerregistryIam resources.
        :param pulumi.Input[builtins.bool] delete_users: Delete existing users from Harbor. IAM can't be enabled if there is at least one user already created. This parameter is only used at IAM configuration creation. **Changing this value recreates the resource.**
        :param pulumi.Input[builtins.str] registry_id: The ID of the Managed Private Registry. **Changing this value recreates the resource.**
        :param pulumi.Input[builtins.str] service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        if delete_users is not None:
            pulumi.set(__self__, "delete_users", delete_users)
        if iam_enabled is not None:
            pulumi.set(__self__, "iam_enabled", iam_enabled)
        if registry_id is not None:
            pulumi.set(__self__, "registry_id", registry_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="deleteUsers")
    def delete_users(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Delete existing users from Harbor. IAM can't be enabled if there is at least one user already created. This parameter is only used at IAM configuration creation. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "delete_users")

    @delete_users.setter
    def delete_users(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "delete_users", value)

    @property
    @pulumi.getter(name="iamEnabled")
    def iam_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "iam_enabled")

    @iam_enabled.setter
    def iam_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "iam_enabled", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the Managed Private Registry. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)


@pulumi.type_token("ovh:index/cloudProjectContainerregistryIam:CloudProjectContainerregistryIam")
class CloudProjectContainerregistryIam(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_users: Optional[pulumi.Input[builtins.bool]] = None,
                 registry_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates an IAM configuration in an OVHcloud Managed Private Registry.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_iam = ovh.CloudProjectContainerregistryIam("my_iam",
            service_name="XXXXXX",
            registry_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
            delete_users=False)
        pulumi.export("iam-enabled", my_iam.iam_enabled)
        ```

        ## Import

        OVHcloud Managed Private Registry IAM can be imported using the tenant `service_name` and registry id `registry_id` separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:index/cloudProjectContainerregistryIam:CloudProjectContainerregistryIam my-iam service_name/registry_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] delete_users: Delete existing users from Harbor. IAM can't be enabled if there is at least one user already created. This parameter is only used at IAM configuration creation. **Changing this value recreates the resource.**
        :param pulumi.Input[builtins.str] registry_id: The ID of the Managed Private Registry. **Changing this value recreates the resource.**
        :param pulumi.Input[builtins.str] service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudProjectContainerregistryIamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an IAM configuration in an OVHcloud Managed Private Registry.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_iam = ovh.CloudProjectContainerregistryIam("my_iam",
            service_name="XXXXXX",
            registry_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
            delete_users=False)
        pulumi.export("iam-enabled", my_iam.iam_enabled)
        ```

        ## Import

        OVHcloud Managed Private Registry IAM can be imported using the tenant `service_name` and registry id `registry_id` separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:index/cloudProjectContainerregistryIam:CloudProjectContainerregistryIam my-iam service_name/registry_id
        ```

        :param str resource_name: The name of the resource.
        :param CloudProjectContainerregistryIamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudProjectContainerregistryIamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_users: Optional[pulumi.Input[builtins.bool]] = None,
                 registry_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudProjectContainerregistryIamArgs.__new__(CloudProjectContainerregistryIamArgs)

            __props__.__dict__["delete_users"] = delete_users
            if registry_id is None and not opts.urn:
                raise TypeError("Missing required property 'registry_id'")
            __props__.__dict__["registry_id"] = registry_id
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["iam_enabled"] = None
        super(CloudProjectContainerregistryIam, __self__).__init__(
            'ovh:index/cloudProjectContainerregistryIam:CloudProjectContainerregistryIam',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delete_users: Optional[pulumi.Input[builtins.bool]] = None,
            iam_enabled: Optional[pulumi.Input[builtins.bool]] = None,
            registry_id: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None) -> 'CloudProjectContainerregistryIam':
        """
        Get an existing CloudProjectContainerregistryIam resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] delete_users: Delete existing users from Harbor. IAM can't be enabled if there is at least one user already created. This parameter is only used at IAM configuration creation. **Changing this value recreates the resource.**
        :param pulumi.Input[builtins.str] registry_id: The ID of the Managed Private Registry. **Changing this value recreates the resource.**
        :param pulumi.Input[builtins.str] service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudProjectContainerregistryIamState.__new__(_CloudProjectContainerregistryIamState)

        __props__.__dict__["delete_users"] = delete_users
        __props__.__dict__["iam_enabled"] = iam_enabled
        __props__.__dict__["registry_id"] = registry_id
        __props__.__dict__["service_name"] = service_name
        return CloudProjectContainerregistryIam(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deleteUsers")
    def delete_users(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Delete existing users from Harbor. IAM can't be enabled if there is at least one user already created. This parameter is only used at IAM configuration creation. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "delete_users")

    @property
    @pulumi.getter(name="iamEnabled")
    def iam_enabled(self) -> pulumi.Output[builtins.bool]:
        return pulumi.get(self, "iam_enabled")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the Managed Private Registry. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "service_name")

