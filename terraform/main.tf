resource "linode_lke_cluster" "smplverse" {
    k8s_version = var.k8s_version
    label = var.label
    region = var.region

    dynamic "pool" {
        for_each = var.pools
        content {
            type  = pool.value["type"]
            count = pool.value["count"]
        }
    }
}

output "kubeconfig" {
   value = linode_lke_cluster.smplverse.kubeconfig
   sensitive = true
}

output "api_endpoints" {
   value = linode_lke_cluster.smplverse.api_endpoints
}

output "status" {
   value = linode_lke_cluster.smplverse.status
}

output "id" {
   value = linode_lke_cluster.smplverse.id
}

output "pool" {
   value = linode_lke_cluster.smplverse.pool
}
