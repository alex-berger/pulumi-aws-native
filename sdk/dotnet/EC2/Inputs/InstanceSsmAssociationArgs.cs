// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html
    /// </summary>
    public sealed class InstanceSsmAssociationArgs : Pulumi.ResourceArgs
    {
        [Input("associationParameters")]
        private InputList<Inputs.InstanceAssociationParameterArgs>? _associationParameters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html#cfn-ec2-instance-ssmassociations-associationparameters
        /// </summary>
        public InputList<Inputs.InstanceAssociationParameterArgs> AssociationParameters
        {
            get => _associationParameters ?? (_associationParameters = new InputList<Inputs.InstanceAssociationParameterArgs>());
            set => _associationParameters = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html#cfn-ec2-instance-ssmassociations-documentname
        /// </summary>
        [Input("documentName", required: true)]
        public Input<string> DocumentName { get; set; } = null!;

        public InstanceSsmAssociationArgs()
        {
        }
    }
}
