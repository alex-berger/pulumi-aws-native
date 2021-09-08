// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-inputrecordtables.html
    /// </summary>
    [OutputType]
    public sealed class MLTransformInputRecordTables
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-mltransform-inputrecordtables.html#cfn-glue-mltransform-inputrecordtables-gluetables
        /// </summary>
        public readonly ImmutableArray<Outputs.MLTransformGlueTables> GlueTables;

        [OutputConstructor]
        private MLTransformInputRecordTables(ImmutableArray<Outputs.MLTransformGlueTables> glueTables)
        {
            GlueTables = glueTables;
        }
    }
}
