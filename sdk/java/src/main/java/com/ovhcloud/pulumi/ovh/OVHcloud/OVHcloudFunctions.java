// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.OVHcloud;

import com.ovhcloud.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
import com.ovhcloud.pulumi.ovh.OVHcloud.inputs.ConnectPlainArgs;
import com.ovhcloud.pulumi.ovh.OVHcloud.outputs.ConnectResult;
import com.ovhcloud.pulumi.ovh.OVHcloud.outputs.ConnectsResult;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.deployment.InvokeOutputOptions;
import com.pulumi.resources.InvokeArgs;
import java.util.concurrent.CompletableFuture;

public final class OVHcloudFunctions {
    /**
     * Use this data source to retrieve information about an Ovhcloud Connect product
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occ = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .serviceName("XXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<ConnectResult> connect(ConnectArgs args) {
        return connect(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an Ovhcloud Connect product
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occ = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .serviceName("XXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<ConnectResult> connectPlain(ConnectPlainArgs args) {
        return connectPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an Ovhcloud Connect product
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occ = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .serviceName("XXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<ConnectResult> connect(ConnectArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:OVHcloud/connect:Connect", TypeShape.of(ConnectResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an Ovhcloud Connect product
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occ = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .serviceName("XXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<ConnectResult> connect(ConnectArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("ovh:OVHcloud/connect:Connect", TypeShape.of(ConnectResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an Ovhcloud Connect product
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occ = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .serviceName("XXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<ConnectResult> connectPlain(ConnectPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:OVHcloud/connect:Connect", TypeShape.of(ConnectResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get the details of your Ovhcloud Connect products.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occs = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<ConnectsResult> connects() {
        return connects(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Get the details of your Ovhcloud Connect products.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occs = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<ConnectsResult> connectsPlain() {
        return connectsPlain(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Get the details of your Ovhcloud Connect products.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occs = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<ConnectsResult> connects(InvokeArgs args) {
        return connects(args, InvokeOptions.Empty);
    }
    /**
     * Get the details of your Ovhcloud Connect products.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occs = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<ConnectsResult> connectsPlain(InvokeArgs args) {
        return connectsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get the details of your Ovhcloud Connect products.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occs = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<ConnectsResult> connects(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:OVHcloud/connects:Connects", TypeShape.of(ConnectsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get the details of your Ovhcloud Connect products.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occs = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<ConnectsResult> connects(InvokeArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("ovh:OVHcloud/connects:Connects", TypeShape.of(ConnectsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get the details of your Ovhcloud Connect products.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.OVHcloud.OVHcloudFunctions;
     * import com.pulumi.ovh.OVHcloud.inputs.ConnectArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var occs = OVHcloudFunctions.Connect(ConnectArgs.builder()
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<ConnectsResult> connectsPlain(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:OVHcloud/connects:Connects", TypeShape.of(ConnectsResult.class), args, Utilities.withVersion(options));
    }
}
