package parsers

import (
	"testing"
)

const singleOutput string =
`1485520026,,ui,message,    builder_aws: AMI: ami-41be5426
1485520026,,ui,say,==> builder_aws: Waiting for AMI to become ready...
1485520158,,ui,say,==> builder_aws: Terminating the source AWS instance...
1485520165,,ui,say,==> builder_aws: Cleaning up any extra volumes...
1485520165,,ui,say,==> builder_aws: No volumes to clean up%!(PACKER_COMMA) skipping
1485520165,,ui,say,==> builder_aws: Deleting temporary security group...
1485520166,,ui,say,==> builder_aws: Deleting temporary keypair...
1485520166,,ui,say,Build 'builder_aws' finished.
1485520166,,ui,say,\n==> Builds finished. The artifacts of successful builds are:
1485520166,builder_aws,artifact-count,1
1485520166,builder_aws,artifact,0,builder-id,mitchellh.amazonebs
1485520166,builder_aws,artifact,0,id,eu-west-1:ami-41be5426
1485520166,builder_aws,artifact,0,string,AMIs were created:\n\neu-west-1: ami-41be5426
1485520166,builder_aws,artifact,0,files-count,0
1485520166,builder_aws,artifact,0,end
1485520166,,ui,say,--> builder_aws: AMIs were created:\n\neu-west-1: ami-41be5426`

const multiOutput string =
`1484841730,,ui,message,    builder1_aws: Creating configuration file 'knife.rb'
1484841730,,ui,message,    builder1_aws: Removing directory: /tmp/packer-chef-client
1484841730,,ui,say,==> builder1_aws: Stopping the source instance...
1484841731,,ui,say,==> builder1_aws: Waiting for the instance to stop...
1484841762,,ui,say,==> builder1_aws: Creating the AMI: builder1 aws 1484840807
1484841763,,ui,message,    builder1_aws: AMI: ami-ne641t7t
1484841763,,ui,say,==> builder1_aws: Waiting for AMI to become ready...
1484841895,,ui,say,==> builder1_aws: Terminating the source AWS instance...
1484841902,,ui,say,==> builder1_aws: Cleaning up any extra volumes...
1484841904,,ui,say,==> builder1_aws: No volumes to clean up%!(PACKER_COMMA) skipping
1484841904,,ui,say,==> builder1_aws: Deleting temporary security group...
1484841904,,ui,say,==> builder1_aws: Deleting temporary keypair...
1484841904,,ui,say,Build 'builder1_aws' finished.
1484841904,,ui,say,\n==> Builds finished. The artifacts of successful builds are:
1484841904,builder2_aws,artifact-count,1
1484841904,builder2_aws,artifact,0,builder-id,mitchellh.amazonebs
1484841904,builder2_aws,artifact,0,id,eu-west-1:ami-dc4f78af
1484841904,builder2_aws,artifact,0,string,AMIs were created:\n\neu-west-1: ami-dc4f78af
1484841904,builder2_aws,artifact,0,files-count,0
1484841904,builder2_aws,artifact,0,end
1484841904,,ui,say,--> builder2_aws: AMIs were created:\n\neu-west-1: ami-dc4f78af
1484841904,builder1_aws,artifact-count,1
1484841904,builder1_aws,artifact,0,builder-id,mitchellh.amazonebs
1484841904,builder1_aws,artifact,0,id,eu-west-1:ami-ne641t7t
1484841904,builder1_aws,artifact,0,string,AMIs were created:\n\neu-west-1: ami-ne641t7t
1484841904,builder1_aws,artifact,0,files-count,0
1484841904,builder1_aws,artifact,0,end
1484841904,,ui,say,--> builder1_aws: AMIs were created:\n\neu-west-1: ami-ne641t7t`

func TestSingleArtifact(t *testing.T) {
	output := []byte(singleOutput)

	artifacts := ParseLines(output)

	builder := "builder_aws"

	if artifacts == nil || len(artifacts) != 1 {
		t.Error("expected 1 builder to produce artifacts artifacts is:", artifacts)
	}

	if artifacts[builder] == nil || len(artifacts[builder]) != 1 {
		t.Error("expected builder", builder, "to create 1 artifact, but found:", artifacts["builder_aws"])
	}

	if artifacts[builder][0].Builder != builder ||
			artifacts[builder][0].Id != "ami-41be5426"  ||
			artifacts[builder][0].Region != "eu-west-1" {
		t.Error("expected artifact to be eu-west-1 ami-41be5426, but it was", artifacts[builder][0])
	}
}

func TestMultiArtifact(t *testing.T) {
	output := []byte(multiOutput)

	artifacts := ParseLines(output)

	if artifacts == nil || len(artifacts) != 2 {
		t.Error("expected artifacts from 2 builders to be produced but artifacts is:", artifacts)
	}

	expectedArtifacts := []Artifact{
		{Builder:"builder1_aws", Region:"eu-west-1", Id:"ami-ne641t7t"},
		{Builder:"builder2_aws", Region:"eu-west-1", Id:"ami-dc4f78af"}}

	for _, artifact := range expectedArtifacts {
		expectedBuilder := artifact.Builder
		if artifacts[expectedBuilder] == nil || len(artifacts[expectedBuilder]) != 1 {
			t.Error("expected builder ", expectedBuilder, "to create 1 artifact but it was", artifacts)
		}

		actualArtifact := artifacts[expectedBuilder][0]
		if actualArtifact.Builder != artifact.Builder ||
				actualArtifact.Region != artifact.Region ||
				actualArtifact.Id != artifact.Id {
			t.Error("expected artifact to be", artifact, "but it was", actualArtifact)
		}
	}
}
