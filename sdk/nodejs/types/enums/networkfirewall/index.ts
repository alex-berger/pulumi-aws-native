// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const FirewallPolicyRuleOrder = {
    DefaultActionOrder: "DEFAULT_ACTION_ORDER",
    StrictOrder: "STRICT_ORDER",
} as const;

export type FirewallPolicyRuleOrder = (typeof FirewallPolicyRuleOrder)[keyof typeof FirewallPolicyRuleOrder];

export const LoggingConfigurationLogDestinationConfigLogDestinationType = {
    S3: "S3",
    CloudWatchLogs: "CloudWatchLogs",
    KinesisDataFirehose: "KinesisDataFirehose",
} as const;

export type LoggingConfigurationLogDestinationConfigLogDestinationType = (typeof LoggingConfigurationLogDestinationConfigLogDestinationType)[keyof typeof LoggingConfigurationLogDestinationConfigLogDestinationType];

export const LoggingConfigurationLogDestinationConfigLogType = {
    Alert: "ALERT",
    Flow: "FLOW",
} as const;

export type LoggingConfigurationLogDestinationConfigLogType = (typeof LoggingConfigurationLogDestinationConfigLogType)[keyof typeof LoggingConfigurationLogDestinationConfigLogType];

export const RuleGroupGeneratedRulesType = {
    Allowlist: "ALLOWLIST",
    Denylist: "DENYLIST",
} as const;

export type RuleGroupGeneratedRulesType = (typeof RuleGroupGeneratedRulesType)[keyof typeof RuleGroupGeneratedRulesType];

export const RuleGroupHeaderDirection = {
    Forward: "FORWARD",
    Any: "ANY",
} as const;

export type RuleGroupHeaderDirection = (typeof RuleGroupHeaderDirection)[keyof typeof RuleGroupHeaderDirection];

export const RuleGroupHeaderProtocol = {
    Ip: "IP",
    Tcp: "TCP",
    Udp: "UDP",
    Icmp: "ICMP",
    Http: "HTTP",
    Ftp: "FTP",
    Tls: "TLS",
    Smb: "SMB",
    Dns: "DNS",
    Dcerpc: "DCERPC",
    Ssh: "SSH",
    Smtp: "SMTP",
    Imap: "IMAP",
    Msn: "MSN",
    Krb5: "KRB5",
    Ikev2: "IKEV2",
    Tftp: "TFTP",
    Ntp: "NTP",
    Dhcp: "DHCP",
} as const;

export type RuleGroupHeaderProtocol = (typeof RuleGroupHeaderProtocol)[keyof typeof RuleGroupHeaderProtocol];

export const RuleGroupRuleOrder = {
    DefaultActionOrder: "DEFAULT_ACTION_ORDER",
    StrictOrder: "STRICT_ORDER",
} as const;

export type RuleGroupRuleOrder = (typeof RuleGroupRuleOrder)[keyof typeof RuleGroupRuleOrder];

export const RuleGroupStatefulRuleAction = {
    Pass: "PASS",
    Drop: "DROP",
    Alert: "ALERT",
} as const;

export type RuleGroupStatefulRuleAction = (typeof RuleGroupStatefulRuleAction)[keyof typeof RuleGroupStatefulRuleAction];

export const RuleGroupTCPFlag = {
    Fin: "FIN",
    Syn: "SYN",
    Rst: "RST",
    Psh: "PSH",
    Ack: "ACK",
    Urg: "URG",
    Ece: "ECE",
    Cwr: "CWR",
} as const;

export type RuleGroupTCPFlag = (typeof RuleGroupTCPFlag)[keyof typeof RuleGroupTCPFlag];

export const RuleGroupTargetType = {
    TlsSni: "TLS_SNI",
    HttpHost: "HTTP_HOST",
} as const;

export type RuleGroupTargetType = (typeof RuleGroupTargetType)[keyof typeof RuleGroupTargetType];

export const RuleGroupTypeEnum = {
    Stateless: "STATELESS",
    Stateful: "STATEFUL",
} as const;

export type RuleGroupTypeEnum = (typeof RuleGroupTypeEnum)[keyof typeof RuleGroupTypeEnum];
