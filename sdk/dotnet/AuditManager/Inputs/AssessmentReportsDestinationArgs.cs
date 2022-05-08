// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AuditManager.Inputs
{

    /// <summary>
    /// The destination in which evidence reports are stored for the specified assessment.
    /// </summary>
    public sealed class AssessmentReportsDestinationArgs : Pulumi.ResourceArgs
    {
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        [Input("destinationType")]
        public Input<Pulumi.AwsNative.AuditManager.AssessmentReportDestinationType>? DestinationType { get; set; }

        public AssessmentReportsDestinationArgs()
        {
        }
    }
}
