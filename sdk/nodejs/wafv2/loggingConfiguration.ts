// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * A WAFv2 Logging Configuration Resource Provider
 */
export class LoggingConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing LoggingConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LoggingConfiguration {
        return new LoggingConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wafv2:LoggingConfiguration';

    /**
     * Returns true if the given object is an instance of LoggingConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoggingConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoggingConfiguration.__pulumiType;
    }

    /**
     * The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL.
     */
    public readonly logDestinationConfigs!: pulumi.Output<string[]>;
    /**
     * Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
     */
    public readonly loggingFilter!: pulumi.Output<any | undefined>;
    /**
     * Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.
     */
    public /*out*/ readonly managedByFirewallManager!: pulumi.Output<boolean>;
    /**
     * The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
     */
    public readonly redactedFields!: pulumi.Output<outputs.wafv2.LoggingConfigurationFieldToMatch[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
     */
    public readonly resourceArn!: pulumi.Output<string>;

    /**
     * Create a LoggingConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoggingConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.logDestinationConfigs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logDestinationConfigs'");
            }
            if ((!args || args.resourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceArn'");
            }
            inputs["logDestinationConfigs"] = args ? args.logDestinationConfigs : undefined;
            inputs["loggingFilter"] = args ? args.loggingFilter : undefined;
            inputs["redactedFields"] = args ? args.redactedFields : undefined;
            inputs["resourceArn"] = args ? args.resourceArn : undefined;
            inputs["managedByFirewallManager"] = undefined /*out*/;
        } else {
            inputs["logDestinationConfigs"] = undefined /*out*/;
            inputs["loggingFilter"] = undefined /*out*/;
            inputs["managedByFirewallManager"] = undefined /*out*/;
            inputs["redactedFields"] = undefined /*out*/;
            inputs["resourceArn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LoggingConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LoggingConfiguration resource.
 */
export interface LoggingConfigurationArgs {
    /**
     * The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL.
     */
    logDestinationConfigs: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
     */
    loggingFilter?: any;
    /**
     * The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
     */
    redactedFields?: pulumi.Input<pulumi.Input<inputs.wafv2.LoggingConfigurationFieldToMatchArgs>[]>;
    /**
     * The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
     */
    resourceArn: pulumi.Input<string>;
}