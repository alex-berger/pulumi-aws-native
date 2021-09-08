// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineversion.html
    /// </summary>
    [OutputType]
    public sealed class WorkGroupEngineVersion
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineversion.html#cfn-athena-workgroup-engineversion-effectiveengineversion
        /// </summary>
        public readonly string? EffectiveEngineVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineversion.html#cfn-athena-workgroup-engineversion-selectedengineversion
        /// </summary>
        public readonly string? SelectedEngineVersion;

        [OutputConstructor]
        private WorkGroupEngineVersion(
            string? effectiveEngineVersion,

            string? selectedEngineVersion)
        {
            EffectiveEngineVersion = effectiveEngineVersion;
            SelectedEngineVersion = selectedEngineVersion;
        }
    }
}