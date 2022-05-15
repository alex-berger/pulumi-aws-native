// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS
{
    public static class GetTaskSet
    {
        /// <summary>
        /// Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.
        /// </summary>
        public static Task<GetTaskSetResult> InvokeAsync(GetTaskSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTaskSetResult>("aws-native:ecs:getTaskSet", args ?? new GetTaskSetArgs(), options.WithDefaults());

        /// <summary>
        /// Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.
        /// </summary>
        public static Output<GetTaskSetResult> Invoke(GetTaskSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTaskSetResult>("aws-native:ecs:getTaskSet", args ?? new GetTaskSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTaskSetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
        /// </summary>
        [Input("cluster", required: true)]
        public string Cluster { get; set; } = null!;

        /// <summary>
        /// The ID of the task set.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
        /// </summary>
        [Input("service", required: true)]
        public string Service { get; set; } = null!;

        public GetTaskSetArgs()
        {
        }
    }

    public sealed class GetTaskSetInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        /// <summary>
        /// The ID of the task set.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public GetTaskSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTaskSetResult
    {
        /// <summary>
        /// The ID of the task set.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// A floating-point percentage of the desired number of tasks to place and keep running in the task set.
        /// </summary>
        public readonly Outputs.TaskSetScale? Scale;

        [OutputConstructor]
        private GetTaskSetResult(
            string? id,

            Outputs.TaskSetScale? scale)
        {
            Id = id;
            Scale = scale;
        }
    }
}
