// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Kinesis
{
    /// <summary>
    /// The encryption type to use. The only valid value is KMS. 
    /// </summary>
    [EnumType]
    public readonly struct StreamEncryptionEncryptionType : IEquatable<StreamEncryptionEncryptionType>
    {
        private readonly string _value;

        private StreamEncryptionEncryptionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StreamEncryptionEncryptionType Kms { get; } = new StreamEncryptionEncryptionType("KMS");

        public static bool operator ==(StreamEncryptionEncryptionType left, StreamEncryptionEncryptionType right) => left.Equals(right);
        public static bool operator !=(StreamEncryptionEncryptionType left, StreamEncryptionEncryptionType right) => !left.Equals(right);

        public static explicit operator string(StreamEncryptionEncryptionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StreamEncryptionEncryptionType other && Equals(other);
        public bool Equals(StreamEncryptionEncryptionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The mode of the stream
    /// </summary>
    [EnumType]
    public readonly struct StreamModeDetailsStreamMode : IEquatable<StreamModeDetailsStreamMode>
    {
        private readonly string _value;

        private StreamModeDetailsStreamMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StreamModeDetailsStreamMode OnDemand { get; } = new StreamModeDetailsStreamMode("ON_DEMAND");
        public static StreamModeDetailsStreamMode Provisioned { get; } = new StreamModeDetailsStreamMode("PROVISIONED");

        public static bool operator ==(StreamModeDetailsStreamMode left, StreamModeDetailsStreamMode right) => left.Equals(right);
        public static bool operator !=(StreamModeDetailsStreamMode left, StreamModeDetailsStreamMode right) => !left.Equals(right);

        public static explicit operator string(StreamModeDetailsStreamMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StreamModeDetailsStreamMode other && Equals(other);
        public bool Equals(StreamModeDetailsStreamMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}