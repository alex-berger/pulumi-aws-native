// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.Connect
{
    /// <summary>
    /// The state of the contact flow module.
    /// </summary>
    [EnumType]
    public readonly struct ContactFlowModuleState : IEquatable<ContactFlowModuleState>
    {
        private readonly string _value;

        private ContactFlowModuleState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContactFlowModuleState Active { get; } = new ContactFlowModuleState("ACTIVE");
        public static ContactFlowModuleState Archived { get; } = new ContactFlowModuleState("ARCHIVED");

        public static bool operator ==(ContactFlowModuleState left, ContactFlowModuleState right) => left.Equals(right);
        public static bool operator !=(ContactFlowModuleState left, ContactFlowModuleState right) => !left.Equals(right);

        public static explicit operator string(ContactFlowModuleState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContactFlowModuleState other && Equals(other);
        public bool Equals(ContactFlowModuleState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the contact flow module.
    /// </summary>
    [EnumType]
    public readonly struct ContactFlowModuleStatus : IEquatable<ContactFlowModuleStatus>
    {
        private readonly string _value;

        private ContactFlowModuleStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContactFlowModuleStatus Published { get; } = new ContactFlowModuleStatus("PUBLISHED");
        public static ContactFlowModuleStatus Saved { get; } = new ContactFlowModuleStatus("SAVED");

        public static bool operator ==(ContactFlowModuleStatus left, ContactFlowModuleStatus right) => left.Equals(right);
        public static bool operator !=(ContactFlowModuleStatus left, ContactFlowModuleStatus right) => !left.Equals(right);

        public static explicit operator string(ContactFlowModuleStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContactFlowModuleStatus other && Equals(other);
        public bool Equals(ContactFlowModuleStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The state of the contact flow.
    /// </summary>
    [EnumType]
    public readonly struct ContactFlowState : IEquatable<ContactFlowState>
    {
        private readonly string _value;

        private ContactFlowState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContactFlowState Active { get; } = new ContactFlowState("ACTIVE");
        public static ContactFlowState Archived { get; } = new ContactFlowState("ARCHIVED");

        public static bool operator ==(ContactFlowState left, ContactFlowState right) => left.Equals(right);
        public static bool operator !=(ContactFlowState left, ContactFlowState right) => !left.Equals(right);

        public static explicit operator string(ContactFlowState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContactFlowState other && Equals(other);
        public bool Equals(ContactFlowState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the contact flow.
    /// </summary>
    [EnumType]
    public readonly struct ContactFlowType : IEquatable<ContactFlowType>
    {
        private readonly string _value;

        private ContactFlowType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContactFlowType ContactFlow { get; } = new ContactFlowType("CONTACT_FLOW");
        public static ContactFlowType CustomerQueue { get; } = new ContactFlowType("CUSTOMER_QUEUE");
        public static ContactFlowType CustomerHold { get; } = new ContactFlowType("CUSTOMER_HOLD");
        public static ContactFlowType CustomerWhisper { get; } = new ContactFlowType("CUSTOMER_WHISPER");
        public static ContactFlowType AgentHold { get; } = new ContactFlowType("AGENT_HOLD");
        public static ContactFlowType AgentWhisper { get; } = new ContactFlowType("AGENT_WHISPER");
        public static ContactFlowType OutboundWhisper { get; } = new ContactFlowType("OUTBOUND_WHISPER");
        public static ContactFlowType AgentTransfer { get; } = new ContactFlowType("AGENT_TRANSFER");
        public static ContactFlowType QueueTransfer { get; } = new ContactFlowType("QUEUE_TRANSFER");

        public static bool operator ==(ContactFlowType left, ContactFlowType right) => left.Equals(right);
        public static bool operator !=(ContactFlowType left, ContactFlowType right) => !left.Equals(right);

        public static explicit operator string(ContactFlowType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContactFlowType other && Equals(other);
        public bool Equals(ContactFlowType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The day that the hours of operation applies to.
    /// </summary>
    [EnumType]
    public readonly struct HoursOfOperationConfigDay : IEquatable<HoursOfOperationConfigDay>
    {
        private readonly string _value;

        private HoursOfOperationConfigDay(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HoursOfOperationConfigDay Sunday { get; } = new HoursOfOperationConfigDay("SUNDAY");
        public static HoursOfOperationConfigDay Monday { get; } = new HoursOfOperationConfigDay("MONDAY");
        public static HoursOfOperationConfigDay Tuesday { get; } = new HoursOfOperationConfigDay("TUESDAY");
        public static HoursOfOperationConfigDay Wednesday { get; } = new HoursOfOperationConfigDay("WEDNESDAY");
        public static HoursOfOperationConfigDay Thursday { get; } = new HoursOfOperationConfigDay("THURSDAY");
        public static HoursOfOperationConfigDay Friday { get; } = new HoursOfOperationConfigDay("FRIDAY");
        public static HoursOfOperationConfigDay Saturday { get; } = new HoursOfOperationConfigDay("SATURDAY");

        public static bool operator ==(HoursOfOperationConfigDay left, HoursOfOperationConfigDay right) => left.Equals(right);
        public static bool operator !=(HoursOfOperationConfigDay left, HoursOfOperationConfigDay right) => !left.Equals(right);

        public static explicit operator string(HoursOfOperationConfigDay value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HoursOfOperationConfigDay other && Equals(other);
        public bool Equals(HoursOfOperationConfigDay other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).
    /// </summary>
    [EnumType]
    public readonly struct QuickConnectType : IEquatable<QuickConnectType>
    {
        private readonly string _value;

        private QuickConnectType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static QuickConnectType PhoneNumber { get; } = new QuickConnectType("PHONE_NUMBER");
        public static QuickConnectType Queue { get; } = new QuickConnectType("QUEUE");
        public static QuickConnectType User { get; } = new QuickConnectType("USER");

        public static bool operator ==(QuickConnectType left, QuickConnectType right) => left.Equals(right);
        public static bool operator !=(QuickConnectType left, QuickConnectType right) => !left.Equals(right);

        public static explicit operator string(QuickConnectType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is QuickConnectType other && Equals(other);
        public bool Equals(QuickConnectType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The phone type.
    /// </summary>
    [EnumType]
    public readonly struct UserPhoneType : IEquatable<UserPhoneType>
    {
        private readonly string _value;

        private UserPhoneType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static UserPhoneType SoftPhone { get; } = new UserPhoneType("SOFT_PHONE");
        public static UserPhoneType DeskPhone { get; } = new UserPhoneType("DESK_PHONE");

        public static bool operator ==(UserPhoneType left, UserPhoneType right) => left.Equals(right);
        public static bool operator !=(UserPhoneType left, UserPhoneType right) => !left.Equals(right);

        public static explicit operator string(UserPhoneType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UserPhoneType other && Equals(other);
        public bool Equals(UserPhoneType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
