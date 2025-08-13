import "plugin" "plugin-demo" {
  source = "./bin/sentinel-plugin-demo"
}

policy "demo-plugin-demo" {
  source            = "./policies/plugin-demo.sentinel"
  enforcement_level = "advisory"
}