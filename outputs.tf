output "bucket_name" {
    description = "Static website bucket name"
    value = module.terrahouse_aws.bucket_name
}

output "s3_website_endpoint" {
  description = "s3 static website hosting endpoint"
  value = module.terrahouse_aws.website_endpoint
}

locals {
  root_path = path.root
  module_path = path.module
}

output "root_path" {
  value = local.root_path
}

output "module_path" {
  value = local.root_path
}