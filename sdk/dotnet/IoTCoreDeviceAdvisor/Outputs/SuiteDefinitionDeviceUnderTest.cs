// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTCoreDeviceAdvisor.Outputs
{

    [OutputType]
    public sealed class SuiteDefinitionDeviceUnderTest
    {
        public readonly string? CertificateArn;
        public readonly string? ThingArn;

        [OutputConstructor]
        private SuiteDefinitionDeviceUnderTest(
            string? certificateArn,

            string? thingArn)
        {
            CertificateArn = certificateArn;
            ThingArn = thingArn;
        }
    }
}
