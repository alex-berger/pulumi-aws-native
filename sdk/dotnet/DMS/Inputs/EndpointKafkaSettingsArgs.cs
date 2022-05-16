// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS.Inputs
{

    public sealed class EndpointKafkaSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("broker")]
        public Input<string>? Broker { get; set; }

        [Input("includeControlDetails")]
        public Input<bool>? IncludeControlDetails { get; set; }

        [Input("includeNullAndEmpty")]
        public Input<bool>? IncludeNullAndEmpty { get; set; }

        [Input("includePartitionValue")]
        public Input<bool>? IncludePartitionValue { get; set; }

        [Input("includeTableAlterOperations")]
        public Input<bool>? IncludeTableAlterOperations { get; set; }

        [Input("includeTransactionDetails")]
        public Input<bool>? IncludeTransactionDetails { get; set; }

        [Input("messageFormat")]
        public Input<string>? MessageFormat { get; set; }

        [Input("messageMaxBytes")]
        public Input<int>? MessageMaxBytes { get; set; }

        [Input("noHexPrefix")]
        public Input<bool>? NoHexPrefix { get; set; }

        [Input("partitionIncludeSchemaTable")]
        public Input<bool>? PartitionIncludeSchemaTable { get; set; }

        [Input("saslPassword")]
        public Input<string>? SaslPassword { get; set; }

        [Input("saslUserName")]
        public Input<string>? SaslUserName { get; set; }

        [Input("securityProtocol")]
        public Input<string>? SecurityProtocol { get; set; }

        [Input("sslCaCertificateArn")]
        public Input<string>? SslCaCertificateArn { get; set; }

        [Input("sslClientCertificateArn")]
        public Input<string>? SslClientCertificateArn { get; set; }

        [Input("sslClientKeyArn")]
        public Input<string>? SslClientKeyArn { get; set; }

        [Input("sslClientKeyPassword")]
        public Input<string>? SslClientKeyPassword { get; set; }

        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public EndpointKafkaSettingsArgs()
        {
        }
    }
}