// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM.Inputs
{

    public sealed class AssociationInstanceAssociationOutputLocationArgs : Pulumi.ResourceArgs
    {
        [Input("s3Location")]
        public Input<Inputs.AssociationS3OutputLocationArgs>? S3Location { get; set; }

        public AssociationInstanceAssociationOutputLocationArgs()
        {
        }
    }
}
