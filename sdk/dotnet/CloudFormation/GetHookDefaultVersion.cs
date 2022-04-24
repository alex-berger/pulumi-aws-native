// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    public static class GetHookDefaultVersion
    {
        /// <summary>
        /// Set a version as default version for a hook in CloudFormation Registry.
        /// </summary>
        public static Task<GetHookDefaultVersionResult> InvokeAsync(GetHookDefaultVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHookDefaultVersionResult>("aws-native:cloudformation:getHookDefaultVersion", args ?? new GetHookDefaultVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Set a version as default version for a hook in CloudFormation Registry.
        /// </summary>
        public static Output<GetHookDefaultVersionResult> Invoke(GetHookDefaultVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetHookDefaultVersionResult>("aws-native:cloudformation:getHookDefaultVersion", args ?? new GetHookDefaultVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHookDefaultVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetHookDefaultVersionArgs()
        {
        }
    }

    public sealed class GetHookDefaultVersionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetHookDefaultVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHookDefaultVersionResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The name of the type being registered.
        /// 
        /// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
        /// </summary>
        public readonly string? TypeName;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the type version.
        /// </summary>
        public readonly string? TypeVersionArn;
        /// <summary>
        /// The ID of an existing version of the hook to set as the default.
        /// </summary>
        public readonly string? VersionId;

        [OutputConstructor]
        private GetHookDefaultVersionResult(
            string? arn,

            string? typeName,

            string? typeVersionArn,

            string? versionId)
        {
            Arn = arn;
            TypeName = typeName;
            TypeVersionArn = typeVersionArn;
            VersionId = versionId;
        }
    }
}
