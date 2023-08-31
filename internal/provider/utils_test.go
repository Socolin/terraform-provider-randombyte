package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func testExtractResourceAttr(resourceName string, attributeName string, attributeValue *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("resource name %s not found in state", resourceName)
		}

		attrValue, ok := rs.Primary.Attributes[attributeName]

		if !ok {
			return fmt.Errorf("attribute %s not found in resource %s state", attributeName, resourceName)
		}

		*attributeValue = attrValue

		return nil
	}
}

func testCheckAttributeValuesDiffer(i *string, j *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testStringValue(i) == testStringValue(j) {
			return fmt.Errorf("attribute values are the same")
		}

		return nil
	}
}

func testCheckAttributeValuesEqual(i *string, j *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testStringValue(i) != testStringValue(j) {
			return fmt.Errorf("attribute values are different, got %s and %s", testStringValue(i), testStringValue(j))
		}

		return nil
	}
}

func testStringValue(sPtr *string) string {
	if sPtr == nil {
		return ""
	}

	return *sPtr
}

func testCheckLen(expectedLen int) func(input string) error {
	return func(input string) error {
		if len(input) != expectedLen {
			return fmt.Errorf("expected length %d, actual length %d", expectedLen, len(input))
		}

		return nil
	}
}
