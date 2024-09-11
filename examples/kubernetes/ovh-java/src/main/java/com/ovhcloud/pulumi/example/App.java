package com.ovhcloud.pulumi.example;



import com.ovhcloud.pulumi.ovh.CloudProject.Kube;
import com.ovhcloud.pulumi.ovh.CloudProject.KubeArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.KubeNodePool;
import com.ovhcloud.pulumi.ovh.CloudProject.KubeNodePoolArgs;
import com.pulumi.Context;
import com.pulumi.Pulumi;


/**
 * Minimal Pulumi example with Java SDK.
 */
public class App {

    private final static String OVH_CLOUD_PROJECT_SERVICE = System.getenv("OVH_CLOUD_PROJECT_SERVICE");

    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        KubeArgs kubeDetails = KubeArgs.builder()
                                        .serviceName(OVH_CLOUD_PROJECT_SERVICE)
                                        .name("pulumi-k8s")
                                        .region("GRA7")
                                    .build();
        Kube myKube = new Kube("my-kube", kubeDetails);

        KubeNodePoolArgs nodePoolDetails = KubeNodePoolArgs.builder()
                                            .serviceName(OVH_CLOUD_PROJECT_SERVICE)
                                            .flavorName("d2-4")
                                            .kubeId(myKube.id().asPlaintext())
                                            .minNodes(1)
                                            .maxNodes(1)
                                        .build();

        KubeNodePool nodePool = new KubeNodePool("pulumi-nodepool", nodePoolDetails);
        
        ctx.export("kubeconfig", myKube.kubeconfig());
    }
}
