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

__all__ = ['PrometheusArgs', 'Prometheus']

@pulumi.input_type
class PrometheusArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 engine: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 password_reset: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Prometheus resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] password_reset: Arbitrary string to change to trigger a password update
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "service_name", service_name)
        if password_reset is not None:
            pulumi.set(__self__, "password_reset", password_reset)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input[str]:
        """
        The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="passwordReset")
    def password_reset(self) -> Optional[pulumi.Input[str]]:
        """
        Arbitrary string to change to trigger a password update
        """
        return pulumi.get(self, "password_reset")

    @password_reset.setter
    def password_reset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_reset", value)


@pulumi.input_type
class _PrometheusState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 password_reset: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusTargetArgs']]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Prometheus resources.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
        :param pulumi.Input[str] password: (Sensitive) Password of the user.
        :param pulumi.Input[str] password_reset: Arbitrary string to change to trigger a password update
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[Sequence[pulumi.Input['PrometheusTargetArgs']]] targets: List of all endpoint targets.
               * `Host` - Host of the endpoint.
               * `Port` - Connection port for the endpoint.
        :param pulumi.Input[str] username: name of the prometheus user.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if password_reset is not None:
            pulumi.set(__self__, "password_reset", password_reset)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if targets is not None:
            pulumi.set(__self__, "targets", targets)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        (Sensitive) Password of the user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="passwordReset")
    def password_reset(self) -> Optional[pulumi.Input[str]]:
        """
        Arbitrary string to change to trigger a password update
        """
        return pulumi.get(self, "password_reset")

    @password_reset.setter
    def password_reset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_reset", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusTargetArgs']]]]:
        """
        List of all endpoint targets.
        * `Host` - Host of the endpoint.
        * `Port` - Connection port for the endpoint.
        """
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrometheusTargetArgs']]]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        name of the prometheus user.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Prometheus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 password_reset: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        OVHcloud Managed database clusters prometheus can be imported using the `service_name`, `engine` and `cluster_id`, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/prometheus:Prometheus my_prometheus service_name/engine/cluster_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
        :param pulumi.Input[str] password_reset: Arbitrary string to change to trigger a password update
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrometheusArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        OVHcloud Managed database clusters prometheus can be imported using the `service_name`, `engine` and `cluster_id`, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/prometheus:Prometheus my_prometheus service_name/engine/cluster_id
        ```

        :param str resource_name: The name of the resource.
        :param PrometheusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrometheusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 password_reset: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrometheusArgs.__new__(PrometheusArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            __props__.__dict__["password_reset"] = password_reset
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["password"] = None
            __props__.__dict__["targets"] = None
            __props__.__dict__["username"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Prometheus, __self__).__init__(
            'ovh:CloudProjectDatabase/prometheus:Prometheus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            password_reset: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            targets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrometheusTargetArgs', 'PrometheusTargetArgsDict']]]]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'Prometheus':
        """
        Get an existing Prometheus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
        :param pulumi.Input[str] password: (Sensitive) Password of the user.
        :param pulumi.Input[str] password_reset: Arbitrary string to change to trigger a password update
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrometheusTargetArgs', 'PrometheusTargetArgsDict']]]] targets: List of all endpoint targets.
               * `Host` - Host of the endpoint.
               * `Port` - Connection port for the endpoint.
        :param pulumi.Input[str] username: name of the prometheus user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrometheusState.__new__(_PrometheusState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["engine"] = engine
        __props__.__dict__["password"] = password
        __props__.__dict__["password_reset"] = password_reset
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["targets"] = targets
        __props__.__dict__["username"] = username
        return Prometheus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        (Sensitive) Password of the user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="passwordReset")
    def password_reset(self) -> pulumi.Output[Optional[str]]:
        """
        Arbitrary string to change to trigger a password update
        """
        return pulumi.get(self, "password_reset")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[Sequence['outputs.PrometheusTarget']]:
        """
        List of all endpoint targets.
        * `Host` - Host of the endpoint.
        * `Port` - Connection port for the endpoint.
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        name of the prometheus user.
        """
        return pulumi.get(self, "username")

