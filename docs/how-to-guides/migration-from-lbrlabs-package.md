---
title: "Migration from LbrLabs OVH Package"
h1: "Migrating from the LbrLabs OVH Package"
meta_desc: Practitioner level instructions for migrating from the community LbrLabs OVH package to this official OVH package.
layout: package
---

The types in each of the Pulumi OVH SDKs have been kept as compatible as possible to the LbrLabs community package.
This should make the migration from the community package to this official OVH package almost a drop in replacement.

{{< chooser language "typescript,python,go,csharp" >}}

{{% choosable language typescript %}}

Replace the NPM package in your project first:

```sh
$ npm uninstall --save @lbrlabs/pulumi-ovh
$ npm install --save @ovh-devrelteam/pulumi-ovh
```

After switching the package, update the import statements in your Pulumi infrastructure code:

```diff
- import * as ovh from "@lbrlabs/pulumi-ovh";
+ import * as ovh from "@ovh-devrelteam/pulumi-ovh";
```

{{% /choosable %}}

{{% choosable language python %}}

Replace the Python package in your project first:

```sh
$ pip uninstall lbrlabs-pulumi-ovh
$ pip install pulumi-ovh
```

After switching the package, update the import statements in your Pulumi infrastructure code:

```diff
-  import lbrlabs_pulumi_ovh as ovh
+  import pulumi_ovh as ovh
```

{{% /choosable %}}

{{% choosable language go %}}

Replace the Go package in your project first:

```sh
$ go get github.com/ovh/pulumi-ovh/sdk/go/...
```

After switching the package, update the import statements in your Pulumi infrastructure code:

```diff
- "github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProject"
+ "github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProject"
```

{{% /choosable %}}

{{% choosable language csharp %}}

Replace the Nuget package in your project first:

```sh
$ dotnet remove package Lbrlabs.PulumiPackage.Ovh
$ dotnet add package Pulumi.Ovh
```

After switching the package, update the import statements in your Pulumi infrastructure code:

```diff
- using Ovh = Lbrlabs.PulumiPackage.Ovh;
+ using Ovh = Pulumi.Ovh;
```

{{% /choosable %}}

{{< /chooser >}}

Your project has now been updated to use the new SDK. When running `pulumi preview` with this new setup,
you should see no changes to your OVH resources reported.

If you do encounter something unexpected, please file an issue in [the Github repository](https://github.com/ovh/pulumi-ovh/issues).
