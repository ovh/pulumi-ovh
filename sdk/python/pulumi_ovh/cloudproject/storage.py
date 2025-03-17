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
from ._inputs import *

__all__ = ['StorageArgs', 'Storage']

@pulumi.input_type
class StorageArgs:
    def __init__(__self__, *,
                 region_name: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 encryption: Optional[pulumi.Input['StorageEncryptionArgs']] = None,
                 limit: Optional[pulumi.Input[float]] = None,
                 marker: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[float]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 replication: Optional[pulumi.Input['StorageReplicationArgs']] = None,
                 versioning: Optional[pulumi.Input['StorageVersioningArgs']] = None):
        """
        The set of arguments for constructing a Storage resource.
        :param pulumi.Input[str] region_name: Region name
        :param pulumi.Input[str] service_name: Service name
        :param pulumi.Input['StorageEncryptionArgs'] encryption: Encryption configuration
        :param pulumi.Input[float] limit: Limit the number of objects returned (1000 maximum, defaults to 1000)
        :param pulumi.Input[str] marker: Key to start with when listing objects
        :param pulumi.Input[str] name: Container name
        :param pulumi.Input[float] owner_id: Container owner user ID
        :param pulumi.Input[str] prefix: List objects whose key begins with this prefix
        :param pulumi.Input['StorageReplicationArgs'] replication: Replication configuration
        :param pulumi.Input['StorageVersioningArgs'] versioning: Versioning configuration
        """
        pulumi.set(__self__, "region_name", region_name)
        pulumi.set(__self__, "service_name", service_name)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if limit is not None:
            pulumi.set(__self__, "limit", limit)
        if marker is not None:
            pulumi.set(__self__, "marker", marker)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if replication is not None:
            pulumi.set(__self__, "replication", replication)
        if versioning is not None:
            pulumi.set(__self__, "versioning", versioning)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Input[str]:
        """
        Region name
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "region_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input['StorageEncryptionArgs']]:
        """
        Encryption configuration
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input['StorageEncryptionArgs']]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter
    def limit(self) -> Optional[pulumi.Input[float]]:
        """
        Limit the number of objects returned (1000 maximum, defaults to 1000)
        """
        return pulumi.get(self, "limit")

    @limit.setter
    def limit(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "limit", value)

    @property
    @pulumi.getter
    def marker(self) -> Optional[pulumi.Input[str]]:
        """
        Key to start with when listing objects
        """
        return pulumi.get(self, "marker")

    @marker.setter
    def marker(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "marker", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Container name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[float]]:
        """
        Container owner user ID
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        List objects whose key begins with this prefix
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def replication(self) -> Optional[pulumi.Input['StorageReplicationArgs']]:
        """
        Replication configuration
        """
        return pulumi.get(self, "replication")

    @replication.setter
    def replication(self, value: Optional[pulumi.Input['StorageReplicationArgs']]):
        pulumi.set(self, "replication", value)

    @property
    @pulumi.getter
    def versioning(self) -> Optional[pulumi.Input['StorageVersioningArgs']]:
        """
        Versioning configuration
        """
        return pulumi.get(self, "versioning")

    @versioning.setter
    def versioning(self, value: Optional[pulumi.Input['StorageVersioningArgs']]):
        pulumi.set(self, "versioning", value)


@pulumi.input_type
class _StorageState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input['StorageEncryptionArgs']] = None,
                 limit: Optional[pulumi.Input[float]] = None,
                 marker: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 objects: Optional[pulumi.Input[Sequence[pulumi.Input['StorageObjectArgs']]]] = None,
                 objects_count: Optional[pulumi.Input[float]] = None,
                 objects_size: Optional[pulumi.Input[float]] = None,
                 owner_id: Optional[pulumi.Input[float]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 replication: Optional[pulumi.Input['StorageReplicationArgs']] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 versioning: Optional[pulumi.Input['StorageVersioningArgs']] = None,
                 virtual_host: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Storage resources.
        :param pulumi.Input[str] created_at: The date and timestamp when the resource was created
        :param pulumi.Input['StorageEncryptionArgs'] encryption: Encryption configuration
        :param pulumi.Input[float] limit: Limit the number of objects returned (1000 maximum, defaults to 1000)
        :param pulumi.Input[str] marker: Key to start with when listing objects
        :param pulumi.Input[str] name: Container name
        :param pulumi.Input[Sequence[pulumi.Input['StorageObjectArgs']]] objects: Container objects
        :param pulumi.Input[float] objects_count: Container total objects count
        :param pulumi.Input[float] objects_size: Container total objects size (bytes)
        :param pulumi.Input[float] owner_id: Container owner user ID
        :param pulumi.Input[str] prefix: List objects whose key begins with this prefix
        :param pulumi.Input[str] region: Container region
        :param pulumi.Input[str] region_name: Region name
        :param pulumi.Input['StorageReplicationArgs'] replication: Replication configuration
        :param pulumi.Input[str] service_name: Service name
        :param pulumi.Input['StorageVersioningArgs'] versioning: Versioning configuration
        :param pulumi.Input[str] virtual_host: Container virtual host
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if limit is not None:
            pulumi.set(__self__, "limit", limit)
        if marker is not None:
            pulumi.set(__self__, "marker", marker)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if objects is not None:
            pulumi.set(__self__, "objects", objects)
        if objects_count is not None:
            pulumi.set(__self__, "objects_count", objects_count)
        if objects_size is not None:
            pulumi.set(__self__, "objects_size", objects_size)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)
        if replication is not None:
            pulumi.set(__self__, "replication", replication)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if versioning is not None:
            pulumi.set(__self__, "versioning", versioning)
        if virtual_host is not None:
            pulumi.set(__self__, "virtual_host", virtual_host)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and timestamp when the resource was created
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input['StorageEncryptionArgs']]:
        """
        Encryption configuration
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input['StorageEncryptionArgs']]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter
    def limit(self) -> Optional[pulumi.Input[float]]:
        """
        Limit the number of objects returned (1000 maximum, defaults to 1000)
        """
        return pulumi.get(self, "limit")

    @limit.setter
    def limit(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "limit", value)

    @property
    @pulumi.getter
    def marker(self) -> Optional[pulumi.Input[str]]:
        """
        Key to start with when listing objects
        """
        return pulumi.get(self, "marker")

    @marker.setter
    def marker(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "marker", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Container name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def objects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StorageObjectArgs']]]]:
        """
        Container objects
        """
        return pulumi.get(self, "objects")

    @objects.setter
    def objects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StorageObjectArgs']]]]):
        pulumi.set(self, "objects", value)

    @property
    @pulumi.getter(name="objectsCount")
    def objects_count(self) -> Optional[pulumi.Input[float]]:
        """
        Container total objects count
        """
        return pulumi.get(self, "objects_count")

    @objects_count.setter
    def objects_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "objects_count", value)

    @property
    @pulumi.getter(name="objectsSize")
    def objects_size(self) -> Optional[pulumi.Input[float]]:
        """
        Container total objects size (bytes)
        """
        return pulumi.get(self, "objects_size")

    @objects_size.setter
    def objects_size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "objects_size", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[float]]:
        """
        Container owner user ID
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        List objects whose key begins with this prefix
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Container region
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[pulumi.Input[str]]:
        """
        Region name
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_name", value)

    @property
    @pulumi.getter
    def replication(self) -> Optional[pulumi.Input['StorageReplicationArgs']]:
        """
        Replication configuration
        """
        return pulumi.get(self, "replication")

    @replication.setter
    def replication(self, value: Optional[pulumi.Input['StorageReplicationArgs']]):
        pulumi.set(self, "replication", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def versioning(self) -> Optional[pulumi.Input['StorageVersioningArgs']]:
        """
        Versioning configuration
        """
        return pulumi.get(self, "versioning")

    @versioning.setter
    def versioning(self, value: Optional[pulumi.Input['StorageVersioningArgs']]):
        pulumi.set(self, "versioning", value)

    @property
    @pulumi.getter(name="virtualHost")
    def virtual_host(self) -> Optional[pulumi.Input[str]]:
        """
        Container virtual host
        """
        return pulumi.get(self, "virtual_host")

    @virtual_host.setter
    def virtual_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_host", value)


class Storage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encryption: Optional[pulumi.Input[Union['StorageEncryptionArgs', 'StorageEncryptionArgsDict']]] = None,
                 limit: Optional[pulumi.Input[float]] = None,
                 marker: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[float]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 replication: Optional[pulumi.Input[Union['StorageReplicationArgs', 'StorageReplicationArgsDict']]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 versioning: Optional[pulumi.Input[Union['StorageVersioningArgs', 'StorageVersioningArgsDict']]] = None,
                 __props__=None):
        """
        Create S3™* compatible storage container
        (* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        storage = ovh.cloud_project.Storage("storage",
            region_name="GRA",
            service_name="<public cloud project ID>",
            versioning={
                "status": "enabled",
            })
        ```

        ## Import

        A storage in a public cloud project can be imported using the `service_name`, `region_name` and `name` attributes.

        Using the following configuration:

        hcl

        import {

          id = "<service_name>/<region_name>/<name>"

          to = ovh_cloud_project_storage.storage

        }

        You can then run:

        bash

        $ pulumi preview -generate-config-out=storage.tf

        $ pulumi up

        The file `storage.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above.

        See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['StorageEncryptionArgs', 'StorageEncryptionArgsDict']] encryption: Encryption configuration
        :param pulumi.Input[float] limit: Limit the number of objects returned (1000 maximum, defaults to 1000)
        :param pulumi.Input[str] marker: Key to start with when listing objects
        :param pulumi.Input[str] name: Container name
        :param pulumi.Input[float] owner_id: Container owner user ID
        :param pulumi.Input[str] prefix: List objects whose key begins with this prefix
        :param pulumi.Input[str] region_name: Region name
        :param pulumi.Input[Union['StorageReplicationArgs', 'StorageReplicationArgsDict']] replication: Replication configuration
        :param pulumi.Input[str] service_name: Service name
        :param pulumi.Input[Union['StorageVersioningArgs', 'StorageVersioningArgsDict']] versioning: Versioning configuration
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StorageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create S3™* compatible storage container
        (* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        storage = ovh.cloud_project.Storage("storage",
            region_name="GRA",
            service_name="<public cloud project ID>",
            versioning={
                "status": "enabled",
            })
        ```

        ## Import

        A storage in a public cloud project can be imported using the `service_name`, `region_name` and `name` attributes.

        Using the following configuration:

        hcl

        import {

          id = "<service_name>/<region_name>/<name>"

          to = ovh_cloud_project_storage.storage

        }

        You can then run:

        bash

        $ pulumi preview -generate-config-out=storage.tf

        $ pulumi up

        The file `storage.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above.

        See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.

        :param str resource_name: The name of the resource.
        :param StorageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StorageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encryption: Optional[pulumi.Input[Union['StorageEncryptionArgs', 'StorageEncryptionArgsDict']]] = None,
                 limit: Optional[pulumi.Input[float]] = None,
                 marker: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[float]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 replication: Optional[pulumi.Input[Union['StorageReplicationArgs', 'StorageReplicationArgsDict']]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 versioning: Optional[pulumi.Input[Union['StorageVersioningArgs', 'StorageVersioningArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StorageArgs.__new__(StorageArgs)

            __props__.__dict__["encryption"] = encryption
            __props__.__dict__["limit"] = limit
            __props__.__dict__["marker"] = marker
            __props__.__dict__["name"] = name
            __props__.__dict__["owner_id"] = owner_id
            __props__.__dict__["prefix"] = prefix
            if region_name is None and not opts.urn:
                raise TypeError("Missing required property 'region_name'")
            __props__.__dict__["region_name"] = region_name
            __props__.__dict__["replication"] = replication
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["versioning"] = versioning
            __props__.__dict__["created_at"] = None
            __props__.__dict__["objects"] = None
            __props__.__dict__["objects_count"] = None
            __props__.__dict__["objects_size"] = None
            __props__.__dict__["region"] = None
            __props__.__dict__["virtual_host"] = None
        super(Storage, __self__).__init__(
            'ovh:CloudProject/storage:Storage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            encryption: Optional[pulumi.Input[Union['StorageEncryptionArgs', 'StorageEncryptionArgsDict']]] = None,
            limit: Optional[pulumi.Input[float]] = None,
            marker: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            objects: Optional[pulumi.Input[Sequence[pulumi.Input[Union['StorageObjectArgs', 'StorageObjectArgsDict']]]]] = None,
            objects_count: Optional[pulumi.Input[float]] = None,
            objects_size: Optional[pulumi.Input[float]] = None,
            owner_id: Optional[pulumi.Input[float]] = None,
            prefix: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            region_name: Optional[pulumi.Input[str]] = None,
            replication: Optional[pulumi.Input[Union['StorageReplicationArgs', 'StorageReplicationArgsDict']]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            versioning: Optional[pulumi.Input[Union['StorageVersioningArgs', 'StorageVersioningArgsDict']]] = None,
            virtual_host: Optional[pulumi.Input[str]] = None) -> 'Storage':
        """
        Get an existing Storage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and timestamp when the resource was created
        :param pulumi.Input[Union['StorageEncryptionArgs', 'StorageEncryptionArgsDict']] encryption: Encryption configuration
        :param pulumi.Input[float] limit: Limit the number of objects returned (1000 maximum, defaults to 1000)
        :param pulumi.Input[str] marker: Key to start with when listing objects
        :param pulumi.Input[str] name: Container name
        :param pulumi.Input[Sequence[pulumi.Input[Union['StorageObjectArgs', 'StorageObjectArgsDict']]]] objects: Container objects
        :param pulumi.Input[float] objects_count: Container total objects count
        :param pulumi.Input[float] objects_size: Container total objects size (bytes)
        :param pulumi.Input[float] owner_id: Container owner user ID
        :param pulumi.Input[str] prefix: List objects whose key begins with this prefix
        :param pulumi.Input[str] region: Container region
        :param pulumi.Input[str] region_name: Region name
        :param pulumi.Input[Union['StorageReplicationArgs', 'StorageReplicationArgsDict']] replication: Replication configuration
        :param pulumi.Input[str] service_name: Service name
        :param pulumi.Input[Union['StorageVersioningArgs', 'StorageVersioningArgsDict']] versioning: Versioning configuration
        :param pulumi.Input[str] virtual_host: Container virtual host
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StorageState.__new__(_StorageState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["encryption"] = encryption
        __props__.__dict__["limit"] = limit
        __props__.__dict__["marker"] = marker
        __props__.__dict__["name"] = name
        __props__.__dict__["objects"] = objects
        __props__.__dict__["objects_count"] = objects_count
        __props__.__dict__["objects_size"] = objects_size
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["prefix"] = prefix
        __props__.__dict__["region"] = region
        __props__.__dict__["region_name"] = region_name
        __props__.__dict__["replication"] = replication
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["versioning"] = versioning
        __props__.__dict__["virtual_host"] = virtual_host
        return Storage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and timestamp when the resource was created
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output['outputs.StorageEncryption']:
        """
        Encryption configuration
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter
    def limit(self) -> pulumi.Output[float]:
        """
        Limit the number of objects returned (1000 maximum, defaults to 1000)
        """
        return pulumi.get(self, "limit")

    @property
    @pulumi.getter
    def marker(self) -> pulumi.Output[str]:
        """
        Key to start with when listing objects
        """
        return pulumi.get(self, "marker")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Container name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def objects(self) -> pulumi.Output[Sequence['outputs.StorageObject']]:
        """
        Container objects
        """
        return pulumi.get(self, "objects")

    @property
    @pulumi.getter(name="objectsCount")
    def objects_count(self) -> pulumi.Output[float]:
        """
        Container total objects count
        """
        return pulumi.get(self, "objects_count")

    @property
    @pulumi.getter(name="objectsSize")
    def objects_size(self) -> pulumi.Output[float]:
        """
        Container total objects size (bytes)
        """
        return pulumi.get(self, "objects_size")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[float]:
        """
        Container owner user ID
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[str]:
        """
        List objects whose key begins with this prefix
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Container region
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Output[str]:
        """
        Region name
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter
    def replication(self) -> pulumi.Output['outputs.StorageReplication']:
        """
        Replication configuration
        """
        return pulumi.get(self, "replication")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        Service name
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def versioning(self) -> pulumi.Output['outputs.StorageVersioning']:
        """
        Versioning configuration
        """
        return pulumi.get(self, "versioning")

    @property
    @pulumi.getter(name="virtualHost")
    def virtual_host(self) -> pulumi.Output[str]:
        """
        Container virtual host
        """
        return pulumi.get(self, "virtual_host")

