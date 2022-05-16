// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElastiCache::User
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:elasticache:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Access permissions string used for this user account.
     */
    public readonly accessString!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the user account.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Must be redis.
     */
    public readonly engine!: pulumi.Output<enums.elasticache.UserEngine>;
    /**
     * Indicates a password is not required for this user account.
     */
    public readonly noPasswordRequired!: pulumi.Output<boolean | undefined>;
    /**
     * Passwords used for this user account. You can create up to two passwords for each user.
     */
    public readonly passwords!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates the user status. Can be "active", "modifying" or "deleting".
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the user.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * The username of the user.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["accessString"] = args ? args.accessString : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["noPasswordRequired"] = args ? args.noPasswordRequired : undefined;
            resourceInputs["passwords"] = args ? args.passwords : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["accessString"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["engine"] = undefined /*out*/;
            resourceInputs["noPasswordRequired"] = undefined /*out*/;
            resourceInputs["passwords"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
            resourceInputs["userName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Access permissions string used for this user account.
     */
    accessString?: pulumi.Input<string>;
    /**
     * Must be redis.
     */
    engine: pulumi.Input<enums.elasticache.UserEngine>;
    /**
     * Indicates a password is not required for this user account.
     */
    noPasswordRequired?: pulumi.Input<boolean>;
    /**
     * Passwords used for this user account. You can create up to two passwords for each user.
     */
    passwords?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the user.
     */
    userId: pulumi.Input<string>;
    /**
     * The username of the user.
     */
    userName?: pulumi.Input<string>;
}
