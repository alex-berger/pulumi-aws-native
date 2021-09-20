// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.FinSpace
{
    /// <summary>
    /// Federation mode used with the Environment
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentFederationMode : IEquatable<EnvironmentFederationMode>
    {
        private readonly string _value;

        private EnvironmentFederationMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentFederationMode Local { get; } = new EnvironmentFederationMode("LOCAL");
        public static EnvironmentFederationMode Federated { get; } = new EnvironmentFederationMode("FEDERATED");

        public static bool operator ==(EnvironmentFederationMode left, EnvironmentFederationMode right) => left.Equals(right);
        public static bool operator !=(EnvironmentFederationMode left, EnvironmentFederationMode right) => !left.Equals(right);

        public static explicit operator string(EnvironmentFederationMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentFederationMode other && Equals(other);
        public bool Equals(EnvironmentFederationMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// State of the Environment
    /// </summary>
    [EnumType]
    public readonly struct EnvironmentStatus : IEquatable<EnvironmentStatus>
    {
        private readonly string _value;

        private EnvironmentStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EnvironmentStatus CreateRequested { get; } = new EnvironmentStatus("CREATE_REQUESTED");
        public static EnvironmentStatus Creating { get; } = new EnvironmentStatus("CREATING");
        public static EnvironmentStatus Created { get; } = new EnvironmentStatus("CREATED");
        public static EnvironmentStatus DeleteRequested { get; } = new EnvironmentStatus("DELETE_REQUESTED");
        public static EnvironmentStatus Deleting { get; } = new EnvironmentStatus("DELETING");
        public static EnvironmentStatus Deleted { get; } = new EnvironmentStatus("DELETED");
        public static EnvironmentStatus FailedCreation { get; } = new EnvironmentStatus("FAILED_CREATION");
        public static EnvironmentStatus FailedDeletion { get; } = new EnvironmentStatus("FAILED_DELETION");
        public static EnvironmentStatus RetryDeletion { get; } = new EnvironmentStatus("RETRY_DELETION");
        public static EnvironmentStatus Suspended { get; } = new EnvironmentStatus("SUSPENDED");

        public static bool operator ==(EnvironmentStatus left, EnvironmentStatus right) => left.Equals(right);
        public static bool operator !=(EnvironmentStatus left, EnvironmentStatus right) => !left.Equals(right);

        public static explicit operator string(EnvironmentStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EnvironmentStatus other && Equals(other);
        public bool Equals(EnvironmentStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}