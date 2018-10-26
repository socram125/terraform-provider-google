// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeAddress_addressBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_addressBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_address.ip_address",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeAddress_addressBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_address" "ip_address" {
  name = "my-address-%s"
}
`, val,
	)
}

func TestAccComputeAddress_addressWithSubnetworkExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_addressWithSubnetworkExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_address.internal_with_subnet_and_address",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeAddress_addressWithSubnetworkExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "default" {
  name = "my-network-%s"
}

resource "google_compute_subnetwork" "default" {
  name          = "my-subnet-%s"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = "${google_compute_network.default.self_link}"
}

resource "google_compute_address" "internal_with_subnet_and_address" {
  name         = "my-internal-address-%s"
  subnetwork   = "${google_compute_subnetwork.default.self_link}"
  address_type = "INTERNAL"
  address      = "10.0.42.42"
  region       = "us-central1"
}
`, val, val, val,
	)
}

func TestAccComputeAddress_instanceWithIpExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAddress_instanceWithIpExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_address.static",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeAddress_instanceWithIpExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_address" "static" {
  name = "ipv4-address-%s"
}

data "google_compute_image" "debian_image" {
	family  = "debian-9"
	project = "debian-cloud"
}

resource "google_compute_instance" "instance_with_ip" {
	name         = "vm-instance-%s"
	machine_type = "f1-micro"
	zone         = "us-central1-a"

	boot_disk {
		initialize_params{
			image = "${data.google_compute_image.debian_image.self_link}"
		}
	}

	network_interface {
		network = "default"
		access_config {
			nat_ip = "${google_compute_address.static.address}"
		}
	}
}
`, val, val,
	)
}

func testAccCheckComputeAddressDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_address" {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/addresses/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeAddress still exists at %s", url)
		}
	}

	return nil
}
