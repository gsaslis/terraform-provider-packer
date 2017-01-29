provider "packer" {
  work_dir = "/Users/giorgos/dev/GitHub/oogababy/chef-repo"
  json_path = "/Users/giorgos/dev/GitHub/oogababy/chef-repo/packer/amis.json"
}

resource "packer_build" "foo" {
  builder_name = "backend"
}
