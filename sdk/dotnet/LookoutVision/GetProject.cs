// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutVision
{
    public static class GetProject
    {
        /// <summary>
        /// The AWS::LookoutVision::Project type creates an Amazon Lookout for Vision project. A project is a grouping of the resources needed to create and manage a Lookout for Vision model.
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("aws-native:lookoutvision:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::LookoutVision::Project type creates an Amazon Lookout for Vision project. A project is a grouping of the resources needed to create and manage a Lookout for Vision model.
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProjectResult>("aws-native:lookoutvision:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        [Input("projectName", required: true)]
        public string ProjectName { get; set; } = null!;

        public GetProjectArgs()
        {
        }
    }

    public sealed class GetProjectInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        public GetProjectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        public readonly string? Arn;

        [OutputConstructor]
        private GetProjectResult(string? arn)
        {
            Arn = arn;
        }
    }
}