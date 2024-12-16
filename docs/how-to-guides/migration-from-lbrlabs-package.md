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
$ npm install --save @ovhcloud/pulumi-ovh
```

After switching the package, update the import statements in your Pulumi infrastructure code:

```diff
- import * as ovh from "@lbrlabs/pulumi-ovh";
+ import * as ovh from "@ovhcloud/pulumi-ovh";
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

Some package has changed, and you'll have to update it as well, for example:

```diff
- ovh.DomainZoneRecord
+ ovh.domain.ZoneRecord
```

Or:

```diff
- ovh.CloudProjectUser
+ ovh.cloudproject.User
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

## Provider configuration

The LbrLabs OVH package marked the consumer key as a Pulumi secret. This was incorrect. This OVH package has it marked
as a regular configuration item, but marks the application secret correctly as a Pulumi secret.

If you have an existing setup where the provider configuration was passed using stack configuration, you have to
change the stack configuration. Make sure you have valid credentials at hand. Then run the following commands:

```sh
pulumi config rm ovh:consumerKey
pulumi config set ovh:consumerKey <your-consumer-key>
```

The consumer key is removed from the stack config as the old encrypted value, and re-added as a non-encrypted value.

```sh
pulumi config rm ovh:applicationSecret
pulumi config set ovh:applicationSecret --secret
value: <paste-your-application-secret>
```

The application secret was unencrypted before. This means that the value could have been save in Pulumi state unencrypted.
By removing it from config and re-adding it as a Pulumi secret config value, Pulumi can now track the value wherever it
is used and it will never be saved in clear text.

Because the application secret was processed unencrypted before by the LbrLabs OVH package,
it is adviced to rotate your application credentials and add the new set as an encrypted value.
In this case, it is possible that Pulumi reports a change to the state, without any clear change to a resource.
The clear text value is replaced by the encrypted value. Please accept this Pulumi update to increase
the security posture of your setup.
