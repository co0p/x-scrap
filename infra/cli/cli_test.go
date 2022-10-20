package cli_test

import (
	"testing"

	"github.com/co0p/x-scrap/infra/cli"
)

func Test_Execute_MissingUrl(t *testing.T) {

	givenArgs := []string{"programname", "-tags", "peter"}

	sut := cli.CLI{}
	_, err := sut.Execute(givenArgs)

	if err == nil {
		t.Errorf("expected err not to be nil, got %s\n", err)
	}
}

func Test_Execute_MissingTags(t *testing.T) {

	givenArgs := []string{"programname", "-urls", "https://www.google.de"}

	sut := cli.CLI{}
	_, err := sut.Execute(givenArgs)

	if err == nil {
		t.Errorf("expected err not to be nil, got %s\n", err)
	}
}
func Test_Execute_OneTag_OneUrl(t *testing.T) {

	givenArgs := []string{"programname", "-tags", "peter", "-urls", "https://www.google.de"}

	sut := cli.CLI{}
	cmd, err := sut.Execute(givenArgs)

	if err != nil {
		t.Errorf("expected err to be nil, got %s\n", err)
	}

	if len(cmd.Tags) != 1 && cmd.Tags[0] != "peter" {
		t.Errorf("expected tags to have length 1 and equal to 'peter', got %s\n", cmd.Tags)
	}

	if len(cmd.URLs) != 1 && cmd.URLs[0] != "https://www.google.de" {
		t.Errorf("expected urls to have length 1 and equal to 'https://www.google.de', got %s\n", cmd.URLs)
	}
}

func Test_Execute_MultipleTags_MultipleUrls(t *testing.T) {

	givenArgs := []string{"programname", "-tags", "peter, parker", "-urls", "https://www.google.de,https://www.yahoo.de"}

	sut := cli.CLI{}
	cmd, err := sut.Execute(givenArgs)

	if err != nil {
		t.Errorf("expected err to be nil, got %s\n", err)
	}

	if len(cmd.Tags) != 2 {
		t.Errorf("expected tags to have length 2, got %d\n", len(cmd.Tags))
	}
	if cmd.Tags[0] != "peter" && cmd.Tags[1] != "parker" {
		t.Errorf("expected 'peter', 'parker', got %v\n", cmd.Tags)
	}
	if len(cmd.URLs) != 2 {
		t.Errorf("expected URLs to have length 2, got %d\n", len(cmd.URLs))
	}
	if cmd.URLs[0] != "https://www.google.de" && cmd.URLs[1] != "https://www.yahoo.de" {
		t.Errorf("expected 'https://www.google.de', 'https://www.yahoo.de', got %v\n", cmd.URLs)
	}
}
