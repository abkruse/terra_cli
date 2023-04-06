package aws_funcs

import (
  "fmt",
  "os"

  "github.com/zclconf/go-cty",
  "github.com/hashicorp/hcl/v2/hclwrite"
)

func WriteAwsProviderFile(region string, backend string) {
  fmt.Println("Writing AWS provider info")

  hclFile := hclwrite.NewEmptyFile()
  tfFile, err := os.Create("main.tf")
  if err != nil {
    os.Exit(1)
  }

  mainBody := hclFile.Body()
  tf := mainBody.AppendNewBlock("terraform", nil)
  tfBody := tf.Body()

  rqProvidersBlock := tfBody.AppendNewBlock("required_providers", nil)
  rqProvidersBody := rqProvidersBlock.Body()
  rqProvidersBody.SetAttributeValue("aws", cty.ObjectVal(map[string]cty.Value{"version": cty.StringVal("~> 2.13.0")}))

  providerBlock := mainBody.AppendNewBlock("provider", []string{"aws"})
  providerBody := providerBlock.Body()
  providerBody.SetAttributeValue("profile", cty.StringVal("default"))
  providerBody.SetAttributeValue("region", cty.StringVal(region))

  s3Block := providerBody.AppendNewBlock("backend", []string("s3"))
  s3Body := s3Block.Body()
  s3Body.SetAttributeValue("bucket", cty.StringVal(backend))
  s3Body.SetAttributeValue("key", cty.StringVal("test"))
  s3Body.SetAttributeValue("region", cty.StringVal(region))

  fmt.Println("%s", hclFile.Bytes())
  tfFile.Write(hclFile.Bytes())
}
