// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetStorageObject
    {
        /// <summary>
        /// Get information about an object in a S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
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
        ///     var @object = Ovh.CloudProject.GetStorageObject.Invoke(new()
        ///     {
        ///         ServiceName = "&lt;public cloud project ID&gt;",
        ///         RegionName = "GRA",
        ///         Name = "&lt;bucket name&gt;",
        ///         Key = "&lt;object name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetStorageObjectResult> InvokeAsync(GetStorageObjectArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStorageObjectResult>("ovh:CloudProject/getStorageObject:getStorageObject", args ?? new GetStorageObjectArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about an object in a S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
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
        ///     var @object = Ovh.CloudProject.GetStorageObject.Invoke(new()
        ///     {
        ///         ServiceName = "&lt;public cloud project ID&gt;",
        ///         RegionName = "GRA",
        ///         Name = "&lt;bucket name&gt;",
        ///         Key = "&lt;object name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetStorageObjectResult> Invoke(GetStorageObjectInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageObjectResult>("ovh:CloudProject/getStorageObject:getStorageObject", args ?? new GetStorageObjectInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about an object in a S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
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
        ///     var @object = Ovh.CloudProject.GetStorageObject.Invoke(new()
        ///     {
        ///         ServiceName = "&lt;public cloud project ID&gt;",
        ///         RegionName = "GRA",
        ///         Name = "&lt;bucket name&gt;",
        ///         Key = "&lt;object name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetStorageObjectResult> Invoke(GetStorageObjectInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageObjectResult>("ovh:CloudProject/getStorageObject:getStorageObject", args ?? new GetStorageObjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStorageObjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Key
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        /// <summary>
        /// Name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

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

        public GetStorageObjectArgs()
        {
        }
        public static new GetStorageObjectArgs Empty => new GetStorageObjectArgs();
    }

    public sealed class GetStorageObjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Key
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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

        public GetStorageObjectInvokeArgs()
        {
        }
        public static new GetStorageObjectInvokeArgs Empty => new GetStorageObjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetStorageObjectResult
    {
        /// <summary>
        /// ETag
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether this object is a delete marker
        /// </summary>
        public readonly bool IsDeleteMarker;
        /// <summary>
        /// Whether this is the latest version of the object
        /// </summary>
        public readonly bool IsLatest;
        /// <summary>
        /// Key
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Last modification date
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Region name
        /// </summary>
        public readonly string RegionName;
        /// <summary>
        /// Service name
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Size (bytes)
        /// </summary>
        public readonly double Size;
        /// <summary>
        /// Storage class
        /// </summary>
        public readonly string StorageClass;
        /// <summary>
        /// Version ID of the object
        /// </summary>
        public readonly string VersionId;

        [OutputConstructor]
        private GetStorageObjectResult(
            string etag,

            string id,

            bool isDeleteMarker,

            bool isLatest,

            string key,

            string lastModified,

            string name,

            string regionName,

            string serviceName,

            double size,

            string storageClass,

            string versionId)
        {
            Etag = etag;
            Id = id;
            IsDeleteMarker = isDeleteMarker;
            IsLatest = isLatest;
            Key = key;
            LastModified = lastModified;
            Name = name;
            RegionName = regionName;
            ServiceName = serviceName;
            Size = size;
            StorageClass = storageClass;
            VersionId = versionId;
        }
    }
}
