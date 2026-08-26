package catalog

import (
	"reflect"
	"testing"
)

func TestExpectedCatalogDigest(t *testing.T) {
	identities := ExpectedIdentities()
	digest, err := Digest(identities)
	if err != nil {
		t.Fatalf("Digest: %v", err)
	}
	if digest != ExpectedDigest {
		t.Fatalf("Digest = %q, want %q", digest, ExpectedDigest)
	}
	if ExpectedDigest != "151a7e614abd6b25e643062b26cb9c6af60b0eadedf8e03c1f12b1850458913d" {
		t.Fatalf("ExpectedDigest drifted: %q", ExpectedDigest)
	}
}

func TestLocalSchemaDigestsMatchPinnedIdentities(t *testing.T) {
	want := identitiesByName(ExpectedIdentities())
	for name, schema := range LocalSchemas() {
		digest, err := SchemaDigest(schema)
		if err != nil {
			t.Fatalf("SchemaDigest(%q): %v", name, err)
		}
		if digest != want[name].ToolSchemaDigest {
			t.Errorf("SchemaDigest(%q) = %q, want %q", name, digest, want[name].ToolSchemaDigest)
		}
	}
}

func TestConsumedM2IdentityBytesArePinned(t *testing.T) {
	want := map[string]string{
		"relay.submit":  "6bb7bbf46d8bf5d210cee410fbd0fa59106145425878c065adf0d54b05ace08e",
		"relay.project": "be5c41ec848bd7f6a7afd16af5acc56c65cf39bc113041941bb6747153bd582a",
		"relay.read":    "a84645cb3f57ea1172661ddcc42e8a710f5a320ee3ed6c944f5e469026b3036e",
	}
	got := map[string]string{
		"relay.submit":  RelaySubmitSchemaDigest,
		"relay.project": RelayProjectSchemaDigest,
		"relay.read":    RelayReadSchemaDigest,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("consumed m-2 digests mismatch:\n got: %#v\nwant: %#v", got, want)
	}

	mutated := ExpectedIdentities()
	mutated[identityIndex(t, mutated, "relay.submit")].ToolSchemaDigest = "0" + RelaySubmitSchemaDigest[1:]
	if err := EqualIdentitySet(mutated, ExpectedIdentities()); err == nil {
		t.Fatal("EqualIdentitySet accepted a mismatched consumed m-2 digest")
	}
}

func TestVersionMarkersAndMappingVersionShape(t *testing.T) {
	if CatalogVersion != "m9-catalog-v1" {
		t.Fatalf("CatalogVersion = %q", CatalogVersion)
	}
	if MappingVersion != "m2-mapping-v1" {
		t.Fatalf("MappingVersion = %q", MappingVersion)
	}

	for _, identity := range ExpectedIdentities() {
		if identity.ToolImplCatalogVersion != CatalogVersion {
			t.Errorf("%s catalog version = %q", identity.CanonicalName, identity.ToolImplCatalogVersion)
		}
		isRelay := len(identity.CanonicalName) > len("relay.") && identity.CanonicalName[:len("relay.")] == "relay."
		if isRelay && identity.FormSchemaMappingVersion != MappingVersion {
			t.Errorf("%s mapping version = %q, want %q", identity.CanonicalName, identity.FormSchemaMappingVersion, MappingVersion)
		}
		if !isRelay && identity.FormSchemaMappingVersion != "" {
			t.Errorf("%s local mapping version must be absent, got %q", identity.CanonicalName, identity.FormSchemaMappingVersion)
		}
	}
}

func TestCanonicalSetEqualityRejectsAbsentMember(t *testing.T) {
	expected := ExpectedIdentities()
	missingTool := append([]Identity(nil), expected[1:]...)
	if err := EqualIdentitySet(expected, missingTool); err == nil {
		t.Fatal("EqualIdentitySet accepted a missing tool member")
	}

	missingMappingVersion := ExpectedIdentities()
	missingMappingVersion[identityIndex(t, missingMappingVersion, "relay.read")].FormSchemaMappingVersion = ""
	if err := EqualIdentitySet(expected, missingMappingVersion); err == nil {
		t.Fatal("EqualIdentitySet accepted an absent relay mapping-version member")
	}
}

func TestFakeM10ManifestIdentityFixtureMatchesBothDirections(t *testing.T) {
	manifestToolSet := []Identity{
		{CanonicalName: "write", ToolSchemaDigest: "0863ca49d0670920e458e08df8a8547017f99b2f09b02375d0bdf10f53e24a7b", ToolImplCatalogVersion: "m9-catalog-v1"},
		{CanonicalName: "relay.submit", ToolSchemaDigest: "6bb7bbf46d8bf5d210cee410fbd0fa59106145425878c065adf0d54b05ace08e", ToolImplCatalogVersion: "m9-catalog-v1", FormSchemaMappingVersion: "m2-mapping-v1"},
		{CanonicalName: "read", ToolSchemaDigest: "4dc4e27079b79febaa3ff2b1a91c90df7cbabd864868d4c052cee2d3b2903356", ToolImplCatalogVersion: "m9-catalog-v1"},
		{CanonicalName: "relay.read", ToolSchemaDigest: "a84645cb3f57ea1172661ddcc42e8a710f5a320ee3ed6c944f5e469026b3036e", ToolImplCatalogVersion: "m9-catalog-v1", FormSchemaMappingVersion: "m2-mapping-v1"},
		{CanonicalName: "edit", ToolSchemaDigest: "396e58a8583dc71e366be6cb516258efcc1f666810bcdebee81bcb1fdb9d6f4d", ToolImplCatalogVersion: "m9-catalog-v1"},
		{CanonicalName: "bash", ToolSchemaDigest: "ddd9efb1ca815ee8a71ffda66eefbd627f1b4be32c9955a23eb874ea9cba5271", ToolImplCatalogVersion: "m9-catalog-v1"},
		{CanonicalName: "relay.project", ToolSchemaDigest: "be5c41ec848bd7f6a7afd16af5acc56c65cf39bc113041941bb6747153bd582a", ToolImplCatalogVersion: "m9-catalog-v1", FormSchemaMappingVersion: "m2-mapping-v1"},
		{CanonicalName: "apply_patch", ToolSchemaDigest: "f6594e18aaedfb029106fa669c557027854ec5f86cce436fcec1723791743cd7", ToolImplCatalogVersion: "m9-catalog-v1"},
	}

	if err := EqualIdentitySet(ExpectedIdentities(), manifestToolSet); err != nil {
		t.Fatalf("worker to manifest equality: %v", err)
	}
	if err := EqualIdentitySet(manifestToolSet, ExpectedIdentities()); err != nil {
		t.Fatalf("manifest to worker equality: %v", err)
	}
	digest, err := Digest(manifestToolSet)
	if err != nil {
		t.Fatalf("Digest(manifestToolSet): %v", err)
	}
	if digest != ExpectedDigest {
		t.Fatalf("manifest digest = %q, want %q", digest, ExpectedDigest)
	}
}

func identitiesByName(identities []Identity) map[string]Identity {
	result := make(map[string]Identity, len(identities))
	for _, identity := range identities {
		result[identity.CanonicalName] = identity
	}
	return result
}

func identityIndex(t *testing.T, identities []Identity, name string) int {
	t.Helper()
	for index, identity := range identities {
		if identity.CanonicalName == name {
			return index
		}
	}
	t.Fatalf("identity %q not found", name)
	return -1
}
