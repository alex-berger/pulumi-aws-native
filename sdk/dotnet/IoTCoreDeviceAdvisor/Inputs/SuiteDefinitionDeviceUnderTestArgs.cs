// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTCoreDeviceAdvisor.Inputs
{

    public sealed class SuiteDefinitionDeviceUnderTestArgs : Pulumi.ResourceArgs
    {
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        [Input("thingArn")]
        public Input<string>? ThingArn { get; set; }

        public SuiteDefinitionDeviceUnderTestArgs()
        {
        }
    }
}
