package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudALBServerGroup_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alb_server_group.default"
	ra := resourceAttrInit(resourceId, AlicloudALBServerGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AlbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAlbServerGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%salbservergroup%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudALBServerGroupBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, connectivity.AlbSupportRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"protocol":          "HTTP",
					"vpc_id":            "${data.alicloud_vpcs.default.vpcs.0.id}",
					"server_group_name": "${var.name}",
					"health_check_config": []map[string]interface{}{
						{
							"health_check_enabled": "false",
						},
					},
					"sticky_session_config": []map[string]interface{}{
						{
							"sticky_session_enabled": "false",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"protocol":                "HTTP",
						"server_group_name":       name,
						"sticky_session_config.#": "1",
						"health_check_config.#":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"server_group_name": "tf-testAcc-new",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"server_group_name": "tf-testAcc-new",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"scheduler": "Wlc",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"scheduler": "Wlc",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"health_check_config": []map[string]interface{}{
						{
							"health_check_connect_port": "46325",
							"health_check_enabled":      "true",
							"health_check_host":         "tf-testAcc.com",
							"health_check_codes":        []string{"http_2xx", "http_3xx", "http_4xx"},
							"health_check_http_version": "HTTP1.1",
							"health_check_interval":     "2",
							"health_check_method":       "HEAD",
							"health_check_path":         "/tf-testAcc",
							"health_check_protocol":     "HTTP",
							"health_check_timeout":      "5",
							"healthy_threshold":         "3",
							"unhealthy_threshold":       "3",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"health_check_config.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"protocol": "HTTPS",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"protocol": "HTTPS",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.groups.0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},

			{
				Config: testAccConfig(map[string]interface{}{
					"sticky_session_config": []map[string]interface{}{
						{
							"cookie_timeout":         "2000",
							"sticky_session_enabled": "true",
							"sticky_session_type":    "Insert",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sticky_session_config.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sticky_session_config": []map[string]interface{}{
						{
							"cookie":                 "tf-testAcc",
							"sticky_session_enabled": "true",
							"sticky_session_type":    "Server",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sticky_session_config.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "tfTestAcc7",
						"For":     "Tftestacc7",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "tfTestAcc7",
						"tags.For":     "Tftestacc7",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"server_group_name": "${var.name}",
					"scheduler":         "Wrr",
					"health_check_config": []map[string]interface{}{
						{
							"health_check_enabled": "false",
						},
					},
					"sticky_session_config": []map[string]interface{}{
						{
							"sticky_session_enabled": "false",
						},
					},
					"tags": map[string]string{
						"Created": "tfTestAcc99",
						"For":     "Tftestacc99",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"server_group_name":       name,
						"scheduler":               "Wrr",
						"health_check_config.#":   "1",
						"sticky_session_config.#": "1",
						"tags.%":                  "2",
						"tags.Created":            "tfTestAcc99",
						"tags.For":                "Tftestacc99",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{"dry_run"},
			},
		},
	})
}

func TestAccAlicloudALBServerGroup_basic1(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alb_server_group.default"
	ra := resourceAttrInit(resourceId, AlicloudALBServerGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AlbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAlbServerGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%salbservergroup%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudALBServerGroupBasicDependence1)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.AlbSupportRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"protocol":          "HTTP",
					"vpc_id":            "${alicloud_vpc.default.id}",
					"server_group_name": "${var.name}",
					"health_check_config": []map[string]interface{}{
						{
							"health_check_enabled": "false",
						},
					},
					"sticky_session_config": []map[string]interface{}{
						{
							"sticky_session_enabled": "false",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"protocol":                "HTTP",
						"server_group_name":       name,
						"sticky_session_config.#": "1",
						"health_check_config.#":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"servers": []map[string]interface{}{
						{
							"description": "tf-testAcc",
							"port":        "80",
							"server_id":   "${alicloud_instance.instance.id}",
							"server_ip":   "${alicloud_instance.instance.private_ip}",
							"server_type": "Ecs",
							"weight":      "10",
						},
						{
							"description": "tf-testAcc",
							"port":        "8080",
							"server_id":   "${alicloud_instance.instance.id}",
							"server_ip":   "${alicloud_instance.instance.private_ip}",
							"server_type": "Ecs",
							"weight":      "10",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"servers.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"servers": []map[string]interface{}{
						{
							"description": "tf-testAcc-update",
							"port":        "80",
							"server_id":   "${alicloud_instance.instance.id}",
							"server_ip":   "${alicloud_instance.instance.private_ip}",
							"server_type": "Ecs",
							"weight":      "10",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"servers.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"servers": []map[string]interface{}{
						{
							"description": "tf-testAcc-update",
							"port":        "80",
							"server_id":   "${alicloud_instance.instance.id}",
							"server_ip":   "${alicloud_instance.instance.private_ip}",
							"server_type": "Ecs",
							"weight":      "10",
						},
						{
							"description": "tf-testAcc-update-8056",
							"port":        "8056",
							"server_id":   "${alicloud_instance.instance.id}",
							"server_ip":   "${alicloud_instance.instance.private_ip}",
							"server_type": "Ecs",
							"weight":      "10",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"servers.#": "2",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{"dry_run"},
			},
		},
	})
}

var AlicloudALBServerGroupMap0 = map[string]string{
	"tags.%":            NOSET,
	"dry_run":           NOSET,
	"resource_group_id": CHECKSET,
	"status":            CHECKSET,
	"scheduler":         CHECKSET,
	"vpc_id":            CHECKSET,
}

func AlicloudALBServerGroupBasicDependence0(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}
data "alicloud_vpcs" "default" {
  name_regex = "default-NODELETING"
}
data "alicloud_resource_manager_resource_groups" "default" {}

`, name)

}

func AlicloudALBServerGroupBasicDependence1(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

data "alicloud_zones" "default" {
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
}

data "alicloud_instance_types" "default" {
  availability_zone = data.alicloud_zones.default.zones[0].id
  cpu_core_count    = 1
  memory_size       = 2
}

data "alicloud_images" "default" {
  name_regex  = "^ubuntu_18.*64"
  most_recent = true
  owners      = "system"
}

resource "alicloud_vpc" "default" {
  vpc_name       = var.name
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vpc_id            = alicloud_vpc.default.id
  cidr_block        = "172.16.0.0/16"
  zone_id           = data.alicloud_zones.default.zones[0].id
  vswitch_name              = var.name
}

resource "alicloud_security_group" "default" {
  name   = var.name
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_instance" "instance" {
  image_id                   = data.alicloud_images.default.images[0].id
  instance_type              = data.alicloud_instance_types.default.instance_types[0].id
  instance_name              = var.name
  security_groups            = alicloud_security_group.default.*.id
  internet_charge_type       = "PayByTraffic"
  internet_max_bandwidth_out = "10"
  availability_zone          = data.alicloud_zones.default.zones[0].id
  instance_charge_type       = "PostPaid"
  system_disk_category       = "cloud_efficiency"
  vswitch_id                 = alicloud_vswitch.default.id
}

`, name)

}
