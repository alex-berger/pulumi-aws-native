// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-configuration.html
    /// </summary>
    public sealed class ClusterConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-configuration.html#cfn-elasticmapreduce-cluster-configuration-classification
        /// </summary>
        [Input("classification")]
        public Input<string>? Classification { get; set; }

        [Input("configurationProperties")]
        private InputMap<string>? _configurationProperties;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-configuration.html#cfn-elasticmapreduce-cluster-configuration-configurationproperties
        /// </summary>
        public InputMap<string> ConfigurationProperties
        {
            get => _configurationProperties ?? (_configurationProperties = new InputMap<string>());
            set => _configurationProperties = value;
        }

        [Input("configurations")]
        private InputList<Inputs.ClusterConfigurationArgs>? _configurations;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-configuration.html#cfn-elasticmapreduce-cluster-configuration-configurations
        /// </summary>
        public InputList<Inputs.ClusterConfigurationArgs> Configurations
        {
            get => _configurations ?? (_configurations = new InputList<Inputs.ClusterConfigurationArgs>());
            set => _configurations = value;
        }

        public ClusterConfigurationArgs()
        {
        }
    }
}
