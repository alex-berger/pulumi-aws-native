// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Inputs
{

    public sealed class RecordSetGroupRecordSetArgs : Pulumi.ResourceArgs
    {
        [Input("aliasTarget")]
        public Input<Inputs.RecordSetGroupAliasTargetArgs>? AliasTarget { get; set; }

        [Input("failover")]
        public Input<string>? Failover { get; set; }

        [Input("geoLocation")]
        public Input<Inputs.RecordSetGroupGeoLocationArgs>? GeoLocation { get; set; }

        [Input("healthCheckId")]
        public Input<string>? HealthCheckId { get; set; }

        [Input("hostedZoneId")]
        public Input<string>? HostedZoneId { get; set; }

        [Input("hostedZoneName")]
        public Input<string>? HostedZoneName { get; set; }

        [Input("multiValueAnswer")]
        public Input<bool>? MultiValueAnswer { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("resourceRecords")]
        private InputList<string>? _resourceRecords;
        public InputList<string> ResourceRecords
        {
            get => _resourceRecords ?? (_resourceRecords = new InputList<string>());
            set => _resourceRecords = value;
        }

        [Input("setIdentifier")]
        public Input<string>? SetIdentifier { get; set; }

        [Input("tTL")]
        public Input<string>? TTL { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public RecordSetGroupRecordSetArgs()
        {
        }
    }
}
