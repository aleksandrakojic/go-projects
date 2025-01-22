package main

import (
    "fmt"

    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func tagSubnets(ctx *pulumi.Context, subnetIds pulumi.StringArrayOutput, tagKey, tagValue string) {
    // use ApplyT get subnet ID and tag
    subnetIds.ApplyT(func(ids []string) []string {
       for _, subnetId := range ids {
          resourceName := fmt.Sprintf("subnet-tag-%s", subnetId)
          ctx.Log.Info(fmt.Sprintf("Tagging subnet: %s with %s=%s", subnetId, tagKey, tagValue), nil)

          // create tag
          _, err := ec2.NewTag(ctx, resourceName, &ec2.TagArgs{
             ResourceId: pulumi.String(subnetId),
             Key:        pulumi.String(tagKey),
             Value:      pulumi.String(tagValue),
          })
          if err != nil {
             // remark error
             ctx.Log.Error(fmt.Sprintf("Failed to tag subnet %s: %v", subnetId, err), nil)
          }
       }
       return nil // return nil for ApplyT
    })
}