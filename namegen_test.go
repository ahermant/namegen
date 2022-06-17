package namegen

import (
    "testing"
    "regexp"
)

// Calling GenerateOne with empty lists should return an error
func TestGeneratOneEmpty(t *testing.T) {
	name, err:= GenerateOne(nil, nil)
    if name != "" || err == nil {
        t.Fatalf(`GenerateOne(nil,nil) = %q, %v, want "", error`, name, err)
    }
}

// Calling GenerateOne with proper lists should return 2 strings separated per the default separator
// Example: awesome-wonderwoman
func TestGenerateOne(t *testing.T) {
	list1:= []string{"awesome", "funny", "scary"}
	list2:= []string{"batman", "wonderwoman", "dark-vader"}
	want := regexp.MustCompile(`\b-\b`)
	name, err:= GenerateOne(list1, list2)
    if !want.MatchString(name) || err != nil {
        t.Fatalf(`GenerateOne(list1, list2) = %q, %v, want match for %#q, nil`, name, err, want)
    }
}

// Calling GenerateOne with lists and separator should return should return 2 strings separated by the separator provided
// Example: awesome#wonderwoman

// Calling GenerateOne with prefix should return the prefix string, 1 separator, the first string, 1 other separator, the last string
// Example: 20220616-awesome-wonderwoman