package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceLexBotAlias(t *testing.T) {
	resourceName := "aws_lex_bot_alias.test"
	testId := acctest.RandStringFromCharSet(8, acctest.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testDataSourceLexBotAliasConfig, testId),
				Check: resource.ComposeTestCheckFunc(
					checkResourceStateComputedAttr(resourceName, dataSourceAwsLexBotAlias()),
				),
			},
		},
	})
}

const testDataSourceLexBotAliasConfig = `
resource "aws_lex_intent" "test" {
  fulfillment_activity {
    type = "ReturnIntent"
  }

  name = "test_intent_%[1]s"
}

resource "aws_lex_bot" "test" {
  abort_statement {
    message {
      content      = "Sorry, I am not able to assist at this time"
      content_type = "PlainText"
    }
  }

  child_directed = false

  clarification_prompt {
    max_attempts = 2

    message {
      content      = "I didn't understand you, what would you like to do?"
      content_type = "PlainText"
    }
  }

  name = "test_bot_%[1]s"

  intent {
    intent_name    = "${aws_lex_intent.test.name}"
    intent_version = "${aws_lex_intent.test.version}"
  }
}

resource "aws_lex_bot_alias" "test" {
  bot_name    = "${aws_lex_bot.test.name}"
  bot_version = "${aws_lex_bot.test.version}"
  name        = "test_bot_alias_%[1]s"
}

data "aws_lex_bot_alias" "test" {
  bot_name = "${aws_lex_bot.test.name}"
  name     = "${aws_lex_bot_alias.test.name}"
}
`
