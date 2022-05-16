// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Location::RouteCalculator Resource Type
 */
export class RouteCalculator extends pulumi.CustomResource {
    /**
     * Get an existing RouteCalculator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RouteCalculator {
        return new RouteCalculator(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:location:RouteCalculator';

    /**
     * Returns true if the given object is an instance of RouteCalculator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteCalculator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteCalculator.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly calculatorArn!: pulumi.Output<string>;
    public readonly calculatorName!: pulumi.Output<string>;
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    public readonly dataSource!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly pricingPlan!: pulumi.Output<enums.location.RouteCalculatorPricingPlan | undefined>;
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a RouteCalculator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteCalculatorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.calculatorName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'calculatorName'");
            }
            if ((!args || args.dataSource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataSource'");
            }
            resourceInputs["calculatorName"] = args ? args.calculatorName : undefined;
            resourceInputs["dataSource"] = args ? args.dataSource : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["pricingPlan"] = args ? args.pricingPlan : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["calculatorArn"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["calculatorArn"] = undefined /*out*/;
            resourceInputs["calculatorName"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dataSource"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["pricingPlan"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouteCalculator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RouteCalculator resource.
 */
export interface RouteCalculatorArgs {
    calculatorName: pulumi.Input<string>;
    dataSource: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    pricingPlan?: pulumi.Input<enums.location.RouteCalculatorPricingPlan>;
}
