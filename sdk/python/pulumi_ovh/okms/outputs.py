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

__all__ = [
    'OkmsIam',
    'ServiceKeyJWKIam',
    'ServiceKeyJWKKey',
    'GetOkmsResourceIamResult',
    'GetOkmsServiceKeyIamResult',
    'GetOkmsServiceKeyJwkIamResult',
    'GetOkmsServiceKeyJwkKeyResult',
    'GetOkmsServiceKeyPemIamResult',
    'GetOkmsServiceKeyPemKeysPemResult',
    'GetOvhCloudConnectIamResult',
    'GetOvhCloudConnectsOccResult',
    'GetOvhCloudConnectsOccIamResult',
]

@pulumi.output_type
class OkmsIam(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "displayName":
            suggest = "display_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OkmsIam. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OkmsIam.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OkmsIam.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 display_name: Optional[str] = None,
                 id: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 urn: Optional[str] = None):
        """
        :param str display_name: Resource display name
        :param str id: Unique identifier of the resource
        :param Mapping[str, str] tags: Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: Unique resource name used in policies
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if urn is not None:
            pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> Optional[str]:
        """
        Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class ServiceKeyJWKIam(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "displayName":
            suggest = "display_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceKeyJWKIam. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceKeyJWKIam.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceKeyJWKIam.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 display_name: Optional[str] = None,
                 id: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 urn: Optional[str] = None):
        """
        :param str display_name: Resource display name
        :param str id: Unique identifier of the resource
        :param Mapping[str, str] tags: Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: Unique resource name used in policies
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if urn is not None:
            pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> Optional[str]:
        """
        Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class ServiceKeyJWKKey(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "keyOps":
            suggest = "key_ops"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceKeyJWKKey. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceKeyJWKKey.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceKeyJWKKey.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 key_ops: Sequence[str],
                 kty: str,
                 alg: Optional[str] = None,
                 crv: Optional[str] = None,
                 d: Optional[str] = None,
                 dp: Optional[str] = None,
                 dq: Optional[str] = None,
                 e: Optional[str] = None,
                 k: Optional[str] = None,
                 kid: Optional[str] = None,
                 n: Optional[str] = None,
                 p: Optional[str] = None,
                 q: Optional[str] = None,
                 qi: Optional[str] = None,
                 use: Optional[str] = None,
                 x: Optional[str] = None,
                 y: Optional[str] = None):
        """
        :param Sequence[str] key_ops: The operation for which the key is intended to be used
        :param str kty: Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC
        :param str alg: The algorithm intended to be used with the key
        :param str crv: The cryptographic curve used with the key
        :param str d: The RSA or EC private exponent
        :param str dp: The RSA private key's first factor CRT exponent
        :param str dq: The RSA private key's second factor CRT exponent
        :param str e: The exponent value for the RSA public key
        :param str k: The value of the symmetric (or other single-valued) key
        :param str kid: key ID parameter used to match a specific key
        :param str n: The modulus value for the RSA public key
        :param str p: The first prime factor of the RSA private key
        :param str q: The second prime factor of the RSA private key
        :param str qi: The CRT coefficient of the second factor of the RSA private key
        :param str use: The intended use of the public key
        :param str x: The x coordinate for the Elliptic Curve point
        :param str y: The y coordinate for the Elliptic Curve point
        """
        pulumi.set(__self__, "key_ops", key_ops)
        pulumi.set(__self__, "kty", kty)
        if alg is not None:
            pulumi.set(__self__, "alg", alg)
        if crv is not None:
            pulumi.set(__self__, "crv", crv)
        if d is not None:
            pulumi.set(__self__, "d", d)
        if dp is not None:
            pulumi.set(__self__, "dp", dp)
        if dq is not None:
            pulumi.set(__self__, "dq", dq)
        if e is not None:
            pulumi.set(__self__, "e", e)
        if k is not None:
            pulumi.set(__self__, "k", k)
        if kid is not None:
            pulumi.set(__self__, "kid", kid)
        if n is not None:
            pulumi.set(__self__, "n", n)
        if p is not None:
            pulumi.set(__self__, "p", p)
        if q is not None:
            pulumi.set(__self__, "q", q)
        if qi is not None:
            pulumi.set(__self__, "qi", qi)
        if use is not None:
            pulumi.set(__self__, "use", use)
        if x is not None:
            pulumi.set(__self__, "x", x)
        if y is not None:
            pulumi.set(__self__, "y", y)

    @property
    @pulumi.getter(name="keyOps")
    def key_ops(self) -> Sequence[str]:
        """
        The operation for which the key is intended to be used
        """
        return pulumi.get(self, "key_ops")

    @property
    @pulumi.getter
    def kty(self) -> str:
        """
        Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC
        """
        return pulumi.get(self, "kty")

    @property
    @pulumi.getter
    def alg(self) -> Optional[str]:
        """
        The algorithm intended to be used with the key
        """
        return pulumi.get(self, "alg")

    @property
    @pulumi.getter
    def crv(self) -> Optional[str]:
        """
        The cryptographic curve used with the key
        """
        return pulumi.get(self, "crv")

    @property
    @pulumi.getter
    def d(self) -> Optional[str]:
        """
        The RSA or EC private exponent
        """
        return pulumi.get(self, "d")

    @property
    @pulumi.getter
    def dp(self) -> Optional[str]:
        """
        The RSA private key's first factor CRT exponent
        """
        return pulumi.get(self, "dp")

    @property
    @pulumi.getter
    def dq(self) -> Optional[str]:
        """
        The RSA private key's second factor CRT exponent
        """
        return pulumi.get(self, "dq")

    @property
    @pulumi.getter
    def e(self) -> Optional[str]:
        """
        The exponent value for the RSA public key
        """
        return pulumi.get(self, "e")

    @property
    @pulumi.getter
    def k(self) -> Optional[str]:
        """
        The value of the symmetric (or other single-valued) key
        """
        return pulumi.get(self, "k")

    @property
    @pulumi.getter
    def kid(self) -> Optional[str]:
        """
        key ID parameter used to match a specific key
        """
        return pulumi.get(self, "kid")

    @property
    @pulumi.getter
    def n(self) -> Optional[str]:
        """
        The modulus value for the RSA public key
        """
        return pulumi.get(self, "n")

    @property
    @pulumi.getter
    def p(self) -> Optional[str]:
        """
        The first prime factor of the RSA private key
        """
        return pulumi.get(self, "p")

    @property
    @pulumi.getter
    def q(self) -> Optional[str]:
        """
        The second prime factor of the RSA private key
        """
        return pulumi.get(self, "q")

    @property
    @pulumi.getter
    def qi(self) -> Optional[str]:
        """
        The CRT coefficient of the second factor of the RSA private key
        """
        return pulumi.get(self, "qi")

    @property
    @pulumi.getter
    def use(self) -> Optional[str]:
        """
        The intended use of the public key
        """
        return pulumi.get(self, "use")

    @property
    @pulumi.getter
    def x(self) -> Optional[str]:
        """
        The x coordinate for the Elliptic Curve point
        """
        return pulumi.get(self, "x")

    @property
    @pulumi.getter
    def y(self) -> Optional[str]:
        """
        The y coordinate for the Elliptic Curve point
        """
        return pulumi.get(self, "y")


