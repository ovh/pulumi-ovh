// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetStorage
    {
        /// <summary>
        /// Get S3™* compatible storage container.
        /// \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var storage = Ovh.CloudProject.GetStorage.Invoke(new()
        ///     {
        ///         Name = "my-storage-name",
        ///         RegionName = "GRA",
        ///         ServiceName = "&lt;public cloud project ID&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetStorageResult> InvokeAsync(GetStorageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStorageResult>("ovh:CloudProject/getStorage:getStorage", args ?? new GetStorageArgs(), options.WithDefaults());

        /// <summary>
        /// Get S3™* compatible storage container.
        /// \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var storage = Ovh.CloudProject.GetStorage.Invoke(new()
        ///     {
        ///         Name = "my-storage-name",
        ///         RegionName = "GRA",
        ///         ServiceName = "&lt;public cloud project ID&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetStorageResult> Invoke(GetStorageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageResult>("ovh:CloudProject/getStorage:getStorage", args ?? new GetStorageInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get S3™* compatible storage container.
        /// \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var storage = Ovh.CloudProject.GetStorage.Invoke(new()
        ///     {
        ///         Name = "my-storage-name",
        ///         RegionName = "GRA",
        ///         ServiceName = "&lt;public cloud project ID&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetStorageResult> Invoke(GetStorageInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageResult>("ovh:CloudProject/getStorage:getStorage", args ?? new GetStorageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStorageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Limit the number of objects returned (1000 maximum, defaults to 1000)
        /// </summary>
        [Input("limit")]
        public double? Limit { get; set; }

        /// <summary>
        /// Key to start with when listing objects
        /// </summary>
        [Input("marker")]
        public string? Marker { get; set; }

        /// <summary>
        /// Name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// List objects whose key begins with this prefix
        /// </summary>
        [Input("prefix")]
        public string? Prefix { get; set; }

        /// <summary>
        /// Region name
        /// </summary>
        [Input("regionName", required: true)]
        public string RegionName { get; set; } = null!;

        /// <summary>
        /// Service name
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetStorageArgs()
        {
        }
        public static new GetStorageArgs Empty => new GetStorageArgs();
    }

    public sealed class GetStorageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Limit the number of objects returned (1000 maximum, defaults to 1000)
        /// </summary>
        [Input("limit")]
        public Input<double>? Limit { get; set; }

        /// <summary>
        /// Key to start with when listing objects
        /// </summary>
        [Input("marker")]
        public Input<string>? Marker { get; set; }

        /// <summary>
        /// Name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// List objects whose key begins with this prefix
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Region name
        /// </summary>
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        /// <summary>
        /// Service name
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetStorageInvokeArgs()
        {
        }
        public static new GetStorageInvokeArgs Empty => new GetStorageInvokeArgs();
    }


    [OutputType]
    public sealed class GetStorageResult
    {
        /// <summary>
        /// The date and timestamp when the resource was created
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Encryption configuration
        /// </summary>
        public readonly Outputs.GetStorageEncryptionResult Encryption;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Limit the number of objects returned (1000 maximum, defaults to 1000)
        /// </summary>
        public readonly double Limit;
        /// <summary>
        /// Key to start with when listing objects
        /// </summary>
        public readonly string Marker;
        /// <summary>
        /// Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Container objects
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStorageObjectResult> Objects;
        /// <summary>
        /// Container total objects count
        /// </summary>
        public readonly double ObjectsCount;
        /// <summary>
        /// Container total objects size (bytes)
        /// </summary>
        public readonly double ObjectsSize;
        /// <summary>
        /// Container owner user ID
        /// </summary>
        public readonly double OwnerId;
        /// <summary>
        /// List objects whose key begins with this prefix
        /// </summary>
        public readonly string Prefix;
        /// <summary>
        /// Container region
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Region name
        /// </summary>
        public readonly string RegionName;
        /// <summary>
        /// Replication configuration
        /// </summary>
        public readonly Outputs.GetStorageReplicationResult Replication;
        /// <summary>
        /// Service name
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Container tags
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Versioning configuration
        /// </summary>
        public readonly Outputs.GetStorageVersioningResult Versioning;
        /// <summary>
        /// Container virtual host
        /// </summary>
        public readonly string VirtualHost;

        [OutputConstructor]
        private GetStorageResult(
            string createdAt,

            Outputs.GetStorageEncryptionResult encryption,

            string id,

            double limit,

            string marker,

            string name,

            ImmutableArray<Outputs.GetStorageObjectResult> objects,

            double objectsCount,

            double objectsSize,

            double ownerId,

            string prefix,

            string region,

            string regionName,

            Outputs.GetStorageReplicationResult replication,

            string serviceName,

            ImmutableDictionary<string, string> tags,

            Outputs.GetStorageVersioningResult versioning,

            string virtualHost)
        {
            CreatedAt = createdAt;
            Encryption = encryption;
            Id = id;
            Limit = limit;
            Marker = marker;
            Name = name;
            Objects = objects;
            ObjectsCount = objectsCount;
            ObjectsSize = objectsSize;
            OwnerId = ownerId;
            Prefix = prefix;
            Region = region;
            RegionName = regionName;
            Replication = replication;
            ServiceName = serviceName;
            Tags = tags;
            Versioning = versioning;
            VirtualHost = virtualHost;
        }
    }
}
