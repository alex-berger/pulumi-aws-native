// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::KMS::ReplicaKey resource specifies a multi-region replica customer master key (CMK) in AWS Key Management Service (AWS KMS).
 */
export class ReplicaKey extends pulumi.CustomResource {
    /**
     * Get an existing ReplicaKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReplicaKey {
        return new ReplicaKey(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:kms:ReplicaKey';

    /**
     * Returns true if the given object is an instance of ReplicaKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicaKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicaKey.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly keyId!: pulumi.Output<string>;
    /**
     * The key policy that authorizes use of the CMK. The key policy must observe the following rules.
     */
    public readonly keyPolicy!: pulumi.Output<any>;
    /**
     * Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
     */
    public readonly pendingWindowInDays!: pulumi.Output<number | undefined>;
    /**
     * Identifies the primary CMK to create a replica of. Specify the Amazon Resource Name (ARN) of the CMK. You cannot specify an alias or key ID. For help finding the ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.
     */
    public readonly primaryKeyArn!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.kms.ReplicaKeyTag[] | undefined>;

    /**
     * Create a ReplicaKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicaKeyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.keyPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyPolicy'");
            }
            if ((!args || args.primaryKeyArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'primaryKeyArn'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["keyPolicy"] = args ? args.keyPolicy : undefined;
            resourceInputs["pendingWindowInDays"] = args ? args.pendingWindowInDays : undefined;
            resourceInputs["primaryKeyArn"] = args ? args.primaryKeyArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["keyId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["keyId"] = undefined /*out*/;
            resourceInputs["keyPolicy"] = undefined /*out*/;
            resourceInputs["pendingWindowInDays"] = undefined /*out*/;
            resourceInputs["primaryKeyArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReplicaKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReplicaKey resource.
 */
export interface ReplicaKeyArgs {
    /**
     * A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The key policy that authorizes use of the CMK. The key policy must observe the following rules.
     */
    keyPolicy: any;
    /**
     * Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
     */
    pendingWindowInDays?: pulumi.Input<number>;
    /**
     * Identifies the primary CMK to create a replica of. Specify the Amazon Resource Name (ARN) of the CMK. You cannot specify an alias or key ID. For help finding the ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.
     */
    primaryKeyArn: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.kms.ReplicaKeyTagArgs>[]>;
}
