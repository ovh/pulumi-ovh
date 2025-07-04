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

__all__ = [
    'GetUserS3CredentialsResult',
    'AwaitableGetUserS3CredentialsResult',
    'get_user_s3_credentials',
    'get_user_s3_credentials_output',
]

@pulumi.output_type
class GetUserS3CredentialsResult:
    """
    A collection of values returned by getUserS3Credentials.
    """
    def __init__(__self__, access_key_ids=None, id=None, service_name=None, user_id=None):
        if access_key_ids and not isinstance(access_key_ids, list):
            raise TypeError("Expected argument 'access_key_ids' to be a list")
        pulumi.set(__self__, "access_key_ids", access_key_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="accessKeyIds")
    def access_key_ids(self) -> Sequence[builtins.str]:
        """
        The list of the Access Key ID associated with this user.
        """
        return pulumi.get(self, "access_key_ids")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> builtins.str:
        return pulumi.get(self, "user_id")


class AwaitableGetUserS3CredentialsResult(GetUserS3CredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserS3CredentialsResult(
            access_key_ids=self.access_key_ids,
            id=self.id,
            service_name=self.service_name,
            user_id=self.user_id)


def get_user_s3_credentials(service_name: Optional[builtins.str] = None,
                            user_id: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserS3CredentialsResult:
    """
    Use this data source to retrieve the list of all the S3 access_key_id associated with a public cloud project's user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_s3_credentials = ovh.CloudProject.get_user_s3_credentials(service_name="XXXXXX",
        user_id="1234")
    pulumi.export("accessKeyIds", my_s3_credentials.access_key_ids)
    ```


    :param builtins.str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    :param builtins.str user_id: The ID of a public cloud project's user.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getUserS3Credentials:getUserS3Credentials', __args__, opts=opts, typ=GetUserS3CredentialsResult).value

    return AwaitableGetUserS3CredentialsResult(
        access_key_ids=pulumi.get(__ret__, 'access_key_ids'),
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'),
        user_id=pulumi.get(__ret__, 'user_id'))
def get_user_s3_credentials_output(service_name: Optional[pulumi.Input[builtins.str]] = None,
                                   user_id: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserS3CredentialsResult]:
    """
    Use this data source to retrieve the list of all the S3 access_key_id associated with a public cloud project's user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_s3_credentials = ovh.CloudProject.get_user_s3_credentials(service_name="XXXXXX",
        user_id="1234")
    pulumi.export("accessKeyIds", my_s3_credentials.access_key_ids)
    ```


    :param builtins.str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    :param builtins.str user_id: The ID of a public cloud project's user.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['userId'] = user_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getUserS3Credentials:getUserS3Credentials', __args__, opts=opts, typ=GetUserS3CredentialsResult)
    return __ret__.apply(lambda __response__: GetUserS3CredentialsResult(
        access_key_ids=pulumi.get(__response__, 'access_key_ids'),
        id=pulumi.get(__response__, 'id'),
        service_name=pulumi.get(__response__, 'service_name'),
        user_id=pulumi.get(__response__, 'user_id')))
