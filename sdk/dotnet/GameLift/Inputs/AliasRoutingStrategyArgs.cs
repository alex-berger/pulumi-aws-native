// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Inputs
{

    public sealed class AliasRoutingStrategyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for a fleet that the alias points to. If you specify SIMPLE for the Type property, you must specify this property.
        /// </summary>
        [Input("fleetId")]
        public Input<string>? FleetId { get; set; }

        /// <summary>
        /// The message text to be used with a terminal routing strategy. If you specify TERMINAL for the Type property, you must specify this property.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Simple routing strategy. The alias resolves to one specific fleet. Use this type when routing to active fleets.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.GameLift.AliasRoutingStrategyType> Type { get; set; } = null!;

        public AliasRoutingStrategyArgs()
        {
        }
    }
}
