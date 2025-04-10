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
    'GetKubeOidcResult',
    'AwaitableGetKubeOidcResult',
    'get_kube_oidc',
    'get_kube_oidc_output',
]

@pulumi.output_type
class GetKubeOidcResult:
    """
    A collection of values returned by getKubeOidc.
    """
    def __init__(__self__, client_id=None, id=None, issuer_url=None, kube_id=None, oidc_ca_content=None, oidc_groups_claims=None, oidc_groups_prefix=None, oidc_required_claims=None, oidc_signing_algs=None, oidc_username_claim=None, oidc_username_prefix=None, service_name=None):
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if issuer_url and not isinstance(issuer_url, str):
            raise TypeError("Expected argument 'issuer_url' to be a str")
        pulumi.set(__self__, "issuer_url", issuer_url)
        if kube_id and not isinstance(kube_id, str):
            raise TypeError("Expected argument 'kube_id' to be a str")
        pulumi.set(__self__, "kube_id", kube_id)
        if oidc_ca_content and not isinstance(oidc_ca_content, str):
            raise TypeError("Expected argument 'oidc_ca_content' to be a str")
        pulumi.set(__self__, "oidc_ca_content", oidc_ca_content)
        if oidc_groups_claims and not isinstance(oidc_groups_claims, list):
            raise TypeError("Expected argument 'oidc_groups_claims' to be a list")
        pulumi.set(__self__, "oidc_groups_claims", oidc_groups_claims)
        if oidc_groups_prefix and not isinstance(oidc_groups_prefix, str):
            raise TypeError("Expected argument 'oidc_groups_prefix' to be a str")
        pulumi.set(__self__, "oidc_groups_prefix", oidc_groups_prefix)
        if oidc_required_claims and not isinstance(oidc_required_claims, list):
            raise TypeError("Expected argument 'oidc_required_claims' to be a list")
        pulumi.set(__self__, "oidc_required_claims", oidc_required_claims)
        if oidc_signing_algs and not isinstance(oidc_signing_algs, list):
            raise TypeError("Expected argument 'oidc_signing_algs' to be a list")
        pulumi.set(__self__, "oidc_signing_algs", oidc_signing_algs)
        if oidc_username_claim and not isinstance(oidc_username_claim, str):
            raise TypeError("Expected argument 'oidc_username_claim' to be a str")
        pulumi.set(__self__, "oidc_username_claim", oidc_username_claim)
        if oidc_username_prefix and not isinstance(oidc_username_prefix, str):
            raise TypeError("Expected argument 'oidc_username_prefix' to be a str")
        pulumi.set(__self__, "oidc_username_prefix", oidc_username_prefix)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[builtins.str]:
        """
        The OIDC client ID.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="issuerUrl")
    def issuer_url(self) -> Optional[builtins.str]:
        """
        The OIDC issuer url.
        """
        return pulumi.get(self, "issuer_url")

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "kube_id")

    @property
    @pulumi.getter(name="oidcCaContent")
    def oidc_ca_content(self) -> Optional[builtins.str]:
        """
        Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
        """
        return pulumi.get(self, "oidc_ca_content")

    @property
    @pulumi.getter(name="oidcGroupsClaims")
    def oidc_groups_claims(self) -> Optional[Sequence[builtins.str]]:
        """
        Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
        """
        return pulumi.get(self, "oidc_groups_claims")

    @property
    @pulumi.getter(name="oidcGroupsPrefix")
    def oidc_groups_prefix(self) -> Optional[builtins.str]:
        """
        Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
        """
        return pulumi.get(self, "oidc_groups_prefix")

    @property
    @pulumi.getter(name="oidcRequiredClaims")
    def oidc_required_claims(self) -> Optional[Sequence[builtins.str]]:
        """
        Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
        """
        return pulumi.get(self, "oidc_required_claims")

    @property
    @pulumi.getter(name="oidcSigningAlgs")
    def oidc_signing_algs(self) -> Optional[Sequence[builtins.str]]:
        """
        Array of signing algorithms accepted. Default is \\"RS256\\".
        """
        return pulumi.get(self, "oidc_signing_algs")

    @property
    @pulumi.getter(name="oidcUsernameClaim")
    def oidc_username_claim(self) -> Optional[builtins.str]:
        """
        JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
        """
        return pulumi.get(self, "oidc_username_claim")

    @property
    @pulumi.getter(name="oidcUsernamePrefix")
    def oidc_username_prefix(self) -> Optional[builtins.str]:
        """
        Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidc_username_claim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
        """
        return pulumi.get(self, "oidc_username_prefix")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetKubeOidcResult(GetKubeOidcResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubeOidcResult(
            client_id=self.client_id,
            id=self.id,
            issuer_url=self.issuer_url,
            kube_id=self.kube_id,
            oidc_ca_content=self.oidc_ca_content,
            oidc_groups_claims=self.oidc_groups_claims,
            oidc_groups_prefix=self.oidc_groups_prefix,
            oidc_required_claims=self.oidc_required_claims,
            oidc_signing_algs=self.oidc_signing_algs,
            oidc_username_claim=self.oidc_username_claim,
            oidc_username_prefix=self.oidc_username_prefix,
            service_name=self.service_name)


def get_kube_oidc(client_id: Optional[builtins.str] = None,
                  issuer_url: Optional[builtins.str] = None,
                  kube_id: Optional[builtins.str] = None,
                  oidc_ca_content: Optional[builtins.str] = None,
                  oidc_groups_claims: Optional[Sequence[builtins.str]] = None,
                  oidc_groups_prefix: Optional[builtins.str] = None,
                  oidc_required_claims: Optional[Sequence[builtins.str]] = None,
                  oidc_signing_algs: Optional[Sequence[builtins.str]] = None,
                  oidc_username_claim: Optional[builtins.str] = None,
                  oidc_username_prefix: Optional[builtins.str] = None,
                  service_name: Optional[builtins.str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubeOidcResult:
    """
    Use this data source to get a OVHcloud Managed Kubernetes Service cluster OIDC.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    oidc = ovh.CloudProject.get_kube_oidc(service_name="XXXXXX",
        kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx")
    pulumi.export("oidc-val", oidc.client_id)
    ```


    :param builtins.str client_id: The OIDC client ID.
    :param builtins.str issuer_url: The OIDC issuer url.
    :param builtins.str kube_id: The id of the managed kubernetes cluster.
    :param builtins.str oidc_ca_content: Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
    :param Sequence[builtins.str] oidc_groups_claims: Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
    :param builtins.str oidc_groups_prefix: Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
    :param Sequence[builtins.str] oidc_required_claims: Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
    :param Sequence[builtins.str] oidc_signing_algs: Array of signing algorithms accepted. Default is \\"RS256\\".
    :param builtins.str oidc_username_claim: JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
    :param builtins.str oidc_username_prefix: Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidc_username_claim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clientId'] = client_id
    __args__['issuerUrl'] = issuer_url
    __args__['kubeId'] = kube_id
    __args__['oidcCaContent'] = oidc_ca_content
    __args__['oidcGroupsClaims'] = oidc_groups_claims
    __args__['oidcGroupsPrefix'] = oidc_groups_prefix
    __args__['oidcRequiredClaims'] = oidc_required_claims
    __args__['oidcSigningAlgs'] = oidc_signing_algs
    __args__['oidcUsernameClaim'] = oidc_username_claim
    __args__['oidcUsernamePrefix'] = oidc_username_prefix
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getKubeOidc:getKubeOidc', __args__, opts=opts, typ=GetKubeOidcResult).value

    return AwaitableGetKubeOidcResult(
        client_id=pulumi.get(__ret__, 'client_id'),
        id=pulumi.get(__ret__, 'id'),
        issuer_url=pulumi.get(__ret__, 'issuer_url'),
        kube_id=pulumi.get(__ret__, 'kube_id'),
        oidc_ca_content=pulumi.get(__ret__, 'oidc_ca_content'),
        oidc_groups_claims=pulumi.get(__ret__, 'oidc_groups_claims'),
        oidc_groups_prefix=pulumi.get(__ret__, 'oidc_groups_prefix'),
        oidc_required_claims=pulumi.get(__ret__, 'oidc_required_claims'),
        oidc_signing_algs=pulumi.get(__ret__, 'oidc_signing_algs'),
        oidc_username_claim=pulumi.get(__ret__, 'oidc_username_claim'),
        oidc_username_prefix=pulumi.get(__ret__, 'oidc_username_prefix'),
        service_name=pulumi.get(__ret__, 'service_name'))
def get_kube_oidc_output(client_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                         issuer_url: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                         kube_id: Optional[pulumi.Input[builtins.str]] = None,
                         oidc_ca_content: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                         oidc_groups_claims: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                         oidc_groups_prefix: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                         oidc_required_claims: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                         oidc_signing_algs: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                         oidc_username_claim: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                         oidc_username_prefix: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                         service_name: Optional[pulumi.Input[builtins.str]] = None,
                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetKubeOidcResult]:
    """
    Use this data source to get a OVHcloud Managed Kubernetes Service cluster OIDC.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    oidc = ovh.CloudProject.get_kube_oidc(service_name="XXXXXX",
        kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx")
    pulumi.export("oidc-val", oidc.client_id)
    ```


    :param builtins.str client_id: The OIDC client ID.
    :param builtins.str issuer_url: The OIDC issuer url.
    :param builtins.str kube_id: The id of the managed kubernetes cluster.
    :param builtins.str oidc_ca_content: Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
    :param Sequence[builtins.str] oidc_groups_claims: Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
    :param builtins.str oidc_groups_prefix: Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
    :param Sequence[builtins.str] oidc_required_claims: Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
    :param Sequence[builtins.str] oidc_signing_algs: Array of signing algorithms accepted. Default is \\"RS256\\".
    :param builtins.str oidc_username_claim: JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
    :param builtins.str oidc_username_prefix: Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidc_username_claim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clientId'] = client_id
    __args__['issuerUrl'] = issuer_url
    __args__['kubeId'] = kube_id
    __args__['oidcCaContent'] = oidc_ca_content
    __args__['oidcGroupsClaims'] = oidc_groups_claims
    __args__['oidcGroupsPrefix'] = oidc_groups_prefix
    __args__['oidcRequiredClaims'] = oidc_required_claims
    __args__['oidcSigningAlgs'] = oidc_signing_algs
    __args__['oidcUsernameClaim'] = oidc_username_claim
    __args__['oidcUsernamePrefix'] = oidc_username_prefix
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getKubeOidc:getKubeOidc', __args__, opts=opts, typ=GetKubeOidcResult)
    return __ret__.apply(lambda __response__: GetKubeOidcResult(
        client_id=pulumi.get(__response__, 'client_id'),
        id=pulumi.get(__response__, 'id'),
        issuer_url=pulumi.get(__response__, 'issuer_url'),
        kube_id=pulumi.get(__response__, 'kube_id'),
        oidc_ca_content=pulumi.get(__response__, 'oidc_ca_content'),
        oidc_groups_claims=pulumi.get(__response__, 'oidc_groups_claims'),
        oidc_groups_prefix=pulumi.get(__response__, 'oidc_groups_prefix'),
        oidc_required_claims=pulumi.get(__response__, 'oidc_required_claims'),
        oidc_signing_algs=pulumi.get(__response__, 'oidc_signing_algs'),
        oidc_username_claim=pulumi.get(__response__, 'oidc_username_claim'),
        oidc_username_prefix=pulumi.get(__response__, 'oidc_username_prefix'),
        service_name=pulumi.get(__response__, 'service_name')))
