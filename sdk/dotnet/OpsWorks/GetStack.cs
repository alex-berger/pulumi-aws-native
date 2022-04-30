// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks
{
    public static class GetStack
    {
        /// <summary>
        /// Resource Type definition for AWS::OpsWorks::Stack
        /// </summary>
        public static Task<GetStackResult> InvokeAsync(GetStackArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStackResult>("aws-native:opsworks:getStack", args ?? new GetStackArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::OpsWorks::Stack
        /// </summary>
        public static Output<GetStackResult> Invoke(GetStackInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStackResult>("aws-native:opsworks:getStack", args ?? new GetStackInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStackArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetStackArgs()
        {
        }
    }

    public sealed class GetStackInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetStackInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStackResult
    {
        public readonly string? AgentVersion;
        public readonly object? Attributes;
        public readonly Outputs.StackChefConfiguration? ChefConfiguration;
        public readonly Outputs.StackConfigurationManager? ConfigurationManager;
        public readonly Outputs.StackSource? CustomCookbooksSource;
        public readonly object? CustomJson;
        public readonly string? DefaultAvailabilityZone;
        public readonly string? DefaultInstanceProfileArn;
        public readonly string? DefaultOs;
        public readonly string? DefaultRootDeviceType;
        public readonly string? DefaultSshKeyName;
        public readonly string? DefaultSubnetId;
        public readonly string? EcsClusterArn;
        public readonly ImmutableArray<Outputs.StackElasticIp> ElasticIps;
        public readonly string? HostnameTheme;
        public readonly string? Id;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.StackRdsDbInstance> RdsDbInstances;
        public readonly ImmutableArray<Outputs.StackTag> Tags;
        public readonly bool? UseCustomCookbooks;
        public readonly bool? UseOpsworksSecurityGroups;

        [OutputConstructor]
        private GetStackResult(
            string? agentVersion,

            object? attributes,

            Outputs.StackChefConfiguration? chefConfiguration,

            Outputs.StackConfigurationManager? configurationManager,

            Outputs.StackSource? customCookbooksSource,

            object? customJson,

            string? defaultAvailabilityZone,

            string? defaultInstanceProfileArn,

            string? defaultOs,

            string? defaultRootDeviceType,

            string? defaultSshKeyName,

            string? defaultSubnetId,

            string? ecsClusterArn,

            ImmutableArray<Outputs.StackElasticIp> elasticIps,

            string? hostnameTheme,

            string? id,

            string? name,

            ImmutableArray<Outputs.StackRdsDbInstance> rdsDbInstances,

            ImmutableArray<Outputs.StackTag> tags,

            bool? useCustomCookbooks,

            bool? useOpsworksSecurityGroups)
        {
            AgentVersion = agentVersion;
            Attributes = attributes;
            ChefConfiguration = chefConfiguration;
            ConfigurationManager = configurationManager;
            CustomCookbooksSource = customCookbooksSource;
            CustomJson = customJson;
            DefaultAvailabilityZone = defaultAvailabilityZone;
            DefaultInstanceProfileArn = defaultInstanceProfileArn;
            DefaultOs = defaultOs;
            DefaultRootDeviceType = defaultRootDeviceType;
            DefaultSshKeyName = defaultSshKeyName;
            DefaultSubnetId = defaultSubnetId;
            EcsClusterArn = ecsClusterArn;
            ElasticIps = elasticIps;
            HostnameTheme = hostnameTheme;
            Id = id;
            Name = name;
            RdsDbInstances = rdsDbInstances;
            Tags = tags;
            UseCustomCookbooks = useCustomCookbooks;
            UseOpsworksSecurityGroups = useOpsworksSecurityGroups;
        }
    }
}
