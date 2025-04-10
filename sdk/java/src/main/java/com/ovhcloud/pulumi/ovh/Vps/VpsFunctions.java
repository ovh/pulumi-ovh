// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Vps;

import com.ovhcloud.pulumi.ovh.Utilities;
import com.ovhcloud.pulumi.ovh.Vps.inputs.GetVpsArgs;
import com.ovhcloud.pulumi.ovh.Vps.inputs.GetVpsPlainArgs;
import com.ovhcloud.pulumi.ovh.Vps.outputs.GetVpsResult;
import com.ovhcloud.pulumi.ovh.Vps.outputs.GetVpssResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.resources.InvokeArgs;
import java.util.concurrent.CompletableFuture;

public final class VpsFunctions {
    /**
     * Use this data source to retrieve information about a vps associated with your OVHcloud Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
     * import com.pulumi.ovh.Vps.inputs.GetVpsArgs;
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
     *         final var server = VpsFunctions.getVps(GetVpsArgs.builder()
     *             .serviceName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetVpsResult> getVps(GetVpsArgs args) {
        return getVps(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about a vps associated with your OVHcloud Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
     * import com.pulumi.ovh.Vps.inputs.GetVpsArgs;
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
     *         final var server = VpsFunctions.getVps(GetVpsArgs.builder()
     *             .serviceName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetVpsResult> getVpsPlain(GetVpsPlainArgs args) {
        return getVpsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about a vps associated with your OVHcloud Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
     * import com.pulumi.ovh.Vps.inputs.GetVpsArgs;
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
     *         final var server = VpsFunctions.getVps(GetVpsArgs.builder()
     *             .serviceName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetVpsResult> getVps(GetVpsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Vps/getVps:getVps", TypeShape.of(GetVpsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about a vps associated with your OVHcloud Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
     * import com.pulumi.ovh.Vps.inputs.GetVpsArgs;
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
     *         final var server = VpsFunctions.getVps(GetVpsArgs.builder()
     *             .serviceName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetVpsResult> getVpsPlain(GetVpsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Vps/getVps:getVps", TypeShape.of(GetVpsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the list of VPS associated with your OVH Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
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
     *         final var servers = VpsFunctions.getVpss(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference);
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetVpssResult> getVpss() {
        return getVpss(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the list of VPS associated with your OVH Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
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
     *         final var servers = VpsFunctions.getVpss(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference);
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetVpssResult> getVpssPlain() {
        return getVpssPlain(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the list of VPS associated with your OVH Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
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
     *         final var servers = VpsFunctions.getVpss(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference);
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetVpssResult> getVpss(InvokeArgs args) {
        return getVpss(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the list of VPS associated with your OVH Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
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
     *         final var servers = VpsFunctions.getVpss(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference);
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetVpssResult> getVpssPlain(InvokeArgs args) {
        return getVpssPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the list of VPS associated with your OVH Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
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
     *         final var servers = VpsFunctions.getVpss(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference);
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetVpssResult> getVpss(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Vps/getVpss:getVpss", TypeShape.of(GetVpssResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the list of VPS associated with your OVH Account.
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
     * import com.pulumi.ovh.Vps.VpsFunctions;
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
     *         final var servers = VpsFunctions.getVpss(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference);
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetVpssResult> getVpssPlain(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Vps/getVpss:getVpss", TypeShape.of(GetVpssResult.class), args, Utilities.withVersion(options));
    }
}