@pulumi.output_type
class GetOkmsResourceIamResult(dict):
    def __init__(__self__, *,
                 display_name: str,
                 id: str,
                 tags: Mapping[str, str],
                 urn: str):
        """
        :param str display_name: Resource display name
        :param str id: Unique identifier of the resource
        :param Mapping[str, str] tags: Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: Unique resource name used in policies
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class GetOkmsServiceKeyIamResult(dict):
    def __init__(__self__, *,
                 display_name: str,
                 id: str,
                 tags: Mapping[str, str],
                 urn: str):
        """
        :param str display_name: Resource display name
        :param str id: Unique identifier of the resource
        :param Mapping[str, str] tags: Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: Unique resource name used in policies
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class GetOkmsServiceKeyJwkIamResult(dict):
    def __init__(__self__, *,
                 display_name: str,
                 id: str,
                 tags: Mapping[str, str],
                 urn: str):
        """
        :param str display_name: Resource display name
        :param str id: Unique identifier of the resource
        :param Mapping[str, str] tags: Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: Unique resource name used in policies
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class GetOkmsServiceKeyJwkKeyResult(dict):
    def __init__(__self__, *,
                 alg: str,
                 crv: str,
                 e: str,
                 key_ops: Sequence[str],
                 kid: str,
                 kty: str,
                 n: str,
                 use: str,
                 x: str,
                 y: str):
        """
        :param str alg: The algorithm intended to be used with the key
        :param str crv: The cryptographic curve used with the key
        :param str e: The exponent value for the RSA public key
        :param Sequence[str] key_ops: The operation for which the key is intended to be used
        :param str kid: key ID parameter used to match a specific key
        :param str kty: Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC
        :param str n: The modulus value for the RSA public key
        :param str use: The intended use of the public key
        :param str x: The x coordinate for the Elliptic Curve point
        :param str y: The y coordinate for the Elliptic Curve point
        """
        pulumi.set(__self__, "alg", alg)
        pulumi.set(__self__, "crv", crv)
        pulumi.set(__self__, "e", e)
        pulumi.set(__self__, "key_ops", key_ops)
        pulumi.set(__self__, "kid", kid)
        pulumi.set(__self__, "kty", kty)
        pulumi.set(__self__, "n", n)
        pulumi.set(__self__, "use", use)
        pulumi.set(__self__, "x", x)
        pulumi.set(__self__, "y", y)

    @property
    @pulumi.getter
    def alg(self) -> str:
        """
        The algorithm intended to be used with the key
        """
        return pulumi.get(self, "alg")

    @property
    @pulumi.getter
    def crv(self) -> str:
        """
        The cryptographic curve used with the key
        """
        return pulumi.get(self, "crv")

    @property
    @pulumi.getter
    def e(self) -> str:
        """
        The exponent value for the RSA public key
        """
        return pulumi.get(self, "e")

    @property
    @pulumi.getter(name="keyOps")
    def key_ops(self) -> Sequence[str]:
        """
        The operation for which the key is intended to be used
        """
        return pulumi.get(self, "key_ops")

    @property
    @pulumi.getter
    def kid(self) -> str:
        """
        key ID parameter used to match a specific key
        """
        return pulumi.get(self, "kid")

    @property
    @pulumi.getter
    def kty(self) -> str:
        """
        Key type parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC
        """
        return pulumi.get(self, "kty")

    @property
    @pulumi.getter
    def n(self) -> str:
        """
        The modulus value for the RSA public key
        """
        return pulumi.get(self, "n")

    @property
    @pulumi.getter
    def use(self) -> str:
        """
        The intended use of the public key
        """
        return pulumi.get(self, "use")

    @property
    @pulumi.getter
    def x(self) -> str:
        """
        The x coordinate for the Elliptic Curve point
        """
        return pulumi.get(self, "x")

    @property
    @pulumi.getter
    def y(self) -> str:
        """
        The y coordinate for the Elliptic Curve point
        """
        return pulumi.get(self, "y")


@pulumi.output_type
class GetOkmsServiceKeyPemIamResult(dict):
    def __init__(__self__, *,
                 display_name: str,
                 id: str,
                 tags: Mapping[str, str],
                 urn: str):
        """
        :param str display_name: Resource display name
        :param str id: Unique identifier of the resource
        :param Mapping[str, str] tags: Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: Unique resource name used in policies
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class GetOkmsServiceKeyPemKeysPemResult(dict):
    def __init__(__self__, *,
                 pem: str):
        """
        :param str pem: The key in base64 encoded PEM format
        """
        pulumi.set(__self__, "pem", pem)

    @property
    @pulumi.getter
    def pem(self) -> str:
        """
        The key in base64 encoded PEM format
        """
        return pulumi.get(self, "pem")


