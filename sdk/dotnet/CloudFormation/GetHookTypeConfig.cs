// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    public static class GetHookTypeConfig
    {
        /// <summary>
        /// Specifies the configuration data for a registered hook in CloudFormation Registry.
        /// </summary>
        public static Task<GetHookTypeConfigResult> InvokeAsync(GetHookTypeConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHookTypeConfigResult>("aws-native:cloudformation:getHookTypeConfig", args ?? new GetHookTypeConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies the configuration data for a registered hook in CloudFormation Registry.
        /// </summary>
        public static Output<GetHookTypeConfigResult> Invoke(GetHookTypeConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetHookTypeConfigResult>("aws-native:cloudformation:getHookTypeConfig", args ?? new GetHookTypeConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHookTypeConfigArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the configuration data, in this account and region.
        /// </summary>
        [Input("configurationArn", required: true)]
        public string ConfigurationArn { get; set; } = null!;

        public GetHookTypeConfigArgs()
        {
        }
    }

    public sealed class GetHookTypeConfigInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the configuration data, in this account and region.
        /// </summary>
        [Input("configurationArn", required: true)]
        public Input<string> ConfigurationArn { get; set; } = null!;

        public GetHookTypeConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHookTypeConfigResult
    {
        /// <summary>
        /// The configuration data for the extension, in this account and region.
        /// </summary>
        public readonly string? Configuration;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the configuration data, in this account and region.
        /// </summary>
        public readonly string? ConfigurationArn;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the type without version number.
        /// </summary>
        public readonly string? TypeArn;
        /// <summary>
        /// The name of the type being registered.
        /// 
        /// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
        /// </summary>
        public readonly string? TypeName;

        [OutputConstructor]
        private GetHookTypeConfigResult(
            string? configuration,

            string? configurationArn,

            string? typeArn,

            string? typeName)
        {
            Configuration = configuration;
            ConfigurationArn = configurationArn;
            TypeArn = typeArn;
            TypeName = typeName;
        }
    }
}
