import "plugin" "plugin-demo" {
  source = "./bin/sentinel-plugin-demo"
}

policy "demo-plugin-demo" {
  source            = "./policies/demo-plugin.sentinel"
  enforcement_level = "advisory"
}