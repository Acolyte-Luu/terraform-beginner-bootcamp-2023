variable "teacherseat_user_uuid" {
    type = string
}

#variable "bucket_name" {
#    type = string
#}

#variable "index_html_filepath" {
#    type = string
#}

#variable "error_html_filepath" {
#    type = string
#}

variable "bloodborne" {
  type = object({
    public_path = string
    content_version = number
  })
}

variable "elden_ring" {
  type = object({
    public_path = string
    content_version = number
  })
}

#variable "assets_path" {
#  type = string
#}

variable "terratowns_endpoint" {
    type = string 
}

variable "terratowns_access_token" {
  type = string
}
