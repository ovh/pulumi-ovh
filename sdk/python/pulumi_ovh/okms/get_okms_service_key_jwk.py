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
from . import outputs

__all__ = [
    'GetOkmsServiceKeyJwkResult',
    'AwaitableGetOkmsServiceKeyJwkResult',
    'get_okms_service_key_jwk',
    'get_okms_service_key_jwk_output',
]

@pulumi.output_type
class GetOkmsServiceKeyJwkResult:
    """
    A collection of values returned by getOkmsServiceKeyJwk.
    """
    def __init__(__self__, created_at=None, iam=None, id=None, keys=None, name=None, okms_id=None, size=None, state=None, type=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if iam and not isinstance(iam, dict):
            raise TypeError("Expected argument 'iam' to be a dict")
        pulumi.set(__self__, "iam", iam)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if okms_id and not isinstance(okms_id, str):
            raise TypeError("Expected argument 'okms_id' to be a str")
        pulumi.set(__self__, "okms_id", okms_id)
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        pulumi.set(__self__, "size", size)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def iam(self) -> 'outputs.GetOkmsServiceKeyJwkIamResult':
        return pulumi.get(self, "iam")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keys(self) -> Sequence['outputs.GetOkmsServiceKeyJwkKeyResult']:
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="okmsId")
    def okms_id(self) -> builtins.str:
        return pulumi.get(self, "okms_id")

    @property
    @pulumi.getter
    def size(self) -> builtins.float:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def state(self) -> builtins.str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        return pulumi.get(self, "type")


class AwaitableGetOkmsServiceKeyJwkResult(GetOkmsServiceKeyJwkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOkmsServiceKeyJwkResult(
            created_at=self.created_at,
            iam=self.iam,
            id=self.id,
            keys=self.keys,
            name=self.name,
            okms_id=self.okms_id,
            size=self.size,
            state=self.state,
            type=self.type)


def get_okms_service_key_jwk(id: Optional[builtins.str] = None,
                             okms_id: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOkmsServiceKeyJwkResult:
    """
    Use this data source to retrieve information about a KMS service key, in the JWK format.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    key_info = ovh.Okms.get_okms_service_key(okms_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
    ```


    :param builtins.str id: ID of the service key
    :param builtins.str okms_id: ID of the KMS
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['okmsId'] = okms_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Okms/getOkmsServiceKeyJwk:getOkmsServiceKeyJwk', __args__, opts=opts, typ=GetOkmsServiceKeyJwkResult).value

    return AwaitableGetOkmsServiceKeyJwkResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        iam=pulumi.get(__ret__, 'iam'),
        id=pulumi.get(__ret__, 'id'),
        keys=pulumi.get(__ret__, 'keys'),
        name=pulumi.get(__ret__, 'name'),
        okms_id=pulumi.get(__ret__, 'okms_id'),
        size=pulumi.get(__ret__, 'size'),
        state=pulumi.get(__ret__, 'state'),
        type=pulumi.get(__ret__, 'type'))
def get_okms_service_key_jwk_output(id: Optional[pulumi.Input[builtins.str]] = None,
                                    okms_id: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOkmsServiceKeyJwkResult]:
    """
    Use this data source to retrieve information about a KMS service key, in the JWK format.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    key_info = ovh.Okms.get_okms_service_key(okms_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
    ```


    :param builtins.str id: ID of the service key
    :param builtins.str okms_id: ID of the KMS
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['okmsId'] = okms_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Okms/getOkmsServiceKeyJwk:getOkmsServiceKeyJwk', __args__, opts=opts, typ=GetOkmsServiceKeyJwkResult)
    return __ret__.apply(lambda __response__: GetOkmsServiceKeyJwkResult(
        created_at=pulumi.get(__response__, 'created_at'),
        iam=pulumi.get(__response__, 'iam'),
        id=pulumi.get(__response__, 'id'),
        keys=pulumi.get(__response__, 'keys'),
        name=pulumi.get(__response__, 'name'),
        okms_id=pulumi.get(__response__, 'okms_id'),
        size=pulumi.get(__response__, 'size'),
        state=pulumi.get(__response__, 'state'),
        type=pulumi.get(__response__, 'type')))
