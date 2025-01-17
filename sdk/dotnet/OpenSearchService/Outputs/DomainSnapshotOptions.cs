// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Outputs
{

    [OutputType]
    public sealed class DomainSnapshotOptions
    {
        public readonly int? AutomatedSnapshotStartHour;

        [OutputConstructor]
        private DomainSnapshotOptions(int? automatedSnapshotStartHour)
        {
            AutomatedSnapshotStartHour = automatedSnapshotStartHour;
        }
    }
}
