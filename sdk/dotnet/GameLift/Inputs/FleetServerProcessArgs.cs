// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Inputs
{

    /// <summary>
    /// A set of instructions for launching server processes on each instance in a fleet. Each instruction set identifies the location of the server executable, optional launch parameters, and the number of server processes with this configuration to maintain concurrently on the instance. Server process configurations make up a fleet's RuntimeConfiguration.
    /// </summary>
    public sealed class FleetServerProcessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of server processes that use this configuration to run concurrently on an instance.
        /// </summary>
        [Input("concurrentExecutions", required: true)]
        public Input<int> ConcurrentExecutions { get; set; } = null!;

        /// <summary>
        /// The location of the server executable in a custom game build or the name of the Realtime script file that contains the Init() function. Game builds and Realtime scripts are installed on instances at the root:
        /// 
        /// Windows (for custom game builds only): C:\game. Example: "C:\game\MyGame\server.exe"
        /// 
        /// Linux: /local/game. Examples: "/local/game/MyGame/server.exe" or "/local/game/MyRealtimeScript.js"
        /// </summary>
        [Input("launchPath", required: true)]
        public Input<string> LaunchPath { get; set; } = null!;

        /// <summary>
        /// An optional list of parameters to pass to the server executable or Realtime script on launch.
        /// </summary>
        [Input("parameters")]
        public Input<string>? Parameters { get; set; }

        public FleetServerProcessArgs()
        {
        }
    }
}
