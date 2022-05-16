// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless
{
    public static class GetFuotaTask
    {
        /// <summary>
        /// Create and manage FUOTA tasks.
        /// </summary>
        public static Task<GetFuotaTaskResult> InvokeAsync(GetFuotaTaskArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFuotaTaskResult>("aws-native:iotwireless:getFuotaTask", args ?? new GetFuotaTaskArgs(), options.WithDefaults());

        /// <summary>
        /// Create and manage FUOTA tasks.
        /// </summary>
        public static Output<GetFuotaTaskResult> Invoke(GetFuotaTaskInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFuotaTaskResult>("aws-native:iotwireless:getFuotaTask", args ?? new GetFuotaTaskInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFuotaTaskArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// FUOTA task id. Returned after successful create.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetFuotaTaskArgs()
        {
        }
    }

    public sealed class GetFuotaTaskInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// FUOTA task id. Returned after successful create.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetFuotaTaskInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFuotaTaskResult
    {
        /// <summary>
        /// FUOTA task arn. Returned after successful create.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Multicast group to associate. Only for update request.
        /// </summary>
        public readonly string? AssociateMulticastGroup;
        /// <summary>
        /// Wireless device to associate. Only for update request.
        /// </summary>
        public readonly string? AssociateWirelessDevice;
        /// <summary>
        /// FUOTA task description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Multicast group to disassociate. Only for update request.
        /// </summary>
        public readonly string? DisassociateMulticastGroup;
        /// <summary>
        /// Wireless device to disassociate. Only for update request.
        /// </summary>
        public readonly string? DisassociateWirelessDevice;
        /// <summary>
        /// FUOTA task firmware update image binary S3 link
        /// </summary>
        public readonly string? FirmwareUpdateImage;
        /// <summary>
        /// FUOTA task firmware IAM role for reading S3
        /// </summary>
        public readonly string? FirmwareUpdateRole;
        /// <summary>
        /// FUOTA task status. Returned after successful read.
        /// </summary>
        public readonly string? FuotaTaskStatus;
        /// <summary>
        /// FUOTA task id. Returned after successful create.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// FUOTA task LoRaWAN
        /// </summary>
        public readonly Outputs.FuotaTaskLoRaWAN? LoRaWAN;
        /// <summary>
        /// Name of FUOTA task
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A list of key-value pairs that contain metadata for the FUOTA task.
        /// </summary>
        public readonly ImmutableArray<Outputs.FuotaTaskTag> Tags;

        [OutputConstructor]
        private GetFuotaTaskResult(
            string? arn,

            string? associateMulticastGroup,

            string? associateWirelessDevice,

            string? description,

            string? disassociateMulticastGroup,

            string? disassociateWirelessDevice,

            string? firmwareUpdateImage,

            string? firmwareUpdateRole,

            string? fuotaTaskStatus,

            string? id,

            Outputs.FuotaTaskLoRaWAN? loRaWAN,

            string? name,

            ImmutableArray<Outputs.FuotaTaskTag> tags)
        {
            Arn = arn;
            AssociateMulticastGroup = associateMulticastGroup;
            AssociateWirelessDevice = associateWirelessDevice;
            Description = description;
            DisassociateMulticastGroup = disassociateMulticastGroup;
            DisassociateWirelessDevice = disassociateWirelessDevice;
            FirmwareUpdateImage = firmwareUpdateImage;
            FirmwareUpdateRole = firmwareUpdateRole;
            FuotaTaskStatus = fuotaTaskStatus;
            Id = id;
            LoRaWAN = loRaWAN;
            Name = name;
            Tags = tags;
        }
    }
}