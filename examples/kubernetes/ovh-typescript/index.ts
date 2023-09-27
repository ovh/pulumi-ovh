import * as pulumi from "@pulumi/pulumi";
import * as ovh from "@ovh-devrelteam/pulumi-ovh"

let config = new pulumi.Config();
let serviceName = config.require("serviceName")

console.log(serviceName)

let myKubeCluster = ovh.cloudproject.getKube({
    serviceName: serviceName,
    kubeId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx"
}) 

export const version = myKubeCluster.then(myKubeCluster => myKubeCluster.version);