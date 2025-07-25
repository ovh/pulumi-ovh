// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Hosting;

import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseAllowlistArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseAllowlistPlainArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseDbArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseDbPlainArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabasePlainArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserGrantArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserGrantPlainArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserPlainArgs;
import com.ovhcloud.pulumi.ovh.Hosting.outputs.GetPrivateDatabaseAllowlistResult;
import com.ovhcloud.pulumi.ovh.Hosting.outputs.GetPrivateDatabaseDbResult;
import com.ovhcloud.pulumi.ovh.Hosting.outputs.GetPrivateDatabaseResult;
import com.ovhcloud.pulumi.ovh.Hosting.outputs.GetPrivateDatabaseUserGrantResult;
import com.ovhcloud.pulumi.ovh.Hosting.outputs.GetPrivateDatabaseUserResult;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class HostingFunctions {
    /**
     * Use this data source to retrieve information about an hosting database.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseArgs;
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
     *         final var database = HostingFunctions.getPrivateDatabase(GetPrivateDatabaseArgs.builder()
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
    public static Output<GetPrivateDatabaseResult> getPrivateDatabase(GetPrivateDatabaseArgs args) {
        return getPrivateDatabase(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting database.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseArgs;
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
     *         final var database = HostingFunctions.getPrivateDatabase(GetPrivateDatabaseArgs.builder()
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
    public static CompletableFuture<GetPrivateDatabaseResult> getPrivateDatabasePlain(GetPrivateDatabasePlainArgs args) {
        return getPrivateDatabasePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting database.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseArgs;
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
     *         final var database = HostingFunctions.getPrivateDatabase(GetPrivateDatabaseArgs.builder()
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
    public static Output<GetPrivateDatabaseResult> getPrivateDatabase(GetPrivateDatabaseArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Hosting/getPrivateDatabase:getPrivateDatabase", TypeShape.of(GetPrivateDatabaseResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an hosting database.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseArgs;
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
     *         final var database = HostingFunctions.getPrivateDatabase(GetPrivateDatabaseArgs.builder()
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
    public static CompletableFuture<GetPrivateDatabaseResult> getPrivateDatabasePlain(GetPrivateDatabasePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Hosting/getPrivateDatabase:getPrivateDatabase", TypeShape.of(GetPrivateDatabaseResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase whitelist.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseAllowlistArgs;
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
     *         final var whitelist = HostingFunctions.getPrivateDatabaseAllowlist(GetPrivateDatabaseAllowlistArgs.builder()
     *             .serviceName("XXXXXX")
     *             .ip("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPrivateDatabaseAllowlistResult> getPrivateDatabaseAllowlist(GetPrivateDatabaseAllowlistArgs args) {
        return getPrivateDatabaseAllowlist(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase whitelist.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseAllowlistArgs;
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
     *         final var whitelist = HostingFunctions.getPrivateDatabaseAllowlist(GetPrivateDatabaseAllowlistArgs.builder()
     *             .serviceName("XXXXXX")
     *             .ip("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPrivateDatabaseAllowlistResult> getPrivateDatabaseAllowlistPlain(GetPrivateDatabaseAllowlistPlainArgs args) {
        return getPrivateDatabaseAllowlistPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase whitelist.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseAllowlistArgs;
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
     *         final var whitelist = HostingFunctions.getPrivateDatabaseAllowlist(GetPrivateDatabaseAllowlistArgs.builder()
     *             .serviceName("XXXXXX")
     *             .ip("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPrivateDatabaseAllowlistResult> getPrivateDatabaseAllowlist(GetPrivateDatabaseAllowlistArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Hosting/getPrivateDatabaseAllowlist:getPrivateDatabaseAllowlist", TypeShape.of(GetPrivateDatabaseAllowlistResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase whitelist.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseAllowlistArgs;
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
     *         final var whitelist = HostingFunctions.getPrivateDatabaseAllowlist(GetPrivateDatabaseAllowlistArgs.builder()
     *             .serviceName("XXXXXX")
     *             .ip("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPrivateDatabaseAllowlistResult> getPrivateDatabaseAllowlistPlain(GetPrivateDatabaseAllowlistPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Hosting/getPrivateDatabaseAllowlist:getPrivateDatabaseAllowlist", TypeShape.of(GetPrivateDatabaseAllowlistResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseDbArgs;
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
     *         final var myDatabase = HostingFunctions.getPrivateDatabaseDb(GetPrivateDatabaseDbArgs.builder()
     *             .serviceName("XXXXXX")
     *             .databaseName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPrivateDatabaseDbResult> getPrivateDatabaseDb(GetPrivateDatabaseDbArgs args) {
        return getPrivateDatabaseDb(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseDbArgs;
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
     *         final var myDatabase = HostingFunctions.getPrivateDatabaseDb(GetPrivateDatabaseDbArgs.builder()
     *             .serviceName("XXXXXX")
     *             .databaseName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPrivateDatabaseDbResult> getPrivateDatabaseDbPlain(GetPrivateDatabaseDbPlainArgs args) {
        return getPrivateDatabaseDbPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseDbArgs;
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
     *         final var myDatabase = HostingFunctions.getPrivateDatabaseDb(GetPrivateDatabaseDbArgs.builder()
     *             .serviceName("XXXXXX")
     *             .databaseName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPrivateDatabaseDbResult> getPrivateDatabaseDb(GetPrivateDatabaseDbArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Hosting/getPrivateDatabaseDb:getPrivateDatabaseDb", TypeShape.of(GetPrivateDatabaseDbResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseDbArgs;
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
     *         final var myDatabase = HostingFunctions.getPrivateDatabaseDb(GetPrivateDatabaseDbArgs.builder()
     *             .serviceName("XXXXXX")
     *             .databaseName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPrivateDatabaseDbResult> getPrivateDatabaseDbPlain(GetPrivateDatabaseDbPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Hosting/getPrivateDatabaseDb:getPrivateDatabaseDb", TypeShape.of(GetPrivateDatabaseDbResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase user.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserArgs;
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
     *         final var user = HostingFunctions.getPrivateDatabaseUser(GetPrivateDatabaseUserArgs.builder()
     *             .serviceName("XXXXXX")
     *             .userName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPrivateDatabaseUserResult> getPrivateDatabaseUser(GetPrivateDatabaseUserArgs args) {
        return getPrivateDatabaseUser(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase user.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserArgs;
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
     *         final var user = HostingFunctions.getPrivateDatabaseUser(GetPrivateDatabaseUserArgs.builder()
     *             .serviceName("XXXXXX")
     *             .userName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPrivateDatabaseUserResult> getPrivateDatabaseUserPlain(GetPrivateDatabaseUserPlainArgs args) {
        return getPrivateDatabaseUserPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase user.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserArgs;
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
     *         final var user = HostingFunctions.getPrivateDatabaseUser(GetPrivateDatabaseUserArgs.builder()
     *             .serviceName("XXXXXX")
     *             .userName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPrivateDatabaseUserResult> getPrivateDatabaseUser(GetPrivateDatabaseUserArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Hosting/getPrivateDatabaseUser:getPrivateDatabaseUser", TypeShape.of(GetPrivateDatabaseUserResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase user.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserArgs;
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
     *         final var user = HostingFunctions.getPrivateDatabaseUser(GetPrivateDatabaseUserArgs.builder()
     *             .serviceName("XXXXXX")
     *             .userName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPrivateDatabaseUserResult> getPrivateDatabaseUserPlain(GetPrivateDatabaseUserPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Hosting/getPrivateDatabaseUser:getPrivateDatabaseUser", TypeShape.of(GetPrivateDatabaseUserResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase user grant.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserGrantArgs;
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
     *         final var userGrant = HostingFunctions.getPrivateDatabaseUserGrant(GetPrivateDatabaseUserGrantArgs.builder()
     *             .serviceName("XXXXXX")
     *             .databaseName("XXXXXX")
     *             .userName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPrivateDatabaseUserGrantResult> getPrivateDatabaseUserGrant(GetPrivateDatabaseUserGrantArgs args) {
        return getPrivateDatabaseUserGrant(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase user grant.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserGrantArgs;
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
     *         final var userGrant = HostingFunctions.getPrivateDatabaseUserGrant(GetPrivateDatabaseUserGrantArgs.builder()
     *             .serviceName("XXXXXX")
     *             .databaseName("XXXXXX")
     *             .userName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPrivateDatabaseUserGrantResult> getPrivateDatabaseUserGrantPlain(GetPrivateDatabaseUserGrantPlainArgs args) {
        return getPrivateDatabaseUserGrantPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase user grant.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserGrantArgs;
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
     *         final var userGrant = HostingFunctions.getPrivateDatabaseUserGrant(GetPrivateDatabaseUserGrantArgs.builder()
     *             .serviceName("XXXXXX")
     *             .databaseName("XXXXXX")
     *             .userName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPrivateDatabaseUserGrantResult> getPrivateDatabaseUserGrant(GetPrivateDatabaseUserGrantArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Hosting/getPrivateDatabaseUserGrant:getPrivateDatabaseUserGrant", TypeShape.of(GetPrivateDatabaseUserGrantResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about an hosting privatedatabase user grant.
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
     * import com.pulumi.ovh.Hosting.HostingFunctions;
     * import com.pulumi.ovh.Hosting.inputs.GetPrivateDatabaseUserGrantArgs;
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
     *         final var userGrant = HostingFunctions.getPrivateDatabaseUserGrant(GetPrivateDatabaseUserGrantArgs.builder()
     *             .serviceName("XXXXXX")
     *             .databaseName("XXXXXX")
     *             .userName("XXXXXX")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPrivateDatabaseUserGrantResult> getPrivateDatabaseUserGrantPlain(GetPrivateDatabaseUserGrantPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Hosting/getPrivateDatabaseUserGrant:getPrivateDatabaseUserGrant", TypeShape.of(GetPrivateDatabaseUserGrantResult.class), args, Utilities.withVersion(options));
    }
}
