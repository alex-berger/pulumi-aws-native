// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationInsights.Outputs
{

    /// <summary>
    /// The JMX Prometheus Exporter settings.
    /// </summary>
    [OutputType]
    public sealed class ApplicationJMXPrometheusExporter
    {
        /// <summary>
        /// Java agent host port
        /// </summary>
        public readonly string? HostPort;
        /// <summary>
        /// JMX service URL.
        /// </summary>
        public readonly string? JMXURL;
        /// <summary>
        /// Prometheus exporter port
        /// </summary>
        public readonly string? PrometheusPort;

        [OutputConstructor]
        private ApplicationJMXPrometheusExporter(
            string? hostPort,

            string? jMXURL,

            string? prometheusPort)
        {
            HostPort = hostPort;
            JMXURL = jMXURL;
            PrometheusPort = prometheusPort;
        }
    }
}
