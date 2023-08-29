using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Ovh = Pulumi.Ovh;
using System;

return await Deployment.RunAsync(() => 
{

    var config = new Pulumi.Config();
    var serviceName = config.Require("serviceName");

    System.Console.WriteLine(serviceName);

    var myKubeCluster = Ovh.CloudProject.GetKube.Invoke(new()
    {
        ServiceName = serviceName,
        KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx",
    });

    return new Dictionary<string, object?>
    {
        ["version"] = myKubeCluster.Apply(getKubeResult => getKubeResult.Version),
    };
});