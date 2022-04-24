// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetNotebookInstanceLifecycleConfig
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::NotebookInstanceLifecycleConfig
        /// </summary>
        public static Task<GetNotebookInstanceLifecycleConfigResult> InvokeAsync(GetNotebookInstanceLifecycleConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNotebookInstanceLifecycleConfigResult>("aws-native:sagemaker:getNotebookInstanceLifecycleConfig", args ?? new GetNotebookInstanceLifecycleConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::NotebookInstanceLifecycleConfig
        /// </summary>
        public static Output<GetNotebookInstanceLifecycleConfigResult> Invoke(GetNotebookInstanceLifecycleConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNotebookInstanceLifecycleConfigResult>("aws-native:sagemaker:getNotebookInstanceLifecycleConfig", args ?? new GetNotebookInstanceLifecycleConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotebookInstanceLifecycleConfigArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetNotebookInstanceLifecycleConfigArgs()
        {
        }
    }

    public sealed class GetNotebookInstanceLifecycleConfigInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetNotebookInstanceLifecycleConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNotebookInstanceLifecycleConfigResult
    {
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook> OnCreate;
        public readonly ImmutableArray<Outputs.NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook> OnStart;

        [OutputConstructor]
        private GetNotebookInstanceLifecycleConfigResult(
            string? id,

            ImmutableArray<Outputs.NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook> onCreate,

            ImmutableArray<Outputs.NotebookInstanceLifecycleConfigNotebookInstanceLifecycleHook> onStart)
        {
            Id = id;
            OnCreate = onCreate;
            OnStart = onStart;
        }
    }
}
