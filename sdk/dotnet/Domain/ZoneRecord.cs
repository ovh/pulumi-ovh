// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Domain
{
    /// <summary>
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
    ///     // Add a record to a sub-domain
    ///     var test = new Ovh.Domain.ZoneRecord("test", new()
    ///     {
    ///         Zone = "testdemo.ovh",
    ///         Subdomain = "test",
    ///         Fieldtype = "A",
    ///         Ttl = 3600,
    ///         Target = "0.0.0.0",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud domain zone record can be imported using the `id`, which can be retrieved by using [OVH API portal](https://api.ovh.com/console/#/domain/zone/%7BzoneName%7D/record~GET), and the `zone`, separated by "." E.g.,
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:Domain/zoneRecord:ZoneRecord test id.zone
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Domain/zoneRecord:ZoneRecord")]
    public partial class ZoneRecord : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of the record
        /// </summary>
        [Output("fieldtype")]
        public Output<string> Fieldtype { get; private set; } = null!;

        /// <summary>
        /// The name of the record. It can be an empty string.
        /// </summary>
        [Output("subdomain")]
        public Output<string?> Subdomain { get; private set; } = null!;

        /// <summary>
        /// The value of the record
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;

        /// <summary>
        /// The TTL of the record, it shall be &gt;= to 60.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;

        /// <summary>
        /// The domain to add the record to
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a ZoneRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZoneRecord(string name, ZoneRecordArgs args, CustomResourceOptions? options = null)
            : base("ovh:Domain/zoneRecord:ZoneRecord", name, args ?? new ZoneRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ZoneRecord(string name, Input<string> id, ZoneRecordState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Domain/zoneRecord:ZoneRecord", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ZoneRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZoneRecord Get(string name, Input<string> id, ZoneRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new ZoneRecord(name, id, state, options);
        }
    }

    public sealed class ZoneRecordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the record
        /// </summary>
        [Input("fieldtype", required: true)]
        public Input<string> Fieldtype { get; set; } = null!;

        /// <summary>
        /// The name of the record. It can be an empty string.
        /// </summary>
        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

        /// <summary>
        /// The value of the record
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// The TTL of the record, it shall be &gt;= to 60.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The domain to add the record to
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ZoneRecordArgs()
        {
        }
        public static new ZoneRecordArgs Empty => new ZoneRecordArgs();
    }

    public sealed class ZoneRecordState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the record
        /// </summary>
        [Input("fieldtype")]
        public Input<string>? Fieldtype { get; set; }

        /// <summary>
        /// The name of the record. It can be an empty string.
        /// </summary>
        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

        /// <summary>
        /// The value of the record
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// The TTL of the record, it shall be &gt;= to 60.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The domain to add the record to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ZoneRecordState()
        {
        }
        public static new ZoneRecordState Empty => new ZoneRecordState();
    }
}
