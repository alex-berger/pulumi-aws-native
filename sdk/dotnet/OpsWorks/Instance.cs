// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks
{
    /// <summary>
    /// Resource Type definition for AWS::OpsWorks::Instance
    /// </summary>
    [Obsolete(@"Instance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:opsworks:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        [Output("agentVersion")]
        public Output<string?> AgentVersion { get; private set; } = null!;

        [Output("amiId")]
        public Output<string?> AmiId { get; private set; } = null!;

        [Output("architecture")]
        public Output<string?> Architecture { get; private set; } = null!;

        [Output("autoScalingType")]
        public Output<string?> AutoScalingType { get; private set; } = null!;

        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        [Output("blockDeviceMappings")]
        public Output<ImmutableArray<Outputs.InstanceBlockDeviceMapping>> BlockDeviceMappings { get; private set; } = null!;

        [Output("ebsOptimized")]
        public Output<bool?> EbsOptimized { get; private set; } = null!;

        [Output("elasticIps")]
        public Output<ImmutableArray<string>> ElasticIps { get; private set; } = null!;

        [Output("hostname")]
        public Output<string?> Hostname { get; private set; } = null!;

        [Output("installUpdatesOnBoot")]
        public Output<bool?> InstallUpdatesOnBoot { get; private set; } = null!;

        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        [Output("layerIds")]
        public Output<ImmutableArray<string>> LayerIds { get; private set; } = null!;

        [Output("os")]
        public Output<string?> Os { get; private set; } = null!;

        [Output("privateDnsName")]
        public Output<string> PrivateDnsName { get; private set; } = null!;

        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        [Output("publicDnsName")]
        public Output<string> PublicDnsName { get; private set; } = null!;

        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        [Output("rootDeviceType")]
        public Output<string?> RootDeviceType { get; private set; } = null!;

        [Output("sshKeyName")]
        public Output<string?> SshKeyName { get; private set; } = null!;

        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        [Output("tenancy")]
        public Output<string?> Tenancy { get; private set; } = null!;

        [Output("timeBasedAutoScaling")]
        public Output<Outputs.InstanceTimeBasedAutoScaling?> TimeBasedAutoScaling { get; private set; } = null!;

        [Output("virtualizationType")]
        public Output<string?> VirtualizationType { get; private set; } = null!;

        [Output("volumes")]
        public Output<ImmutableArray<string>> Volumes { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:opsworks:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:opsworks:Instance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        [Input("agentVersion")]
        public Input<string>? AgentVersion { get; set; }

        [Input("amiId")]
        public Input<string>? AmiId { get; set; }

        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        [Input("autoScalingType")]
        public Input<string>? AutoScalingType { get; set; }

        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("blockDeviceMappings")]
        private InputList<Inputs.InstanceBlockDeviceMappingArgs>? _blockDeviceMappings;
        public InputList<Inputs.InstanceBlockDeviceMappingArgs> BlockDeviceMappings
        {
            get => _blockDeviceMappings ?? (_blockDeviceMappings = new InputList<Inputs.InstanceBlockDeviceMappingArgs>());
            set => _blockDeviceMappings = value;
        }

        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        [Input("elasticIps")]
        private InputList<string>? _elasticIps;
        public InputList<string> ElasticIps
        {
            get => _elasticIps ?? (_elasticIps = new InputList<string>());
            set => _elasticIps = value;
        }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("installUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("layerIds", required: true)]
        private InputList<string>? _layerIds;
        public InputList<string> LayerIds
        {
            get => _layerIds ?? (_layerIds = new InputList<string>());
            set => _layerIds = value;
        }

        [Input("os")]
        public Input<string>? Os { get; set; }

        [Input("rootDeviceType")]
        public Input<string>? RootDeviceType { get; set; }

        [Input("sshKeyName")]
        public Input<string>? SshKeyName { get; set; }

        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tenancy")]
        public Input<string>? Tenancy { get; set; }

        [Input("timeBasedAutoScaling")]
        public Input<Inputs.InstanceTimeBasedAutoScalingArgs>? TimeBasedAutoScaling { get; set; }

        [Input("virtualizationType")]
        public Input<string>? VirtualizationType { get; set; }

        [Input("volumes")]
        private InputList<string>? _volumes;
        public InputList<string> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<string>());
            set => _volumes = value;
        }

        public InstanceArgs()
        {
        }
    }
}
