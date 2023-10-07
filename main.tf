terraform {
  required_providers {
    terratowns = {
      source = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
#  cloud {
#  organization = "Luu"
#
#   workspaces {
#    name = "terra-house-acolyteluu"
# }
#}
}

provider "terratowns" {
  endpoint = "http://localhost:4567/api"
  user_uuid="00522451-153b-4131-ab48-43e132f662cc" 
  token="9b49b3fb-b8e9-483c-b703-97ba88eef8e0"
}

#module "terrahouse_aws" {
#  source = "./modules/terrahouse_aws"
#  user_uuid = var.user_uuid
#  bucket_name = var.bucket_name
#  index_html_filepath = var.index_html_filepath
#  error_html_filepath = var.error_html_filepath
#  content_version = var.content_version
#  assets_path = var.assets_path
#}

resource "terratowns_home" "home" {
  name = "How to end the long night"
  description = <<DESCRIPTION
Bloodborne is the best FromSoft game. We need a remastered one.
I am going to show you how to end the long night.
DESCRIPTION
  #domain_name = module.terrahouse_aws.cloudfront_url
  domain_name = "3bfdc.cloudfront.net"
  town = "gamers-grotto"
  content_version = 1
}