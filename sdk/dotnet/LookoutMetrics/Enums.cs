// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.LookoutMetrics
{
    [EnumType]
    public readonly struct AnomalyDetectorCsvFormatDescriptorFileCompression : IEquatable<AnomalyDetectorCsvFormatDescriptorFileCompression>
    {
        private readonly string _value;

        private AnomalyDetectorCsvFormatDescriptorFileCompression(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AnomalyDetectorCsvFormatDescriptorFileCompression None { get; } = new AnomalyDetectorCsvFormatDescriptorFileCompression("NONE");
        public static AnomalyDetectorCsvFormatDescriptorFileCompression Gzip { get; } = new AnomalyDetectorCsvFormatDescriptorFileCompression("GZIP");

        public static bool operator ==(AnomalyDetectorCsvFormatDescriptorFileCompression left, AnomalyDetectorCsvFormatDescriptorFileCompression right) => left.Equals(right);
        public static bool operator !=(AnomalyDetectorCsvFormatDescriptorFileCompression left, AnomalyDetectorCsvFormatDescriptorFileCompression right) => !left.Equals(right);

        public static explicit operator string(AnomalyDetectorCsvFormatDescriptorFileCompression value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AnomalyDetectorCsvFormatDescriptorFileCompression other && Equals(other);
        public bool Equals(AnomalyDetectorCsvFormatDescriptorFileCompression other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Frequency of anomaly detection
    /// </summary>
    [EnumType]
    public readonly struct AnomalyDetectorFrequency : IEquatable<AnomalyDetectorFrequency>
    {
        private readonly string _value;

        private AnomalyDetectorFrequency(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AnomalyDetectorFrequency Pt5m { get; } = new AnomalyDetectorFrequency("PT5M");
        public static AnomalyDetectorFrequency Pt10m { get; } = new AnomalyDetectorFrequency("PT10M");
        public static AnomalyDetectorFrequency Pt1h { get; } = new AnomalyDetectorFrequency("PT1H");
        public static AnomalyDetectorFrequency P1d { get; } = new AnomalyDetectorFrequency("P1D");

        public static bool operator ==(AnomalyDetectorFrequency left, AnomalyDetectorFrequency right) => left.Equals(right);
        public static bool operator !=(AnomalyDetectorFrequency left, AnomalyDetectorFrequency right) => !left.Equals(right);

        public static explicit operator string(AnomalyDetectorFrequency value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AnomalyDetectorFrequency other && Equals(other);
        public bool Equals(AnomalyDetectorFrequency other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct AnomalyDetectorJsonFormatDescriptorFileCompression : IEquatable<AnomalyDetectorJsonFormatDescriptorFileCompression>
    {
        private readonly string _value;

        private AnomalyDetectorJsonFormatDescriptorFileCompression(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AnomalyDetectorJsonFormatDescriptorFileCompression None { get; } = new AnomalyDetectorJsonFormatDescriptorFileCompression("NONE");
        public static AnomalyDetectorJsonFormatDescriptorFileCompression Gzip { get; } = new AnomalyDetectorJsonFormatDescriptorFileCompression("GZIP");

        public static bool operator ==(AnomalyDetectorJsonFormatDescriptorFileCompression left, AnomalyDetectorJsonFormatDescriptorFileCompression right) => left.Equals(right);
        public static bool operator !=(AnomalyDetectorJsonFormatDescriptorFileCompression left, AnomalyDetectorJsonFormatDescriptorFileCompression right) => !left.Equals(right);

        public static explicit operator string(AnomalyDetectorJsonFormatDescriptorFileCompression value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AnomalyDetectorJsonFormatDescriptorFileCompression other && Equals(other);
        public bool Equals(AnomalyDetectorJsonFormatDescriptorFileCompression other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Operator used to aggregate metric values
    /// </summary>
    [EnumType]
    public readonly struct AnomalyDetectorMetricAggregationFunction : IEquatable<AnomalyDetectorMetricAggregationFunction>
    {
        private readonly string _value;

        private AnomalyDetectorMetricAggregationFunction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AnomalyDetectorMetricAggregationFunction Avg { get; } = new AnomalyDetectorMetricAggregationFunction("AVG");
        public static AnomalyDetectorMetricAggregationFunction Sum { get; } = new AnomalyDetectorMetricAggregationFunction("SUM");

        public static bool operator ==(AnomalyDetectorMetricAggregationFunction left, AnomalyDetectorMetricAggregationFunction right) => left.Equals(right);
        public static bool operator !=(AnomalyDetectorMetricAggregationFunction left, AnomalyDetectorMetricAggregationFunction right) => !left.Equals(right);

        public static explicit operator string(AnomalyDetectorMetricAggregationFunction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AnomalyDetectorMetricAggregationFunction other && Equals(other);
        public bool Equals(AnomalyDetectorMetricAggregationFunction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A frequency period to aggregate the data
    /// </summary>
    [EnumType]
    public readonly struct AnomalyDetectorMetricSetMetricSetFrequency : IEquatable<AnomalyDetectorMetricSetMetricSetFrequency>
    {
        private readonly string _value;

        private AnomalyDetectorMetricSetMetricSetFrequency(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AnomalyDetectorMetricSetMetricSetFrequency Pt5m { get; } = new AnomalyDetectorMetricSetMetricSetFrequency("PT5M");
        public static AnomalyDetectorMetricSetMetricSetFrequency Pt10m { get; } = new AnomalyDetectorMetricSetMetricSetFrequency("PT10M");
        public static AnomalyDetectorMetricSetMetricSetFrequency Pt1h { get; } = new AnomalyDetectorMetricSetMetricSetFrequency("PT1H");
        public static AnomalyDetectorMetricSetMetricSetFrequency P1d { get; } = new AnomalyDetectorMetricSetMetricSetFrequency("P1D");

        public static bool operator ==(AnomalyDetectorMetricSetMetricSetFrequency left, AnomalyDetectorMetricSetMetricSetFrequency right) => left.Equals(right);
        public static bool operator !=(AnomalyDetectorMetricSetMetricSetFrequency left, AnomalyDetectorMetricSetMetricSetFrequency right) => !left.Equals(right);

        public static explicit operator string(AnomalyDetectorMetricSetMetricSetFrequency value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AnomalyDetectorMetricSetMetricSetFrequency other && Equals(other);
        public bool Equals(AnomalyDetectorMetricSetMetricSetFrequency other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
