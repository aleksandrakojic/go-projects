package main

import (
 eks2 "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
 "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
 "github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ec2"
 "github.com/pulumi/pulumi-eks/sdk/v3/go/eks"
 "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
 "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
 pulumi.Run(func(ctx *pulumi.Context) error {
  clusterName := "DevCluster"
  value := eks.ResolveConflictsOnCreate("OVERWRITE")
  // Get some configuration values or set default values
  cfg := config.New(ctx, "")
  minClusterSize, err := cfg.TryInt("minClusterSize")
  if err != nil {
   minClusterSize = 1
  }
  maxClusterSize, err := cfg.TryInt("maxClusterSize")
  if err != nil {
   maxClusterSize = 3
  }
  desiredClusterSize, err := cfg.TryInt("desiredClusterSize")
  if err != nil {
   desiredClusterSize = 1
  }
  eksNodeInstanceType, err := cfg.Try("eksNodeInstanceType")
  if err != nil {
   eksNodeInstanceType = "t3a.medium"
  }
  vpcNetworkCidr, err := cfg.Try("vpcNetworkCidr")
  if err != nil {
   vpcNetworkCidr = "10.0.0.0/16"
  }

  // Create a new VPC, subnets, and associated infrastructure
  eksVpc, err := ec2.NewVpc(ctx, "eks-vpc", &ec2.VpcArgs{
   EnableDnsHostnames: pulumi.Bool(true),
   CidrBlock:          &vpcNetworkCidr,
  })
  if err != nil {
   return err
  }

  apiAuthMode := eks.AuthenticationModeApi
  // Create a new EKS cluster
  eksCluster, err := eks.NewCluster(ctx, "eks-cluster", &eks.ClusterArgs{
   Name: pulumi.String(clusterName),
   // Put the cluster in the new VPC created earlier
   VpcId: eksVpc.VpcId,
   // Use the "API" authentication mode to support access entries
   AuthenticationMode: &apiAuthMode,
   // Public subnets will be used for load balancers
   PublicSubnetIds: eksVpc.PublicSubnetIds,
   // Private subnets will be used for cluster nodes
   PrivateSubnetIds: eksVpc.PrivateSubnetIds,
   // Change configuration values above to change any of the following settings
   //InstanceType:    pulumi.String(eksNodeInstanceType),
   //DesiredCapacity: pulumi.Int(desiredClusterSize),
   //MinSize:         pulumi.Int(minClusterSize),
   //MaxSize:         pulumi.Int(maxClusterSize),
   // Do not give the worker nodes a public IP address
   NodeAssociatePublicIpAddress: pulumi.BoolRef(false),
   // Change these values for a private cluster (VPN access required)
   EndpointPrivateAccess: pulumi.Bool(true),
   EndpointPublicAccess:  pulumi.Bool(true),
   SkipDefaultNodeGroup:  pulumi.BoolRef(true),
   CorednsAddonOptions: &eks.CoreDnsAddonOptionsArgs{
    Version:                  pulumi.String("v1.11.3-eksbuild.2"),
    Enabled:                  pulumi.BoolRef(true),
    ResolveConflictsOnCreate: &value,
   },

   //CreateOidcProvider:    pulumi.Bool(true),

  })
  if err != nil {
   return err
  }
  tagSubnets(ctx, eksVpc.PublicSubnetIds, "kubernetes.io/role/elb", "1")

  tagSubnets(ctx, eksVpc.PrivateSubnetIds, "kubernetes.io/role/internal-elb", "1")
  // Create an IAM role for the managed node group
  nodeGroupRole, err := iam.NewRole(ctx, "eks-nodegroup-role", &iam.RoleArgs{
   AssumeRolePolicy: pulumi.String(`{
    "Version": "2012-10-17",
    "Statement": [
     {
      "Effect": "Allow",
      "Principal": {
       "Service": "ec2.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
     }
    ]
   }`),
  })
  if err != nil {
   return err
  }

  _, err = iam.NewRolePolicyAttachment(ctx, "eks-nodegroup-policy-attachment", &iam.RolePolicyAttachmentArgs{
   Role:      nodeGroupRole.Name,
   PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"),
  })
  if err != nil {
   return err
  }

  _, err = iam.NewRolePolicyAttachment(ctx, "eks-cni-policy-attachment", &iam.RolePolicyAttachmentArgs{
   Role:      nodeGroupRole.Name,
   PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"),
  })
  if err != nil {
   return err
  }

  _, err = iam.NewRolePolicyAttachment(ctx, "eks-registry-policy-attachment", &iam.RolePolicyAttachmentArgs{
   Role:      nodeGroupRole.Name,
   PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"),
  })
  if err != nil {
   return err
  }

  // Create a new Node Group
  nodeGroup, err := eks.NewManagedNodeGroup(ctx, "eks-nodegroup", &eks.ManagedNodeGroupArgs{
   Cluster:       eksCluster,
   NodeGroupName: pulumi.String("dev-ng"),
   InstanceTypes: pulumi.StringArray{
    pulumi.String(eksNodeInstanceType),
   },
   DiskSize: pulumi.Int(50),
   ScalingConfig: eks2.NodeGroupScalingConfigArgs{
    MinSize:     pulumi.Int(minClusterSize),
    MaxSize:     pulumi.Int(maxClusterSize),
    DesiredSize: pulumi.Int(desiredClusterSize),
   },
   NodeRoleArn:  nodeGroupRole.Arn,
   CapacityType: pulumi.String("SPOT"),
   AmiId:        pulumi.String("ami-02c0c54a3a6c62ac1"),
  })
  if err != nil {
   return err
  }
  // Export some values in case they are needed elsewhere
  ctx.Export("kubeconfig", eksCluster.Kubeconfig)
  ctx.Export("vpcId", eksVpc.VpcId)
  ctx.Export("eksNodeGroupName", nodeGroup.NodeGroup.NodeGroupName())
  return nil
 })
}