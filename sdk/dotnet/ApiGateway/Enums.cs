// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.ApiGateway
{
    /// <summary>
    /// The method's authorization type.
    /// </summary>
    [EnumType]
    public readonly struct MethodAuthorizationType : IEquatable<MethodAuthorizationType>
    {
        private readonly string _value;

        private MethodAuthorizationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MethodAuthorizationType None { get; } = new MethodAuthorizationType("NONE");
        public static MethodAuthorizationType AwsIam { get; } = new MethodAuthorizationType("AWS_IAM");
        public static MethodAuthorizationType Custom { get; } = new MethodAuthorizationType("CUSTOM");
        public static MethodAuthorizationType CognitoUserPools { get; } = new MethodAuthorizationType("COGNITO_USER_POOLS");

        public static bool operator ==(MethodAuthorizationType left, MethodAuthorizationType right) => left.Equals(right);
        public static bool operator !=(MethodAuthorizationType left, MethodAuthorizationType right) => !left.Equals(right);

        public static explicit operator string(MethodAuthorizationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MethodAuthorizationType other && Equals(other);
        public bool Equals(MethodAuthorizationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the network connection to the integration endpoint.
    /// </summary>
    [EnumType]
    public readonly struct MethodIntegrationConnectionType : IEquatable<MethodIntegrationConnectionType>
    {
        private readonly string _value;

        private MethodIntegrationConnectionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MethodIntegrationConnectionType Internet { get; } = new MethodIntegrationConnectionType("INTERNET");
        public static MethodIntegrationConnectionType VpcLink { get; } = new MethodIntegrationConnectionType("VPC_LINK");

        public static bool operator ==(MethodIntegrationConnectionType left, MethodIntegrationConnectionType right) => left.Equals(right);
        public static bool operator !=(MethodIntegrationConnectionType left, MethodIntegrationConnectionType right) => !left.Equals(right);

        public static explicit operator string(MethodIntegrationConnectionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MethodIntegrationConnectionType other && Equals(other);
        public bool Equals(MethodIntegrationConnectionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies how to handle request payload content type conversions.
    /// </summary>
    [EnumType]
    public readonly struct MethodIntegrationContentHandling : IEquatable<MethodIntegrationContentHandling>
    {
        private readonly string _value;

        private MethodIntegrationContentHandling(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MethodIntegrationContentHandling ConvertToBinary { get; } = new MethodIntegrationContentHandling("CONVERT_TO_BINARY");
        public static MethodIntegrationContentHandling ConvertToText { get; } = new MethodIntegrationContentHandling("CONVERT_TO_TEXT");

        public static bool operator ==(MethodIntegrationContentHandling left, MethodIntegrationContentHandling right) => left.Equals(right);
        public static bool operator !=(MethodIntegrationContentHandling left, MethodIntegrationContentHandling right) => !left.Equals(right);

        public static explicit operator string(MethodIntegrationContentHandling value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MethodIntegrationContentHandling other && Equals(other);
        public bool Equals(MethodIntegrationContentHandling other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates when API Gateway passes requests to the targeted backend.
    /// </summary>
    [EnumType]
    public readonly struct MethodIntegrationPassthroughBehavior : IEquatable<MethodIntegrationPassthroughBehavior>
    {
        private readonly string _value;

        private MethodIntegrationPassthroughBehavior(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MethodIntegrationPassthroughBehavior WhenNoMatch { get; } = new MethodIntegrationPassthroughBehavior("WHEN_NO_MATCH");
        public static MethodIntegrationPassthroughBehavior WhenNoTemplates { get; } = new MethodIntegrationPassthroughBehavior("WHEN_NO_TEMPLATES");
        public static MethodIntegrationPassthroughBehavior Never { get; } = new MethodIntegrationPassthroughBehavior("NEVER");

        public static bool operator ==(MethodIntegrationPassthroughBehavior left, MethodIntegrationPassthroughBehavior right) => left.Equals(right);
        public static bool operator !=(MethodIntegrationPassthroughBehavior left, MethodIntegrationPassthroughBehavior right) => !left.Equals(right);

        public static explicit operator string(MethodIntegrationPassthroughBehavior value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MethodIntegrationPassthroughBehavior other && Equals(other);
        public bool Equals(MethodIntegrationPassthroughBehavior other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies how to handle request payload content type conversions.
    /// </summary>
    [EnumType]
    public readonly struct MethodIntegrationResponseContentHandling : IEquatable<MethodIntegrationResponseContentHandling>
    {
        private readonly string _value;

        private MethodIntegrationResponseContentHandling(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MethodIntegrationResponseContentHandling ConvertToBinary { get; } = new MethodIntegrationResponseContentHandling("CONVERT_TO_BINARY");
        public static MethodIntegrationResponseContentHandling ConvertToText { get; } = new MethodIntegrationResponseContentHandling("CONVERT_TO_TEXT");

        public static bool operator ==(MethodIntegrationResponseContentHandling left, MethodIntegrationResponseContentHandling right) => left.Equals(right);
        public static bool operator !=(MethodIntegrationResponseContentHandling left, MethodIntegrationResponseContentHandling right) => !left.Equals(right);

        public static explicit operator string(MethodIntegrationResponseContentHandling value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MethodIntegrationResponseContentHandling other && Equals(other);
        public bool Equals(MethodIntegrationResponseContentHandling other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of backend that your method is running.
    /// </summary>
    [EnumType]
    public readonly struct MethodIntegrationType : IEquatable<MethodIntegrationType>
    {
        private readonly string _value;

        private MethodIntegrationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MethodIntegrationType Aws { get; } = new MethodIntegrationType("AWS");
        public static MethodIntegrationType AwsProxy { get; } = new MethodIntegrationType("AWS_PROXY");
        public static MethodIntegrationType Http { get; } = new MethodIntegrationType("HTTP");
        public static MethodIntegrationType HttpProxy { get; } = new MethodIntegrationType("HTTP_PROXY");
        public static MethodIntegrationType Mock { get; } = new MethodIntegrationType("MOCK");

        public static bool operator ==(MethodIntegrationType left, MethodIntegrationType right) => left.Equals(right);
        public static bool operator !=(MethodIntegrationType left, MethodIntegrationType right) => !left.Equals(right);

        public static explicit operator string(MethodIntegrationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MethodIntegrationType other && Equals(other);
        public bool Equals(MethodIntegrationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of usage plan key. Currently, the only valid key type is API_KEY.
    /// </summary>
    [EnumType]
    public readonly struct UsagePlanKeyKeyType : IEquatable<UsagePlanKeyKeyType>
    {
        private readonly string _value;

        private UsagePlanKeyKeyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static UsagePlanKeyKeyType ApiKey { get; } = new UsagePlanKeyKeyType("API_KEY");

        public static bool operator ==(UsagePlanKeyKeyType left, UsagePlanKeyKeyType right) => left.Equals(right);
        public static bool operator !=(UsagePlanKeyKeyType left, UsagePlanKeyKeyType right) => !left.Equals(right);

        public static explicit operator string(UsagePlanKeyKeyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UsagePlanKeyKeyType other && Equals(other);
        public bool Equals(UsagePlanKeyKeyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}