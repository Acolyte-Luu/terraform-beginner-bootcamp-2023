output "bucket_name" {
    description = "Static website bucket name"
    value = module.terrahouse_aws.bucket_name
}

output "s3_website_endpoint" {
  description = "s3 static website hosting endpoint"
  value = module.terrahouse_aws.website_endpoint
}