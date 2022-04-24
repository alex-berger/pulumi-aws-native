// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetDevice
    {
        /// <summary>
        /// Resource schema for AWS::SageMaker::Device
        /// </summary>
        public static Task<GetDeviceResult> InvokeAsync(GetDeviceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDeviceResult>("aws-native:sagemaker:getDevice", args ?? new GetDeviceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::SageMaker::Device
        /// </summary>
        public static Output<GetDeviceResult> Invoke(GetDeviceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDeviceResult>("aws-native:sagemaker:getDevice", args ?? new GetDeviceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeviceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the edge device fleet
        /// </summary>
        [Input("deviceFleetName", required: true)]
        public string DeviceFleetName { get; set; } = null!;

        public GetDeviceArgs()
        {
        }
    }

    public sealed class GetDeviceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the edge device fleet
        /// </summary>
        [Input("deviceFleetName", required: true)]
        public Input<string> DeviceFleetName { get; set; } = null!;

        public GetDeviceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDeviceResult
    {
        /// <summary>
        /// The Edge Device you want to register against a device fleet
        /// </summary>
        public readonly Outputs.Device? DeviceValue;
        /// <summary>
        /// Associate tags with the resource
        /// </summary>
        public readonly ImmutableArray<Outputs.DeviceTag> Tags;

        [OutputConstructor]
        private GetDeviceResult(
            Outputs.Device? device,

            ImmutableArray<Outputs.DeviceTag> tags)
        {
            DeviceValue = device;
            Tags = tags;
        }
    }
}