@pulumi.output_type
class GetOvhCloudConnectIamResult(dict):
    def __init__(__self__, *,
                 display_name: str,
                 id: str,
                 tags: Mapping[str, str],
                 urn: str):
        """
        :param str display_name: Resource display name
        :param str id: Unique identifier of the resource
        :param Mapping[str, str] tags: Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: Unique resource name used in policies
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class GetOvhCloudConnectsOccResult(dict):
    def __init__(__self__, *,
                 bandwidth: str,
                 description: str,
                 iam: 'outputs.GetOvhCloudConnectsOccIamResult',
                 interface_lists: Sequence[float],
                 pop: str,
                 port_quantity: str,
                 product: str,
                 provider_name: str,
                 service_name: str,
                 status: str,
                 uuid: str,
                 vrack: str):
        """
        :param str bandwidth: Service bandwidth
        :param str description: Service description
        :param 'GetOvhCloudConnectsOccIamArgs' iam: IAM resource metadata
        :param Sequence[float] interface_lists: List of interfaces linked to a service
        :param str pop: Pop reference where the service is delivered
        :param str port_quantity: Port quantity
        :param str product: Product name of the service
        :param str provider_name: Service provider
        :param str service_name: Service name
        :param str status: Service status
        :param str uuid: Service UUID
        :param str vrack: vrack linked to the service
        """
        pulumi.set(__self__, "bandwidth", bandwidth)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "iam", iam)
        pulumi.set(__self__, "interface_lists", interface_lists)
        pulumi.set(__self__, "pop", pop)
        pulumi.set(__self__, "port_quantity", port_quantity)
        pulumi.set(__self__, "product", product)
        pulumi.set(__self__, "provider_name", provider_name)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "uuid", uuid)
        pulumi.set(__self__, "vrack", vrack)

    @property
    @pulumi.getter
    def bandwidth(self) -> str:
        """
        Service bandwidth
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Service description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def iam(self) -> 'outputs.GetOvhCloudConnectsOccIamResult':
        """
        IAM resource metadata
        """
        return pulumi.get(self, "iam")

    @property
    @pulumi.getter(name="interfaceLists")
    def interface_lists(self) -> Sequence[float]:
        """
        List of interfaces linked to a service
        """
        return pulumi.get(self, "interface_lists")

    @property
    @pulumi.getter
    def pop(self) -> str:
        """
        Pop reference where the service is delivered
        """
        return pulumi.get(self, "pop")

    @property
    @pulumi.getter(name="portQuantity")
    def port_quantity(self) -> str:
        """
        Port quantity
        """
        return pulumi.get(self, "port_quantity")

    @property
    @pulumi.getter
    def product(self) -> str:
        """
        Product name of the service
        """
        return pulumi.get(self, "product")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> str:
        """
        Service provider
        """
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Service status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        Service UUID
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vrack(self) -> str:
        """
        vrack linked to the service
        """
        return pulumi.get(self, "vrack")


@pulumi.output_type
class GetOvhCloudConnectsOccIamResult(dict):
    def __init__(__self__, *,
                 display_name: str,
                 id: str,
                 tags: Mapping[str, str],
                 urn: str):
        """
        :param str display_name: Resource display name
        :param str id: Unique identifier of the resource
        :param Mapping[str, str] tags: Resource tags. Tags that were internally computed are prefixed with ovh:
        :param str urn: Unique resource name used in policies
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


