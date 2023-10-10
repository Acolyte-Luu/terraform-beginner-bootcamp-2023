output "bucket_name" {
    description = "Static website bucket name"
    value = module.home_bloodborne_hosting.bucket_name
}

output "s3_website_endpoint" {
  description = "s3 static website hosting endpoint"
  value = module.home_bloodborne_hosting.website_endpoint
}

output "domain_name" {
  description = "Cloudfront distribution domain name"
  value = module.home_bloodborne_hosting.domain_name
}