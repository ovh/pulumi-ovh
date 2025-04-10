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

__all__ = ['S3CredentialArgs', 'S3Credential']

@pulumi.input_type
class S3CredentialArgs:
    def __init__(__self__, *,
                 service_name: pulumi.Input[builtins.str],
                 user_id: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a S3Credential resource.
        :param pulumi.Input[builtins.str] service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] user_id: The ID of a public cloud project's user.
        """
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of a public cloud project's user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _S3CredentialState:
    def __init__(__self__, *,
                 access_key_id: Optional[pulumi.Input[builtins.str]] = None,
                 internal_user_id: Optional[pulumi.Input[builtins.str]] = None,
                 secret_access_key: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 user_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering S3Credential resources.
        :param pulumi.Input[builtins.str] access_key_id: the Access Key ID
        :param pulumi.Input[builtins.str] secret_access_key: (Sensitive) the Secret Access Key
        :param pulumi.Input[builtins.str] service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] user_id: The ID of a public cloud project's user.
        """
        if access_key_id is not None:
            pulumi.set(__self__, "access_key_id", access_key_id)
        if internal_user_id is not None:
            pulumi.set(__self__, "internal_user_id", internal_user_id)
        if secret_access_key is not None:
            pulumi.set(__self__, "secret_access_key", secret_access_key)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="accessKeyId")
    def access_key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        the Access Key ID
        """
        return pulumi.get(self, "access_key_id")

    @access_key_id.setter
    def access_key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_key_id", value)

    @property
    @pulumi.getter(name="internalUserId")
    def internal_user_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "internal_user_id")

    @internal_user_id.setter
    def internal_user_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "internal_user_id", value)

    @property
    @pulumi.getter(name="secretAccessKey")
    def secret_access_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        (Sensitive) the Secret Access Key
        """
        return pulumi.get(self, "secret_access_key")

    @secret_access_key.setter
    def secret_access_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secret_access_key", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of a public cloud project's user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "user_id", value)


class S3Credential(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 user_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates an S3 Credential for a user in a public cloud project.

        ## Import

        OVHcloud User S3 Credentials can be imported using the `service_name`, `user_id` and `access_key_id` of the credential, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/s3Credential:S3Credential s3_credential service_name/user_id/access_key_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] user_id: The ID of a public cloud project's user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: S3CredentialArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an S3 Credential for a user in a public cloud project.

        ## Import

        OVHcloud User S3 Credentials can be imported using the `service_name`, `user_id` and `access_key_id` of the credential, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/s3Credential:S3Credential s3_credential service_name/user_id/access_key_id
        ```

        :param str resource_name: The name of the resource.
        :param S3CredentialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(S3CredentialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 user_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = S3CredentialArgs.__new__(S3CredentialArgs)

            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["access_key_id"] = None
            __props__.__dict__["internal_user_id"] = None
            __props__.__dict__["secret_access_key"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretAccessKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(S3Credential, __self__).__init__(
            'ovh:CloudProject/s3Credential:S3Credential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key_id: Optional[pulumi.Input[builtins.str]] = None,
            internal_user_id: Optional[pulumi.Input[builtins.str]] = None,
            secret_access_key: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            user_id: Optional[pulumi.Input[builtins.str]] = None) -> 'S3Credential':
        """
        Get an existing S3Credential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] access_key_id: the Access Key ID
        :param pulumi.Input[builtins.str] secret_access_key: (Sensitive) the Secret Access Key
        :param pulumi.Input[builtins.str] service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] user_id: The ID of a public cloud project's user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _S3CredentialState.__new__(_S3CredentialState)

        __props__.__dict__["access_key_id"] = access_key_id
        __props__.__dict__["internal_user_id"] = internal_user_id
        __props__.__dict__["secret_access_key"] = secret_access_key
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["user_id"] = user_id
        return S3Credential(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKeyId")
    def access_key_id(self) -> pulumi.Output[builtins.str]:
        """
        the Access Key ID
        """
        return pulumi.get(self, "access_key_id")

    @property
    @pulumi.getter(name="internalUserId")
    def internal_user_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "internal_user_id")

    @property
    @pulumi.getter(name="secretAccessKey")
    def secret_access_key(self) -> pulumi.Output[builtins.str]:
        """
        (Sensitive) the Secret Access Key
        """
        return pulumi.get(self, "secret_access_key")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of a public cloud project's user.
        """
        return pulumi.get(self, "user_id")

