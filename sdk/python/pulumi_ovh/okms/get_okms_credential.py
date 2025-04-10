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

__all__ = [
    'GetOkmsCredentialResult',
    'AwaitableGetOkmsCredentialResult',
    'get_okms_credential',
    'get_okms_credential_output',
]

@pulumi.output_type
class GetOkmsCredentialResult:
    """
    A collection of values returned by getOkmsCredential.
    """
    def __init__(__self__, certificate_pem=None, created_at=None, description=None, expired_at=None, from_csr=None, id=None, identity_urns=None, name=None, okms_id=None, status=None):
        if certificate_pem and not isinstance(certificate_pem, str):
            raise TypeError("Expected argument 'certificate_pem' to be a str")
        pulumi.set(__self__, "certificate_pem", certificate_pem)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if expired_at and not isinstance(expired_at, str):
            raise TypeError("Expected argument 'expired_at' to be a str")
        pulumi.set(__self__, "expired_at", expired_at)
        if from_csr and not isinstance(from_csr, bool):
            raise TypeError("Expected argument 'from_csr' to be a bool")
        pulumi.set(__self__, "from_csr", from_csr)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity_urns and not isinstance(identity_urns, list):
            raise TypeError("Expected argument 'identity_urns' to be a list")
        pulumi.set(__self__, "identity_urns", identity_urns)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if okms_id and not isinstance(okms_id, str):
            raise TypeError("Expected argument 'okms_id' to be a str")
        pulumi.set(__self__, "okms_id", okms_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="certificatePem")
    def certificate_pem(self) -> builtins.str:
        """
        (String) PEM encoded certificate of the credential
        """
        return pulumi.get(self, "certificate_pem")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        """
        (String) Creation time of the credential
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        (String) Description of the credential
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expiredAt")
    def expired_at(self) -> builtins.str:
        """
        (String) Expiration time of the credential
        """
        return pulumi.get(self, "expired_at")

    @property
    @pulumi.getter(name="fromCsr")
    def from_csr(self) -> builtins.bool:
        """
        (Boolean) Is the credential generated from CSR
        """
        return pulumi.get(self, "from_csr")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identityUrns")
    def identity_urns(self) -> Sequence[builtins.str]:
        """
        (List of String) List of identity URNs associated with the credential
        """
        return pulumi.get(self, "identity_urns")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        (String) Name of the credential
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="okmsId")
    def okms_id(self) -> builtins.str:
        return pulumi.get(self, "okms_id")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        (String) Status of the credential
        """
        return pulumi.get(self, "status")


class AwaitableGetOkmsCredentialResult(GetOkmsCredentialResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOkmsCredentialResult(
            certificate_pem=self.certificate_pem,
            created_at=self.created_at,
            description=self.description,
            expired_at=self.expired_at,
            from_csr=self.from_csr,
            id=self.id,
            identity_urns=self.identity_urns,
            name=self.name,
            okms_id=self.okms_id,
            status=self.status)


def get_okms_credential(id: Optional[builtins.str] = None,
                        okms_id: Optional[builtins.str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOkmsCredentialResult:
    """
    Use this data source to retrieve data associated with a KMS credential, such as the PEM encoded certificate.


    :param builtins.str id: ID of the credential
    :param builtins.str okms_id: ID of the KMS
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['okmsId'] = okms_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Okms/getOkmsCredential:getOkmsCredential', __args__, opts=opts, typ=GetOkmsCredentialResult).value

    return AwaitableGetOkmsCredentialResult(
        certificate_pem=pulumi.get(__ret__, 'certificate_pem'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        expired_at=pulumi.get(__ret__, 'expired_at'),
        from_csr=pulumi.get(__ret__, 'from_csr'),
        id=pulumi.get(__ret__, 'id'),
        identity_urns=pulumi.get(__ret__, 'identity_urns'),
        name=pulumi.get(__ret__, 'name'),
        okms_id=pulumi.get(__ret__, 'okms_id'),
        status=pulumi.get(__ret__, 'status'))
def get_okms_credential_output(id: Optional[pulumi.Input[builtins.str]] = None,
                               okms_id: Optional[pulumi.Input[builtins.str]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOkmsCredentialResult]:
    """
    Use this data source to retrieve data associated with a KMS credential, such as the PEM encoded certificate.


    :param builtins.str id: ID of the credential
    :param builtins.str okms_id: ID of the KMS
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['okmsId'] = okms_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Okms/getOkmsCredential:getOkmsCredential', __args__, opts=opts, typ=GetOkmsCredentialResult)
    return __ret__.apply(lambda __response__: GetOkmsCredentialResult(
        certificate_pem=pulumi.get(__response__, 'certificate_pem'),
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        expired_at=pulumi.get(__response__, 'expired_at'),
        from_csr=pulumi.get(__response__, 'from_csr'),
        id=pulumi.get(__response__, 'id'),
        identity_urns=pulumi.get(__response__, 'identity_urns'),
        name=pulumi.get(__response__, 'name'),
        okms_id=pulumi.get(__response__, 'okms_id'),
        status=pulumi.get(__response__, 'status')))
