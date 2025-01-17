// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    public static class GetFleet
    {
        /// <summary>
        /// The AWS::GameLift::Fleet resource creates an Amazon GameLift (GameLift) fleet to host game servers.  A fleet is a set of EC2 instances, each of which can host multiple game sessions.
        /// </summary>
        public static Task<GetFleetResult> InvokeAsync(GetFleetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFleetResult>("aws-native:gamelift:getFleet", args ?? new GetFleetArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::GameLift::Fleet resource creates an Amazon GameLift (GameLift) fleet to host game servers.  A fleet is a set of EC2 instances, each of which can host multiple game sessions.
        /// </summary>
        public static Output<GetFleetResult> Invoke(GetFleetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFleetResult>("aws-native:gamelift:getFleet", args ?? new GetFleetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFleetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique fleet ID
        /// </summary>
        [Input("fleetId", required: true)]
        public string FleetId { get; set; } = null!;

        public GetFleetArgs()
        {
        }
    }

    public sealed class GetFleetInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique fleet ID
        /// </summary>
        [Input("fleetId", required: true)]
        public Input<string> FleetId { get; set; } = null!;

        public GetFleetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFleetResult
    {
        /// <summary>
        /// A human-readable description of a fleet.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// [DEPRECATED] The number of EC2 instances that you want this fleet to host. When creating a new fleet, GameLift automatically sets this value to "1" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.
        /// </summary>
        public readonly int? DesiredEC2Instances;
        /// <summary>
        /// A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift server.
        /// </summary>
        public readonly ImmutableArray<Outputs.FleetIpPermission> EC2InboundPermissions;
        /// <summary>
        /// Unique fleet ID
        /// </summary>
        public readonly string? FleetId;
        public readonly ImmutableArray<Outputs.FleetLocationConfiguration> Locations;
        /// <summary>
        /// [DEPRECATED] The maximum value that is allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to "1". Once the fleet is active, you can change this value.
        /// </summary>
        public readonly int? MaxSize;
        /// <summary>
        /// The name of an Amazon CloudWatch metric group. A metric group aggregates the metrics for all fleets in the group. Specify a string containing the metric group name. You can use an existing name or use a new name to create a new metric group. Currently, this parameter can have only one string.
        /// </summary>
        public readonly ImmutableArray<string> MetricGroups;
        /// <summary>
        /// [DEPRECATED] The minimum value allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to "0". After the fleet is active, you can change this value.
        /// </summary>
        public readonly int? MinSize;
        /// <summary>
        /// A descriptive label that is associated with a fleet. Fleet names do not need to be unique.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A game session protection policy to apply to all game sessions hosted on instances in this fleet. When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.
        /// </summary>
        public readonly Pulumi.AwsNative.GameLift.FleetNewGameSessionProtectionPolicy? NewGameSessionProtectionPolicy;
        /// <summary>
        /// A policy that limits the number of game sessions an individual player can create over a span of time for this fleet.
        /// </summary>
        public readonly Outputs.FleetResourceCreationLimitPolicy? ResourceCreationLimitPolicy;
        /// <summary>
        /// Instructions for launching server processes on each instance in the fleet. Server processes run either a custom game build executable or a Realtime script. The runtime configuration defines the server executables or launch script file, launch parameters, and the number of processes to run concurrently on each instance. When creating a fleet, the runtime configuration must have at least one server process configuration; otherwise the request fails with an invalid request exception.
        /// 
        /// This parameter is required unless the parameters ServerLaunchPath and ServerLaunchParameters are defined. Runtime configuration has replaced these parameters, but fleets that use them will continue to work.
        /// </summary>
        public readonly Outputs.FleetRuntimeConfiguration? RuntimeConfiguration;

        [OutputConstructor]
        private GetFleetResult(
            string? description,

            int? desiredEC2Instances,

            ImmutableArray<Outputs.FleetIpPermission> eC2InboundPermissions,

            string? fleetId,

            ImmutableArray<Outputs.FleetLocationConfiguration> locations,

            int? maxSize,

            ImmutableArray<string> metricGroups,

            int? minSize,

            string? name,

            Pulumi.AwsNative.GameLift.FleetNewGameSessionProtectionPolicy? newGameSessionProtectionPolicy,

            Outputs.FleetResourceCreationLimitPolicy? resourceCreationLimitPolicy,

            Outputs.FleetRuntimeConfiguration? runtimeConfiguration)
        {
            Description = description;
            DesiredEC2Instances = desiredEC2Instances;
            EC2InboundPermissions = eC2InboundPermissions;
            FleetId = fleetId;
            Locations = locations;
            MaxSize = maxSize;
            MetricGroups = metricGroups;
            MinSize = minSize;
            Name = name;
            NewGameSessionProtectionPolicy = newGameSessionProtectionPolicy;
            ResourceCreationLimitPolicy = resourceCreationLimitPolicy;
            RuntimeConfiguration = runtimeConfiguration;
        }
    }
}
