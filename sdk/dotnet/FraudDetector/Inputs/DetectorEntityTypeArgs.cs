// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector.Inputs
{

    public sealed class DetectorEntityTypeArgs : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The time when the entity type was created.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// The description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inline")]
        public Input<bool>? Inline { get; set; }

        /// <summary>
        /// The time when the entity type was last updated.
        /// </summary>
        [Input("lastUpdatedTime")]
        public Input<string>? LastUpdatedTime { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.DetectorTagArgs>? _tags;

        /// <summary>
        /// Tags associated with this entity type.
        /// </summary>
        public InputList<Inputs.DetectorTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DetectorTagArgs>());
            set => _tags = value;
        }

        public DetectorEntityTypeArgs()
        {
        }
    }
}
