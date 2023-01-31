job "chroot_docker" {
  datacenters = ["dc1"]
  type        = "batch"

  constraint {
    attribute = "${attr.kernel.name}"
    value     = "linux"
  }

  group "print" {
    task "env" {
      driver = "docker"
      config {
        image = "busybox:1"
        args = [
          "/bin/sh",
          "-c",
          "echo $NOMAD_ALLOC_DIR; echo $NOMAD_TASK_DIR; echo $NOMAD_SECRETS_DIR; echo $PATH"
        ]
      }
      resources {
        cpu    = 50
        memory = 32
      }
    }
  }
}