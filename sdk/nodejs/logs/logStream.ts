// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Logs::LogStream
 *
 * @deprecated LogStream is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class LogStream extends pulumi.CustomResource {
    /**
     * Get an existing LogStream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LogStream {
        pulumi.log.warn("LogStream is deprecated: LogStream is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new LogStream(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:logs:LogStream';

    /**
     * Returns true if the given object is an instance of LogStream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogStream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogStream.__pulumiType;
    }

    public readonly logGroupName!: pulumi.Output<string>;
    public readonly logStreamName!: pulumi.Output<string | undefined>;

    /**
     * Create a LogStream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated LogStream is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: LogStreamArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LogStream is deprecated: LogStream is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.logGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logGroupName'");
            }
            resourceInputs["logGroupName"] = args ? args.logGroupName : undefined;
            resourceInputs["logStreamName"] = args ? args.logStreamName : undefined;
        } else {
            resourceInputs["logGroupName"] = undefined /*out*/;
            resourceInputs["logStreamName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogStream.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LogStream resource.
 */
export interface LogStreamArgs {
    logGroupName: pulumi.Input<string>;
    logStreamName?: pulumi.Input<string>;
}
