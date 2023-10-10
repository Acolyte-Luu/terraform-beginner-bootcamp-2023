terraform {
  required_providers {
    terratowns = {
      source = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
  cloud {
  organization = "Luu"

   workspaces {
    name = "terra-house-acolyteluu"
 }
}
}

provider "terratowns" {
  endpoint = var.terratowns_endpoint
  user_uuid = var.teacherseat_user_uuid 
  token = var.terratowns_access_token
}

module "home_bloodborne_hosting" {
  source = "./modules/terrahome_aws"
  user_uuid = var.teacherseat_user_uuid
  public_path = var.bloodborne.public_path
  content_version = var.bloodborne.content_version
}

resource "terratowns_home" "home_bloodborne" {
  name = "How to end the long night"
  description = <<DESCRIPTION
Bloodborne is the best FromSoft game. We need a remastered one.
I am going to show you how to end the long night.
DESCRIPTION
  domain_name = module.home_bloodborne_hosting.domain_name
  #domain_name = "3dfgffg.cloudfront.net"
  town = "missingo"
  content_version = var.bloodborne.content_version
}

module "home_elden_ring_hosting" {
  source = "./modules/terrahome_aws"
  user_uuid = var.teacherseat_user_uuid
  public_path = var.elden_ring.public_path
  content_version = var.elden_ring.content_version
}

resource "terratowns_home" "home_elden_ring" {
  name = "How to become the Elden Lord"
  description = <<DESCRIPTION
Tarnished! Are you ready to face the demi-gods and becom the rightful Elden Lord?
Pick up your sword and find grace in the Lands Between.
DESCRIPTION
  domain_name = module.home_elden_ring_hosting.domain_name
  #domain_name = "3elde4ing.cloudfront.net"
  town = "missingo"
  content_version = var.elden_ring.content_version
}
