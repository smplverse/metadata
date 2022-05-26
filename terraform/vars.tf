variable "LINODE_TOKEN" {}

variable "k8s_version" {}

variable "label" {}

variable "region" {}

variable "pools" {
    type = list(object({
        type = string
        count = number
    }))
}
