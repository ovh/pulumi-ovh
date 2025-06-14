// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetImage
    {
        /// <summary>
        /// Get information about an image in the given public cloud project.
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
        ///     var image = Ovh.CloudProject.GetImage.Invoke(new()
        ///     {
        ///         ServiceName = "&lt;public cloud project ID&gt;",
        ///         ImageId = "&lt;image ID&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("ovh:CloudProject/getImage:getImage", args ?? new GetImageArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about an image in the given public cloud project.
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
        ///     var image = Ovh.CloudProject.GetImage.Invoke(new()
        ///     {
        ///         ServiceName = "&lt;public cloud project ID&gt;",
        ///         ImageId = "&lt;image ID&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetImageResult> Invoke(GetImageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageResult>("ovh:CloudProject/getImage:getImage", args ?? new GetImageInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about an image in the given public cloud project.
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
        ///     var image = Ovh.CloudProject.GetImage.Invoke(new()
        ///     {
        ///         ServiceName = "&lt;public cloud project ID&gt;",
        ///         ImageId = "&lt;image ID&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetImageResult> Invoke(GetImageInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageResult>("ovh:CloudProject/getImage:getImage", args ?? new GetImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Image ID
        /// </summary>
        [Input("imageId", required: true)]
        public string ImageId { get; set; } = null!;

        /// <summary>
        /// Public cloud project ID
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetImageArgs()
        {
        }
        public static new GetImageArgs Empty => new GetImageArgs();
    }

    public sealed class GetImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Image ID
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        /// <summary>
        /// Public cloud project ID
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetImageInvokeArgs()
        {
        }
        public static new GetImageInvokeArgs Empty => new GetImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetImageResult
    {
        /// <summary>
        /// Image creation date
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// Image usable only for this type of flavor if not null
        /// </summary>
        public readonly string FlavorType;
        /// <summary>
        /// Image ID
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Image ID
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// Minimum disks required to use image
        /// </summary>
        public readonly double MinDisk;
        /// <summary>
        /// Minimum RAM required to use image
        /// </summary>
        public readonly double MinRam;
        /// <summary>
        /// Image name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Order plan code
        /// </summary>
        public readonly string PlanCode;
        /// <summary>
        /// Image region
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Public cloud project ID
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Image size (in GiB)
        /// </summary>
        public readonly double Size;
        /// <summary>
        /// Image status
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Tags about the image
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// Image type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// User to connect with
        /// </summary>
        public readonly string User;
        /// <summary>
        /// Image visibility
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetImageResult(
            string creationDate,

            string flavorType,

            string id,

            string imageId,

            double minDisk,

            double minRam,

            string name,

            string planCode,

            string region,

            string serviceName,

            double size,

            string status,

            ImmutableArray<string> tags,

            string type,

            string user,

            string visibility)
        {
            CreationDate = creationDate;
            FlavorType = flavorType;
            Id = id;
            ImageId = imageId;
            MinDisk = minDisk;
            MinRam = minRam;
            Name = name;
            PlanCode = planCode;
            Region = region;
            ServiceName = serviceName;
            Size = size;
            Status = status;
            Tags = tags;
            Type = type;
            User = user;
            Visibility = visibility;
        }
    }
}
