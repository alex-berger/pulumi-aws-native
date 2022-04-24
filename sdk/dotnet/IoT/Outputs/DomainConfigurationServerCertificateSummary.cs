// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    [OutputType]
    public sealed class DomainConfigurationServerCertificateSummary
    {
        public readonly string? ServerCertificateArn;
        public readonly Pulumi.AwsNative.IoT.DomainConfigurationServerCertificateSummaryServerCertificateStatus? ServerCertificateStatus;
        public readonly string? ServerCertificateStatusDetail;

        [OutputConstructor]
        private DomainConfigurationServerCertificateSummary(
            string? serverCertificateArn,

            Pulumi.AwsNative.IoT.DomainConfigurationServerCertificateSummaryServerCertificateStatus? serverCertificateStatus,

            string? serverCertificateStatusDetail)
        {
            ServerCertificateArn = serverCertificateArn;
            ServerCertificateStatus = serverCertificateStatus;
            ServerCertificateStatusDetail = serverCertificateStatusDetail;
        }
    }
}
