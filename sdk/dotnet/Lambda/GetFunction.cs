// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda
{
    public static class GetFunction
    {
        /// <summary>
        /// Resource Type definition for AWS::Lambda::Function
        /// </summary>
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("aws-native:lambda:getFunction", args ?? new GetFunctionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Lambda::Function
        /// </summary>
        public static Output<GetFunctionResult> Invoke(GetFunctionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFunctionResult>("aws-native:lambda:getFunction", args ?? new GetFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.
        /// </summary>
        [Input("functionName", required: true)]
        public string FunctionName { get; set; } = null!;

        public GetFunctionArgs()
        {
        }
    }

    public sealed class GetFunctionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        public GetFunctionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionResult
    {
        public readonly ImmutableArray<Pulumi.AwsNative.Lambda.FunctionArchitecturesItem> Architectures;
        /// <summary>
        /// Unique identifier for function resources
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// A unique Arn for CodeSigningConfig resource
        /// </summary>
        public readonly string? CodeSigningConfigArn;
        /// <summary>
        /// A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.
        /// </summary>
        public readonly Outputs.FunctionDeadLetterConfig? DeadLetterConfig;
        /// <summary>
        /// A description of the function.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Environment variables that are accessible from function code during execution.
        /// </summary>
        public readonly Outputs.FunctionEnvironment? Environment;
        /// <summary>
        /// A function's ephemeral storage settings.
        /// </summary>
        public readonly Outputs.FunctionEphemeralStorage? EphemeralStorage;
        /// <summary>
        /// Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function.
        /// </summary>
        public readonly ImmutableArray<Outputs.FunctionFileSystemConfig> FileSystemConfigs;
        /// <summary>
        /// The name of the method within your code that Lambda calls to execute your function. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime
        /// </summary>
        public readonly string? Handler;
        /// <summary>
        /// ImageConfig
        /// </summary>
        public readonly Outputs.FunctionImageConfig? ImageConfig;
        /// <summary>
        /// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
        /// </summary>
        public readonly string? KmsKeyArn;
        /// <summary>
        /// A list of function layers to add to the function's execution environment. Specify each layer by its ARN, including the version.
        /// </summary>
        public readonly ImmutableArray<string> Layers;
        /// <summary>
        /// The amount of memory that your function has access to. Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.
        /// </summary>
        public readonly int? MemorySize;
        /// <summary>
        /// PackageType.
        /// </summary>
        public readonly Pulumi.AwsNative.Lambda.FunctionPackageType? PackageType;
        /// <summary>
        /// The number of simultaneous executions to reserve for the function.
        /// </summary>
        public readonly int? ReservedConcurrentExecutions;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the function's execution role.
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// The identifier of the function's runtime.
        /// </summary>
        public readonly string? Runtime;
        /// <summary>
        /// A list of tags to apply to the function.
        /// </summary>
        public readonly ImmutableArray<Outputs.FunctionTag> Tags;
        /// <summary>
        /// The amount of time that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// Set Mode to Active to sample and trace a subset of incoming requests with AWS X-Ray.
        /// </summary>
        public readonly Outputs.FunctionTracingConfig? TracingConfig;
        /// <summary>
        /// For network connectivity to AWS resources in a VPC, specify a list of security groups and subnets in the VPC.
        /// </summary>
        public readonly Outputs.FunctionVpcConfig? VpcConfig;

        [OutputConstructor]
        private GetFunctionResult(
            ImmutableArray<Pulumi.AwsNative.Lambda.FunctionArchitecturesItem> architectures,

            string? arn,

            string? codeSigningConfigArn,

            Outputs.FunctionDeadLetterConfig? deadLetterConfig,

            string? description,

            Outputs.FunctionEnvironment? environment,

            Outputs.FunctionEphemeralStorage? ephemeralStorage,

            ImmutableArray<Outputs.FunctionFileSystemConfig> fileSystemConfigs,

            string? handler,

            Outputs.FunctionImageConfig? imageConfig,

            string? kmsKeyArn,

            ImmutableArray<string> layers,

            int? memorySize,

            Pulumi.AwsNative.Lambda.FunctionPackageType? packageType,

            int? reservedConcurrentExecutions,

            string? role,

            string? runtime,

            ImmutableArray<Outputs.FunctionTag> tags,

            int? timeout,

            Outputs.FunctionTracingConfig? tracingConfig,

            Outputs.FunctionVpcConfig? vpcConfig)
        {
            Architectures = architectures;
            Arn = arn;
            CodeSigningConfigArn = codeSigningConfigArn;
            DeadLetterConfig = deadLetterConfig;
            Description = description;
            Environment = environment;
            EphemeralStorage = ephemeralStorage;
            FileSystemConfigs = fileSystemConfigs;
            Handler = handler;
            ImageConfig = imageConfig;
            KmsKeyArn = kmsKeyArn;
            Layers = layers;
            MemorySize = memorySize;
            PackageType = packageType;
            ReservedConcurrentExecutions = reservedConcurrentExecutions;
            Role = role;
            Runtime = runtime;
            Tags = tags;
            Timeout = timeout;
            TracingConfig = tracingConfig;
            VpcConfig = vpcConfig;
        }
    }
}
