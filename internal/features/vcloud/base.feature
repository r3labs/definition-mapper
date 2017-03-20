@vcloud @vcloud_base
Feature: Service apply

  Scenario: Applying a basic service
    Given I setup ernest with target "https://ernest.local"
    And I setup a new service name
    When I'm logged in as "usr" / "pwd"
    And I start recording
    And I apply the definition "vcloud1.yml"
    And I stop recording
    Then an event "router.create.vcloud-fake" should be called exactly "1" times
    And all "router.create.vcloud-fake" messages should contain a field "_provider" with "vcloud-fake"
    And all "router.create.vcloud-fake" messages should contain a field "name" with "vse2"
    And all "router.create.vcloud-fake" messages should contain a field "vcloud_url" with "https://vcloud.net"
    And all "router.create.vcloud-fake" messages should contain a field "datacenter_name" with "fakevcloud"
    And all "router.create.vcloud-fake" messages should contain an encrypted field "datacenter_username" with "fakeuser@test"
    And all "router.create.vcloud-fake" messages should contain an encrypted field "datacenter_password" with "test123"
    And message "router.create.vcloud-fake" number "0" should contain "in_in_any" as json field "firewall_rules.0.name"
    And message "router.create.vcloud-fake" number "0" should contain "office2_in_22" as json field "firewall_rules.1.name"
    And message "router.create.vcloud-fake" number "0" should contain "office1_in_22" as json field "firewall_rules.2.name"
    And message "router.create.vcloud-fake" number "0" should contain "home_in_22" as json field "firewall_rules.3.name"
    Then an event "network.create.vcloud-fake" should be called exactly "1" times
    And all "network.create.vcloud-fake" messages should contain a field "_provider" with "vcloud-fake"
    And all "network.create.vcloud-fake" messages should contain a field "range" with "10.1.0.0/24"
    And all "network.create.vcloud-fake" messages should contain a field "vcloud_url" with "https://vcloud.net"
    And all "network.create.vcloud-fake" messages should contain a field "datacenter_name" with "fakevcloud"
    And all "network.create.vcloud-fake" messages should contain an encrypted field "datacenter_username" with "fakeuser@test"
    And all "network.create.vcloud-fake" messages should contain an encrypted field "datacenter_password" with "test123"
    Then an event "instance.create.vcloud-fake" should be called exactly "1" times
    And all "instance.create.vcloud-fake" messages should contain a field "_provider" with "vcloud-fake"
    And all "instance.create.vcloud-fake" messages should contain a field "vcloud_url" with "https://vcloud.net"
    And all "instance.create.vcloud-fake" messages should contain a field "datacenter_name" with "fakevcloud"
    And all "instance.create.vcloud-fake" messages should contain an encrypted field "datacenter_username" with "fakeuser@test"
    And all "instance.create.vcloud-fake" messages should contain an encrypted field "datacenter_password" with "test123"
    And all "instance.create.vcloud-fake" messages should contain a field "name" with "web-1"
    And all "instance.create.vcloud-fake" messages should contain a field "hostname" with "web-1"
    And all "instance.create.vcloud-fake" messages should contain a field "ip" with "10.1.0.11"
    And all "instance.create.vcloud-fake" messages should contain a field "network" with "web"
    And all "instance.create.vcloud-fake" messages should contain a field "cpus" with "1"
    And all "instance.create.vcloud-fake" messages should contain a field "ram" with "1024"
    And all "instance.create.vcloud-fake" messages should contain a field "reference_image" with "ubuntu-1404"
    And all "instance.create.vcloud-fake" messages should contain a field "reference_catalog" with "r3"
    And all "instance.create.vcloud-fake" messages should contain a field "_state" with "running"
