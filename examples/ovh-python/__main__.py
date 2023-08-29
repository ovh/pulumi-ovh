"""A Python Pulumi program"""

import pulumi
import pulumi_ovh as ovh

config = pulumi.Config();
service_name = config.require('serviceName')

print(service_name);

"""Get a Kubernetes cluster version"""
my_kube_cluster = ovh.cloudproject.get_kube(service_name=service_name,
    kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx");
pulumi.export("version", my_kube_cluster.version)