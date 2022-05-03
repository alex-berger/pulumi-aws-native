// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Redshift
{
    [EnumType]
    public readonly struct EventSubscriptionEventCategoriesItem : IEquatable<EventSubscriptionEventCategoriesItem>
    {
        private readonly string _value;

        private EventSubscriptionEventCategoriesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EventSubscriptionEventCategoriesItem Configuration { get; } = new EventSubscriptionEventCategoriesItem("configuration");
        public static EventSubscriptionEventCategoriesItem Management { get; } = new EventSubscriptionEventCategoriesItem("management");
        public static EventSubscriptionEventCategoriesItem Monitoring { get; } = new EventSubscriptionEventCategoriesItem("monitoring");
        public static EventSubscriptionEventCategoriesItem Security { get; } = new EventSubscriptionEventCategoriesItem("security");
        public static EventSubscriptionEventCategoriesItem Pending { get; } = new EventSubscriptionEventCategoriesItem("pending");

        public static bool operator ==(EventSubscriptionEventCategoriesItem left, EventSubscriptionEventCategoriesItem right) => left.Equals(right);
        public static bool operator !=(EventSubscriptionEventCategoriesItem left, EventSubscriptionEventCategoriesItem right) => !left.Equals(right);

        public static explicit operator string(EventSubscriptionEventCategoriesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EventSubscriptionEventCategoriesItem other && Equals(other);
        public bool Equals(EventSubscriptionEventCategoriesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies the Amazon Redshift event severity to be published by the event notification subscription.
    /// </summary>
    [EnumType]
    public readonly struct EventSubscriptionSeverity : IEquatable<EventSubscriptionSeverity>
    {
        private readonly string _value;

        private EventSubscriptionSeverity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EventSubscriptionSeverity Error { get; } = new EventSubscriptionSeverity("ERROR");
        public static EventSubscriptionSeverity Info { get; } = new EventSubscriptionSeverity("INFO");

        public static bool operator ==(EventSubscriptionSeverity left, EventSubscriptionSeverity right) => left.Equals(right);
        public static bool operator !=(EventSubscriptionSeverity left, EventSubscriptionSeverity right) => !left.Equals(right);

        public static explicit operator string(EventSubscriptionSeverity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EventSubscriptionSeverity other && Equals(other);
        public bool Equals(EventSubscriptionSeverity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of source that will be generating the events.
    /// </summary>
    [EnumType]
    public readonly struct EventSubscriptionSourceType : IEquatable<EventSubscriptionSourceType>
    {
        private readonly string _value;

        private EventSubscriptionSourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EventSubscriptionSourceType Cluster { get; } = new EventSubscriptionSourceType("cluster");
        public static EventSubscriptionSourceType ClusterParameterGroup { get; } = new EventSubscriptionSourceType("cluster-parameter-group");
        public static EventSubscriptionSourceType ClusterSecurityGroup { get; } = new EventSubscriptionSourceType("cluster-security-group");
        public static EventSubscriptionSourceType ClusterSnapshot { get; } = new EventSubscriptionSourceType("cluster-snapshot");
        public static EventSubscriptionSourceType ScheduledAction { get; } = new EventSubscriptionSourceType("scheduled-action");

        public static bool operator ==(EventSubscriptionSourceType left, EventSubscriptionSourceType right) => left.Equals(right);
        public static bool operator !=(EventSubscriptionSourceType left, EventSubscriptionSourceType right) => !left.Equals(right);

        public static explicit operator string(EventSubscriptionSourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EventSubscriptionSourceType other && Equals(other);
        public bool Equals(EventSubscriptionSourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the Amazon Redshift event notification subscription.
    /// </summary>
    [EnumType]
    public readonly struct EventSubscriptionStatus : IEquatable<EventSubscriptionStatus>
    {
        private readonly string _value;

        private EventSubscriptionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EventSubscriptionStatus Active { get; } = new EventSubscriptionStatus("active");
        public static EventSubscriptionStatus NoPermission { get; } = new EventSubscriptionStatus("no-permission");
        public static EventSubscriptionStatus TopicNotExist { get; } = new EventSubscriptionStatus("topic-not-exist");

        public static bool operator ==(EventSubscriptionStatus left, EventSubscriptionStatus right) => left.Equals(right);
        public static bool operator !=(EventSubscriptionStatus left, EventSubscriptionStatus right) => !left.Equals(right);

        public static explicit operator string(EventSubscriptionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EventSubscriptionStatus other && Equals(other);
        public bool Equals(EventSubscriptionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The state of the scheduled action.
    /// </summary>
    [EnumType]
    public readonly struct ScheduledActionState : IEquatable<ScheduledActionState>
    {
        private readonly string _value;

        private ScheduledActionState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduledActionState Active { get; } = new ScheduledActionState("ACTIVE");
        public static ScheduledActionState Disabled { get; } = new ScheduledActionState("DISABLED");

        public static bool operator ==(ScheduledActionState left, ScheduledActionState right) => left.Equals(right);
        public static bool operator !=(ScheduledActionState left, ScheduledActionState right) => !left.Equals(right);

        public static explicit operator string(ScheduledActionState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduledActionState other && Equals(other);
        public bool Equals(ScheduledActionState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
