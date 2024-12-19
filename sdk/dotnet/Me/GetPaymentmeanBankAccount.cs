// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me
{
    public static class GetPaymentmeanBankAccount
    {
        /// <summary>
        /// Use this data source to retrieve information about a bank account
        /// payment mean associated with an OVHcloud account.
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
        ///     var ba = Ovh.Me.GetPaymentmeanBankAccount.Invoke(new()
        ///     {
        ///         UseDefault = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetPaymentmeanBankAccountResult> InvokeAsync(GetPaymentmeanBankAccountArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPaymentmeanBankAccountResult>("ovh:Me/getPaymentmeanBankAccount:getPaymentmeanBankAccount", args ?? new GetPaymentmeanBankAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a bank account
        /// payment mean associated with an OVHcloud account.
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
        ///     var ba = Ovh.Me.GetPaymentmeanBankAccount.Invoke(new()
        ///     {
        ///         UseDefault = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPaymentmeanBankAccountResult> Invoke(GetPaymentmeanBankAccountInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPaymentmeanBankAccountResult>("ovh:Me/getPaymentmeanBankAccount:getPaymentmeanBankAccount", args ?? new GetPaymentmeanBankAccountInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a bank account
        /// payment mean associated with an OVHcloud account.
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
        ///     var ba = Ovh.Me.GetPaymentmeanBankAccount.Invoke(new()
        ///     {
        ///         UseDefault = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPaymentmeanBankAccountResult> Invoke(GetPaymentmeanBankAccountInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPaymentmeanBankAccountResult>("ovh:Me/getPaymentmeanBankAccount:getPaymentmeanBankAccount", args ?? new GetPaymentmeanBankAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPaymentmeanBankAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// a regexp used to filter bank accounts 
        /// on their `description` attributes.
        /// </summary>
        [Input("descriptionRegexp")]
        public string? DescriptionRegexp { get; set; }

        /// <summary>
        /// Filter bank accounts on their `state` attribute.
        /// Can be "blockedForIncidents", "valid", "pendingValidation"
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        /// <summary>
        /// Retrieve bank account marked as default payment mean.
        /// </summary>
        [Input("useDefault")]
        public bool? UseDefault { get; set; }

        /// <summary>
        /// Retrieve oldest bank account.
        /// project.
        /// </summary>
        [Input("useOldest")]
        public bool? UseOldest { get; set; }

        public GetPaymentmeanBankAccountArgs()
        {
        }
        public static new GetPaymentmeanBankAccountArgs Empty => new GetPaymentmeanBankAccountArgs();
    }

    public sealed class GetPaymentmeanBankAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// a regexp used to filter bank accounts 
        /// on their `description` attributes.
        /// </summary>
        [Input("descriptionRegexp")]
        public Input<string>? DescriptionRegexp { get; set; }

        /// <summary>
        /// Filter bank accounts on their `state` attribute.
        /// Can be "blockedForIncidents", "valid", "pendingValidation"
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Retrieve bank account marked as default payment mean.
        /// </summary>
        [Input("useDefault")]
        public Input<bool>? UseDefault { get; set; }

        /// <summary>
        /// Retrieve oldest bank account.
        /// project.
        /// </summary>
        [Input("useOldest")]
        public Input<bool>? UseOldest { get; set; }

        public GetPaymentmeanBankAccountInvokeArgs()
        {
        }
        public static new GetPaymentmeanBankAccountInvokeArgs Empty => new GetPaymentmeanBankAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetPaymentmeanBankAccountResult
    {
        /// <summary>
        /// a boolean which tells if the retrieved bank account
        /// is marked as the default payment mean
        /// </summary>
        public readonly bool Default;
        /// <summary>
        /// the description attribute of the bank account
        /// </summary>
        public readonly string Description;
        public readonly string? DescriptionRegexp;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string State;
        public readonly bool? UseDefault;
        public readonly bool? UseOldest;

        [OutputConstructor]
        private GetPaymentmeanBankAccountResult(
            bool @default,

            string description,

            string? descriptionRegexp,

            string id,

            string state,

            bool? useDefault,

            bool? useOldest)
        {
            Default = @default;
            Description = description;
            DescriptionRegexp = descriptionRegexp;
            Id = id;
            State = state;
            UseDefault = useDefault;
            UseOldest = useOldest;
        }
    }
}
