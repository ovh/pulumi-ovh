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
    'GetKafkaTopicResult',
    'AwaitableGetKafkaTopicResult',
    'get_kafka_topic',
    'get_kafka_topic_output',
]

@pulumi.output_type
class GetKafkaTopicResult:
    """
    A collection of values returned by getKafkaTopic.
    """
    def __init__(__self__, cluster_id=None, id=None, min_insync_replicas=None, name=None, partitions=None, replication=None, retention_bytes=None, retention_hours=None, service_name=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if min_insync_replicas and not isinstance(min_insync_replicas, int):
            raise TypeError("Expected argument 'min_insync_replicas' to be a int")
        pulumi.set(__self__, "min_insync_replicas", min_insync_replicas)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partitions and not isinstance(partitions, int):
            raise TypeError("Expected argument 'partitions' to be a int")
        pulumi.set(__self__, "partitions", partitions)
        if replication and not isinstance(replication, int):
            raise TypeError("Expected argument 'replication' to be a int")
        pulumi.set(__self__, "replication", replication)
        if retention_bytes and not isinstance(retention_bytes, int):
            raise TypeError("Expected argument 'retention_bytes' to be a int")
        pulumi.set(__self__, "retention_bytes", retention_bytes)
        if retention_hours and not isinstance(retention_hours, int):
            raise TypeError("Expected argument 'retention_hours' to be a int")
        pulumi.set(__self__, "retention_hours", retention_hours)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="minInsyncReplicas")
    def min_insync_replicas(self) -> builtins.int:
        """
        Minimum insync replica accepted for this topic.
        """
        return pulumi.get(self, "min_insync_replicas")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of the topic.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partitions(self) -> builtins.int:
        """
        Number of partitions for this topic.
        """
        return pulumi.get(self, "partitions")

    @property
    @pulumi.getter
    def replication(self) -> builtins.int:
        """
        Number of replication for this topic.
        """
        return pulumi.get(self, "replication")

    @property
    @pulumi.getter(name="retentionBytes")
    def retention_bytes(self) -> builtins.int:
        """
        Number of bytes for the retention of the data for this topic. Inferior to 0 mean Unlimited
        """
        return pulumi.get(self, "retention_bytes")

    @property
    @pulumi.getter(name="retentionHours")
    def retention_hours(self) -> builtins.int:
        """
        Number of hours for the retention of the data for this topic. Inferior to 0 mean Unlimited
        """
        return pulumi.get(self, "retention_hours")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetKafkaTopicResult(GetKafkaTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKafkaTopicResult(
            cluster_id=self.cluster_id,
            id=self.id,
            min_insync_replicas=self.min_insync_replicas,
            name=self.name,
            partitions=self.partitions,
            replication=self.replication,
            retention_bytes=self.retention_bytes,
            retention_hours=self.retention_hours,
            service_name=self.service_name)


def get_kafka_topic(cluster_id: Optional[builtins.str] = None,
                    id: Optional[builtins.str] = None,
                    service_name: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKafkaTopicResult:
    """
    Use this data source to get information about a topic of a kafka cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    topic = ovh.CloudProjectDatabase.get_kafka_topic(service_name="XXX",
        cluster_id="YYY",
        id="ZZZ")
    pulumi.export("topicName", topic.name)
    ```


    :param builtins.str cluster_id: Cluster ID
    :param builtins.str id: Topic ID
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['id'] = id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getKafkaTopic:getKafkaTopic', __args__, opts=opts, typ=GetKafkaTopicResult).value

    return AwaitableGetKafkaTopicResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        id=pulumi.get(__ret__, 'id'),
        min_insync_replicas=pulumi.get(__ret__, 'min_insync_replicas'),
        name=pulumi.get(__ret__, 'name'),
        partitions=pulumi.get(__ret__, 'partitions'),
        replication=pulumi.get(__ret__, 'replication'),
        retention_bytes=pulumi.get(__ret__, 'retention_bytes'),
        retention_hours=pulumi.get(__ret__, 'retention_hours'),
        service_name=pulumi.get(__ret__, 'service_name'))
def get_kafka_topic_output(cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                           id: Optional[pulumi.Input[builtins.str]] = None,
                           service_name: Optional[pulumi.Input[builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetKafkaTopicResult]:
    """
    Use this data source to get information about a topic of a kafka cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    topic = ovh.CloudProjectDatabase.get_kafka_topic(service_name="XXX",
        cluster_id="YYY",
        id="ZZZ")
    pulumi.export("topicName", topic.name)
    ```


    :param builtins.str cluster_id: Cluster ID
    :param builtins.str id: Topic ID
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['id'] = id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProjectDatabase/getKafkaTopic:getKafkaTopic', __args__, opts=opts, typ=GetKafkaTopicResult)
    return __ret__.apply(lambda __response__: GetKafkaTopicResult(
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        id=pulumi.get(__response__, 'id'),
        min_insync_replicas=pulumi.get(__response__, 'min_insync_replicas'),
        name=pulumi.get(__response__, 'name'),
        partitions=pulumi.get(__response__, 'partitions'),
        replication=pulumi.get(__response__, 'replication'),
        retention_bytes=pulumi.get(__response__, 'retention_bytes'),
        retention_hours=pulumi.get(__response__, 'retention_hours'),
        service_name=pulumi.get(__response__, 'service_name')))
