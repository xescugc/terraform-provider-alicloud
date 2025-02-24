package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccAlicloudRdsDBInstanceClasses_base(t *testing.T) {
	rand := acctest.RandInt()
	ZoneIDConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id": `"${data.alicloud_db_zones.default.zones.0.id}"`,
		}),
		fakeConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id": `"fake_zoneid"`,
		}),
	}
	EngineVersionConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":        `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"engine":         `"MySQL"`,
			"engine_version": `"8.0"`,
		}),
		fakeConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":        `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"engine":         `"MySQL"`,
			"engine_version": `"3.0"`,
		}),
	}

	ChargeTypeConf_Prepaid := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":              `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"instance_charge_type": `"PrePaid"`,
		}),
	}
	ChargeTypeConf_Postpaid := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":              `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"instance_charge_type": `"PostPaid"`,
		}),
	}
	StorageTypeConf_local_ssd := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":      `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"storage_type": `"local_ssd"`,
		}),
	}

	StorageTypeConf_cloud_ssd := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":      `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"storage_type": `"cloud_ssd"`,
		}),
	}
	multiZoneConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":    `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"multi_zone": `"true"`,
		}),
	}
	falseMultiZoneConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":    `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"multi_zone": `"false"`,
		}),
	}
	CategoryConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":  `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"category": `"HighAvailability"`,
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":              `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"instance_charge_type": `"PostPaid"`,
			"storage_type":         `"local_ssd"`,
			"category":             `"HighAvailability"`,
			"engine":               `"MySQL"`,
			"engine_version":       `"8.0"`,
		}),
		fakeConfig: testAccCheckAlicloudDBInstanceClassesDataSourceConfig(map[string]string{
			"zone_id":              `"${data.alicloud_db_zones.default.zones.0.id}"`,
			"instance_charge_type": `"PostPaid"`,
			"engine":               `"MySQL"`,
			"engine_version":       `"5.0"`,
		}),
	}

	var existDBInstanceMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"instance_classes.#":                    CHECKSET,
			"instance_classes.0.instance_class":     CHECKSET,
			"instance_classes.0.storage_range.min":  CHECKSET,
			"instance_classes.0.storage_range.max":  CHECKSET,
			"instance_classes.0.storage_range.step": CHECKSET,
			"instance_classes.0.zone_ids.0.id":      CHECKSET,
		}
	}

	var fakeDBInstanceMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"instance_classes.#": "0",
		}
	}

	var DBInstanceCheckInfo = dataSourceAttr{
		resourceId:   "data.alicloud_db_instance_classes.default",
		existMapFunc: existDBInstanceMapFunc,
		fakeMapFunc:  fakeDBInstanceMapFunc,
	}

	//DBInstanceCheckInfo.dataSourceTestCheck(t, rand, EngineVersionConf, prePaidSortedByConf, postPaidSortedByConf,
	//	ChargeTypeConf_Prepaid, ChargeTypeConf_Postpaid, CategoryConf, DBInstanceClassConf, multiZoneConf, falseMultiZoneConf, StorageTypeConf_local_ssd, StorageTypeConf_cloud_ssd, allConf)
	DBInstanceCheckInfo.dataSourceTestCheck(t, rand, ZoneIDConf, EngineVersionConf, ChargeTypeConf_Prepaid,
		ChargeTypeConf_Postpaid, CategoryConf, multiZoneConf, falseMultiZoneConf, StorageTypeConf_local_ssd, StorageTypeConf_cloud_ssd, allConf)
}

func testAccCheckAlicloudDBInstanceClassesDataSourceConfig(attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}
	config := fmt.Sprintf(`
data "alicloud_db_zones" "default" {
  instance_charge_type= "PostPaid"
}
data "alicloud_db_instance_classes" "default" {
  %s
}
`, strings.Join(pairs, "\n  "))
	return config
}
