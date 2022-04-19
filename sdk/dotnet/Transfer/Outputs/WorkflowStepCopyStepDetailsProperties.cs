// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer.Outputs
{

    /// <summary>
    /// Details for a step that performs a file copy.
    /// </summary>
    [OutputType]
    public sealed class WorkflowStepCopyStepDetailsProperties
    {
        public readonly Outputs.WorkflowInputFileLocation? DestinationFileLocation;
        /// <summary>
        /// The name of the step, used as an identifier.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.
        /// </summary>
        public readonly Pulumi.AwsNative.Transfer.WorkflowStepCopyStepDetailsPropertiesOverwriteExisting? OverwriteExisting;
        /// <summary>
        /// Specifies which file to use as input to the workflow step.
        /// </summary>
        public readonly string? SourceFileLocation;

        [OutputConstructor]
        private WorkflowStepCopyStepDetailsProperties(
            Outputs.WorkflowInputFileLocation? destinationFileLocation,

            string? name,

            Pulumi.AwsNative.Transfer.WorkflowStepCopyStepDetailsPropertiesOverwriteExisting? overwriteExisting,

            string? sourceFileLocation)
        {
            DestinationFileLocation = destinationFileLocation;
            Name = name;
            OverwriteExisting = overwriteExisting;
            SourceFileLocation = sourceFileLocation;
        }
    }
}
