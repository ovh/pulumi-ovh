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
    'GetUserS3CredentialResult',
    'AwaitableGetUserS3CredentialResult',
    'get_user_s3_credential',
    'get_user_s3_credential_output',
]

@pulumi.output_type
class GetUserS3CredentialResult:
    """
    A collection of values returned by getUserS3Credential.
    """
    def __init__(__self__, access_key_id=None, id=None, secret_access_key=None, service_name=None, user_id=None):
        if access_key_id and not isinstance(access_key_id, str):
            raise TypeError("Expected argument 'access_key_id' to be a str")
        pulumi.set(__self__, "access_key_id", access_key_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if secret_access_key and not isinstance(secret_access_key, str):
            raise TypeError("Expected argument 'secret_access_key' to be a str")
        pulumi.set(__self__, "secret_access_key", secret_access_key)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="accessKeyId")
    def access_key_id(self) -> builtins.str:
        return pulumi.get(self, "access_key_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="secretAccessKey")
    def secret_access_key(self) -> builtins.str:
        """
        (Sensitive) the Secret Access Key
        """
        return pulumi.get(self, "secret_access_key")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> builtins.str:
        return pulumi.get(self, "user_id")


class AwaitableGetUserS3CredentialResult(GetUserS3CredentialResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserS3CredentialResult(
            access_key_id=self.access_key_id,
            id=self.id,
            secret_access_key=self.secret_access_key,
            service_name=self.service_name,
            user_id=self.user_id)


def get_user_s3_credential(access_key_id: Optional[builtins.str] = None,
                           service_name: Optional[builtins.str] = None,
                           user_id: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserS3CredentialResult:
    """
    Use this data source to retrieve the Secret Access Key of an Access Key ID associated with a public cloud project's user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    project_users = ovh.CloudProject.get_users(service_name="XXX")
    # Get the user ID of a previously created user with the description "S3-User"
    users = [user.user_id for user in project_users.users if user.description == "S3-User"]
    s3_user_id = users[0]
    my_s3_credentials = ovh.CloudProject.get_user_s3_credentials(service_name=project_users.service_name,
        user_id=s3_user_id)
    my_s3_credential = ovh.CloudProject.get_user_s3_credential(service_name=my_s3_credentials.service_name,
        user_id=my_s3_credentials.user_id,
        access_key_id=my_s3_credentials.access_key_ids[0])
    pulumi.export("myAccessKeyId", my_s3_credential.access_key_id)
    pulumi.export("mySecretAccessKey", my_s3_credential.secret_access_key)
    ```


    :param builtins.str access_key_id: the Access Key ID
    :param builtins.str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    :param builtins.str user_id: The ID of a public cloud project's user.
    """
    __args__ = dict()
    __args__['accessKeyId'] = access_key_id
    __args__['serviceName'] = service_name
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getUserS3Credential:getUserS3Credential', __args__, opts=opts, typ=GetUserS3CredentialResult).value

    return AwaitableGetUserS3CredentialResult(
        access_key_id=pulumi.get(__ret__, 'access_key_id'),
        id=pulumi.get(__ret__, 'id'),
        secret_access_key=pulumi.get(__ret__, 'secret_access_key'),
        service_name=pulumi.get(__ret__, 'service_name'),
        user_id=pulumi.get(__ret__, 'user_id'))
def get_user_s3_credential_output(access_key_id: Optional[pulumi.Input[builtins.str]] = None,
                                  service_name: Optional[pulumi.Input[builtins.str]] = None,
                                  user_id: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserS3CredentialResult]:
    """
    Use this data source to retrieve the Secret Access Key of an Access Key ID associated with a public cloud project's user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    project_users = ovh.CloudProject.get_users(service_name="XXX")
    # Get the user ID of a previously created user with the description "S3-User"
    users = [user.user_id for user in project_users.users if user.description == "S3-User"]
    s3_user_id = users[0]
    my_s3_credentials = ovh.CloudProject.get_user_s3_credentials(service_name=project_users.service_name,
        user_id=s3_user_id)
    my_s3_credential = ovh.CloudProject.get_user_s3_credential(service_name=my_s3_credentials.service_name,
        user_id=my_s3_credentials.user_id,
        access_key_id=my_s3_credentials.access_key_ids[0])
    pulumi.export("myAccessKeyId", my_s3_credential.access_key_id)
    pulumi.export("mySecretAccessKey", my_s3_credential.secret_access_key)
    ```


    :param builtins.str access_key_id: the Access Key ID
    :param builtins.str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    :param builtins.str user_id: The ID of a public cloud project's user.
    """
    __args__ = dict()
    __args__['accessKeyId'] = access_key_id
    __args__['serviceName'] = service_name
    __args__['userId'] = user_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getUserS3Credential:getUserS3Credential', __args__, opts=opts, typ=GetUserS3CredentialResult)
    return __ret__.apply(lambda __response__: GetUserS3CredentialResult(
        access_key_id=pulumi.get(__response__, 'access_key_id'),
        id=pulumi.get(__response__, 'id'),
        secret_access_key=pulumi.get(__response__, 'secret_access_key'),
        service_name=pulumi.get(__response__, 'service_name'),
        user_id=pulumi.get(__response__, 'user_id')))
