// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.VMware;

import com.ovhcloud.pulumi.ovh.Utilities;
import com.ovhcloud.pulumi.ovh.VMware.inputs.GetCloudDirectorBackupArgs;
import com.ovhcloud.pulumi.ovh.VMware.inputs.GetCloudDirectorBackupPlainArgs;
import com.ovhcloud.pulumi.ovh.VMware.inputs.GetCloudDirectorOrganizationArgs;
import com.ovhcloud.pulumi.ovh.VMware.inputs.GetCloudDirectorOrganizationPlainArgs;
import com.ovhcloud.pulumi.ovh.VMware.outputs.GetCloudDirectorBackupResult;
import com.ovhcloud.pulumi.ovh.VMware.outputs.GetCloudDirectorOrganizationResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class VMwareFunctions {
    /**
     * Get information about a VMware Cloud Director Backup service
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
     * import com.pulumi.ovh.VMware.VMwareFunctions;
     * import com.pulumi.ovh.VMware.inputs.GetCloudDirectorBackupArgs;
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
     *         final var backup = VMwareFunctions.getCloudDirectorBackup(GetCloudDirectorBackupArgs.builder()
     *             .backupId("<VCD backup ID>")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetCloudDirectorBackupResult> getCloudDirectorBackup(GetCloudDirectorBackupArgs args) {
        return getCloudDirectorBackup(args, InvokeOptions.Empty);
    }
    /**
     * Get information about a VMware Cloud Director Backup service
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
     * import com.pulumi.ovh.VMware.VMwareFunctions;
     * import com.pulumi.ovh.VMware.inputs.GetCloudDirectorBackupArgs;
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
     *         final var backup = VMwareFunctions.getCloudDirectorBackup(GetCloudDirectorBackupArgs.builder()
     *             .backupId("<VCD backup ID>")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetCloudDirectorBackupResult> getCloudDirectorBackupPlain(GetCloudDirectorBackupPlainArgs args) {
        return getCloudDirectorBackupPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get information about a VMware Cloud Director Backup service
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
     * import com.pulumi.ovh.VMware.VMwareFunctions;
     * import com.pulumi.ovh.VMware.inputs.GetCloudDirectorBackupArgs;
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
     *         final var backup = VMwareFunctions.getCloudDirectorBackup(GetCloudDirectorBackupArgs.builder()
     *             .backupId("<VCD backup ID>")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetCloudDirectorBackupResult> getCloudDirectorBackup(GetCloudDirectorBackupArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:VMware/getCloudDirectorBackup:getCloudDirectorBackup", TypeShape.of(GetCloudDirectorBackupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get information about a VMware Cloud Director Backup service
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
     * import com.pulumi.ovh.VMware.VMwareFunctions;
     * import com.pulumi.ovh.VMware.inputs.GetCloudDirectorBackupArgs;
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
     *         final var backup = VMwareFunctions.getCloudDirectorBackup(GetCloudDirectorBackupArgs.builder()
     *             .backupId("<VCD backup ID>")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetCloudDirectorBackupResult> getCloudDirectorBackupPlain(GetCloudDirectorBackupPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:VMware/getCloudDirectorBackup:getCloudDirectorBackup", TypeShape.of(GetCloudDirectorBackupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get VMware Cloud Director organization details
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
     * import com.pulumi.ovh.VMware.VMwareFunctions;
     * import com.pulumi.ovh.VMware.inputs.GetCloudDirectorOrganizationArgs;
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
     *         final var organization = VMwareFunctions.getCloudDirectorOrganization(GetCloudDirectorOrganizationArgs.builder()
     *             .organizationId("<VCD organization ID>")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetCloudDirectorOrganizationResult> getCloudDirectorOrganization(GetCloudDirectorOrganizationArgs args) {
        return getCloudDirectorOrganization(args, InvokeOptions.Empty);
    }
    /**
     * Get VMware Cloud Director organization details
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
     * import com.pulumi.ovh.VMware.VMwareFunctions;
     * import com.pulumi.ovh.VMware.inputs.GetCloudDirectorOrganizationArgs;
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
     *         final var organization = VMwareFunctions.getCloudDirectorOrganization(GetCloudDirectorOrganizationArgs.builder()
     *             .organizationId("<VCD organization ID>")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetCloudDirectorOrganizationResult> getCloudDirectorOrganizationPlain(GetCloudDirectorOrganizationPlainArgs args) {
        return getCloudDirectorOrganizationPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get VMware Cloud Director organization details
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
     * import com.pulumi.ovh.VMware.VMwareFunctions;
     * import com.pulumi.ovh.VMware.inputs.GetCloudDirectorOrganizationArgs;
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
     *         final var organization = VMwareFunctions.getCloudDirectorOrganization(GetCloudDirectorOrganizationArgs.builder()
     *             .organizationId("<VCD organization ID>")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetCloudDirectorOrganizationResult> getCloudDirectorOrganization(GetCloudDirectorOrganizationArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:VMware/getCloudDirectorOrganization:getCloudDirectorOrganization", TypeShape.of(GetCloudDirectorOrganizationResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get VMware Cloud Director organization details
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
     * import com.pulumi.ovh.VMware.VMwareFunctions;
     * import com.pulumi.ovh.VMware.inputs.GetCloudDirectorOrganizationArgs;
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
     *         final var organization = VMwareFunctions.getCloudDirectorOrganization(GetCloudDirectorOrganizationArgs.builder()
     *             .organizationId("<VCD organization ID>")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetCloudDirectorOrganizationResult> getCloudDirectorOrganizationPlain(GetCloudDirectorOrganizationPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:VMware/getCloudDirectorOrganization:getCloudDirectorOrganization", TypeShape.of(GetCloudDirectorOrganizationResult.class), args, Utilities.withVersion(options));
    }
}
