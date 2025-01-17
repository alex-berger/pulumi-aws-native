// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:opensearchservice:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    public readonly accessPolicies!: pulumi.Output<any | undefined>;
    public readonly advancedOptions!: pulumi.Output<any | undefined>;
    public readonly advancedSecurityOptions!: pulumi.Output<outputs.opensearchservice.DomainAdvancedSecurityOptionsInput | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly clusterConfig!: pulumi.Output<outputs.opensearchservice.DomainClusterConfig | undefined>;
    public readonly cognitoOptions!: pulumi.Output<outputs.opensearchservice.DomainCognitoOptions | undefined>;
    public /*out*/ readonly domainArn!: pulumi.Output<string>;
    public /*out*/ readonly domainEndpoint!: pulumi.Output<string>;
    public readonly domainEndpointOptions!: pulumi.Output<outputs.opensearchservice.DomainEndpointOptions | undefined>;
    public /*out*/ readonly domainEndpoints!: pulumi.Output<any>;
    public readonly domainName!: pulumi.Output<string | undefined>;
    public readonly eBSOptions!: pulumi.Output<outputs.opensearchservice.DomainEBSOptions | undefined>;
    public readonly encryptionAtRestOptions!: pulumi.Output<outputs.opensearchservice.DomainEncryptionAtRestOptions | undefined>;
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    public readonly logPublishingOptions!: pulumi.Output<any | undefined>;
    public readonly nodeToNodeEncryptionOptions!: pulumi.Output<outputs.opensearchservice.DomainNodeToNodeEncryptionOptions | undefined>;
    public /*out*/ readonly serviceSoftwareOptions!: pulumi.Output<outputs.opensearchservice.DomainServiceSoftwareOptions>;
    public readonly snapshotOptions!: pulumi.Output<outputs.opensearchservice.DomainSnapshotOptions | undefined>;
    /**
     * An arbitrary set of tags (key-value pairs) for this Domain.
     */
    public readonly tags!: pulumi.Output<outputs.opensearchservice.DomainTag[] | undefined>;
    public readonly vPCOptions!: pulumi.Output<outputs.opensearchservice.DomainVPCOptions | undefined>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DomainArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["accessPolicies"] = args ? args.accessPolicies : undefined;
            resourceInputs["advancedOptions"] = args ? args.advancedOptions : undefined;
            resourceInputs["advancedSecurityOptions"] = args ? args.advancedSecurityOptions : undefined;
            resourceInputs["clusterConfig"] = args ? args.clusterConfig : undefined;
            resourceInputs["cognitoOptions"] = args ? args.cognitoOptions : undefined;
            resourceInputs["domainEndpointOptions"] = args ? args.domainEndpointOptions : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["eBSOptions"] = args ? args.eBSOptions : undefined;
            resourceInputs["encryptionAtRestOptions"] = args ? args.encryptionAtRestOptions : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["logPublishingOptions"] = args ? args.logPublishingOptions : undefined;
            resourceInputs["nodeToNodeEncryptionOptions"] = args ? args.nodeToNodeEncryptionOptions : undefined;
            resourceInputs["snapshotOptions"] = args ? args.snapshotOptions : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vPCOptions"] = args ? args.vPCOptions : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["domainArn"] = undefined /*out*/;
            resourceInputs["domainEndpoint"] = undefined /*out*/;
            resourceInputs["domainEndpoints"] = undefined /*out*/;
            resourceInputs["serviceSoftwareOptions"] = undefined /*out*/;
        } else {
            resourceInputs["accessPolicies"] = undefined /*out*/;
            resourceInputs["advancedOptions"] = undefined /*out*/;
            resourceInputs["advancedSecurityOptions"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["clusterConfig"] = undefined /*out*/;
            resourceInputs["cognitoOptions"] = undefined /*out*/;
            resourceInputs["domainArn"] = undefined /*out*/;
            resourceInputs["domainEndpoint"] = undefined /*out*/;
            resourceInputs["domainEndpointOptions"] = undefined /*out*/;
            resourceInputs["domainEndpoints"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["eBSOptions"] = undefined /*out*/;
            resourceInputs["encryptionAtRestOptions"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["logPublishingOptions"] = undefined /*out*/;
            resourceInputs["nodeToNodeEncryptionOptions"] = undefined /*out*/;
            resourceInputs["serviceSoftwareOptions"] = undefined /*out*/;
            resourceInputs["snapshotOptions"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vPCOptions"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    accessPolicies?: any;
    advancedOptions?: any;
    advancedSecurityOptions?: pulumi.Input<inputs.opensearchservice.DomainAdvancedSecurityOptionsInputArgs>;
    clusterConfig?: pulumi.Input<inputs.opensearchservice.DomainClusterConfigArgs>;
    cognitoOptions?: pulumi.Input<inputs.opensearchservice.DomainCognitoOptionsArgs>;
    domainEndpointOptions?: pulumi.Input<inputs.opensearchservice.DomainEndpointOptionsArgs>;
    domainName?: pulumi.Input<string>;
    eBSOptions?: pulumi.Input<inputs.opensearchservice.DomainEBSOptionsArgs>;
    encryptionAtRestOptions?: pulumi.Input<inputs.opensearchservice.DomainEncryptionAtRestOptionsArgs>;
    engineVersion?: pulumi.Input<string>;
    logPublishingOptions?: any;
    nodeToNodeEncryptionOptions?: pulumi.Input<inputs.opensearchservice.DomainNodeToNodeEncryptionOptionsArgs>;
    snapshotOptions?: pulumi.Input<inputs.opensearchservice.DomainSnapshotOptionsArgs>;
    /**
     * An arbitrary set of tags (key-value pairs) for this Domain.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.opensearchservice.DomainTagArgs>[]>;
    vPCOptions?: pulumi.Input<inputs.opensearchservice.DomainVPCOptionsArgs>;
}
